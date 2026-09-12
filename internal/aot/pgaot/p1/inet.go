package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_fetch(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v10 = F_palloc0(m, int32(22))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(1)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		v18 = v16 & v14
		if v18 != 0 {
			v19 = v14
		} else {
			v19 = int32(4)
		}
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(v10+v19))) = uint8(v21)
		v24 = v10 + int32(1)
		v26 = v10 + int32(4)
		if v18 != 0 {
			v27 = v24
		} else {
			v27 = v26
		}
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)))
		*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v28)
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		v32 = v30 & int32(1)
		if v32 != 0 {
			v33 = v24
		} else {
			v33 = v26
		}
		v36 = int32(4)
		if v32 != 0 {
			v42 = int32(1)
		} else {
			v42 = v36
		}
		v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v42))))
		if v44 == int32(2) {
			v47 = v36
		} else {
			v47 = int32(16)
		}
		if v47 != 0 {
			v48 = F__emscripten_memcpy_bulkmem(m, v33+int32(2), v8+v36, v47)
			mBase = m.M
		} else {
		}
		v52 = int32(1)
		v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		if v54&v52 != 0 {
			v57 = v52
		} else {
			v57 = int32(4)
		}
		v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v57))))
		if v59 == int32(2) {
			v62 = int32(40)
		} else {
			v62 = int32(88)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v62
		v65 = F_palloc(m, int32(16))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v65))) = v10
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v70
			v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
			v73 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v65)+14)) = uint8(v73)
			*(*uint16)(unsafe.Add(mBase, uint32(v65)+12)) = uint16(v72)
			return v65
		}
	}
}
func F_inet_out(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_network_out(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_inet_set_masklen(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v15 == int32(-1) {
			v20 = int32(1)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v22&v20 != 0 {
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
			v34 = v30
			v35 = v22
			v38 = int32(1)
			v41 = v35 & v38
			if v41 != 0 {
				v42 = v38
			} else {
				v42 = int32(4)
			}
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v42))))
			if v44 == int32(2) {
				v47 = int32(32)
			} else {
				v47 = int32(128)
			}
			if base.Ui32(v47) < base.Ui32(v34) {
				v118 = v34
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v118
						F_errmsg(m, int32(474685), v8)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488675), int32(334), int32(277127))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
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
				if v35 == int32(1) {
					v51 = int32(6)
					v53 = int32(18)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
					if v55 == v53 {
						v58 = v53
					} else {
						v58 = int32(2)
					}
					if v55&int32(254) == int32(2) {
						v63 = v51
					} else {
						v63 = v58
					}
					if v55 == int32(1) {
						v66 = v51
					} else {
						v66 = v63
					}
					v73 = v66
				} else {
					if v41 != 0 {
						v73 = int32(base.Ui32(v35) >> (uint(int32(1)) % 32))
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						v73 = int32(base.Ui32(v69) >> (uint(int32(2)) % 32))
					}
				}
				v74 = F_palloc(m, v73)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if v76 == int32(1) {
						v79 = int32(6)
						v81 = int32(18)
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
						if v83 == v81 {
							v86 = v81
						} else {
							v86 = int32(2)
						}
						if v83&int32(254) == int32(2) {
							v91 = v79
						} else {
							v91 = v86
						}
						if v83 == int32(1) {
							v94 = v79
						} else {
							v94 = v91
						}
						v103 = v94
					} else {
						v95 = int32(1)
						if v76&v95 != 0 {
							v103 = int32(base.Ui32(v76) >> (uint(v95) % 32))
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v103 = int32(base.Ui32(v99) >> (uint(int32(2)) % 32))
						}
					}
					if v103 != 0 {
						v104 = F__emscripten_memcpy_bulkmem(m, v74, v11, v103)
						mBase = m.M
						v105 = v104
					} else {
						v105 = v74
					}
					v106 = int32(1)
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
					if v108&v106 != 0 {
						v111 = v106
					} else {
						v111 = int32(4)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v105+v111)+1)) = uint8(v34)
					m.G0 = v8 + int32(16)
					return v105
				}
			}
		} else {
			if v15 < int32(0) {
				v118 = v15
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v118
						F_errmsg(m, int32(474685), v8)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488675), int32(334), int32(277127))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
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
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				v34 = v15
				v35 = v33
				v38 = int32(1)
				v41 = v35 & v38
				if v41 != 0 {
					v42 = v38
				} else {
					v42 = int32(4)
				}
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v42))))
				if v44 == int32(2) {
					v47 = int32(32)
				} else {
					v47 = int32(128)
				}
				if base.Ui32(v47) < base.Ui32(v34) {
					v118 = v34
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v118
							F_errmsg(m, int32(474685), v8)
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(488675), int32(334), int32(277127))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
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
					if v35 == int32(1) {
						v51 = int32(6)
						v53 = int32(18)
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
						if v55 == v53 {
							v58 = v53
						} else {
							v58 = int32(2)
						}
						if v55&int32(254) == int32(2) {
							v63 = v51
						} else {
							v63 = v58
						}
						if v55 == int32(1) {
							v66 = v51
						} else {
							v66 = v63
						}
						v73 = v66
					} else {
						if v41 != 0 {
							v73 = int32(base.Ui32(v35) >> (uint(int32(1)) % 32))
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v73 = int32(base.Ui32(v69) >> (uint(int32(2)) % 32))
						}
					}
					v74 = F_palloc(m, v73)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if v76 == int32(1) {
							v79 = int32(6)
							v81 = int32(18)
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
							if v83 == v81 {
								v86 = v81
							} else {
								v86 = int32(2)
							}
							if v83&int32(254) == int32(2) {
								v91 = v79
							} else {
								v91 = v86
							}
							if v83 == int32(1) {
								v94 = v79
							} else {
								v94 = v91
							}
							v103 = v94
						} else {
							v95 = int32(1)
							if v76&v95 != 0 {
								v103 = int32(base.Ui32(v76) >> (uint(v95) % 32))
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								v103 = int32(base.Ui32(v99) >> (uint(int32(2)) % 32))
							}
						}
						if v103 != 0 {
							v104 = F__emscripten_memcpy_bulkmem(m, v74, v11, v103)
							mBase = m.M
							v105 = v104
						} else {
							v105 = v74
						}
						v106 = int32(1)
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
						if v108&v106 != 0 {
							v111 = v106
						} else {
							v111 = int32(4)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v105+v111)+1)) = uint8(v34)
						m.G0 = v8 + int32(16)
						return v105
					}
				}
			}
		}
	}
}
func F_inet_spg_choose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v382 int32
	_ = v382
	var v393 int32
	_ = v393
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
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
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v393+v13))) = v382
	return int32(0)
