package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_numeric_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v30 = int32(49152)
			v31 = v29 & v30
			if v31 == v30 {
				if v29 != int32(53248) {
					if v29 != int32(49152) {
						if v28 != int32(61440) {
							v50 = int32(-1)
						} else {
							v50 = int32(0)
						}
						v169 = v50
					} else {
						v169 = base.B2i32(v28 != int32(49152))
					}
				} else {
					if v28 == int32(49152) {
						v45 = int32(-1)
					} else {
						v45 = base.B2i32(v28 != int32(53248))
					}
					v169 = v45
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v28) {
					if v28 == int32(61440) {
						v57 = int32(1)
					} else {
						v57 = int32(-1)
					}
					v169 = v57
				} else {
					v59 = v7 + int32(6)
					v64 = base.B2i32(int32(0) <= base.I32_extend16_s(v29))
					if int32(0) <= base.I32_extend16_s(v29) {
						v65 = int32(-8)
					} else {
						v65 = int32(-6)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if int32(0) <= base.I32_extend16_s(v29) {
						v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
						v79 = v69
					} else {
						v79 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = v65 + int32(base.Ui32(v66)>>(uint(int32(2))%32))
					v82 = v14 + int32(6)
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v28))
					if int32(0) <= base.I32_extend16_s(v28) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if int32(0) <= base.I32_extend16_s(v28) {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
					}
					v103 = v88 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
					v109 = v28 & int32(49152)
					if v109 == int32(32768) {
						v112 = v28 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v112 = v109
					}
					if base.Ui32(v80) <= base.Ui32(int32(1)) {
						if base.Ui32(v103) < base.Ui32(int32(2)) {
							v169 = int32(0)
						} else {
							if v112 == int32(16384) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v169 = v122
						}
					} else {
						if v31 == int32(32768) {
							v129 = v29 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v129 = v31
						}
						if base.Ui32(v103) <= base.Ui32(int32(1)) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v169 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v29) {
								v137 = v7 + int32(8)
							} else {
								v137 = v59
							}
							v139 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v28) {
								v142 = v14 + int32(8)
							} else {
								v142 = v82
							}
							v144 = int32(base.Ui32(v103) >> (uint(int32(1)) % 32))
							if v129 == int32(0) {
								if v112 == int32(16384) {
									v169 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v137, v139, v79, v142, v144, v102)
									mBase = m.M
									v169 = v150
								}
							} else {
								if v112 == int32(0) {
									v169 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v142, v144, v102, v137, v139, v79)
									mBase = m.M
									v169 = v154
								}
							}
						}
					}
				}
			}
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v170 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v174 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							return v169
						}
					} else {
						return v169
					}
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v174 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						return v169
					}
				} else {
					return v169
				}
			}
		}
	}
}
func F_numeric_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v30 = int32(49152)
			v31 = v29 & v30
			if v31 == v30 {
				if v29 != int32(53248) {
					if v29 != int32(49152) {
						if v28 != int32(61440) {
							v50 = int32(-1)
						} else {
							v50 = int32(0)
						}
						v169 = v50
					} else {
						v169 = base.B2i32(v28 != int32(49152))
					}
				} else {
					if v28 == int32(49152) {
						v45 = int32(-1)
					} else {
						v45 = base.B2i32(v28 != int32(53248))
					}
					v169 = v45
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v28) {
					if v28 == int32(61440) {
						v57 = int32(1)
					} else {
						v57 = int32(-1)
					}
					v169 = v57
				} else {
					v59 = v7 + int32(6)
					v64 = base.B2i32(int32(0) <= base.I32_extend16_s(v29))
					if int32(0) <= base.I32_extend16_s(v29) {
						v65 = int32(-8)
					} else {
						v65 = int32(-6)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if int32(0) <= base.I32_extend16_s(v29) {
						v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
						v79 = v69
					} else {
						v79 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = v65 + int32(base.Ui32(v66)>>(uint(int32(2))%32))
					v82 = v14 + int32(6)
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v28))
					if int32(0) <= base.I32_extend16_s(v28) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if int32(0) <= base.I32_extend16_s(v28) {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
					}
					v103 = v88 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
					v109 = v28 & int32(49152)
					if v109 == int32(32768) {
						v112 = v28 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v112 = v109
					}
					if base.Ui32(v80) <= base.Ui32(int32(1)) {
						if base.Ui32(v103) < base.Ui32(int32(2)) {
							v169 = int32(0)
						} else {
							if v112 == int32(16384) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v169 = v122
						}
					} else {
						if v31 == int32(32768) {
							v129 = v29 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v129 = v31
						}
						if base.Ui32(v103) <= base.Ui32(int32(1)) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v169 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v29) {
								v137 = v7 + int32(8)
							} else {
								v137 = v59
							}
							v139 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v28) {
								v142 = v14 + int32(8)
							} else {
								v142 = v82
							}
							v144 = int32(base.Ui32(v103) >> (uint(int32(1)) % 32))
							if v129 == int32(0) {
								if v112 == int32(16384) {
									v169 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v137, v139, v79, v142, v144, v102)
									mBase = m.M
									v169 = v150
								}
							} else {
								if v112 == int32(0) {
									v169 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v142, v144, v102, v137, v139, v79)
									mBase = m.M
									v169 = v154
								}
							}
						}
					}
				}
			}
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v170 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v174 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v169 == int32(0))
						}
					} else {
						return base.B2i32(v169 == int32(0))
					}
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v174 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v169 == int32(0))
					}
				} else {
					return base.B2i32(v169 == int32(0))
				}
			}
		}
	}
}
func F_numeric_floor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = int32(base.Ui32(v20) >> (uint(int32(2)) % 32))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v23) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(48)
	return v353
