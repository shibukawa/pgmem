package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_plpgsql_add_initdatums(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	v2 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1461]))
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1459]))
	if v17 < v20 {
		v22 = v20 - v17
		v24 = v22 & int32(3)
		v26 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
		if base.Ui32(int32(-4)) < base.Ui32(v17-v20) {
			v85 = int32(0)
			v86 = v17
		} else {
			v41 = int32(0)
			v42 = v17
			v48 = v2
			for {
				v54 = v42 << (uint(int32(2)) % 32)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+v54)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				switch v57 {
				case 0, 2:
					v60 = v41 + int32(1)
				default:
					v60 = v41
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v26+int32(4)))))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
				switch v63 {
				case 0, 2:
					v66 = v60 + int32(1)
				default:
					v66 = v60
				}
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v26+int32(8)))))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
				switch v69 {
				case 0, 2:
					v72 = v66 + int32(1)
				default:
					v72 = v66
				}
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v26+int32(12)))))
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
				switch v75 {
				case 0, 2:
					v78 = v72 + int32(1)
				default:
					v78 = v72
				}
				v79 = int32(4)
				v80 = v42 + v79
				v82 = v48 + v79
				if v82 != v22&int32(-4) {
					v41 = v78
					v42 = v80
					v48 = v82
					continue
				} else {
					break
				}
				break
			}
			v85 = v78
			v86 = v80
		}
		if v24 != 0 {
			v98 = v85
			v99 = v86
			v103 = v2
			for {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v26+v99<<(uint(int32(2))%32))))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
				switch v114 {
				case 0, 2:
					v117 = v98 + int32(1)
				default:
					v117 = v98
				}
				v118 = int32(1)
				v121 = v103 + v118
				if v121 != v24 {
					v98 = v117
					v99 = v99 + v118
					v103 = v121
					continue
				} else {
					break
				}
				break
			}
			v124 = v117
		} else {
			v124 = v85
		}
		if l0 == int32(0) {
			v208 = v124
			v212 = v20
			*(*int32)(unsafe.Add(mBase, _consts[1461])) = v212
			return v208
		} else {
			if v124 <= int32(0) {
				v193 = v124
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v208 = v193
				v212 = v20
				*(*int32)(unsafe.Add(mBase, _consts[1461])) = v212
				return v208
			} else {
				v142 = F_palloc(m, v124<<(uint(int32(2))%32))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v142
					v147 = int32(0)
					v151 = *(*int32)(unsafe.Add(mBase, _consts[1461]))
					v154 = *(*int32)(unsafe.Add(mBase, _consts[1459]))
					if v154 <= v151 {
						v208 = v147
						v212 = v154
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
						v159 = v147
						v160 = v151
						v163 = v154
						for {
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v157+v160<<(uint(int32(2))%32))))
							v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
							switch v175 {
							case 0, 2:
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v142+v159<<(uint(int32(2))%32)))) = v179
								v184 = *(*int32)(unsafe.Add(mBase, _consts[1459]))
								v185 = v159 + int32(1)
								v186 = v184
							default:
								v185 = v159
								v186 = v163
							}
							v188 = v160 + int32(1)
							if v188 < v186 {
								v159 = v185
								v160 = v188
								v163 = v186
								continue
							} else {
								break
							}
							break
						}
						v208 = v185
						v212 = v186
					}
					*(*int32)(unsafe.Add(mBase, _consts[1461])) = v212
					return v208
				}
			}
		}
	} else {
		if l0 == int32(0) {
			v208 = v2
			v212 = v20
		} else {
			v193 = v2
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v208 = v193
			v212 = v20
		}
		*(*int32)(unsafe.Add(mBase, _consts[1461])) = v212
		return v208
	}
}
func F_plpgsql_build_record(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v8 = F_palloc0(m, int32(40))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(2)
		v14 = F_pstrdup(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(4294967295)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14
			v25 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
			v28 = *(*int32)(unsafe.Add(mBase, _consts[1459]))
			v31 = *(*int32)(unsafe.Add(mBase, _consts[1460]))
			if v28 == v31 {
				*(*int32)(unsafe.Add(mBase, _consts[1460])) = v28 << (uint(int32(1)) % 32)
				v43 = F_repalloc(m, v25, v28<<(uint(int32(3))%32))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[1458])) = v43
					v48 = *(*int32)(unsafe.Add(mBase, _consts[1459]))
					v49 = v48
					v50 = v43
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v49
					*(*int32)(unsafe.Add(mBase, _consts[1459])) = v49 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v50+v49<<(uint(int32(2))%32)))) = v8
					if l4 != 0 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						F_plpgsql_ns_additem(m, int32(2), v49, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						return v8
					}
				}
			} else {
				v49 = v28
				v50 = v25
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v49
				*(*int32)(unsafe.Add(mBase, _consts[1459])) = v49 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v50+v49<<(uint(int32(2))%32)))) = v8
				if l4 != 0 {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					F_plpgsql_ns_additem(m, int32(2), v49, v61)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						return v8
					}
				} else {
					return v8
				}
			}
		}
	}
}
func F_plpgsql_extra_warnings_assign_hook(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[1463])) = v4
	return
}
func F_plpgsql_fulfill_promise(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v13 != 0 {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L30
	} else {
		goto L131
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L30
	} else {
		goto L128
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L30
	} else {
		goto L125
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L30
	} else {
		goto L122
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L30
	} else {
		goto L119
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L30
	} else {
		goto L116
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L30
	} else {
		goto L113
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L30
	} else {
		goto L110
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L30
	} else {
		goto L107
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L30
	} else {
		goto L104
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L30
	} else {
		goto L101
	}
L12:
	;
	v14 = int32(4464496)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v16
	switch v13 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	case 5:
		goto L23
	case 6:
		goto L22
	case 7:
		goto L21
	case 8:
		goto L20
	case 9:
		goto L19
	case 10:
		goto L18
	default:
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	m.G0 = v11 + int32(16)
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v15
	goto L14
L16:
	;
	F_assign_text_var(m, l0, l1, int32(527628))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L30
	} else {
		goto L100
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L30
	} else {
		goto L97
	}
