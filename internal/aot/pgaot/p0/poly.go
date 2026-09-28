package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_above(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+32))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.F64_lt(v14, v15))
						}
					} else {
						return base.I64_extend_i32_u(base.F64_lt(v14, v15))
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.F64_lt(v14, v15))
					}
				} else {
					return base.I64_extend_i32_u(base.F64_lt(v14, v15))
				}
			}
		}
	}
}
func F_poly_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 float64
	_ = v50
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v93 int32
	_ = v93
	var v96 float64
	_ = v96
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v109 float64
	_ = v109
	var v110 int32
	_ = v110
	var v126 float64
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v143 float64
	_ = v143
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v156 int32
	_ = v156
	var v159 float64
	_ = v159
	var v163 float64
	_ = v163
	var v165 float64
	_ = v165
	var v167 float64
	_ = v167
	var v172 float64
	_ = v172
	var v173 int32
	_ = v173
	var v180 float64
	_ = v180
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v204 float64
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v230 int32
	_ = v230
	var v244 float64
	_ = v244
	var v262 int64
	_ = v262
	v2 = int32(0)
	v15 = int64(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 - int32(-64)
	return v262
L2:
	;
	return int64(0)
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = F_poly_overlap_internal(m, v21, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v28 != 0 {
		v262 = v15
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if int32(0) < v30 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v262 = base.I64_reinterpret_f64(v244)
	goto L1
L8:
	;
	v33 = int32(40)
	v34 = v26 + v33
	v36 = v21 + v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v40 = v37
	v41 = v30
	v43 = v2
	v44 = v2
	v50 = float64(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v230)
	v262 = v15
	goto L1
L11:
	;
	v60 = v43
	goto L13
L12:
	;
	if v44 != 0 {
		v244 = v50
		goto L7
	} else {
		goto L46
	}
L13:
	;
	if base.B2i32(v40 <= int32(0)) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v213 = v60 + int32(1)
	if v213 < v41 {
		v60 = v213
		goto L13
	} else {
		goto L45
	}
L18:
	;
	v72 = v60
	goto L20
L19:
	;
	v72 = v41
	goto L20
L20:
	;
	v73 = int32(4)
	v75 = v36 + v72<<(uint(v73)%32)
	v76 = int32(16)
	v77 = v75 - v76
	v78 = *(*float64)(unsafe.Add(mBase, uint32(v77)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v78
	v80 = int32(8)
	v81 = v75 - v80
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v81)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v82
	v86 = v36 + v60<<(uint(v73)%32)
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+48)) = v87
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v86)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+56)) = v89
	v93 = v34 + v40<<(uint(v73)%32)
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v93-v76)))
	*(*float64)(unsafe.Add(mBase, uint32(v18))) = v96
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v93-v80)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+8)) = v100
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v26)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v102
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v26)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v104
	v109 = F_lseg_closept_lseg(m, int32(0), v16+int32(-32), v18)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v44 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if int32(2) <= v128 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v126 = v109
	goto L22
L24:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)) {
		v126 = v50
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v50)&int64(9223372036854775807)) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if base.F64_gt(v50, v109) == int32(0) {
		v126 = v50
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v137 = int32(1)
	v143 = v126
	goto L31
L29:
	;
	v194 = v128
	v204 = v126
	goto L30
L30:
	;
	v207 = int32(1)
	v209 = v60 + v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v209 < v210 {
		v40 = v194
		v41 = v210
		v43 = v209
		v44 = v207
		v50 = v204
		goto L11
	} else {
		goto L44
	}
L31:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v77)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v146
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v81)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v148
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+48)) = v150
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v86)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+56)) = v152
	v156 = v34 + v137<<(uint(int32(4))%32)
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v156-int32(16))))
	*(*float64)(unsafe.Add(mBase, uint32(v18))) = v159
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v156-int32(8))))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+8)) = v163
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v156)))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v165
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v156)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v167
	v172 = F_lseg_closept_lseg(m, int32(0), v16+int32(-32), v18)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L33
	}
