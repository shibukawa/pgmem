package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_numeric_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
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
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(_a_F_numeric_cmp_0)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(_a_F_numeric_cmp_1) {
					if v25 != int32(_a_F_numeric_cmp_0) {
						if v24 != int32(_a_F_numeric_cmp_2) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v163 = v46
					} else {
						v163 = base.B2i32(v24 != int32(_a_F_numeric_cmp_0))
					}
				} else {
					if v24 == int32(_a_F_numeric_cmp_0) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(_a_F_numeric_cmp_1))
					}
					v163 = v41
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_cmp_0)) <= base.Ui32(v24) {
					if v24 == int32(_a_F_numeric_cmp_2) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v163 = v53
				} else {
					v55 = v6 + int32(6)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v63 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v64 = int32(-8)
					} else {
						v64 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v25) {
						v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v76 = v66
					} else {
						v76 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v78 = int32(base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
					v80 = v11 + int32(6)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v88 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v89 = int32(-8)
					} else {
						v89 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v101 = v91
					} else {
						v101 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v102 = int32(1)
					v103 = int32(base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) >> (uint(v102) % 32))
					v109 = v24 & int32(_a_F_numeric_cmp_0)
					if v109 == int32(_a_F_numeric_cmp_3) {
						v112 = v24 << (uint(v102) % 32) & int32(_a_F_numeric_cmp_4)
					} else {
						v112 = v109
					}
					if v78 == int32(0) {
						if v103 == int32(0) {
							v163 = int32(0)
						} else {
							if v112 == int32(_a_F_numeric_cmp_4) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v163 = v122
						}
					} else {
						if v27 == int32(_a_F_numeric_cmp_3) {
							v129 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_cmp_4)
						} else {
							v129 = v27
						}
						if v103 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v163 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v137 = v6 + int32(8)
							} else {
								v137 = v55
							}
							if int32(0) <= base.I32_extend16_s(v24) {
								v140 = v11 + int32(8)
							} else {
								v140 = v80
							}
							if v129 == int32(0) {
								if v112 == int32(_a_F_numeric_cmp_4) {
									v163 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v137, v78, v76, v140, v103, v101)
									mBase = m.M
									v163 = v146
								}
							} else {
								if v112 == int32(0) {
									v163 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v140, v103, v101, v137, v78, v76)
									mBase = m.M
									v163 = v150
								}
							}
						}
					}
				}
			}
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v164 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v168 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return v163
						}
					} else {
						return v163
					}
				}
			} else {
				v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v168 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return v163
					}
				} else {
					return v163
				}
			}
		}
	}
}
func F_numeric_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
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
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(_a_F_numeric_eq_0)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(_a_F_numeric_eq_1) {
					if v25 != int32(_a_F_numeric_eq_0) {
						if v24 != int32(_a_F_numeric_eq_2) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v163 = v46
					} else {
						v163 = base.B2i32(v24 != int32(_a_F_numeric_eq_0))
					}
				} else {
					if v24 == int32(_a_F_numeric_eq_0) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(_a_F_numeric_eq_1))
					}
					v163 = v41
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_eq_0)) <= base.Ui32(v24) {
					if v24 == int32(_a_F_numeric_eq_2) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v163 = v53
				} else {
					v55 = v6 + int32(6)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v63 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v64 = int32(-8)
					} else {
						v64 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v25) {
						v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v76 = v66
					} else {
						v76 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v78 = int32(base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
					v80 = v11 + int32(6)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v88 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v89 = int32(-8)
					} else {
						v89 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v101 = v91
					} else {
						v101 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v102 = int32(1)
					v103 = int32(base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) >> (uint(v102) % 32))
					v109 = v24 & int32(_a_F_numeric_eq_0)
					if v109 == int32(_a_F_numeric_eq_3) {
						v112 = v24 << (uint(v102) % 32) & int32(_a_F_numeric_eq_4)
					} else {
						v112 = v109
					}
					if v78 == int32(0) {
						if v103 == int32(0) {
							v163 = int32(0)
						} else {
							if v112 == int32(_a_F_numeric_eq_4) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v163 = v122
						}
					} else {
						if v27 == int32(_a_F_numeric_eq_3) {
							v129 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_eq_4)
						} else {
							v129 = v27
						}
						if v103 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v163 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v137 = v6 + int32(8)
							} else {
								v137 = v55
							}
							if int32(0) <= base.I32_extend16_s(v24) {
								v140 = v11 + int32(8)
							} else {
								v140 = v80
							}
							if v129 == int32(0) {
								if v112 == int32(_a_F_numeric_eq_4) {
									v163 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v137, v78, v76, v140, v103, v101)
									mBase = m.M
									v163 = v146
								}
							} else {
								if v112 == int32(0) {
									v163 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v140, v103, v101, v137, v78, v76)
									mBase = m.M
									v163 = v150
								}
							}
						}
					}
				}
			}
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v164 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v168 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v163 == int32(0))
						}
					} else {
						return base.B2i32(v163 == int32(0))
					}
				}
			} else {
				v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v168 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v163 == int32(0))
					}
				} else {
					return base.B2i32(v163 == int32(0))
				}
			}
		}
	}
}
func F_numeric_floor(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
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
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v21 = int32(base.Ui32(v19) >> (uint(int32(2)) % 32))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	if base.Ui32(int32(_a_F_numeric_floor_0)) <= base.Ui32(v22) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 + int32(48)
	return v359
L4:
	;
	v25 = F_palloc(m, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v35 = base.I32_extend16_s(v22)
	v37 = base.B2i32(int32(0) <= v35)
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v29 = int32(base.Ui32(v27) >> (uint(int32(2)) % 32))
	if v29 == int32(0) {
		v359 = v25
		goto L3
	} else {
		goto L8
	}
L8:
	;
	base.MemoryCopy(m, v25, v15, v29)
	v359 = v25
	goto L3
L9:
	;
	v38 = int32(-8)
	goto L11
L10:
	;
	v38 = int32(-6)
	goto L11
L11:
	;
	v39 = v21 + v38
	v41 = int32(base.Ui32(v39) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v41
	if int32(0) <= v35 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+6)))
	v53 = v43
	goto L14
L13:
	;
	v53 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v53
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	v60 = base.B2i32(v35 < v55)
	if v35 < v55 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = int32(6)
	goto L17
L16:
	;
	v61 = int32(8)
	goto L17
L17:
	;
	v62 = v15 + v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v62
	if v35 < v55 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = int32(base.Ui32(v22)>>(uint(int32(7))%32)) & int32(63)
	goto L20
