package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_addUnicodeChar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
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
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l0 == int32(0) {
		v14 = int32(0)
		v15 = F_errsave_start(m, l1)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v110 = v14
				m.G0 = v10 + int32(32)
				return v110
			} else {
				F_errcode(m, int32(84017282))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_addUnicodeChar_0), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v30 = F_errdetail(m, int32(_a_F_addUnicodeChar_1), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l1, int32(_a_F_addUnicodeChar_2), int32(593), int32(_a_F_addUnicodeChar_3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v110 = v14
								m.G0 = v10 + int32(32)
								return v110
							}
						}
					}
				}
			}
		}
	} else {
		if l1 != 0 {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v37 == int32(453) {
				v42 = F_pg_unicode_to_server_noerror(m, l0, v10)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					if v42 != 0 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
						v63 = F_strlen(m, v10)
						mBase = m.M
						v65 = v63 + int32(1)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
						if v67 <= v62+v65 {
							v71 = v67
							v74 = v61 + int32(8)
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = v71 << (uint(int32(1)) % 32)
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
								if v84 <= v85+v65 {
									v71 = v84
									v74 = v81 + int32(8)
									continue
								} else {
									break
								}
								break
							}
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
							v89 = F_repalloc(m, v88, v84)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*int32)(unsafe.Add(mBase, uint32(v91))) = v89
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
								v96 = v93
								v98 = v94
								if v63 != 0 {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
									base.MemoryCopy(m, v102+v98, v10, v63)
								} else {
								}
								v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v106 + v63
								v110 = int32(1)
								m.G0 = v10 + int32(32)
								return v110
							}
						} else {
							v96 = v61
							v98 = v62
							if v63 != 0 {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								base.MemoryCopy(m, v102+v98, v10, v63)
							} else {
							}
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v106 + v63
							v110 = int32(1)
							m.G0 = v10 + int32(32)
							return v110
						}
					} else {
						v44 = int32(0)
						v45 = F_errsave_start(m, l1)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if v45 == int32(0) {
								v110 = v44
								m.G0 = v10 + int32(32)
								return v110
							} else {
								F_errcode(m, int32(16801924))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_addUnicodeChar_4), int32(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, l1, int32(_a_F_addUnicodeChar_2), int32(610), int32(_a_F_addUnicodeChar_3))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v110 = v44
											m.G0 = v10 + int32(32)
											return v110
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_pg_unicode_to_server(m, l0, v10)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
					v63 = F_strlen(m, v10)
					mBase = m.M
					v65 = v63 + int32(1)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
					if v67 <= v62+v65 {
						v71 = v67
						v74 = v61 + int32(8)
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v71 << (uint(int32(1)) % 32)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
							if v84 <= v85+v65 {
								v71 = v84
								v74 = v81 + int32(8)
								continue
							} else {
								break
							}
							break
						}
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
						v89 = F_repalloc(m, v88, v84)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*int32)(unsafe.Add(mBase, uint32(v91))) = v89
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
							v96 = v93
							v98 = v94
							if v63 != 0 {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								base.MemoryCopy(m, v102+v98, v10, v63)
							} else {
							}
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v106 + v63
							v110 = int32(1)
							m.G0 = v10 + int32(32)
							return v110
						}
					} else {
						v96 = v61
						v98 = v62
						if v63 != 0 {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							base.MemoryCopy(m, v102+v98, v10, v63)
						} else {
						}
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v106 + v63
						v110 = int32(1)
						m.G0 = v10 + int32(32)
						return v110
					}
				}
			}
		} else {
			F_pg_unicode_to_server(m, l0, v10)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				v63 = F_strlen(m, v10)
				mBase = m.M
				v65 = v63 + int32(1)
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
				if v67 <= v62+v65 {
					v71 = v67
					v74 = v61 + int32(8)
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = v71 << (uint(int32(1)) % 32)
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
						if v84 <= v85+v65 {
							v71 = v84
							v74 = v81 + int32(8)
							continue
						} else {
							break
						}
						break
					}
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
					v89 = F_repalloc(m, v88, v84)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						*(*int32)(unsafe.Add(mBase, uint32(v91))) = v89
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
						v96 = v93
						v98 = v94
						if v63 != 0 {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							base.MemoryCopy(m, v102+v98, v10, v63)
						} else {
						}
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v106 + v63
						v110 = int32(1)
						m.G0 = v10 + int32(32)
						return v110
					}
				} else {
					v96 = v61
					v98 = v62
					if v63 != 0 {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
						base.MemoryCopy(m, v102+v98, v10, v63)
					} else {
					}
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v106 + v63
					v110 = int32(1)
					m.G0 = v10 + int32(32)
					return v110
				}
			}
		}
	}
}
func F_unicode_is_normalized(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v165 int32
	_ = v165
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v527 int32
	_ = v527
	var v532 int64
	_ = v532
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v24 = F_pg_detoast_datum_packed(m, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v26 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v32 == int32(18) {
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
	if v26&v43 != 0 {
		v55 = int32(base.Ui32(v26)>>(uint(v43)%32)) - v43
		goto L3
	} else {
		goto L15
	}
L9:
	;
	v35 = int32(16)
	goto L11
L10:
	;
	v35 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = int32(4)
	goto L14
L13:
	;
	v42 = v35
	goto L14
L14:
	;
	v55 = v42
	goto L3
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L16:
	;
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v60 = int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v62&v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55+v58))) = uint8(v69)
	if v24 != v22 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v65 = v60
	goto L22
