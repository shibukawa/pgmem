package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_cmp(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_copy(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_copy(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_CompareTSQ(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return v13
							}
						} else {
							return v13
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return v13
						}
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_tsquery_eq(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_copy(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_copy(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_CompareTSQ(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v13 == int32(0))
							}
						} else {
							return base.B2i32(v13 == int32(0))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v13 == int32(0))
						}
					} else {
						return base.B2i32(v13 == int32(0))
					}
				}
			}
		}
	}
}
func F_tsquery_ge(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_copy(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_copy(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_CompareTSQ(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v13^int32(-1)) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v13^int32(-1)) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v13^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v13^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_tsquery_lt(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_copy(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_copy(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_CompareTSQ(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v13) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v13) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v13) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v13) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_tsquery_opr_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32) float64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v149 float32
	_ = v149
	var v150 float64
	_ = v150
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v157 float32
	_ = v157
	var v158 float64
	_ = v158
	var v161 float64
	_ = v161
	var v163 int32
	_ = v163
	var v165 float64
	_ = v165
	var v174 float64
	_ = v174
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v186 float64
	_ = v186
	var v189 float64
	_ = v189
	var v196 float64
	_ = v196
	var v198 float64
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 float32
	_ = v210
	var v212 int32
	_ = v212
	var v218 float64
	_ = v218
	var v219 int32
	_ = v219
	var v223 float64
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 float64
	_ = v229
	var v230 int32
	_ = v230
	var v234 float64
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 float64
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 float64
	_ = v259
	var v262 float64
	_ = v262
	var v265 float64
	_ = v265
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v295 float64
	_ = v295
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	F_check_stack_depth(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0)
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v24 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v18 + int32(16)
	return v295
L4:
	;
	v282 = float64(0)
	if base.F64_lt(v281, v282) != 0 {
		v295 = v282
		goto L3
	} else {
		goto L78
	}
L5:
	;
	v259 = float64(0.005)
	v262 = base.F64_promote_f32(base.F32_mul(l4, float32(0.5)))
	if base.F64_gt(v262, v259) != 0 {
		goto L75
	} else {
		goto L76
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = v27 & int32(4095)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v29
	v33 = l1 + int32(base.Ui32(v27)>>(uint(int32(12))%32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v35 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	switch v212 - int32(1) {
	case 0:
		goto L66
	case 1, 3:
		goto L65
	case 2:
		goto L64
	default:
		goto L63
	}
L9:
	;
	if base.B2i32(l2 == int32(0))|base.B2i32(l3 < int32(100)) != 0 {
		v295 = float64(0.02)
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if l2 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v46 = int32(0)
	v51 = float64(0)
	v53 = float64(0)
	v58 = int32(0)
	goto L13
L13:
	;
	v63 = l2 + v46<<(uint(int32(3))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v65 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v165 = float64(0)
	if base.F64_lt(v155, v165) != 0 {
		v174 = v165
		goto L46
	} else {
		goto L47
	}
L15:
	;
	v157 = *(*float32)(unsafe.Add(mBase, uint32(v63)+4))
	v158 = base.F64_promote_f32(v157)
	v161 = base.F64_add(v51, base.F64_sub(v158, base.F64_mul(v51, v158)))
	v163 = v46 + int32(1)
	if v163 != l3 {
		v46 = v163
		v51 = v161
		v53 = v155
		v58 = v156
		goto L13
	} else {
		goto L45
	}
L16:
	;
	if v94 < v29 {
		v155 = v53
		v156 = v58
		goto L15
	} else {
		goto L27
	}
L17:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v71 == int32(18) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v82 = int32(1)
	if v65&v82 != 0 {
		v94 = int32(base.Ui32(v65)>>(uint(v82)%32)) - v82
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v74 = int32(16)
	goto L22
L21:
	;
	v74 = int32(0)
	goto L22
L22:
	;
	if base.Ui32((v71-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v81 = int32(4)
	goto L25
L24:
	;
	v81 = v74
	goto L25
L25:
	;
	v94 = v81
	goto L16
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v94 = int32(base.Ui32(v88)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	v96 = int32(1)
	if v65&v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v100 = v96
	goto L30
L29:
	;
	v100 = int32(4)
	goto L30
L30:
	;
	v101 = v64 + v100
	if v29 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v146 != 0 {
		v155 = v53
		v156 = v58
		goto L15
	} else {
		goto L44
	}
L32:
	;
	v146 = int32(0)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v107 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v108 = v33
	v109 = v101
	v110 = v29
	v111 = v107
	goto L39
L36:
	;
	v134 = v101
	v138 = int32(0)
	goto L37
L37:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v146 = v138 - v139
	goto L31
L38:
	;
	v134 = v129
	v138 = v131
	goto L37
L39:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if base.B2i32(v111 != v113)|base.B2i32(v113 == int32(0)) != 0 {
		v129 = v109
		v131 = v111
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v129 = v123
	v131 = int32(0)
	goto L38
L41:
	;
	v119 = v110 - int32(1)
	if v119 == int32(0) {
		v129 = v109
		v131 = v111
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v122 = int32(1)
	v123 = v109 + v122
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v124 != 0 {
		v108 = v108 + v122
		v109 = v123
		v110 = v119
		v111 = v124
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v149 = *(*float32)(unsafe.Add(mBase, uint32(v63)+4))
	v150 = base.F64_promote_f32(v149)
	v155 = base.F64_add(v53, base.F64_sub(v150, base.F64_mul(v53, v150)))
	v156 = v58 + int32(1)
	goto L15
L45:
	;
	goto L14
L46:
	;
	if base.F64_lt(v161, float64(0)) != 0 {
		v182 = v165
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if base.F64_gt(v155, float64(1)) == int32(0) {
		v174 = v155
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v174 = float64(1)
	goto L46
L49:
	;
	v183 = float64(0.005)
	v186 = base.F64_promote_f32(base.F32_mul(l4, float32(0.5)))
	if base.F64_gt(v186, v183) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if base.F64_gt(v161, float64(1)) == int32(0) {
		v182 = v161
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v182 = float64(1)
	goto L49
L52:
	;
	v189 = v183
	goto L54
L53:
	;
	v189 = v186
	goto L54
L54:
	;
	v196 = base.F64_add(base.F64_mul(base.F64_sub(float64(1), v182), base.F64_div(base.F64_convert_i32_s(v156), base.F64_convert_i32_u(l3))), v174)
	if base.F64_gt(v189, v196) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v198 = v189
	goto L57
L56:
	;
	v198 = v196
	goto L57
L57:
	;
	v281 = v198
	goto L4
L58:
	;
	v295 = float64(0.005)
	goto L3
L59:
	;
	goto L60
L60:
	;
	v202 = int32(8)
	v206 = F_bsearch(m, v18+v202, l2, l3, v202, int32(1164))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v206 == int32(0) {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v210 = *(*float32)(unsafe.Add(mBase, uint32(v206)+4))
	v281 = base.F64_promote_f32(v210)
	goto L4
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L72
	}
L64:
	;
	v234 = F_tsquery_opr_selec(m, l0+int32(12), l1, l2, l3, l4)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L70
	}
L65:
	;
	v223 = F_tsquery_opr_selec(m, l0+int32(12), l1, l2, l3, l4)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	v218 = F_tsquery_opr_selec(m, l0+int32(12), l1, l2, l3, l4)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v281 = base.F64_sub(float64(1), v218)
	goto L4
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = F_tsquery_opr_selec(m, l0+v225*int32(12), l1, l2, l3, l4)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v281 = base.F64_mul(v223, v229)
	goto L4
L70:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v240 = F_tsquery_opr_selec(m, l0+v236*int32(12), l1, l2, l3, l4)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v281 = base.F64_sub(base.F64_add(v234, v240), base.F64_mul(v234, v240))
	goto L4
L72:
	;
	v249 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v249
	F_errmsg_internal(m, int32(_a_F_tsquery_opr_selec_0), v18)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_tsquery_opr_selec_1), int32(417), int32(_a_F_tsquery_opr_selec_2))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v265 = v259
	goto L77
L76:
	;
	v265 = v262
	goto L77
L77:
	;
	v281 = v265
	goto L4
L78:
	;
	if base.F64_gt(v281, float64(1)) == int32(0) {
		v295 = v281
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v295 = float64(1)
	goto L3
}