L18:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v247 == int32(0) {
		goto L1
	} else {
		goto L93
	}
L19:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v237 == int32(0) {
		goto L2
	} else {
		goto L90
	}
L20:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v156 == int32(0) {
		goto L3
	} else {
		goto L70
	}
L21:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v147 == int32(0) {
		goto L4
	} else {
		goto L68
	}
L22:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v131 == int32(0) {
		goto L5
	} else {
		goto L64
	}
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v116 == int32(0) {
		goto L6
	} else {
		goto L61
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v107 == int32(0) {
		goto L7
	} else {
		goto L59
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v87 == int32(0) {
		goto L8
	} else {
		goto L50
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 == int32(0) {
		goto L9
	} else {
		goto L42
	}
L27:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v33 == int32(0) {
		goto L10
	} else {
		goto L33
	}
L28:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return
L31:
	;
	F_assign_simple_var(m, l0, l1, v27, int32(0), int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L15
L33:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v37 = int32(3)
	switch int32(base.Ui32(v36)>>(uint(v37)%32))&v37 - int32(1) {
	case 0:
		goto L16
	case 1:
		goto L35
	case 2:
		goto L34
	default:
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L30
	} else {
		goto L39
	}
L35:
	;
	F_assign_text_var(m, l0, l1, int32(525144))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L30
	} else {
		goto L38
	}
L36:
	;
	F_assign_text_var(m, l0, l1, int32(513701))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	goto L15
L38:
	;
	goto L15
L39:
	;
	F_errmsg_internal(m, int32(525085), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(488778), int32(1422), int32(353298))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L30
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)))
	if v70&int32(4) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v74 = F_cstring_to_text(m, int32(504552))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L30
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v81 = F_cstring_to_text(m, int32(506707))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L30
	} else {
		goto L48
	}
L46:
	;
	F_assign_simple_var(m, l0, l1, v74, int32(0), int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	goto L15
L48:
	;
	F_assign_simple_var(m, l0, l1, v81, int32(0), int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L30
	} else {
		goto L49
	}
L49:
	;
	goto L15
L50:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	switch v90&int32(3) - int32(1) {
	case 0:
		goto L52
	case 1:
		goto L53
	case 2:
		goto L51
	default:
		goto L54
	}
L51:
	;
	F_assign_text_var(m, l0, l1, int32(527036))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L30
	} else {
		goto L58
	}
