package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v10 = F_palloc0(m, int32(22))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = int32(1)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		if v16&v14 != 0 {
			v19 = v14
		} else {
			v19 = int32(4)
		}
		v20 = v10 + v19
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v21)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)))
		*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v23)
		v28 = base.B2i32(v21 == int32(2))
		if v21 == int32(2) {
			v29 = int32(4)
		} else {
			v29 = int32(16)
		}
		if v29 != 0 {
			base.MemoryCopy(m, v20+int32(2), v8+int32(4), v29)
		} else {
		}
		if v21 == int32(2) {
			v37 = int32(40)
		} else {
			v37 = int32(88)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v37
		v40 = F_palloc(m, int32(24))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_extend_i32_u(v10)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v44
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v46
			v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
			v49 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+18)) = uint8(v49)
			*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)) = uint16(v48)
			return base.I64_extend_i32_u(v40)
		}
	}
}
func F_inet_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_network_out(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_inet_set_masklen(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v15 == int32(-1) {
			v20 = int32(1)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v24 = v22 & v20
			if v24 != 0 {
				v25 = v20
			} else {
				v25 = int32(4)
			}
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v25))))
			if v27 == int32(2) {
				v30 = int32(32)
			} else {
				v30 = int32(128)
			}
			v36 = v22
			v37 = v24
			v38 = v30
			if v37 != 0 {
				v43 = int32(1)
			} else {
				v43 = int32(4)
			}
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v43))))
			if v45 == int32(2) {
				v48 = int32(32)
			} else {
				v48 = int32(128)
			}
			if base.Ui32(v48) < base.Ui32(v38) {
				v119 = v38
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v119
						F_errmsg(m, int32(_a_F_inet_set_masklen_0), v8)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_inet_set_masklen_1), int32(332), int32(_a_F_inet_set_masklen_2))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
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
				if v36 == int32(1) {
					v53 = int32(18)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
					if v55 == v53 {
						v58 = v53
					} else {
						v58 = int32(2)
					}
					if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v65 = int32(6)
					} else {
						v65 = v58
					}
					v74 = v65
				} else {
					v66 = int32(1)
					if v36&v66 != 0 {
						v74 = int32(base.Ui32(v36) >> (uint(v66) % 32))
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						v74 = int32(base.Ui32(v70) >> (uint(int32(2)) % 32))
					}
				}
				v75 = F_palloc(m, v74)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int64(0)
				} else {
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if v77 == int32(1) {
						v81 = int32(18)
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
						if v83 == v81 {
							v86 = v81
						} else {
							v86 = int32(2)
						}
						if base.Ui32((v83-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v93 = int32(6)
						} else {
							v93 = v86
						}
						v102 = v93
					} else {
						v94 = int32(1)
						if v77&v94 != 0 {
							v102 = int32(base.Ui32(v77) >> (uint(v94) % 32))
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v102 = int32(base.Ui32(v98) >> (uint(int32(2)) % 32))
						}
					}
					if v102 != 0 {
						base.MemoryCopy(m, v75, v11, v102)
					} else {
					}
					v104 = int32(1)
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
					if v106&v104 != 0 {
						v109 = v104
					} else {
						v109 = int32(4)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v75+v109)+1)) = uint8(v38)
					m.G0 = v8 + int32(16)
					return base.I64_extend_i32_u(v75)
				}
			}
		} else {
			if v15 < int32(0) {
				v119 = v15
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v119
						F_errmsg(m, int32(_a_F_inet_set_masklen_0), v8)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_inet_set_masklen_1), int32(332), int32(_a_F_inet_set_masklen_2))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
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
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				v36 = v33
				v37 = v33 & int32(1)
				v38 = v15
				if v37 != 0 {
					v43 = int32(1)
				} else {
					v43 = int32(4)
				}
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v43))))
				if v45 == int32(2) {
					v48 = int32(32)
				} else {
					v48 = int32(128)
				}
				if base.Ui32(v48) < base.Ui32(v38) {
					v119 = v38
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v119
							F_errmsg(m, int32(_a_F_inet_set_masklen_0), v8)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_inet_set_masklen_1), int32(332), int32(_a_F_inet_set_masklen_2))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
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
					if v36 == int32(1) {
						v53 = int32(18)
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
						if v55 == v53 {
							v58 = v53
						} else {
							v58 = int32(2)
						}
						if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v65 = int32(6)
						} else {
							v65 = v58
						}
						v74 = v65
					} else {
						v66 = int32(1)
						if v36&v66 != 0 {
							v74 = int32(base.Ui32(v36) >> (uint(v66) % 32))
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v74 = int32(base.Ui32(v70) >> (uint(int32(2)) % 32))
						}
					}
					v75 = F_palloc(m, v74)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int64(0)
					} else {
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if v77 == int32(1) {
							v81 = int32(18)
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
							if v83 == v81 {
								v86 = v81
							} else {
								v86 = int32(2)
							}
							if base.Ui32((v83-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v93 = int32(6)
							} else {
								v93 = v86
							}
							v102 = v93
						} else {
							v94 = int32(1)
							if v77&v94 != 0 {
								v102 = int32(base.Ui32(v77) >> (uint(v94) % 32))
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								v102 = int32(base.Ui32(v98) >> (uint(int32(2)) % 32))
							}
						}
						if v102 != 0 {
							base.MemoryCopy(m, v75, v11, v102)
						} else {
						}
						v104 = int32(1)
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
						if v106&v104 != 0 {
							v109 = v104
						} else {
							v109 = int32(4)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v75+v109)+1)) = uint8(v38)
						m.G0 = v8 + int32(16)
						return base.I64_extend_i32_u(v75)
					}
				}
			}
		}
	}
}
func F_inet_spg_choose(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
	if v16 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v378+v9))) = base.I64_extend_i32_u(v371)
	return int64(0)
