package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_VectorArraySet(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v6 == int32(1) {
		v9 = int32(6)
		v11 = int32(18)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
		if v13 == v11 {
			v16 = v11
		} else {
			v16 = int32(2)
		}
		if v13&int32(254) == int32(2) {
			v21 = v9
		} else {
			v21 = v16
		}
		if v13 == int32(1) {
			v24 = v9
		} else {
			v24 = v21
		}
		v33 = v24
	} else {
		v25 = int32(1)
		if v6&v25 != 0 {
			v33 = int32(base.Ui32(v6) >> (uint(v25) % 32))
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v33 = int32(base.Ui32(v29) >> (uint(int32(2)) % 32))
		}
	}
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v33) <= base.Ui32(v34) {
		if l1 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(476713), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errfinish(m, int32(343466), int32(326), int32(116085))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v38 <= l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(476713), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(343466), int32(326), int32(116085))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v33 != 0 {
					v43 = F__emscripten_memcpy_bulkmem(m, v40+l1*v34, l2, v33)
					mBase = m.M
				} else {
				}
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(476713), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_errfinish(m, int32(343466), int32(337), int32(115812))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
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
func F_vector_avg(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v74 float32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L21
	}
L2:
	;
	return int32(0)
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v18 != int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v25 != int32(701) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	if base.F64_eq(v28, float64(0)) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v11 + int32(16)
	return v82
L9:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	v82 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v34 = v21 - int32(1)
	F_CheckDim_3(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v39 = F_mul_size(m, int32(4), v34)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v41 = F_add_size(m, int32(8), v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v43 = F_palloc0(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+4)) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v41 << (uint(int32(2)) % 32)
	if v21 == int32(1) {
		v82 = v43
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v57 = int32(0)
	goto L17
L17:
	;
	v68 = v57 + int32(1)
	v72 = *(*float64)(unsafe.Add(mBase, uint32(v14+int32(24)+v68<<(uint(int32(3))%32))))
	v74 = base.F32_demote_f64(base.F64_div(v72, v28))
	*(*float32)(unsafe.Add(mBase, uint32(v43+int32(8)+v57<<(uint(int32(2))%32)))) = v74
	F_CheckElement_3(m, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	v82 = v43
	goto L8
L19:
	;
	if v34 != v68 {
		v57 = v68
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(344177)
	F_errmsg_internal(m, int32(26392), v11)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(520291), int32(169), int32(26883))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_concat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 float32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 float32
	_ = v74
	var v77 int32
	_ = v77
	var v80 float32
	_ = v80
	var v83 int32
	_ = v83
	var v86 float32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v123 float32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v177 float32
	_ = v177
	var v180 int32
	_ = v180
	var v188 float32
	_ = v188
	var v191 int32
	_ = v191
	var v199 float32
	_ = v199
	var v202 int32
	_ = v202
	var v210 float32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v250 float32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	v2 = int32(0)
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
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
			v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
			v23 = v21 + v22
			F_CheckDim_3(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_mul_size(m, int32(4), v23)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = F_add_size(m, int32(8), v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = F_palloc0(m, v30)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)) = uint16(v23)
							*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30 << (uint(int32(2)) % 32)
							v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
							if v38 <= int32(0) {
							} else {
								v42 = v38 & int32(3)
								v43 = int32(8)
								v44 = v32 + v43
								v46 = v14 + v43
								v47 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v38) {
									v53 = v47
									v58 = v2
									for {
										v65 = v53 << (uint(int32(2)) % 32)
										v68 = *(*float32)(unsafe.Add(mBase, uint32(v46+v65)))
										*(*float32)(unsafe.Add(mBase, uint32(v44+v65))) = v68
										v70 = int32(4)
										v71 = v65 | v70
										v74 = *(*float32)(unsafe.Add(mBase, uint32(v46+v71)))
										*(*float32)(unsafe.Add(mBase, uint32(v44+v71))) = v74
										v77 = v65 | int32(8)
										v80 = *(*float32)(unsafe.Add(mBase, uint32(v46+v77)))
										*(*float32)(unsafe.Add(mBase, uint32(v44+v77))) = v80
										v83 = v65 | int32(12)
										v86 = *(*float32)(unsafe.Add(mBase, uint32(v46+v83)))
										*(*float32)(unsafe.Add(mBase, uint32(v44+v83))) = v86
										v89 = v53 + v70
										v91 = v58 + v70
										if v91 != v38&int32(32764) {
											v53 = v89
											v58 = v91
											continue
										} else {
											break
										}
										break
									}
									v94 = v89
								} else {
									v94 = v47
								}
								if v42 == int32(0) {
								} else {
									v108 = v94
									v112 = v2
									for {
										v120 = v108 << (uint(int32(2)) % 32)
										v123 = *(*float32)(unsafe.Add(mBase, uint32(v46+v120)))
										*(*float32)(unsafe.Add(mBase, uint32(v44+v120))) = v123
										v125 = int32(1)
										v128 = v112 + v125
										if v128 != v42 {
											v108 = v108 + v125
											v112 = v128
											continue
										} else {
											break
										}
										break
									}
								}
							}
							v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
							if v142 <= int32(0) {
							} else {
								v146 = v142 & int32(3)
								v147 = int32(8)
								v148 = v32 + v147
								v150 = v19 + v147
								v151 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v142) {
									v159 = v151
									v163 = int32(0)
									for {
										v171 = int32(2)
										v177 = *(*float32)(unsafe.Add(mBase, uint32(v150+v159<<(uint(v171)%32))))
										*(*float32)(unsafe.Add(mBase, uint32(v148+(v38+v159)<<(uint(v171)%32)))) = v177
										v180 = v159 | int32(1)
										v188 = *(*float32)(unsafe.Add(mBase, uint32(v150+v180<<(uint(v171)%32))))
										*(*float32)(unsafe.Add(mBase, uint32(v148+(v180+v38)<<(uint(v171)%32)))) = v188
										v191 = v159 | v171
										v199 = *(*float32)(unsafe.Add(mBase, uint32(v150+v191<<(uint(v171)%32))))
										*(*float32)(unsafe.Add(mBase, uint32(v148+(v191+v38)<<(uint(v171)%32)))) = v199
										v202 = v159 | int32(3)
										v210 = *(*float32)(unsafe.Add(mBase, uint32(v150+v202<<(uint(v171)%32))))
										*(*float32)(unsafe.Add(mBase, uint32(v148+(v202+v38)<<(uint(v171)%32)))) = v210
										v212 = int32(4)
										v213 = v159 + v212
										v215 = v163 + v212
										if v215 != v142&int32(32764) {
											v159 = v213
											v163 = v215
											continue
										} else {
											break
										}
										break
									}
									v218 = v213
								} else {
									v218 = v151
								}
								if v146 == int32(0) {
								} else {
									v232 = v218
									v235 = v151
									for {
										v244 = int32(2)
										v250 = *(*float32)(unsafe.Add(mBase, uint32(v150+v232<<(uint(v244)%32))))
										*(*float32)(unsafe.Add(mBase, uint32(v148+(v38+v232)<<(uint(v244)%32)))) = v250
										v252 = int32(1)
										v255 = v235 + v252
										if v255 != v146 {
											v232 = v232 + v252
											v235 = v255
											continue
										} else {
											break
										}
										break
									}
								}
							}
							return v32
						}
					}
				}
			}
		}
	}
}
func F_vector_l2_squared_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 float32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 float32
	_ = v57
	var v59 float32
	_ = v59
	var v60 float32
	_ = v60
	var v63 float32
	_ = v63
	var v65 float32
	_ = v65
	var v66 float32
	_ = v66
	var v69 float32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 float32
	_ = v84
	var v89 int32
	_ = v89
	var v91 float32
	_ = v91
	var v93 float32
	_ = v93
	var v94 float32
	_ = v94
	var v99 float32
	_ = v99
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v10 = float32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
			if v24 == v25 {
				v27 = base.I32_extend16_s(v24)
				if v27 <= int32(0) {
					v112 = float64(0)
				} else {
					v31 = int32(8)
					v32 = v22 + v31
					v34 = v17 + v31
					if v27 == int32(1) {
						v75 = int32(0)
						v84 = v10
					} else {
						v41 = int32(0)
						v47 = int32(0)
						v50 = v10
						for {
							v52 = int32(2)
							v53 = v41 << (uint(v52) % 32)
							v55 = v53 | int32(4)
							v57 = *(*float32)(unsafe.Add(mBase, uint32(v34+v55)))
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v32+v55)))
							v60 = base.F32_sub(v57, v59)
							v63 = *(*float32)(unsafe.Add(mBase, uint32(v53+v34)))
							v65 = *(*float32)(unsafe.Add(mBase, uint32(v53+v32)))
							v66 = base.F32_sub(v63, v65)
							v69 = base.F32_add(base.F32_mul(v60, v60), base.F32_add(base.F32_mul(v66, v66), v50))
							v71 = v41 + v52
							v73 = v47 + v52
							if v73 != v27&int32(32766) {
								v41 = v71
								v47 = v73
								v50 = v69
								continue
							} else {
								break
							}
							break
						}
						v75 = v71
						v84 = v69
					}
					if v27&int32(1) != 0 {
						v89 = v75 << (uint(int32(2)) % 32)
						v91 = *(*float32)(unsafe.Add(mBase, uint32(v34+v89)))
						v93 = *(*float32)(unsafe.Add(mBase, uint32(v89+v32)))
						v94 = base.F32_sub(v91, v93)
						v99 = base.F32_add(base.F32_mul(v94, v94), v84)
					} else {
						v99 = v84
					}
					v112 = base.F64_promote_f32(v99)
				}
				v113 = F_Float8GetDatum(m, v112)
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(16)
					return v113
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v127
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v126
						F_errmsg(m, int32(501556), v14)
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(520291), int32(76), int32(160460))
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
			}
		}
	}
}
func F_vector_norm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 float64
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 float32
	_ = v40
	var v41 float64
	_ = v41
	var v43 float32
	_ = v43
	var v44 float64
	_ = v44
	var v46 float32
	_ = v46
	var v47 float64
	_ = v47
	var v49 float32
	_ = v49
	var v50 float64
	_ = v50
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 float64
	_ = v63
	var v72 int32
	_ = v72
	var v74 float64
	_ = v74
	var v78 int32
	_ = v78
	var v84 float32
	_ = v84
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 float64
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v2 = float64(0)
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+4)))
		if v15 <= int32(0) {
			v95 = v2
		} else {
			v19 = v11 + int32(8)
			v21 = v15 & int32(3)
			if base.Ui32(v15) < base.Ui32(int32(4)) {
				v61 = int32(0)
				v63 = v2
			} else {
				v28 = int32(0)
				v30 = v2
				v35 = v4
				for {
					v39 = v19 + v28<<(uint(int32(2))%32)
					v40 = *(*float32)(unsafe.Add(mBase, uint32(v39)+12))
					v41 = base.F64_promote_f32(v40)
					v43 = *(*float32)(unsafe.Add(mBase, uint32(v39)+8))
					v44 = base.F64_promote_f32(v43)
					v46 = *(*float32)(unsafe.Add(mBase, uint32(v39)+4))
					v47 = base.F64_promote_f32(v46)
					v49 = *(*float32)(unsafe.Add(mBase, uint32(v39)))
					v50 = base.F64_promote_f32(v49)
					v55 = base.F64_add(base.F64_mul(v41, v41), base.F64_add(base.F64_mul(v44, v44), base.F64_add(base.F64_mul(v47, v47), base.F64_add(base.F64_mul(v50, v50), v30))))
					v56 = int32(4)
					v57 = v28 + v56
					v59 = v35 + v56
					if v59 != v15&int32(32764) {
						v28 = v57
						v30 = v55
						v35 = v59
						continue
					} else {
						break
					}
					break
				}
				v61 = v57
				v63 = v55
			}
			if v21 == int32(0) {
				v95 = v63
			} else {
				v72 = v61
				v74 = v63
				v78 = v4
				for {
					v84 = *(*float32)(unsafe.Add(mBase, uint32(v19+v72<<(uint(int32(2))%32))))
					v85 = base.F64_promote_f32(v84)
					v87 = base.F64_add(base.F64_mul(v85, v85), v74)
					v88 = int32(1)
					v91 = v78 + v88
					if v91 != v21 {
						v72 = v72 + v88
						v74 = v87
						v78 = v91
						continue
					} else {
						break
					}
					break
				}
				v95 = v87
			}
		}
		v103 = F_Float8GetDatum(m, base.F64_sqrt(v95))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return int32(0)
		} else {
			return v103
		}
	}
}
func F_vector_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 float32
	_ = v29
	var v30 int32
	_ = v30
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int64
	_ = v213
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v382 int32
	_ = v382
	var v383 int64
	_ = v383
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v425 int64
	_ = v425
	var v429 int64
	_ = v429
	var v431 int32
	_ = v431
	var v434 int64
	_ = v434
	var v436 int32
	_ = v436
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v691 int32
	_ = v691
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v965 int64
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1167 float32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1273 int64
	_ = v1273
	var v1275 int64
	_ = v1275
	var v1276 int64
	_ = v1276
	var v1278 int64
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1282 int64
	_ = v1282
	var v1283 int64
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int64
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1309 int64
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1319 int32
	_ = v1319
	var v1320 int64
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1353 int64
	_ = v1353
	var v1357 int64
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1362 int64
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1375 int32
	_ = v1375
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1488 int64
	_ = v1488
	var v1490 int64
	_ = v1490
	var v1491 int64
	_ = v1491
	var v1493 int64
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1497 int64
	_ = v1497
	var v1498 int64
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1518 int64
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1523 int64
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1533 int32
	_ = v1533
	var v1534 int64
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1539 int64
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1565 int64
	_ = v1565
	var v1569 int64
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1574 int64
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1589 int32
	_ = v1589
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1604 int32
	_ = v1604
	var v1610 int32
	_ = v1610
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1802 int32
	_ = v1802
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1831 int32
	_ = v1831
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1873 int32
	_ = v1873
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1906 int32
	_ = v1906
	var v1911 int32
	_ = v1911
	var v1930 int32
	_ = v1930
	var v1937 int32
	_ = v1937
	var v1943 int32
	_ = v1943
	var v1953 int32
	_ = v1953
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1977 int32
	_ = v1977
	var v1981 int32
	_ = v1981
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1991 int32
	_ = v1991
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2105 int64
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2147 int32
	_ = v2147
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2203 int32
	_ = v2203
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2220 int32
	_ = v2220
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2231 int32
	_ = v2231
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2263 int32
	_ = v2263
	var v2266 int32
	_ = v2266
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
	v15 = F_mul_size(m, int32(16), v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = F_add_size(m, v15, int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = F_palloc(m, v18)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = int32(91)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v22)
	v24 = int32(1)
	v26 = v20 + v24
	if v14 <= int32(0) {
		v2295 = v26
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v2301 = int32(93)
	*(*uint16)(unsafe.Add(mBase, uint32(v2295))) = uint16(v2301)
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2303 != v10 {
		goto L346
	} else {
		goto L347
	}
L7:
	;
	v29 = *(*float32)(unsafe.Add(mBase, uint32(v10)+8))
	v30 = int32(0)
	v47 = base.I32_reinterpret_f32(v29)
	v49 = v47 & int32(8388607)
	v52 = int32(255)
	v53 = int32(base.Ui32(v47)>>(uint(int32(23))%32)) & v52
	if v49|v53 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v1150 = v1149 + v26
	if v14 == int32(1) {
		v2295 = v1150
		goto L6
	} else {
		goto L175
	}
L9:
	;
	v58 = base.B2i32(v53 != v52)
	goto L11
L10:
	;
	v58 = v30
	goto L11
L11:
	;
	if v58 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(int32(23)) < base.Ui32(v53-int32(127)) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1614])))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2)) = uint8(v62)
	v65 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1295])))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v65)
	v1149 = int32(3)
	goto L8