L52:
	;
	F_assign_text_var(m, l0, l1, int32(526348))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L30
	} else {
		goto L57
	}
L53:
	;
	F_assign_text_var(m, l0, l1, int32(527014))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L30
	} else {
		goto L56
	}
L54:
	;
	F_assign_text_var(m, l0, l1, int32(505937))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L30
	} else {
		goto L55
	}
L55:
	;
	goto L15
L56:
	;
	goto L15
L57:
	;
	goto L15
L58:
	;
	goto L15
L59:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+56))
	v112 = int32(0)
	F_assign_simple_var(m, l0, l1, v111, v112, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L30
	} else {
		goto L60
	}
L60:
	;
	goto L15
L61:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+48))
	v125 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v122+int32(4))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	F_assign_simple_var(m, l0, l1, v125, int32(0), int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	goto L15
L64:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+48))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+68))
	v139 = F_get_namespace_name(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L30
	} else {
		goto L65
	}
L65:
	;
	v141 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L30
	} else {
		goto L66
	}
L66:
	;
	F_assign_simple_var(m, l0, l1, v141, int32(0), int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L30
	} else {
		goto L67
	}
L67:
	;
	goto L15
L68:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+34)))
	v152 = int32(0)
	F_assign_simple_var(m, l0, l1, v151, v152, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	goto L15
L70:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v159)+34)))
	if int32(0) < v160 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v166 = F_palloc(m, v160<<(uint(int32(2))%32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L30
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
	if v209 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v170 = int32(0)
	goto L75
L75:
	;
	v177 = v170 << (uint(int32(2)) % 32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+44))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v177)))
	v184 = F_cstring_to_text(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L30
	} else {
		goto L77
	}
L76:
	;
	v190 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v160
	v203 = F_construct_md_array(m, v166, v190, int32(1), v11+int32(12), v11+int32(8), int32(25), int32(-1), v190, int32(105))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L30
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166+v177))) = v184
	v188 = v170 + int32(1)
	if v188 != v160 {
		v170 = v188
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	F_assign_simple_var(m, l0, l1, v203, int32(0), int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L30
	} else {
		goto L80
	}
L80:
	;
	goto L15
L81:
	;
	v231 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v231
	v233 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v233)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v231
	goto L15
L82:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	if v212 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_pfree(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L30
	} else {
		goto L89
	}
L84:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+12)))
	if v214 != int32(65535) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v218 != int32(1) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v221 != int32(3) {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	F_DeleteExpandedObject(m, v217)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L30
	} else {
		goto L88
	}
L88:
	;
	goto L81
L89:
	;
	goto L81
L90:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	v241 = F_cstring_to_text(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L30
	} else {
		goto L91
	}
L91:
	;
	F_assign_simple_var(m, l0, l1, v241, int32(0), int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L30
	} else {
		goto L92
	}
L92:
	;
	goto L15
L93:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250<<(uint(int32(3))%32))+uint32(_consts[495])))
	goto L94
L94:
	;
	v256 = F_cstring_to_text(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L30
	} else {
		goto L95
	}
L95:
	;
	F_assign_simple_var(m, l0, l1, v256, int32(0), int32(1))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L30
	} else {
		goto L96
	}
L96:
	;
	goto L15