L4:
	;
	v26 = F_palloc(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v35 = base.I32_extend16_s(v23)
	v37 = base.B2i32(int32(0) <= v35)
	if int32(0) <= v35 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v30 = int32(base.Ui32(v28) >> (uint(int32(2)) % 32))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v353 = v26
	goto L3
L9:
	;
	v31 = F__emscripten_memcpy_bulkmem(m, v26, v16, v30)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v38 = int32(-8)
	goto L14
L13:
	;
	v38 = int32(-6)
	goto L14
L14:
	;
	v39 = v22 + v38
	v41 = int32(base.Ui32(v39) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
	if int32(0) <= v35 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
	v53 = v43
	goto L17
L16:
	;
	v53 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v53
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v55
	v60 = base.B2i32(v35 < v55)
	if v35 < v55 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = int32(6)
	goto L20
L19:
	;
	v61 = int32(8)
	goto L20
L20:
	;
	v62 = v16 + v61
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v62
	if v35 < v55 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = int32(base.Ui32(v23)>>(uint(int32(7))%32)) & int32(63)
	goto L23
L22:
	;
	v70 = v23 & int32(16383)
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v70
	v77 = v23 & int32(49152)
	if v77 == int32(32768) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = v23 << (uint(int32(1)) % 32) & int32(16384)
	goto L26
L25:
	;
	v80 = v77
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v80
	v83 = v39 & int32(-2)
	v86 = F_palloc(m, v83+int32(2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v88 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v88)
	if base.Ui32(int32(2)) <= base.Ui32(v39) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v83 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v86
	v101 = int32(2)
	v102 = v86 + v101
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v105
	v108 = v104 << (uint(v101) % 32)
	if v108+int32(4) <= v105 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	goto L30
L32:
	;
	v94 = F__emscripten_memcpy_bulkmem(m, v86+int32(2), v62, v83)
	mBase = m.M
	goto L34
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v127
	if v80 != int32(16384) {
		v321 = v127
		goto L42
	} else {
		goto L43
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = int64(0)
	v115 = int32(0)
	v125 = v115
	v127 = v115
	goto L35
L37:
	;
	goto L38
L38:
	;
	v120 = base.I32_div_s(v108+int32(7), int32(4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v120 < v121 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v123 = v120
	goto L41
L40:
	;
	v123 = v121
	goto L41
L41:
	;
	v125 = v104
	v127 = v123
	goto L35
L42:
	;
	v323 = v321 << (uint(int32(1)) % 32)
	v326 = F_palloc(m, v323+int32(2))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L100
	}
L43:
	;
	if base.Ui32(v39) <= base.Ui32(int32(1)) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v314 = v13 + int32(24)
	F_sub_var(m, v314, int32(1770600), v314)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L99
	}
L45:
	;
	if v127 != 0 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v127 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v321 = int32(0)
	goto L42
L49:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	if v136 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v139 = int32(0)
	if base.B2i32(v53 < v125)&base.B2i32(v139 < v127) == v139 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	if v310 == int32(0) {
		v321 = v127
		goto L42
	} else {
		goto L98
	}
L52:
	;
	v310 = v300
	goto L51
L53:
	;
	if v53 <= v170 {
		v205 = v53
		v207 = v139
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v170 = v125
	v174 = v139
	goto L53
L55:
	;
	goto L56
L56:
	;
	v151 = v125
	v155 = v139
	goto L57
L57:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v155<<(uint(int32(1))%32)))))
	if v161 != 0 {
		v300 = int32(1)
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v170 = v165
	v174 = v163
	goto L53
L59:
	;
	v162 = int32(1)
	v163 = v155 + v162
	v165 = v151 - v162
	if v165 <= v53 {
		v170 = v165
		v174 = v163
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v163 < v127 {
		v151 = v165
		v155 = v163
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	if v170 != v205 {
		v246 = v174
		v247 = v207
		goto L70
	} else {
		goto L71
	}
L63:
	;
	if v41 <= int32(0) {
		v205 = v53
		v207 = v139
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v186 = v53
	v188 = v139
	goto L65
L65:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v188<<(uint(int32(1))%32)))))
	if v193 != 0 {
		v300 = int32(-1)
		goto L52
	} else {
		goto L67
	}
L66:
	;
	v205 = v197
	v207 = v195
	goto L62
L67:
	;
	v194 = int32(1)
	v195 = v188 + v194
	v197 = v186 - v194
	if v197 <= v170 {
		v205 = v197
		v207 = v195
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v195 < v41 {
		v186 = v197
		v188 = v195
		goto L65
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	if v127 < v246 {
		goto L80
	} else {
		goto L81
	}
L71:
	;
	v216 = v174
	v217 = v207
	goto L72
L72:
	;
	if v127 <= v216 {
		v246 = v216
		v247 = v217
		goto L70
	} else {
		goto L74
	}
L73:
	;
	if base.I32_extend16_s(v232) < base.I32_extend16_s(v230) {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	if v41 <= v217 {
		v246 = v216
		v247 = v217
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v221 = int32(1)
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v216<<(uint(v221)%32)))))
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217<<(uint(v221)%32)+v62))))
	if v230 == v232 {
		v216 = v216 + v221
		v217 = v217 + v221
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	v239 = int32(1)
	goto L79
L78:
	;
	v239 = int32(-1)
	goto L79
L79:
	;
	v310 = v239
	goto L51
L80:
	;
	v250 = v246
	goto L82
L81:
	;
	v250 = v127
	goto L82
L82:
	;
	v257 = v246
	goto L83
L83:
	;
	if v250 == v257 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v300 = v283
	goto L52
L85:
	;
	if v41 < v247 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v283 = int32(1)
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v257<<(uint(v283)%32)))))
	if v289 == int32(0) {
		v257 = v257 + v283
		goto L83
	} else {
		goto L97
	}
L88:
	;
	v262 = v247
	goto L90
L89:
	;
	v262 = v41
	goto L90
L90:
	;
	v270 = v247
	goto L91
L91:
	;
	if v262 == v270 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v300 = int32(-1)
	goto L52
L93:
	;
	v310 = int32(0)
	goto L51
L94:
	;
	goto L95
