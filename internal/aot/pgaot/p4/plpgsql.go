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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[1]))
	if v12 < v14 {
		v16 = v14 - v12
		v18 = v16 & int32(3)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[2]))
		if base.Ui32(int32(-4)) < base.Ui32(v12-v14) {
			v69 = int32(0)
			v70 = v12
			v79 = v69
			v80 = v70
			v87 = v2
			for {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v20+v80<<(uint(int32(2))%32))))
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				switch v92 {
				case 0, 2:
					v95 = v79 + int32(1)
				default:
					v95 = v79
				}
				v96 = int32(1)
				v99 = v87 + v96
				if v99 != v18 {
					v79 = v95
					v80 = v80 + v96
					v87 = v99
					continue
				} else {
					break
				}
				break
			}
			v102 = v95
		} else {
			v29 = int32(0)
			v30 = v12
			v36 = v2
			for {
				v40 = v20 + v30<<(uint(int32(2))%32)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				switch v42 {
				case 0, 2:
					v45 = v29 + int32(1)
				default:
					v45 = v29
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
				switch v47 {
				case 0, 2:
					v50 = v45 + int32(1)
				default:
					v50 = v45
				}
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
				switch v52 {
				case 0, 2:
					v55 = v50 + int32(1)
				default:
					v55 = v50
				}
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				switch v57 {
				case 0, 2:
					v60 = v55 + int32(1)
				default:
					v60 = v55
				}
				v61 = int32(4)
				v62 = v30 + v61
				v64 = v36 + v61
				if v64 != v16&int32(-4) {
					v29 = v60
					v30 = v62
					v36 = v64
					continue
				} else {
					break
				}
				break
			}
			if v18 == int32(0) {
				v102 = v60
			} else {
				v69 = v60
				v70 = v62
				v79 = v69
				v80 = v70
				v87 = v2
				for {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v20+v80<<(uint(int32(2))%32))))
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
					switch v92 {
					case 0, 2:
						v95 = v79 + int32(1)
					default:
						v95 = v79
					}
					v96 = int32(1)
					v99 = v87 + v96
					if v99 != v18 {
						v79 = v95
						v80 = v80 + v96
						v87 = v99
						continue
					} else {
						break
					}
					break
				}
				v102 = v95
			}
		}
		if l0 == int32(0) {
			v174 = v102
			v176 = v14
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[0])) = v176
			return v174
		} else {
			if v102 <= int32(0) {
				v162 = v102
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v174 = v162
				v176 = v14
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[0])) = v176
				return v174
			} else {
				v117 = F_palloc(m, v102<<(uint(int32(2))%32))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v117
					v122 = int32(0)
					v124 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[0]))
					v126 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[1]))
					if v126 <= v124 {
						v174 = v122
						v176 = v126
					} else {
						v129 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[2]))
						v131 = v122
						v132 = v124
						v133 = v126
						for {
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v129+v132<<(uint(int32(2))%32))))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							switch v144 {
							case 0, 2:
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v117+v131<<(uint(int32(2))%32)))) = v148
								v153 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[1]))
								v154 = v131 + int32(1)
								v155 = v153
							default:
								v154 = v131
								v155 = v133
							}
							v157 = v132 + int32(1)
							if v157 < v155 {
								v131 = v154
								v132 = v157
								v133 = v155
								continue
							} else {
								break
							}
							break
						}
						v174 = v154
						v176 = v155
					}
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[0])) = v176
					return v174
				}
			}
		}
	} else {
		if l0 == int32(0) {
			v174 = v2
			v176 = v14
		} else {
			v162 = v2
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v174 = v162
			v176 = v14
		}
		*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_add_initdatums[0])) = v176
		return v174
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[0]))
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[1]))
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[2]))
			if v25 == v27 {
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[2])) = v25 << (uint(int32(1)) % 32)
				v36 = F_repalloc(m, v23, v25<<(uint(int32(3))%32))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[0])) = v36
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[1]))
					v41 = v40
					v42 = v36
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v41
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[1])) = v41 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v42+v41<<(uint(int32(2))%32)))) = v8
					if l4 != 0 {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						F_plpgsql_ns_additem(m, int32(2), v41, v53)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						return v8
					}
				}
			} else {
				v41 = v25
				v42 = v23
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v41
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_record[1])) = v41 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v42+v41<<(uint(int32(2))%32)))) = v8
				if l4 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					F_plpgsql_ns_additem(m, int32(2), v41, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_extra_warnings_assign_hook[0])) = v4
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
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
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
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
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L30
	} else {
		goto L131
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L30
	} else {
		goto L128
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L30
	} else {
		goto L125
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L30
	} else {
		goto L122
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L30
	} else {
		goto L119
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L30
	} else {
		goto L116
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L30
	} else {
		goto L113
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L30
	} else {
		goto L110
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L30
	} else {
		goto L107
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L30
	} else {
		goto L104
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L30
	} else {
		goto L101
	}