L16:
	;
	goto L17
L17:
	;
	if v47 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v70)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v74 = v26 + int32(base.Ui32(v47)>>(uint(int32(31))%32))
	if v53 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v74))) = int64(8751735898823355977)
	if v47 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v82 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v82)
	if v47 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v81 = int32(9)
	goto L26
L25:
	;
	v81 = int32(8)
	goto L26
L26:
	;
	v1149 = v81
	goto L8
L27:
	;
	v88 = int32(2)
	goto L29
L28:
	;
	v88 = int32(1)
	goto L29
L29:
	;
	v1149 = v88
	goto L8
L30:
	;
	v741 = v740 + v733
	v742 = int32(0)
	if v47 < v742 {
		goto L100
	} else {
		goto L101
	}
L31:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v681) {
		v723 = v681
		v733 = v691
		v740 = int32(8)
		goto L30
	} else {
		goto L91
	}
L32:
	;
	v106 = int32(2)
	v112 = v49 << (uint(v106) % 32)
	if v53 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v93 = int32(-1)
	v95 = int32(150) - v53
	if v49&(v93<<(uint(v95)%32)^v93) != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v681 = int32(base.Ui32(v49|int32(8388608)) >> (uint(v95) % 32))
	v691 = v30
	goto L31
L35:
	;
	v115 = v112 | int32(33554432)
	goto L37
L36:
	;
	v115 = v112
	goto L37
L37:
	;
	v116 = base.B2i32(v49 != int32(0)) | base.B2i32(base.Ui32(v53) < base.Ui32(v106)) ^ int32(-1) + v115
	v118 = v115 | int32(2)
	if v53 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v674 = v656 + v662
	v675 = v673 + v658
	if base.Ui32(v675) <= base.Ui32(int32(99999999)) {
		v681 = v675
		v691 = v674
		goto L31
	} else {
		goto L90
	}
L39:
	;
	v592 = int32(0)
	v593 = int32(10)
	v594 = base.I32_div_u_s(v582, v593)
	v596 = base.I32_div_u_s(v584, v593)
	if base.Ui32(v596) < base.Ui32(v594) {
		goto L84
	} else {
		goto L85
	}
L40:
	;
	v496 = int32(0)
	v497 = int32(10)
	v498 = base.I32_div_u_s(v486, v497)
	v500 = base.I32_div_u_s(v488, v497)
	if base.Ui32(v498) <= base.Ui32(v500) {
		goto L78
	} else {
		goto L79
	}
L41:
	;
	v122 = v53 - int32(152)
	goto L43
L42:
	;
	v122 = int32(-151)
	goto L43