L19:
	;
	v70 = v22 & int32(_a_F_numeric_floor_1)
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v70
	v77 = v22 & int32(_a_F_numeric_floor_0)
	if v77 == int32(_a_F_numeric_floor_2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = v22 << (uint(int32(1)) % 32) & int32(_a_F_numeric_floor_3)
	goto L23
L22:
	;
	v80 = v77
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v80
	v83 = v39 & int32(-2)
	v86 = F_palloc(m, v83+int32(2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v88)
	if base.B2i32(v41 == v88)|base.B2i32(v83 == v88) == v88 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	base.MemoryCopy(m, v86+int32(2), v62, v83)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v86
	v107 = int32(2)
	v108 = v86 + v107
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v112 = v110 << (uint(v107) % 32)
	if base.Ui32(int32(2147483644)) <= base.Ui32(v112) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v129
	if v80 != int32(_a_F_numeric_floor_3) {
		v323 = v129
		goto L35
	} else {
		goto L36
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+28)) = int64(0)
	v117 = int32(0)
	v127 = v117
	v129 = v117
	goto L28
L30:
	;
	goto L31
L31:
	;
	v122 = base.I32_div_s(v112+int32(7), int32(4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v122 < v123 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v125 = v122
	goto L34
L33:
	;
	v125 = v123
	goto L34
L34:
	;
	v127 = v110
	v129 = v125
	goto L28
L35:
	;
	v325 = v323 << (uint(int32(1)) % 32)
	v328 = F_palloc(m, v325+int32(2))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L89
	}
L36:
	;
	if v41 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v318 = v12 + int32(24)
	F_sub_var(m, v318, int32(_a_F_numeric_floor_4), v318)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L88
	}
L38:
	;
	if v129 != 0 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v129 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v323 = int32(0)
	goto L35
L42:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v138 == int32(0) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v141 = int32(0)
	if base.B2i32(v53 < v127)&base.B2i32(v141 < v129) == v141 {
		v172 = v127
		v176 = v141
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v314 == int32(0) {
		v323 = v129
		goto L35
	} else {
		goto L87
	}
L45:
	;
	v314 = v304
	goto L44
L46:
	;
	if base.B2i32(v41 <= int32(0))|base.B2i32(v53 <= v172) != 0 {
		v208 = v53
		v210 = v141
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v153 = v127
	v157 = v141
	goto L48
L48:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+v157<<(uint(int32(1))%32)))))
	if v163 != 0 {
		v304 = int32(1)
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v172 = v167
	v176 = v165
	goto L46
L50:
	;
	v164 = int32(1)
	v165 = v157 + v164
	v167 = v153 - v164
	if v167 <= v53 {
		v172 = v167
		v176 = v165
		goto L46
	} else {
		goto L51
	}
L51:
	;
	if v165 < v129 {
		v153 = v167
		v157 = v165
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	if v172 != v208 {
		v250 = v176
		v251 = v210
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v189 = v53
	v191 = v141
	goto L55
L55:
	;
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v191<<(uint(int32(1))%32)))))
	if v196 != 0 {
		v304 = int32(-1)
		goto L45
	} else {
		goto L57
	}
L56:
	;
	v208 = v200
	v210 = v198
	goto L53
L57:
	;
	v197 = int32(1)
	v198 = v191 + v197
	v200 = v189 - v197
	if v200 <= v172 {
		v208 = v200
		v210 = v198
		goto L53
	} else {
		goto L58
	}
L58:
	;
	if v198 < v41 {
		v189 = v200
		v191 = v198
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	if v129 < v250 {
		goto L69
	} else {
		goto L70
	}
L61:
	;
	v219 = v176
	v220 = v210
	goto L62
L62:
	;
	if base.B2i32(v129 <= v219)|base.B2i32(v41 <= v220) != 0 {
		v250 = v219
		v251 = v220
		goto L60
	} else {
		goto L64
	}
L63:
	;
	if base.I32_extend16_s(v236) < base.I32_extend16_s(v234) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v225 = int32(1)
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+v219<<(uint(v225)%32)))))
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220<<(uint(v225)%32)+v62))))
	if v234 == v236 {
		v219 = v219 + v225
		v220 = v220 + v225
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v243 = int32(1)
	goto L68
L67:
	;
	v243 = int32(-1)
	goto L68
L68:
	;
	v314 = v243
	goto L44
L69:
	;
	v254 = v250
	goto L71
L70:
	;
	v254 = v129
	goto L71
L71:
	;
	v261 = v250
	goto L72
L72:
	;
	if v254 == v261 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v304 = v287
	goto L45
L74:
	;
	if v41 < v251 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v287 = int32(1)
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+v261<<(uint(v287)%32)))))
	if v293 == int32(0) {
		v261 = v261 + v287
		goto L72
	} else {
		goto L86
	}
L77:
	;
	v266 = v251
	goto L79
L78:
	;
	v266 = v41
	goto L79
L79:
	;
	v274 = v251
	goto L80
L80:
	;
	if v266 == v274 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v304 = int32(-1)
	goto L45
L82:
	;
	v314 = int32(0)
	goto L44
L83:
	;
	goto L84