L95:
	;
	v274 = int32(1)
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v270<<(uint(v274)%32)))))
	if v279 == int32(0) {
		v270 = v270 + v274
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	goto L84
L98:
	;
	goto L44
L99:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v321 = v320
	goto L42
L100:
	;
	v328 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v326))) = uint16(v328)
	if v328 < v321 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v323 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v337
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v326 + int32(2)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v345 != 0 {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	goto L103
L105:
	;
	v335 = F__emscripten_memcpy_bulkmem(m, v326+int32(2), v334, v323)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	F_pfree(m, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v349 = F_make_result_opt_error(m, v13, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	F_pfree(m, v326)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v353 = v349
	goto L3
}
func F_numeric_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v30 = int32(49152)
			v31 = v29 & v30
			if v31 == v30 {
				if v29 != int32(53248) {
					if v29 != int32(49152) {
						if v28 != int32(61440) {
							v50 = int32(-1)
						} else {
							v50 = int32(0)
						}
						v169 = v50
					} else {
						v169 = base.B2i32(v28 != int32(49152))
					}
				} else {
					if v28 == int32(49152) {
						v45 = int32(-1)
					} else {
						v45 = base.B2i32(v28 != int32(53248))
					}
					v169 = v45
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v28) {
					if v28 == int32(61440) {
						v57 = int32(1)
					} else {
						v57 = int32(-1)
					}
					v169 = v57
				} else {
					v59 = v7 + int32(6)
					v64 = base.B2i32(int32(0) <= base.I32_extend16_s(v29))
					if int32(0) <= base.I32_extend16_s(v29) {
						v65 = int32(-8)
					} else {
						v65 = int32(-6)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if int32(0) <= base.I32_extend16_s(v29) {
						v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
						v79 = v69
					} else {
						v79 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = v65 + int32(base.Ui32(v66)>>(uint(int32(2))%32))
					v82 = v14 + int32(6)
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v28))
					if int32(0) <= base.I32_extend16_s(v28) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if int32(0) <= base.I32_extend16_s(v28) {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
					}
					v103 = v88 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
					v109 = v28 & int32(49152)
					if v109 == int32(32768) {
						v112 = v28 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v112 = v109
					}
					if base.Ui32(v80) <= base.Ui32(int32(1)) {
						if base.Ui32(v103) < base.Ui32(int32(2)) {
							v169 = int32(0)
						} else {
							if v112 == int32(16384) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v169 = v122
						}
					} else {
						if v31 == int32(32768) {
							v129 = v29 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v129 = v31
						}
						if base.Ui32(v103) <= base.Ui32(int32(1)) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v169 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v29) {
								v137 = v7 + int32(8)
							} else {
								v137 = v59
							}
							v139 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v28) {
								v142 = v14 + int32(8)
							} else {
								v142 = v82
							}
							v144 = int32(base.Ui32(v103) >> (uint(int32(1)) % 32))
							if v129 == int32(0) {
								if v112 == int32(16384) {
									v169 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v137, v139, v79, v142, v144, v102)
									mBase = m.M
									v169 = v150
								}
							} else {
								if v112 == int32(0) {
									v169 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v142, v144, v102, v137, v139, v79)
									mBase = m.M
									v169 = v154
								}
							}
						}
					}
				}
			}
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v170 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v174 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v169^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v169^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v174 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						return int32(base.Ui32(v169^int32(-1)) >> (uint(int32(31)) % 32))
					}
				} else {
					return int32(base.Ui32(v169^int32(-1)) >> (uint(int32(31)) % 32))
				}
			}
		}
	}
}
func F_numeric_int8_opt_error(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l1 == int32(0) {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		if base.Ui32(v14) <= base.Ui32(int32(49151)) {
			v48 = v14
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v54 = base.I32_extend16_s(v48)
			v56 = base.B2i32(int32(0) <= v54)
			if int32(0) <= v54 {
				v57 = int32(-8)
			} else {
				v57 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(base.Ui32(int32(base.Ui32(v49)>>(uint(int32(2))%32))+v57) >> (uint(int32(1)) % 32))
			if int32(0) <= v54 {
				v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v72 = v62
			} else {
				v72 = v48<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v48&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v72
			v74 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v74
			v83 = base.B2i32(v54 < v74)
			if v54 < v74 {
				v84 = int32(base.Ui32(v48)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v84 = v48 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v84
			v91 = v48 & int32(49152)
			if v91 == int32(32768) {
				v94 = v48 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v94 = v91
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v94
			if v54 < v74 {
				v98 = int32(6)
			} else {
				v98 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = l0 + v98
			v105 = F_numericvar_to_int64(m, v8+int32(-24), v8+int32(-32))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int64(0)
			} else {
				if v105 == int32(0) {
					if l1 != 0 {
						v109 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v109)
						v132 = int64(0)
						m.G0 = v10 - int32(-64)
						return v132
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(420400), int32(0))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(523892), int32(4646), int32(222451))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int64(0)
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
					v128 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
					v132 = v128
					m.G0 = v10 - int32(-64)
					return v132
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					if v14 == int32(49152) {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(95661)
						F_errmsg(m, int32(193475), v10)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(523892), int32(4624), int32(222451))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(95661)
						F_errmsg(m, int32(192798), v8+int32(-48))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(523892), int32(4628), int32(222451))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
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
	} else {
		v40 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v40)
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		if base.Ui32(v42) < base.Ui32(int32(49152)) {
			v48 = v42
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v54 = base.I32_extend16_s(v48)
			v56 = base.B2i32(int32(0) <= v54)
			if int32(0) <= v54 {
				v57 = int32(-8)
			} else {
				v57 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(base.Ui32(int32(base.Ui32(v49)>>(uint(int32(2))%32))+v57) >> (uint(int32(1)) % 32))
			if int32(0) <= v54 {
				v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v72 = v62
			} else {
				v72 = v48<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v48&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v72
			v74 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v74
			v83 = base.B2i32(v54 < v74)
			if v54 < v74 {
				v84 = int32(base.Ui32(v48)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v84 = v48 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v84
			v91 = v48 & int32(49152)
			if v91 == int32(32768) {
				v94 = v48 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v94 = v91
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v94
			if v54 < v74 {
				v98 = int32(6)
			} else {
				v98 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = l0 + v98
			v105 = F_numericvar_to_int64(m, v8+int32(-24), v8+int32(-32))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int64(0)
			} else {
				if v105 == int32(0) {
					if l1 != 0 {
						v109 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v109)
						v132 = int64(0)
						m.G0 = v10 - int32(-64)
						return v132
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(420400), int32(0))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(523892), int32(4646), int32(222451))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int64(0)
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
					v128 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
					v132 = v128
					m.G0 = v10 - int32(-64)
					return v132
				}
			}
		} else {
			v45 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v45)
			v132 = int64(0)
			m.G0 = v10 - int32(-64)
			return v132
		}
	}
}
func F_numeric_is_inf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	return base.B2i32(v2&int32(57343) == int32(53248))
}
func F_numeric_log(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v22 = base.I32_extend16_s(v21)
			if base.Ui32(v21) <= base.Ui32(int32(49151)) {
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
				v26 = base.I32_extend16_s(v25)
				if base.Ui32(int32(49151)) < base.Ui32(v25) {
					v54 = v26
					if v54&int32(65535) != int32(49152) {
						if v22 == int32(-12288) {
							v69 = int32(1)
						} else {
							v69 = int32(-1)
						}
						v70 = int32(49152)
						v71 = v21 & v70
						if v71 == v70 {
							v98 = v69
						} else {
							v74 = int32(0)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							if v74 <= v22 {
								v82 = int32(-8)
							} else {
								v82 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v75)>>(uint(int32(2))%32))+v82) < base.Ui32(int32(2)) {
								v98 = v74
							} else {
								v87 = int32(1)
								if v71 == int32(32768) {
									v94 = v21 << (uint(v87) % 32) & int32(16384)
								} else {
									v94 = v71
								}
								if v94 == int32(16384) {
									v97 = int32(-1)
								} else {
									v97 = v87
								}
								v98 = v97
							}
						}
						v100 = v54 & int32(65535)
						if v100 == int32(53248) {
							v105 = int32(1)
						} else {
							v105 = int32(-1)
						}
						v106 = int32(49152)
						v107 = v54 & v106
						if v107 == v106 {
							v135 = v105
						} else {
							v110 = int32(0)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							if v110 <= base.I32_extend16_s(v54) {
								v119 = int32(-8)
							} else {
								v119 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v111)>>(uint(int32(2))%32))+v119) < base.Ui32(int32(2)) {
								v135 = v110
							} else {
								v124 = int32(1)
								if v107 == int32(32768) {
									v131 = v100 << (uint(v124) % 32) & int32(16384)
								} else {
									v131 = v107
								}
								if v131 == int32(16384) {
									v134 = int32(-1)
								} else {
									v134 = v124
								}
								v135 = v134
							}
						}
						if v98 < int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v349 = m.ExcPending
							if v349 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(352583810))
								mBase = m.M
								v352 = m.ExcPending
								if v352 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(239569), int32(0))
									mBase = m.M
									v356 = m.ExcPending
									if v356 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523892), int32(4008), int32(343390))
										mBase = m.M
										v361 = m.ExcPending
										if v361 != 0 {
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
							if v135 < int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v349 = m.ExcPending
								if v349 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(352583810))
									mBase = m.M
									v352 = m.ExcPending
									if v352 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(239569), int32(0))
										mBase = m.M
										v356 = m.ExcPending
										if v356 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(523892), int32(4008), int32(343390))
											mBase = m.M
											v361 = m.ExcPending
											if v361 != 0 {
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
								if v98 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v365 = m.ExcPending
									if v365 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(352583810))
										mBase = m.M
										v368 = m.ExcPending
										if v368 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(251854), int32(0))
											mBase = m.M
											v372 = m.ExcPending
											if v372 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(523892), int32(4013), int32(343390))
												mBase = m.M
												v377 = m.ExcPending
												if v377 != 0 {
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
									if v135 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v365 = m.ExcPending
										if v365 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(352583810))
											mBase = m.M
											v368 = m.ExcPending
											if v368 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(251854), int32(0))
												mBase = m.M
												v372 = m.ExcPending
												if v372 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(523892), int32(4013), int32(343390))
													mBase = m.M
													v377 = m.ExcPending
													if v377 != 0 {
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
										if v22 == int32(-12288) {
											if v54&int32(65535) == int32(53248) {
												v152 = F_make_result_opt_error(m, int32(1770476), int32(0))
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return int32(0)
												} else {
													v335 = v152
													m.G0 = v11 + int32(128)
													return v335
												}
											} else {
												v156 = F_make_result_opt_error(m, int32(1770548), int32(0))
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													v335 = v156
													m.G0 = v11 + int32(128)
													return v335
												}
											}
										} else {
											v160 = F_make_result_opt_error(m, int32(1770500), int32(0))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int32(0)
											} else {
												v335 = v160
												m.G0 = v11 + int32(128)
												return v335
											}
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(1770476), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v335 = v63
							m.G0 = v11 + int32(128)
							return v335
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v35 = base.B2i32(int32(0) <= v22)
					if int32(0) <= v22 {
						v36 = int32(-8)
					} else {
						v36 = int32(-6)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = int32(base.Ui32(int32(base.Ui32(v29)>>(uint(int32(2))%32))+v36) >> (uint(int32(1)) % 32))
					if int32(0) <= v22 {
						v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+6)))
						v163 = v162
					} else {
						v163 = v21<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v21&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v163
					v165 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v165
					v174 = base.B2i32(v22 < v165)
					if v22 < v165 {
						v175 = int32(base.Ui32(v21)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v175 = v21 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v175
					v182 = v21 & int32(49152)
					if v182 == int32(32768) {
						v185 = v21 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v185 = v182
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v185
					if v22 < v165 {
						v189 = int32(6)
					} else {
						v189 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v14 + v189
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v198 = base.B2i32(int32(0) <= v26)
					if int32(0) <= v26 {
						v199 = int32(-8)
					} else {
						v199 = int32(-6)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(base.Ui32(int32(base.Ui32(v192)>>(uint(int32(2))%32))+v199) >> (uint(int32(1)) % 32))
					if int32(0) <= v26 {
						v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+6)))
						v214 = v204
					} else {
						v214 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v215 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v215
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v215
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v214
					v220 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v220
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v215
					v227 = base.B2i32(v26 < v220)
					if v26 < v220 {
						v228 = int32(6)
					} else {
						v228 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v19 + v228
					v236 = v25 & int32(49152)
					if v236 == int32(32768) {
						v239 = v25 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v239 = v236
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v239
					if v26 < v220 {
						v247 = int32(base.Ui32(v25)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v247 = v25 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v247
					v249 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+120)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v249
					v256 = v11 + int32(96)
					*(*int64)(unsafe.Add(mBase, uint32(v256))) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v249
					v265 = F_estimate_ln_dweight(m, v11+int32(56))
					mBase = m.M
					v266 = m.ExcPending
					if v266 != 0 {
						return int32(0)
					} else {
						v275 = F_estimate_ln_dweight(m, v11+int32(32))
						mBase = m.M
						v276 = m.ExcPending
						if v276 != 0 {
							return int32(0)
						} else {
							v277 = v275 - v265
							v278 = int32(16) - v277
							if v175 < v278 {
								v280 = v278
							} else {
								v280 = v175
							}
							if base.Ui32(v247) < base.Ui32(v280) {
								v282 = v280
							} else {
								v282 = v247
							}
							if base.Ui32(int32(1000)) <= base.Ui32(v282) {
								v285 = int32(1000)
							} else {
								v285 = v282
							}
							v289 = v285 + v277 - v265 + int32(8)
							v290 = int32(0)
							if v290 < v289 {
								v293 = v289
							} else {
								v293 = v290
							}
							F_ln_var(m, v11+int32(56), v11+int32(104), v293)
							mBase = m.M
							v295 = m.ExcPending
							if v295 != 0 {
								return int32(0)
							} else {
								v302 = v285 - v265 + int32(8)
								v303 = int32(0)
								if v303 < v302 {
									v306 = v302
								} else {
									v306 = v303
								}
								F_ln_var(m, v11+int32(32), v11+int32(80), v306)
								mBase = m.M
								v308 = m.ExcPending
								if v308 != 0 {
									return int32(0)
								} else {
									F_div_var(m, v11+int32(80), v11+int32(104), v11+int32(8), v285, int32(1), int32(0))
									mBase = m.M
									v318 = m.ExcPending
									if v318 != 0 {
										return int32(0)
									} else {
										v319 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
										if v319 != 0 {
											F_pfree(m, v319)
											mBase = m.M
											v321 = m.ExcPending
											if v321 != 0 {
												return int32(0)
											} else {
												v322 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
												if v322 != 0 {
													F_pfree(m, v322)
													mBase = m.M
													v324 = m.ExcPending
													if v324 != 0 {
														return int32(0)
													} else {
														v328 = F_make_result_opt_error(m, v11+int32(8), int32(0))
														mBase = m.M
														v329 = m.ExcPending
														if v329 != 0 {
															return int32(0)
														} else {
															v330 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
															if v330 == int32(0) {
																v335 = v328
																m.G0 = v11 + int32(128)
																return v335
															} else {
																F_pfree(m, v330)
																mBase = m.M
																v334 = m.ExcPending
																if v334 != 0 {
																	return int32(0)
																} else {
																	v335 = v328
																	m.G0 = v11 + int32(128)
																	return v335
																}
															}
														}
													}
												} else {
													v328 = F_make_result_opt_error(m, v11+int32(8), int32(0))
													mBase = m.M
													v329 = m.ExcPending
													if v329 != 0 {
														return int32(0)
													} else {
														v330 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
														if v330 == int32(0) {
															v335 = v328
															m.G0 = v11 + int32(128)
															return v335
														} else {
															F_pfree(m, v330)
															mBase = m.M
															v334 = m.ExcPending
															if v334 != 0 {
																return int32(0)
															} else {
																v335 = v328
																m.G0 = v11 + int32(128)
																return v335
															}
														}
													}
												}
											}
										} else {
											v322 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
											if v322 != 0 {
												F_pfree(m, v322)
												mBase = m.M
												v324 = m.ExcPending
												if v324 != 0 {
													return int32(0)
												} else {
													v328 = F_make_result_opt_error(m, v11+int32(8), int32(0))
													mBase = m.M
													v329 = m.ExcPending
													if v329 != 0 {
														return int32(0)
													} else {
														v330 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
														if v330 == int32(0) {
															v335 = v328
															m.G0 = v11 + int32(128)
															return v335
														} else {
															F_pfree(m, v330)
															mBase = m.M
															v334 = m.ExcPending
															if v334 != 0 {
																return int32(0)
															} else {
																v335 = v328
																m.G0 = v11 + int32(128)
																return v335
															}
														}
													}
												}
											} else {
												v328 = F_make_result_opt_error(m, v11+int32(8), int32(0))
												mBase = m.M
												v329 = m.ExcPending
												if v329 != 0 {
													return int32(0)
												} else {
													v330 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
													if v330 == int32(0) {
														v335 = v328
														m.G0 = v11 + int32(128)
														return v335
													} else {
														F_pfree(m, v330)
														mBase = m.M
														v334 = m.ExcPending
														if v334 != 0 {
															return int32(0)
														} else {
															v335 = v328
															m.G0 = v11 + int32(128)
															return v335
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
			} else {
				if v22 == int32(-16384) {
					v63 = F_make_result_opt_error(m, int32(1770476), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v335 = v63
						m.G0 = v11 + int32(128)
						return v335
					}
				} else {
					v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
					v54 = v52
					if v54&int32(65535) != int32(49152) {
						if v22 == int32(-12288) {
							v69 = int32(1)
						} else {
							v69 = int32(-1)
						}
						v70 = int32(49152)
						v71 = v21 & v70
						if v71 == v70 {
							v98 = v69
						} else {
							v74 = int32(0)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							if v74 <= v22 {
								v82 = int32(-8)
							} else {
								v82 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v75)>>(uint(int32(2))%32))+v82) < base.Ui32(int32(2)) {
								v98 = v74
							} else {
								v87 = int32(1)
								if v71 == int32(32768) {
									v94 = v21 << (uint(v87) % 32) & int32(16384)
								} else {
									v94 = v71
								}
								if v94 == int32(16384) {
									v97 = int32(-1)
								} else {
									v97 = v87
								}
								v98 = v97
							}
						}
						v100 = v54 & int32(65535)
						if v100 == int32(53248) {
							v105 = int32(1)
						} else {
							v105 = int32(-1)
						}
						v106 = int32(49152)
						v107 = v54 & v106
						if v107 == v106 {
							v135 = v105
						} else {
							v110 = int32(0)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							if v110 <= base.I32_extend16_s(v54) {
								v119 = int32(-8)
							} else {
								v119 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v111)>>(uint(int32(2))%32))+v119) < base.Ui32(int32(2)) {
								v135 = v110
							} else {
								v124 = int32(1)
								if v107 == int32(32768) {
									v131 = v100 << (uint(v124) % 32) & int32(16384)
								} else {
									v131 = v107
								}
								if v131 == int32(16384) {
									v134 = int32(-1)
								} else {
									v134 = v124
								}
								v135 = v134
							}
						}
						if v98 < int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v349 = m.ExcPending
							if v349 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(352583810))
								mBase = m.M
								v352 = m.ExcPending
								if v352 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(239569), int32(0))
									mBase = m.M
									v356 = m.ExcPending
									if v356 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523892), int32(4008), int32(343390))
										mBase = m.M
										v361 = m.ExcPending
										if v361 != 0 {
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
							if v135 < int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v349 = m.ExcPending
								if v349 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(352583810))
									mBase = m.M
									v352 = m.ExcPending
									if v352 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(239569), int32(0))
										mBase = m.M
										v356 = m.ExcPending
										if v356 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(523892), int32(4008), int32(343390))
											mBase = m.M
											v361 = m.ExcPending
											if v361 != 0 {
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
								if v98 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v365 = m.ExcPending
									if v365 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(352583810))
										mBase = m.M
										v368 = m.ExcPending
										if v368 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(251854), int32(0))
											mBase = m.M
											v372 = m.ExcPending
											if v372 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(523892), int32(4013), int32(343390))
												mBase = m.M
												v377 = m.ExcPending
												if v377 != 0 {
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
									if v135 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v365 = m.ExcPending
										if v365 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(352583810))
											mBase = m.M
											v368 = m.ExcPending
											if v368 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(251854), int32(0))
												mBase = m.M
												v372 = m.ExcPending
												if v372 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(523892), int32(4013), int32(343390))
													mBase = m.M
													v377 = m.ExcPending
													if v377 != 0 {
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
										if v22 == int32(-12288) {
											if v54&int32(65535) == int32(53248) {
												v152 = F_make_result_opt_error(m, int32(1770476), int32(0))
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return int32(0)
												} else {
													v335 = v152
													m.G0 = v11 + int32(128)
													return v335
												}
											} else {
												v156 = F_make_result_opt_error(m, int32(1770548), int32(0))
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													v335 = v156
													m.G0 = v11 + int32(128)
													return v335
												}
											}
										} else {
											v160 = F_make_result_opt_error(m, int32(1770500), int32(0))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int32(0)
											} else {
												v335 = v160
												m.G0 = v11 + int32(128)
												return v335
											}
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(1770476), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v335 = v63
							m.G0 = v11 + int32(128)
							return v335
						}
					}
				}
			}
		}
	}
}
func F_numeric_poly_var_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v15 = F_numeric_poly_stddev_internal(m, v10, int32(1), int32(0), v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
func F_numeric_round(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
		if base.Ui32(int32(49152)) <= base.Ui32(v20) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v26 = F_palloc(m, int32(base.Ui32(v23)>>(uint(int32(2))%32)))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v30 = int32(base.Ui32(v28) >> (uint(int32(2)) % 32))
				if v30 != 0 {
					v31 = F__emscripten_memcpy_bulkmem(m, v26, v16, v30)
					mBase = m.M
				} else {
				}
				v185 = v26
				m.G0 = v13 + int32(32)
				return v185
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v34 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v34
			*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v34
			*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v34
			F_set_var_from_num(m, v16, v13+int32(8))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v45 = int32(-131073)
				if v33 <= v45 {
					v48 = v45
				} else {
					v48 = v33
				}
				if int32(16383) <= v48 {
					v51 = int32(16383)
				} else {
					v51 = v48
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v51
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				v56 = v51 + v53<<(uint(int32(2))%32)
				if v56+int32(4) < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
					v67 = v51 & int32(3)
					v71 = base.I32_div_s(v56+int32(7), int32(4))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					if v72 <= v71 {
						if v67 == int32(0) {
						} else {
							if v72 != v71 {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v71
								v85 = int32(1)
								v86 = v71 - v85
								v89 = v65 + v86<<(uint(v85)%32)
								v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89))))
								v91 = int32(2)
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(v91)%32))+uint32(_consts[1304])))
								v96 = base.I32_rem_s(v90, v95)
								v97 = v90 - v96
								*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v97)
								v100 = base.I32_div_s(v95, v91)
								if v96 < v100 {
									v142 = v86
								} else {
									v103 = v95 + base.I32_extend16_s(v97)
									if int32(9999) < v103 {
										v108 = v103 + int32(55536)
									} else {
										v108 = v103
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v108)
									if v103 < int32(10000) {
										v142 = v86
									} else {
										v113 = v86
										v119 = v113
										for {
											v127 = int32(1)
											v128 = v119 - v127
											v131 = v65 + v128<<(uint(v127)%32)
											v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131))))
											v136 = base.B2i32(int32(9998) < v134)
											if int32(9998) < v134 {
												v137 = int32(-9999)
											} else {
												v137 = v127
											}
											v138 = v137 + v134
											*(*uint16)(unsafe.Add(mBase, uint32(v131))) = uint16(v138)
											if int32(9998) < v134 {
												v119 = v128
												continue
											} else {
												break
											}
											break
										}
										v142 = v128
									}
								}
								if int32(0) <= v142 {
								} else {
									v152 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v53 + v152
									*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v71 + v152
									*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v65 - int32(2)
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v71
						if v67 != 0 {
							v85 = int32(1)
							v86 = v71 - v85
							v89 = v65 + v86<<(uint(v85)%32)
							v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89))))
							v91 = int32(2)
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(v91)%32))+uint32(_consts[1304])))
							v96 = base.I32_rem_s(v90, v95)
							v97 = v90 - v96
							*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v97)
							v100 = base.I32_div_s(v95, v91)
							if v96 < v100 {
								v142 = v86
							} else {
								v103 = v95 + base.I32_extend16_s(v97)
								if int32(9999) < v103 {
									v108 = v103 + int32(55536)
								} else {
									v108 = v103
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v108)
								if v103 < int32(10000) {
									v142 = v86
								} else {
									v113 = v86
									v119 = v113
									for {
										v127 = int32(1)
										v128 = v119 - v127
										v131 = v65 + v128<<(uint(v127)%32)
										v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131))))
										v136 = base.B2i32(int32(9998) < v134)
										if int32(9998) < v134 {
											v137 = int32(-9999)
										} else {
											v137 = v127
										}
										v138 = v137 + v134
										*(*uint16)(unsafe.Add(mBase, uint32(v131))) = uint16(v138)
										if int32(9998) < v134 {
											v119 = v128
											continue
										} else {
											break
										}
										break
									}
									v142 = v128
								}
							}
						} else {
							v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65+v71<<(uint(int32(1))%32)))))
							if v82 <= int32(4999) {
								v142 = v71
							} else {
								v113 = v71
								v119 = v113
								for {
									v127 = int32(1)
									v128 = v119 - v127
									v131 = v65 + v128<<(uint(v127)%32)
									v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131))))
									v136 = base.B2i32(int32(9998) < v134)
									if int32(9998) < v134 {
										v137 = int32(-9999)
									} else {
										v137 = v127
									}
									v138 = v137 + v134
									*(*uint16)(unsafe.Add(mBase, uint32(v131))) = uint16(v138)
									if int32(9998) < v134 {
										v119 = v128
										continue
									} else {
										break
									}
									break
								}
								v142 = v128
							}
						}
						if int32(0) <= v142 {
						} else {
							v152 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v53 + v152
							*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v71 + v152
							*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v65 - int32(2)
						}
					}
				}
				if v33 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(0)
				} else {
				}
				v178 = F_make_result_opt_error(m, v13+int32(8), int32(0))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int32(0)
				} else {
					v180 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					if v180 == int32(0) {
						v185 = v178
						m.G0 = v13 + int32(32)
						return v185
					} else {
						F_pfree(m, v180)
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
							return int32(0)
						} else {
							v185 = v178
							m.G0 = v13 + int32(32)
							return v185
						}
					}
				}
			}
		}
	}
}
func F_numeric_sqrt(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v17 = base.I32_extend16_s(v16)
		if base.Ui32(int32(49152)) <= base.Ui32(v16) {
			if v17 == int32(-4096) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(369361026))
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(239524), int32(0))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523892), int32(3813), int32(84217))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
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
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v25 = F_palloc(m, int32(base.Ui32(v22)>>(uint(int32(2))%32)))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v29 = int32(base.Ui32(v27) >> (uint(int32(2)) % 32))
					if v29 != 0 {
						v30 = F__emscripten_memcpy_bulkmem(m, v25, v12, v29)
						mBase = m.M
					} else {
					}
					v113 = v25
					m.G0 = v9 + int32(48)
					return v113
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v38 = base.B2i32(int32(0) <= v17)
			if int32(0) <= v17 {
				v39 = int32(-8)
			} else {
				v39 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(int32(base.Ui32(v32)>>(uint(int32(2))%32))+v39) >> (uint(int32(1)) % 32))
			if int32(0) <= v17 {
				v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v54 = v44
			} else {
				v54 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
			}
			v55 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v55
			v58 = v9 + int32(16)
			*(*int64)(unsafe.Add(mBase, uint32(v58))) = v55
			v61 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v61
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v55
			v68 = base.B2i32(v17 < v61)
			if v17 < v61 {
				v69 = int32(6)
			} else {
				v69 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v12 + v69
			v77 = v16 & int32(49152)
			if v77 == int32(32768) {
				v80 = v16 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v80 = v77
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v80
			if v17 < v61 {
				v88 = int32(base.Ui32(v16)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v88 = v16 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v88
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v54
			v97 = int32(15) - v54<<(uint(int32(1))%32)
			if v88 < v97 {
				v99 = v97
			} else {
				v99 = v88
			}
			if int32(1000) <= v99 {
				v102 = int32(1000)
			} else {
				v102 = v99
			}
			F_sqrt_var(m, v9+int32(24), v9, v102)
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				v106 = F_make_result_opt_error(m, v9, int32(0))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
					if v108 == int32(0) {
						v113 = v106
						m.G0 = v9 + int32(48)
						return v113
					} else {
						F_pfree(m, v108)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v113 = v106
							m.G0 = v9 + int32(48)
							return v113
						}
					}
				}
			}
		}
	}
}
func F_numeric_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v96 int64
	_ = v96
	var v100 int64
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(96)
	return v421