L43:
	;
	if int32(0) <= v122 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v128 = int32(base.Ui32(v122*int32(78913)) >> (uint(int32(18)) % 32))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v128<<(uint(int32(3))%32))+uint32(_consts[1615])))
	v135 = v133 & int64(4294967295)
	v136 = base.I64_extend_i32_u(v116)
	v138 = int64(32)
	v140 = base.I32_wrap_i64(int64(base.Ui64(v135*v136) >> (uint(v138) % 64)))
	v142 = int64(base.Ui64(v133) >> (uint(v138) % 64))
	v143 = v142 * v136
	v145 = v140 + base.I32_wrap_i64(v143)
	v152 = v128 - v122
	v157 = v152 + int32(base.Ui32(v128*int32(1217359))>>(uint(int32(19))%32))
	v158 = int32(5) - v157
	v161 = v157 + int32(27)
	v163 = (base.B2i32(base.Ui32(v145) < base.Ui32(v140))+base.I32_wrap_i64(int64(base.Ui64(v143)>>(uint(v138)%64))))<<(uint(v158)%32) | int32(base.Ui32(v145)>>(uint(v161)%32))
	v164 = base.I64_extend_i32_u(v118)
	v168 = base.I32_wrap_i64(int64(base.Ui64(v135*v164) >> (uint(v138) % 64)))
	v169 = v164 * v142
	v171 = v168 + base.I32_wrap_i64(v169)
	v179 = (base.B2i32(base.Ui32(v171) < base.Ui32(v168))+base.I32_wrap_i64(int64(base.Ui64(v169)>>(uint(v138)%64))))<<(uint(v158)%32) | int32(base.Ui32(v171)>>(uint(v161)%32))
	v180 = base.I64_extend_i32_u(v115)
	v184 = base.I32_wrap_i64(int64(base.Ui64(v135*v180) >> (uint(v138) % 64)))
	v185 = v180 * v142
	v187 = v184 + base.I32_wrap_i64(v185)
	v195 = (base.B2i32(base.Ui32(v187) < base.Ui32(v184))+base.I32_wrap_i64(int64(base.Ui64(v185)>>(uint(v138)%64))))<<(uint(v158)%32) | int32(base.Ui32(v187)>>(uint(v161)%32))
	v196 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v122) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v341 = v122 * int32(-732923)
	v343 = int32(base.Ui32(v341) >> (uint(int32(20)) % 32))
	v344 = v343 + v122
	v348 = *(*int64)(unsafe.Add(mBase, uint32(int32(1880608)-v344<<(uint(int32(3))%32))))
	v350 = v348 & int64(4294967295)
	v351 = base.I64_extend_i32_u(v116)
	v353 = int64(32)
	v355 = base.I32_wrap_i64(int64(base.Ui64(v350*v351) >> (uint(v353) % 64)))
	v357 = int64(base.Ui64(v348) >> (uint(v353) % 64))
	v358 = v357 * v351
	v360 = v355 + base.I32_wrap_i64(v358)
	v371 = v343 - int32(base.Ui32(v344*int32(-1217359))>>(uint(int32(19))%32))
	v372 = int32(4) - v371
	v375 = v371 + int32(28)
	v377 = (base.B2i32(base.Ui32(v360) < base.Ui32(v355))+base.I32_wrap_i64(int64(base.Ui64(v358)>>(uint(v353)%64))))<<(uint(v372)%32) | int32(base.Ui32(v360)>>(uint(v375)%32))
	v378 = base.I64_extend_i32_u(v115)
	v382 = base.I32_wrap_i64(int64(base.Ui64(v350*v378) >> (uint(v353) % 64)))
	v383 = v378 * v357
	v385 = v382 + base.I32_wrap_i64(v383)
	v393 = (base.B2i32(base.Ui32(v385) < base.Ui32(v382))+base.I32_wrap_i64(int64(base.Ui64(v383)>>(uint(v353)%64))))<<(uint(v372)%32) | int32(base.Ui32(v385)>>(uint(v375)%32))
	v394 = base.I64_extend_i32_u(v118)
	v398 = base.I32_wrap_i64(int64(base.Ui64(v350*v394) >> (uint(v353) % 64)))
	v399 = v357 * v394
	v401 = v398 + base.I32_wrap_i64(v399)
	v409 = (base.B2i32(base.Ui32(v401) < base.Ui32(v398))+base.I32_wrap_i64(int64(base.Ui64(v399)>>(uint(v353)%64))))<<(uint(v372)%32) | int32(base.Ui32(v401)>>(uint(v375)%32))
	v411 = v409 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v341) {
		goto L68
	} else {
		goto L69
	}
L47:
	;
	v202 = int32(10)
	v203 = base.I32_div_u_s(v179-int32(1), v202)
	v205 = base.I32_div_u_s(v163, v202)
	if base.Ui32(v203) <= base.Ui32(v205) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v250 = v196
	goto L49
L49:
	;
	v255 = base.I32_rem_u_s(v115, int32(5))
	if v255 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v208 = v128 - int32(1)
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v208<<(uint(int32(3))%32))+uint32(_consts[1615])))
	v217 = int64(32)
	v219 = base.I32_wrap_i64(int64(base.Ui64(v213&int64(4294967295)*v180) >> (uint(v217) % 64)))
	v222 = int64(base.Ui64(v213)>>(uint(v217)%64)) * v180
	v224 = v219 + base.I32_wrap_i64(v222)
	v235 = v152 + int32(base.Ui32(v208*int32(1217359))>>(uint(int32(19))%32))
	v243 = base.I32_rem_u_s((base.B2i32(base.Ui32(v224) < base.Ui32(v219))+base.I32_wrap_i64(int64(base.Ui64(v222)>>(uint(v217)%64))))<<(uint(int32(6)-v235)%32)|int32(base.Ui32(v224)>>(uint(v235+int32(26))%32)), int32(10))
	v244 = v243
	goto L52
L51:
	;
	v244 = v196
	goto L52
L52:
	;
	if base.Ui32(int32(33)) < base.Ui32(v122) {
		v576 = v195
		v578 = v244
		v581 = v128
		v582 = v179
		v584 = v163
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v250 = v244
	goto L49
L54:
	;
	v260 = v115
	v262 = v196
	goto L57
L55:
	;
	goto L56
L56:
	;
	v286 = int32(0)
	v288 = base.I32_rem_u_s(v118, int32(5))
	if v288 == v286 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v278 = v262 + int32(1)
	v279 = int32(5)
	v280 = base.I32_div_u_s(v260, v279)
	v282 = base.I32_rem_u_s(v280, v279)
	if v282 == int32(0) {
		v260 = v280
		v262 = v278
		goto L57
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(v278) < base.Ui32(v128) {
		v576 = v195
		v578 = v250
		v581 = v128
		v582 = v179
		v584 = v163
		goto L39
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v480 = v195
	v482 = v250
	v485 = v128
	v486 = v179
	v488 = v163
	goto L40
L61:
	;
	v293 = v286
	v297 = v118
	goto L64
L62:
	;
	v320 = v286
	goto L63
L63:
	;
	v576 = v195
	v578 = v250
	v581 = v128
	v582 = v179 - base.B2i32(base.Ui32(v128) <= base.Ui32(v320))
	v584 = v163
	goto L39
L64:
	;
	v311 = v293 + int32(1)
	v312 = int32(5)
	v313 = base.I32_div_u_s(v297, v312)
	v315 = base.I32_rem_u_s(v313, v312)
	if v315 == int32(0) {
		v293 = v311
		v297 = v313
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v320 = v311
	goto L63
L66:
	;
	goto L65
L67:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v341) {
		v576 = v393
		v578 = v458
		v581 = v344
		v582 = v409
		v584 = v377
		goto L39
	} else {
		goto L75
	}
L68:
	;
	v414 = int32(10)
	v415 = base.I32_div_u_s(v411, v414)
	v417 = base.I32_div_u_s(v377, v414)
	if base.Ui32(v415) <= base.Ui32(v417) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v464 = v30
	goto L70
L70:
	;
	v480 = v393
	v482 = v464
	v485 = v344
	v486 = v411
	v488 = v377
	goto L40
L71:
	;
	v420 = int32(1) - v344
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v420<<(uint(int32(3))%32))+uint32(_consts[1616])))
	v429 = int64(32)
	v431 = base.I32_wrap_i64(int64(base.Ui64(v425&int64(4294967295)*v378) >> (uint(v429) % 64)))
	v434 = int64(base.Ui64(v425)>>(uint(v429)%64)) * v378
	v436 = v431 + base.I32_wrap_i64(v434)
	v449 = v343 + (int32(base.Ui32(v420*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v457 = base.I32_rem_u_s((base.B2i32(base.Ui32(v436) < base.Ui32(v431))+base.I32_wrap_i64(int64(base.Ui64(v434)>>(uint(v429)%64))))<<(uint(int32(4)-v449)%32)|int32(base.Ui32(v436)>>(uint(v449+int32(28))%32)), int32(10))
	v458 = v457
	goto L73
L72:
	;
	v458 = v30
	goto L73
L73:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v341) {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	v464 = v458
	goto L70
L75:
	;
	v470 = int32(-1)
	if v115&(v470<<(uint(v343-int32(1))%32)^v470) != 0 {
		v576 = v393
		v578 = v458
		v581 = v344
		v582 = v409
		v584 = v377
		goto L39
	} else {
		goto L76
	}
L76:
	;
	v480 = v393
	v482 = v458
	v485 = v344
	v486 = v409
	v488 = v377
	goto L40
L77:
	;
	v563 = v548 & int32(255)
	v656 = v544
	v658 = v546
	v662 = v485
	v673 = (v561|base.B2i32(v563 != int32(5))|v546)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v563)) | base.B2i32(v546 == v553)
	goto L38
L78:
	;
	v544 = v496
	v546 = v480
	v548 = v482
	v553 = v488
	v561 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v506 = v496
	v507 = v480
	v509 = v482
	v511 = v498
	v513 = v500
	v514 = int32(1)
	goto L81
L81:
	;
	v524 = v506 + int32(1)
	v525 = int32(10)
	v526 = base.I32_div_u_s(v507, v525)
	v529 = v507 - v526*v525
	v534 = v514 & base.B2i32(v509&int32(255) == int32(0))
	v536 = base.I32_div_u_s(v511, v525)
	v538 = base.I32_div_u_s(v513, v525)
	if base.Ui32(v538) < base.Ui32(v536) {
		v506 = v524
		v507 = v526
		v509 = v529
		v511 = v536
		v513 = v538
		v514 = v534
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v544 = v524
	v546 = v526
	v548 = v529
	v553 = v513
	v561 = v534 ^ int32(1)
	goto L77
L83:
	;
	goto L82
L84:
	;
	v600 = v592
	v601 = v576
	v602 = v594
	v604 = v596
	goto L87
L85:
	;
	v631 = v592
	v632 = v576
	v634 = v578
	v640 = v584
	goto L86
L86:
	;
	v656 = v631
	v658 = v632
	v662 = v581
	v673 = base.B2i32(v632 == v640) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v634&int32(255)))
	goto L38