L21:
	;
	v65 = int32(4)
	goto L22
L22:
	;
	base.MemoryCopy(m, v58, v24+v65, v55)
	goto L19
L23:
	;
	F_pfree(m, v24)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v74 = F_unicode_norm_form_from_string(m, v58)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v76 = int32(4)
	v77 = int32(1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v81 = v79 & v77
	if v81 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v82 = v77
	goto L30
L29:
	;
	v82 = v76
	goto L30
L30:
	;
	if v79 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v111 = F_pg_mbstrlen_with_len(m, v17+v82, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L42
	}
L32:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v89 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v100 = int32(1)
	if v81 != 0 {
		v110 = int32(base.Ui32(v79)>>(uint(v100)%32)) - v100
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v92 = int32(16)
	goto L37
L36:
	;
	v92 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v99 = int32(4)
	goto L40
L39:
	;
	v99 = v92
	goto L40
L40:
	;
	v110 = v99
	goto L31
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	v115 = F_palloc_mul(m, v76, v111+int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v111 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v117 = int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v119&v117 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115+v111<<(uint(int32(2))%32)))) = int32(0)
	v254 = int32(1)
	if v74&int32(-3) == v254 {
		v438 = int32(-1)
		goto L79
	} else {
		goto L80
	}
L47:
	;
	v122 = v117
	goto L49
L48:
	;
	v122 = int32(4)
	goto L49
L49:
	;
	v125 = int32(0)
	v126 = v17 + v122
	goto L50
L50:
	;
	v143 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
	v145 = v143 & int32(255)
	if int32(0) <= v143 {
		v202 = v145
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115+v125<<(uint(int32(2))%32)))) = v202
	v204 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
	if int32(0) <= v204 {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if v145&int32(224) == int32(192) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v126))))
	v202 = v197&int32(63) | v194
	goto L52
L55:
	;
	v194 = v145 << (uint(int32(6)) % 32) & int32(1984)
	v195 = int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	if v145&int32(240) == int32(224) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	v194 = v145<<(uint(int32(12))%32)&int32(_a_F_unicode_is_normalized_0) | v165&int32(63)<<(uint(int32(6))%32)
	v195 = int32(2)
	goto L54
L59:
	;
	goto L60
L60:
	;
	if v145&int32(248) != int32(240) {
		v202 = int32(-1)
		goto L52
	} else {
		goto L61
	}
L61:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	v182 = int32(63)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
	v194 = v145<<(uint(int32(18))%32)&int32(_a_F_unicode_is_normalized_1) | v181&v182<<(uint(int32(12))%32) | v187&v182<<(uint(int32(6))%32)
	v195 = int32(3)
	goto L54
L62:
	;
	v231 = v125 + int32(1)
	if v231 != v111 {
		v125 = v231
		v126 = v228 + v126
		goto L50
	} else {
		goto L75
	}