L2:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = v16 + (v17 + (v18 + v19))
	if v22 != int64(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v26)
	v421 = int32(0)
	goto L1
L6:
	;
	goto L5
L7:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v36)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) < v38 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if int64(1) < v22 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v33)
	v421 = int32(0)
	goto L1
L10:
	;
	v53 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v53
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v75 = v14 + int32(72)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v76 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v49 = F_make_result_opt_error(m, int32(1770476), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	if int64(0) < v41 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	if v44 <= int64(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	return int32(0)
L16:
	;
	v421 = v49
	goto L1
L17:
	;
	F_pfree(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v80 = F_palloc(m, int32(12))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v80
	v83 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v80))) = uint16(v83)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v86 + int32(2)
	if v73 < int64(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v130
	F_accum_sum_final(m, l0+int32(16), v14+int32(48))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L15
	} else {
		goto L31
	}
L23:
	;
	v106 = v83
	v109 = v86 + int32(12)
	v112 = v100
	goto L28
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = int64(16384)
	v100 = int64(0) - v73
	goto L23
L25:
	;
	goto L26
L26:
	;
	v96 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v96
	if v73 == v96 {
		v130 = v83
		v134 = int32(0)
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v100 = v73
	goto L23
L28:
	;
	v115 = v109 - int32(2)
	v117 = base.I64_div_u_s(v112, int64(10000))
	v120 = v117*int64(55536) + v112
	*(*uint16)(unsafe.Add(mBase, uint32(v115))) = uint16(v120)
	v123 = v106 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v112) {
		v106 = v123
		v109 = v115
		v112 = v117
		goto L28
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v115
	v130 = v123
	v134 = v106
	goto L22
L30:
	;
	goto L29
L31:
	;
	F_accum_sum_final(m, l0+int32(44), v14+int32(24))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v152 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v152
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v152
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v152
	F_sub_var(m, v14+int32(72), int32(1770600), v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	v164 = v14 + int32(48)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v171 = v169 << (uint(int32(1)) % 32)
	F_mul_var(m, v164, v164, v164, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	v177 = v14 + int32(24)
	F_mul_var(m, v14+int32(72), v177, v177, v171)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	v183 = v14 + int32(24)
	F_sub_var(m, v183, v14+int32(48), v183)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L15
	} else {
		goto L36
	}
L36:
	;
	v191 = v14 + int32(24)
	v192 = int32(1770548)
	v199 = *(*int32)(unsafe.Add(mBase, _consts[1306]))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[1307]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v201 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v410 != 0 {
		goto L109
	} else {
		goto L110
	}
L38:
	;
	if v237 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L39:
	;
	if v200 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v200 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v237 = int32(0)
	goto L38
L43:
	;
	goto L44
L44:
	;
	if v199 == int32(16384) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v211 = int32(1)
	goto L47
L46:
	;
	v211 = int32(-1)
	goto L47
L47:
	;
	v237 = v211
	goto L38
L48:
	;
	if v212 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[1308]))
	v219 = *(*int32)(unsafe.Add(mBase, _consts[1309]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	if v212 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v217 = int32(-1)
	goto L53
L52:
	;
	v217 = int32(1)
	goto L53
L53:
	;
	v237 = v217
	goto L38
L54:
	;
	if v199 == int32(16384) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if v199 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v237 = int32(1)
	goto L38
L58:
	;
	goto L59
L59:
	;
	v227 = F_cmp_abs_common(m, v221, v201, v220, v219, v200, v218)
	mBase = m.M
	v237 = v227
	goto L38
L60:
	;
	v237 = int32(-1)
	goto L38
L61:
	;
	goto L62
L62:
	;
	v231 = F_cmp_abs_common(m, v219, v200, v218, v221, v201, v220)
	mBase = m.M
	v237 = v231
	goto L38
L63:
	;
	v242 = F_make_result_opt_error(m, int32(1770548), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L15
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v245 = v14 + int32(72)
	if l2 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v409 = v242
	goto L37
L67:
	;
	v248 = v14
	goto L69
L68:
	;
	v248 = v245
	goto L69
L69:
	;
	F_mul_var(m, v245, v248, v14, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L15
	} else {
		goto L70
	}
L70:
	;
	v252 = int32(0)
	v255 = v14 + int32(24)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v252 < v256 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if int32(0) < v305 {
		goto L82
	} else {
		goto L83
	}
L72:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v260 = v252
	goto L75
L73:
	;
	goto L74
L74:
	;
	v291 = int32(0)
	v295 = v291
	v304 = v291
	goto L71
L75:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259+v260<<(uint(int32(1))%32)))))
	if v274 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L74