L87:
	;
	v618 = v600 + int32(1)
	v619 = int32(10)
	v620 = base.I32_div_u_s(v601, v619)
	v622 = base.I32_div_u_s(v602, v619)
	v624 = base.I32_div_u_s(v604, v619)
	if base.Ui32(v624) < base.Ui32(v622) {
		v600 = v618
		v601 = v620
		v602 = v622
		v604 = v624
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v631 = v618
	v632 = v620
	v634 = v601 - v620*int32(10)
	v640 = v604
	goto L86
L89:
	;
	goto L88
L90:
	;
	v723 = v675
	v733 = v674
	v740 = int32(9)
	goto L30
L91:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v681) {
		v723 = v681
		v733 = v691
		v740 = int32(7)
		goto L30
	} else {
		goto L92
	}
L92:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v681) {
		v723 = v681
		v733 = v691
		v740 = int32(6)
		goto L30
	} else {
		goto L93
	}
L93:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v681) {
		v723 = v681
		v733 = v691
		v740 = int32(5)
		goto L30
	} else {
		goto L94
	}
L94:
	;
	if base.Ui32(int32(999)) < base.Ui32(v681) {
		v723 = v681
		v733 = v691
		v740 = int32(4)
		goto L30
	} else {
		goto L95
	}
L95:
	;
	if base.Ui32(int32(99)) < base.Ui32(v681) {
		v723 = v681
		v733 = v691
		v740 = int32(3)
		goto L30
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(int32(9)) < base.Ui32(v681) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v720 = int32(2)
	goto L99
L98:
	;
	v720 = int32(1)
	goto L99
L99:
	;
	v723 = v681
	v733 = v691
	v740 = v720
	goto L30
L100:
	;
	v745 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v745)
	v748 = int32(1)
	goto L102
L101:
	;
	v748 = v742
	goto L102
L102:
	;
	if base.Ui32(v741+int32(3)) <= base.Ui32(int32(9)) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v969 = int32(0)
	if base.Ui32(v723) < base.Ui32(int32(10000)) {
		goto L145
	} else {
		goto L146
	}
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v753))) = v965
	v967 = v964
	goto L103
L105:
	;
	v753 = v26 + v748
	v754 = int32(0)
	if v741 <= v754 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	if v733 != 0 {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v964 = int32(2) - v741
	v965 = int64(3472328296227679792)
	goto L104
L109:
	;
	goto L110
L110:
	;
	if int32(0) <= v733 {
		v964 = v754
		v965 = int64(3472328296227680304)
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v967 = int32(1)
	goto L103
L112:
	;
	v813 = int32(0)
	if base.Ui32(v797) < base.Ui32(int32(10000)) {
		goto L121
	} else {
		goto L122
	}
L113:
	;
	v797 = v723
	v803 = v740
	goto L112
L114:
	;
	goto L115
L115:
	;
	v766 = v723
	v771 = v740
	goto L116
L116:
	;
	if v766&int32(1) != 0 {
		v797 = v766
		v803 = v771
		goto L112
	} else {
		goto L118
	}
L117:
	;
	v797 = v766
	v803 = v771
	goto L112
L118:
	;
	v790 = base.I32_div_u_s(v766, int32(10))
	if int32(0)-v766 == v790*int32(-10) {
		v766 = v790
		v771 = v771 - int32(1)
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	if base.Ui32(v872) < base.Ui32(int32(100)) {
		goto L128
	} else {
		goto L129
	}
L121:
	;
	v870 = v813
	v872 = v797
	goto L120
L122:
	;
	goto L123
L123:
	;
	v820 = v813
	v821 = v797
	goto L124
L124:
	;
	v837 = v26 + v748 + v803 - v820
	v841 = base.I32_div_u_s(v821, int32(10000))
	v844 = v841*int32(-10000) + v821
	v845 = int32(100)
	v846 = base.I32_div_u_s(v844, v845)
	v847 = int32(1)
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v846<<(uint(v847)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v837-int32(3)))) = uint16(v851)
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v844-v846*v845)<<(uint(v847)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v837-v847))) = uint16(v862)
	v865 = v820 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v821) {
		v820 = v865
		v821 = v841
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v870 = v865
	v872 = v841
	goto L120
L126:
	;
	goto L125
L127:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v912) {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	v911 = v870
	v912 = v872
	goto L127
L129:
	;
	goto L130
L130:
	;
	v894 = int32(65535)
	v896 = int32(100)
	v897 = base.I32_div_u_s(v872&v894, v896)
	v907 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v872-v897*v896)&v894<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v26+v748+v803+(v870^int32(-1))))) = uint16(v907)
	v911 = v870 | int32(2)
	v912 = v897
	goto L127
L131:
	;
	v931 = int32(1)
	v932 = v741 - v931
	v933 = v26 + v748
	*(*uint8)(unsafe.Add(mBase, uint32(v933))) = uint8(v930)
	if base.Ui32(int32(2)) <= base.Ui32(v803) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v919 = v912 << (uint(int32(1)) % 32)
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+uint32(_consts[1618]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26+(v748+v803-v911)))) = uint8(v922)
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+uint32(_consts[1617]))))
	v930 = v926
	goto L131
L133:
	;
	goto L134
L134:
	;
	v930 = v912 | int32(48)
	goto L131
L135:
	;
	v938 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v933)+1)) = uint8(v938)
	v942 = v803 + int32(1)
	goto L137
L136:
	;
	v942 = v931
	goto L137
L137:
	;
	v943 = v942 + v748
	v944 = v26 + v943
	v945 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v944))) = uint8(v945)
	v950 = base.B2i32(v932 < int32(0))
	if v932 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v951 = int32(45)
	goto L140
L139:
	;
	v951 = int32(43)
	goto L140
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v944)+1)) = uint8(v951)
	if v932 < int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v955 = int32(1) - v741
	goto L143
L142:
	;
	v955 = v932
	goto L143
L143:
	;
	v960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v955<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v944)+2)) = uint16(v960)
	v1149 = v943 + int32(4)
	goto L8
L144:
	;
	if base.Ui32(v1028) < base.Ui32(int32(100)) {
		goto L152
	} else {
		goto L153
	}
L145:
	;
	v1027 = v969
	v1028 = v723
	goto L144
L146:
	;
	goto L147
L147:
	;
	v976 = v723
	v977 = v969
	goto L148
L148:
	;
	v993 = v753 + v967 + v740 - v977
	v994 = int32(4)
	v997 = base.I32_div_u_s(v976, int32(10000))
	v1000 = v997*int32(-10000) + v976
	v1001 = int32(100)
	v1002 = base.I32_div_u_s(v1000, v1001)
	v1003 = int32(1)
	v1007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1002<<(uint(v1003)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v993-v994))) = uint16(v1007)
	v1018 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1000-v1002*v1001)<<(uint(v1003)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v993-int32(2)))) = uint16(v1018)
	v1021 = v977 + v994
	if base.Ui32(int32(99999999)) < base.Ui32(v976) {
		v976 = v997
		v977 = v1021
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v1027 = v1021
	v1028 = v997
	goto L144
L150:
	;
	goto L149
L151:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1067) {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	v1067 = v1028
	v1068 = v1027
	goto L151
L153:
	;
	goto L154
L154:
	;
	v1048 = int32(2)
	v1050 = int32(65535)
	v1052 = int32(100)
	v1053 = base.I32_div_u_s(v1028&v1050, v1052)
	v1063 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1028-v1053*v1052)&v1050<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v753+v967+v740-v1027-v1048))) = uint16(v1063)
	v1067 = v1053
	v1068 = v1027 | v1048
	goto L151
