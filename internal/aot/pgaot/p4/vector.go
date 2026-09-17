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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v6 == int32(1) {
		v10 = int32(18)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
		if v12 == v10 {
			v15 = v10
		} else {
			v15 = int32(2)
		}
		if base.Ui32((v12-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v22 = int32(6)
		} else {
			v22 = v15
		}
		v31 = v22
	} else {
		v23 = int32(1)
		if v6&v23 != 0 {
			v31 = int32(base.Ui32(v6) >> (uint(v23) % 32))
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v31 = int32(base.Ui32(v27) >> (uint(int32(2)) % 32))
		}
	}
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v31) <= base.Ui32(v32) {
		if l1 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_VectorArraySet_0), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_VectorArraySet_1), int32(326), int32(_a_F_VectorArraySet_2))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v36 <= l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_VectorArraySet_0), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_VectorArraySet_1), int32(326), int32(_a_F_VectorArraySet_2))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if v31 != 0 {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					base.MemoryCopy(m, v38+l1*v32, l2, v31)
				} else {
				}
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_VectorArraySet_0), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_VectorArraySet_1), int32(337), int32(_a_F_VectorArraySet_3))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
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
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v74 float32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
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
	return v84
L9:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	v84 = int32(0)
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
		v84 = v43
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v56 = int32(0)
	goto L17
L17:
	;
	v68 = v56 + int32(1)
	v72 = *(*float64)(unsafe.Add(mBase, uint32(v14+int32(24)+v68<<(uint(int32(3))%32))))
	v74 = base.F32_demote_f64(base.F64_div(v72, v28))
	*(*float32)(unsafe.Add(mBase, uint32(v43+int32(8)+v56<<(uint(int32(2))%32)))) = v74
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
	v84 = v43
	goto L8
L19:
	;
	if v68 != v34 {
		v56 = v68
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_vector_avg_0)
	F_errmsg_internal(m, int32(_a_F_vector_avg_1), v11)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_vector_avg_2), int32(169), int32(_a_F_vector_avg_3))
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 float32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 float32
	_ = v70
	var v73 int32
	_ = v73
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v82 float32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 float32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v171 float32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 float32
	_ = v177
	var v180 int32
	_ = v180
	var v183 float32
	_ = v183
	var v186 int32
	_ = v186
	var v189 float32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 float32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
			v22 = v20 + v21
			F_CheckDim_3(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v27 = F_mul_size(m, int32(4), v22)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = F_add_size(m, int32(8), v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = F_palloc0(m, v29)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)) = uint16(v22)
							*(*int32)(unsafe.Add(mBase, uint32(v31))) = v29 << (uint(int32(2)) % 32)
							v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
							if v37 <= int32(0) {
							} else {
								v40 = int32(8)
								v41 = v31 + v40
								v43 = v13 + v40
								v44 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v37) {
									v49 = v44
									v54 = v2
									for {
										v61 = v49 << (uint(int32(2)) % 32)
										v64 = *(*float32)(unsafe.Add(mBase, uint32(v61+v43)))
										*(*float32)(unsafe.Add(mBase, uint32(v41+v61))) = v64
										v66 = int32(4)
										v67 = v61 | v66
										v70 = *(*float32)(unsafe.Add(mBase, uint32(v43+v67)))
										*(*float32)(unsafe.Add(mBase, uint32(v41+v67))) = v70
										v73 = v61 | int32(8)
										v76 = *(*float32)(unsafe.Add(mBase, uint32(v43+v73)))
										*(*float32)(unsafe.Add(mBase, uint32(v41+v73))) = v76
										v79 = v61 | int32(12)
										v82 = *(*float32)(unsafe.Add(mBase, uint32(v79+v43)))
										*(*float32)(unsafe.Add(mBase, uint32(v41+v79))) = v82
										v85 = v49 + v66
										v87 = v54 + v66
										if v87 != v37&int32(_a_F_vector_concat_0) {
											v49 = v85
											v54 = v87
											continue
										} else {
											break
										}
										break
									}
									if v37&int32(3) == int32(0) {
									} else {
										v93 = v85
										v106 = v93
										v115 = v2
										for {
											v118 = v106 << (uint(int32(2)) % 32)
											v121 = *(*float32)(unsafe.Add(mBase, uint32(v43+v118)))
											*(*float32)(unsafe.Add(mBase, uint32(v41+v118))) = v121
											v123 = int32(1)
											v126 = v115 + v123
											if v126 != v37&int32(3) {
												v106 = v106 + v123
												v115 = v126
												continue
											} else {
												break
											}
											break
										}
									}
								} else {
									v93 = v44
									v106 = v93
									v115 = v2
									for {
										v118 = v106 << (uint(int32(2)) % 32)
										v121 = *(*float32)(unsafe.Add(mBase, uint32(v43+v118)))
										*(*float32)(unsafe.Add(mBase, uint32(v41+v118))) = v121
										v123 = int32(1)
										v126 = v115 + v123
										if v126 != v37&int32(3) {
											v106 = v106 + v123
											v115 = v126
											continue
										} else {
											break
										}
										break
									}
								}
							}
							v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
							if v139 <= int32(0) {
							} else {
								v142 = int32(8)
								v143 = v18 + v142
								v148 = v31 + v37<<(uint(int32(2))%32) + v142
								v149 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v139) {
									v156 = v149
									v161 = int32(0)
									for {
										v168 = v156 << (uint(int32(2)) % 32)
										v171 = *(*float32)(unsafe.Add(mBase, uint32(v168+v143)))
										*(*float32)(unsafe.Add(mBase, uint32(v148+v168))) = v171
										v173 = int32(4)
										v174 = v168 | v173
										v177 = *(*float32)(unsafe.Add(mBase, uint32(v143+v174)))
										*(*float32)(unsafe.Add(mBase, uint32(v148+v174))) = v177
										v180 = v168 | int32(8)
										v183 = *(*float32)(unsafe.Add(mBase, uint32(v143+v180)))
										*(*float32)(unsafe.Add(mBase, uint32(v148+v180))) = v183
										v186 = v168 | int32(12)
										v189 = *(*float32)(unsafe.Add(mBase, uint32(v186+v143)))
										*(*float32)(unsafe.Add(mBase, uint32(v148+v186))) = v189
										v192 = v156 + v173
										v194 = v161 + v173
										if v194 != v139&int32(_a_F_vector_concat_0) {
											v156 = v192
											v161 = v194
											continue
										} else {
											break
										}
										break
									}
									if v139&int32(3) == int32(0) {
									} else {
										v200 = v192
										v213 = v200
										v222 = v149
										for {
											v225 = v213 << (uint(int32(2)) % 32)
											v228 = *(*float32)(unsafe.Add(mBase, uint32(v225+v143)))
											*(*float32)(unsafe.Add(mBase, uint32(v148+v225))) = v228
											v230 = int32(1)
											v233 = v222 + v230
											if v233 != v139&int32(3) {
												v213 = v213 + v230
												v222 = v233
												continue
											} else {
												break
											}
											break
										}
									}
								} else {
									v200 = v149
									v213 = v200
									v222 = v149
									for {
										v225 = v213 << (uint(int32(2)) % 32)
										v228 = *(*float32)(unsafe.Add(mBase, uint32(v225+v143)))
										*(*float32)(unsafe.Add(mBase, uint32(v148+v225))) = v228
										v230 = int32(1)
										v233 = v222 + v230
										if v233 != v139&int32(3) {
											v213 = v213 + v230
											v222 = v233
											continue
										} else {
											break
										}
										break
									}
								}
							}
							return v31
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
	var v49 int32
	_ = v49
	var v51 float32
	_ = v51
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
	var v79 int32
	_ = v79
	var v89 float32
	_ = v89
	var v91 int32
	_ = v91
	var v93 float32
	_ = v93
	var v95 float32
	_ = v95
	var v96 float32
	_ = v96
	var v109 float32
	_ = v109
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
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
					v122 = float64(0)
				} else {
					v31 = int32(8)
					v32 = v22 + v31
					v34 = v17 + v31
					if v27 == int32(1) {
						v79 = int32(0)
						v89 = v10
						v91 = v79 << (uint(int32(2)) % 32)
						v93 = *(*float32)(unsafe.Add(mBase, uint32(v34+v91)))
						v95 = *(*float32)(unsafe.Add(mBase, uint32(v91+v32)))
						v96 = base.F32_sub(v93, v95)
						v109 = base.F32_add(base.F32_mul(v96, v96), v89)
					} else {
						v41 = int32(0)
						v49 = int32(0)
						v51 = v10
						for {
							v52 = int32(2)
							v53 = v41 << (uint(v52) % 32)
							v55 = v53 | int32(4)
							v57 = *(*float32)(unsafe.Add(mBase, uint32(v34+v55)))
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v32+v55)))
							v60 = base.F32_sub(v57, v59)
							v63 = *(*float32)(unsafe.Add(mBase, uint32(v34+v53)))
							v65 = *(*float32)(unsafe.Add(mBase, uint32(v32+v53)))
							v66 = base.F32_sub(v63, v65)
							v69 = base.F32_add(base.F32_mul(v60, v60), base.F32_add(base.F32_mul(v66, v66), v51))
							v71 = v41 + v52
							v73 = v49 + v52
							if v73 != v27&int32(_a_F_vector_l2_squared_distance_0) {
								v41 = v71
								v49 = v73
								v51 = v69
								continue
							} else {
								break
							}
							break
						}
						if v27&int32(1) == int32(0) {
							v109 = v69
						} else {
							v79 = v71
							v89 = v69
							v91 = v79 << (uint(int32(2)) % 32)
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v34+v91)))
							v95 = *(*float32)(unsafe.Add(mBase, uint32(v91+v32)))
							v96 = base.F32_sub(v93, v95)
							v109 = base.F32_add(base.F32_mul(v96, v96), v89)
						}
					}
					v122 = base.F64_promote_f32(v109)
				}
				v123 = F_Float8GetDatum(m, v122)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(16)
					return v123
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int32(0)
					} else {
						v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v137
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v136
						F_errmsg(m, int32(_a_F_vector_l2_squared_distance_1), v14)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_vector_l2_squared_distance_2), int32(76), int32(_a_F_vector_l2_squared_distance_3))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 float32
	_ = v41
	var v42 float64
	_ = v42
	var v44 float32
	_ = v44
	var v45 float64
	_ = v45
	var v47 float32
	_ = v47
	var v48 float64
	_ = v48
	var v50 float32
	_ = v50
	var v51 float64
	_ = v51
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 float64
	_ = v68
	var v77 int32
	_ = v77
	var v79 float64
	_ = v79
	var v85 int32
	_ = v85
	var v89 float32
	_ = v89
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 float64
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
			v19 = F_Float8GetDatum(m, float64(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v19
			}
		} else {
			v23 = v11 + int32(8)
			v24 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v15) {
				v29 = v24
				v31 = v2
				v36 = v4
				for {
					v40 = v23 + v29<<(uint(int32(2))%32)
					v41 = *(*float32)(unsafe.Add(mBase, uint32(v40)+12))
					v42 = base.F64_promote_f32(v41)
					v44 = *(*float32)(unsafe.Add(mBase, uint32(v40)+8))
					v45 = base.F64_promote_f32(v44)
					v47 = *(*float32)(unsafe.Add(mBase, uint32(v40)+4))
					v48 = base.F64_promote_f32(v47)
					v50 = *(*float32)(unsafe.Add(mBase, uint32(v40)))
					v51 = base.F64_promote_f32(v50)
					v56 = base.F64_add(base.F64_mul(v42, v42), base.F64_add(base.F64_mul(v45, v45), base.F64_add(base.F64_mul(v48, v48), base.F64_add(base.F64_mul(v51, v51), v31))))
					v57 = int32(4)
					v58 = v29 + v57
					v60 = v36 + v57
					if v60 != v15&int32(_a_F_vector_norm_0) {
						v29 = v58
						v31 = v56
						v36 = v60
						continue
					} else {
						break
					}
					break
				}
				if v15&int32(3) == int32(0) {
					v100 = v56
				} else {
					v66 = v58
					v68 = v56
					v77 = v66
					v79 = v68
					v85 = v4
					for {
						v89 = *(*float32)(unsafe.Add(mBase, uint32(v23+v77<<(uint(int32(2))%32))))
						v90 = base.F64_promote_f32(v89)
						v92 = base.F64_add(base.F64_mul(v90, v90), v79)
						v93 = int32(1)
						v96 = v85 + v93
						if v96 != v15&int32(3) {
							v77 = v77 + v93
							v79 = v92
							v85 = v96
							continue
						} else {
							break
						}
						break
					}
					v100 = v92
				}
			} else {
				v66 = v24
				v68 = v2
				v77 = v66
				v79 = v68
				v85 = v4
				for {
					v89 = *(*float32)(unsafe.Add(mBase, uint32(v23+v77<<(uint(int32(2))%32))))
					v90 = base.F64_promote_f32(v89)
					v92 = base.F64_add(base.F64_mul(v90, v90), v79)
					v93 = int32(1)
					v96 = v85 + v93
					if v96 != v15&int32(3) {
						v77 = v77 + v93
						v79 = v92
						v85 = v96
						continue
					} else {
						break
					}
					break
				}
				v100 = v92
			}
			v108 = F_Float8GetDatum(m, base.F64_sqrt(v100))
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return int32(0)
			} else {
				return v108
			}
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
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int64
	_ = v207
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v351 int32
	_ = v351
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int64
	_ = v374
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int64
	_ = v390
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int64
	_ = v417
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v426 int64
	_ = v426
	var v428 int32
	_ = v428
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int64
	_ = v950
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1144 float32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1248 int64
	_ = v1248
	var v1250 int64
	_ = v1250
	var v1251 int64
	_ = v1251
	var v1253 int64
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int64
	_ = v1257
	var v1258 int64
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1279 int64
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1284 int64
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1294 int32
	_ = v1294
	var v1295 int64
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1300 int64
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1324 int64
	_ = v1324
	var v1328 int64
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1333 int64
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1346 int32
	_ = v1346
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1434 int32
	_ = v1434
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1461 int64
	_ = v1461
	var v1463 int64
	_ = v1463
	var v1464 int64
	_ = v1464
	var v1466 int64
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int64
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int64
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1496 int64
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1506 int32
	_ = v1506
	var v1507 int64
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1512 int64
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1534 int64
	_ = v1534
	var v1538 int64
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1543 int64
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1558 int32
	_ = v1558
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1773 int32
	_ = v1773
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1841 int32
	_ = v1841
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1902 int32
	_ = v1902
	var v1908 int32
	_ = v1908
	var v1913 int32
	_ = v1913
	var v1925 int32
	_ = v1925
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1970 int32
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2013 int32
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2067 int64
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2116 int32
	_ = v2116
	var v2119 int32
	_ = v2119
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2180 int32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2255 int32
	_ = v2255
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
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
		v2249 = v26
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v2255 = int32(93)
	*(*uint16)(unsafe.Add(mBase, uint32(v2249))) = uint16(v2255)
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2257 != v10 {
		goto L338
	} else {
		goto L339
	}