L97:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v268
	F_errmsg_internal(m, int32(474802), v11)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L30
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(488778), int32(1532), int32(353298))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L30
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	goto L15
L101:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L30
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(488778), int32(1405), int32(353298))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L30
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L30
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(488778), int32(1414), int32(353298))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L30
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L30
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(488778), int32(1427), int32(353298))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L30
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L30
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(488778), int32(1438), int32(353298))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L30
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L30
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(488778), int32(1453), int32(353298))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L30
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L30
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(488778), int32(1461), int32(353298))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L30
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L30
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(488778), int32(1470), int32(353298))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L30
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L30
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(488778), int32(1479), int32(353298))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L30
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errmsg_internal(m, int32(247051), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L30
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(488778), int32(1487), int32(353298))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L30
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errmsg_internal(m, int32(246878), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L30
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(488778), int32(1521), int32(353298))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L30
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errmsg_internal(m, int32(246878), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L30
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(488778), int32(1527), int32(353298))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L30
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_mark_local_assignment_targets(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v5 <= v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	F_mark_stmt(m, v28, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L10
	}
L2:
	;
	v25 = v2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v11 = v2
	v12 = v2
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(72)+v12<<(uint(int32(2))%32))))
	v18 = F_bms_add_member(m, v11, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v25 = v18
	goto L1
L7:
	;
	return
L8:
	;
	v21 = v12 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v21 < v22 {
		v11 = v18
		v12 = v21
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	F_bms_free(m, v25)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	return
}
func F_plpgsql_ns_lookup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v207 int32
	_ = v207
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v207
L2:
	;
	v207 = v193
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v185
	v193 = v183
	goto L2
L4:
	;
	v175 = int32(0)
	if l5 == v175 {
		v207 = v175
		goto L1
	} else {
		goto L60
	}
L5:
	;
	v14 = l0
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	v28 = v14
	v30 = v23
	goto L11
L9:
	;
	v74 = v14
	goto L10
L10:
	;
	if l3 == int32(0) {
		v160 = v74
		goto L26
	} else {
		goto L27
	}
L11:
	;
	v34 = v28 + int32(12)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v38 == int32(0) {
		v57 = v37
		v58 = v38
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v74 = v68
	goto L10
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 != 0 {
		v28 = v68
		v30 = v69
		goto L11
	} else {
		goto L25
	}
L14:
	;
	if v58-v57 != 0 {
		goto L13
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v37 != v38 {
		v57 = v37
		v58 = v38
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v42 = v34
	v43 = l2
	goto L18
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v46
		v58 = v47
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v57 = v46
	v58 = v47
	goto L15
L20:
	;
	v50 = int32(1)
	if v46 == v47 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if base.B2i32(v30 == int32(1))&base.B2i32(l3 != int32(0)) != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	if l5 == int32(0) {
		v193 = v28
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v183 = v28
	v185 = int32(1)
	goto L3
L25:
	;
	goto L12
L26:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L58
	}
L27:
	;
	v82 = v74 + int32(12)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v86 == int32(0) {
		v105 = v85
		v106 = v86
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v107 != 0 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v107 = v106 - v105
	goto L28
L30:
	;
	if v85 != v86 {
		v105 = v85
		v106 = v86
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v90 = v82
	v91 = l2
	goto L32
L32:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v95 == int32(0) {
		v105 = v94
		v106 = v95
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v105 = v94
	v106 = v95
	goto L29
L34:
	;
	v98 = int32(1)
	if v94 == v95 {
		v90 = v90 + v98
		v91 = v91 + v98
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v108 = v74
	goto L38
L37:
	;
	v108 = v14
	goto L38
L38:
	;
	if v107 != 0 {
		v160 = v108
		goto L26
	} else {
		goto L39
	}
L39:
	;
	if v23 == int32(0) {
		v160 = v108
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v111 = v14
	v118 = v23
	goto L41
L41:
	;
	v121 = v111 + int32(12)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v125 == int32(0) {
		v144 = v124
		v145 = v125
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v160 = v154
	goto L26
L43:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v155 != 0 {
		v111 = v154
		v118 = v155
		goto L41
	} else {
		goto L57
	}
L44:
	;
	if v145-v144 != 0 {
		goto L43
	} else {
		goto L52
	}
L45:
	;
	goto L44
L46:
	;
	if v124 != v125 {
		v144 = v124
		v145 = v125
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v129 = v121
	v130 = l3
	goto L48
L48:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	if v134 == int32(0) {
		v144 = v133
		v145 = v134
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v144 = v133
	v145 = v134
	goto L45
L50:
	;
	v137 = int32(1)
	if v133 == v134 {
		v129 = v129 + v137
		v130 = v130 + v137
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if base.B2i32(l4 != int32(0))&base.B2i32(v118 == int32(1)) != 0 {
		goto L43
	} else {
		goto L53
	}
L53:
	;
	if l5 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return v111
L55:
	;
	goto L56
L56:
	;
	v183 = v111
	v185 = int32(2)
	goto L3
L57:
	;
	goto L42
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v165 != 0 {
		v14 = v165
		goto L6
	} else {
		goto L59
	}
L59:
	;
	goto L7
L60:
	;
	v183 = v175
	v185 = v175
	goto L3
}
func F_plpgsql_param_eval_generic_ro(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13<<(uint(int32(2))%32)-int32(4))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_exec_eval_datum(m, v11, v19, v8+int32(28), v8+int32(24), v24, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		if v28 == v29 {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
			if v34 != 0 {
				v46 = v32
				v47 = v31
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
				if v35 != int32(1) {
					v44 = v32
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
					if v38 != int32(3) {
						v44 = v32
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+2))
						v44 = v41 + int32(18)
					}
				}
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v46 = v44
				v47 = v45
			}
			*(*int32)(unsafe.Add(mBase, uint32(v47))) = v46
			m.G0 = v8 + int32(32)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(537468))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					v62 = F_format_type_be(m, v28)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v65 = F_format_type_be(m, v64)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v62
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
							F_errmsg(m, int32(646091), v8)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								F_errfinish(m, int32(488778), int32(6883), int32(235046))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
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
func F_plpgsql_param_eval_var_transfer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32)-int32(4))))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+44)))
	if v17 != 0 {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v40))) = v16
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v17)
		return
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
		if v18 != int32(1) {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = v16
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v17)
			return
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
			if v21 != int32(3) {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v40))) = v16
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v17)
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+120))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+2))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
				F_MemoryContextSetParent(m, v27, v25)
				mBase = m.M
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v26 + int32(12)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v34)
				v36 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v15)+44)) = uint16(v36)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v34
				return
			}
		}
	}
}
func F_plpgsql_param_fetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v15 = l1 - int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15<<(uint(int32(2))%32))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v22 = F_bms_is_member(m, v15, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if l2 == int32(0) {
			if v22 != 0 {
				v70 = l3 + int32(4)
				F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v70)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v73)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					switch v75 {
					case 0:
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
						if v76 == int32(0) {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
							v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+12)))
							if v80 == int32(65535) {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
								if v86 != int32(1) {
									v95 = v85
								} else {
									v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
									if v89 != int32(3) {
										v95 = v85
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+2))
										v95 = v92 + int32(18)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v95
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
							}
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
						}
					default:
					case 2:
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
						if v98 != 0 {
							v109 = v97
						} else {
							v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
							if v99 != int32(1) {
								v108 = v97
							} else {
								v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
								if v102 != int32(3) {
									v108 = v97
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+2))
									v108 = v105 + int32(18)
								}
							}
							v109 = v108
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
					}
					m.G0 = v10 + int32(16)
					return l3
				}
			} else {
				v55 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v55
				*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v55)
				v59 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v59)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
				m.G0 = v10 + int32(16)
				return l3
			}
		} else {
			if v22 == int32(0) {
				if v22 != 0 {
					v70 = l3 + int32(4)
					F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v73)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						switch v75 {
						case 0:
							v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
							if v76 == int32(0) {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
								v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+12)))
								if v80 == int32(65535) {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
									if v86 != int32(1) {
										v95 = v85
									} else {
										v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
										if v89 != int32(3) {
											v95 = v85
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+2))
											v95 = v92 + int32(18)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v95
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
								}
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
							}
						default:
						case 2:
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
							if v98 != 0 {
								v109 = v97
							} else {
								v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
								if v99 != int32(1) {
									v108 = v97
								} else {
									v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
									if v102 != int32(3) {
										v108 = v97
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+2))
										v108 = v105 + int32(18)
									}
								}
								v109 = v108
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
						}
						m.G0 = v10 + int32(16)
						return l3
					}
				} else {
					v55 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v55
					*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v55)
					v59 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v59)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
					m.G0 = v10 + int32(16)
					return l3
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				switch v30 {
				case 0, 1, 2, 4:
					v70 = l3 + int32(4)
					F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v73)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						switch v75 {
						case 0:
							v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
							if v76 == int32(0) {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
								v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+12)))
								if v80 == int32(65535) {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
									if v86 != int32(1) {
										v95 = v85
									} else {
										v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
										if v89 != int32(3) {
											v95 = v85
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+2))
											v95 = v92 + int32(18)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v95
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
								}
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
							}
						default:
						case 2:
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
							if v98 != 0 {
								v109 = v97
							} else {
								v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
								if v99 != int32(1) {
									v108 = v97
								} else {
									v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
									if v102 != int32(3) {
										v108 = v97
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+2))
										v108 = v105 + int32(18)
									}
								}
								v109 = v108
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
						}
						m.G0 = v10 + int32(16)
						return l3
					}
				case 3:
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v31+v32<<(uint(int32(2))%32))))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
					if v37 == int32(0) {
						v55 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v55
						*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v55)
						v59 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v59)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
						m.G0 = v10 + int32(16)
						return l3
					} else {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v37)+48))
						if v40 == v41 {
							v70 = l3 + int32(4)
							F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v73)
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
								switch v75 {
								case 0:
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
									if v76 == int32(0) {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
										v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+12)))
										if v80 == int32(65535) {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
											if v86 != int32(1) {
												v95 = v85
											} else {
												v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
												if v89 != int32(3) {
													v95 = v85
												} else {
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+2))
													v95 = v92 + int32(18)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v95
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
										}
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
									}
								default:
								case 2:
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
									if v98 != 0 {
										v109 = v97
									} else {
										v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
										if v99 != int32(1) {
											v108 = v97
										} else {
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
											if v102 != int32(3) {
												v108 = v97
											} else {
												v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+2))
												v108 = v105 + int32(18)
											}
										}
										v109 = v108
									}
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
								}
								m.G0 = v10 + int32(16)
								return l3
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
							v46 = F_expanded_record_lookup_field(m, v37, v43, v19+int32(32))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if v46 == int32(0) {
									v55 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v55
									*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v55)
									v59 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v59)
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
									m.G0 = v10 + int32(16)
									return l3
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
									v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v51
									v70 = l3 + int32(4)
									F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										v73 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v73)
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
										switch v75 {
										case 0:
											v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
											if v76 == int32(0) {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
												v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+12)))
												if v80 == int32(65535) {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
													v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
													if v86 != int32(1) {
														v95 = v85
													} else {
														v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
														if v89 != int32(3) {
															v95 = v85
														} else {
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+2))
															v95 = v92 + int32(18)
														}
													}
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = v95
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
												}
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
											}
										default:
										case 2:
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
											if v98 != 0 {
												v109 = v97
											} else {
												v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
												if v99 != int32(1) {
													v108 = v97
												} else {
													v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
													if v102 != int32(3) {
														v108 = v97
													} else {
														v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+2))
														v108 = v105 + int32(18)
													}
												}
												v109 = v108
											}
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
										}
										m.G0 = v10 + int32(16)
										return l3
									}
								}
							}
						}
					}
				default:
					v55 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v55
					*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v55)
					v59 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v59)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v55
					m.G0 = v10 + int32(16)
					return l3
				}
			}
		}
	}
}
func F_plpgsql_parse_dblword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = F_makeString(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v12
	v17 = F_makeString(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v21
	v27 = F_list_make2_impl(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[1456]))
	if v30 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v10 + int32(32)
	return v219
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v27
	v219 = int32(0)
	goto L5
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v38 = v10 + int32(28)
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v176 == int32(0) {
		goto L6
	} else {
		goto L45
	}
L9:
	;
	v176 = v166
	goto L8
L10:
	;
	v166 = v152
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v144
	v152 = v142
	goto L10
L12:
	;
	v134 = int32(0)
	if v38 == v134 {
		v166 = v134
		goto L9
	} else {
		goto L44
	}
L13:
	;
	v46 = v34
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v55 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L12
L16:
	;
	v60 = v46
	v62 = v55
	goto L19
L17:
	;
	v82 = v46
	goto L18
L18:
	;
	if l1 == int32(0) {
		v119 = v82
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v67 = F_strcmp(m, v60+int32(12), l0)
	mBase = m.M
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v82 = v76
	goto L18
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 != 0 {
		v60 = v76
		v62 = v77
		goto L19
	} else {
		goto L25
	}
L22:
	;
	if base.B2i32(v62 == int32(1))&base.B2i32(l1 != int32(0)) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v38 == int32(0) {
		v152 = v60
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v142 = v60
	v144 = int32(1)
	goto L11
L25:
	;
	goto L20
L26:
	;
	goto L42
L27:
	;
	v91 = F_strcmp(m, v82+int32(12), l0)
	mBase = m.M
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v92 = v82
	goto L30
L29:
	;
	v92 = v46
	goto L30
L30:
	;
	if v91 != 0 {
		v119 = v92
		goto L26
	} else {
		goto L31
	}
L31:
	;
	if v55 == int32(0) {
		v119 = v92
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v95 = v46
	v102 = v55
	goto L33
L33:
	;
	v106 = F_strcmp(m, v95+int32(12), l1)
	mBase = m.M
	if v106 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v119 = v113
	goto L26
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v114 != 0 {
		v95 = v113
		v102 = v114
		goto L33
	} else {
		goto L41
	}
L36:
	;
	if int32(0)&base.B2i32(v102 == int32(1)) != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if v38 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v176 = v95
	goto L8
L39:
	;
	goto L40
L40:
	;
	v142 = v95
	v144 = int32(2)
	goto L11
L41:
	;
	goto L34
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	if v124 != 0 {
		v46 = v124
		goto L14
	} else {
		goto L43
	}
L43:
	;
	goto L15
L44:
	;
	v142 = v134
	v144 = v134
	goto L11
L45:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	switch v179 - int32(1) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L6
	}
L46:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197+v198<<(uint(int32(2))%32))))
	v203 = int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v204 == v203 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183+v184<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v27
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v188
	v219 = int32(1)
	goto L5