L32:
	;
	v194 = v190
	v204 = v187
	goto L30
L33:
	;
	if base.Ui64(base.I64_reinterpret_f64(v172)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if base.F64_gt(v143, v172) != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v187 = v143
	goto L36
L36:
	;
	v189 = v137 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v189 < v190 {
		v137 = v189
		v143 = v187
		goto L31
	} else {
		goto L43
	}
L37:
	;
	v180 = v172
	goto L39
L38:
	;
	v180 = v143
	goto L39
L39:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v143)&int64(9223372036854775807)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v186 = v172
	goto L42
L41:
	;
	v186 = v180
	goto L42
L42:
	;
	v187 = v186
	goto L36
L43:
	;
	goto L32
L44:
	;
	v244 = v204
	goto L7
L45:
	;
	goto L14
L46:
	;
	goto L10
}
func F_poly_left(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.F64_gt(v14, v15))
						}
					} else {
						return base.I64_extend_i32_u(base.F64_gt(v14, v15))
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.F64_gt(v14, v15))
					}
				} else {
					return base.I64_extend_i32_u(base.F64_gt(v14, v15))
				}
			}
		}
	}
}
func F_poly_overabove(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+32))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+32))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.F64_le(v14, v15))
						}
					} else {
						return base.I64_extend_i32_u(base.F64_le(v14, v15))
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.F64_le(v14, v15))
					}
				} else {
					return base.I64_extend_i32_u(base.F64_le(v14, v15))
				}
			}
		}
	}
}
func F_poly_overbelow(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.F64_ge(v14, v15))
						}
					} else {
						return base.I64_extend_i32_u(base.F64_ge(v14, v15))
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.F64_ge(v14, v15))
					}
				} else {
					return base.I64_extend_i32_u(base.F64_ge(v14, v15))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_le(v18, base.F64_add(v19, float64(1e-06))) == v3 {
		v199 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 - int32(-64)
	return v199
L2:
	;
	v25 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v26 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F64_le(v25, base.F64_add(v26, float64(1e-06))) == int32(0) {
		v199 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_le(v32, base.F64_add(v33, float64(1e-06))) == int32(0) {
		v199 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.F64_le(v39, base.F64_add(v40, float64(1e-06))) == int32(0) {
		v199 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = l0 + int32(40)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = v47 + v48<<(uint(int32(4))%32) - int32(16)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v56
	if v48 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v187 = l1 + int32(40)
	v188 = F_point_inside(m, v47, v176, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L27
	}
L7:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v176 = v60
	goto L6
L8:
	;
	goto L9
L9:
	;
	v62 = l1 + int32(40)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v65 = v14 + int32(-48)
	v67 = v14 + int32(-16)
	v71 = v63
	v72 = v48
	v76 = v3
	goto L10
L10:
	;
	v85 = v62 + v71<<(uint(int32(4))%32) - int32(16)
	v96 = v76
	goto L12
L11:
	;
	v176 = v71
	goto L6
L12:
	;
	v103 = v47 + v96<<(uint(int32(4))%32)
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+8)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v110
	if base.B2i32(v71 <= int32(0)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v119 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v168
	v171 = v96 + int32(1)
	if v171 < v72 {
		v96 = v171
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v130 = v62 + v119<<(uint(int32(4))%32)
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v130)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v133
	v138 = F_lseg_interpt_lseg(m, int32(0), v14+int32(-32), v16)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v155
	v158 = v96 + int32(1)
	if base.B2i32(v152 <= v158)|v138 == int32(0) {
		v71 = v148
		v72 = v152
		v76 = v158
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
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v142
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v144
	v147 = v119 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v148 <= v147 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if v138 == int32(0) {
		v119 = v147
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	if v138 == int32(0) {
		v176 = v148
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v199 = int32(1)
	goto L1
L26:
	;
	goto L13
L27:
	;
	if v188 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v199 = int32(1)
	goto L1
L29:
	;
	goto L30
L30:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v192 = F_point_inside(m, v187, v191, v47)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v199 = base.B2i32(v192 != int32(0))
	goto L1
}