L84:
	;
	v278 = int32(1)
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274<<(uint(v278)%32)+v62))))
	if v283 == int32(0) {
		v274 = v274 + v278
		goto L80
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	goto L73
L87:
	;
	goto L37
L88:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v323 = v322
	goto L35
L89:
	;
	v330 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v328))) = uint16(v330)
	if base.B2i32(v325 == v330)|base.B2i32(v323 <= v330) == v330 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	base.MemoryCopy(m, v328+int32(2), v341, v325)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v343
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v328 + int32(2)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v351 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_pfree(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v355 = F_make_result_opt_error(m, v12, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	F_pfree(m, v328)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v359 = v355
	goto L3
}
func F_numeric_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
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
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(_a_F_numeric_ge_0)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(_a_F_numeric_ge_1) {
					if v25 != int32(_a_F_numeric_ge_0) {
						if v24 != int32(_a_F_numeric_ge_2) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v163 = v46
					} else {
						v163 = base.B2i32(v24 != int32(_a_F_numeric_ge_0))
					}
				} else {
					if v24 == int32(_a_F_numeric_ge_0) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(_a_F_numeric_ge_1))
					}
					v163 = v41
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_ge_0)) <= base.Ui32(v24) {
					if v24 == int32(_a_F_numeric_ge_2) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v163 = v53
				} else {
					v55 = v6 + int32(6)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v63 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v64 = int32(-8)
					} else {
						v64 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v25) {
						v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v76 = v66
					} else {
						v76 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v78 = int32(base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
					v80 = v11 + int32(6)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v88 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v89 = int32(-8)
					} else {
						v89 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v101 = v91
					} else {
						v101 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v102 = int32(1)
					v103 = int32(base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) >> (uint(v102) % 32))
					v109 = v24 & int32(_a_F_numeric_ge_0)
					if v109 == int32(_a_F_numeric_ge_3) {
						v112 = v24 << (uint(v102) % 32) & int32(_a_F_numeric_ge_4)
					} else {
						v112 = v109
					}
					if v78 == int32(0) {
						if v103 == int32(0) {
							v163 = int32(0)
						} else {
							if v112 == int32(_a_F_numeric_ge_4) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v163 = v122
						}
					} else {
						if v27 == int32(_a_F_numeric_ge_3) {
							v129 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_ge_4)
						} else {
							v129 = v27
						}
						if v103 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v163 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v137 = v6 + int32(8)
							} else {
								v137 = v55
							}
							if int32(0) <= base.I32_extend16_s(v24) {
								v140 = v11 + int32(8)
							} else {
								v140 = v80
							}
							if v129 == int32(0) {
								if v112 == int32(_a_F_numeric_ge_4) {
									v163 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v137, v78, v76, v140, v103, v101)
									mBase = m.M
									v163 = v146
								}
							} else {
								if v112 == int32(0) {
									v163 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v140, v103, v101, v137, v78, v76)
									mBase = m.M
									v163 = v150
								}
							}
						}
					}
				}
			}
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v164 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v168 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v163^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v163^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			} else {
				v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v168 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return int32(base.Ui32(v163^int32(-1)) >> (uint(int32(31)) % 32))
					}
				} else {
					return int32(base.Ui32(v163^int32(-1)) >> (uint(int32(31)) % 32))
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
		if base.Ui32(v14) <= base.Ui32(int32(_a_F_numeric_int8_opt_error_0)) {
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
				v84 = v48 & int32(_a_F_numeric_int8_opt_error_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v84
			v91 = v48 & int32(_a_F_numeric_int8_opt_error_2)
			if v91 == int32(_a_F_numeric_int8_opt_error_3) {
				v94 = v48 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int8_opt_error_4)
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
								F_errmsg(m, int32(_a_F_numeric_int8_opt_error_5), int32(0))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_numeric_int8_opt_error_6), int32(_a_F_numeric_int8_opt_error_7), int32(_a_F_numeric_int8_opt_error_8))
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
					if v14 == int32(_a_F_numeric_int8_opt_error_2) {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_numeric_int8_opt_error_9)
						F_errmsg(m, int32(_a_F_numeric_int8_opt_error_10), v10)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_int8_opt_error_6), int32(_a_F_numeric_int8_opt_error_11), int32(_a_F_numeric_int8_opt_error_8))
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
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_numeric_int8_opt_error_9)
						F_errmsg(m, int32(_a_F_numeric_int8_opt_error_12), v8+int32(-48))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_int8_opt_error_6), int32(_a_F_numeric_int8_opt_error_13), int32(_a_F_numeric_int8_opt_error_8))
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
		if base.Ui32(v42) < base.Ui32(int32(_a_F_numeric_int8_opt_error_2)) {
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
				v84 = v48 & int32(_a_F_numeric_int8_opt_error_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v84
			v91 = v48 & int32(_a_F_numeric_int8_opt_error_2)
			if v91 == int32(_a_F_numeric_int8_opt_error_3) {
				v94 = v48 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int8_opt_error_4)
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
								F_errmsg(m, int32(_a_F_numeric_int8_opt_error_5), int32(0))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_numeric_int8_opt_error_6), int32(_a_F_numeric_int8_opt_error_7), int32(_a_F_numeric_int8_opt_error_8))
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
	return base.B2i32(v2&int32(_a_F_numeric_is_inf_0) == int32(_a_F_numeric_is_inf_1))
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
	var v102 int32
	_ = v102
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
	var v139 int32
	_ = v139
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
	var v216 int32
	_ = v216
	var v218 int64
	_ = v218
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
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
			if base.Ui32(v21) <= base.Ui32(int32(_a_F_numeric_log_0)) {
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
				v26 = base.I32_extend16_s(v25)
				if base.Ui32(int32(_a_F_numeric_log_0)) < base.Ui32(v25) {
					v54 = v26
					if v54&int32(_a_F_numeric_log_1) != int32(_a_F_numeric_log_2) {
						if v22 == int32(-12288) {
							v69 = int32(1)
						} else {
							v69 = int32(-1)
						}
						v70 = int32(_a_F_numeric_log_2)
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
								if v71 == int32(_a_F_numeric_log_3) {
									v94 = v21 << (uint(v87) % 32) & int32(_a_F_numeric_log_4)
								} else {
									v94 = v71
								}
								if v94 == int32(_a_F_numeric_log_4) {
									v97 = int32(-1)
								} else {
									v97 = v87
								}
								v98 = v97
							}
						}
						v102 = v54 & int32(_a_F_numeric_log_1)
						if v102 == int32(_a_F_numeric_log_5) {
							v105 = int32(1)
						} else {
							v105 = int32(-1)
						}
						v106 = int32(_a_F_numeric_log_2)
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
								if v107 == int32(_a_F_numeric_log_3) {
									v131 = v102 << (uint(v124) % 32) & int32(_a_F_numeric_log_4)
								} else {
									v131 = v107
								}
								if v131 == int32(_a_F_numeric_log_4) {
									v134 = int32(-1)
								} else {
									v134 = v124
								}
								v135 = v134
							}
						}
						if v98|v135 < int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v339 = m.ExcPending
							if v339 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(352583810))
								mBase = m.M
								v342 = m.ExcPending
								if v342 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_numeric_log_6), int32(0))
									mBase = m.M
									v346 = m.ExcPending
									if v346 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_numeric_log_7), int32(4008), int32(_a_F_numeric_log_8))
										mBase = m.M
										v351 = m.ExcPending
										if v351 != 0 {
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
							v139 = int32(0)
							if base.B2i32(v98 == v139)|base.B2i32(v135 == v139) != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v355 = m.ExcPending
								if v355 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(352583810))
									mBase = m.M
									v358 = m.ExcPending
									if v358 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_numeric_log_9), int32(0))
										mBase = m.M
										v362 = m.ExcPending
										if v362 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_numeric_log_7), int32(4013), int32(_a_F_numeric_log_8))
											mBase = m.M
											v367 = m.ExcPending
											if v367 != 0 {
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
									if v54&int32(_a_F_numeric_log_1) == int32(_a_F_numeric_log_5) {
										v152 = F_make_result_opt_error(m, int32(_a_F_numeric_log_10), int32(0))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											v325 = v152
											m.G0 = v11 + int32(128)
											return v325
										}
									} else {
										v156 = F_make_result_opt_error(m, int32(_a_F_numeric_log_11), int32(0))
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											v325 = v156
											m.G0 = v11 + int32(128)
											return v325
										}
									}
								} else {
									v160 = F_make_result_opt_error(m, int32(_a_F_numeric_log_12), int32(0))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										v325 = v160
										m.G0 = v11 + int32(128)
										return v325
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(_a_F_numeric_log_10), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v325 = v63
							m.G0 = v11 + int32(128)
							return v325
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
						v175 = v21 & int32(_a_F_numeric_log_13)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v175
					v182 = v21 & int32(_a_F_numeric_log_2)
					if v182 == int32(_a_F_numeric_log_3) {
						v185 = v21 << (uint(int32(1)) % 32) & int32(_a_F_numeric_log_4)
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
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v214
					v216 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v216
					v218 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v218
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v218
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v218
					v227 = base.B2i32(v26 < v216)
					if v26 < v216 {
						v228 = int32(6)
					} else {
						v228 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v19 + v228
					v236 = v25 & int32(_a_F_numeric_log_2)
					if v236 == int32(_a_F_numeric_log_3) {
						v239 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_log_4)
					} else {
						v239 = v236
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v239
					if v26 < v216 {
						v247 = int32(base.Ui32(v25)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v247 = v25 & int32(_a_F_numeric_log_13)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v247
					v249 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+120)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v249
					v262 = v11 + int32(56)
					v263 = F_estimate_ln_dweight(m, v262)
					mBase = m.M
					v264 = m.ExcPending
					if v264 != 0 {
						return int32(0)
					} else {
						v266 = v11 + int32(104)
						v270 = v11 + int32(32)
						v271 = F_estimate_ln_dweight(m, v270)
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int32(0)
						} else {
							v273 = v271 - v263
							v274 = int32(16) - v273
							if v175 < v274 {
								v276 = v274
							} else {
								v276 = v175
							}
							if base.Ui32(v247) < base.Ui32(v276) {
								v278 = v276
							} else {
								v278 = v247
							}
							if base.Ui32(int32(1000)) <= base.Ui32(v278) {
								v281 = int32(1000)
							} else {
								v281 = v278
							}
							v285 = v281 + v273 - v263 + int32(8)
							v286 = int32(0)
							if v286 < v285 {
								v289 = v285
							} else {
								v289 = v286
							}
							F_ln_var(m, v262, v266, v289)
							mBase = m.M
							v291 = m.ExcPending
							if v291 != 0 {
								return int32(0)
							} else {
								v293 = v11 + int32(80)
								v296 = v281 - v263 + int32(8)
								v297 = int32(0)
								if v297 < v296 {
									v300 = v296
								} else {
									v300 = v297
								}
								F_ln_var(m, v270, v293, v300)
								mBase = m.M
								v302 = m.ExcPending
								if v302 != 0 {
									return int32(0)
								} else {
									F_div_var(m, v293, v266, v11+int32(8), v281, int32(1), int32(0))
									mBase = m.M
									v308 = m.ExcPending
									if v308 != 0 {
										return int32(0)
									} else {
										v309 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
										if v309 != 0 {
											F_pfree(m, v309)
											mBase = m.M
											v311 = m.ExcPending
											if v311 != 0 {
												return int32(0)
											} else {
												v312 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
												if v312 != 0 {
													F_pfree(m, v312)
													mBase = m.M
													v314 = m.ExcPending
													if v314 != 0 {
														return int32(0)
													} else {
														v318 = F_make_result_opt_error(m, v11+int32(8), int32(0))
														mBase = m.M
														v319 = m.ExcPending
														if v319 != 0 {
															return int32(0)
														} else {
															v320 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
															if v320 == int32(0) {
																v325 = v318
																m.G0 = v11 + int32(128)
																return v325
															} else {
																F_pfree(m, v320)
																mBase = m.M
																v324 = m.ExcPending
																if v324 != 0 {
																	return int32(0)
																} else {
																	v325 = v318
																	m.G0 = v11 + int32(128)
																	return v325
																}
															}
														}
													}
												} else {
													v318 = F_make_result_opt_error(m, v11+int32(8), int32(0))
													mBase = m.M
													v319 = m.ExcPending
													if v319 != 0 {
														return int32(0)
													} else {
														v320 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
														if v320 == int32(0) {
															v325 = v318
															m.G0 = v11 + int32(128)
															return v325
														} else {
															F_pfree(m, v320)
															mBase = m.M
															v324 = m.ExcPending
															if v324 != 0 {
																return int32(0)
															} else {
																v325 = v318
																m.G0 = v11 + int32(128)
																return v325
															}
														}
													}
												}
											}
										} else {
											v312 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
											if v312 != 0 {
												F_pfree(m, v312)
												mBase = m.M
												v314 = m.ExcPending
												if v314 != 0 {
													return int32(0)
												} else {
													v318 = F_make_result_opt_error(m, v11+int32(8), int32(0))
													mBase = m.M
													v319 = m.ExcPending
													if v319 != 0 {
														return int32(0)
													} else {
														v320 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
														if v320 == int32(0) {
															v325 = v318
															m.G0 = v11 + int32(128)
															return v325
														} else {
															F_pfree(m, v320)
															mBase = m.M
															v324 = m.ExcPending
															if v324 != 0 {
																return int32(0)
															} else {
																v325 = v318
																m.G0 = v11 + int32(128)
																return v325
															}
														}
													}
												}
											} else {
												v318 = F_make_result_opt_error(m, v11+int32(8), int32(0))
												mBase = m.M
												v319 = m.ExcPending
												if v319 != 0 {
													return int32(0)
												} else {
													v320 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
													if v320 == int32(0) {
														v325 = v318
														m.G0 = v11 + int32(128)
														return v325
													} else {
														F_pfree(m, v320)
														mBase = m.M
														v324 = m.ExcPending
														if v324 != 0 {
															return int32(0)
														} else {
															v325 = v318
															m.G0 = v11 + int32(128)
															return v325
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
					v63 = F_make_result_opt_error(m, int32(_a_F_numeric_log_10), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v325 = v63
						m.G0 = v11 + int32(128)
						return v325
					}
				} else {
					v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
					v54 = v52
					if v54&int32(_a_F_numeric_log_1) != int32(_a_F_numeric_log_2) {
						if v22 == int32(-12288) {
							v69 = int32(1)
						} else {
							v69 = int32(-1)
						}
						v70 = int32(_a_F_numeric_log_2)
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
								if v71 == int32(_a_F_numeric_log_3) {
									v94 = v21 << (uint(v87) % 32) & int32(_a_F_numeric_log_4)
								} else {
									v94 = v71
								}
								if v94 == int32(_a_F_numeric_log_4) {
									v97 = int32(-1)
								} else {
									v97 = v87
								}
								v98 = v97
							}
						}
						v102 = v54 & int32(_a_F_numeric_log_1)
						if v102 == int32(_a_F_numeric_log_5) {
							v105 = int32(1)
						} else {
							v105 = int32(-1)
						}
						v106 = int32(_a_F_numeric_log_2)
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
								if v107 == int32(_a_F_numeric_log_3) {
									v131 = v102 << (uint(v124) % 32) & int32(_a_F_numeric_log_4)
								} else {
									v131 = v107
								}
								if v131 == int32(_a_F_numeric_log_4) {
									v134 = int32(-1)
								} else {
									v134 = v124
								}
								v135 = v134
							}
						}
						if v98|v135 < int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v339 = m.ExcPending
							if v339 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(352583810))
								mBase = m.M
								v342 = m.ExcPending
								if v342 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_numeric_log_6), int32(0))
									mBase = m.M
									v346 = m.ExcPending
									if v346 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_numeric_log_7), int32(4008), int32(_a_F_numeric_log_8))
										mBase = m.M
										v351 = m.ExcPending
										if v351 != 0 {
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
							v139 = int32(0)
							if base.B2i32(v98 == v139)|base.B2i32(v135 == v139) != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v355 = m.ExcPending
								if v355 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(352583810))
									mBase = m.M
									v358 = m.ExcPending
									if v358 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_numeric_log_9), int32(0))
										mBase = m.M
										v362 = m.ExcPending
										if v362 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_numeric_log_7), int32(4013), int32(_a_F_numeric_log_8))
											mBase = m.M
											v367 = m.ExcPending
											if v367 != 0 {
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
									if v54&int32(_a_F_numeric_log_1) == int32(_a_F_numeric_log_5) {
										v152 = F_make_result_opt_error(m, int32(_a_F_numeric_log_10), int32(0))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											v325 = v152
											m.G0 = v11 + int32(128)
											return v325
										}
									} else {
										v156 = F_make_result_opt_error(m, int32(_a_F_numeric_log_11), int32(0))
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											v325 = v156
											m.G0 = v11 + int32(128)
											return v325
										}
									}
								} else {
									v160 = F_make_result_opt_error(m, int32(_a_F_numeric_log_12), int32(0))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										v325 = v160
										m.G0 = v11 + int32(128)
										return v325
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(_a_F_numeric_log_10), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v325 = v63
							m.G0 = v11 + int32(128)
							return v325
						}
					}
				}
			}
		}
	}
}
func F_numeric_poly_var_pop(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13972(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
		if base.Ui32(int32(_a_F_numeric_round_0)) <= base.Ui32(v20) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v26 = F_palloc(m, int32(base.Ui32(v23)>>(uint(int32(2))%32)))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v30 = int32(base.Ui32(v28) >> (uint(int32(2)) % 32))
				if v30 == int32(0) {
					v181 = v26
				} else {
					base.MemoryCopy(m, v26, v16, v30)
					v181 = v26
				}
				m.G0 = v13 + int32(32)
				return v181
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(0)
			F_set_var_from_num(m, v16, v13+int32(8))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v42 = int32(-131073)
				if v34 <= v42 {
					v45 = v42
				} else {
					v45 = v34
				}
				if int32(_a_F_numeric_round_1) <= v45 {
					v48 = int32(_a_F_numeric_round_1)
				} else {
					v48 = v45
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v48
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				v53 = v48 + v50<<(uint(int32(2))%32)
				if v53+int32(4) < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
					v64 = v48 & int32(3)
					v68 = base.I32_div_s(v53+int32(7), int32(4))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					if v69 <= v68 {
						if base.B2i32(v64 == int32(0))|base.B2i32(v68 != v69) != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v68
							v83 = int32(1)
							v84 = v68 - v83
							v87 = v62 + v84<<(uint(v83)%32)
							v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87))))
							v89 = int32(2)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(v89)%32))+uint32(_c_F_numeric_round[0])))
							v92 = base.I32_rem_s(v88, v91)
							v93 = v88 - v92
							*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v93)
							v96 = base.I32_div_s(v91, v89)
							if v92 < v96 {
								v138 = v84
							} else {
								v99 = v91 + base.I32_extend16_s(v93)
								if int32(_a_F_numeric_round_2) < v99 {
									v104 = v99 + int32(_a_F_numeric_round_3)
								} else {
									v104 = v99
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v104)
								if v99 < int32(_a_F_numeric_round_4) {
									v138 = v84
								} else {
									v108 = v84
									v115 = v108
									for {
										v123 = int32(1)
										v124 = v115 - v123
										v127 = v62 + v124<<(uint(v123)%32)
										v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
										v132 = base.B2i32(int32(_a_F_numeric_round_5) < v130)
										if int32(_a_F_numeric_round_5) < v130 {
											v133 = int32(-9999)
										} else {
											v133 = v123
										}
										v134 = v133 + v130
										*(*uint16)(unsafe.Add(mBase, uint32(v127))) = uint16(v134)
										if int32(_a_F_numeric_round_5) < v130 {
											v115 = v124
											continue
										} else {
											break
										}
										break
									}
									v138 = v124
								}
							}
							if int32(0) <= v138 {
							} else {
								v148 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v50 + v148
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v68 + v148
								*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v62 - int32(2)
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v68
						if v64 != 0 {
							v83 = int32(1)
							v84 = v68 - v83
							v87 = v62 + v84<<(uint(v83)%32)
							v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87))))
							v89 = int32(2)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(v89)%32))+uint32(_c_F_numeric_round[0])))
							v92 = base.I32_rem_s(v88, v91)
							v93 = v88 - v92
							*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v93)
							v96 = base.I32_div_s(v91, v89)
							if v92 < v96 {
								v138 = v84
							} else {
								v99 = v91 + base.I32_extend16_s(v93)
								if int32(_a_F_numeric_round_2) < v99 {
									v104 = v99 + int32(_a_F_numeric_round_3)
								} else {
									v104 = v99
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v104)
								if v99 < int32(_a_F_numeric_round_4) {
									v138 = v84
								} else {
									v108 = v84
									v115 = v108
									for {
										v123 = int32(1)
										v124 = v115 - v123
										v127 = v62 + v124<<(uint(v123)%32)
										v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
										v132 = base.B2i32(int32(_a_F_numeric_round_5) < v130)
										if int32(_a_F_numeric_round_5) < v130 {
											v133 = int32(-9999)
										} else {
											v133 = v123
										}
										v134 = v133 + v130
										*(*uint16)(unsafe.Add(mBase, uint32(v127))) = uint16(v134)
										if int32(_a_F_numeric_round_5) < v130 {
											v115 = v124
											continue
										} else {
											break
										}
										break
									}
									v138 = v124
								}
							}
						} else {
							v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62+v68<<(uint(int32(1))%32)))))
							if v80 <= int32(_a_F_numeric_round_6) {
								v138 = v68
							} else {
								v108 = v68
								v115 = v108
								for {
									v123 = int32(1)
									v124 = v115 - v123
									v127 = v62 + v124<<(uint(v123)%32)
									v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
									v132 = base.B2i32(int32(_a_F_numeric_round_5) < v130)
									if int32(_a_F_numeric_round_5) < v130 {
										v133 = int32(-9999)
									} else {
										v133 = v123
									}
									v134 = v133 + v130
									*(*uint16)(unsafe.Add(mBase, uint32(v127))) = uint16(v134)
									if int32(_a_F_numeric_round_5) < v130 {
										v115 = v124
										continue
									} else {
										break
									}
									break
								}
								v138 = v124
							}
						}
						if int32(0) <= v138 {
						} else {
							v148 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v50 + v148
							*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v68 + v148
							*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v62 - int32(2)
						}
					}
				}
				if v34 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(0)
				} else {
				}
				v174 = F_make_result_opt_error(m, v13+int32(8), int32(0))
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return int32(0)
				} else {
					v176 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					if v176 == int32(0) {
						v181 = v174
						m.G0 = v13 + int32(32)
						return v181
					} else {
						F_pfree(m, v176)
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return int32(0)
						} else {
							v181 = v174
							m.G0 = v13 + int32(32)
							return v181
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		v16 = base.I32_extend16_s(v15)
		if base.Ui32(int32(_a_F_numeric_sqrt_0)) <= base.Ui32(v15) {
			if v16 == int32(-4096) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(369361026))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_numeric_sqrt_1), int32(0))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_sqrt_2), int32(3813), int32(_a_F_numeric_sqrt_3))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v24 = F_palloc(m, int32(base.Ui32(v21)>>(uint(int32(2))%32)))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v28 = int32(base.Ui32(v26) >> (uint(int32(2)) % 32))
					if v28 == int32(0) {
						v111 = v24
					} else {
						base.MemoryCopy(m, v24, v11, v28)
						v111 = v24
					}
					m.G0 = v8 + int32(48)
					return v111
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v38 = base.B2i32(int32(0) <= v16)
			if int32(0) <= v16 {
				v39 = int32(-8)
			} else {
				v39 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(base.Ui32(int32(base.Ui32(v32)>>(uint(int32(2))%32))+v39) >> (uint(int32(1)) % 32))
			if int32(0) <= v16 {
				v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+6)))
				v54 = v44
			} else {
				v54 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
			}
			v55 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v55
			v57 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v57
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v57
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v57
			v66 = base.B2i32(v16 < v55)
			if v16 < v55 {
				v67 = int32(6)
			} else {
				v67 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v11 + v67
			v75 = v15 & int32(_a_F_numeric_sqrt_0)
			if v75 == int32(_a_F_numeric_sqrt_4) {
				v78 = v15 << (uint(int32(1)) % 32) & int32(_a_F_numeric_sqrt_5)
			} else {
				v78 = v75
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v78
			if v16 < v55 {
				v86 = int32(base.Ui32(v15)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v86 = v15 & int32(_a_F_numeric_sqrt_6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v86
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v54
			v95 = int32(15) - v54<<(uint(int32(1))%32)
			if v86 < v95 {
				v97 = v95
			} else {
				v97 = v86
			}
			if int32(1000) <= v97 {
				v100 = int32(1000)
			} else {
				v100 = v97
			}
			F_sqrt_var(m, v8+int32(24), v8, v100)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v104 = F_make_result_opt_error(m, v8, int32(0))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					if v106 == int32(0) {
						v111 = v104
						m.G0 = v8 + int32(48)
						return v111
					} else {
						F_pfree(m, v106)
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							v111 = v104
							m.G0 = v8 + int32(48)
							return v111
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
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v74 int64
	_ = v74
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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
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
	return v377
L2:
	;
	v29 = int32(0)
	if base.B2i32(l2 == v29)|base.B2i32(int64(1) < v22) == v29 {
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
	v377 = int32(0)
	goto L1
L6:
	;
	goto L5
L7:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v36)
	v377 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v39)
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) < v41 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v56 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v56
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v76 = v14 + int32(72)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
	if v77 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v52 = F_make_result_opt_error(m, int32(_a_F_numeric_stddev_internal_0), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	if int64(0) < v44 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	if v47 <= int64(0) {
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
	v377 = v52
	goto L1
L17:
	;
	F_pfree(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v81 = F_palloc(m, int32(12))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L15
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v81
	v84 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v81))) = uint16(v84)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = v87 + int32(2)
	if v74 < int64(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v131
	v144 = v14 + int32(48)
	F_accum_sum_final(m, l0+int32(16), v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L15
	} else {
		goto L31
	}
L23:
	;
	v107 = v84
	v111 = v87 + int32(12)
	v113 = v101
	goto L28
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = int64(16384)
	v101 = int64(0) - v74
	goto L23
L25:
	;
	goto L26
L26:
	;
	v97 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v97
	if v74 == v97 {
		v131 = v84
		v134 = int32(0)
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v101 = v74
	goto L23
L28:
	;
	v116 = v111 - int32(2)
	v118 = base.I64_div_u_s(v113, int64(10000))
	v121 = v118*int64(55536) + v113
	*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v121)
	v124 = v107 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v113) {
		v107 = v124
		v111 = v116
		v113 = v118
		goto L28
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = v116
	v131 = v124
	v134 = v107
	goto L22
L30:
	;
	goto L29
L31:
	;
	v150 = v14 + int32(24)
	F_accum_sum_final(m, l0+int32(44), v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v153 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v153
	v160 = v14 + int32(72)
	F_sub_var(m, v160, int32(_a_F_numeric_stddev_internal_1), v14)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v166 = v164 << (uint(int32(1)) % 32)
	F_mul_var(m, v144, v144, v144, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	F_mul_var(m, v160, v150, v150, v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	F_sub_var(m, v150, v144, v150)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L15
	} else {
		goto L36
	}
L36:
	;
	v173 = int32(_a_F_numeric_stddev_internal_2)
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_stddev_internal[0]))
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_stddev_internal[1]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v182 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v366 != 0 {
		goto L107
	} else {
		goto L108
	}
L38:
	;
	if v218 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L39:
	;
	if v181 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	if v181 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v218 = int32(0)
	goto L38
L43:
	;
	goto L44
L44:
	;
	if v180 == int32(_a_F_numeric_stddev_internal_3) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v192 = int32(1)
	goto L47
L46:
	;
	v192 = int32(-1)
	goto L47
L47:
	;
	v218 = v192
	goto L38
L48:
	;
	if v193 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_stddev_internal[2]))
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_stddev_internal[3]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
	if v193 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v198 = int32(-1)
	goto L53
L52:
	;
	v198 = int32(1)
	goto L53
L53:
	;
	v218 = v198
	goto L38
L54:
	;
	if v180 == int32(_a_F_numeric_stddev_internal_3) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if v180 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v218 = int32(1)
	goto L38
L58:
	;
	goto L59
L59:
	;
	v208 = F_cmp_abs_common(m, v202, v182, v201, v200, v181, v199)
	mBase = m.M
	v218 = v208
	goto L38
L60:
	;
	v218 = int32(-1)
	goto L38
L61:
	;
	goto L62
L62:
	;
	v212 = F_cmp_abs_common(m, v200, v181, v199, v202, v182, v201)
	mBase = m.M
	v218 = v212
	goto L38
L63:
	;
	v223 = F_make_result_opt_error(m, int32(_a_F_numeric_stddev_internal_2), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L15
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v226 = v14 + int32(72)
	if l2 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v365 = v223
	goto L37
L67:
	;
	v227 = v14
	goto L69
L68:
	;
	v227 = v226
	goto L69
L69:
	;
	F_mul_var(m, v226, v227, v14, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L15
	} else {
		goto L70
	}
L70:
	;
	v231 = int32(0)
	v236 = v14 + int32(24)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	if v237 <= v231 {
		v265 = v231
		v274 = v231
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v275 <= int32(0) {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236)+20))
	v241 = v231
	goto L73
L73:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240+v241<<(uint(int32(1))%32)))))
	if v255 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v261 = int32(0)
	v265 = v261
	v274 = v261
	goto L71