L77:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v295 = v274
	v304 = v260 - v275
	goto L71
L78:
	;
	goto L79
L79:
	;
	v278 = v260 + int32(1)
	if v278 != v256 {
		v260 = v278
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v367 = (v346+v304+base.B2i32(base.I32_extend16_s(v295) <= base.I32_extend16_s(v349)))<<(uint(int32(2))%32) + int32(16)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	if v368 < v367 {
		goto L91
	} else {
		goto L92
	}
L82:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v310 = int32(0)
	goto L85
L83:
	;
	v334 = v252
	goto L84
L84:
	;
	v346 = v334
	v349 = int32(0)
	goto L81
L85:
	;
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308+v310<<(uint(int32(1))%32)))))
	if v324 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v334 = int32(0)
	goto L84
L87:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v346 = v325 - v310
	v349 = v324
	goto L81
L88:
	;
	goto L89
L89:
	;
	v328 = v310 + int32(1)
	if v328 != v305 {
		v310 = v328
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v370 = v367
	goto L93
L92:
	;
	v370 = v368
	goto L93
L93:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v371 < v370 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v373 = v370
	goto L96
L95:
	;
	v373 = v371
	goto L96
L96:
	;
	v374 = int32(0)
	if v374 < v373 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v377 = v373
	goto L99
L98:
	;
	v377 = v374
	goto L99
L99:
	;
	if int32(1000) <= v377 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v380 = int32(1000)
	goto L102
L101:
	;
	v380 = v377
	goto L102
L102:
	;
	v381 = int32(1)
	F_div_var(m, v14+int32(24), v14, v14+int32(48), v380, v381, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L15
	} else {
		goto L103
	}
L103:
	;
	if l1 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v388 = v14 + int32(48)
	F_sqrt_var(m, v388, v388, v380)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L15
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v396 = F_make_result_opt_error(m, v14+int32(48), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L15
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v409 = v396
	goto L37
L109:
	;
	F_pfree(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L15
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	if v413 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L111
L113:
	;
	F_pfree(m, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L15
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v416 == int32(0) {
		v421 = v409
		goto L1
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	F_pfree(m, v416)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L15
	} else {
		goto L118
	}
L118:
	;
	v421 = v409
	goto L1
}
func F_numeric_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
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
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 + int32(48)
	return v387
L4:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v52-int32(1)) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v22 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = int32(4)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v27&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v40 = int32(1)
	if v22&v40 != 0 {
		v52 = int32(base.Ui32(v22)>>(uint(v40)%32)) - v40
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v36 = v25
	goto L11
L10:
	;
	v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
	goto L11
L11:
	;
	if v27 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = v25
	goto L14
L13:
	;
	v39 = v36
	goto L14
L14:
	;
	v52 = v39
	goto L4
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L16:
	;
	v58 = F_cstring_to_text(m, int32(790230))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v64 = F_palloc0(m, v52<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v387 = v58
	goto L3
L20:
	;
	v70 = F_NUM_cache(m, v52, v12+int32(12), v20, v12+int32(11))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v72&int32(1024) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v365 = v64 + int32(4)
	F_NUM_processor(m, v70, v12+int32(12), v365, v357, int32(0), v356, v359, int32(1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L120
	}
L23:
	;
	v78 = F_numeric_int4_opt_error(m, v15, v12+int32(10))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v72&int32(16384) != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)))
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = int32(2147483647)
	goto L29
L28:
	;
	v81 = v78
	goto L29
L29:
	;
	v82 = F_int_to_roman(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v356 = int32(0)
	v357 = v82
	v359 = v2
	goto L22
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v88 = F_numeric_out_sci(m, v15, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	if v72&int32(2048) != 0 {
		goto L91
	} else {
		goto L92
	}
L34:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v192 == int32(45) {
		goto L66
	} else {
		goto L67
	}
L35:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v173 = v172 + v87
	v176 = F_palloc(m, v173+int32(7))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L64
	}
L36:
	;
	v90 = int32(552455)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1293])))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v94 == int32(0) {
		v113 = v93
		v114 = v94
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v114-v113 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	if v93 != v94 {
		v113 = v93
		v114 = v94
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v98 = v88
	v99 = v90
	goto L41
L41:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v103 == int32(0) {
		v113 = v102
		v114 = v103
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v113 = v102
	v114 = v103
	goto L38
L43:
	;
	v106 = int32(1)
	if v102 == v103 {
		v98 = v98 + v106
		v99 = v99 + v106
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v118 = int32(11507)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1294])))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v122 == int32(0) {
		v141 = v121
		v142 = v122
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v142-v141 == int32(0) {
		goto L35
	} else {
		goto L54
	}
L47:
	;
	goto L46
L48:
	;
	if v121 != v122 {
		v141 = v121
		v142 = v122
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v126 = v88
	v127 = v118
	goto L50
L50:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v131 == int32(0) {
		v141 = v130
		v142 = v131
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v141 = v130
	v142 = v131
	goto L47
L52:
	;
	v134 = int32(1)
	if v130 == v131 {
		v126 = v126 + v134
		v127 = v127 + v134
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v146 = int32(11496)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1295])))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v150 == int32(0) {
		v169 = v149
		v170 = v150
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v170-v169 != 0 {
		goto L34
	} else {
		goto L63
	}
L56:
	;
	goto L55
L57:
	;
	if v149 != v150 {
		v169 = v149
		v170 = v150
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v154 = v88
	v155 = v146
	goto L59
L59:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v159 == int32(0) {
		v169 = v158
		v170 = v159
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v169 = v158
	v170 = v159
	goto L56
L61:
	;
	v162 = int32(1)
	if v158 == v159 {
		v154 = v154 + v162
		v155 = v155 + v162
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L35
L64:
	;
	v180 = v173 + int32(6)
	v182 = F__emscripten_memset_bulkmem(m, v176, base.I32_extend8_s(int32(35)), v180)
	mBase = m.M
	goto L65
L65:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v180))) = uint8(v184)
	v186 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v186)
	v189 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v172)+1)) = uint8(v189)
	v356 = v184
	v357 = v176
	v359 = v2
	goto L22
