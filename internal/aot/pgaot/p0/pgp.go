package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_cfb_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	m.T0[v4].(func(*base.Module, int32))(m, v3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v9 = F___memset(m, l0, int32(0), int32(116))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgp_create_pkt_reader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v7 = F_palloc(m, int32(8))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l3
		v14 = F_pullf_create(m, l0, int32(4394608), v7, l1)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v14 < int32(0) {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v14
				}
			} else {
				return v14
			}
		}
	}
}
func F_pgp_get_unicode_mode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2
}
func F_pgp_key_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	if l0 != 0 {
		v10 = int32(24)
		v11 = int32(12)
		v12 = int32(8)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
		switch v13 - int32(1) {
		case 0, 1, 2:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v17 = F_pgp_mpi_free(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v24 = int32(24)
				v25 = int32(12)
				v26 = int32(28)
				v27 = int32(32)
				v28 = int32(36)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0+v25)))
				v31 = F_pgp_mpi_free(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = v26
					v34 = v24
					v37 = v28
					v38 = v27
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
					v41 = F_pgp_mpi_free(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
						v45 = F_pgp_mpi_free(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38)))
							v49 = F_pgp_mpi_free(m, v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37)))
								v53 = F_pgp_mpi_free(m, v52)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v63 = F___memset(m, l0, int32(0), int32(52))
									mBase = m.M
									F_pfree(m, l0)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		default:
			v63 = F___memset(m, l0, int32(0), int32(52))
			mBase = m.M
			F_pfree(m, l0)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				return
			}
		case 15:
			v33 = v11
			v34 = v12
			v37 = v10
			v38 = v13
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
			v41 = F_pgp_mpi_free(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
				v45 = F_pgp_mpi_free(m, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38)))
					v49 = F_pgp_mpi_free(m, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37)))
						v53 = F_pgp_mpi_free(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v63 = F___memset(m, l0, int32(0), int32(52))
							mBase = m.M
							F_pfree(m, l0)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		case 16:
			v24 = v11
			v25 = v12
			v26 = int32(16)
			v27 = int32(20)
			v28 = v10
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0+v25)))
			v31 = F_pgp_mpi_free(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = v26
				v34 = v24
				v37 = v28
				v38 = v27
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
				v41 = F_pgp_mpi_free(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
					v45 = F_pgp_mpi_free(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38)))
						v49 = F_pgp_mpi_free(m, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37)))
							v53 = F_pgp_mpi_free(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v63 = F___memset(m, l0, int32(0), int32(52))
								mBase = m.M
								F_pfree(m, l0)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F_pgp_mpi_write(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v10 = int32(8)
	v14 = v9<<(uint(v10)%32) | int32(base.Ui32(v9)>>(uint(v10)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v14)
	v19 = F_pushf_write(m, l0, v7+int32(14), int32(2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v19 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v27 = F_pushf_write(m, l0, v25, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = v27
				m.G0 = v7 + int32(16)
				return v29
			}
		} else {
			v29 = v19
			m.G0 = v7 + int32(16)
			return v29
		}
	}
}
func F_pgp_rsa_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = int32(-109)
	v13 = F_mpi_check(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v58 = v12
			return v58
		} else {
			v19 = F_mpi_check(m, v9)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v58 = v12
					return v58
				} else {
					v23 = F_mpi_check(m, v10)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 == int32(0) {
							v58 = v12
							return v58
						} else {
							v27 = int32(1)
							if v11 <= v27 {
								v30 = v27
							} else {
								v30 = v11
							}
							v31 = F_palloc(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v40 = m.Env.Pgmem_bn_op(m, int32(1), v34, v35, v36, v37, v38, v39, v31, v11)
								mBase = m.M
								if v40 < int32(0) {
									v51 = int32(-109)
									if v31 == int32(0) {
										v58 = v51
										return v58
									} else {
										v55 = F___memset(m, v31, int32(0), v30)
										mBase = m.M
										F_pfree(m, v31)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = v51
											return v58
										}
									}
								} else {
									v44 = F_bytes_to_mpi(m, v31, v40)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v44
										if v44 != 0 {
											v49 = int32(0)
										} else {
											v49 = int32(-109)
										}
										v51 = v49
										if v31 == int32(0) {
											v58 = v51
											return v58
										} else {
											v55 = F___memset(m, v31, int32(0), v30)
											mBase = m.M
											F_pfree(m, v31)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v58 = v51
												return v58
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
func F_pgp_s2k_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v23 = l1 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v23) {
		v38 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)) = uint8(v38)
	if v38&int32(255) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	if int32(base.Ui32(int32(487))>>(uint(v23)%32))&int32(1) == int32(0) {
		v38 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_consts[1096])))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v38 = v37
	goto L2
L5:
	;
	m.G0 = v19 + int32(80)
	return v428
L6:
	;
	v428 = int32(-103)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v48 = F_pgp_load_digest(m, v45, v19+int32(12))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v48 < int32(0) {
		v428 = v48
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v55 {
	case 0:
		goto L16
	case 1:
		goto L14
	default:
		v408 = int32(-121)
		goto L12
	case 3:
		goto L15
	}
L12:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+20))
	m.T0[v424].(func(*base.Module, int32))(m, v423)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L9
	} else {
		goto L112
	}