L4:
	;
	v382 = v16
	v393 = int32(12)
	goto L3
L5:
	;
	v23 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v23
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v27&v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v37 = F_pg_detoast_datum_packed(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v30 = v23
	goto L10
L9:
	;
	v30 = int32(4)
	goto L10
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v30))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = base.B2i32(v32 != int32(2))
	goto L4
L11:
	;
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v43 = v41 & v39
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v39
	goto L14
L13:
	;
	v44 = int32(4)
	goto L14
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v44))))
	v47 = int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v51 = v49 & v47
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = v47
	goto L17
L16:
	;
	v52 = int32(4)
	goto L17
L17:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v52))))
	if v46 != v54 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = int64(2)
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
	v62 = int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v64&v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v77 = v37 + int32(1)
	v79 = v37 + int32(4)
	if v51 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v67 = v62
	goto L23
L22:
	;
	v67 = int32(4)
	goto L23
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v67))))
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = base.B2i32(v69 != int32(2))
	v382 = v37
	v393 = int32(28)
	goto L3
L24:
	;
	v336 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v336
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v343&v336 != 0 {
		goto L94
	} else {
		goto L95
	}
L25:
	;
	v80 = v77
	goto L27
L26:
	;
	v80 = v79
	goto L27
L27:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	v83 = v16 + int32(1)
	v85 = v16 + int32(4)
	if v43 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v86 = v83
	goto L30