L66:
	;
	v356 = int32(0)
	v357 = v88
	v359 = v2
	goto L22
L67:
	;
	goto L68
L68:
	;
	v196 = F_strlen(m, v88)
	mBase = m.M
	v199 = F_palloc(m, v196+int32(2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v201 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v201)
	v204 = v199 + int32(1)
	if (v88^v204)&int32(3) != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v356 = int32(0)
	v357 = v199
	v359 = v2
	goto L22
L71:
	;
	goto L70
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v258)
	if v258&int32(255) == int32(0) {
		goto L71
	} else {
		goto L87
	}
L73:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v257 = v88
	v258 = v210
	v259 = v204
	goto L72
L74:
	;
	goto L75
L75:
	;
	if v88&int32(3) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v214 = v88
	v216 = v204
	goto L79
L77:
	;
	v228 = v88
	v230 = v204
	goto L78
L78:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v235 = int32(-2139062144)
	if (int32(16843008)-v232|v232)&v235 != v235 {
		v257 = v228
		v258 = v232
		v259 = v230
		goto L72
	} else {
		goto L83
	}
L79:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v217)
	if v217 == int32(0) {
		goto L71
	} else {
		goto L81
	}
L80:
	;
	v228 = v224
	v230 = v222
	goto L78
L81:
	;
	v221 = int32(1)
	v222 = v216 + v221
	v224 = v214 + v221
	if v224&int32(3) != 0 {
		v214 = v224
		v216 = v222
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v240 = v228
	v241 = v232
	v242 = v230
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v241
	v244 = int32(4)
	v245 = v242 + v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v248 = v240 + v244
	v252 = int32(-2139062144)
	if (v246|(int32(16843008)-v246))&v252 == v252 {
		v240 = v248
		v241 = v246
		v242 = v245
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v257 = v248
	v258 = v246
	v259 = v245
	goto L72
L86:
	;
	goto L85
L87:
	;
	v266 = v257
	v268 = v259
	goto L88
L88:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)) = uint8(v269)
	v271 = int32(1)
	if v269 != 0 {
		v266 = v266 + v271
		v268 = v268 + v271
		goto L88
	} else {
		goto L90
	}