L155:
	;
	v1086 = int32(1)
	if v967 == v1086 {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	v1080 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1067<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v753+v967+v740-v1068-int32(2)))) = uint16(v1080)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v1084 = v1067 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v753+v967))) = uint8(v1084)
	goto L155
L159:
	;
	v1149 = v1126 + int32(base.Ui32(v47)>>(uint(int32(31))%32))
	goto L8
L160:
	;
	if v741&int32(4) != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	if v733 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L163:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v753)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v753))) = v1091
	v1094 = int32(5)
	goto L165
L164:
	;
	v1094 = v1086
	goto L165
L165:
	;
	if v741&int32(2) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1097 = v753 + v1094
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1097))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1097-int32(1)))) = uint16(v1100)
	v1105 = v1094 | int32(2)
	goto L168
L167:
	;
	v1105 = v1094
	goto L168
L168:
	;
	if v741&int32(1) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1108 = v753 + v1105
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1108-int32(1)))) = uint8(v1111)
	goto L171
L170:
	;
	goto L171
L171:
	;
	v1115 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v753+v741))) = uint8(v1115)
	v1126 = v740 + int32(1)
	goto L159
L172:
	;
	v1123 = int32(2) - v733
	goto L174
L173:
	;
	v1123 = v741
	goto L174
L174:
	;
	v1126 = v1123
	goto L159
L175:
	;
	v1156 = v1150
	v1159 = v24
	goto L176
L176:
	;
	v1162 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1156))) = uint8(v1162)
	v1167 = *(*float32)(unsafe.Add(mBase, uint32(v10+int32(8)+v1159<<(uint(int32(2))%32))))
	v1169 = v1156 + int32(1)
	v1170 = int32(0)
	v1187 = base.I32_reinterpret_f32(v1167)
	v1189 = v1187 & int32(8388607)
	v1192 = int32(255)
	v1193 = int32(base.Ui32(v1187)>>(uint(int32(23))%32)) & v1192
	if v1189|v1193 != 0 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v2295 = v2290
	goto L6
L178:
	;
	v2290 = v2289 + v1169
	v2292 = v1159 + int32(1)
	if v2292 != v14 {
		v1156 = v2290
		v1159 = v2292
		goto L176
	} else {
		goto L345
	}
L179:
	;
	v1198 = base.B2i32(v1193 != v1192)
	goto L181
L180:
	;
	v1198 = v1170
	goto L181
L181:
	;
	if v1198 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	if v1189 != 0 {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L184
L184:
	;
	if base.Ui32(int32(23)) < base.Ui32(v1193-int32(127)) {
		goto L202
	} else {
		goto L203
	}
L185:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1614])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1169)+2)) = uint8(v1202)
	v1205 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1295])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1169))) = uint16(v1205)
	v2289 = int32(3)
	goto L178
L186:
	;
	goto L187
L187:
	;
	if v1187 < int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1210 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1169))) = uint8(v1210)
	goto L190
L189:
	;
	goto L190
L190:
	;
	v1214 = v1169 + int32(base.Ui32(v1187)>>(uint(int32(31))%32))
	if v1193 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1214))) = int64(8751735898823355977)
	if v1187 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	goto L193
L193:
	;
	v1222 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1214))) = uint8(v1222)
	if v1187 < int32(0) {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v1221 = int32(9)
	goto L196
L195:
	;
	v1221 = int32(8)
	goto L196
L196:
	;
	v2289 = v1221
	goto L178
L197:
	;
	v1228 = int32(2)
	goto L199
L198:
	;
	v1228 = int32(1)
	goto L199
L199:
	;
	v2289 = v1228
	goto L178
L200:
	;
	v1881 = v1880 + v1873
	v1882 = int32(0)
	if v1187 < v1882 {
		goto L270
	} else {
		goto L271
	}
L201:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v1821) {
		v1863 = v1821
		v1873 = v1831
		v1880 = int32(8)
		goto L200
	} else {
		goto L261
	}
L202:
	;
	v1246 = int32(2)
	v1252 = v1189 << (uint(v1246) % 32)
	if v1193 != 0 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v1233 = int32(-1)
	v1235 = int32(150) - v1193
	if v1189&(v1233<<(uint(v1235)%32)^v1233) != 0 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1821 = int32(base.Ui32(v1189|int32(8388608)) >> (uint(v1235) % 32))
	v1831 = v1170
	goto L201
L205:
	;
	v1255 = v1252 | int32(33554432)
	goto L207
L206:
	;
	v1255 = v1252
	goto L207
L207:
	;
	v1256 = base.B2i32(v1189 != int32(0)) | base.B2i32(base.Ui32(v1193) < base.Ui32(v1246)) ^ int32(-1) + v1255
	v1258 = v1255 | int32(2)
	if v1193 != 0 {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v1814 = v1796 + v1802
	v1815 = v1813 + v1798
	if base.Ui32(v1815) <= base.Ui32(int32(99999999)) {
		v1821 = v1815
		v1831 = v1814
		goto L201
	} else {
		goto L260
	}
L209:
	;
	v1732 = int32(0)
	v1733 = int32(10)
	v1734 = base.I32_div_u_s(v1722, v1733)
	v1736 = base.I32_div_u_s(v1724, v1733)
	if base.Ui32(v1736) < base.Ui32(v1734) {
		goto L254
	} else {
		goto L255
	}
L210:
	;
	v1636 = int32(0)
	v1637 = int32(10)
	v1638 = base.I32_div_u_s(v1626, v1637)
	v1640 = base.I32_div_u_s(v1628, v1637)
	if base.Ui32(v1638) <= base.Ui32(v1640) {
		goto L248
	} else {
		goto L249
	}
L211:
	;
	v1262 = v1193 - int32(152)
	goto L213
L212:
	;
	v1262 = int32(-151)
	goto L213
L213:
	;
	if int32(0) <= v1262 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1268 = int32(base.Ui32(v1262*int32(78913)) >> (uint(int32(18)) % 32))
	v1273 = *(*int64)(unsafe.Add(mBase, uint32(v1268<<(uint(int32(3))%32))+uint32(_consts[1615])))
	v1275 = v1273 & int64(4294967295)
	v1276 = base.I64_extend_i32_u(v1256)
	v1278 = int64(32)
	v1280 = base.I32_wrap_i64(int64(base.Ui64(v1275*v1276) >> (uint(v1278) % 64)))
	v1282 = int64(base.Ui64(v1273) >> (uint(v1278) % 64))
	v1283 = v1282 * v1276
	v1285 = v1280 + base.I32_wrap_i64(v1283)
	v1292 = v1268 - v1262
	v1297 = v1292 + int32(base.Ui32(v1268*int32(1217359))>>(uint(int32(19))%32))
	v1298 = int32(5) - v1297
	v1301 = v1297 + int32(27)
	v1303 = (base.B2i32(base.Ui32(v1285) < base.Ui32(v1280))+base.I32_wrap_i64(int64(base.Ui64(v1283)>>(uint(v1278)%64))))<<(uint(v1298)%32) | int32(base.Ui32(v1285)>>(uint(v1301)%32))
	v1304 = base.I64_extend_i32_u(v1258)
	v1308 = base.I32_wrap_i64(int64(base.Ui64(v1275*v1304) >> (uint(v1278) % 64)))
	v1309 = v1304 * v1282
	v1311 = v1308 + base.I32_wrap_i64(v1309)
	v1319 = (base.B2i32(base.Ui32(v1311) < base.Ui32(v1308))+base.I32_wrap_i64(int64(base.Ui64(v1309)>>(uint(v1278)%64))))<<(uint(v1298)%32) | int32(base.Ui32(v1311)>>(uint(v1301)%32))
	v1320 = base.I64_extend_i32_u(v1255)
	v1324 = base.I32_wrap_i64(int64(base.Ui64(v1275*v1320) >> (uint(v1278) % 64)))
	v1325 = v1320 * v1282
	v1327 = v1324 + base.I32_wrap_i64(v1325)
	v1335 = (base.B2i32(base.Ui32(v1327) < base.Ui32(v1324))+base.I32_wrap_i64(int64(base.Ui64(v1325)>>(uint(v1278)%64))))<<(uint(v1298)%32) | int32(base.Ui32(v1327)>>(uint(v1301)%32))
	v1336 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1262) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v1481 = v1262 * int32(-732923)
	v1483 = int32(base.Ui32(v1481) >> (uint(int32(20)) % 32))
	v1484 = v1483 + v1262
	v1488 = *(*int64)(unsafe.Add(mBase, uint32(int32(1880608)-v1484<<(uint(int32(3))%32))))
	v1490 = v1488 & int64(4294967295)
	v1491 = base.I64_extend_i32_u(v1256)
	v1493 = int64(32)
	v1495 = base.I32_wrap_i64(int64(base.Ui64(v1490*v1491) >> (uint(v1493) % 64)))
	v1497 = int64(base.Ui64(v1488) >> (uint(v1493) % 64))
	v1498 = v1497 * v1491
	v1500 = v1495 + base.I32_wrap_i64(v1498)
	v1511 = v1483 - int32(base.Ui32(v1484*int32(-1217359))>>(uint(int32(19))%32))
	v1512 = int32(4) - v1511
	v1515 = v1511 + int32(28)
	v1517 = (base.B2i32(base.Ui32(v1500) < base.Ui32(v1495))+base.I32_wrap_i64(int64(base.Ui64(v1498)>>(uint(v1493)%64))))<<(uint(v1512)%32) | int32(base.Ui32(v1500)>>(uint(v1515)%32))
	v1518 = base.I64_extend_i32_u(v1255)
	v1522 = base.I32_wrap_i64(int64(base.Ui64(v1490*v1518) >> (uint(v1493) % 64)))
	v1523 = v1518 * v1497
	v1525 = v1522 + base.I32_wrap_i64(v1523)
	v1533 = (base.B2i32(base.Ui32(v1525) < base.Ui32(v1522))+base.I32_wrap_i64(int64(base.Ui64(v1523)>>(uint(v1493)%64))))<<(uint(v1512)%32) | int32(base.Ui32(v1525)>>(uint(v1515)%32))
	v1534 = base.I64_extend_i32_u(v1258)
	v1538 = base.I32_wrap_i64(int64(base.Ui64(v1490*v1534) >> (uint(v1493) % 64)))
	v1539 = v1497 * v1534
	v1541 = v1538 + base.I32_wrap_i64(v1539)
	v1549 = (base.B2i32(base.Ui32(v1541) < base.Ui32(v1538))+base.I32_wrap_i64(int64(base.Ui64(v1539)>>(uint(v1493)%64))))<<(uint(v1512)%32) | int32(base.Ui32(v1541)>>(uint(v1515)%32))
	v1551 = v1549 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v1481) {
		goto L238
	} else {
		goto L239
	}
