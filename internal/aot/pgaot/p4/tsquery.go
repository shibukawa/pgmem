package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_cmp(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_CompareTSQ(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return v16
							}
						} else {
							return v16
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v16
						}
					} else {
						return v16
					}
				}
			}
		}
	}
}
func F_tsquery_eq(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_CompareTSQ(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v16 == int32(0))
							}
						} else {
							return base.B2i32(v16 == int32(0))
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v16 == int32(0))
						}
					} else {
						return base.B2i32(v16 == int32(0))
					}
				}
			}
		}
	}
}
func F_tsquery_ge(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_CompareTSQ(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v16^int32(-1)) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v16^int32(-1)) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v16^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v16^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_tsquery_lt(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_CompareTSQ(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v16) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v16) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v16) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v16) >> (uint(int32(31)) % 32))
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
	var v38 float64
	_ = v38
	var v45 int32
	_ = v45
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
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
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 float32
	_ = v148
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v156 float32
	_ = v156
	var v157 float64
	_ = v157
	var v160 float64
	_ = v160
	var v162 int32
	_ = v162
	var v164 float64
	_ = v164
	var v173 float64
	_ = v173
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v185 float64
	_ = v185
	var v188 float64
	_ = v188
	var v195 float64
	_ = v195
	var v197 float64
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 float32
	_ = v209
	var v211 int32
	_ = v211
	var v217 float64
	_ = v217
	var v218 int32
	_ = v218
	var v222 float64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 float64
	_ = v228
	var v229 int32
	_ = v229
	var v233 float64
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 float64
	_ = v258
	var v261 float64
	_ = v261
	var v264 float64
	_ = v264
	var v280 float64
	_ = v280
	var v281 float64
	_ = v281
	var v294 float64
	_ = v294
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
	return v294
L4:
	;
	v281 = float64(0)
	if base.F64_lt(v280, v281) != 0 {
		v294 = v281
		goto L3
	} else {
		goto L80
	}