L13:
	;
	v401 = int32(0)
	v406 = F___memset(m, v19+int32(16), v401, int32(64))
	mBase = m.M
	goto L111
L14:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v299 = m.T0[v298].(func(*base.Module, int32) int32)(m, v297)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L85
	}
L15:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v119 = m.T0[v118].(func(*base.Module, int32) int32)(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L40
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = m.T0[v57].(func(*base.Module, int32) int32)(m, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v60 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v66 = int32(0)
	v72 = v60
	v73 = l0 + int32(11)
	goto L19
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	m.T0[v82].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	if v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v89 = F__emscripten_memset_bulkmem(m, v19+int32(16), base.I32_extend8_s(int32(0)), v66)
	mBase = m.M
	goto L25
L23:
	;
	goto L24
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	m.T0[v95].(func(*base.Module, int32, int32, int32))(m, v56, l2, l3)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L27
	}
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	m.T0[v92].(func(*base.Module, int32, int32, int32))(m, v56, v19+int32(16), v66)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	m.T0[v100].(func(*base.Module, int32, int32))(m, v56, v19+int32(16))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v72) <= base.Ui32(v58) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v72 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	if v58 != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	goto L13
L33:
	;
	v106 = F__emscripten_memcpy_bulkmem(m, v73, v19+int32(16), v72)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v66 = v66 + int32(1)
	v72 = v72 - v58
	v73 = v112 + v58
	goto L19
L37:
	;
	v111 = F__emscripten_memcpy_bulkmem(m, v73, v19+int32(16), v58)
	mBase = m.M
	v112 = v111
	goto L39
L38:
	;
	v112 = v73
	goto L39
L39:
	;
	goto L36
L40:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v121 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v291 = int32(0)
	v296 = F___memset(m, v19+int32(16), v291, int32(64))
	mBase = m.M
	goto L84
L42:
	;
	v124 = int32(8)
	v127 = int32(16)
	v133 = (v116&int32(15) | v127) << (uint(int32(base.Ui32(v116)>>(uint(int32(4))%32))+int32(6)) % 32)
	v135 = l3 + v124
	if base.Ui32(l3+v127) < base.Ui32(v133) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v140 = v124
	goto L45
L44:
	;
	v140 = v133 - v135
	goto L45
L45:
	;
	v142 = v140 + (l3 + v135)
	v145 = l0 + int32(2)
	v156 = v5
	v158 = v121
	v160 = l0 + int32(11)
	goto L46
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	m.T0[v164].(func(*base.Module, int32))(m, v117)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	if v156 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v171 = F__emscripten_memset_bulkmem(m, v19+int32(16), base.I32_extend8_s(int32(0)), v156)
	mBase = m.M
	goto L52
L50:
	;
	goto L51
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v178].(func(*base.Module, int32, int32, int32))(m, v117, v145, int32(8))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L54
	}
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v174].(func(*base.Module, int32, int32, int32))(m, v117, v19+int32(16), v156)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v181].(func(*base.Module, int32, int32, int32))(m, v117, l2, l3)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	if base.Ui32(v133) <= base.Ui32(v135) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	m.T0[v259].(func(*base.Module, int32, int32))(m, v117, v19+int32(16))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L9
	} else {
		goto L72
	}
L57:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v185].(func(*base.Module, int32, int32, int32))(m, v117, v145, v140)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	if base.Ui32(v142) < base.Ui32(v133) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v189 = v142
	goto L62
L60:
	;
	v226 = v135 + v140
	goto L61
L61:
	;
	if base.Ui32(v133) <= base.Ui32(v226) {
		goto L56
	} else {
		goto L70
	}
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v205].(func(*base.Module, int32, int32, int32))(m, v117, l2, l3)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L64
	}
L63:
	;
	v226 = v217
	goto L61
L64:
	;
	v208 = int32(8)
	if base.Ui32(v189+v208) < base.Ui32(v133) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v213 = v208
	goto L67
L66:
	;
	v213 = v133 - v189
	goto L67