L217:
	;
	v1342 = int32(10)
	v1343 = base.I32_div_u_s(v1319-int32(1), v1342)
	v1345 = base.I32_div_u_s(v1303, v1342)
	if base.Ui32(v1343) <= base.Ui32(v1345) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	v1390 = v1336
	goto L219
L219:
	;
	v1395 = base.I32_rem_u_s(v1255, int32(5))
	if v1395 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	v1348 = v1268 - int32(1)
	v1353 = *(*int64)(unsafe.Add(mBase, uint32(v1348<<(uint(int32(3))%32))+uint32(_consts[1615])))
	v1357 = int64(32)
	v1359 = base.I32_wrap_i64(int64(base.Ui64(v1353&int64(4294967295)*v1320) >> (uint(v1357) % 64)))
	v1362 = int64(base.Ui64(v1353)>>(uint(v1357)%64)) * v1320
	v1364 = v1359 + base.I32_wrap_i64(v1362)
	v1375 = v1292 + int32(base.Ui32(v1348*int32(1217359))>>(uint(int32(19))%32))
	v1383 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1364) < base.Ui32(v1359))+base.I32_wrap_i64(int64(base.Ui64(v1362)>>(uint(v1357)%64))))<<(uint(int32(6)-v1375)%32)|int32(base.Ui32(v1364)>>(uint(v1375+int32(26))%32)), int32(10))
	v1384 = v1383
	goto L222
L221:
	;
	v1384 = v1336
	goto L222
L222:
	;
	if base.Ui32(int32(33)) < base.Ui32(v1262) {
		v1716 = v1335
		v1718 = v1384
		v1721 = v1268
		v1722 = v1319
		v1724 = v1303
		goto L209
	} else {
		goto L223
	}
L223:
	;
	v1390 = v1384
	goto L219
L224:
	;
	v1400 = v1255
	v1402 = v1336
	goto L227
L225:
	;
	goto L226
L226:
	;
	v1426 = int32(0)
	v1428 = base.I32_rem_u_s(v1258, int32(5))
	if v1428 == v1426 {
		goto L231
	} else {
		goto L232
	}
L227:
	;
	v1418 = v1402 + int32(1)
	v1419 = int32(5)
	v1420 = base.I32_div_u_s(v1400, v1419)
	v1422 = base.I32_rem_u_s(v1420, v1419)
	if v1422 == int32(0) {
		v1400 = v1420
		v1402 = v1418
		goto L227
	} else {
		goto L229
	}
L228:
	;
	if base.Ui32(v1418) < base.Ui32(v1268) {
		v1716 = v1335
		v1718 = v1390
		v1721 = v1268
		v1722 = v1319
		v1724 = v1303
		goto L209
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	v1620 = v1335
	v1622 = v1390
	v1625 = v1268
	v1626 = v1319
	v1628 = v1303
	goto L210
L231:
	;
	v1433 = v1426
	v1437 = v1258
	goto L234
L232:
	;
	v1460 = v1426
	goto L233
L233:
	;
	v1716 = v1335
	v1718 = v1390
	v1721 = v1268
	v1722 = v1319 - base.B2i32(base.Ui32(v1268) <= base.Ui32(v1460))
	v1724 = v1303
	goto L209
L234:
	;
	v1451 = v1433 + int32(1)
	v1452 = int32(5)
	v1453 = base.I32_div_u_s(v1437, v1452)
	v1455 = base.I32_rem_u_s(v1453, v1452)
	if v1455 == int32(0) {
		v1433 = v1451
		v1437 = v1453
		goto L234
	} else {
		goto L236
	}
L235:
	;
	v1460 = v1451
	goto L233
L236:
	;
	goto L235
L237:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v1481) {
		v1716 = v1533
		v1718 = v1598
		v1721 = v1484
		v1722 = v1549
		v1724 = v1517
		goto L209
	} else {
		goto L245
	}
L238:
	;
	v1554 = int32(10)
	v1555 = base.I32_div_u_s(v1551, v1554)
	v1557 = base.I32_div_u_s(v1517, v1554)
	if base.Ui32(v1555) <= base.Ui32(v1557) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v1604 = v1170
	goto L240
L240:
	;
	v1620 = v1533
	v1622 = v1604
	v1625 = v1484
	v1626 = v1551
	v1628 = v1517
	goto L210
L241:
	;
	v1560 = int32(1) - v1484
	v1565 = *(*int64)(unsafe.Add(mBase, uint32(v1560<<(uint(int32(3))%32))+uint32(_consts[1616])))
	v1569 = int64(32)
	v1571 = base.I32_wrap_i64(int64(base.Ui64(v1565&int64(4294967295)*v1518) >> (uint(v1569) % 64)))
	v1574 = int64(base.Ui64(v1565)>>(uint(v1569)%64)) * v1518
	v1576 = v1571 + base.I32_wrap_i64(v1574)
	v1589 = v1483 + (int32(base.Ui32(v1560*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v1597 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1576) < base.Ui32(v1571))+base.I32_wrap_i64(int64(base.Ui64(v1574)>>(uint(v1569)%64))))<<(uint(int32(4)-v1589)%32)|int32(base.Ui32(v1576)>>(uint(v1589+int32(28))%32)), int32(10))
	v1598 = v1597
	goto L243
L242:
	;
	v1598 = v1170
	goto L243
L243:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v1481) {
		goto L237
	} else {
		goto L244
	}
L244:
	;
	v1604 = v1598
	goto L240
L245:
	;
	v1610 = int32(-1)
	if v1255&(v1610<<(uint(v1483-int32(1))%32)^v1610) != 0 {
		v1716 = v1533
		v1718 = v1598
		v1721 = v1484
		v1722 = v1549
		v1724 = v1517
		goto L209
	} else {
		goto L246
	}
L246:
	;
	v1620 = v1533
	v1622 = v1598
	v1625 = v1484
	v1626 = v1549
	v1628 = v1517
	goto L210
L247:
	;
	v1703 = v1688 & int32(255)
	v1796 = v1684
	v1798 = v1686
	v1802 = v1625
	v1813 = (v1701|base.B2i32(v1703 != int32(5))|v1686)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1703)) | base.B2i32(v1686 == v1693)
	goto L208
L248:
	;
	v1684 = v1636
	v1686 = v1620
	v1688 = v1622
	v1693 = v1628
	v1701 = int32(0)
	goto L247
L249:
	;
	goto L250
L250:
	;
	v1646 = v1636
	v1647 = v1620
	v1649 = v1622
	v1651 = v1638
	v1653 = v1640
	v1654 = int32(1)
	goto L251
L251:
	;
	v1664 = v1646 + int32(1)
	v1665 = int32(10)
	v1666 = base.I32_div_u_s(v1647, v1665)
	v1669 = v1647 - v1666*v1665
	v1674 = v1654 & base.B2i32(v1649&int32(255) == int32(0))
	v1676 = base.I32_div_u_s(v1651, v1665)
	v1678 = base.I32_div_u_s(v1653, v1665)
	if base.Ui32(v1678) < base.Ui32(v1676) {
		v1646 = v1664
		v1647 = v1666
		v1649 = v1669
		v1651 = v1676
		v1653 = v1678
		v1654 = v1674
		goto L251
	} else {
		goto L253
	}
L252:
	;
	v1684 = v1664
	v1686 = v1666
	v1688 = v1669
	v1693 = v1653
	v1701 = v1674 ^ int32(1)
	goto L247
L253:
	;
	goto L252
L254:
	;
	v1740 = v1732
	v1741 = v1716
	v1742 = v1734
	v1744 = v1736
	goto L257
L255:
	;
	v1771 = v1732
	v1772 = v1716
	v1774 = v1718
	v1780 = v1724
	goto L256