L29:
	;
	v86 = v85
	goto L30
L30:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if base.Ui32(v81) <= base.Ui32(v87) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = int32(2)
	v90 = v80 + v89
	v92 = v86 + v89
	v97 = base.I32_div_s(v81, int32(8))
	v98 = F_memcmp(m, v90, v92, v97)
	mBase = m.M
	if v98 != 0 {
		v184 = v98
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v203 = v43
	v204 = v51
	goto L33
L33:
	;
	if v204 != 0 {
		goto L56
	} else {
		goto L57
	}
L34:
	;
	if v194 == int32(0) {
		goto L24
	} else {
		goto L55
	}
L35:
	;
	if v185 != 0 {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v194 = v184
	goto L34
L37:
	;
	v99 = int32(0)
	v102 = v81 - v97<<(uint(int32(3))%32)
	if v102 <= v99 {
		v184 = v99
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v97))))
	v107 = int32(128)
	v108 = v106 & v107
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v97))))
	if v108 != v110&v107 {
		v185 = v108
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v102 == int32(1) {
		v184 = v99
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v116 = int32(1)
	v118 = int32(128)
	v119 = v106 << (uint(v116) % 32) & v118
	if v119 != v110<<(uint(v116)%32)&v118 {
		v185 = v119
		goto L35
	} else {
		goto L41
	}
L41:
	;
	if v102 < int32(3) {
		v184 = v99
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v127 = int32(2)
	v129 = int32(128)
	v130 = v106 << (uint(v127) % 32) & v129
	if v130 != v110<<(uint(v127)%32)&v129 {
		v185 = v130
		goto L35
	} else {
		goto L43
	}
L43:
	;
	if v102 == int32(3) {
		v184 = v99
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v138 = int32(3)
	v140 = int32(128)
	v141 = v106 << (uint(v138) % 32) & v140
	if v141 != v110<<(uint(v138)%32)&v140 {
		v185 = v141
		goto L35
	} else {
		goto L45
	}
L45:
	;
	if v102 < int32(5) {
		v184 = v99
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v149 = int32(4)
	v151 = int32(128)
	v152 = v106 << (uint(v149) % 32) & v151
	if v152 != v110<<(uint(v149)%32)&v151 {
		v185 = v152
		goto L35
	} else {
		goto L47
	}
L47:
	;
	if v102 == int32(5) {
		v184 = v99
		goto L36
	} else {
		goto L48
	}
L48:
	;
	v160 = int32(5)
	v162 = int32(128)
	v163 = v106 << (uint(v160) % 32) & v162
	if v163 != v110<<(uint(v160)%32)&v162 {
		v185 = v163
		goto L35
	} else {
		goto L49
	}
L49:
	;
	if v102 < int32(7) {
		v184 = v99
		goto L36
	} else {
		goto L50
	}
L50:
	;
	v171 = int32(6)
	v173 = int32(128)
	v174 = v106 << (uint(v171) % 32) & v173
	if v174 != v110<<(uint(v171)%32)&v173 {
		v185 = v174
		goto L35
	} else {
		goto L51
	}
L51:
	;
	v184 = v99
	goto L36
L52:
	;
	v188 = int32(1)
	goto L54
L53:
	;
	v188 = int32(-1)
	goto L54
L54:
	;
	v194 = v188
	goto L34
L55:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v198 = int32(1)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v203 = v200 & v198
	v204 = v197 & v198
	goto L33
L56:
	;
	v205 = v77
	goto L58
L57:
	;
	v205 = v79
	goto L58
L58:
	;
	v207 = v205 + int32(2)
	if v203 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v208 = v83
	goto L61
L60:
	;
	v208 = v85
	goto L61
L61:
	;
	v210 = v208 + int32(2)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if base.Ui32(v211) < base.Ui32(v81) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v213 = v211
	goto L64
L63:
	;
	v213 = v81
	goto L64
L64:
	;
	v214 = int32(0)
	v219 = int32(8)
	v220 = base.I32_div_s(v213, v219)
	if v219 <= v213 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v290)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
	v294 = F_cidr_set_masklen_internal(m, v16, v289)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L81
	}
L66:
	;
	v289 = v284 + v282<<(uint(int32(3))%32)
	goto L65
L67:
	;
	v270 = v261
	goto L78
L68:
	;
	v226 = v214
	goto L71
L69:
	;
	v243 = v214
	goto L70
L70:
	;
	v250 = v213 - v220<<(uint(int32(3))%32)
	if v250 == int32(0) {
		v282 = v243
		v284 = v214
		goto L66
	} else {
		goto L77
	}
L71:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v226))))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+v226))))
	if v232 != v234 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v243 = v220
	goto L70
