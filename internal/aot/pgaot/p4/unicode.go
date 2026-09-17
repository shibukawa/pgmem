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
						F_errdetail(m, int32(_a_F_addUnicodeChar_1), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l1, int32(_a_F_addUnicodeChar_2), int32(585), int32(_a_F_addUnicodeChar_3))
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
			if v37 == int32(447) {
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
										F_errsave_finish(m, l1, int32(_a_F_addUnicodeChar_2), int32(602), int32(_a_F_addUnicodeChar_3))
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
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
	v57 = F_palloc(m, v54+int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v31 == int32(18) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v42 = int32(1)
	if v25&v42 != 0 {
		v54 = int32(base.Ui32(v25)>>(uint(v42)%32)) - v42
		goto L3
	} else {
		goto L15
	}
L9:
	;
	v34 = int32(16)
	goto L11
L10:
	;
	v34 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = int32(4)
	goto L14
L13:
	;
	v41 = v34
	goto L14
L14:
	;
	v54 = v41
	goto L3
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L16:
	;
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = int32(1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v61&v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54+v57))) = uint8(v68)
	if v23 != v21 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v64 = v59
	goto L22
L21:
	;
	v64 = int32(4)
	goto L22
L22:
	;
	base.MemoryCopy(m, v57, v23+v64, v54)
	goto L19
L23:
	;
	F_pfree(m, v23)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v73 = F_unicode_norm_form_from_string(m, v57)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v75 = int32(1)
	v76 = v16 + v75
	v78 = v16 + int32(4)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v81 = v79 & v75
	if v81 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v82 = v76
	goto L30
L29:
	;
	v82 = v78
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
	v110 = F_pg_mbstrlen_with_len(m, v82, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L42
	}
L32:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v88 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v99 = int32(1)
	if v81 != 0 {
		v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v91 = int32(16)
	goto L37
L36:
	;
	v91 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v98 = int32(4)
	goto L40
L39:
	;
	v98 = v91
	goto L40
L40:
	;
	v109 = v98
	goto L31
L41:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	v113 = v110 << (uint(int32(2)) % 32)
	v116 = F_palloc(m, v113+int32(4))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v118 = int32(0)
	if v118 < v110 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v121&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v247 = int32(0)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v247<<(uint(int32(2))%32)))) = int32(0)
	v253 = int32(1)
	if v73&int32(-3) == v253 {
		v417 = int32(-1)
		goto L77
	} else {
		goto L78
	}
L47:
	;
	v124 = v76
	goto L49
L48:
	;
	v124 = v78
	goto L49
L49:
	;
	v125 = v118
	v126 = v124
	goto L50
L50:
	;
	v142 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
	v144 = v142 & int32(255)
	if int32(0) <= v142 {
		v201 = v144
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v247 = v110
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v125<<(uint(int32(2))%32)))) = v201
	v203 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
	if int32(0) <= v203 {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if v144&int32(224) == int32(192) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v126))))
	v201 = v196&int32(63) | v193
	goto L52
L55:
	;
	v193 = v144 << (uint(int32(6)) % 32) & int32(1984)
	v194 = int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	if v144&int32(240) == int32(224) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	v193 = v144<<(uint(int32(12))%32)&int32(_a_F_unicode_is_normalized_0) | v164&int32(63)<<(uint(int32(6))%32)
	v194 = int32(2)
	goto L54
L59:
	;
	goto L60
L60:
	;
	if v144&int32(248) != int32(240) {
		v201 = int32(-1)
		goto L52
	} else {
		goto L61
	}
L61:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	v181 = int32(63)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
	v193 = v144<<(uint(int32(18))%32)&int32(_a_F_unicode_is_normalized_1) | v180&v181<<(uint(int32(12))%32) | v186&v181<<(uint(int32(6))%32)
	v194 = int32(3)
	goto L54