L256:
	;
	v1796 = v1771
	v1798 = v1772
	v1802 = v1721
	v1813 = base.B2i32(v1772 == v1780) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1774&int32(255)))
	goto L208
L257:
	;
	v1758 = v1740 + int32(1)
	v1759 = int32(10)
	v1760 = base.I32_div_u_s(v1741, v1759)
	v1762 = base.I32_div_u_s(v1742, v1759)
	v1764 = base.I32_div_u_s(v1744, v1759)
	if base.Ui32(v1764) < base.Ui32(v1762) {
		v1740 = v1758
		v1741 = v1760
		v1742 = v1762
		v1744 = v1764
		goto L257
	} else {
		goto L259
	}
L258:
	;
	v1771 = v1758
	v1772 = v1760
	v1774 = v1741 - v1760*int32(10)
	v1780 = v1744
	goto L256
L259:
	;
	goto L258
L260:
	;
	v1863 = v1815
	v1873 = v1814
	v1880 = int32(9)
	goto L200
L261:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v1821) {
		v1863 = v1821
		v1873 = v1831
		v1880 = int32(7)
		goto L200
	} else {
		goto L262
	}
L262:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v1821) {
		v1863 = v1821
		v1873 = v1831
		v1880 = int32(6)
		goto L200
	} else {
		goto L263
	}
L263:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v1821) {
		v1863 = v1821
		v1873 = v1831
		v1880 = int32(5)
		goto L200
	} else {
		goto L264
	}
L264:
	;
	if base.Ui32(int32(999)) < base.Ui32(v1821) {
		v1863 = v1821
		v1873 = v1831
		v1880 = int32(4)
		goto L200
	} else {
		goto L265
	}
L265:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1821) {
		v1863 = v1821
		v1873 = v1831
		v1880 = int32(3)
		goto L200
	} else {
		goto L266
	}
L266:
	;
	if base.Ui32(int32(9)) < base.Ui32(v1821) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1860 = int32(2)
	goto L269
L268:
	;
	v1860 = int32(1)
	goto L269
L269:
	;
	v1863 = v1821
	v1873 = v1831
	v1880 = v1860
	goto L200
L270:
	;
	v1885 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1169))) = uint8(v1885)
	v1888 = int32(1)
	goto L272
L271:
	;
	v1888 = v1882
	goto L272
L272:
	;
	if base.Ui32(v1881+int32(3)) <= base.Ui32(int32(9)) {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	v2109 = int32(0)
	if base.Ui32(v1863) < base.Ui32(int32(10000)) {
		goto L315
	} else {
		goto L316
	}
L274:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1893))) = v2105
	v2107 = v2104
	goto L273
L275:
	;
	v1893 = v1169 + v1888
	v1894 = int32(0)
	if v1881 <= v1894 {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	goto L277
L277:
	;
	if v1873 != 0 {
		goto L283
	} else {
		goto L284
	}
L278:
	;
	v2104 = int32(2) - v1881
	v2105 = int64(3472328296227679792)
	goto L274
L279:
	;
	goto L280
L280:
	;
	if int32(0) <= v1873 {
		v2104 = v1894
		v2105 = int64(3472328296227680304)
		goto L274
	} else {
		goto L281
	}
L281:
	;
	v2107 = int32(1)
	goto L273
L282:
	;
	v1953 = int32(0)
	if base.Ui32(v1937) < base.Ui32(int32(10000)) {
		goto L291
	} else {
		goto L292
	}
L283:
	;
	v1937 = v1863
	v1943 = v1880
	goto L282
L284:
	;
	goto L285
L285:
	;
	v1906 = v1863
	v1911 = v1880
	goto L286
L286:
	;
	if v1906&int32(1) != 0 {
		v1937 = v1906
		v1943 = v1911
		goto L282
	} else {
		goto L288
	}
L287:
	;
	v1937 = v1906
	v1943 = v1911
	goto L282
L288:
	;
	v1930 = base.I32_div_u_s(v1906, int32(10))
	if int32(0)-v1906 == v1930*int32(-10) {
		v1906 = v1930
		v1911 = v1911 - int32(1)
		goto L286
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	if base.Ui32(v2012) < base.Ui32(int32(100)) {
		goto L298
	} else {
		goto L299
	}
L291:
	;
	v2010 = v1953
	v2012 = v1937
	goto L290
L292:
	;
	goto L293
L293:
	;
	v1960 = v1953
	v1961 = v1937
	goto L294
L294:
	;
	v1977 = v1169 + v1888 + v1943 - v1960
	v1981 = base.I32_div_u_s(v1961, int32(10000))
	v1984 = v1981*int32(-10000) + v1961
	v1985 = int32(100)
	v1986 = base.I32_div_u_s(v1984, v1985)
	v1987 = int32(1)
	v1991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1986<<(uint(v1987)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1977-int32(3)))) = uint16(v1991)
	v2002 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1984-v1986*v1985)<<(uint(v1987)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1977-v1987))) = uint16(v2002)
	v2005 = v1960 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v1961) {
		v1960 = v2005
		v1961 = v1981
		goto L294
	} else {
		goto L296
	}
L295:
	;
	v2010 = v2005
	v2012 = v1981
	goto L290
L296:
	;
	goto L295
L297:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2052) {
		goto L302
	} else {
		goto L303
	}
L298:
	;
	v2051 = v2010
	v2052 = v2012
	goto L297
L299:
	;
	goto L300
L300:
	;
	v2034 = int32(65535)
	v2036 = int32(100)
	v2037 = base.I32_div_u_s(v2012&v2034, v2036)
	v2047 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2012-v2037*v2036)&v2034<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1169+v1888+v1943+(v2010^int32(-1))))) = uint16(v2047)
	v2051 = v2010 | int32(2)
	v2052 = v2037
	goto L297
L301:
	;
	v2071 = int32(1)
	v2072 = v1881 - v2071
	v2073 = v1169 + v1888
	*(*uint8)(unsafe.Add(mBase, uint32(v2073))) = uint8(v2070)
	if base.Ui32(int32(2)) <= base.Ui32(v1943) {
		goto L305
	} else {
		goto L306
	}
L302:
	;
	v2059 = v2052 << (uint(int32(1)) % 32)
	v2062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2059)+uint32(_consts[1618]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1169+(v1888+v1943-v2051)))) = uint8(v2062)
	v2066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2059)+uint32(_consts[1617]))))
	v2070 = v2066
	goto L301
L303:
	;
	goto L304
L304:
	;
	v2070 = v2052 | int32(48)
	goto L301
L305:
	;
	v2078 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2073)+1)) = uint8(v2078)
	v2082 = v1943 + int32(1)
	goto L307
L306:
	;
	v2082 = v2071
	goto L307
L307:
	;
	v2083 = v2082 + v1888
	v2084 = v1169 + v2083
	v2085 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2084))) = uint8(v2085)
	v2090 = base.B2i32(v2072 < int32(0))
	if v2072 < int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v2091 = int32(45)
	goto L310
L309:
	;
	v2091 = int32(43)
	goto L310
L310:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2084)+1)) = uint8(v2091)
	if v2072 < int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v2095 = int32(1) - v1881
	goto L313
L312:
	;
	v2095 = v2072
	goto L313
L313:
	;
	v2100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2095<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2084)+2)) = uint16(v2100)
	v2289 = v2083 + int32(4)
	goto L178
L314:
	;
	if base.Ui32(v2168) < base.Ui32(int32(100)) {
		goto L322
	} else {
		goto L323
	}
L315:
	;
	v2167 = v2109
	v2168 = v1863
	goto L314
L316:
	;
	goto L317
L317:
	;
	v2116 = v1863
	v2117 = v2109
	goto L318
L318:
	;
	v2133 = v1893 + v2107 + v1880 - v2117
	v2134 = int32(4)
	v2137 = base.I32_div_u_s(v2116, int32(10000))
	v2140 = v2137*int32(-10000) + v2116
	v2141 = int32(100)
	v2142 = base.I32_div_u_s(v2140, v2141)
	v2143 = int32(1)
	v2147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2142<<(uint(v2143)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2133-v2134))) = uint16(v2147)
	v2158 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2140-v2142*v2141)<<(uint(v2143)%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2133-int32(2)))) = uint16(v2158)
	v2161 = v2117 + v2134
	if base.Ui32(int32(99999999)) < base.Ui32(v2116) {
		v2116 = v2137
		v2117 = v2161
		goto L318
	} else {
		goto L320
	}
L319:
	;
	v2167 = v2161
	v2168 = v2137
	goto L314
L320:
	;
	goto L319
L321:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2207) {
		goto L326
	} else {
		goto L327
	}
L322:
	;
	v2207 = v2168
	v2208 = v2167
	goto L321
L323:
	;
	goto L324
L324:
	;
	v2188 = int32(2)
	v2190 = int32(65535)
	v2192 = int32(100)
	v2193 = base.I32_div_u_s(v2168&v2190, v2192)
	v2203 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2168-v2193*v2192)&v2190<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1893+v2107+v1880-v2167-v2188))) = uint16(v2203)
	v2207 = v2193
	v2208 = v2167 | v2188
	goto L321
L325:
	;
	v2226 = int32(1)
	if v2107 == v2226 {
		goto L330
	} else {
		goto L331
	}