L63:
	;
	v228 = int32(1)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v209 = v204 & int32(255)
	if v209&int32(224) == int32(192) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v228 = int32(2)
	goto L62
L67:
	;
	goto L68
L68:
	;
	if v209&int32(240) == int32(224) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v228 = int32(3)
	goto L62
L70:
	;
	goto L71
L71:
	;
	if v209&int32(248) == int32(240) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v226 = int32(4)
	goto L74
L73:
	;
	v226 = int32(1)
	goto L74
L74:
	;
	v228 = v226
	goto L62
L75:
	;
	goto L51
L76:
	;
	return v532
L77:
	;
	v441 = F_unicode_normalize(m, v74, v115)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L99
	}
L78:
	;
	return int64(0)
L79:
	;
	switch v438 {
	case 0:
		goto L78
	case 1:
		v532 = int64(1)
		goto L76
	default:
		goto L77
	}
L80:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v261 == int32(0) {
		v438 = int32(1)
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v266 = v254
	v268 = v115
	v269 = v261
	v277 = int32(0)
	goto L82
L82:
	;
	v279 = int32(0)
	v280 = int32(16711935)
	v282 = int32(8)
	v283 = base.I32_rotr(v269&v280, v282)
	v284 = int32(24)
	v285 = base.I32_rotr(v269, v284)
	v287 = int32(255)
	v288 = (v283 | v285) & v287
	v289 = int32(127)
	v294 = int32(base.Ui32(v283)>>(uint(v282)%32)) & v287
	v301 = int32(base.Ui32(v285&v280) >> (uint(int32(16)) % 32))
	v306 = int32(base.Ui32(v283) >> (uint(v284) % 32))
	v310 = int32(_a_F_unicode_is_normalized_2)
	v311 = base.I32_rem_u_s(((v288*v289+v294)*v289+v301)*v289+v306+int32(260144641), v310)
	v312 = int32(1)
	v314 = int32(*(*int16)(unsafe.Add(mBase, uint32(v311<<(uint(v312)%32))+uint32(_c_F_unicode_is_normalized[0]))))
	v315 = int32(257)
	v323 = ((v288*v315+v294)*v315+v301)*v315 + v306
	v325 = base.I32_rem_u_s(v323, v310)
	v328 = int32(*(*int16)(unsafe.Add(mBase, uint32(v325<<(uint(v312)%32))+uint32(_c_F_unicode_is_normalized[0]))))
	v329 = v314 + v328
	if base.Ui32(int32(_a_F_unicode_is_normalized_3)) < base.Ui32(v329) {
		v346 = v279
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v438 = v418
	goto L79
L84:
	;
	switch v74 {
	case 0:
		goto L93
	default:
		goto L90
	case 2:
		goto L92
	}
L85:
	;
	v333 = v329 << (uint(int32(3)) % 32)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_c_F_unicode_is_normalized[1])))
	if v269 != v334 {
		v346 = v279
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+uint32(_c_F_unicode_is_normalized[2]))))
	if base.B2i32(base.Ui32(v277&int32(255)) <= base.Ui32(v338))|base.B2i32(v338 == int32(0)) != 0 {
		v346 = v338
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v438 = int32(0)
	goto L79
L88:
	;
	goto L83
L89:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v415 != 0 {
		v266 = v414
		v268 = v268 + int32(4)
		v269 = v415
		v277 = v346
		goto L82
	} else {
		goto L98
	}
L90:
	;
	v414 = v266
	goto L89
L91:
	;
	v410 = v406 << (uint(int32(7)) % 32) >> (uint(int32(28)) % 32)
	switch v410 + int32(1) {
	case 0:
		v414 = v410
		goto L89
	case 1:
		v418 = v410
		goto L88
	default:
		goto L90
	}
L92:
	;
	v376 = int32(_a_F_unicode_is_normalized_4)
	v387 = int32(_a_F_unicode_is_normalized_5)
	v388 = base.I32_rem_u_s(((v288*v376+v294)*v376+v301)*v376+v306+int32(402620417), v387)
	v389 = int32(1)
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v388<<(uint(v389)%32))+uint32(_c_F_unicode_is_normalized[3]))))
	v393 = base.I32_rem_u_s(v323, v387)
	v396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v393<<(uint(v389)%32))+uint32(_c_F_unicode_is_normalized[3]))))
	v397 = v391 + v396
	if base.Ui32(int32(_a_F_unicode_is_normalized_6)) < base.Ui32(v397) {
		goto L90
	} else {
		goto L96
	}