L7:
	;
	v29 = *(*float32)(unsafe.Add(mBase, uint32(v10)+8))
	v30 = int32(0)
	v47 = base.I32_reinterpret_f32(v29)
	v49 = v47 & int32(_a_F_vector_out_0)
	v52 = int32(255)
	v53 = int32(base.Ui32(v47)>>(uint(int32(23))%32)) & v52
	if v53|v49 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v1127 = v1126 + v26
	if v14 == int32(1) {
		v2249 = v1127
		goto L6
	} else {
		goto L171
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
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vector_out[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2)) = uint8(v62)
	v65 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_vector_out[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v65)
	v1126 = int32(3)
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
	v1126 = v81
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
	v1126 = v88
	goto L8
L30:
	;
	v736 = int32(0)
	if v47 < v736 {
		goto L98
	} else {
		goto L99
	}
L31:
	;
	if base.Ui32(int32(_a_F_vector_out_1)) < base.Ui32(v677) {
		v719 = v677
		v724 = v682
		v735 = int32(8)
		goto L30
	} else {
		goto L89
	}
L32:
	;
	v105 = v49 << (uint(int32(2)) % 32)
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
	v677 = int32(base.Ui32(v49|int32(_a_F_vector_out_2)) >> (uint(v95) % 32))
	v682 = v30
	goto L31
L35:
	;
	v108 = v105 | int32(33554432)
	goto L37
L36:
	;
	v108 = v105
	goto L37
L37:
	;
	v111 = int32(2)
	v116 = v108 + (base.B2i32(v49 != int32(0)) | base.B2i32(base.Ui32(v53) < base.Ui32(v111)) ^ int32(-1))
	v118 = v108 | v111
	if v53 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v669 = v652 + v656
	v670 = v651 + v668
	if base.Ui32(v670) <= base.Ui32(int32(99999999)) {
		v677 = v670
		v682 = v669
		goto L31
	} else {
		goto L88
	}
L39:
	;
	v587 = int32(0)
	v588 = int32(10)
	v589 = base.I32_div_u_s(v573, v588)
	v591 = base.I32_div_u_s(v577, v588)
	if base.Ui32(v591) < base.Ui32(v589) {
		goto L82
	} else {
		goto L83
	}