L89:
	;
	goto L71
L90:
	;
	goto L89
L91:
	;
	v283 = int32(0)
	v287 = F_int64_to_numeric(m, int64(10))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	v304 = v15
	goto L93
L93:
	;
	v307 = int32(0)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v311 = F_DirectFunctionCall2Coll(m, int32(1275), v307, v304, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L101
	}
L94:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v291 = F_int64_to_numeric(m, base.I64_extend_i32_s(v289))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v293 = F_DirectFunctionCall2Coll(m, int32(1313), v283, v287, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v295 = F_pg_detoast_datum(m, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v297 = F_DirectFunctionCall2Coll(m, int32(1278), v283, v15, v295)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v299 = F_pg_detoast_datum(m, v297)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v289 + v301
	v304 = v299
	goto L93
L100:
	;
	if v317 == int32(45) {
		goto L111
	} else {
		goto L112
	}
L101:
	;
	v313 = F_pg_detoast_datum(m, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v315 = F_DirectFunctionCall1Coll(m, int32(618), v307, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	v319 = base.B2i32(v317 == int32(45))
	v320 = v315 + v319
	v321 = int32(46)
	v322 = F___strchrnul(m, v320, v321)
	mBase = m.M
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v324 == v321 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v328 != 0 {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v328 = v322
	goto L107
L106:
	;
	v328 = int32(0)
	goto L107
L107:
	;
	goto L104
L108:
	;
	v331 = v328 - v320
	goto L100
L109:
	;
	goto L110
L110:
	;
	v330 = F_strlen(m, v320)
	mBase = m.M
	v331 = v330
	goto L100
L111:
	;
	v334 = int32(45)
	goto L113
L112:
	;
	v334 = int32(43)
	goto L113
L113:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v331 < v335 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v356 = v335 - v331
	v357 = v320
	v359 = v334
	goto L22
L115:
	;
	goto L116
L116:
	;
	v338 = int32(0)
	if v331 <= v335 {
		v356 = v338
		v357 = v320
		v359 = v334
		goto L22
	} else {
		goto L117
	}
L117:
	;
	v340 = v335 + v310
	v343 = F_palloc(m, v340+int32(2))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v347 = v340 + int32(1)
	v349 = F__emscripten_memset_bulkmem(m, v343, base.I32_extend8_s(int32(35)), v347)
	mBase = m.M
	goto L119
L119:
	;
	v351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v349+v347))) = uint8(v351)
	v354 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v335+v349))) = uint8(v354)
	v356 = v338
	v357 = v343
	v359 = v334
	goto L22