L5:
	;
	v258 = float64(0.005)
	v261 = base.F64_promote_f32(base.F32_mul(l4, float32(0.5)))
	if base.F64_gt(v261, v258) != 0 {
		goto L77
	} else {
		goto L78
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
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	switch v211 - int32(1) {
	case 0:
		goto L68
	case 1, 3:
		goto L67
	case 2:
		goto L66
	default:
		goto L65
	}
L9:
	;
	v38 = float64(0.02)
	if l2 == int32(0) {
		v294 = v38
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
		goto L60
	} else {
		goto L61
	}
L12:
	;
	if l3 < int32(100) {
		v294 = v38
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v45 = int32(0)
	v50 = float64(0)
	v52 = float64(0)
	v55 = int32(0)
	goto L14
L14:
	;
	v62 = l2 + v45<<(uint(int32(3))%32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v164 = float64(0)
	if base.F64_lt(v154, v164) != 0 {
		v173 = v164
		goto L48
	} else {
		goto L49
	}
L16:
	;
	v156 = *(*float32)(unsafe.Add(mBase, uint32(v62)+4))
	v157 = base.F64_promote_f32(v156)
	v160 = base.F64_add(v50, base.F64_sub(v157, base.F64_mul(v50, v157)))
	v162 = v45 + int32(1)
	if v162 != l3 {
		v45 = v162
		v50 = v160
		v52 = v154
		v55 = v155
		goto L14
	} else {
		goto L47
	}
L17:
	;
	if v94 < v29 {
		v154 = v52
		v155 = v55
		goto L16
	} else {
		goto L28
	}
L18:
	;
	v67 = int32(4)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v69&int32(254) == int32(2) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v82 = int32(1)
	if v64&v82 != 0 {
		v94 = int32(base.Ui32(v64)>>(uint(v82)%32)) - v82
		goto L17
	} else {
		goto L27
	}
L21:
	;
	v78 = v67
	goto L23
L22:
	;
	v78 = base.B2i32(v69 == int32(18)) << (uint(v67) % 32)
	goto L23
L23:
	;
	if v69 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = v67
	goto L26
L25:
	;
	v81 = v78
	goto L26
L26:
	;
	v94 = v81
	goto L17
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v94 = int32(base.Ui32(v88)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L28:
	;
	v96 = int32(1)
	if v64&v96 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v100 = v96
	goto L31
L30:
	;
	v100 = int32(4)
	goto L31
L31:
	;
	v101 = v63 + v100
	if v29 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v145 != 0 {
		v154 = v52
		v155 = v55
		goto L16
	} else {
		goto L46
	}
L33:
	;
	v145 = int32(0)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v107 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v108 = v33
	v109 = v101
	v110 = v29
	v111 = v107
	goto L40
L37:
	;
	v133 = v101
	v137 = int32(0)
	goto L38
L38:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v145 = v137 - v138
	goto L32
L39:
	;
	v133 = v128
	v137 = v130
	goto L38
L40:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v111 != v113 {
		v128 = v109
		v130 = v111
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v128 = v122
	v130 = int32(0)
	goto L39
L42:
	;
	if v113 == int32(0) {
		v128 = v109
		v130 = v111
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v118 = v110 - int32(1)
	if v118 == int32(0) {
		v128 = v109
		v130 = v111
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v121 = int32(1)
	v122 = v109 + v121
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v123 != 0 {
		v108 = v108 + v121
		v109 = v122
		v110 = v118
		v111 = v123
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v148 = *(*float32)(unsafe.Add(mBase, uint32(v62)+4))
	v149 = base.F64_promote_f32(v148)
	v154 = base.F64_add(v52, base.F64_sub(v149, base.F64_mul(v52, v149)))
	v155 = v55 + int32(1)
	goto L16
L47:
	;
	goto L15
L48:
	;
	if base.F64_lt(v160, float64(0)) != 0 {
		v181 = v164
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if base.F64_gt(v154, float64(1)) == int32(0) {
		v173 = v154
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v173 = float64(1)
	goto L48
L51:
	;
	v182 = float64(0.005)
	v185 = base.F64_promote_f32(base.F32_mul(l4, float32(0.5)))
	if base.F64_gt(v185, v182) != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if base.F64_gt(v160, float64(1)) == int32(0) {
		v181 = v160
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v181 = float64(1)
	goto L51
L54:
	;
	v188 = v182
	goto L56
L55:
	;
	v188 = v185
	goto L56
L56:
	;
	v195 = base.F64_add(base.F64_mul(base.F64_sub(float64(1), v181), base.F64_div(base.F64_convert_i32_s(v155), base.F64_convert_i32_u(l3))), v173)
	if base.F64_gt(v188, v195) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v197 = v188
	goto L59
L58:
	;
	v197 = v195
	goto L59
L59:
	;
	v280 = v197
	goto L4
L60:
	;
	v294 = float64(0.005)
	goto L3
L61:
	;
	goto L62
L62:
	;
	v201 = int32(8)
	v205 = F_bsearch(m, v18+v201, l2, l3, v201, int32(1180))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v205 == int32(0) {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v209 = *(*float32)(unsafe.Add(mBase, uint32(v205)+4))
	v280 = base.F64_promote_f32(v209)
	goto L4
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L74
	}
L66:
	;
	v233 = F_tsquery_opr_selec(m, l0+int32(12), l1, l2, l3, l4)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L72
	}
L67:
	;
	v222 = F_tsquery_opr_selec(m, l0+int32(12), l1, l2, l3, l4)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	v217 = F_tsquery_opr_selec(m, l0+int32(12), l1, l2, l3, l4)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v280 = base.F64_sub(float64(1), v217)
	goto L4
L70:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v228 = F_tsquery_opr_selec(m, l0+v224*int32(12), l1, l2, l3, l4)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v280 = base.F64_mul(v222, v228)
	goto L4
L72:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v239 = F_tsquery_opr_selec(m, l0+v235*int32(12), l1, l2, l3, l4)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v280 = base.F64_sub(base.F64_add(v233, v239), base.F64_mul(v233, v239))
	goto L4
L74:
	;
	v248 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v248
	F_errmsg_internal(m, int32(492275), v18)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(506179), int32(417), int32(502517))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v264 = v258
	goto L79
L78:
	;
	v264 = v261
	goto L79
L79:
	;
	v280 = v264
	goto L4
L80:
	;
	if base.F64_gt(v280, float64(1)) == int32(0) {
		v294 = v280
		goto L3
	} else {
		goto L81
	}
L81:
	;
	v294 = float64(1)
	goto L3
}