L12:
	;
	v14 = int32(_a_F_plpgsql_fulfill_promise_1)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_fulfill_promise[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_fulfill_promise[0])) = v17
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
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_fulfill_promise[0])) = v15
	goto L14
L16:
	;
	F_assign_text_var(m, l0, l1, int32(_a_F_plpgsql_fulfill_promise_2))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L30
	} else {
		goto L100
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L30
	} else {
		goto L97
	}
L18:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v242 == int32(0) {
		goto L1
	} else {
		goto L93
	}
L19:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v232 == int32(0) {
		goto L2
	} else {
		goto L90
	}
L20:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v152 == int32(0) {
		goto L3
	} else {
		goto L70
	}
L21:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v143 == int32(0) {
		goto L4
	} else {
		goto L68
	}
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v127 == int32(0) {
		goto L5
	} else {
		goto L64
	}
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v112 == int32(0) {
		goto L6
	} else {
		goto L61
	}
L24:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v103 == int32(0) {
		goto L7
	} else {
		goto L59
	}
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v83 == int32(0) {
		goto L8
	} else {
		goto L50
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v63 == int32(0) {
		goto L9
	} else {
		goto L42
	}
L27:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v34 == int32(0) {
		goto L10
	} else {
		goto L33
	}
L28:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v28 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return
L31:
	;
	F_assign_simple_var(m, l0, l1, v28, int32(0), int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L15
L33:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v38 = int32(3)
	switch int32(base.Ui32(v37)>>(uint(v38)%32))&v38 - int32(1) {
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
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_fulfill_promise_0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L30
	} else {
		goto L39
	}
L35:
	;
	F_assign_text_var(m, l0, l1, int32(_a_F_plpgsql_fulfill_promise_3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L30
	} else {
		goto L38
	}
L36:
	;
	F_assign_text_var(m, l0, l1, int32(_a_F_plpgsql_fulfill_promise_4))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_5), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1422), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
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
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v66&int32(4) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v70 = F_cstring_to_text(m, int32(_a_F_plpgsql_fulfill_promise_8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L30
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v77 = F_cstring_to_text(m, int32(_a_F_plpgsql_fulfill_promise_9))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L30
	} else {
		goto L48
	}
L46:
	;
	F_assign_simple_var(m, l0, l1, v70, int32(0), int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	goto L15
L48:
	;
	F_assign_simple_var(m, l0, l1, v77, int32(0), int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L30
	} else {
		goto L49
	}
L49:
	;
	goto L15
L50:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	switch v86&int32(3) - int32(1) {
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
	F_assign_text_var(m, l0, l1, int32(_a_F_plpgsql_fulfill_promise_10))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L30
	} else {
		goto L58
	}
L52:
	;
	F_assign_text_var(m, l0, l1, int32(_a_F_plpgsql_fulfill_promise_11))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L30
	} else {
		goto L57
	}
L53:
	;
	F_assign_text_var(m, l0, l1, int32(_a_F_plpgsql_fulfill_promise_12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L30
	} else {
		goto L56
	}
L54:
	;
	F_assign_text_var(m, l0, l1, int32(_a_F_plpgsql_fulfill_promise_13))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+56))
	v108 = int32(0)
	F_assign_simple_var(m, l0, l1, v107, v108, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L30
	} else {
		goto L60
	}
L60:
	;
	goto L15
L61:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+48))
	v121 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v118+int32(4))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	F_assign_simple_var(m, l0, l1, v121, int32(0), int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	goto L15
L64:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+48))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+68))
	v135 = F_get_namespace_name(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L30
	} else {
		goto L65
	}