L73:
	;
	v260 = v226
	v261 = int32(7)
	v263 = v232
	v264 = v234
	goto L67
L74:
	;
	goto L75
L75:
	;
	v238 = v226 + int32(1)
	if v238 != v220 {
		v226 = v238
		goto L71
	} else {
		goto L76
	}
L76:
	;
	goto L72
L77:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+v243))))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v243))))
	v260 = v243
	v261 = v250
	v263 = v256
	v264 = v254
	goto L67
L78:
	;
	if int32(base.Ui32(v263^v264)>>(uint(int32(8)-v270)%32)) != 0 {
		v270 = v270 - int32(1)
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v282 = v260
	v284 = v270
	goto L66
L80:
	;
	goto L79
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = int64(4)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v294
	v302 = int32(1)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v304&v302 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v307 = v302
	goto L84
L83:
	;
	v307 = int32(4)
	goto L84
L84:
	;
	v308 = v37 + v307
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	if v309 == int32(2) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v312 = int32(32)
	goto L87
L86:
	;
	v312 = int32(128)
	goto L87
L87:
	;
	if v289 < v312 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v315 = base.I32_div_s(v289, int32(8))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308+v315)+2)))
	v326 = int32(base.Ui32(v317)>>(uint(v315<<(uint(int32(3))%32)-v289+int32(7))%32)) & int32(1)
	goto L90
L89:
	;
	v326 = int32(0)
	goto L90
L90:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v328)
	if v289 < v327 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v333 = v326 | int32(2)
	goto L93
L92:
	;
	v333 = v326
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v333
	v382 = v37
	v393 = int32(28)
	goto L3
L94:
	;
	v346 = v336
	goto L96
L95:
	;
	v346 = int32(4)
	goto L96
L96:
	;
	v347 = v16 + v346
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	if v348 == int32(2) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v351 = int32(32)
	goto L99
L98:
	;
	v351 = int32(128)
	goto L99
L99:
	;
	if base.Ui32(v81) < base.Ui32(v351) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+int32(base.Ui32(v81)>>(uint(int32(3))%32)))+2)))
	v364 = int32(base.Ui32(v356)>>(uint((v81^int32(-1))&int32(7))%32)) & int32(1)
	goto L102
L101:
	;
	v364 = int32(0)
	goto L102
L102:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	if base.Ui32(v81) < base.Ui32(v367) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v369 = v364 | int32(2)
	goto L105
L104:
	;
	v369 = v364
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v369
	goto L4
}
func F_inet_spg_leaf_consistent(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = v7
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)) = uint8(v12)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v17 = F_inet_spg_consistent_bitmap(m, v7, v14, v15, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v17 != int32(0))
		}
	}
}