L40:
	;
	v491 = int32(0)
	v492 = int32(10)
	v493 = base.I32_div_u_s(v477, v492)
	v495 = base.I32_div_u_s(v481, v492)
	if base.Ui32(v493) <= base.Ui32(v495) {
		goto L76
	} else {
		goto L77
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
	v128 = int32(base.Ui32(v122*int32(_a_F_vector_out_3)) >> (uint(int32(18)) % 32))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v128<<(uint(int32(3))%32))+uint32(_c_F_vector_out[2])))
	v133 = v131 & int64(4294967295)
	v134 = base.I64_extend_i32_u(v116)
	v136 = int64(32)
	v138 = base.I32_wrap_i64(int64(base.Ui64(v133*v134) >> (uint(v136) % 64)))
	v140 = int64(base.Ui64(v131) >> (uint(v136) % 64))
	v141 = v134 * v140
	v143 = v138 + base.I32_wrap_i64(v141)
	v150 = v128 - v122
	v155 = v150 + int32(base.Ui32(v128*int32(_a_F_vector_out_4))>>(uint(int32(19))%32))
	v156 = int32(5) - v155
	v159 = v155 + int32(27)
	v161 = (base.B2i32(base.Ui32(v143) < base.Ui32(v138))+base.I32_wrap_i64(int64(base.Ui64(v141)>>(uint(v136)%64))))<<(uint(v156)%32) | int32(base.Ui32(v143)>>(uint(v159)%32))
	v162 = base.I64_extend_i32_u(v118)
	v166 = base.I32_wrap_i64(int64(base.Ui64(v133*v162) >> (uint(v136) % 64)))
	v167 = v162 * v140
	v169 = v166 + base.I32_wrap_i64(v167)
	v177 = (base.B2i32(base.Ui32(v169) < base.Ui32(v166))+base.I32_wrap_i64(int64(base.Ui64(v167)>>(uint(v136)%64))))<<(uint(v156)%32) | int32(base.Ui32(v169)>>(uint(v159)%32))
	v178 = base.I64_extend_i32_u(v108)
	v182 = base.I32_wrap_i64(int64(base.Ui64(v133*v178) >> (uint(v136) % 64)))
	v183 = v178 * v140
	v185 = v182 + base.I32_wrap_i64(v183)
	v193 = (base.B2i32(base.Ui32(v185) < base.Ui32(v182))+base.I32_wrap_i64(int64(base.Ui64(v183)>>(uint(v136)%64))))<<(uint(v156)%32) | int32(base.Ui32(v185)>>(uint(v159)%32))
	v194 = int32(0)
	if v128 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v337 = v122 * int32(-732923)
	v339 = int32(base.Ui32(v337) >> (uint(int32(20)) % 32))
	v340 = v122 + v339
	v344 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_vector_out_5)-v340<<(uint(int32(3))%32))))
	v346 = v344 & int64(4294967295)
	v347 = base.I64_extend_i32_u(v116)
	v349 = int64(32)
	v351 = base.I32_wrap_i64(int64(base.Ui64(v346*v347) >> (uint(v349) % 64)))
	v353 = int64(base.Ui64(v344) >> (uint(v349) % 64))
	v354 = v347 * v353
	v356 = v351 + base.I32_wrap_i64(v354)
	v367 = v339 - int32(base.Ui32(v340*int32(-1217359))>>(uint(int32(19))%32))
	v368 = int32(4) - v367
	v371 = v367 + int32(28)
	v373 = (base.B2i32(base.Ui32(v356) < base.Ui32(v351))+base.I32_wrap_i64(int64(base.Ui64(v354)>>(uint(v349)%64))))<<(uint(v368)%32) | int32(base.Ui32(v356)>>(uint(v371)%32))
	v374 = base.I64_extend_i32_u(v108)
	v378 = base.I32_wrap_i64(int64(base.Ui64(v346*v374) >> (uint(v349) % 64)))
	v379 = v374 * v353
	v381 = v378 + base.I32_wrap_i64(v379)
	v389 = (base.B2i32(base.Ui32(v381) < base.Ui32(v378))+base.I32_wrap_i64(int64(base.Ui64(v379)>>(uint(v349)%64))))<<(uint(v368)%32) | int32(base.Ui32(v381)>>(uint(v371)%32))
	v390 = base.I64_extend_i32_u(v118)
	v394 = base.I32_wrap_i64(int64(base.Ui64(v346*v390) >> (uint(v349) % 64)))
	v395 = v353 * v390
	v397 = v394 + base.I32_wrap_i64(v395)
	v405 = (base.B2i32(base.Ui32(v397) < base.Ui32(v394))+base.I32_wrap_i64(int64(base.Ui64(v395)>>(uint(v349)%64))))<<(uint(v368)%32) | int32(base.Ui32(v397)>>(uint(v371)%32))
	v407 = v405 - int32(1)
	if v339 != 0 {
		goto L68
	} else {
		goto L69
	}
L47:
	;
	v198 = int32(10)
	v199 = base.I32_div_u_s(v177-int32(1), v198)
	v201 = base.I32_div_u_s(v161, v198)
	if base.Ui32(v199) <= base.Ui32(v201) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v245 = v194
	goto L49
L49:
	;
	v251 = base.I32_rem_u_s(v108, int32(5))
	if v251 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v204 = v128 - int32(1)
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v204<<(uint(int32(3))%32))+uint32(_c_F_vector_out[2])))
	v211 = int64(32)
	v213 = base.I32_wrap_i64(int64(base.Ui64(v207&int64(4294967295)*v178) >> (uint(v211) % 64)))
	v216 = int64(base.Ui64(v207)>>(uint(v211)%64)) * v178
	v218 = v213 + base.I32_wrap_i64(v216)
	v229 = v150 + int32(base.Ui32(v204*int32(_a_F_vector_out_4))>>(uint(int32(19))%32))
	v237 = base.I32_rem_u_s((base.B2i32(base.Ui32(v218) < base.Ui32(v213))+base.I32_wrap_i64(int64(base.Ui64(v216)>>(uint(v211)%64))))<<(uint(int32(6)-v229)%32)|int32(base.Ui32(v218)>>(uint(v229+int32(26))%32)), int32(10))
	v238 = v237
	goto L52
L51:
	;
	v238 = v194
	goto L52
L52:
	;
	if base.Ui32(int32(33)) < base.Ui32(v122) {
		v570 = v193
		v573 = v177
		v575 = v128
		v576 = v238
		v577 = v161
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v245 = v238
	goto L49
L54:
	;
	v257 = v108
	v258 = v194
	goto L57
L55:
	;
	goto L56
L56:
	;
	v282 = int32(0)
	v284 = base.I32_rem_u_s(v118, int32(5))
	if v284 == v282 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v274 = v258 + int32(1)
	v275 = int32(5)
	v276 = base.I32_div_u_s(v257, v275)
	v278 = base.I32_rem_u_s(v276, v275)
	if v278 == int32(0) {
		v257 = v276
		v258 = v274
		goto L57
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(v274) < base.Ui32(v128) {
		v570 = v193
		v573 = v177
		v575 = v128
		v576 = v245
		v577 = v161
		goto L39
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v474 = v193
	v477 = v177
	v479 = v128
	v480 = v245
	v481 = v161
	goto L40
L61:
	;
	v290 = v282
	v293 = v118
	goto L64
L62:
	;
	v317 = v282
	goto L63
L63:
	;
	v570 = v193
	v573 = v177 - base.B2i32(base.Ui32(v128) <= base.Ui32(v317))
	v575 = v128
	v576 = v245
	v577 = v161
	goto L39
L64:
	;
	v307 = v290 + int32(1)
	v308 = int32(5)
	v309 = base.I32_div_u_s(v293, v308)
	v311 = base.I32_rem_u_s(v309, v308)
	if v311 == int32(0) {
		v290 = v307
		v293 = v309
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v317 = v307
	goto L63
L66:
	;
	goto L65
L67:
	;
	v462 = int32(-1)
	if v108&(v462<<(uint(v339-int32(1))%32)^v462)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v337)) != 0 {
		v570 = v389
		v573 = v405
		v575 = v340
		v576 = v451
		v577 = v373
		goto L39
	} else {
		goto L75
	}
L68:
	;
	v408 = int32(10)
	v409 = base.I32_div_u_s(v407, v408)
	v411 = base.I32_div_u_s(v373, v408)
	if base.Ui32(v409) <= base.Ui32(v411) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v458 = v30
	goto L70
L70:
	;
	v474 = v389
	v477 = v407
	v479 = v340
	v480 = v458
	v481 = v373
	goto L40
L71:
	;
	v414 = int32(1) - v340
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v414<<(uint(int32(3))%32))+uint32(_c_F_vector_out[3])))
	v421 = int64(32)
	v423 = base.I32_wrap_i64(int64(base.Ui64(v417&int64(4294967295)*v374) >> (uint(v421) % 64)))
	v426 = int64(base.Ui64(v417)>>(uint(v421)%64)) * v374
	v428 = v423 + base.I32_wrap_i64(v426)
	v441 = v339 + (int32(base.Ui32(v414*int32(_a_F_vector_out_4))>>(uint(int32(19))%32)) ^ int32(-1))
	v449 = base.I32_rem_u_s((base.B2i32(base.Ui32(v428) < base.Ui32(v423))+base.I32_wrap_i64(int64(base.Ui64(v426)>>(uint(v421)%64))))<<(uint(int32(4)-v441)%32)|int32(base.Ui32(v428)>>(uint(v441+int32(28))%32)), int32(10))
	v451 = v449
	goto L73
L72:
	;
	v451 = v30
	goto L73