L65:
	;
	v137 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L66
	}
L66:
	;
	F_assign_simple_var(m, l0, l1, v137, int32(0), int32(1))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L30
	} else {
		goto L67
	}
L67:
	;
	goto L15
L68:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(v146)+34)))
	v148 = int32(0)
	F_assign_simple_var(m, l0, l1, v147, v148, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	goto L15
L70:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+34)))
	if int32(0) < v156 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v161 = F_palloc(m, v156<<(uint(int32(2))%32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L30
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
	if v204 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v167 = int32(0)
	goto L75
L75:
	;
	v172 = v167 << (uint(int32(2)) % 32)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+44))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176+v172)))
	v179 = F_cstring_to_text(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L30
	} else {
		goto L77
	}
L76:
	;
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v156
	v198 = F_construct_md_array(m, v161, v185, int32(1), v11+int32(12), v11+int32(8), int32(25), int32(-1), v185, int32(105))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L30
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161+v172))) = v179
	v183 = v167 + int32(1)
	if v183 != v156 {
		v167 = v183
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	F_assign_simple_var(m, l0, l1, v198, int32(0), int32(1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L30
	} else {
		goto L80
	}
L80:
	;
	goto L15
L81:
	;
	v226 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v226
	v228 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v228)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v226
	goto L15
L82:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	if v207 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_pfree(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L30
	} else {
		goto L89
	}
L84:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+12)))
	if v209 != int32(_a_F_plpgsql_fulfill_promise_14) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v213 != int32(1) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
	if v216 != int32(3) {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	F_DeleteExpandedObject(m, v212)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
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
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v236 = F_cstring_to_text(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L30
	} else {
		goto L91
	}
L91:
	;
	F_assign_simple_var(m, l0, l1, v236, int32(0), int32(1))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L30
	} else {
		goto L92
	}
L92:
	;
	goto L15
L93:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245<<(uint(int32(3))%32))+uint32(_c_F_plpgsql_fulfill_promise[1])))
	goto L94
L94:
	;
	v249 = F_cstring_to_text(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L30
	} else {
		goto L95
	}
L95:
	;
	F_assign_simple_var(m, l0, l1, v249, int32(0), int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L30
	} else {
		goto L96
	}
L96:
	;
	goto L15
L97:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v259
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_15), v11)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L30
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1532), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L30
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1405), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L30
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1414), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L30
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1427), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L30
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1438), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L30
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1453), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L30
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1461), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L30
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1470), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L30
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1479), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_16), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L30
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1487), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_17), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L30
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1521), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_fulfill_promise_17), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L30
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_plpgsql_fulfill_promise_6), int32(1527), int32(_a_F_plpgsql_fulfill_promise_7))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
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
	if v2 < v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v2
	v12 = v2
	goto L4