L75:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v265 = v255
	v274 = v241 - v256
	goto L71
L76:
	;
	goto L77
L77:
	;
	v259 = v241 + int32(1)
	if v259 != v237 {
		v241 = v259
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	v317 = v14 + int32(48)
	v327 = (v306+v274+base.B2i32(base.I32_extend16_s(v265) <= base.I32_extend16_s(v309)))<<(uint(int32(2))%32) + int32(16)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	if v328 < v327 {
		goto L89
	} else {
		goto L90
	}
L80:
	;
	v306 = v231
	v309 = int32(0)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v281 = int32(0)
	goto L83
L83:
	;
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279+v281<<(uint(int32(1))%32)))))
	if v295 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v301 = int32(0)
	v306 = v301
	v309 = v301
	goto L79
L85:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v306 = v296 - v281
	v309 = v295
	goto L79
L86:
	;
	goto L87
L87:
	;
	v299 = v281 + int32(1)
	if v299 != v275 {
		v281 = v299
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	v330 = v327
	goto L91
L90:
	;
	v330 = v328
	goto L91
L91:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v331 < v330 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v333 = v330
	goto L94
L93:
	;
	v333 = v331
	goto L94
L94:
	;
	v334 = int32(0)
	if v334 < v333 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v337 = v333
	goto L97
L96:
	;
	v337 = v334
	goto L97
L97:
	;
	if int32(1000) <= v337 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v340 = int32(1000)
	goto L100
L99:
	;
	v340 = v337
	goto L100
L100:
	;
	v341 = int32(1)
	F_div_var(m, v14+int32(24), v14, v317, v340, v341, v341)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L15
	} else {
		goto L101
	}