L73:
	;
	if v339 != int32(1) {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	v458 = v451
	goto L70
L75:
	;
	v474 = v389
	v477 = v405
	v479 = v340
	v480 = v451
	v481 = v373
	goto L40
L76:
	;
	v539 = v474
	v540 = v491
	v545 = v480
	v546 = v481
	v556 = int32(0)
	goto L78
L77:
	;
	v501 = v474
	v502 = v491
	v503 = v493
	v504 = int32(1)
	v505 = v495
	v507 = v480
	goto L79
L78:
	;
	v558 = v545 & int32(255)
	v651 = v539
	v652 = v540
	v656 = v479
	v668 = (v556|base.B2i32(v558 != int32(5))|v539)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v558)) | base.B2i32(v539 == v546)
	goto L38
L79:
	;
	v519 = v502 + int32(1)
	v520 = int32(10)
	v521 = base.I32_div_u_s(v501, v520)
	v524 = v501 - v521*v520
	v529 = v504 & base.B2i32(v507&int32(255) == int32(0))
	v531 = base.I32_div_u_s(v503, v520)
	v533 = base.I32_div_u_s(v505, v520)
	if base.Ui32(v533) < base.Ui32(v531) {
		v501 = v521
		v502 = v519
		v503 = v531
		v504 = v529
		v505 = v533
		v507 = v524
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v539 = v521
	v540 = v519
	v545 = v524
	v546 = v505
	v556 = v529 ^ int32(1)
	goto L78
L81:
	;
	goto L80
L82:
	;
	v595 = v570
	v596 = v587
	v597 = v589
	v599 = v591
	goto L85
L83:
	;
	v626 = v570
	v627 = v587
	v632 = v576
	v633 = v577
	goto L84
L84:
	;
	v651 = v626
	v652 = v627
	v656 = v575
	v668 = base.B2i32(v626 == v633) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v632&int32(255)))
	goto L38
L85:
	;
	v613 = v596 + int32(1)
	v614 = int32(10)
	v615 = base.I32_div_u_s(v595, v614)
	v617 = base.I32_div_u_s(v597, v614)
	v619 = base.I32_div_u_s(v599, v614)
	if base.Ui32(v619) < base.Ui32(v617) {
		v595 = v615
		v596 = v613
		v597 = v617
		v599 = v619
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v626 = v615
	v627 = v613
	v632 = v595 - v615*int32(10)
	v633 = v599
	goto L84
L87:
	;
	goto L86
L88:
	;
	v719 = v670
	v724 = v669
	v735 = int32(9)
	goto L30
L89:
	;
	if base.Ui32(int32(_a_F_vector_out_6)) < base.Ui32(v677) {
		v719 = v677
		v724 = v682
		v735 = int32(7)
		goto L30
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(int32(_a_F_vector_out_7)) < base.Ui32(v677) {
		v719 = v677
		v724 = v682
		v735 = int32(6)
		goto L30
	} else {
		goto L91
	}
L91:
	;
	if base.Ui32(int32(_a_F_vector_out_8)) < base.Ui32(v677) {
		v719 = v677
		v724 = v682
		v735 = int32(5)
		goto L30
	} else {
		goto L92
	}
L92:
	;
	if base.Ui32(int32(999)) < base.Ui32(v677) {
		v719 = v677
		v724 = v682
		v735 = int32(4)
		goto L30
	} else {
		goto L93
	}
L93:
	;
	if base.Ui32(int32(99)) < base.Ui32(v677) {
		v719 = v677
		v724 = v682
		v735 = int32(3)
		goto L30
	} else {
		goto L94
	}
L94:
	;
	if base.Ui32(int32(9)) < base.Ui32(v677) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v715 = int32(2)
	goto L97
L96:
	;
	v715 = int32(1)
	goto L97
L97:
	;
	v719 = v677
	v724 = v682
	v735 = v715
	goto L30
L98:
	;
	v739 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v739)
	v742 = int32(1)
	goto L100
L99:
	;
	v742 = v736
	goto L100
L100:
	;
	v743 = v735 + v724
	if base.Ui32(v743+int32(3)) <= base.Ui32(int32(9)) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v954 = int32(0)
	if base.Ui32(int32(_a_F_vector_out_9)) <= base.Ui32(v719) {
		goto L141
	} else {
		goto L142
	}
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v748))) = v950
	v952 = v949
	goto L101
L103:
	;
	v748 = v26 + v742
	v749 = int32(0)
	if v743 <= v749 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if v724 != 0 {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v949 = int32(2) - v743
	v950 = int64(3472328296227679792)
	goto L102
L107:
	;
	goto L108
L108:
	;
	if int32(0) <= v724 {
		v949 = v749
		v950 = int64(3472328296227680304)
		goto L102
	} else {
		goto L109
	}
L109:
	;
	v952 = int32(1)
	goto L101
L110:
	;
	v808 = int32(0)
	if base.Ui32(int32(_a_F_vector_out_9)) <= base.Ui32(v791) {
		goto L118
	} else {
		goto L119
	}
L111:
	;
	v791 = v719
	v796 = v735
	goto L110
L112:
	;
	goto L113
L113:
	;
	v762 = v719
	v764 = v735
	goto L114
L114:
	;
	if v762&int32(1) != 0 {
		v791 = v762
		v796 = v764
		goto L110
	} else {
		goto L116
	}
L115:
	;
	v791 = v762
	v796 = v764
	goto L110
L116:
	;
	v785 = base.I32_div_u_s(v762, int32(10))
	if int32(0)-v762 == v785*int32(-10) {
		v762 = v785
		v764 = v764 - int32(1)
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v815 = v791
	v816 = v808
	goto L121
L119:
	;
	v861 = v791
	v862 = v808
	goto L120
L120:
	;
	if base.Ui32(v861) < base.Ui32(int32(100)) {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	v832 = v26 + v742 + v796 - v816
	v836 = base.I32_div_u_s(v815, int32(_a_F_vector_out_9))
	v839 = v815 + v836*int32(-10000)
	v840 = int32(100)
	v841 = base.I32_div_u_s(v839, v840)
	v842 = int32(1)
	v844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v841<<(uint(v842)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v832-int32(3)))) = uint16(v844)
	v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v839-v841*v840)<<(uint(v842)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v832-v842))) = uint16(v853)
	v856 = v816 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v815) {
		v815 = v836
		v816 = v856
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v861 = v836
	v862 = v856
	goto L120
L123:
	;
	goto L122
L124:
	;
	v903 = v743 - int32(1)
	v904 = v26 + v742
	if base.Ui32(int32(10)) <= base.Ui32(v901) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v900 = v862
	v901 = v861
	goto L124
L126:
	;
	goto L127
L127:
	;
	v885 = int32(_a_F_vector_out_10)
	v887 = int32(100)
	v888 = base.I32_div_u_s(v861&v885, v887)
	v896 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v861-v888*v887)&v885<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v26+v742+v796+(v862^int32(-1))))) = uint16(v896)
	v900 = v862 | int32(2)
	v901 = v888
	goto L124
L128:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v904))) = uint8(v918)
	if base.Ui32(int32(2)) <= base.Ui32(v796) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v911 = v901 << (uint(int32(1)) % 32)
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911)+uint32(_c_F_vector_out[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26+(v742+v796-v900)))) = uint8(v912)
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911)+uint32(_c_F_vector_out[4]))))
	v918 = v914
	goto L128
L130:
	;
	goto L131
L131:
	;
	v918 = v901 | int32(48)
	goto L128
L132:
	;
	v922 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v904)+1)) = uint8(v922)
	v927 = v796 + int32(1)
	goto L134
L133:
	;
	v927 = int32(1)
	goto L134
L134:
	;
	v928 = v927 + v742
	v929 = v26 + v928
	v930 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v929))) = uint8(v930)
	v935 = base.B2i32(v903 < int32(0))
	if v903 < int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v936 = int32(45)
	goto L137
L136:
	;
	v936 = int32(43)
	goto L137
L137:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v929)+1)) = uint8(v936)
	if v903 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v940 = int32(1) - v743
	goto L140
L139:
	;
	v940 = v903
	goto L140
L140:
	;
	v945 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v940<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v929)+2)) = uint16(v945)
	v1126 = v928 + int32(4)
	goto L8
L141:
	;
	v961 = v954
	v962 = v719
	goto L144
L142:
	;
	v1007 = v954
	v1008 = v719
	goto L143
L143:
	;
	if base.Ui32(v1008) < base.Ui32(int32(100)) {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	v978 = v952 + v748 + v735 - v961
	v979 = int32(4)
	v982 = base.I32_div_u_s(v962, int32(_a_F_vector_out_9))
	v985 = v962 + v982*int32(-10000)
	v986 = int32(100)
	v987 = base.I32_div_u_s(v985, v986)
	v988 = int32(1)
	v990 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v987<<(uint(v988)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v978-v979))) = uint16(v990)
	v999 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v985-v987*v986)<<(uint(v988)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v978-int32(2)))) = uint16(v999)
	v1002 = v961 + v979
	if base.Ui32(int32(99999999)) < base.Ui32(v962) {
		v961 = v1002
		v962 = v982
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v1007 = v1002
	v1008 = v982
	goto L143
L146:
	;
	goto L145
L147:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1047) {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	v1046 = v1007
	v1047 = v1008
	goto L147
L149:
	;
	goto L150
L150:
	;
	v1029 = int32(2)
	v1031 = int32(_a_F_vector_out_10)
	v1033 = int32(100)
	v1034 = base.I32_div_u_s(v1008&v1031, v1033)
	v1042 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1008-v1034*v1033)&v1031<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v952+v748+v735-v1007-v1029))) = uint16(v1042)
	v1046 = v1007 | v1029
	v1047 = v1034
	goto L147