L93:
	;
	v348 = int32(_a_F_unicode_is_normalized_4)
	v357 = int32(2505)
	v358 = base.I32_rem_u_s(((v288*v348+v294)*v348+v301)*v348+v306, v357)
	v359 = int32(1)
	v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v358<<(uint(v359)%32))+uint32(_c_F_unicode_is_normalized[4]))))
	v363 = base.I32_rem_u_s(v323, v357)
	v366 = int32(*(*int16)(unsafe.Add(mBase, uint32(v363<<(uint(v359)%32))+uint32(_c_F_unicode_is_normalized[4]))))
	v367 = v361 + v366
	if base.Ui32(int32(1251)) < base.Ui32(v367) {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v367<<(uint(int32(2))%32))+uint32(_c_F_unicode_is_normalized[5])))
	if v269 == v372&int32(_a_F_unicode_is_normalized_7) {
		v406 = v372
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L90
L96:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v397<<(uint(int32(2))%32))+uint32(_c_F_unicode_is_normalized[6])))
	if v269 != v402&int32(_a_F_unicode_is_normalized_7) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v406 = v402
	goto L91
L98:
	;
	v418 = v414
	goto L88
L99:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	if v443 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v445 = v441 + int32(4)
	v449 = v445
	goto L104
L101:
	;
	v460 = int32(0)
	goto L102