L62:
	;
	v230 = v125 + int32(1)
	if v230 != v110 {
		v125 = v230
		v126 = v227 + v126
		goto L50
	} else {
		goto L75
	}
L63:
	;
	v227 = int32(1)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v208 = v203 & int32(255)
	if v208&int32(224) == int32(192) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v227 = int32(2)
	goto L62
L67:
	;
	goto L68
L68:
	;
	if v208&int32(240) == int32(224) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v227 = int32(3)
	goto L62
L70:
	;
	goto L71
L71:
	;
	if v208&int32(248) == int32(240) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v225 = int32(4)
	goto L74
L73:
	;
	v225 = int32(1)
	goto L74
L74:
	;
	v227 = v225
	goto L62
L75:
	;
	goto L51
L76:
	;
	return v507
L77:
	;
	if base.Ui32(v417) < base.Ui32(int32(2)) {
		v507 = v417
		goto L76
	} else {
		goto L97
	}
L78:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v260 == int32(0) {
		v417 = int32(1)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v265 = v253
	v267 = v116
	v268 = v260
	v276 = int32(0)
	goto L80
L80:
	;
	v277 = int32(0)
	v278 = int32(16711935)
	v280 = int32(8)
	v281 = base.I32_rotr(v268&v278, v280)
	v284 = int32(255)
	v285 = int32(base.Ui32(v281)>>(uint(v280)%32)) & v284
	v286 = int32(24)
	v287 = base.I32_rotr(v268, v286)
	v290 = (v281 | v287) & v284
	v291 = int32(_a_F_unicode_is_normalized_2)
	v299 = int32(base.Ui32(v287&v278) >> (uint(int32(16)) % 32))
	v304 = int32(base.Ui32(v281) >> (uint(v286) % 32))
	v305 = ((v285+v290*v291)*v291+v299)*v291 + v304
	v308 = int32(_a_F_unicode_is_normalized_3)
	v309 = base.I32_rem_u_s(v305+int32(402620417), v308)
	v310 = int32(1)
	v312 = int32(*(*int16)(unsafe.Add(mBase, uint32(v309<<(uint(v310)%32))+uint32(_c_F_unicode_is_normalized[0]))))
	v313 = int32(257)
	v321 = ((v290*v313+v285)*v313+v299)*v313 + v304
	v323 = base.I32_rem_u_s(v321, v308)
	v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v323<<(uint(v310)%32))+uint32(_c_F_unicode_is_normalized[0]))))
	v327 = v312 + v326
	if base.Ui32(int32(_a_F_unicode_is_normalized_4)) < base.Ui32(v327) {
		v344 = v277
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v417 = v398
	goto L77
L82:
	;
	switch v73 {
	case 0:
		goto L91
	default:
		goto L88
	case 2:
		goto L90
	}
L83:
	;
	v331 = v327 << (uint(int32(3)) % 32)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_c_F_unicode_is_normalized[1])))
	if v268 != v332 {
		v344 = v277
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+uint32(_c_F_unicode_is_normalized[2]))))
	if base.B2i32(base.Ui32(v276&int32(255)) <= base.Ui32(v336))|base.B2i32(v336 == int32(0)) != 0 {
		v344 = v336
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v417 = int32(0)
	goto L77
L86:
	;
	goto L81
L87:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v395 != 0 {
		v265 = v394
		v267 = v267 + int32(4)
		v268 = v395
		v276 = v344
		goto L80
	} else {
		goto L96
	}
L88:
	;
	v394 = v265
	goto L87
L89:
	;
	v390 = v386 << (uint(int32(7)) % 32) >> (uint(int32(28)) % 32)
	switch v390 + int32(1) {
	case 0:
		v394 = v390
		goto L87
	case 1:
		v398 = v390
		goto L86
	default:
		goto L88
	}
L90:
	;
	v367 = int32(_a_F_unicode_is_normalized_5)
	v368 = base.I32_rem_u_s(v305+int32(1207861251), v367)
	v369 = int32(1)
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368<<(uint(v369)%32))+uint32(_c_F_unicode_is_normalized[3]))))
	v373 = base.I32_rem_u_s(v321, v367)
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373<<(uint(v369)%32))+uint32(_c_F_unicode_is_normalized[3]))))
	v377 = v371 + v376
	if base.Ui32(int32(_a_F_unicode_is_normalized_6)) < base.Ui32(v377) {
		goto L88
	} else {
		goto L94
	}