L151:
	;
	v1063 = int32(1)
	if v952 == v1063 {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	v1057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1047<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v952+v748+v735-v1046-int32(2)))) = uint16(v1057)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v1061 = v1047 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v952+v748))) = uint8(v1061)
	goto L151
L155:
	;
	v1126 = v1103 + int32(base.Ui32(v47)>>(uint(int32(31))%32))
	goto L8
L156:
	;
	if v743&int32(4) != 0 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	if v724 < int32(0) {
		goto L168
	} else {
		goto L169
	}
L159:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v748)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v748))) = v1068
	v1071 = int32(5)
	goto L161
L160:
	;
	v1071 = v1063
	goto L161
L161:
	;
	if v743&int32(2) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1074 = v1071 + v748
	v1077 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1074))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1074-int32(1)))) = uint16(v1077)
	v1082 = v1071 | int32(2)
	goto L164
L163:
	;
	v1082 = v1071
	goto L164
L164:
	;
	if v743&int32(1) != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1085 = v1082 + v748
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1085-int32(1)))) = uint8(v1088)
	goto L167
L166:
	;
	goto L167
L167:
	;
	v1092 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v748+v743))) = uint8(v1092)
	v1103 = v735 + int32(1)
	goto L155
L168:
	;
	v1100 = int32(2) - v724
	goto L170
L169:
	;
	v1100 = v743
	goto L170
L170:
	;
	v1103 = v1100
	goto L155
L171:
	;
	v1133 = v1127
	v1136 = v24
	goto L172
L172:
	;
	v1139 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1133))) = uint8(v1139)
	v1144 = *(*float32)(unsafe.Add(mBase, uint32(v10+int32(8)+v1136<<(uint(int32(2))%32))))
	v1146 = v1133 + int32(1)
	v1147 = int32(0)
	v1164 = base.I32_reinterpret_f32(v1144)
	v1166 = v1164 & int32(_a_F_vector_out_0)
	v1169 = int32(255)
	v1170 = int32(base.Ui32(v1164)>>(uint(int32(23))%32)) & v1169
	if v1170|v1166 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v2249 = v2244
	goto L6
L174:
	;
	v2244 = v2243 + v1146
	v2246 = v1136 + int32(1)
	if v2246 != v14 {
		v1133 = v2244
		v1136 = v2246
		goto L172
	} else {
		goto L337
	}
L175:
	;
	v1175 = base.B2i32(v1170 != v1169)
	goto L177
L176:
	;
	v1175 = v1147
	goto L177
L177:
	;
	if v1175 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	if v1166 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	if base.Ui32(int32(23)) < base.Ui32(v1170-int32(127)) {
		goto L198
	} else {
		goto L199
	}
L181:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vector_out[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1146)+2)) = uint8(v1179)
	v1182 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_vector_out[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1146))) = uint16(v1182)
	v2243 = int32(3)
	goto L174
L182:
	;
	goto L183
L183:
	;
	if v1164 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1187 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1146))) = uint8(v1187)
	goto L186
L185:
	;
	goto L186
L186:
	;
	v1191 = v1146 + int32(base.Ui32(v1164)>>(uint(int32(31))%32))
	if v1170 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1191))) = int64(8751735898823355977)
	if v1164 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	v1199 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1191))) = uint8(v1199)
	if v1164 < int32(0) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v1198 = int32(9)
	goto L192
L191:
	;
	v1198 = int32(8)
	goto L192
L192:
	;
	v2243 = v1198
	goto L174
L193:
	;
	v1205 = int32(2)
	goto L195
L194:
	;
	v1205 = int32(1)
	goto L195
L195:
	;
	v2243 = v1205
	goto L174
L196:
	;
	v1853 = int32(0)
	if v1164 < v1853 {
		goto L264
	} else {
		goto L265
	}
L197:
	;
	if base.Ui32(int32(_a_F_vector_out_1)) < base.Ui32(v1794) {
		v1836 = v1794
		v1841 = v1799
		v1852 = int32(8)
		goto L196
	} else {
		goto L255
	}
L198:
	;
	v1222 = v1166 << (uint(int32(2)) % 32)
	if v1170 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1210 = int32(-1)
	v1212 = int32(150) - v1170
	if v1166&(v1210<<(uint(v1212)%32)^v1210) != 0 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1794 = int32(base.Ui32(v1166|int32(_a_F_vector_out_2)) >> (uint(v1212) % 32))
	v1799 = v1147
	goto L197
L201:
	;
	v1225 = v1222 | int32(33554432)
	goto L203
L202:
	;
	v1225 = v1222
	goto L203
L203:
	;
	v1228 = int32(2)
	v1233 = v1225 + (base.B2i32(v1166 != int32(0)) | base.B2i32(base.Ui32(v1170) < base.Ui32(v1228)) ^ int32(-1))
	v1235 = v1225 | v1228
	if v1170 != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v1786 = v1769 + v1773
	v1787 = v1768 + v1785
	if base.Ui32(v1787) <= base.Ui32(int32(99999999)) {
		v1794 = v1787
		v1799 = v1786
		goto L197
	} else {
		goto L254
	}
L205:
	;
	v1704 = int32(0)
	v1705 = int32(10)
	v1706 = base.I32_div_u_s(v1690, v1705)
	v1708 = base.I32_div_u_s(v1694, v1705)
	if base.Ui32(v1708) < base.Ui32(v1706) {
		goto L248
	} else {
		goto L249
	}
L206:
	;
	v1608 = int32(0)
	v1609 = int32(10)
	v1610 = base.I32_div_u_s(v1594, v1609)
	v1612 = base.I32_div_u_s(v1598, v1609)
	if base.Ui32(v1610) <= base.Ui32(v1612) {
		goto L242
	} else {
		goto L243
	}
L207:
	;
	v1239 = v1170 - int32(152)
	goto L209
L208:
	;
	v1239 = int32(-151)
	goto L209
L209:
	;
	if int32(0) <= v1239 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1245 = int32(base.Ui32(v1239*int32(_a_F_vector_out_3)) >> (uint(int32(18)) % 32))
	v1248 = *(*int64)(unsafe.Add(mBase, uint32(v1245<<(uint(int32(3))%32))+uint32(_c_F_vector_out[2])))
	v1250 = v1248 & int64(4294967295)
	v1251 = base.I64_extend_i32_u(v1233)
	v1253 = int64(32)
	v1255 = base.I32_wrap_i64(int64(base.Ui64(v1250*v1251) >> (uint(v1253) % 64)))
	v1257 = int64(base.Ui64(v1248) >> (uint(v1253) % 64))
	v1258 = v1251 * v1257
	v1260 = v1255 + base.I32_wrap_i64(v1258)
	v1267 = v1245 - v1239
	v1272 = v1267 + int32(base.Ui32(v1245*int32(_a_F_vector_out_4))>>(uint(int32(19))%32))
	v1273 = int32(5) - v1272
	v1276 = v1272 + int32(27)
	v1278 = (base.B2i32(base.Ui32(v1260) < base.Ui32(v1255))+base.I32_wrap_i64(int64(base.Ui64(v1258)>>(uint(v1253)%64))))<<(uint(v1273)%32) | int32(base.Ui32(v1260)>>(uint(v1276)%32))
	v1279 = base.I64_extend_i32_u(v1235)
	v1283 = base.I32_wrap_i64(int64(base.Ui64(v1250*v1279) >> (uint(v1253) % 64)))
	v1284 = v1279 * v1257
	v1286 = v1283 + base.I32_wrap_i64(v1284)
	v1294 = (base.B2i32(base.Ui32(v1286) < base.Ui32(v1283))+base.I32_wrap_i64(int64(base.Ui64(v1284)>>(uint(v1253)%64))))<<(uint(v1273)%32) | int32(base.Ui32(v1286)>>(uint(v1276)%32))
	v1295 = base.I64_extend_i32_u(v1225)
	v1299 = base.I32_wrap_i64(int64(base.Ui64(v1250*v1295) >> (uint(v1253) % 64)))
	v1300 = v1295 * v1257
	v1302 = v1299 + base.I32_wrap_i64(v1300)
	v1310 = (base.B2i32(base.Ui32(v1302) < base.Ui32(v1299))+base.I32_wrap_i64(int64(base.Ui64(v1300)>>(uint(v1253)%64))))<<(uint(v1273)%32) | int32(base.Ui32(v1302)>>(uint(v1276)%32))
	v1311 = int32(0)
	if v1245 != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	goto L212