L48:
	;
	v207 = F_plpgsql_build_recfield(m, v202, l1)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v209 = v202
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v27
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v211)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v209
	v219 = v203
	goto L5
L51:
	;
	v209 = v207
	goto L50
}
func F_plpgsql_parse_word(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v106 int32
	_ = v106
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v207
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l0
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(base.B2i32(v200 == int32(34)))
	v207 = int32(0)
	goto L1
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1456]))
	if v14 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	if v16 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if v148 == int32(0) {
		goto L2
	} else {
		goto L42
	}
L6:
	;
	goto L5
L7:
	;
	v148 = v42
	goto L6
L9:
	;
	v148 = int32(0)
	goto L6
L10:
	;
	v28 = v16
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L9
L13:
	;
	v42 = v28
	v44 = v37
	goto L16
L14:
	;
	v64 = v28
	goto L15
L15:
	;
	goto L23
L16:
	;
	v49 = F_strcmp(m, v42+int32(12), l0)
	mBase = m.M
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v64 = v58
	goto L15
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v59 != 0 {
		v42 = v58
		v44 = v59
		goto L16
	} else {
		goto L22
	}
L19:
	;
	if base.B2i32(v44 == int32(1))&int32(0) != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	goto L7
L22:
	;
	goto L17
L23:
	;
	goto L39
