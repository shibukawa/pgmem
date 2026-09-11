package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_above(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+32))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_lt(v17, v18)
						}
					} else {
						return base.F64_lt(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_lt(v17, v18)
					}
				} else {
					return base.F64_lt(v17, v18)
				}
			}
		}
	}
}
func F_poly_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 float64
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v63 float64
	_ = v63
	var v67 int32
	_ = v67
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v158 float64
	_ = v158
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v182 float64
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 float64
	_ = v198
	var v201 int32
	_ = v201
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 - int32(-64)
	return v241
L2:
	;
	return int32(0)
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v27 = F_poly_overlap_internal(m, v20, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(0) < v31 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v225 = F_Float8GetDatum(m, float64(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L47
	}
L9:
	;
	v222 = F_Float8GetDatum(m, v198)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L46
	}
L10:
	;
	v34 = int32(40)
	v35 = v25 + v34
	v37 = v20 + v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v41 = v31
	v42 = v38
	v44 = v2
	v45 = v2
	v51 = float64(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v219)
	v241 = int32(0)
	goto L1
L13:
	;
	if int32(0) < v42 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v192&int32(1) != 0 {
		goto L9
	} else {
		goto L45
	}
L15:
	;
	if v44 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v188 = v41
	v189 = v42
	v192 = v45
	v198 = v51
	goto L17
L17:
	;
	v201 = v44 + int32(1)
	if v201 < v188 {
		v41 = v188
		v42 = v189
		v44 = v201
		v45 = v192
		v51 = v198
		goto L13
	} else {
		goto L44
	}
L18:
	;
	v55 = v44
	goto L20
L19:
	;
	v55 = v41
	goto L20
L20:
	;
	v56 = int32(4)
	v59 = int32(16)
	v60 = v55<<(uint(v56)%32) + v37 - v59
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+32)) = v61
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v60)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+40)) = v63
	v67 = v37 + v44<<(uint(v56)%32)
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = v68
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+56)) = v70
	v76 = v42<<(uint(v56)%32) + v35 - v59
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v76)))
	*(*float64)(unsafe.Add(mBase, uint32(v17))) = v77
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v76)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v79
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v25)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+16)) = v81
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v25)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+24)) = v83
	v88 = F_lseg_closept_lseg(m, int32(0), v15+int32(-32), v17)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v45&int32(1) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if int32(2) <= v109 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v107 = v88
	goto L22
L24:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v88)&int64(9223372036854775807)) {
		v107 = v51
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v51)&int64(9223372036854775807)) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if base.F64_gt(v51, v88) == int32(0) {
		v107 = v51
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v114 = int32(1)
	v124 = v107
	goto L31
L29:
	;
	v173 = v109
	v182 = v107
	goto L30
L30:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v188 = v185
	v189 = v173
	v192 = int32(1)
	v198 = v182
	goto L17
L31:
	;
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+32)) = v126
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v60)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+40)) = v128
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = v130
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+56)) = v132
	v136 = v35 + v114<<(uint(int32(4))%32)
	v138 = v136 - int32(16)
	v139 = *(*float64)(unsafe.Add(mBase, uint32(v138)))
	*(*float64)(unsafe.Add(mBase, uint32(v17))) = v139
	v141 = *(*float64)(unsafe.Add(mBase, uint32(v138)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v141
	v143 = *(*float64)(unsafe.Add(mBase, uint32(v136)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+16)) = v143
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v136)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+24)) = v145
	v150 = F_lseg_closept_lseg(m, int32(0), v15+int32(-32), v17)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L33
	}
L32:
	;
	v173 = v168
	v182 = v165
	goto L30
L33:
	;
	if base.Ui64(base.I64_reinterpret_f64(v150)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if base.F64_gt(v124, v150) != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v165 = v124
	goto L36
L36:
	;
	v167 = v114 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v167 < v168 {
		v114 = v167
		v124 = v165
		goto L31
	} else {
		goto L43
	}
L37:
	;
	v158 = v150
	goto L39
L38:
	;
	v158 = v124
	goto L39
L39:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v124)&int64(9223372036854775807)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v164 = v150
	goto L42