L212:
	;
	v1454 = v1239 * int32(-732923)
	v1456 = int32(base.Ui32(v1454) >> (uint(int32(20)) % 32))
	v1457 = v1239 + v1456
	v1461 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_vector_out_5)-v1457<<(uint(int32(3))%32))))
	v1463 = v1461 & int64(4294967295)
	v1464 = base.I64_extend_i32_u(v1233)
	v1466 = int64(32)
	v1468 = base.I32_wrap_i64(int64(base.Ui64(v1463*v1464) >> (uint(v1466) % 64)))
	v1470 = int64(base.Ui64(v1461) >> (uint(v1466) % 64))
	v1471 = v1464 * v1470
	v1473 = v1468 + base.I32_wrap_i64(v1471)
	v1484 = v1456 - int32(base.Ui32(v1457*int32(-1217359))>>(uint(int32(19))%32))
	v1485 = int32(4) - v1484
	v1488 = v1484 + int32(28)
	v1490 = (base.B2i32(base.Ui32(v1473) < base.Ui32(v1468))+base.I32_wrap_i64(int64(base.Ui64(v1471)>>(uint(v1466)%64))))<<(uint(v1485)%32) | int32(base.Ui32(v1473)>>(uint(v1488)%32))
	v1491 = base.I64_extend_i32_u(v1225)
	v1495 = base.I32_wrap_i64(int64(base.Ui64(v1463*v1491) >> (uint(v1466) % 64)))
	v1496 = v1491 * v1470
	v1498 = v1495 + base.I32_wrap_i64(v1496)
	v1506 = (base.B2i32(base.Ui32(v1498) < base.Ui32(v1495))+base.I32_wrap_i64(int64(base.Ui64(v1496)>>(uint(v1466)%64))))<<(uint(v1485)%32) | int32(base.Ui32(v1498)>>(uint(v1488)%32))
	v1507 = base.I64_extend_i32_u(v1235)
	v1511 = base.I32_wrap_i64(int64(base.Ui64(v1463*v1507) >> (uint(v1466) % 64)))
	v1512 = v1470 * v1507
	v1514 = v1511 + base.I32_wrap_i64(v1512)
	v1522 = (base.B2i32(base.Ui32(v1514) < base.Ui32(v1511))+base.I32_wrap_i64(int64(base.Ui64(v1512)>>(uint(v1466)%64))))<<(uint(v1485)%32) | int32(base.Ui32(v1514)>>(uint(v1488)%32))
	v1524 = v1522 - int32(1)
	if v1456 != 0 {
		goto L234
	} else {
		goto L235
	}
L213:
	;
	v1315 = int32(10)
	v1316 = base.I32_div_u_s(v1294-int32(1), v1315)
	v1318 = base.I32_div_u_s(v1278, v1315)
	if base.Ui32(v1316) <= base.Ui32(v1318) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v1362 = v1311
	goto L215
L215:
	;
	v1368 = base.I32_rem_u_s(v1225, int32(5))
	if v1368 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L216:
	;
	v1321 = v1245 - int32(1)
	v1324 = *(*int64)(unsafe.Add(mBase, uint32(v1321<<(uint(int32(3))%32))+uint32(_c_F_vector_out[2])))
	v1328 = int64(32)
	v1330 = base.I32_wrap_i64(int64(base.Ui64(v1324&int64(4294967295)*v1295) >> (uint(v1328) % 64)))
	v1333 = int64(base.Ui64(v1324)>>(uint(v1328)%64)) * v1295
	v1335 = v1330 + base.I32_wrap_i64(v1333)
	v1346 = v1267 + int32(base.Ui32(v1321*int32(_a_F_vector_out_4))>>(uint(int32(19))%32))
	v1354 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1335) < base.Ui32(v1330))+base.I32_wrap_i64(int64(base.Ui64(v1333)>>(uint(v1328)%64))))<<(uint(int32(6)-v1346)%32)|int32(base.Ui32(v1335)>>(uint(v1346+int32(26))%32)), int32(10))
	v1355 = v1354
	goto L218
L217:
	;
	v1355 = v1311
	goto L218
L218:
	;
	if base.Ui32(int32(33)) < base.Ui32(v1239) {
		v1687 = v1310
		v1690 = v1294
		v1692 = v1245
		v1693 = v1355
		v1694 = v1278
		goto L205
	} else {
		goto L219
	}
L219:
	;
	v1362 = v1355
	goto L215
L220:
	;
	v1374 = v1225
	v1375 = v1311
	goto L223
L221:
	;
	goto L222
L222:
	;
	v1399 = int32(0)
	v1401 = base.I32_rem_u_s(v1235, int32(5))
	if v1401 == v1399 {
		goto L227
	} else {
		goto L228
	}
L223:
	;
	v1391 = v1375 + int32(1)
	v1392 = int32(5)
	v1393 = base.I32_div_u_s(v1374, v1392)
	v1395 = base.I32_rem_u_s(v1393, v1392)
	if v1395 == int32(0) {
		v1374 = v1393
		v1375 = v1391
		goto L223
	} else {
		goto L225
	}
L224:
	;
	if base.Ui32(v1391) < base.Ui32(v1245) {
		v1687 = v1310
		v1690 = v1294
		v1692 = v1245
		v1693 = v1362
		v1694 = v1278
		goto L205
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v1591 = v1310
	v1594 = v1294
	v1596 = v1245
	v1597 = v1362
	v1598 = v1278
	goto L206
L227:
	;
	v1407 = v1399
	v1410 = v1235
	goto L230
L228:
	;
	v1434 = v1399
	goto L229
L229:
	;
	v1687 = v1310
	v1690 = v1294 - base.B2i32(base.Ui32(v1245) <= base.Ui32(v1434))
	v1692 = v1245
	v1693 = v1362
	v1694 = v1278
	goto L205
L230:
	;
	v1424 = v1407 + int32(1)
	v1425 = int32(5)
	v1426 = base.I32_div_u_s(v1410, v1425)
	v1428 = base.I32_rem_u_s(v1426, v1425)
	if v1428 == int32(0) {
		v1407 = v1424
		v1410 = v1426
		goto L230
	} else {
		goto L232
	}
L231:
	;
	v1434 = v1424
	goto L229
L232:
	;
	goto L231
L233:
	;
	v1579 = int32(-1)
	if v1225&(v1579<<(uint(v1456-int32(1))%32)^v1579)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v1454)) != 0 {
		v1687 = v1506
		v1690 = v1522
		v1692 = v1457
		v1693 = v1568
		v1694 = v1490
		goto L205
	} else {
		goto L241
	}
L234:
	;
	v1525 = int32(10)
	v1526 = base.I32_div_u_s(v1524, v1525)
	v1528 = base.I32_div_u_s(v1490, v1525)
	if base.Ui32(v1526) <= base.Ui32(v1528) {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	v1575 = v1147
	goto L236
L236:
	;
	v1591 = v1506
	v1594 = v1524
	v1596 = v1457
	v1597 = v1575
	v1598 = v1490
	goto L206
L237:
	;
	v1531 = int32(1) - v1457
	v1534 = *(*int64)(unsafe.Add(mBase, uint32(v1531<<(uint(int32(3))%32))+uint32(_c_F_vector_out[3])))
	v1538 = int64(32)
	v1540 = base.I32_wrap_i64(int64(base.Ui64(v1534&int64(4294967295)*v1491) >> (uint(v1538) % 64)))
	v1543 = int64(base.Ui64(v1534)>>(uint(v1538)%64)) * v1491
	v1545 = v1540 + base.I32_wrap_i64(v1543)
	v1558 = v1456 + (int32(base.Ui32(v1531*int32(_a_F_vector_out_4))>>(uint(int32(19))%32)) ^ int32(-1))
	v1566 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1545) < base.Ui32(v1540))+base.I32_wrap_i64(int64(base.Ui64(v1543)>>(uint(v1538)%64))))<<(uint(int32(4)-v1558)%32)|int32(base.Ui32(v1545)>>(uint(v1558+int32(28))%32)), int32(10))
	v1568 = v1566
	goto L239
L238:
	;
	v1568 = v1147
	goto L239
L239:
	;
	if v1456 != int32(1) {
		goto L233
	} else {
		goto L240
	}
L240:
	;
	v1575 = v1568
	goto L236
L241:
	;
	v1591 = v1506
	v1594 = v1522
	v1596 = v1457
	v1597 = v1568
	v1598 = v1490
	goto L206
L242:
	;
	v1656 = v1591
	v1657 = v1608
	v1662 = v1597
	v1663 = v1598
	v1673 = int32(0)
	goto L244
L243:
	;
	v1618 = v1591
	v1619 = v1608
	v1620 = v1610
	v1621 = int32(1)
	v1622 = v1612
	v1624 = v1597
	goto L245
L244:
	;
	v1675 = v1662 & int32(255)
	v1768 = v1656
	v1769 = v1657
	v1773 = v1596
	v1785 = (v1673|base.B2i32(v1675 != int32(5))|v1656)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1675)) | base.B2i32(v1656 == v1663)
	goto L204