L2:
	;
	v25 = v2
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	F_mark_stmt(m, v28, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L9
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(72)+v12<<(uint(int32(2))%32))))
	v18 = F_bms_add_member(m, v11, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v25 = v18
	goto L3
L6:
	;
	return
L7:
	;
	v21 = v12 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v21 < v22 {
		v11 = v18
		v12 = v21
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	F_bms_free(m, v25)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v217 int32
	_ = v217
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v217
L2:
	;
	v217 = v203
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v195
	v203 = v193
	goto L2
L4:
	;
	v185 = int32(0)
	if l5 == v185 {
		v217 = v185
		goto L1
	} else {
		goto L56
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
	v78 = v14
	goto L10
L10:
	;
	if l3 == int32(0) {
		v170 = v78
		goto L25
	} else {
		goto L26
	}
L11:
	;
	v34 = v28 + int32(12)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v37 == int32(0))|base.B2i32(v37 != v40) != 0 {
		v58 = v37
		v59 = v40
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v78 = v72
	goto L10
L13:
	;
	v63 = int32(0)
	if v58-v59|base.B2i32(v30 == int32(1))&base.B2i32(l3 != v63) == v63 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	v43 = v34
	v44 = l2
	goto L16
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v48
		v59 = v47
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v58 = v48
	v59 = v47
	goto L14
L18:
	;
	v51 = int32(1)
	if v48 == v47 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if l5 == int32(0) {
		v203 = v28
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 != 0 {
		v28 = v72
		v30 = v73
		goto L11
	} else {
		goto L24
	}
L23:
	;
	v193 = v28
	v195 = int32(1)
	goto L3
L24:
	;
	goto L12
L25:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L54
	}
L26:
	;
	v86 = v78 + int32(12)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v89 == int32(0))|base.B2i32(v89 != v92) != 0 {
		v110 = v89
		v111 = v92
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v112 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v112 = v110 - v111
	goto L27
L29:
	;
	v95 = v86
	v96 = l2
	goto L30
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v100 == int32(0) {
		v110 = v100
		v111 = v99
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v110 = v100
	v111 = v99
	goto L28
L32:
	;
	v103 = int32(1)
	if v100 == v99 {
		v95 = v95 + v103
		v96 = v96 + v103
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v113 = v78
	goto L36
L35:
	;
	v113 = v14
	goto L36
L36:
	;
	if v112|base.B2i32(v23 == int32(0)) != 0 {
		v170 = v113
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v117 = v14
	v124 = v23
	goto L38
L38:
	;
	v127 = v117 + int32(12)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v130 == int32(0))|base.B2i32(v130 != v133) != 0 {
		v151 = v130
		v152 = v133
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v170 = v164
	goto L25
L40:
	;
	if v151-v152|base.B2i32(l4 != int32(0))&base.B2i32(v124 == int32(1)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v136 = v127
	v137 = l3
	goto L43
L43:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v141 == int32(0) {
		v151 = v141
		v152 = v140
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v151 = v141
	v152 = v140
	goto L41
L45:
	;
	v144 = int32(1)
	if v141 == v140 {
		v136 = v136 + v144
		v137 = v137 + v144
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if l5 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v165 != 0 {
		v117 = v164
		v124 = v165
		goto L38
	} else {
		goto L53
	}
L50:
	;
	return v117
L51:
	;
	goto L52
L52:
	;
	v193 = v117
	v195 = int32(2)
	goto L3
L53:
	;
	goto L39
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	if v175 != 0 {
		v14 = v175
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L7
L56:
	;
	v193 = v185
	v195 = v185
	goto L3
}
func F_plpgsql_param_eval_generic_ro(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12<<(uint(int32(2))%32)-int32(4))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_exec_eval_datum(m, v10, v18, v7+int32(28), v7+int32(24), v23, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		if v27 == v28 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			if v33 != 0 {
				v45 = v31
				v46 = v30
			} else {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
				if v34 != int32(1) {
					v43 = v31
				} else {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
					if v37 != int32(3) {
						v43 = v31
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+2))
						v43 = v40 + int32(18)
					}
				}
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v45 = v43
				v46 = v44
			}
			*(*int32)(unsafe.Add(mBase, uint32(v46))) = v45
			m.G0 = v7 + int32(32)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_param_eval_generic_ro_0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					v59 = F_format_type_be(m, v27)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v62 = F_format_type_be(m, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v62
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v59
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v58
							F_errmsg(m, int32(_a_F_plpgsql_param_eval_generic_ro_1), v7)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_param_eval_generic_ro_2), int32(_a_F_plpgsql_param_eval_generic_ro_3), int32(_a_F_plpgsql_param_eval_generic_ro_4))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v15 = l1 - int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15<<(uint(int32(2))%32))))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v24 = F_bms_is_member(m, v15, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = int32(0)
		if base.B2i32(l2 == int32(0))|base.B2i32(v24 == v28) == v28 {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			switch v33 {
			case 0, 1, 2, 4:
				v73 = l3 + int32(4)
				F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v73)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v76)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					switch v78 {
					case 0:
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
						if v79 == int32(0) {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
							v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)))
							if v83 == int32(_a_F_plpgsql_param_fetch_0) {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
								if v89 != int32(1) {
									v98 = v88
								} else {
									v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
									if v92 != int32(3) {
										v98 = v88
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+2))
										v98 = v95 + int32(18)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v98
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
							}
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
						}
					default:
					case 2:
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
						if v101 != 0 {
							v112 = v100
						} else {
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
							if v102 != int32(1) {
								v111 = v100
							} else {
								v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
								if v105 != int32(3) {
									v111 = v100
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)+2))
									v111 = v108 + int32(18)
								}
							}
							v112 = v111
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v112
					}
					m.G0 = v10 + int32(16)
					return l3
				}
			case 3:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35<<(uint(int32(2))%32))))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
				if v40 == int32(0) {
					v58 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v58
					*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v58)
					v62 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v62)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
					m.G0 = v10 + int32(16)
					return l3
				} else {
					v43 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v40)+48))
					if v43 == v44 {
						v73 = l3 + int32(4)
						F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v73)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v76)
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							switch v78 {
							case 0:
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
								if v79 == int32(0) {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
									v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)))
									if v83 == int32(_a_F_plpgsql_param_fetch_0) {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
										if v89 != int32(1) {
											v98 = v88
										} else {
											v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
											if v92 != int32(3) {
												v98 = v88
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+2))
												v98 = v95 + int32(18)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v98
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
									}
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
								}
							default:
							case 2:
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
								if v101 != 0 {
									v112 = v100
								} else {
									v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
									if v102 != int32(1) {
										v111 = v100
									} else {
										v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
										if v105 != int32(3) {
											v111 = v100
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)+2))
											v111 = v108 + int32(18)
										}
									}
									v112 = v111
								}
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v112
							}
							m.G0 = v10 + int32(16)
							return l3
						}
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
						v49 = F_expanded_record_lookup_field(m, v40, v46, v19+int32(32))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 == int32(0) {
								v58 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v58
								*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v58)
								v62 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v62)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
								m.G0 = v10 + int32(16)
								return l3
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
								v54 = *(*int64)(unsafe.Add(mBase, uint32(v53)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v54
								v73 = l3 + int32(4)
								F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v73)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v76 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v76)
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
									switch v78 {
									case 0:
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
										if v79 == int32(0) {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
											v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)))
											if v83 == int32(_a_F_plpgsql_param_fetch_0) {
												v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
												if v89 != int32(1) {
													v98 = v88
												} else {
													v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
													if v92 != int32(3) {
														v98 = v88
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+2))
														v98 = v95 + int32(18)
													}
												}
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v98
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
											}
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
										}
									default:
									case 2:
										v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
										if v101 != 0 {
											v112 = v100
										} else {
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
											if v102 != int32(1) {
												v111 = v100
											} else {
												v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
												if v105 != int32(3) {
													v111 = v100
												} else {
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)+2))
													v111 = v108 + int32(18)
												}
											}
											v112 = v111
										}
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v112
									}
									m.G0 = v10 + int32(16)
									return l3
								}
							}
						}
					}
				}
			default:
				v58 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v58
				*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v58)
				v62 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v62)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
				m.G0 = v10 + int32(16)
				return l3
			}
		} else {
			if v24 != 0 {
				v73 = l3 + int32(4)
				F_exec_eval_datum(m, v12, v19, l3+int32(8), v10+int32(12), l3, v73)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v76)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					switch v78 {
					case 0:
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
						if v79 == int32(0) {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
							v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)))
							if v83 == int32(_a_F_plpgsql_param_fetch_0) {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
								if v89 != int32(1) {
									v98 = v88
								} else {
									v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
									if v92 != int32(3) {
										v98 = v88
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+2))
										v98 = v95 + int32(18)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v98
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
							}
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
						}
					default:
					case 2:
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
						if v101 != 0 {
							v112 = v100
						} else {
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
							if v102 != int32(1) {
								v111 = v100
							} else {
								v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
								if v105 != int32(3) {
									v111 = v100
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)+2))
									v111 = v108 + int32(18)
								}
							}
							v112 = v111
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v112
					}
					m.G0 = v10 + int32(16)
					return l3
				}
			} else {
				v58 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v58
				*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v58)
				v62 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v62)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
				m.G0 = v10 + int32(16)
				return l3
			}
		}
	}
}
func F_plpgsql_parse_dblword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = F_makeString(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v11
	v16 = F_makeString(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v20
	v26 = F_list_make2_impl(m, v9+int32(16), v9+int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_dblword[0]))
	if v29 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v9 + int32(32)
	return v225
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v26
	v225 = int32(0)
	goto L5
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_dblword[1]))
	v37 = v9 + int32(28)
	if v33 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v182 == int32(0) {
		goto L6
	} else {
		goto L44
	}