L4:
	;
	v371 = v12
	v378 = int32(16)
	goto L3
L5:
	;
	v19 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v23&v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v33 = F_pg_detoast_datum_packed(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v26 = v19
	goto L10
L9:
	;
	v26 = int32(4)
	goto L10
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v26))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = base.B2i32(v28 != int32(2))
	goto L4
L11:
	;
	v35 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v39 = v37 & v35
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = v35
	goto L14
L13:
	;
	v40 = int32(4)
	goto L14
L14:
	;
	v41 = v12 + v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v43 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v47 = v45 & v43
	if v47 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = v43
	goto L17
L16:
	;
	v48 = int32(4)
	goto L17
L17:
	;
	v49 = v33 + v48
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v42 != v50 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(2)
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(3)
	v58 = int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v60&v58 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if base.Ui32(v72) <= base.Ui32(v73) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v63 = v58
	goto L23
L22:
	;
	v63 = int32(4)
	goto L23
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v63))))
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v66)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = base.B2i32(v65 != int32(2))
	v371 = v33
	v378 = int32(40)
	goto L3
L24:
	;
	v329 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v329
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v336&v329 != 0 {
		goto L88
	} else {
		goto L89
	}
L25:
	;
	v75 = int32(2)
	v76 = v49 + v75
	v78 = v41 + v75
	v83 = base.I32_div_s(v72, int32(8))
	v84 = F_memcmp(m, v76, v78, v83)
	mBase = m.M
	if v84 != 0 {
		v170 = v84
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v189 = v39
	v190 = v47
	goto L27
L27:
	;
	if v190 != 0 {
		goto L50
	} else {
		goto L51
	}
L28:
	;
	if v180 == int32(0) {
		goto L24
	} else {
		goto L49
	}
L29:
	;
	if v171 != 0 {
		goto L46
	} else {
		goto L47
	}
L30:
	;
	v180 = v170
	goto L28
L31:
	;
	v85 = int32(0)
	v88 = v72 - v83<<(uint(int32(3))%32)
	if v88 <= v85 {
		v170 = v85
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v83))))
	v93 = int32(128)
	v94 = v92 & v93
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v83))))
	if v94 != v96&v93 {
		v171 = v94
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v88 == int32(1) {
		v170 = v85
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v102 = int32(1)
	v104 = int32(128)
	v105 = v92 << (uint(v102) % 32) & v104
	if v105 != v96<<(uint(v102)%32)&v104 {
		v171 = v105
		goto L29
	} else {
		goto L35
	}
L35:
	;
	if v88 < int32(3) {
		v170 = v85
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v113 = int32(2)
	v115 = int32(128)
	v116 = v92 << (uint(v113) % 32) & v115
	if v116 != v96<<(uint(v113)%32)&v115 {
		v171 = v116
		goto L29
	} else {
		goto L37
	}
L37:
	;
	if v88 == int32(3) {
		v170 = v85
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v124 = int32(3)
	v126 = int32(128)
	v127 = v92 << (uint(v124) % 32) & v126
	if v127 != v96<<(uint(v124)%32)&v126 {
		v171 = v127
		goto L29
	} else {
		goto L39
	}
L39:
	;
	if v88 < int32(5) {
		v170 = v85
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v135 = int32(4)
	v137 = int32(128)
	v138 = v92 << (uint(v135) % 32) & v137
	if v138 != v96<<(uint(v135)%32)&v137 {
		v171 = v138
		goto L29
	} else {
		goto L41
	}
L41:
	;
	if v88 == int32(5) {
		v170 = v85
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v146 = int32(5)
	v148 = int32(128)
	v149 = v92 << (uint(v146) % 32) & v148
	if v149 != v96<<(uint(v146)%32)&v148 {
		v171 = v149
		goto L29
	} else {
		goto L43
	}
L43:
	;
	if v88 < int32(7) {
		v170 = v85
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v157 = int32(6)
	v159 = int32(128)
	v160 = v92 << (uint(v157) % 32) & v159
	if v160 != v96<<(uint(v157)%32)&v159 {
		v171 = v160
		goto L29
	} else {
		goto L45
	}
L45:
	;
	v170 = v85
	goto L30
L46:
	;
	v174 = int32(1)
	goto L48
L47:
	;
	v174 = int32(-1)
	goto L48
L48:
	;
	v180 = v174
	goto L28
L49:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v184 = int32(1)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v189 = v186 & v184
	v190 = v183 & v184
	goto L27
L50:
	;
	v193 = int32(1)
	goto L52
L51:
	;
	v193 = int32(4)
	goto L52
L52:
	;
	v196 = v33 + v193 + int32(2)
	if v189 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v199 = int32(1)
	goto L55
L54:
	;
	v199 = int32(4)
	goto L55
L55:
	;
	v200 = v12 + v199
	v202 = v200 + int32(2)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)))
	if base.Ui32(v203) < base.Ui32(v72) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v205 = v203
	goto L58
L57:
	;
	v205 = v72
	goto L58
L58:
	;
	v206 = int32(0)
	v211 = int32(8)
	v212 = base.I32_div_s(v205, v211)
	if v211 <= v205 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v282)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(3)
	v286 = F_cidr_set_masklen_internal(m, v12, v281)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L75
	}
L60:
	;
	v281 = v278 + v274<<(uint(int32(3))%32)
	goto L59
L61:
	;
	v260 = v251
	goto L72
L62:
	;
	v218 = v206
	goto L65
L63:
	;
	v235 = v206
	goto L64
L64:
	;
	v242 = v205 - v212<<(uint(int32(3))%32)
	if v242 == int32(0) {
		v274 = v235
		v278 = v206
		goto L60
	} else {
		goto L71
	}
L65:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+v218))))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v218))))
	if v224 != v226 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v235 = v212
	goto L64