L101:
	;
	if l1 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_sqrt_var(m, v317, v317, v340)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L15
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v352 = F_make_result_opt_error(m, v14+int32(48), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L15
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v365 = v352
	goto L37
L107:
	;
	F_pfree(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L15
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	if v369 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	F_pfree(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L15
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v372 == int32(0) {
		v377 = v365
		goto L1
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	F_pfree(m, v372)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L15
	} else {
		goto L116
	}
L116:
	;
	v377 = v365
	goto L1
}
func F_numeric_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
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
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
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
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
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
	var v305 int32
	_ = v305
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
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
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
	var v384 int32
	_ = v384
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(48)
	return v384
L4:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v50-int32(1)) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v27 == int32(18) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v38 = int32(1)
	if v21&v38 != 0 {
		v50 = int32(base.Ui32(v21)>>(uint(v38)%32)) - v38
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v30 = int32(16)
	goto L11
L10:
	;
	v30 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = int32(4)
	goto L14
L13:
	;
	v37 = v30
	goto L14
L14:
	;
	v50 = v37
	goto L4
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L16:
	;
	v56 = F_cstring_to_text(m, int32(_a_F_numeric_to_char_0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v62 = F_palloc0(m, v50<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v384 = v56
	goto L3
L20:
	;
	v68 = F_NUM_cache(m, v50, v11+int32(12), v19, v11+int32(11))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v70&int32(1024) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v365 = v62 + int32(4)
	F_NUM_processor(m, v68, v11+int32(12), v365, v356, int32(0), v361, v360, int32(1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L121
	}
L23:
	;
	v76 = F_numeric_int4_opt_error(m, v14, v11+int32(10))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v70&int32(_a_F_numeric_to_char_1) != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v78 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = int32(2147483647)
	goto L29
L28:
	;
	v79 = v76
	goto L29
L29:
	;
	v80 = F_int_to_roman(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v356 = v80
	v360 = v2
	v361 = int32(0)
	goto L22
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v86 = F_numeric_out_sci(m, v14, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	if v70&int32(2048) != 0 {
		goto L90
	} else {
		goto L91
	}
L34:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v192 == int32(45) {
		goto L65
	} else {
		goto L66
	}
L35:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v174 = v173 + v85
	v177 = F_palloc(m, v174+int32(7))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L61
	}
L36:
	;
	v88 = int32(_a_F_numeric_to_char_2)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_numeric_to_char[0])))
	if base.B2i32(v91 == int32(0))|base.B2i32(v91 != v94) != 0 {
		v112 = v91
		v113 = v94
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v112-v113 == int32(0) {
		goto L35
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v97 = v86
	v98 = v88
	goto L40
L40:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v102 == int32(0) {
		v112 = v102
		v113 = v101
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v112 = v102
	v113 = v101
	goto L38
L42:
	;
	v105 = int32(1)
	if v102 == v101 {
		v97 = v97 + v105
		v98 = v98 + v105
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v117 = int32(_a_F_numeric_to_char_3)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_numeric_to_char[1])))
	if base.B2i32(v120 == int32(0))|base.B2i32(v120 != v123) != 0 {
		v141 = v120
		v142 = v123
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v141-v142 == int32(0) {
		goto L35
	} else {
		goto L52
	}
L46:
	;
	goto L45
L47:
	;
	v126 = v86
	v127 = v117
	goto L48
L48:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v131 == int32(0) {
		v141 = v131
		v142 = v130
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v141 = v131
	v142 = v130
	goto L46
L50:
	;
	v134 = int32(1)
	if v131 == v130 {
		v126 = v126 + v134
		v127 = v127 + v134
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v146 = int32(_a_F_numeric_to_char_4)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_numeric_to_char[2])))
	if base.B2i32(v149 == int32(0))|base.B2i32(v149 != v152) != 0 {
		v170 = v149
		v171 = v152
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v170-v171 != 0 {
		goto L34
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v155 = v86
	v156 = v146
	goto L56
L56:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	if v160 == int32(0) {
		v170 = v160
		v171 = v159
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v170 = v160
	v171 = v159
	goto L54
L58:
	;
	v163 = int32(1)
	if v160 == v159 {
		v155 = v155 + v163
		v156 = v156 + v163
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L35
L61:
	;
	v180 = v174 + int32(6)
	if v180 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryFill(m, v177, int32(35), v180)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v177+v180))) = uint8(v184)
	v186 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v186)
	v189 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v177+v173)+1)) = uint8(v189)
	v356 = v177
	v360 = v2
	v361 = v184
	goto L22
L65:
	;
	v356 = v86
	v360 = v2
	v361 = int32(0)
	goto L22
L66:
	;
	goto L67
L67:
	;
	v196 = F_strlen(m, v86)
	mBase = m.M
	v199 = F_palloc(m, v196+int32(2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v201 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v201)
	v204 = v199 + int32(1)
	if (v86^v204)&int32(3) != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v356 = v199
	v360 = v2
	v361 = int32(0)
	goto L22
L70:
	;
	goto L69
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v258)
	if v258&int32(255) == int32(0) {
		goto L70
	} else {
		goto L86
	}
L72:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v257 = v86
	v258 = v210
	v259 = v204
	goto L71
L73:
	;
	goto L74
L74:
	;
	if v86&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v214 = v86
	v216 = v204
	goto L78
L76:
	;
	v228 = v86
	v230 = v204
	goto L77
L77:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v235 = int32(-2139062144)
	if (int32(16843008)-v232|v232)&v235 != v235 {
		v257 = v228
		v258 = v232
		v259 = v230
		goto L71
	} else {
		goto L82
	}
L78:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v217)
	if v217 == int32(0) {
		goto L70
	} else {
		goto L80
	}
L79:
	;
	v228 = v224
	v230 = v222
	goto L77
L80:
	;
	v221 = int32(1)
	v222 = v216 + v221
	v224 = v214 + v221
	if v224&int32(3) != 0 {
		v214 = v224
		v216 = v222
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v240 = v228
	v241 = v232
	v242 = v230
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v241
	v244 = int32(4)
	v245 = v242 + v244
	v247 = v240 + v244
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v252 = int32(-2139062144)
	if (int32(16843008)-v249|v249)&v252 == v252 {
		v240 = v247
		v241 = v249
		v242 = v245
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v257 = v247
	v258 = v249
	v259 = v245
	goto L71
L85:
	;
	goto L84
L86:
	;
	v266 = v257
	v268 = v259
	goto L87
L87:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)) = uint8(v269)
	v271 = int32(1)
	if v269 != 0 {
		v266 = v266 + v271
		v268 = v268 + v271
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L70
L89:
	;
	goto L88
L90:
	;
	v283 = int32(0)
	v287 = F_int64_to_numeric(m, int64(10))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	v305 = v14
	goto L92
L92:
	;
	v307 = int32(0)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v311 = F_DirectFunctionCall2Coll(m, int32(1259), v307, v305, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L100
	}
L93:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v291 = F_int64_to_numeric(m, base.I64_extend_i32_s(v289))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v293 = F_DirectFunctionCall2Coll(m, int32(1297), v283, v287, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v295 = F_pg_detoast_datum(m, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v297 = F_DirectFunctionCall2Coll(m, int32(1262), v283, v14, v295)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v299 = F_pg_detoast_datum(m, v297)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v289 + v301
	v305 = v299
	goto L92
L99:
	;
	if v317 == int32(45) {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v313 = F_pg_detoast_datum(m, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v315 = F_DirectFunctionCall1Coll(m, int32(618), v307, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	v319 = base.B2i32(v317 == int32(45))
	v320 = v315 + v319
	v321 = int32(46)
	v322 = F___strchrnul(m, v320, v321)
	mBase = m.M
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v324 == v321 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v328 != 0 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v328 = v322
	goto L106
L105:
	;
	v328 = int32(0)
	goto L106
L106:
	;
	goto L103
L107:
	;
	v331 = v328 - v320
	goto L99
L108:
	;
	goto L109
L109:
	;
	v330 = F_strlen(m, v320)
	mBase = m.M
	v331 = v330
	goto L99
L110:
	;
	v334 = int32(45)
	goto L112
L111:
	;
	v334 = int32(43)
	goto L112
L112:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v331 < v335 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v356 = v320
	v360 = v334
	v361 = v335 - v331
	goto L22
L114:
	;
	goto L115
L115:
	;
	if v331 <= v335 {
		v356 = v320
		v360 = v334
		v361 = int32(0)
		goto L22
	} else {
		goto L116
	}
L116:
	;
	v340 = v335 + v310
	v343 = F_palloc(m, v340+int32(2))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v346 = v340 + int32(1)
	if v346 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	base.MemoryFill(m, v343, int32(35), v346)
	goto L120
L119:
	;
	goto L120
L120:
	;
	v350 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v343+v346))) = uint8(v350)
	v353 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v343+v335))) = uint8(v353)
	v356 = v343
	v360 = v334
	v361 = v350
	goto L22
L121:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v370 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_pfree(m, v68)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v375 = F_strlen(m, v365)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v375<<(uint(int32(2))%32) + int32(16)
	v384 = v62
	goto L3
L125:
	;
	goto L124
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
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
		if base.Ui32(int32(_a_F_numeric_trunc_0)) <= base.Ui32(v14) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v20 = F_palloc(m, int32(base.Ui32(v17)>>(uint(int32(2))%32)))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v24 = int32(base.Ui32(v22) >> (uint(int32(2)) % 32))
				if v24 == int32(0) {
					v96 = v20
				} else {
					base.MemoryCopy(m, v20, v10, v24)
					v96 = v20
				}
				m.G0 = v7 + int32(32)
				return v96
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
			F_set_var_from_num(m, v10, v7+int32(8))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v36 = int32(-131072)
				if v28 <= v36 {
					v39 = v36
				} else {
					v39 = v28
				}
				if int32(_a_F_numeric_trunc_1) <= v39 {
					v42 = int32(_a_F_numeric_trunc_1)
				} else {
					v42 = v39
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v42
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v47 = v42 + v44<<(uint(int32(2))%32)
				if v47+int32(4) <= int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
				} else {
					v59 = base.I32_div_s(v47+int32(7), int32(4))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					if v60 < v59 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v59
						v64 = v42 & int32(3)
						if v64 == int32(0) {
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
							v71 = int32(2)
							v72 = v67 + v59<<(uint(int32(1))%32) - v71
							v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72))))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(v71)%32))+uint32(_c_F_numeric_trunc[0])))
							v77 = base.I32_rem_s(v73, v76)
							v78 = v73 - v77
							*(*uint16)(unsafe.Add(mBase, uint32(v72))) = uint16(v78)
						}
					}
				}
				if v28 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(0)
				} else {
				}
				v89 = F_make_result_opt_error(m, v7+int32(8), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
					if v91 == int32(0) {
						v96 = v89
						m.G0 = v7 + int32(32)
						return v96
					} else {
						F_pfree(m, v91)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = v89
							m.G0 = v7 + int32(32)
							return v96
						}
					}
				}
			}
		}
	}
}