L9:
	;
	v182 = v172
	goto L8
L10:
	;
	v172 = v158
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v150
	v158 = v148
	goto L10
L12:
	;
	v140 = int32(0)
	if v37 == v140 {
		v172 = v140
		goto L9
	} else {
		goto L43
	}
L13:
	;
	v45 = v33
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L12
L16:
	;
	v59 = v45
	v61 = v54
	goto L19
L17:
	;
	v84 = v45
	goto L18
L18:
	;
	if l1 == int32(0) {
		v125 = v84
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v66 = F_strcmp(m, v59+int32(12), l0)
	mBase = m.M
	v69 = int32(0)
	if v66|base.B2i32(v61 == int32(1))&base.B2i32(l1 != v69) == v69 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v84 = v78
	goto L18
L21:
	;
	if v37 == int32(0) {
		v158 = v59
		goto L10
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 != 0 {
		v59 = v78
		v61 = v79
		goto L19
	} else {
		goto L25
	}
L24:
	;
	v148 = v59
	v150 = int32(1)
	goto L11
L25:
	;
	goto L20
L26:
	;
	goto L41
L27:
	;
	v93 = F_strcmp(m, v84+int32(12), l0)
	mBase = m.M
	if v93 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v94 = v84
	goto L30
L29:
	;
	v94 = v45
	goto L30
L30:
	;
	if v93|base.B2i32(v54 == int32(0)) != 0 {
		v125 = v94
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v98 = v45
	v105 = v54
	goto L32
L32:
	;
	v109 = F_strcmp(m, v98+int32(12), l1)
	mBase = m.M
	if v109|int32(0)&base.B2i32(v105 == int32(1)) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v125 = v119
	goto L26
L34:
	;
	if v37 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v120 != 0 {
		v98 = v119
		v105 = v120
		goto L32
	} else {
		goto L40
	}
L37:
	;
	v182 = v98
	goto L8
L38:
	;
	goto L39
L39:
	;
	v148 = v98
	v150 = int32(2)
	goto L11
L40:
	;
	goto L33
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	if v130 != 0 {
		v45 = v130
		goto L14
	} else {
		goto L42
	}
L42:
	;
	goto L15
L43:
	;
	v148 = v140
	v150 = v140
	goto L11
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	switch v185 - int32(1) {
	case 0:
		goto L46
	case 1:
		goto L45
	default:
		goto L6
	}
L45:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_dblword[2]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203+v204<<(uint(int32(2))%32))))
	v209 = int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v210 == v209 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_dblword[2]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v26
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v196)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v194
	v225 = int32(1)
	goto L5