L41:
	;
	v164 = v158
	goto L42
L42:
	;
	v165 = v164
	goto L36
L43:
	;
	goto L32
L44:
	;
	goto L14
L45:
	;
	goto L12
L46:
	;
	v241 = v222
	goto L1
L47:
	;
	v241 = v225
	goto L1
}
func F_poly_left(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_gt(v17, v18)
						}
					} else {
						return base.F64_gt(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_gt(v17, v18)
					}
				} else {
					return base.F64_gt(v17, v18)
				}
			}
		}
	}
}
func F_poly_overabove(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+32))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_le(v17, v18)
						}
					} else {
						return base.F64_le(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_le(v17, v18)
					}
				} else {
					return base.F64_le(v17, v18)
				}
			}
		}
	}
}
func F_poly_overbelow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_ge(v17, v18)
						}
					} else {
						return base.F64_ge(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_ge(v17, v18)
					}
				} else {
					return base.F64_ge(v17, v18)
				}
			}
		}
	}
}
func F_poly_overlap_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_le(v22, base.F64_add(v23, float64(1e-06))) == v3 {
		v229 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 - int32(-64)
	return v229
L2:
	;
	v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v30 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F64_le(v29, base.F64_add(v30, float64(1e-06))) == int32(0) {
		v229 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_le(v36, base.F64_add(v37, float64(1e-06))) == int32(0) {
		v229 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v43 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.F64_le(v43, base.F64_add(v44, float64(1e-06))) == int32(0) {
		v229 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v51 = v18 + int32(-24)
	v53 = l0 + int32(40)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v59 = v53 + v54<<(uint(int32(4))%32) - int32(16)
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v62
	if v54 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v217 = l1 + int32(40)
	v218 = F_point_inside(m, v53, v205, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L20
	} else {
		goto L27
	}
L7:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v205 = v66
	goto L6
L8:
	;
	goto L9
L9:
	;
	v68 = l1 + int32(40)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v71 = v18 + int32(-48)
	v73 = v18 + int32(-16)
	v80 = v69
	v81 = v3
	v82 = v54
	goto L10
L10:
	;
	v95 = v80<<(uint(int32(4))%32) + v68 - int32(16)
	v105 = v81
	goto L12
L11:
	;
	v205 = v80
	goto L6
L12:
	;
	v117 = v53 + v105<<(uint(int32(4))%32)
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v117)))
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = v118
	v121 = v18 + int32(-8)
	v123 = v117 + int32(8)
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = v124
	v127 = v18 + int32(-56)
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v95)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v130
	if base.B2i32(v80 <= int32(0)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v140 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v192
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v117)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v194
	v197 = v105 + int32(1)
	if v197 < v82 {
		v105 = v197
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v154 = v68 + v140<<(uint(int32(4))%32)
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v71))) = v155
	v158 = v18 + int32(-40)
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v154)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = v159
	v164 = F_lseg_interpt_lseg(m, int32(0), v18+int32(-32), v20)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v179
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v181
	v184 = v105 + int32(1)
	if base.B2i32(v178 <= v184)|v164 == int32(0) {
		v80 = v174
		v81 = v184
		v82 = v178
		goto L10
	} else {
		goto L24
	}
L19:
	;
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v170
	v173 = v140 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v174 <= v173 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if v164 == int32(0) {
		v140 = v173
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	if v164 == int32(0) {
		v205 = v174
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v229 = int32(1)
	goto L1
L26:
	;
	goto L13
L27:
	;
	if v218 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v229 = int32(1)
	goto L1
L29:
	;
	goto L30
L30:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = F_point_inside(m, v217, v221, v53)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v229 = base.B2i32(v222 != int32(0))
	goto L1
}