L91:
	;
	v346 = int32(2505)
	v347 = base.I32_rem_u_s(v305, v346)
	v348 = int32(1)
	v350 = int32(*(*int16)(unsafe.Add(mBase, uint32(v347<<(uint(v348)%32))+uint32(_c_F_unicode_is_normalized[4]))))
	v352 = base.I32_rem_u_s(v321, v346)
	v355 = int32(*(*int16)(unsafe.Add(mBase, uint32(v352<<(uint(v348)%32))+uint32(_c_F_unicode_is_normalized[4]))))
	v356 = v350 + v355
	if base.Ui32(int32(1251)) < base.Ui32(v356) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v356<<(uint(int32(2))%32))+uint32(_c_F_unicode_is_normalized[5])))
	if v268 == v361&int32(_a_F_unicode_is_normalized_7) {
		v386 = v361
		goto L89
	} else {
		goto L93
	}
L93:
	;
	goto L88
L94:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v377<<(uint(int32(2))%32))+uint32(_c_F_unicode_is_normalized[6])))
	if v268 != v382&int32(_a_F_unicode_is_normalized_7) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v386 = v382
	goto L89
L96:
	;
	v398 = v394
	goto L86
L97:
	;
	v421 = F_unicode_normalize(m, v73, v116)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	if v423 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v425 = v421 + int32(4)
	v429 = v425
	goto L103
L100:
	;
	v440 = int32(0)
	goto L101
L101:
	;
	if v440 != v110 {
		v507 = int32(0)
		goto L76
	} else {
		goto L106
	}
L102:
	;
	v440 = (v429-v425)>>(uint(int32(2))%32) + int32(1)
	goto L101
L103:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v433 != 0 {
		v429 = v429 + int32(4)
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	goto L104
L106:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v113) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v507 = base.B2i32(v503 == int32(0))
	goto L76
L108:
	;
	v503 = int32(0)
	goto L107
L109:
	;
	v477 = v472
	v478 = v473
	v479 = v474
	goto L119
L110:
	;
	if (v116|v421)&int32(3) != 0 {
		v472 = v116
		v473 = v421
		v474 = v113
		goto L109
	} else {
		goto L113
	}
L111:
	;
	v465 = v116
	v466 = v421
	v467 = v113
	goto L112
L112:
	;
	if v467 == int32(0) {
		goto L108
	} else {
		goto L118
	}
L113:
	;
	v449 = v116
	v450 = v421
	v451 = v113
	goto L114
L114:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	if v454 != v455 {
		v472 = v449
		v473 = v450
		v474 = v451
		goto L109
	} else {
		goto L116
	}
L115:
	;
	v465 = v460
	v466 = v458
	v467 = v462
	goto L112