L39:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v106 != 0 {
		v28 = v106
		goto L11
	} else {
		goto L40
	}
L40:
	;
	goto L12
L42:
	;
	v161 = int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if base.Ui32(v162-v161) <= base.Ui32(v161) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v168+v169<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v173
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(base.B2i32(v176 == int32(34)))
	v207 = v161
	goto L1
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return int32(0)
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v188
	F_errmsg_internal(m, int32(473305), v9)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(484732), int32(1328), int32(411041))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_parse_wordrowtype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	F_recomputeNamespacePath(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L28
	}
L3:
	;
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	return int32(0)
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	if v15 == int32(0) {
		v43 = v2
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v49 = v43
	goto L3
L7:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= v18 {
		v49 = v18
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v25 = v2
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v25<<(uint(int32(2))%32))))
	v32 = F_get_relname_relid(m, l0, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v49 = int32(0)
	goto L3
L11:
	;
	if v32 != 0 {
		v43 = v32
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v35 = v25 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v35 < v36 {
		v25 = v35
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v50 = F_get_rel_type_id(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L24
	}
L17:
	;
	if v50 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v54 = F_makeTypeName(m, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v57 = F_SearchSysCache1(m, int32(82), v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if v57 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v63 = F_build_datatype(m, v57, int32(-1), int32(0), v54)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_ReleaseCatCache(m, v57)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	m.G0 = v8 + int32(48)
	return v63
L24:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg(m, int32(70091), v8)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(484732), int32(1683), int32(357402))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	F_errmsg(m, int32(361683), v8+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(484732), int32(1691), int32(357402))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v50
	F_errmsg_internal(m, int32(49363), v8+int32(32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(484732), int32(1960), int32(358007))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_parse_wordtype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v99 int32
	_ = v99
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	if v9 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176+v177<<(uint(int32(2))%32))))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	m.G0 = v6 + int32(16)
	return v182
L2:
	;
	if v141 != 0 {
		goto L39
	} else {
		goto L40
	}
L3:
	;
	goto L2
L4:
	;
	v141 = v35
	goto L3
L6:
	;
	v141 = int32(0)
	goto L3
L7:
	;
	v21 = v9
	goto L8
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v30 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	v35 = v21
	v37 = v30
	goto L13
L11:
	;
	v57 = v21
	goto L12
L12:
	;
	goto L20
L13:
	;
	v42 = F_strcmp(m, v35+int32(12), l0)
	mBase = m.M
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v57 = v51
	goto L12
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 != 0 {
		v35 = v51
		v37 = v52
		goto L13
	} else {
		goto L19
	}
L16:
	;
	if base.B2i32(v37 == int32(1))&int32(0) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	goto L4
L19:
	;
	goto L14
L20:
	;
	goto L36
L36:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v99 != 0 {
		v21 = v99
		goto L8
	} else {
		goto L37
	}
L37:
	;
	goto L9
L39:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if base.Ui32(v152-int32(1)) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_errstart_cold(m, int32(21), int32(537468))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	return int32(0)
L44:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(70968), v6)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(484732), int32(1541), int32(357910))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_pre_column_ref(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+488))
	if v7 == int32(1) {
		v11 = F_resolve_column_ref(m, l0, v5, l1, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			return v15
		}
	} else {
		v15 = int32(0)
		return v15
	}
}
func F_plpgsql_scanner_finish(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_scanner_finish(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_plpgsql_stmt_typename(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 {
	case 0:
		v70 = int32(309431)
		return v70
	case 1:
		return int32(92941)
	case 2:
		return int32(525167)
	case 3:
		return int32(527485)
	case 4:
		return int32(515057)
	case 5:
		return int32(528440)
	case 6:
		return int32(387792)
	case 7:
		return int32(111068)
	case 8:
		return int32(206096)
	case 9:
		return int32(24288)
	case 10:
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if v25 != 0 {
			v26 = int32(508955)
		} else {
			v26 = int32(525692)
		}
		return v26
	case 11:
		return int32(516682)
	case 12:
		return int32(505323)
	case 13:
		return int32(497085)
	case 14:
		return int32(527468)
	case 15:
		return int32(505721)
	case 16:
		return int32(93826)
	case 17:
		return int32(525991)
	case 18:
		return int32(93840)
	case 19:
		v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if v46 != 0 {
			v47 = int32(512157)
		} else {
			v47 = int32(512103)
		}
		return v47
	case 20:
		return int32(518974)
	case 21:
		v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v53 != 0 {
			v54 = int32(525578)
		} else {
			v54 = int32(524150)
		}
		return v54
	case 22:
		return int32(527355)
	case 23:
		return int32(519503)
	case 24:
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v62 != 0 {
			v63 = int32(521715)
		} else {
			v63 = int32(515842)
		}
		return v63
	case 25:
		return int32(509298)
	case 26:
		return int32(522305)
	default:
		v70 = int32(238283)
		return v70
	}
}
func F_plpgsql_token_is_unreserved_keyword(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v5 = int32(0)
	goto L1
L1:
	;
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5<<(uint(int32(1))%32))+uint32(_consts[1465]))))
	v12 = base.B2i32(l0 == v11)
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v12
L3:
	;
	v16 = v5 + int32(1)
	if v16 != int32(83) {
		v5 = v16
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	goto L5
}
func F_plpgsql_token_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+64))
	return v3
}
func F_plpgsql_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v12 + v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 == int32(0) {
		F_errstart_cold(m, int32(21), int32(537468))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
				F_errmsg(m, int32(63732), v9)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = F_plpgsql_scanner_errposition(m, v32, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(484215), int32(544), int32(206161))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14+v42))) = uint8(v44)
		F_errstart_cold(m, int32(21), int32(537468))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
				F_errmsg(m, int32(673167), v9+int32(16))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v64 = F_plpgsql_scanner_errposition(m, v63, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(484215), int32(560), int32(206161))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