L47:
	;
	v213 = F_plpgsql_build_recfield(m, v208, l1)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v215 = v208
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v26
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v217)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v215
	v225 = v209
	goto L5
L50:
	;
	v215 = v213
	goto L49
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v113 int32
	_ = v113
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
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
	return v214
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l0
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(base.B2i32(v207 == int32(34)))
	v214 = int32(0)
	goto L1
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_word[0]))
	if v14 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_word[1]))
	if v16 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if v155 == int32(0) {
		goto L2
	} else {
		goto L41
	}
L6:
	;
	goto L5
L7:
	;
	v155 = v42
	goto L6
L9:
	;
	v155 = int32(0)
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
	v67 = v28
	goto L15
L15:
	;
	goto L23
L16:
	;
	v49 = F_strcmp(m, v42+int32(12), l0)
	mBase = m.M
	if v49|base.B2i32(v44 == int32(1))&int32(0) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v67 = v61
	goto L15
L18:
	;
	goto L7
L19:
	;
	goto L20
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != 0 {
		v42 = v61
		v44 = v62
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L17
L23:
	;
	goto L38
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v113 != 0 {
		v28 = v113
		goto L11
	} else {
		goto L39
	}
L39:
	;
	goto L12
L41:
	;
	v168 = int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if base.Ui32(v169-v168) <= base.Ui32(v168) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_word[2]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175+v176<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v180
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(base.B2i32(v183 == int32(34)))
	v214 = v168
	goto L1
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_word_0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v195
	F_errmsg_internal(m, int32(_a_F_plpgsql_parse_word_1), v9)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_word_2), int32(1328), int32(_a_F_plpgsql_parse_word_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
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
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_wordrowtype_0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_wordrowtype_0))
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_wordrowtype[0]))
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
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_wordrowtype_0))
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
	F_errmsg(m, int32(_a_F_plpgsql_parse_wordrowtype_1), v8)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_wordrowtype_2), int32(1683), int32(_a_F_plpgsql_parse_wordrowtype_3))
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
	F_errmsg(m, int32(_a_F_plpgsql_parse_wordrowtype_4), v8+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_wordrowtype_2), int32(1691), int32(_a_F_plpgsql_parse_wordrowtype_3))
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
	F_errmsg_internal(m, int32(_a_F_plpgsql_parse_wordrowtype_5), v8+int32(32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_wordrowtype_2), int32(1960), int32(_a_F_plpgsql_parse_wordrowtype_6))
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v106 int32
	_ = v106
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_wordtype[0]))
	if v9 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_wordtype[1]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183+v184<<(uint(int32(2))%32))))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+24))
	m.G0 = v6 + int32(16)
	return v189