L116:
	;
	v457 = int32(4)
	v458 = v450 + v457
	v460 = v449 + v457
	v462 = v451 - v457
	if base.Ui32(int32(3)) < base.Ui32(v462) {
		v449 = v460
		v450 = v458
		v451 = v462
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v472 = v465
	v473 = v466
	v474 = v467
	goto L109
L119:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	if v482 == v483 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v503 = v482 - v483
	goto L107
L121:
	;
	v485 = int32(1)
	v490 = v479 - v485
	if v490 != 0 {
		v477 = v477 + v485
		v478 = v478 + v485
		v479 = v490
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	goto L120
L124:
	;
	goto L108
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v159 int32
	_ = v159
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
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
					return int32(0)
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
							return int32(0)
						} else {
							v72 = F_unicode_norm_form_from_string(m, v56)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v74 = int32(4)
								v75 = int32(1)
								v76 = v15 + v75
								v78 = v15 + v74
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
								v81 = v79 & v75
								if v81 != 0 {
									v82 = v76
								} else {
									v82 = v78
								}
								if v79 == int32(1) {
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
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
									if v81 != 0 {
										v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
										v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v110 = F_pg_mbstrlen_with_len(m, v82, v109)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v116 = F_palloc(m, v110<<(uint(int32(2))%32)+int32(4))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										v118 = int32(0)
										if v118 < v110 {
											v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
											if v121&int32(1) != 0 {
												v124 = v76
											} else {
												v124 = v78
											}
											v125 = v118
											v126 = v124
											for {
												v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
												v139 = v137 & int32(255)
												if int32(0) <= v137 {
													v196 = v139
												} else {
													if v139&int32(224) == int32(192) {
														v188 = v139 << (uint(int32(6)) % 32) & int32(1984)
														v189 = int32(1)
														v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v126))))
														v196 = v191&int32(63) | v188
													} else {
														if v139&int32(240) == int32(224) {
															v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
															v188 = v139<<(uint(int32(12))%32)&int32(_a_F_unicode_normalize_func_0) | v159&int32(63)<<(uint(int32(6))%32)
															v189 = int32(2)
															v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v126))))
															v196 = v191&int32(63) | v188
														} else {
															if v139&int32(248) != int32(240) {
																v196 = int32(-1)
															} else {
																v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
																v176 = int32(63)
																v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
																v188 = v139<<(uint(int32(18))%32)&int32(_a_F_unicode_normalize_func_1) | v175&v176<<(uint(int32(12))%32) | v181&v176<<(uint(int32(6))%32)
																v189 = int32(3)
																v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v126))))
																v196 = v191&int32(63) | v188
															}
														}
													}
												}
												*(*int32)(unsafe.Add(mBase, uint32(v116+v125<<(uint(int32(2))%32)))) = v196
												v198 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
												if int32(0) <= v198 {
													v222 = int32(1)
												} else {
													v203 = v198 & int32(255)
													if v203&int32(224) == int32(192) {
														v222 = int32(2)
													} else {
														if v203&int32(240) == int32(224) {
															v222 = int32(3)
														} else {
															if v203&int32(248) == int32(240) {
																v220 = int32(4)
															} else {
																v220 = int32(1)
															}
															v222 = v220
														}
													}
												}
												v225 = v125 + int32(1)
												if v225 != v110 {
													v125 = v225
													v126 = v222 + v126
													continue
												} else {
													break
												}
												break
											}
											v237 = v110
										} else {
											v237 = int32(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v116+v237<<(uint(int32(2))%32)))) = int32(0)
										v243 = F_unicode_normalize(m, v72, v116)
										mBase = m.M
										v244 = m.ExcPending
										if v244 != 0 {
											return int32(0)
										} else {
											v245 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
											if v245 != 0 {
												v247 = v243
												v248 = v245
												v251 = int32(0)
												for {
													if base.Ui32(v248) <= base.Ui32(int32(127)) {
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v248)
													} else {
														if base.Ui32(v248) <= base.Ui32(int32(2047)) {
															v264 = v248&int32(63) | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v264)
															v269 = int32(base.Ui32(v248)>>(uint(int32(6))%32)) | int32(192)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v269)
														} else {
															if base.Ui32(v248) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
																v273 = int32(63)
																v275 = int32(128)
																v276 = v248&v273 | v275
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v276)
																v281 = int32(base.Ui32(v248)>>(uint(int32(12))%32)) | int32(224)
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v281)
																v288 = int32(base.Ui32(v248)>>(uint(int32(6))%32))&v273 | v275
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v288)
															} else {
																v290 = int32(63)
																v292 = int32(128)
																v293 = v248&v290 | v292
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v293)
																v300 = int32(base.Ui32(v248)>>(uint(int32(6))%32))&v290 | v292
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v300)
																v307 = int32(base.Ui32(v248)>>(uint(int32(12))%32))&v290 | v292
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v307)
																v314 = int32(base.Ui32(v248)>>(uint(int32(18))%32))&int32(7) | int32(240)
																*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v314)
															}
														}
													}
													v320 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
													if int32(0) <= v320 {
														v344 = int32(1)
													} else {
														v325 = v320 & int32(255)
														if v325&int32(224) == int32(192) {
															v344 = int32(2)
														} else {
															if v325&int32(240) == int32(224) {
																v344 = int32(3)
															} else {
																if v325&int32(248) == int32(240) {
																	v342 = int32(4)
																} else {
																	v342 = int32(1)
																}
																v344 = v342
															}
														}
													}
													v345 = v344 + v251
													v346 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
													if v346 != 0 {
														v247 = v247 + int32(4)
														v248 = v346
														v251 = v345
														continue
													} else {
														break
													}
													break
												}
												v351 = v345 + int32(4)
											} else {
												v351 = v74
											}
											v358 = F_palloc(m, v351)
											mBase = m.M
											v359 = m.ExcPending
											if v359 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v358))) = v351 << (uint(int32(2)) % 32)
												v363 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
												if v363 != 0 {
													v366 = v363
													v367 = v358 + int32(4)
													v369 = v243
													for {
														if base.Ui32(v366) <= base.Ui32(int32(127)) {
															*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v366)
														} else {
															if base.Ui32(v366) <= base.Ui32(int32(2047)) {
																v383 = v366&int32(63) | int32(128)
																*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)) = uint8(v383)
																v388 = int32(base.Ui32(v366)>>(uint(int32(6))%32)) | int32(192)
																*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v388)
															} else {
																if base.Ui32(v366) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
																	v392 = int32(63)
																	v394 = int32(128)
																	v395 = v366&v392 | v394
																	*(*uint8)(unsafe.Add(mBase, uint32(v367)+2)) = uint8(v395)
																	v400 = int32(base.Ui32(v366)>>(uint(int32(12))%32)) | int32(224)
																	*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v400)
																	v407 = int32(base.Ui32(v366)>>(uint(int32(6))%32))&v392 | v394
																	*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)) = uint8(v407)
																} else {
																	v409 = int32(63)
																	v411 = int32(128)
																	v412 = v366&v409 | v411
																	*(*uint8)(unsafe.Add(mBase, uint32(v367)+3)) = uint8(v412)
																	v419 = int32(base.Ui32(v366)>>(uint(int32(6))%32))&v409 | v411
																	*(*uint8)(unsafe.Add(mBase, uint32(v367)+2)) = uint8(v419)
																	v426 = int32(base.Ui32(v366)>>(uint(int32(12))%32))&v409 | v411
																	*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)) = uint8(v426)
																	v433 = int32(base.Ui32(v366)>>(uint(int32(18))%32))&int32(7) | int32(240)
																	*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v433)
																}
															}
														}
														v437 = int32(*(*int8)(unsafe.Add(mBase, uint32(v367))))
														if int32(0) <= v437 {
															v461 = int32(1)
														} else {
															v442 = v437 & int32(255)
															if v442&int32(224) == int32(192) {
																v461 = int32(2)
															} else {
																if v442&int32(240) == int32(224) {
																	v461 = int32(3)
																} else {
																	if v442&int32(248) == int32(240) {
																		v459 = int32(4)
																	} else {
																		v459 = int32(1)
																	}
																	v461 = v459
																}
															}
														}
														v463 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
														if v463 != 0 {
															v366 = v463
															v367 = v461 + v367
															v369 = v369 + int32(4)
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												m.G0 = v12 + int32(16)
												return v358
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
							return int32(0)
						} else {
							v74 = int32(4)
							v75 = int32(1)
							v76 = v15 + v75
							v78 = v15 + v74
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
							v81 = v79 & v75
							if v81 != 0 {
								v82 = v76
							} else {
								v82 = v78
							}
							if v79 == int32(1) {
								v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
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
								if v81 != 0 {
									v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
									v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v110 = F_pg_mbstrlen_with_len(m, v82, v109)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v116 = F_palloc(m, v110<<(uint(int32(2))%32)+int32(4))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									v118 = int32(0)
									if v118 < v110 {
										v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
										if v121&int32(1) != 0 {
											v124 = v76
										} else {
											v124 = v78
										}
										v125 = v118
										v126 = v124
										for {
											v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
											v139 = v137 & int32(255)
											if int32(0) <= v137 {
												v196 = v139
											} else {
												if v139&int32(224) == int32(192) {
													v188 = v139 << (uint(int32(6)) % 32) & int32(1984)
													v189 = int32(1)
													v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v126))))
													v196 = v191&int32(63) | v188
												} else {
													if v139&int32(240) == int32(224) {
														v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
														v188 = v139<<(uint(int32(12))%32)&int32(_a_F_unicode_normalize_func_0) | v159&int32(63)<<(uint(int32(6))%32)
														v189 = int32(2)
														v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v126))))
														v196 = v191&int32(63) | v188
													} else {
														if v139&int32(248) != int32(240) {
															v196 = int32(-1)
														} else {
															v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
															v176 = int32(63)
															v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
															v188 = v139<<(uint(int32(18))%32)&int32(_a_F_unicode_normalize_func_1) | v175&v176<<(uint(int32(12))%32) | v181&v176<<(uint(int32(6))%32)
															v189 = int32(3)
															v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v126))))
															v196 = v191&int32(63) | v188
														}
													}
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v116+v125<<(uint(int32(2))%32)))) = v196
											v198 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
											if int32(0) <= v198 {
												v222 = int32(1)
											} else {
												v203 = v198 & int32(255)
												if v203&int32(224) == int32(192) {
													v222 = int32(2)
												} else {
													if v203&int32(240) == int32(224) {
														v222 = int32(3)
													} else {
														if v203&int32(248) == int32(240) {
															v220 = int32(4)
														} else {
															v220 = int32(1)
														}
														v222 = v220
													}
												}
											}
											v225 = v125 + int32(1)
											if v225 != v110 {
												v125 = v225
												v126 = v222 + v126
												continue
											} else {
												break
											}
											break
										}
										v237 = v110
									} else {
										v237 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v116+v237<<(uint(int32(2))%32)))) = int32(0)
									v243 = F_unicode_normalize(m, v72, v116)
									mBase = m.M
									v244 = m.ExcPending
									if v244 != 0 {
										return int32(0)
									} else {
										v245 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
										if v245 != 0 {
											v247 = v243
											v248 = v245
											v251 = int32(0)
											for {
												if base.Ui32(v248) <= base.Ui32(int32(127)) {
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v248)
												} else {
													if base.Ui32(v248) <= base.Ui32(int32(2047)) {
														v264 = v248&int32(63) | int32(128)
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v264)
														v269 = int32(base.Ui32(v248)>>(uint(int32(6))%32)) | int32(192)
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v269)
													} else {
														if base.Ui32(v248) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
															v273 = int32(63)
															v275 = int32(128)
															v276 = v248&v273 | v275
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v276)
															v281 = int32(base.Ui32(v248)>>(uint(int32(12))%32)) | int32(224)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v281)
															v288 = int32(base.Ui32(v248)>>(uint(int32(6))%32))&v273 | v275
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v288)
														} else {
															v290 = int32(63)
															v292 = int32(128)
															v293 = v248&v290 | v292
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v293)
															v300 = int32(base.Ui32(v248)>>(uint(int32(6))%32))&v290 | v292
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v300)
															v307 = int32(base.Ui32(v248)>>(uint(int32(12))%32))&v290 | v292
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v307)
															v314 = int32(base.Ui32(v248)>>(uint(int32(18))%32))&int32(7) | int32(240)
															*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v314)
														}
													}
												}
												v320 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
												if int32(0) <= v320 {
													v344 = int32(1)
												} else {
													v325 = v320 & int32(255)
													if v325&int32(224) == int32(192) {
														v344 = int32(2)
													} else {
														if v325&int32(240) == int32(224) {
															v344 = int32(3)
														} else {
															if v325&int32(248) == int32(240) {
																v342 = int32(4)
															} else {
																v342 = int32(1)
															}
															v344 = v342
														}
													}
												}
												v345 = v344 + v251
												v346 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
												if v346 != 0 {
													v247 = v247 + int32(4)
													v248 = v346
													v251 = v345
													continue
												} else {
													break
												}
												break
											}
											v351 = v345 + int32(4)
										} else {
											v351 = v74
										}
										v358 = F_palloc(m, v351)
										mBase = m.M
										v359 = m.ExcPending
										if v359 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v358))) = v351 << (uint(int32(2)) % 32)
											v363 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
											if v363 != 0 {
												v366 = v363
												v367 = v358 + int32(4)
												v369 = v243
												for {
													if base.Ui32(v366) <= base.Ui32(int32(127)) {
														*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v366)
													} else {
														if base.Ui32(v366) <= base.Ui32(int32(2047)) {
															v383 = v366&int32(63) | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)) = uint8(v383)
															v388 = int32(base.Ui32(v366)>>(uint(int32(6))%32)) | int32(192)
															*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v388)
														} else {
															if base.Ui32(v366) <= base.Ui32(int32(_a_F_unicode_normalize_func_2)) {
																v392 = int32(63)
																v394 = int32(128)
																v395 = v366&v392 | v394
																*(*uint8)(unsafe.Add(mBase, uint32(v367)+2)) = uint8(v395)
																v400 = int32(base.Ui32(v366)>>(uint(int32(12))%32)) | int32(224)
																*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v400)
																v407 = int32(base.Ui32(v366)>>(uint(int32(6))%32))&v392 | v394
																*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)) = uint8(v407)
															} else {
																v409 = int32(63)
																v411 = int32(128)
																v412 = v366&v409 | v411
																*(*uint8)(unsafe.Add(mBase, uint32(v367)+3)) = uint8(v412)
																v419 = int32(base.Ui32(v366)>>(uint(int32(6))%32))&v409 | v411
																*(*uint8)(unsafe.Add(mBase, uint32(v367)+2)) = uint8(v419)
																v426 = int32(base.Ui32(v366)>>(uint(int32(12))%32))&v409 | v411
																*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)) = uint8(v426)
																v433 = int32(base.Ui32(v366)>>(uint(int32(18))%32))&int32(7) | int32(240)
																*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v433)
															}
														}
													}
													v437 = int32(*(*int8)(unsafe.Add(mBase, uint32(v367))))
													if int32(0) <= v437 {
														v461 = int32(1)
													} else {
														v442 = v437 & int32(255)
														if v442&int32(224) == int32(192) {
															v461 = int32(2)
														} else {
															if v442&int32(240) == int32(224) {
																v461 = int32(3)
															} else {
																if v442&int32(248) == int32(240) {
																	v459 = int32(4)
																} else {
																	v459 = int32(1)
																}
																v461 = v459
															}
														}
													}
													v463 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
													if v463 != 0 {
														v366 = v463
														v367 = v461 + v367
														v369 = v369 + int32(4)
														continue
													} else {
														break
													}
													break
												}
											} else {
											}
											m.G0 = v12 + int32(16)
											return v358
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