L67:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v214].(func(*base.Module, int32, int32, int32))(m, v117, v145, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v217 = v189 + v213
	v218 = v217 + l3
	if base.Ui32(v218) < base.Ui32(v133) {
		v189 = v218
		goto L62
	} else {
		goto L69
	}
L69:
	;
	goto L63
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v238].(func(*base.Module, int32, int32, int32))(m, v117, l2, v133-v226)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	goto L56
L72:
	;
	if base.Ui32(v158) <= base.Ui32(v119) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v158 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	goto L75
L75:
	;
	if v119 != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	goto L41
L77:
	;
	v265 = F__emscripten_memcpy_bulkmem(m, v160, v19+int32(16), v158)
	mBase = m.M
	goto L79
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	v156 = v156 + int32(1)
	v158 = v158 - v119
	v160 = v272 + v119
	goto L46
L81:
	;
	v271 = F__emscripten_memcpy_bulkmem(m, v160, v19+int32(16), v119)
	mBase = m.M
	v272 = v271
	goto L83
L82:
	;
	v272 = v160
	goto L83
L83:
	;
	goto L80
L84:
	;
	v408 = v291
	goto L12
L85:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v301 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v379 = int32(0)
	v384 = F___memset(m, v19+int32(16), v379, int32(64))
	mBase = m.M
	goto L110
L87:
	;
	v309 = int32(0)
	v315 = v301
	v316 = l0 + int32(11)
	goto L88
L88:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	m.T0[v325].(func(*base.Module, int32))(m, v297)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	if v309 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v332 = F__emscripten_memset_bulkmem(m, v19+int32(16), base.I32_extend8_s(int32(0)), v309)
	mBase = m.M
	goto L94
L92:
	;
	goto L93
L93:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	m.T0[v339].(func(*base.Module, int32, int32, int32))(m, v297, l0+int32(2), int32(8))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L9
	} else {
		goto L96
	}
L94:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	m.T0[v335].(func(*base.Module, int32, int32, int32))(m, v297, v19+int32(16), v309)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	m.T0[v342].(func(*base.Module, int32, int32, int32))(m, v297, l2, l3)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
	m.T0[v347].(func(*base.Module, int32, int32))(m, v297, v19+int32(16))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	if base.Ui32(v315) <= base.Ui32(v299) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v315 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	if v299 != 0 {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	goto L86
L103:
	;
	v353 = F__emscripten_memcpy_bulkmem(m, v316, v19+int32(16), v315)
	mBase = m.M
	goto L105
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v309 = v309 + int32(1)
	v315 = v315 - v299
	v316 = v358 + v299
	goto L88
L107:
	;
	v357 = F__emscripten_memcpy_bulkmem(m, v316, v19+int32(16), v299)
	mBase = m.M
	v358 = v357
	goto L109
L108:
	;
	v358 = v316
	goto L109
L109:
	;
	goto L106
L110:
	;
	v408 = v379
	goto L12
L111:
	;
	v408 = v401
	goto L12
L112:
	;
	v428 = v408
	goto L5
}
func F_pgp_set_s2k_digest_algo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v8 = F_pg_strcasecmp(m, int32(556880), l1)
	mBase = m.M
	if v8 == int32(0) {
		v61 = int32(4394736)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
		v63 = v62
	} else {
		v16 = F_pg_strcasecmp(m, int32(561202), l1)
		mBase = m.M
		if v16 == int32(0) {
			v61 = int32(4394744)
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
			v63 = v62
		} else {
			v24 = F_pg_strcasecmp(m, int32(561623), l1)
			mBase = m.M
			if v24 == int32(0) {
				v61 = int32(4394752)
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				v63 = v62
			} else {
				v32 = F_pg_strcasecmp(m, int32(562552), l1)
				mBase = m.M
				if v32 == int32(0) {
					v61 = int32(4394760)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
					v63 = v62
				} else {
					v40 = F_pg_strcasecmp(m, int32(556466), l1)
					mBase = m.M
					if v40 == int32(0) {
						v61 = int32(4394768)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
						v63 = v62
					} else {
						v48 = F_pg_strcasecmp(m, int32(558360), l1)
						mBase = m.M
						if v48 == int32(0) {
							v61 = int32(4394776)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
							v63 = v62
						} else {
							v57 = F_pg_strcasecmp(m, int32(560656), l1)
							mBase = m.M
							if v57 != 0 {
								v63 = int32(-104)
							} else {
								v61 = int32(4394784)
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
								v63 = v62
							}
						}
					}
				}
			}
		}
	}
	if v63 < int32(0) {
		return v63
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v63
		return int32(0)
	}
}
func F_pgp_set_text_mode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = l1
	return int32(0)
}