L2:
	;
	if v148 != 0 {
		goto L38
	} else {
		goto L39
	}
L3:
	;
	goto L2
L4:
	;
	v148 = v35
	goto L3
L6:
	;
	v148 = int32(0)
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
	v60 = v21
	goto L12
L12:
	;
	goto L20
L13:
	;
	v42 = F_strcmp(m, v35+int32(12), l0)
	mBase = m.M
	if v42|base.B2i32(v37 == int32(1))&int32(0) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v60 = v54
	goto L12
L15:
	;
	goto L4
L16:
	;
	goto L17
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v55 != 0 {
		v35 = v54
		v37 = v55
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	goto L35
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v106 != 0 {
		v21 = v106
		goto L8
	} else {
		goto L36
	}
L36:
	;
	goto L9
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if base.Ui32(v159-int32(1)) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_wordtype_0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	return int32(0)
L43:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(_a_F_plpgsql_parse_wordtype_1), v6)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_wordtype_2), int32(1541), int32(_a_F_plpgsql_parse_wordtype_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_pre_column_ref(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+488))
	if v6 == int32(1) {
		v10 = F_resolve_column_ref(m, l0, v4, l1, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = v10
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
		v70 = int32(_a_F_plpgsql_stmt_typename_0)
		return v70
	case 1:
		return int32(_a_F_plpgsql_stmt_typename_1)
	case 2:
		return int32(_a_F_plpgsql_stmt_typename_2)
	case 3:
		return int32(_a_F_plpgsql_stmt_typename_3)
	case 4:
		return int32(_a_F_plpgsql_stmt_typename_4)
	case 5:
		return int32(_a_F_plpgsql_stmt_typename_5)
	case 6:
		return int32(_a_F_plpgsql_stmt_typename_6)
	case 7:
		return int32(_a_F_plpgsql_stmt_typename_7)
	case 8:
		return int32(_a_F_plpgsql_stmt_typename_8)
	case 9:
		return int32(_a_F_plpgsql_stmt_typename_9)
	case 10:
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if v25 != 0 {
			v26 = int32(_a_F_plpgsql_stmt_typename_10)
		} else {
			v26 = int32(_a_F_plpgsql_stmt_typename_11)
		}
		return v26
	case 11:
		return int32(_a_F_plpgsql_stmt_typename_12)
	case 12:
		return int32(_a_F_plpgsql_stmt_typename_13)
	case 13:
		return int32(_a_F_plpgsql_stmt_typename_14)
	case 14:
		return int32(_a_F_plpgsql_stmt_typename_15)
	case 15:
		return int32(_a_F_plpgsql_stmt_typename_16)
	case 16:
		return int32(_a_F_plpgsql_stmt_typename_17)
	case 17:
		return int32(_a_F_plpgsql_stmt_typename_18)
	case 18:
		return int32(_a_F_plpgsql_stmt_typename_19)
	case 19:
		v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if v46 != 0 {
			v47 = int32(_a_F_plpgsql_stmt_typename_20)
		} else {
			v47 = int32(_a_F_plpgsql_stmt_typename_21)
		}
		return v47
	case 20:
		return int32(_a_F_plpgsql_stmt_typename_22)
	case 21:
		v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v53 != 0 {
			v54 = int32(_a_F_plpgsql_stmt_typename_23)
		} else {
			v54 = int32(_a_F_plpgsql_stmt_typename_24)
		}
		return v54
	case 22:
		return int32(_a_F_plpgsql_stmt_typename_25)
	case 23:
		return int32(_a_F_plpgsql_stmt_typename_26)
	case 24:
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v62 != 0 {
			v63 = int32(_a_F_plpgsql_stmt_typename_27)
		} else {
			v63 = int32(_a_F_plpgsql_stmt_typename_28)
		}
		return v63
	case 25:
		return int32(_a_F_plpgsql_stmt_typename_29)
	case 26:
		return int32(_a_F_plpgsql_stmt_typename_30)
	default:
		v70 = int32(_a_F_plpgsql_stmt_typename_31)
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
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_token_is_unreserved_keyword[0]))))
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v12 + v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 == int32(0) {
		F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyerror_0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
				F_errmsg(m, int32(_a_F_plpgsql_yyerror_1), v9)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v30 = F_plpgsql_scanner_errposition(m, v29, l2)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_plpgsql_yyerror_2), int32(544), int32(_a_F_plpgsql_yyerror_3))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
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
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v39 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14+v37))) = uint8(v39)
		F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyerror_0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
				F_errmsg(m, int32(_a_F_plpgsql_yyerror_4), v9+int32(16))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v56 = F_plpgsql_scanner_errposition(m, v55, l2)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_plpgsql_yyerror_2), int32(560), int32(_a_F_plpgsql_yyerror_3))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
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