L245:
	;
	v1636 = v1619 + int32(1)
	v1637 = int32(10)
	v1638 = base.I32_div_u_s(v1618, v1637)
	v1641 = v1618 - v1638*v1637
	v1646 = v1621 & base.B2i32(v1624&int32(255) == int32(0))
	v1648 = base.I32_div_u_s(v1620, v1637)
	v1650 = base.I32_div_u_s(v1622, v1637)
	if base.Ui32(v1650) < base.Ui32(v1648) {
		v1618 = v1638
		v1619 = v1636
		v1620 = v1648
		v1621 = v1646
		v1622 = v1650
		v1624 = v1641
		goto L245
	} else {
		goto L247
	}
L246:
	;
	v1656 = v1638
	v1657 = v1636
	v1662 = v1641
	v1663 = v1622
	v1673 = v1646 ^ int32(1)
	goto L244
L247:
	;
	goto L246
L248:
	;
	v1712 = v1687
	v1713 = v1704
	v1714 = v1706
	v1716 = v1708
	goto L251
L249:
	;
	v1743 = v1687
	v1744 = v1704
	v1749 = v1693
	v1750 = v1694
	goto L250
L250:
	;
	v1768 = v1743
	v1769 = v1744
	v1773 = v1692
	v1785 = base.B2i32(v1743 == v1750) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1749&int32(255)))
	goto L204
L251:
	;
	v1730 = v1713 + int32(1)
	v1731 = int32(10)
	v1732 = base.I32_div_u_s(v1712, v1731)
	v1734 = base.I32_div_u_s(v1714, v1731)
	v1736 = base.I32_div_u_s(v1716, v1731)
	if base.Ui32(v1736) < base.Ui32(v1734) {
		v1712 = v1732
		v1713 = v1730
		v1714 = v1734
		v1716 = v1736
		goto L251
	} else {
		goto L253
	}
L252:
	;
	v1743 = v1732
	v1744 = v1730
	v1749 = v1712 - v1732*int32(10)
	v1750 = v1716
	goto L250
L253:
	;
	goto L252
L254:
	;
	v1836 = v1787
	v1841 = v1786
	v1852 = int32(9)
	goto L196
L255:
	;
	if base.Ui32(int32(_a_F_vector_out_6)) < base.Ui32(v1794) {
		v1836 = v1794
		v1841 = v1799
		v1852 = int32(7)
		goto L196
	} else {
		goto L256
	}
L256:
	;
	if base.Ui32(int32(_a_F_vector_out_7)) < base.Ui32(v1794) {
		v1836 = v1794
		v1841 = v1799
		v1852 = int32(6)
		goto L196
	} else {
		goto L257
	}
L257:
	;
	if base.Ui32(int32(_a_F_vector_out_8)) < base.Ui32(v1794) {
		v1836 = v1794
		v1841 = v1799
		v1852 = int32(5)
		goto L196
	} else {
		goto L258
	}
L258:
	;
	if base.Ui32(int32(999)) < base.Ui32(v1794) {
		v1836 = v1794
		v1841 = v1799
		v1852 = int32(4)
		goto L196
	} else {
		goto L259
	}
L259:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1794) {
		v1836 = v1794
		v1841 = v1799
		v1852 = int32(3)
		goto L196
	} else {
		goto L260
	}
L260:
	;
	if base.Ui32(int32(9)) < base.Ui32(v1794) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1832 = int32(2)
	goto L263
L262:
	;
	v1832 = int32(1)
	goto L263
L263:
	;
	v1836 = v1794
	v1841 = v1799
	v1852 = v1832
	goto L196
L264:
	;
	v1856 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1146))) = uint8(v1856)
	v1859 = int32(1)
	goto L266
L265:
	;
	v1859 = v1853
	goto L266
L266:
	;
	v1860 = v1852 + v1841
	if base.Ui32(v1860+int32(3)) <= base.Ui32(int32(9)) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	v2071 = int32(0)
	if base.Ui32(int32(_a_F_vector_out_9)) <= base.Ui32(v1836) {
		goto L307
	} else {
		goto L308
	}
L268:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1865))) = v2067
	v2069 = v2066
	goto L267
L269:
	;
	v1865 = v1146 + v1859
	v1866 = int32(0)
	if v1860 <= v1866 {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	goto L271
L271:
	;
	if v1841 != 0 {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	v2066 = int32(2) - v1860
	v2067 = int64(3472328296227679792)
	goto L268
L273:
	;
	goto L274
L274:
	;
	if int32(0) <= v1841 {
		v2066 = v1866
		v2067 = int64(3472328296227680304)
		goto L268
	} else {
		goto L275
	}
L275:
	;
	v2069 = int32(1)
	goto L267
L276:
	;
	v1925 = int32(0)
	if base.Ui32(int32(_a_F_vector_out_9)) <= base.Ui32(v1908) {
		goto L284
	} else {
		goto L285
	}
L277:
	;
	v1908 = v1836
	v1913 = v1852
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1879 = v1836
	v1881 = v1852
	goto L280
L280:
	;
	if v1879&int32(1) != 0 {
		v1908 = v1879
		v1913 = v1881
		goto L276
	} else {
		goto L282
	}
L281:
	;
	v1908 = v1879
	v1913 = v1881
	goto L276
L282:
	;
	v1902 = base.I32_div_u_s(v1879, int32(10))
	if int32(0)-v1879 == v1902*int32(-10) {
		v1879 = v1902
		v1881 = v1881 - int32(1)
		goto L280
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	v1932 = v1908
	v1933 = v1925
	goto L287
L285:
	;
	v1978 = v1908
	v1979 = v1925
	goto L286
L286:
	;
	if base.Ui32(v1978) < base.Ui32(int32(100)) {
		goto L291
	} else {
		goto L292
	}
L287:
	;
	v1949 = v1146 + v1859 + v1913 - v1933
	v1953 = base.I32_div_u_s(v1932, int32(_a_F_vector_out_9))
	v1956 = v1932 + v1953*int32(-10000)
	v1957 = int32(100)
	v1958 = base.I32_div_u_s(v1956, v1957)
	v1959 = int32(1)
	v1961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1958<<(uint(v1959)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1949-int32(3)))) = uint16(v1961)
	v1970 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1956-v1958*v1957)<<(uint(v1959)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1949-v1959))) = uint16(v1970)
	v1973 = v1933 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v1932) {
		v1932 = v1953
		v1933 = v1973
		goto L287
	} else {
		goto L289
	}
L288:
	;
	v1978 = v1953
	v1979 = v1973
	goto L286
L289:
	;
	goto L288
L290:
	;
	v2020 = v1860 - int32(1)
	v2021 = v1146 + v1859
	if base.Ui32(int32(10)) <= base.Ui32(v2018) {
		goto L295
	} else {
		goto L296
	}
L291:
	;
	v2017 = v1979
	v2018 = v1978
	goto L290
L292:
	;
	goto L293
L293:
	;
	v2002 = int32(_a_F_vector_out_10)
	v2004 = int32(100)
	v2005 = base.I32_div_u_s(v1978&v2002, v2004)
	v2013 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1978-v2005*v2004)&v2002<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1146+v1859+v1913+(v1979^int32(-1))))) = uint16(v2013)
	v2017 = v1979 | int32(2)
	v2018 = v2005
	goto L290
L294:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2021))) = uint8(v2035)
	if base.Ui32(int32(2)) <= base.Ui32(v1913) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	v2028 = v2018 << (uint(int32(1)) % 32)
	v2029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2028)+uint32(_c_F_vector_out[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1146+(v1859+v1913-v2017)))) = uint8(v2029)
	v2031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2028)+uint32(_c_F_vector_out[4]))))
	v2035 = v2031
	goto L294
L296:
	;
	goto L297
L297:
	;
	v2035 = v2018 | int32(48)
	goto L294
L298:
	;
	v2039 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2021)+1)) = uint8(v2039)
	v2044 = v1913 + int32(1)
	goto L300
L299:
	;
	v2044 = int32(1)
	goto L300
L300:
	;
	v2045 = v2044 + v1859
	v2046 = v1146 + v2045
	v2047 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2046))) = uint8(v2047)
	v2052 = base.B2i32(v2020 < int32(0))
	if v2020 < int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v2053 = int32(45)
	goto L303
L302:
	;
	v2053 = int32(43)
	goto L303
L303:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2046)+1)) = uint8(v2053)
	if v2020 < int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v2057 = int32(1) - v1860
	goto L306
L305:
	;
	v2057 = v2020
	goto L306
L306:
	;
	v2062 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2057<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2046)+2)) = uint16(v2062)
	v2243 = v2045 + int32(4)
	goto L174
L307:
	;
	v2078 = v2071
	v2079 = v1836
	goto L310
L308:
	;
	v2124 = v2071
	v2125 = v1836
	goto L309
L309:
	;
	if base.Ui32(v2125) < base.Ui32(int32(100)) {
		goto L314
	} else {
		goto L315
	}