L120:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if v370 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_pfree(m, v70)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v375 = F_strlen(m, v365)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v375<<(uint(int32(2))%32) + int32(16)
	v387 = v64
	goto L3
L124:
	;
	goto L123
}
func F_numeric_trunc(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
		if base.Ui32(int32(49152)) <= base.Ui32(v14) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v20 = F_palloc(m, int32(base.Ui32(v17)>>(uint(int32(2))%32)))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v24 = int32(base.Ui32(v22) >> (uint(int32(2)) % 32))
				if v24 != 0 {
					v25 = F__emscripten_memcpy_bulkmem(m, v20, v10, v24)
					mBase = m.M
				} else {
				}
				v101 = v20
				m.G0 = v7 + int32(32)
				return v101
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v28
			F_set_var_from_num(m, v10, v7+int32(8))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v39 = int32(-131072)
				if v27 <= v39 {
					v42 = v39
				} else {
					v42 = v27
				}
				if int32(16383) <= v42 {
					v45 = int32(16383)
				} else {
					v45 = v42
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v45
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v50 = v45 + v47<<(uint(int32(2))%32)
				if v50+int32(4) <= int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
				} else {
					v62 = base.I32_div_s(v50+int32(7), int32(4))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					if v63 < v62 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v62
						v67 = v45 & int32(3)
						if v67 == int32(0) {
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
							v74 = int32(2)
							v75 = v70 + v62<<(uint(int32(1))%32) - v74
							v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75))))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(v74)%32))+uint32(_consts[1304])))
							v82 = base.I32_rem_s(v76, v81)
							v83 = v76 - v82
							*(*uint16)(unsafe.Add(mBase, uint32(v75))) = uint16(v83)
						}
					}
				}
				if v27 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(0)
				} else {
				}
				v94 = F_make_result_opt_error(m, v7+int32(8), int32(0))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
					if v96 == int32(0) {
						v101 = v94
						m.G0 = v7 + int32(32)
						return v101
					} else {
						F_pfree(m, v96)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v101 = v94
							m.G0 = v7 + int32(32)
							return v101
						}
					}
				}
			}
		}
	}
}