L67:
	;
	v251 = int32(7)
	v252 = v218
	v254 = v224
	v255 = v226
	goto L61
L68:
	;
	goto L69
L69:
	;
	v230 = v218 + int32(1)
	if v230 != v212 {
		v218 = v230
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v235))))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+v235))))
	v251 = v242
	v252 = v235
	v254 = v248
	v255 = v246
	goto L61
L72:
	;
	if int32(base.Ui32(v254^v255)>>(uint(int32(8)-v260)%32)) != 0 {
		v260 = v260 - int32(1)
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v274 = v252
	v278 = v260
	goto L60
L74:
	;
	goto L73
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_extend_i32_u(v286)
	v295 = int32(1)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v297&v295 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v300 = v295
	goto L78
L77:
	;
	v300 = int32(4)
	goto L78
L78:
	;
	v301 = v33 + v300
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v302 == int32(2) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v305 = int32(32)
	goto L81
L80:
	;
	v305 = int32(128)
	goto L81
L81:
	;
	if v281 < v305 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v308 = base.I32_div_s(v281, int32(8))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v308)+2)))
	v319 = int32(base.Ui32(v310)>>(uint(v308<<(uint(int32(3))%32)-v281+int32(7))%32)) & int32(1)
	goto L84
L83:
	;
	v319 = int32(0)
	goto L84
L84:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v321)
	if v281 < v320 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v326 = v319 | int32(2)
	goto L87
L86:
	;
	v326 = v319
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v326
	v371 = v33
	v378 = int32(40)
	goto L3
L88:
	;
	v339 = v329
	goto L90
L89:
	;
	v339 = int32(4)
	goto L90
L90:
	;
	v340 = v12 + v339
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v341 == int32(2) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v344 = int32(32)
	goto L93
L92:
	;
	v344 = int32(128)
	goto L93
L93:
	;
	if base.Ui32(v72) < base.Ui32(v344) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340+int32(base.Ui32(v72)>>(uint(int32(3))%32)))+2)))
	v357 = int32(base.Ui32(v349)>>(uint((v72^int32(-1))&int32(7))%32)) & int32(1)
	goto L96
L95:
	;
	v357 = int32(0)
	goto L96
L96:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	if base.Ui32(v72) < base.Ui32(v360) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v362 = v357 | int32(2)
	goto L99
L98:
	;
	v362 = v357
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v362
	goto L4
}
func F_inet_spg_leaf_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+8)) = uint8(v11)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = base.I64_extend_i32_u(v7)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v18 = F_inet_spg_consistent_bitmap(m, v7, v15, v16, int32(1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.B2i32(v18 != int32(0)))
		}
	}
}