L310:
	;
	v2095 = v2069 + v1865 + v1852 - v2078
	v2096 = int32(4)
	v2099 = base.I32_div_u_s(v2079, int32(_a_F_vector_out_9))
	v2102 = v2079 + v2099*int32(-10000)
	v2103 = int32(100)
	v2104 = base.I32_div_u_s(v2102, v2103)
	v2105 = int32(1)
	v2107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2104<<(uint(v2105)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2095-v2096))) = uint16(v2107)
	v2116 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2102-v2104*v2103)<<(uint(v2105)%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2095-int32(2)))) = uint16(v2116)
	v2119 = v2078 + v2096
	if base.Ui32(int32(99999999)) < base.Ui32(v2079) {
		v2078 = v2119
		v2079 = v2099
		goto L310
	} else {
		goto L312
	}
L311:
	;
	v2124 = v2119
	v2125 = v2099
	goto L309
L312:
	;
	goto L311
L313:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2164) {
		goto L318
	} else {
		goto L319
	}
L314:
	;
	v2163 = v2124
	v2164 = v2125
	goto L313
L315:
	;
	goto L316
L316:
	;
	v2146 = int32(2)
	v2148 = int32(_a_F_vector_out_10)
	v2150 = int32(100)
	v2151 = base.I32_div_u_s(v2125&v2148, v2150)
	v2159 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2125-v2151*v2150)&v2148<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2069+v1865+v1852-v2124-v2146))) = uint16(v2159)
	v2163 = v2124 | v2146
	v2164 = v2151
	goto L313
L317:
	;
	v2180 = int32(1)
	if v2069 == v2180 {
		goto L322
	} else {
		goto L323
	}
L318:
	;
	v2174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2164<<(uint(int32(1))%32))+uint32(_c_F_vector_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2069+v1865+v1852-v2163-int32(2)))) = uint16(v2174)
	goto L317
L319:
	;
	goto L320
L320:
	;
	v2178 = v2164 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2069+v1865))) = uint8(v2178)
	goto L317
L321:
	;
	v2243 = v2220 + int32(base.Ui32(v1164)>>(uint(int32(31))%32))
	goto L174
L322:
	;
	if v1860&int32(4) != 0 {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	goto L324
L324:
	;
	if v1841 < int32(0) {
		goto L334
	} else {
		goto L335
	}
L325:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v1865))) = v2185
	v2188 = int32(5)
	goto L327
L326:
	;
	v2188 = v2180
	goto L327
L327:
	;
	if v1860&int32(2) != 0 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v2191 = v2188 + v1865
	v2194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2191))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2191-int32(1)))) = uint16(v2194)
	v2199 = v2188 | int32(2)
	goto L330
L329:
	;
	v2199 = v2188
	goto L330
L330:
	;
	if v1860&int32(1) != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v2202 = v2199 + v1865
	v2205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2202))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2202-int32(1)))) = uint8(v2205)
	goto L333
L332:
	;
	goto L333
L333:
	;
	v2209 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1865+v1860))) = uint8(v2209)
	v2220 = v1852 + int32(1)
	goto L321
L334:
	;
	v2217 = int32(2) - v1841
	goto L336
L335:
	;
	v2217 = v1860
	goto L336
L336:
	;
	v2220 = v2217
	goto L321
L337:
	;
	goto L173
L338:
	;
	F_pfree(m, v10)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L1
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	return v20
L341:
	;
	goto L340
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
	F_errmsg(m, int32(_a_F_vector_recv_0), v10+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_vector_recv_1), int32(88), int32(_a_F_vector_recv_2))
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
	F_errmsg(m, int32(_a_F_vector_recv_3), v10)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_vector_recv_1), int32(393), int32(_a_F_vector_recv_4))
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
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
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
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v104 float32
	_ = v104
	var v106 float32
	_ = v106
	var v121 int32
	_ = v121
	var v133 float32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
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
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L27
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
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v121 = int32(0)
	goto L20
L13:
	;
	v54 = v48
	v62 = int32(0)
	goto L16
L14:
	;
	v91 = v48
	goto L15
L15:
	;
	v101 = v91 << (uint(int32(2)) % 32)
	v104 = *(*float32)(unsafe.Add(mBase, uint32(v101+v43)))
	v106 = *(*float32)(unsafe.Add(mBase, uint32(v101+v45)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v101))) = base.F32_sub(v104, v106)
	goto L12
L16:
	;
	v63 = int32(2)
	v64 = v54 << (uint(v63) % 32)
	v67 = *(*float32)(unsafe.Add(mBase, uint32(v64+v43)))
	v69 = *(*float32)(unsafe.Add(mBase, uint32(v64+v45)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v64))) = base.F32_sub(v67, v69)
	v73 = v64 | int32(4)
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v73+v43)))
	v78 = *(*float32)(unsafe.Add(mBase, uint32(v73+v45)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v73))) = base.F32_sub(v76, v78)
	v82 = v54 + v63
	v84 = v62 + v63
	if v84 != v39&int32(_a_F_vector_sub_0) {
		v54 = v82
		v62 = v84
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v39&int32(1) == int32(0) {
		goto L12
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v91 = v82
	goto L15
L20:
	;
	v133 = *(*float32)(unsafe.Add(mBase, uint32(v47+v121<<(uint(int32(2))%32))))
	if base.F32_ne(base.F32_abs(v133), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	v138 = v121 + int32(1)
	if v39 != v138 {
		v121 = v138
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L4
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v149
	F_errmsg(m, int32(_a_F_vector_sub_1), v13)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_vector_sub_2), int32(76), int32(_a_F_vector_sub_3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 float32
	_ = v56
	var v57 float32
	_ = v57
	var v60 float32
	_ = v60
	var v64 float32
	_ = v64
	var v68 float32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v110 float32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v177 float32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	F_CheckDim_2(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.B2i32(v22 != int32(-1))&base.B2i32(v22 != v23) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = base.B2i32(v23 <= int32(0))
	if v23 <= int32(0) {
		v120 = v2
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L39
	}
L7:
	;
	F_CheckNnz(m, v120, v23)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L19
	}
L8:
	;
	v35 = v18 + int32(8)
	v36 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v23) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = v36
	v42 = v2
	v50 = v2
	goto L12
L10:
	;
	v81 = v36
	v82 = v2
	goto L11
L11:
	;
	v95 = v81
	v96 = v82
	v105 = v2
	goto L16
L12:
	;
	v55 = v35 + v41<<(uint(int32(2))%32)
	v56 = *(*float32)(unsafe.Add(mBase, uint32(v55)))
	v57 = float32(0)
	v60 = *(*float32)(unsafe.Add(mBase, uint32(v55)+4))
	v64 = *(*float32)(unsafe.Add(mBase, uint32(v55)+8))
	v68 = *(*float32)(unsafe.Add(mBase, uint32(v55)+12))
	v71 = v42 + base.F32_ne(v56, v57) + base.F32_ne(v60, v57) + base.F32_ne(v64, v57) + base.F32_ne(v68, v57)
	v72 = int32(4)
	v73 = v41 + v72
	v75 = v50 + v72
	if v75 != v23&int32(_a_F_vector_to_sparsevec_0) {
		v41 = v73
		v42 = v71
		v50 = v75
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v23&int32(3) == int32(0) {
		v120 = v71
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v81 = v73
	v82 = v71
	goto L11
L16:
	;
	v110 = *(*float32)(unsafe.Add(mBase, uint32(v35+v95<<(uint(int32(2))%32))))
	v113 = v96 + base.F32_ne(v110, float32(0))
	v114 = int32(1)
	v117 = v105 + v114
	if v117 != v23&int32(3) {
		v95 = v95 + v114
		v96 = v113
		v105 = v117
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v120 = v113
	goto L7
L18:
	;
	goto L17
L19:
	;
	v135 = F_mul_size(m, int32(4), v120)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v137 = F_add_size(m, int32(16), v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v140 = F_mul_size(m, int32(4), v120)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v142 = F_add_size(m, v137, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v144 = F_palloc0(m, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v142 << (uint(int32(2)) % 32)
	if v33 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L36
	}
L26:
	;
	v154 = v144 + int32(16)
	v160 = int32(0)
	v162 = v160
	v163 = v160
	goto L29
L27:
	;
	goto L28
L28:
	;
	m.G0 = v15 + int32(16)
	return v144
L29:
	;
	v177 = *(*float32)(unsafe.Add(mBase, uint32(v18+int32(8)+v162<<(uint(int32(2))%32))))
	if base.F32_ne(v177, float32(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	if v180 <= v163 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	v190 = v163
	goto L33
L33:
	;
	v193 = v162 + int32(1)
	if v193 != v23 {
		v162 = v193
		v163 = v190
		goto L29
	} else {
		goto L35
	}
L34:
	;
	v183 = v163 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v154+v183))) = v162
	*(*float32)(unsafe.Add(mBase, uint32(v154+v120<<(uint(int32(2))%32)+v183))) = v177
	v190 = v163 + int32(1)
	goto L33
L35:
	;
	goto L30
L36:
	;
	F_errmsg_internal(m, int32(_a_F_vector_to_sparsevec_1), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_vector_to_sparsevec_2), int32(632), int32(_a_F_vector_to_sparsevec_3))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
	F_errmsg(m, int32(_a_F_vector_to_sparsevec_4), v15)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_vector_to_sparsevec_2), int32(62), int32(_a_F_vector_to_sparsevec_5))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
