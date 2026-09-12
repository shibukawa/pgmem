package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sparsevec(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v14 == int32(-1) {
			m.G0 = v7 + int32(16)
			return v10
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if v14 == v17 {
				m.G0 = v7 + int32(16)
				return v10
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg(m, int32(_a_F_sparsevec_0), v7)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_sparsevec_1), int32(62), int32(_a_F_sparsevec_2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
func F_sparsevec_l2_norm(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 float64
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v52 float32
	_ = v52
	var v53 float64
	_ = v53
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 float64
	_ = v66
	var v75 int32
	_ = v75
	var v77 float64
	_ = v77
	var v81 int32
	_ = v81
	var v87 float32
	_ = v87
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 float64
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v2 = float64(0)
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		if v15 <= int32(0) {
			v98 = v2
		} else {
			v22 = v11 + v15<<(uint(int32(2))%32) + int32(16)
			v24 = v15 & int32(3)
			if base.Ui32(v15) < base.Ui32(int32(4)) {
				v64 = int32(0)
				v66 = v2
			} else {
				v31 = int32(0)
				v33 = v2
				v38 = v4
				for {
					v42 = v22 + v31<<(uint(int32(2))%32)
					v43 = *(*float32)(unsafe.Add(mBase, uint32(v42)+12))
					v44 = base.F64_promote_f32(v43)
					v46 = *(*float32)(unsafe.Add(mBase, uint32(v42)+8))
					v47 = base.F64_promote_f32(v46)
					v49 = *(*float32)(unsafe.Add(mBase, uint32(v42)+4))
					v50 = base.F64_promote_f32(v49)
					v52 = *(*float32)(unsafe.Add(mBase, uint32(v42)))
					v53 = base.F64_promote_f32(v52)
					v58 = base.F64_add(base.F64_mul(v44, v44), base.F64_add(base.F64_mul(v47, v47), base.F64_add(base.F64_mul(v50, v50), base.F64_add(base.F64_mul(v53, v53), v33))))
					v59 = int32(4)
					v60 = v31 + v59
					v62 = v38 + v59
					if v62 != v15&int32(2147483644) {
						v31 = v60
						v33 = v58
						v38 = v62
						continue
					} else {
						break
					}
					break
				}
				v64 = v60
				v66 = v58
			}
			if v24 == int32(0) {
				v98 = v66
			} else {
				v75 = v64
				v77 = v66
				v81 = v4
				for {
					v87 = *(*float32)(unsafe.Add(mBase, uint32(v22+v75<<(uint(int32(2))%32))))
					v88 = base.F64_promote_f32(v87)
					v90 = base.F64_add(base.F64_mul(v88, v88), v77)
					v91 = int32(1)
					v94 = v81 + v91
					if v94 != v24 {
						v75 = v75 + v91
						v77 = v90
						v81 = v94
						continue
					} else {
						break
					}
					break
				}
				v98 = v90
			}
		}
		v106 = F_Float8GetDatum(m, base.F64_sqrt(v98))
		mBase = m.M
		v107 = m.ExcPending
		if v107 != 0 {
			return int32(0)
		} else {
			return v106
		}
	}
}
func F_sparsevec_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 float32
	_ = v66
	var v69 int32
	_ = v69
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v81 float32
	_ = v81
	var v83 float32
	_ = v83
	var v89 int32
	_ = v89
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 float32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 float32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v157 int32
	_ = v157
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(16)
	v24 = v8 + v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v26 = int32(2)
	v28 = v24 + v25<<(uint(v26)%32)
	v30 = v3 + v23
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v34 = v30 + v31<<(uint(v26)%32)
	if v31 < v25 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return base.B2i32(v157 != int32(0))
L5:
	;
	v36 = v31
	goto L7
L6:
	;
	v36 = v25
	goto L7
L7:
	;
	if int32(0) < v36 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	if v25 <= v31 {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v55 = v41 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v24)))
	if v57 < v59 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v66 = *(*float32)(unsafe.Add(mBase, uint32(v34+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v66, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v59 < v57 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v69 = int32(-1)
	goto L18
L17:
	;
	v69 = int32(1)
	goto L18
L18:
	;
	v157 = v69
	goto L4
L19:
	;
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v28+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v76, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v81, v83) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v79 = int32(1)
	goto L24
L23:
	;
	v79 = int32(-1)
	goto L24
L24:
	;
	v157 = v79
	goto L4
L25:
	;
	v157 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v81, v83) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v157 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v89 = v41 + int32(1)
	if v89 != v36 {
		v41 = v89
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L12
L32:
	;
	if v31 <= v25 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v108 = v31 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24+v108)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v111 <= v110 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*float32)(unsafe.Add(mBase, uint32(v108+v28)))
	if base.F32_lt(v116, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v119 = int32(1)
	goto L37
L36:
	;
	v119 = int32(-1)
	goto L37
L37:
	;
	v157 = v119
	goto L4
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v138 < v137 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v137 = v122
	goto L38
L40:
	;
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v125 = v36 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v30+v125)))
	if v123 <= v127 {
		v137 = v123
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v132 = *(*float32)(unsafe.Add(mBase, uint32(v125+v34)))
	if base.F32_lt(v132, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v135 = int32(-1)
	goto L45
L44:
	;
	v135 = int32(1)
	goto L45
L45:
	;
	v157 = v135
	goto L4
L46:
	;
	v157 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v157 = base.B2i32(v137 < v138)
	goto L4
}
func F_sparsevec_to_halfvec(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 float32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v10 = m.G0
	v12 = v10 - int32(16)
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	F_CheckDim_1(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.B2i32(v19 != int32(-1))&base.B2i32(v21 != v19) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = F_mul_size(m, int32(2), v21)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v34 = F_add_size(m, int32(8), v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v36 = F_palloc0(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v34 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if int32(0) < v42 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = v15 + int32(16)
	v53 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v12 + int32(16)
	return v36
L13:
	;
	v63 = v53 << (uint(int32(2)) % 32)
	v65 = *(*float32)(unsafe.Add(mBase, uint32(v46+v20<<(uint(int32(2))%32)+v63)))
	v66 = F_Float4ToHalf(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+v46)))
	v70 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(8)+v69<<(uint(v70)%32)))) = uint16(v66)
	v75 = v53 + v70
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v75 < v76 {
		v53 = v75
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v19
	F_errmsg(m, int32(_a_F_sparsevec_to_halfvec_0), v12)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_sparsevec_to_halfvec_1), int32(92), int32(_a_F_sparsevec_to_halfvec_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_typmod_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F_ArrayGetIntegerTypmods(m, v8, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v19 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_sparsevec_typmod_in_0), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_sparsevec_typmod_in_1), int32(495), int32(_a_F_sparsevec_typmod_in_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
					if base.Ui32(int32(1000000001)) <= base.Ui32(v19) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(1000000000)
								F_errmsg(m, int32(_a_F_sparsevec_typmod_in_3), v5)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_sparsevec_typmod_in_1), int32(500), int32(_a_F_sparsevec_typmod_in_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
						m.G0 = v5 + int32(16)
						return v19
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_sparsevec_typmod_in_4), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_sparsevec_typmod_in_1), int32(490), int32(_a_F_sparsevec_typmod_in_2))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