L102:
	;
	if v460 != v111 {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v460 = (v449-v445)>>(uint(int32(2))%32) + int32(1)
	goto L102
L104:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	if v453 != 0 {
		v449 = v449 + int32(4)
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	goto L105
L107:
	;
	return int64(0)
L108:
	;
	goto L109
L109:
	;
	v465 = v111 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v465) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v532 = base.I64_extend_i32_u(base.B2i32(v527 == int32(0)))
	goto L76
L111:
	;
	v527 = int32(0)
	goto L110
L112:
	;
	v501 = v496
	v502 = v497
	v503 = v498
	goto L122
L113:
	;
	if (v115|v441)&int32(3) != 0 {
		v496 = v115
		v497 = v441
		v498 = v465
		goto L112
	} else {
		goto L116
	}
L114:
	;
	v489 = v115
	v490 = v441
	v491 = v465
	goto L115
L115:
	;
	if v491 == int32(0) {
		goto L111
	} else {
		goto L121
	}
L116:
	;
	v473 = v115
	v474 = v441
	v475 = v465
	goto L117
L117:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	if v478 != v479 {
		v496 = v473
		v497 = v474
		v498 = v475
		goto L112
	} else {
		goto L119
	}
L118:
	;
	v489 = v484
	v490 = v482
	v491 = v486
	goto L115
L119:
	;
	v481 = int32(4)
	v482 = v474 + v481
	v484 = v473 + v481
	v486 = v475 - v481
	if base.Ui32(int32(3)) < base.Ui32(v486) {
		v473 = v484
		v474 = v482
		v475 = v486
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v496 = v489
	v497 = v490
	v498 = v491
	goto L112
L122:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	if v506 == v507 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v527 = v506 - v507
	goto L110
L124:
	;
	v509 = int32(1)
	v514 = v503 - v509
	if v514 != 0 {
		v501 = v501 + v509
		v502 = v502 + v509
		v503 = v514
		goto L122
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	goto L123
L127:
	;
	goto L111
}
func F_unicode_normalize_func(m *base.Module, l0 int32) int64 {
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v158 int32
	_ = v158
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				if v24 == int32(1) {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
					if v30 == int32(18) {
						v33 = int32(16)
					} else {
						v33 = int32(0)
					}
					if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v40 = int32(4)
					} else {
						v40 = v33
					}
					v53 = v40
				} else {
					v41 = int32(1)
					if v24&v41 != 0 {
						v53 = int32(base.Ui32(v24)>>(uint(v41)%32)) - v41
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v56 = F_palloc(m, v53+int32(1))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					if v53 != 0 {
						v58 = int32(1)
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
						if v60&v58 != 0 {
							v63 = v58
						} else {
							v63 = int32(4)
						}
						base.MemoryCopy(m, v56, v22+v63, v53)
					} else {
					}
					v67 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v56+v53))) = uint8(v67)
					if v22 != v20 {
						F_pfree(m, v22)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							v72 = F_unicode_norm_form_from_string(m, v56)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								v74 = int32(4)
								v76 = int32(1)
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
								v80 = v78 & v76
								if v80 != 0 {
									v81 = v76
								} else {
									v81 = v74
								}
								if v78 == int32(1) {
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
									if v88 == int32(18) {
										v91 = int32(16)
									} else {
										v91 = int32(0)
									}
									if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v98 = int32(4)
									} else {
										v98 = v91
									}
									v109 = v98
								} else {
									v99 = int32(1)
									if v80 != 0 {
										v109 = int32(base.Ui32(v78)>>(uint(v99)%32)) - v99
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
										v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v110 = F_pg_mbstrlen_with_len(m, v15+v81, v109)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int64(0)
								} else {
									v114 = F_palloc_mul(m, v74, v110+int32(1))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int64(0)
									} else {
										if v110 != 0 {
											v116 = int32(1)
											v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
											if v118&v116 != 0 {
												v121 = v116
											} else {
												v121 = int32(4)
											}
											v124 = int32(0)
											v125 = v15 + v121
											for {
												v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125))))
												v138 = v136 & int32(255)
												if int32(0) <= v136 {
													v195 = v138
												} else {
													if v138&int32(224) == int32(192) {
														v187 = v138 << (uint(int32(6)) % 32) & int32(1984)
														v188 = int32(1)
														v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v125))))
														v195 = v190&int32(63) | v187
													} else {
														if v138&int32(240) == int32(224) {
															v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
															v187 = v138<<(uint(int32(12))%32)&int32(_a_F_unicode_normalize_func_0) | v158&int32(63)<<(uint(int32(6))%32)
															v188 = int32(2)
															v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v125))))
															v195 = v190&int32(63) | v187
														} else {
															if v138&int32(248) != int32(240) {
																v195 = int32(-1)
															} else {
																v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
																v175 = int32(63)
																v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+2)))
																v187 = v138<<(uint(int32(18))%32)&int32(_a_F_unicode_normalize_func_1) | v174&v175<<(uint(int32(12))%32) | v180&v175<<(uint(int32(6))%32)
																v188 = int32(3)
																v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v125))))
																v195 = v190&int32(63) | v187
															}
														}
													}
												}
												*(*int32)(unsafe.Add(mBase, uint32(v114+v124<<(uint(int32(2))%32)))) = v195
												v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125))))
												if int32(0) <= v197 {
													v221 = int32(1)
												} else {
													v202 = v197 & int32(255)
													if v202&int32(224) == int32(192) {
														v221 = int32(2)
													} else {
														if v202&int32(240) == int32(224) {
															v221 = int32(3)
														} else {
															if v202&int32(248) == int32(240) {
																v219 = int32(4)
															} else {
																v219 = int32(1)
															}
															v221 = v219
														}
													}
												}
												v224 = v124 + int32(1)
												if v224 != v110 {
													v124 = v224
													v125 = v221 + v125
													continue
												} else {
													break
												}
												break
											}
										} else {
										}
										v235 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v114+v110<<(uint(int32(2))%32)))) = v235
										v241 = F_unicode_normalize(m, v72, v114)
										mBase = m.M
										v242 = m.ExcPending
										if v242 != 0 {
											return int64(0)
										} else {
											v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
											if v243 != 0 {
												v244 = v241
												v245 = v243
												v248 = v235
												for {
													if base.Ui32(v245) <= base.Ui32(int32(127)) {
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v245)
													} else {
														if base.Ui32(v245) <= base.Ui32(int32(2047)) {
															v261 = v245&int32(63) | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v261)
															v266 = int32(base.Ui32(v245)>>(uint(int32(6))%32)) | int32(192)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v266)
														} else {
															if base.Ui32(v245) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
																v270 = int32(63)
																v272 = int32(128)
																v273 = v245&v270 | v272
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v273)
																v278 = int32(base.Ui32(v245)>>(uint(int32(12))%32)) | int32(224)
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v278)
																v285 = int32(base.Ui32(v245)>>(uint(int32(6))%32))&v270 | v272
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v285)
															} else {
																v287 = int32(63)
																v289 = int32(128)
																v290 = v245&v287 | v289
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v290)
																v297 = int32(base.Ui32(v245)>>(uint(int32(6))%32))&v287 | v289
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v297)
																v304 = int32(base.Ui32(v245)>>(uint(int32(12))%32))&v287 | v289
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v304)
																v311 = int32(base.Ui32(v245)>>(uint(int32(18))%32))&int32(7) | int32(240)
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v311)
															}
														}
													}
													v317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
													if int32(0) <= v317 {
														v341 = int32(1)
													} else {
														v322 = v317 & int32(255)
														if v322&int32(224) == int32(192) {
															v341 = int32(2)
														} else {
															if v322&int32(240) == int32(224) {
																v341 = int32(3)
															} else {
																if v322&int32(248) == int32(240) {
																	v339 = int32(4)
																} else {
																	v339 = int32(1)
																}
																v341 = v339
															}
														}
													}
													v342 = v341 + v248
													v343 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
													if v343 != 0 {
														v244 = v244 + int32(4)
														v245 = v343
														v248 = v342
														continue
													} else {
														break
													}
													break
												}
												v349 = v342 + int32(4)
											} else {
												v349 = v74
											}
											v355 = F_palloc(m, v349)
											mBase = m.M
											v356 = m.ExcPending
											if v356 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v355))) = v349 << (uint(int32(2)) % 32)
												v360 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
												if v360 != 0 {
													v363 = v360
													v364 = v355 + int32(4)
													v365 = v241
													for {
														if base.Ui32(v363) <= base.Ui32(int32(127)) {
															*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v363)
														} else {
															if base.Ui32(v363) <= base.Ui32(int32(2047)) {
																v380 = v363&int32(63) | int32(128)
																*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)) = uint8(v380)
																v385 = int32(base.Ui32(v363)>>(uint(int32(6))%32)) | int32(192)
																*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v385)
															} else {
																if base.Ui32(v363) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
																	v389 = int32(63)
																	v391 = int32(128)
																	v392 = v363&v389 | v391
																	*(*uint8)(unsafe.Add(mBase, uint32(v364)+2)) = uint8(v392)
																	v397 = int32(base.Ui32(v363)>>(uint(int32(12))%32)) | int32(224)
																	*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v397)
																	v404 = int32(base.Ui32(v363)>>(uint(int32(6))%32))&v389 | v391
																	*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)) = uint8(v404)
																} else {
																	v406 = int32(63)
																	v408 = int32(128)
																	v409 = v363&v406 | v408
																	*(*uint8)(unsafe.Add(mBase, uint32(v364)+3)) = uint8(v409)
																	v416 = int32(base.Ui32(v363)>>(uint(int32(6))%32))&v406 | v408
																	*(*uint8)(unsafe.Add(mBase, uint32(v364)+2)) = uint8(v416)
																	v423 = int32(base.Ui32(v363)>>(uint(int32(12))%32))&v406 | v408
																	*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)) = uint8(v423)
																	v430 = int32(base.Ui32(v363)>>(uint(int32(18))%32))&int32(7) | int32(240)
																	*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v430)
																}
															}
														}
														v434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v364))))
														if int32(0) <= v434 {
															v458 = int32(1)
														} else {
															v439 = v434 & int32(255)
															if v439&int32(224) == int32(192) {
																v458 = int32(2)
															} else {
																if v439&int32(240) == int32(224) {
																	v458 = int32(3)
																} else {
																	if v439&int32(248) == int32(240) {
																		v456 = int32(4)
																	} else {
																		v456 = int32(1)
																	}
																	v458 = v456
																}
															}
														}
														v460 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
														if v460 != 0 {
															v363 = v460
															v364 = v458 + v364
															v365 = v365 + int32(4)
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												m.G0 = v12 + int32(16)
												return base.I64_extend_i32_u(v355)
											}
										}
									}
								}
							}
						}
					} else {
						v72 = F_unicode_norm_form_from_string(m, v56)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							v74 = int32(4)
							v76 = int32(1)
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
							v80 = v78 & v76
							if v80 != 0 {
								v81 = v76
							} else {
								v81 = v74
							}
							if v78 == int32(1) {
								v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
								if v88 == int32(18) {
									v91 = int32(16)
								} else {
									v91 = int32(0)
								}
								if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v98 = int32(4)
								} else {
									v98 = v91
								}
								v109 = v98
							} else {
								v99 = int32(1)
								if v80 != 0 {
									v109 = int32(base.Ui32(v78)>>(uint(v99)%32)) - v99
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
									v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v110 = F_pg_mbstrlen_with_len(m, v15+v81, v109)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int64(0)
							} else {
								v114 = F_palloc_mul(m, v74, v110+int32(1))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int64(0)
								} else {
									if v110 != 0 {
										v116 = int32(1)
										v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
										if v118&v116 != 0 {
											v121 = v116
										} else {
											v121 = int32(4)
										}
										v124 = int32(0)
										v125 = v15 + v121
										for {
											v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125))))
											v138 = v136 & int32(255)
											if int32(0) <= v136 {
												v195 = v138
											} else {
												if v138&int32(224) == int32(192) {
													v187 = v138 << (uint(int32(6)) % 32) & int32(1984)
													v188 = int32(1)
													v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v125))))
													v195 = v190&int32(63) | v187
												} else {
													if v138&int32(240) == int32(224) {
														v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
														v187 = v138<<(uint(int32(12))%32)&int32(_a_F_unicode_normalize_func_0) | v158&int32(63)<<(uint(int32(6))%32)
														v188 = int32(2)
														v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v125))))
														v195 = v190&int32(63) | v187
													} else {
														if v138&int32(248) != int32(240) {
															v195 = int32(-1)
														} else {
															v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
															v175 = int32(63)
															v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+2)))
															v187 = v138<<(uint(int32(18))%32)&int32(_a_F_unicode_normalize_func_1) | v174&v175<<(uint(int32(12))%32) | v180&v175<<(uint(int32(6))%32)
															v188 = int32(3)
															v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v125))))
															v195 = v190&int32(63) | v187
														}
													}
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v114+v124<<(uint(int32(2))%32)))) = v195
											v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125))))
											if int32(0) <= v197 {
												v221 = int32(1)
											} else {
												v202 = v197 & int32(255)
												if v202&int32(224) == int32(192) {
													v221 = int32(2)
												} else {
													if v202&int32(240) == int32(224) {
														v221 = int32(3)
													} else {
														if v202&int32(248) == int32(240) {
															v219 = int32(4)
														} else {
															v219 = int32(1)
														}
														v221 = v219
													}
												}
											}
											v224 = v124 + int32(1)
											if v224 != v110 {
												v124 = v224
												v125 = v221 + v125
												continue
											} else {
												break
											}
											break
										}
									} else {
									}
									v235 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v114+v110<<(uint(int32(2))%32)))) = v235
									v241 = F_unicode_normalize(m, v72, v114)
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int64(0)
									} else {
										v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
										if v243 != 0 {
											v244 = v241
											v245 = v243
											v248 = v235
											for {
												if base.Ui32(v245) <= base.Ui32(int32(127)) {
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v245)
												} else {
													if base.Ui32(v245) <= base.Ui32(int32(2047)) {
														v261 = v245&int32(63) | int32(128)
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v261)
														v266 = int32(base.Ui32(v245)>>(uint(int32(6))%32)) | int32(192)
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v266)
													} else {
														if base.Ui32(v245) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
															v270 = int32(63)
															v272 = int32(128)
															v273 = v245&v270 | v272
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v273)
															v278 = int32(base.Ui32(v245)>>(uint(int32(12))%32)) | int32(224)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v278)
															v285 = int32(base.Ui32(v245)>>(uint(int32(6))%32))&v270 | v272
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v285)
														} else {
															v287 = int32(63)
															v289 = int32(128)
															v290 = v245&v287 | v289
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v290)
															v297 = int32(base.Ui32(v245)>>(uint(int32(6))%32))&v287 | v289
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v297)
															v304 = int32(base.Ui32(v245)>>(uint(int32(12))%32))&v287 | v289
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v304)
															v311 = int32(base.Ui32(v245)>>(uint(int32(18))%32))&int32(7) | int32(240)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v311)
														}
													}
												}
												v317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
												if int32(0) <= v317 {
													v341 = int32(1)
												} else {
													v322 = v317 & int32(255)
													if v322&int32(224) == int32(192) {
														v341 = int32(2)
													} else {
														if v322&int32(240) == int32(224) {
															v341 = int32(3)
														} else {
															if v322&int32(248) == int32(240) {
																v339 = int32(4)
															} else {
																v339 = int32(1)
															}
															v341 = v339
														}
													}
												}
												v342 = v341 + v248
												v343 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
												if v343 != 0 {
													v244 = v244 + int32(4)
													v245 = v343
													v248 = v342
													continue
												} else {
													break
												}
												break
											}
											v349 = v342 + int32(4)
										} else {
											v349 = v74
										}
										v355 = F_palloc(m, v349)
										mBase = m.M
										v356 = m.ExcPending
										if v356 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v355))) = v349 << (uint(int32(2)) % 32)
											v360 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
											if v360 != 0 {
												v363 = v360
												v364 = v355 + int32(4)
												v365 = v241
												for {
													if base.Ui32(v363) <= base.Ui32(int32(127)) {
														*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v363)
													} else {
														if base.Ui32(v363) <= base.Ui32(int32(2047)) {
															v380 = v363&int32(63) | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)) = uint8(v380)
															v385 = int32(base.Ui32(v363)>>(uint(int32(6))%32)) | int32(192)
															*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v385)
														} else {
															if base.Ui32(v363) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
																v389 = int32(63)
																v391 = int32(128)
																v392 = v363&v389 | v391
																*(*uint8)(unsafe.Add(mBase, uint32(v364)+2)) = uint8(v392)
																v397 = int32(base.Ui32(v363)>>(uint(int32(12))%32)) | int32(224)
																*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v397)
																v404 = int32(base.Ui32(v363)>>(uint(int32(6))%32))&v389 | v391
																*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)) = uint8(v404)
															} else {
																v406 = int32(63)
																v408 = int32(128)
																v409 = v363&v406 | v408
																*(*uint8)(unsafe.Add(mBase, uint32(v364)+3)) = uint8(v409)
																v416 = int32(base.Ui32(v363)>>(uint(int32(6))%32))&v406 | v408
																*(*uint8)(unsafe.Add(mBase, uint32(v364)+2)) = uint8(v416)
																v423 = int32(base.Ui32(v363)>>(uint(int32(12))%32))&v406 | v408
																*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)) = uint8(v423)
																v430 = int32(base.Ui32(v363)>>(uint(int32(18))%32))&int32(7) | int32(240)
																*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v430)
															}
														}
													}
													v434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v364))))
													if int32(0) <= v434 {
														v458 = int32(1)
													} else {
														v439 = v434 & int32(255)
														if v439&int32(224) == int32(192) {
															v458 = int32(2)
														} else {
															if v439&int32(240) == int32(224) {
																v458 = int32(3)
															} else {
																if v439&int32(248) == int32(240) {
																	v456 = int32(4)
																} else {
																	v456 = int32(1)
																}
																v458 = v456
															}
														}
													}
													v460 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
													if v460 != 0 {
														v363 = v460
														v364 = v458 + v364
														v365 = v365 + int32(4)
														continue
													} else {
														break
													}
													break
												}
											} else {
											}
											m.G0 = v12 + int32(16)
											return base.I64_extend_i32_u(v355)
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
func F_unicode_version(m *base.Module, l0 int32) int64 {
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
		return int64(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(3471773046342156320)
		return base.I64_extend_i32_u(v3)
	}
}