L326:
	;
	v2220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2207<<(uint(int32(1))%32))+uint32(_consts[1617]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1893+v2107+v1880-v2208-int32(2)))) = uint16(v2220)
	goto L325
L327:
	;
	goto L328
L328:
	;
	v2224 = v2207 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1893+v2107))) = uint8(v2224)
	goto L325
L329:
	;
	v2289 = v2266 + int32(base.Ui32(v1187)>>(uint(int32(31))%32))
	goto L178
L330:
	;
	if v1881&int32(4) != 0 {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	goto L332
L332:
	;
	if v1873 < int32(0) {
		goto L342
	} else {
		goto L343
	}
L333:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v1893))) = v2231
	v2234 = int32(5)
	goto L335
L334:
	;
	v2234 = v2226
	goto L335
L335:
	;
	if v1881&int32(2) != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2237 = v1893 + v2234
	v2240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2237))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2237-int32(1)))) = uint16(v2240)
	v2245 = v2234 | int32(2)
	goto L338
L337:
	;
	v2245 = v2234
	goto L338
L338:
	;
	if v1881&int32(1) != 0 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v2248 = v1893 + v2245
	v2251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2248))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2248-int32(1)))) = uint8(v2251)
	goto L341
L340:
	;
	goto L341
L341:
	;
	v2255 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1893+v1881))) = uint8(v2255)
	v2266 = v1880 + int32(1)
	goto L329
L342:
	;
	v2263 = int32(2) - v1873
	goto L344
L343:
	;
	v2263 = v1881
	goto L344
L344:
	;
	v2266 = v2263
	goto L329
L345:
	;
	goto L177
L346:
	;
	F_pfree(m, v10)
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L1
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	return v20
L349:
	;
	goto L348
}
func F_vector_recv(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v60 float32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pq_getmsgint(m, v13, int32(2))
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
	v20 = F_pq_getmsgint(m, v13, int32(2))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = base.I32_extend16_s(v15)
	F_CheckDim_3(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.B2i32(v12 != int32(-1))&base.B2i32(v12 != v22) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L25
	}
L6:
	;
	v32 = v20 << (uint(int32(16)) % 32)
	if v32 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L9:
	;
	v35 = F_mul_size(m, int32(4), v22)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v37 = F_add_size(m, int32(8), v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = F_palloc0(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)) = uint16(v15)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v37 << (uint(int32(2)) % 32)
	if int32(0) < v22 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	m.G0 = v10 + int32(32)
	return v39
L16:
	;
	v60 = F_pq_getmsgfloat4(m, v13)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v39+int32(8)+v50<<(uint(int32(2))%32)))) = v60
	F_CheckElement_3(m, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v66 = v50 + int32(1)
	if v66 != v22 {
		v50 = v66
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	F_errmsg(m, int32(490583), v10+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(520291), int32(88), int32(303246))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32 >> (uint(int32(16)) % 32)
	F_errmsg(m, int32(490614), v10)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(520291), int32(393), int32(37909))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_sub(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
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
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 float32
	_ = v67
	var v69 float32
	_ = v69
	var v73 int32
	_ = v73
	var v76 float32
	_ = v76
	var v78 float32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 float32
	_ = v102
	var v104 float32
	_ = v104
	var v110 int32
	_ = v110
	var v122 float32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v11 = m.G0
	v13 = v11 - int32(16)
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	if v23 == v24 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v13 + int32(16)
	return v33
L5:
	;
	v29 = F_mul_size(m, int32(4), base.I32_extend16_s(v23))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L28
	}
L8:
	;
	v31 = F_add_size(m, int32(8), v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v33 = F_palloc0(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)) = uint16(v23)
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v31 << (uint(int32(2)) % 32)
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	if v39 <= int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v42 = int32(8)
	v43 = v16 + v42
	v45 = v21 + v42
	v47 = v33 + v42
	v48 = int32(0)
	if v39 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v55 = v48
	v61 = int32(0)
	goto L15
L13:
	;
	v88 = v48
	goto L14
L14:
	;
	if v39&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v63 = int32(2)
	v64 = v55 << (uint(v63) % 32)
	v67 = *(*float32)(unsafe.Add(mBase, uint32(v64+v43)))
	v69 = *(*float32)(unsafe.Add(mBase, uint32(v64+v45)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v64))) = base.F32_sub(v67, v69)
	v73 = v64 | int32(4)
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v73+v43)))
	v78 = *(*float32)(unsafe.Add(mBase, uint32(v73+v45)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v73))) = base.F32_sub(v76, v78)
	v82 = v55 + v63
	v84 = v61 + v63
	if v84 != v39&int32(32766) {
		v55 = v82
		v61 = v84
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v88 = v82
	goto L14
L17:
	;
	goto L16
L18:
	;
	v99 = v88 << (uint(int32(2)) % 32)
	v102 = *(*float32)(unsafe.Add(mBase, uint32(v99+v43)))
	v104 = *(*float32)(unsafe.Add(mBase, uint32(v99+v45)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v99))) = base.F32_sub(v102, v104)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v110 = int32(0)
	goto L21
L21:
	;
	v122 = *(*float32)(unsafe.Add(mBase, uint32(v47+v110<<(uint(int32(2))%32))))
	if base.F32_ne(base.F32_abs(v122), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L27
	}
L23:
	;
	v127 = v110 + int32(1)
	if v39 != v127 {
		v110 = v127
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L4
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v138
	F_errmsg(m, int32(501556), v13)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(520291), int32(76), int32(160460))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_to_sparsevec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 float32
	_ = v60
	var v61 float32
	_ = v61
	var v64 float32
	_ = v64
	var v68 float32
	_ = v68
	var v72 float32
	_ = v72
	var v75 int32
	_ = v75
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v112 float32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v181 float32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
	F_CheckDim_2(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.B2i32(v23 != int32(-1))&base.B2i32(v23 != v24) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = base.B2i32(v24 <= int32(0))
	if v24 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L41
	}
L7:
	;
	F_CheckNnz(m, v122, v24)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L21
	}
L8:
	;
	v122 = v2
	goto L7
L9:
	;
	goto L10
L10:
	;
	v36 = v24 & int32(3)
	v38 = v19 + int32(8)
	v39 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v24) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = v39
	v45 = v2
	v52 = v2
	goto L14
L12:
	;
	v81 = v39
	v82 = v2
	goto L13
L13:
	;
	if v36 == int32(0) {
		v122 = v82
		goto L7
	} else {
		goto L17
	}
L14:
	;
	v59 = v38 + v44<<(uint(int32(2))%32)
	v60 = *(*float32)(unsafe.Add(mBase, uint32(v59)))
	v61 = float32(0)
	v64 = *(*float32)(unsafe.Add(mBase, uint32(v59)+4))
	v68 = *(*float32)(unsafe.Add(mBase, uint32(v59)+8))
	v72 = *(*float32)(unsafe.Add(mBase, uint32(v59)+12))
	v75 = v45 + base.F32_ne(v60, v61) + base.F32_ne(v64, v61) + base.F32_ne(v68, v61) + base.F32_ne(v72, v61)
	v76 = int32(4)
	v77 = v44 + v76
	v79 = v52 + v76
	if v79 != v24&int32(32764) {
		v44 = v77
		v45 = v75
		v52 = v79
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v81 = v77
	v82 = v75
	goto L13
L16:
	;
	goto L15
L17:
	;
	v96 = v81
	v97 = v82
	v103 = v2
	goto L18
L18:
	;
	v112 = *(*float32)(unsafe.Add(mBase, uint32(v38+v96<<(uint(int32(2))%32))))
	v115 = v97 + base.F32_ne(v112, float32(0))
	v116 = int32(1)
	v119 = v103 + v116
	if v119 != v36 {
		v96 = v96 + v116
		v97 = v115
		v103 = v119
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v122 = v115
	goto L7
L20:
	;
	goto L19
L21:
	;
	v138 = F_mul_size(m, int32(4), v122)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v140 = F_add_size(m, int32(16), v138)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v143 = F_mul_size(m, int32(4), v122)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v145 = F_add_size(m, v140, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v147 = F_palloc0(m, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v145 << (uint(int32(2)) % 32)
	if v34 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L38
	}
L28:
	;
	v157 = v147 + int32(16)
	v163 = int32(0)
	v165 = v163
	v166 = v163
	goto L31
L29:
	;
	goto L30
L30:
	;
	m.G0 = v16 + int32(16)
	return v147
L31:
	;
	v181 = *(*float32)(unsafe.Add(mBase, uint32(v19+int32(8)+v165<<(uint(int32(2))%32))))
	if base.F32_ne(v181, float32(0)) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	if v184 <= v166 {
		goto L27
	} else {
		goto L36
	}
L34:
	;
	v194 = v166
	goto L35
L35:
	;
	v197 = v165 + int32(1)
	if v197 != v24 {
		v165 = v197
		v166 = v194
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v187 = v166 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v157+v187))) = v165
	*(*float32)(unsafe.Add(mBase, uint32(v187+(v157+v122<<(uint(int32(2))%32))))) = v181
	v194 = v166 + int32(1)
	goto L35
L37:
	;
	goto L32
L38:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(526480), int32(632), int32(515799))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v23
	F_errmsg(m, int32(490583), v16)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(526480), int32(62), int32(303246))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
