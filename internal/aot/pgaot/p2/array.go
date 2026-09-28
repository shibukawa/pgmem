package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayCheckBounds(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v20 = v4
	goto L3
L3:
	;
	v24 = v20 << (uint(int32(2)) % 32)
	v25 = l2 + v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
	if base.B2i32(v26 < int32(0)) == base.B2i32(v30+v26 < v30) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v38 = F_errsave_start(m, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v35 = v20 + int32(1)
	if l0 != v35 {
		v20 = v35
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	goto L1
L9:
	;
	return
L10:
	;
	if v38 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v45
	F_errmsg(m, int32(_a_F_ArrayCheckBounds_0), v11)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, int32(0), int32(_a_F_ArrayCheckBounds_1), int32(141), int32(_a_F_ArrayCheckBounds_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L1
}
func F_ArrayGetOffset(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	v5 = int32(0)
	v14 = l0 - int32(1)
	if v14 < v5 {
		v92 = v5
	} else {
		v17 = int32(1)
		if v14 != 0 {
			v26 = v14
			v27 = v17
			v28 = v5
			v33 = v5
			for {
				v34 = int32(2)
				v35 = v26 << (uint(v34) % 32)
				v37 = v35 - int32(4)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l3+v37)))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l2+v37)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v35+l1)))
				v45 = v44 * v27
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v35+l3)))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v35+l2)))
				v54 = (v39-v41)*v45 + ((v48-v50)*v27 + v28)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l1+v37)))
				v57 = v56 * v45
				v59 = v26 - v34
				v61 = v33 + v34
				if v61 != l0&int32(-2) {
					v26 = v59
					v27 = v57
					v28 = v54
					v33 = v61
					continue
				} else {
					break
				}
				break
			}
			if l0&int32(1) == int32(0) {
				v92 = v54
			} else {
				v69 = v59
				v70 = v57
				v71 = v54
				v78 = v69 << (uint(int32(2)) % 32)
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l3+v78)))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+l2)))
				v92 = (v80-v82)*v70 + v71
			}
		} else {
			v69 = v14
			v70 = v17
			v71 = v5
			v78 = v69 << (uint(int32(2)) % 32)
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l3+v78)))
			v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+l2)))
			v92 = (v80-v82)*v70 + v71
		}
	}
	return v92
}
func F_CopyArrayEls(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v20 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	switch l6 - int32(99) {
	case 0:
		v58 = int32(1)
		goto L5
	case 1:
		goto L8
	default:
		goto L7
	case 6:
		goto L9
	case 16:
		goto L6
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = (v23<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v38 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v20
	v38 = l0 + v31<<(uint(int32(3))%32) + int32(16)
	goto L1
L5:
	;
	if l3 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v58 = int32(2)
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v58 = int32(8)
	goto L5
L9:
	;
	v58 = int32(4)
	goto L5
L10:
	;
	return
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l6
	F_errmsg_internal(m, int32(_a_F_CopyArrayEls_0), v18)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_CopyArrayEls_1), int32(322), int32(_a_F_CopyArrayEls_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	m.G0 = v18 + int32(16)
	return
L15:
	;
	v62 = int32(1)
	v65 = int32(0)
	v68 = v65
	v74 = v62
	v75 = v65
	v76 = v38
	v77 = l0 + v37
	goto L16
L16:
	;
	if l2 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if base.B2i32(v129 == int32(0))|base.B2i32(v127 == int32(1)) != 0 {
		goto L14
	} else {
		goto L35
	}
L18:
	;
	v134 = v68 + int32(1)
	if v134 != l3 {
		v68 = v134
		v74 = v127
		v75 = v128
		v76 = v129
		v77 = v130
		goto L16
	} else {
		goto L34
	}
L19:
	;
	v119 = v74 << (uint(int32(1)) % 32)
	if v119 != int32(256) {
		v127 = v119
		v128 = v114
		v129 = v76
		v130 = v115
		goto L18
	} else {
		goto L33
	}
L20:
	;
	v104 = l1 + v68<<(uint(int32(3))%32)
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	v106 = F_ArrayCastAndSet(m, v105, l4, l5, v58, v77)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L27
	}
L21:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+l2))))
	if v86 != int32(1) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v76 != 0 {
		v114 = v75
		v115 = v77
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	F_errmsg_internal(m, int32(_a_F_CopyArrayEls_3), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_CopyArrayEls_4), int32(989), int32(_a_F_CopyArrayEls_5))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	if l7&(l5^v62) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	F_pfree(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v111 = v74 | v75
	v112 = v77 + v106
	if v76 != 0 {
		v114 = v111
		v115 = v112
		goto L19
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v127 = v74
	v128 = v111
	v129 = int32(0)
	v130 = v112
	goto L18
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v114)
	v123 = int32(1)
	v127 = v123
	v128 = int32(0)
	v129 = v76 + v123
	v130 = v115
	goto L18
L34:
	;
	goto L17
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v128)
	goto L14
}
func F_array_agg_array_transfn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_get_fn_expr_argtype(m, v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		if v10 != 0 {
			v15 = v6 + int32(12)
			v16 = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 == v16 {
				v34 = int32(0)
				if v15 == v34 {
					v42 = v34
				} else {
					v37 = v34
					v38 = v16
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
				}
				v45 = v42
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				switch v20 - int32(435) {
				case 0:
					if v15 == int32(0) {
						v45 = int32(1)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
						v37 = v27
						v38 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				case 1:
					if v15 == int32(0) {
						v45 = int32(2)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+376))
						v37 = v32
						v38 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				default:
					v34 = int32(0)
					if v15 == v34 {
						v42 = v34
					} else {
						v37 = v34
						v38 = v16
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
					}
					v45 = v42
				}
			}
			if v45 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_array_agg_array_transfn_0), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_array_transfn_1), int32(954), int32(_a_F_array_agg_array_transfn_2))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v48 == int32(1) {
					v51 = int32(0)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					v54 = F_initArrayResultArr(m, v10, v51, v52, v51)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						v57 = v54
						v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						v61 = F_accumArrayResultArr(m, v57, v58, v59, v10, v60)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							m.G0 = v6 + int32(16)
							return base.I64_extend_i32_u(v61)
						}
					}
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v57 = v56
					v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					v61 = F_accumArrayResultArr(m, v57, v58, v59, v10, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int64(0)
					} else {
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v61)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_array_agg_array_transfn_3), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_array_transfn_1), int32(943), int32(_a_F_array_agg_array_transfn_2))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
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
}
func F_array_append_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	v3 = int64(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v5 != int32(472) {
		v22 = v3
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v10 == int32(0) {
			v22 = v3
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(8) {
				v22 = v3
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v16 != 0 {
					v22 = v3
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
					if v17 != v18 {
						v22 = v3
					} else {
						v22 = base.I64_extend_i32_u(v10)
					}
				}
			}
		}
	}
	return v22
}
func F_array_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	v18 = m.G0
	v20 = v18 - int32(128)
	m.G0 = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = F_DatumGetAnyArrayP(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = F_DatumGetAnyArrayP(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v32 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = int32(28)
	goto L6
L5:
	;
	v35 = int32(4)
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v39 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v42 = int32(28)
	goto L9
L8:
	;
	v42 = int32(4)
	goto L9
L9:
	;
	if v39 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28+v35)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23+v42)))
	if v32 == int32(-1) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v49 = v46
	goto L10
L12:
	;
	goto L13
L13:
	;
	v49 = v23 + int32(16)
	goto L10
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = F_ArrayGetNItemsSafe(m, v51, v49)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v57 = v54
	goto L14
L16:
	;
	goto L17
L17:
	;
	v57 = v28 + int32(16)
	goto L14
L18:
	;
	v61 = F_ArrayGetNItemsSafe(m, v50, v57)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v65 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L118
	}
L21:
	;
	v68 = int32(40)
	goto L23
L22:
	;
	v68 = int32(12)
	goto L23
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v23+v68)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v73 == int32(-1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v76 = int32(40)
	goto L26
L25:
	;
	v76 = int32(12)
	goto L26
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v28+v76)))
	if v70 == v78 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	if v81 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L114
	}
L30:
	;
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v92)+11)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+10)))
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92)+8)))
	v96 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+90)) = uint16(v96)
	v98 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+88)) = uint8(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v20)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v92 + int32(104)
	F_array_iter_setup(m, v20+int32(44), v23, v95, v94, v93)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L37
	}
L31:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v82 == v70 {
		v92 = v81
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v85 = F_lookup_type_cache(m, v70, int32(64))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+108))
	if v87 == int32(0) {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v85
	v92 = v85
	goto L30
L37:
	;
	F_array_iter_setup(m, v20+int32(16), v28, v95, v94, v93)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v114 = base.B2i32(v59 < v61)
	if v59 < v61 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v346 == int32(-1) {
		goto L106
	} else {
		goto L107
	}
L40:
	;
	v115 = v59
	goto L42
L41:
	;
	v115 = v61
	goto L42
L42:
	;
	if int32(0) < v115 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v122 = int32(0)
	goto L48
L44:
	;
	goto L45
L45:
	;
	if v59 != v61 {
		goto L64
	} else {
		goto L65
	}
L46:
	;
	if v194 != 0 {
		v330 = v194
		goto L39
	} else {
		goto L63
	}
L47:
	;
	v186 = int32(-1)
	if v148 != 0 {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v140 = F_array_iter_next(m, v20+int32(44), v20+int32(15), v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v194 = int32(0)
	goto L46
L50:
	;
	v146 = F_array_iter_next(m, v20+int32(16), v20+int32(14), v122)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if v148|v149&int32(1) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v183 = v122 + int32(1)
	if v183 != v115 {
		v122 = v183
		goto L48
	} else {
		goto L59
	}
L53:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+120)) = uint8(v155)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+112)) = v146
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+104)) = uint8(v155)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+96)) = v140
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = m.T0[v164].(func(*base.Module, int32) int64)(m, v20+int32(72))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v174 = int32(0)
	if base.B2i32(v148 == v174)|base.B2i32(v149&int32(1) == v174) != 0 {
		goto L47
	} else {
		goto L58
	}
L56:
	;
	v167 = base.I32_wrap_i64(v165)
	if v167 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v330 = v167>>(uint(int32(31))%32) | int32(1)
	goto L39
L58:
	;
	goto L52
L59:
	;
	goto L49
L60:
	;
	v191 = (v149 ^ v186) & int32(1)
	goto L62
L61:
	;
	v191 = v186
	goto L62
L62:
	;
	v194 = v191
	goto L46
L63:
	;
	goto L45
L64:
	;
	if v59 < v61 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if v51 == v50 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v215 = int32(-1)
	goto L69
L68:
	;
	v215 = int32(1)
	goto L69
L69:
	;
	v330 = v215
	goto L39
L70:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v272 == int32(-1) {
		goto L89
	} else {
		goto L90
	}
L71:
	;
	v227 = v217
	goto L79
L72:
	;
	v217 = int32(0)
	if v51 <= v217 {
		goto L70
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v51 < v50 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L71
L76:
	;
	v223 = int32(-1)
	goto L78
L77:
	;
	v223 = int32(1)
	goto L78
L78:
	;
	v330 = v223
	goto L39
L79:
	;
	v242 = v227 << (uint(int32(2)) % 32)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v49+v242)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242+v57)))
	if v244 == v246 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v244 < v246 {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v249 = v227 + int32(1)
	if v51 != v249 {
		v227 = v249
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	goto L70
L85:
	;
	v254 = int32(-1)
	goto L87
L86:
	;
	v254 = int32(1)
	goto L87
L87:
	;
	v330 = v254
	goto L39
L88:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v283 == int32(-1) {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v282 = v275
	goto L88
L90:
	;
	goto L91
L91:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v282 = v23 + v276<<(uint(int32(2))%32) + int32(16)
	goto L88
L92:
	;
	v294 = int32(0)
	if v51 <= v294 {
		v330 = v294
		goto L39
	} else {
		goto L96
	}
L93:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v293 = v286
	goto L92
L94:
	;
	goto L95
L95:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v293 = v28 + v287<<(uint(int32(2))%32) + int32(16)
	goto L92
L96:
	;
	v301 = int32(0)
	goto L97
L97:
	;
	v316 = v301 << (uint(int32(2)) % 32)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v282+v316)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v293+v316)))
	if v318 == v320 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v318 < v320 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v323 = v301 + int32(1)
	if v51 != v323 {
		v301 = v323
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L98
L102:
	;
	v330 = v294
	goto L39
L103:
	;
	v328 = int32(-1)
	goto L105
L104:
	;
	v328 = int32(1)
	goto L105
L105:
	;
	v330 = v328
	goto L39
L106:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v353 == int32(-1) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v23 == v349 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	F_pfree(m, v23)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	m.G0 = v20 + int32(128)
	return v330
L111:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v28 == v356 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	F_pfree(m, v28)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L110
L114:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_array_cmp_0), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_array_cmp_1), int32(4023), int32(_a_F_array_cmp_2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v387 = F_format_type_be(m, v70)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v387
	F_errmsg(m, int32(_a_F_array_cmp_3), v20)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_array_cmp_1), int32(4041), int32(_a_F_array_cmp_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_iter_next(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v126 int64
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 != 0 {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v11+l2<<(uint(int32(3))%32))))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v16 != 0 {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+l2))))
			v20 = v18
		} else {
			v20 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
		v126 = v15
		m.G0 = v9 + int32(16)
		return v126
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v22 == int32(0) {
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v34 == int32(1) {
				if base.I32_popcnt(v32) != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
						F_errmsg_internal(m, int32(_a_F_array_iter_next_0), v9)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_array_iter_next_1), int32(123), int32(_a_F_array_iter_next_2))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					switch base.I32_ctz(v32) {
					case 0:
						v41 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33))))
						v61 = v41
						if int32(0) < v32 {
							v99 = v33 + v32
						} else {
							if v32 == int32(-1) {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								if v67 == int32(1) {
									v71 = int32(18)
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
									if v73 == v71 {
										v76 = v71
									} else {
										v76 = int32(2)
									}
									if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v83 = int32(6)
									} else {
										v83 = v76
									}
									v99 = v33 + v83
								} else {
									v85 = int32(1)
									if v67&v85 != 0 {
										v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
									}
								}
							} else {
								v94 = F_strlen(m, v33)
								mBase = m.M
								v99 = v94 + v33 + int32(1)
							}
						}
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
						v110 = v61
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v113 = v111 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
						if v113 != int32(256) {
							v126 = v110
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v117 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v126 = v110
						}
						m.G0 = v9 + int32(16)
						return v126
					case 1:
						v42 = int64(*(*int16)(unsafe.Add(mBase, uint32(v33))))
						v61 = v42
						if int32(0) < v32 {
							v99 = v33 + v32
						} else {
							if v32 == int32(-1) {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								if v67 == int32(1) {
									v71 = int32(18)
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
									if v73 == v71 {
										v76 = v71
									} else {
										v76 = int32(2)
									}
									if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v83 = int32(6)
									} else {
										v83 = v76
									}
									v99 = v33 + v83
								} else {
									v85 = int32(1)
									if v67&v85 != 0 {
										v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
									}
								}
							} else {
								v94 = F_strlen(m, v33)
								mBase = m.M
								v99 = v94 + v33 + int32(1)
							}
						}
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
						v110 = v61
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v113 = v111 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
						if v113 != int32(256) {
							v126 = v110
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v117 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v126 = v110
						}
						m.G0 = v9 + int32(16)
						return v126
					case 2:
						v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33))))
						v61 = v43
						if int32(0) < v32 {
							v99 = v33 + v32
						} else {
							if v32 == int32(-1) {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								if v67 == int32(1) {
									v71 = int32(18)
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
									if v73 == v71 {
										v76 = v71
									} else {
										v76 = int32(2)
									}
									if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v83 = int32(6)
									} else {
										v83 = v76
									}
									v99 = v33 + v83
								} else {
									v85 = int32(1)
									if v67&v85 != 0 {
										v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
									}
								}
							} else {
								v94 = F_strlen(m, v33)
								mBase = m.M
								v99 = v94 + v33 + int32(1)
							}
						}
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
						v110 = v61
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v113 = v111 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
						if v113 != int32(256) {
							v126 = v110
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v117 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v126 = v110
						}
						m.G0 = v9 + int32(16)
						return v126
					case 3:
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
						v61 = v44
						if int32(0) < v32 {
							v99 = v33 + v32
						} else {
							if v32 == int32(-1) {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								if v67 == int32(1) {
									v71 = int32(18)
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
									if v73 == v71 {
										v76 = v71
									} else {
										v76 = int32(2)
									}
									if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v83 = int32(6)
									} else {
										v83 = v76
									}
									v99 = v33 + v83
								} else {
									v85 = int32(1)
									if v67&v85 != 0 {
										v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
									}
								}
							} else {
								v94 = F_strlen(m, v33)
								mBase = m.M
								v99 = v94 + v33 + int32(1)
							}
						}
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
						v110 = v61
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v113 = v111 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
						if v113 != int32(256) {
							v126 = v110
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v117 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v126 = v110
						}
						m.G0 = v9 + int32(16)
						return v126
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
							F_errmsg_internal(m, int32(_a_F_array_iter_next_0), v9)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_array_iter_next_1), int32(123), int32(_a_F_array_iter_next_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
				v61 = base.I64_extend_i32_u(v33)
				if int32(0) < v32 {
					v99 = v33 + v32
				} else {
					if v32 == int32(-1) {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
						if v67 == int32(1) {
							v71 = int32(18)
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
							if v73 == v71 {
								v76 = v71
							} else {
								v76 = int32(2)
							}
							if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v83 = int32(6)
							} else {
								v83 = v76
							}
							v99 = v33 + v83
						} else {
							v85 = int32(1)
							if v67&v85 != 0 {
								v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
							} else {
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
							}
						}
					} else {
						v94 = F_strlen(m, v33)
						mBase = m.M
						v99 = v94 + v33 + int32(1)
					}
				}
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
				v110 = v61
				v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v113 = v111 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
				if v113 != int32(256) {
					v126 = v110
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v117 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
					v126 = v110
				}
				m.G0 = v9 + int32(16)
				return v126
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
			if v25&v26 != 0 {
				v30 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v34 == int32(1) {
					if base.I32_popcnt(v32) != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
							F_errmsg_internal(m, int32(_a_F_array_iter_next_0), v9)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_array_iter_next_1), int32(123), int32(_a_F_array_iter_next_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						switch base.I32_ctz(v32) {
						case 0:
							v41 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33))))
							v61 = v41
							if int32(0) < v32 {
								v99 = v33 + v32
							} else {
								if v32 == int32(-1) {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
									if v67 == int32(1) {
										v71 = int32(18)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
										if v73 == v71 {
											v76 = v71
										} else {
											v76 = int32(2)
										}
										if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v83 = int32(6)
										} else {
											v83 = v76
										}
										v99 = v33 + v83
									} else {
										v85 = int32(1)
										if v67&v85 != 0 {
											v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
										}
									}
								} else {
									v94 = F_strlen(m, v33)
									mBase = m.M
									v99 = v94 + v33 + int32(1)
								}
							}
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
							v110 = v61
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v113 = v111 << (uint(int32(1)) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
							if v113 != int32(256) {
								v126 = v110
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v117 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
								v126 = v110
							}
							m.G0 = v9 + int32(16)
							return v126
						case 1:
							v42 = int64(*(*int16)(unsafe.Add(mBase, uint32(v33))))
							v61 = v42
							if int32(0) < v32 {
								v99 = v33 + v32
							} else {
								if v32 == int32(-1) {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
									if v67 == int32(1) {
										v71 = int32(18)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
										if v73 == v71 {
											v76 = v71
										} else {
											v76 = int32(2)
										}
										if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v83 = int32(6)
										} else {
											v83 = v76
										}
										v99 = v33 + v83
									} else {
										v85 = int32(1)
										if v67&v85 != 0 {
											v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
										}
									}
								} else {
									v94 = F_strlen(m, v33)
									mBase = m.M
									v99 = v94 + v33 + int32(1)
								}
							}
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
							v110 = v61
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v113 = v111 << (uint(int32(1)) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
							if v113 != int32(256) {
								v126 = v110
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v117 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
								v126 = v110
							}
							m.G0 = v9 + int32(16)
							return v126
						case 2:
							v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33))))
							v61 = v43
							if int32(0) < v32 {
								v99 = v33 + v32
							} else {
								if v32 == int32(-1) {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
									if v67 == int32(1) {
										v71 = int32(18)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
										if v73 == v71 {
											v76 = v71
										} else {
											v76 = int32(2)
										}
										if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v83 = int32(6)
										} else {
											v83 = v76
										}
										v99 = v33 + v83
									} else {
										v85 = int32(1)
										if v67&v85 != 0 {
											v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
										}
									}
								} else {
									v94 = F_strlen(m, v33)
									mBase = m.M
									v99 = v94 + v33 + int32(1)
								}
							}
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
							v110 = v61
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v113 = v111 << (uint(int32(1)) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
							if v113 != int32(256) {
								v126 = v110
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v117 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
								v126 = v110
							}
							m.G0 = v9 + int32(16)
							return v126
						case 3:
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
							v61 = v44
							if int32(0) < v32 {
								v99 = v33 + v32
							} else {
								if v32 == int32(-1) {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
									if v67 == int32(1) {
										v71 = int32(18)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
										if v73 == v71 {
											v76 = v71
										} else {
											v76 = int32(2)
										}
										if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v83 = int32(6)
										} else {
											v83 = v76
										}
										v99 = v33 + v83
									} else {
										v85 = int32(1)
										if v67&v85 != 0 {
											v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
										}
									}
								} else {
									v94 = F_strlen(m, v33)
									mBase = m.M
									v99 = v94 + v33 + int32(1)
								}
							}
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
							v110 = v61
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v113 = v111 << (uint(int32(1)) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
							if v113 != int32(256) {
								v126 = v110
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v117 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
								v126 = v110
							}
							m.G0 = v9 + int32(16)
							return v126
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
								F_errmsg_internal(m, int32(_a_F_array_iter_next_0), v9)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_array_iter_next_1), int32(123), int32(_a_F_array_iter_next_2))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
					v61 = base.I64_extend_i32_u(v33)
					if int32(0) < v32 {
						v99 = v33 + v32
					} else {
						if v32 == int32(-1) {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
							if v67 == int32(1) {
								v71 = int32(18)
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
								if v73 == v71 {
									v76 = v71
								} else {
									v76 = int32(2)
								}
								if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v83 = int32(6)
								} else {
									v83 = v76
								}
								v99 = v33 + v83
							} else {
								v85 = int32(1)
								if v67&v85 != 0 {
									v99 = v33 + int32(base.Ui32(v67)>>(uint(v85)%32))
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									v99 = v33 + int32(base.Ui32(v90)>>(uint(int32(2))%32))
								}
							}
						} else {
							v94 = F_strlen(m, v33)
							mBase = m.M
							v99 = v94 + v33 + int32(1)
						}
					}
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = (v99 + v100 - int32(1)) & (int32(0) - v100)
					v110 = v61
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v113 = v111 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
					if v113 != int32(256) {
						v126 = v110
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v117 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v126 = v110
					}
					m.G0 = v9 + int32(16)
					return v126
				}
			} else {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v28)
				v110 = int64(0)
				v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v113 = v111 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
				if v113 != int32(256) {
					v126 = v110
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v117 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117 + int32(1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
					v126 = v110
				}
				m.G0 = v9 + int32(16)
				return v126
			}
		}
	}
}
func F_array_iter_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	v4 = l3
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13 == int32(-1) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		if v16 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v16
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18
			v74 = v19
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			if v26 != 0 {
				v34 = v26
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34 + v25
			v37 = int32(0)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
			if v39 == v37 {
				v74 = v37
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v74 = v38 + v42<<(uint(int32(3))%32) + int32(16)
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v50 != 0 {
			v58 = v50
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v58 + l1
		v61 = int32(0)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v62 == v61 {
			v74 = v61
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v74 = l1 + v65<<(uint(int32(3))%32) + int32(16)
		}
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l2
	v77 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v74
	switch l4 - int32(99) {
	case 0:
		v99 = v77
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v99)
		m.G0 = v11 + int32(16)
		return
	case 1:
		v99 = int32(8)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v99)
		m.G0 = v11 + int32(16)
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = l4
			F_errmsg_internal(m, int32(_a_F_array_iter_setup_0), v11)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_array_iter_setup_1), int32(322), int32(_a_F_array_iter_setup_2))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 6:
		v99 = int32(4)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v99)
		m.G0 = v11 + int32(16)
		return
	case 16:
		v99 = int32(2)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v99)
		m.G0 = v11 + int32(16)
		return
	}
}
func F_array_iterate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
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
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 <= v17 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return base.B2i32(v17 < v18)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v17 + int32(1)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(0) < v126 {
		goto L40
	} else {
		goto L41
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v43)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v45 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v30 = base.I32_div_s(v17, int32(8))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v30))))
	if int32(base.Ui32(v32)>>(uint(v17&int32(7))%32))&int32(1) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v38)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	goto L1
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v74
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v76 {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.I32_popcnt(v48) != int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v74 = base.I64_extend_i32_u(v42)
	goto L9
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	switch base.I32_ctz(v48) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	default:
		goto L13
	}
L15:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v74 = v56
	goto L9
L16:
	;
	v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42))))
	v74 = v55
	goto L9
L17:
	;
	v54 = int64(*(*int16)(unsafe.Add(mBase, uint32(v42))))
	v74 = v54
	goto L9
L18:
	;
	v53 = int64(*(*int8)(unsafe.Add(mBase, uint32(v42))))
	v74 = v53
	goto L9
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v48
	F_errmsg_internal(m, int32(_a_F_array_iterate_0), v15)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_array_iterate_1), int32(123), int32(_a_F_array_iterate_2))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = (v114 + v115 - int32(1)) & (int32(0) - v115)
	goto L1
L24:
	;
	v114 = v76 + v42
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v76 == int32(-1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v82 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v109 = F_strlen(m, v42)
	mBase = m.M
	v114 = v109 + v42 + int32(1)
	goto L23
L30:
	;
	v86 = int32(18)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v88 == v86 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v100 = int32(1)
	if v82&v100 != 0 {
		v114 = v42 + int32(base.Ui32(v82)>>(uint(v100)%32))
		goto L23
	} else {
		goto L39
	}
L33:
	;
	v91 = v86
	goto L35
L34:
	;
	v91 = int32(2)
	goto L35
L35:
	;
	if base.Ui32((v88-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v98 = int32(6)
	goto L38
L37:
	;
	v98 = v91
	goto L38
L38:
	;
	v114 = v42 + v98
	goto L23
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v114 = v42 + int32(base.Ui32(v105)>>(uint(int32(2))%32))
	goto L23
L40:
	;
	v133 = v123
	v135 = int32(0)
	goto L43
L41:
	;
	v259 = v123
	v261 = v20
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v259
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v273 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v275 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+15)))
	v276 = F_construct_md_array(m, v125, v124, v261, v269, v270, v272, v273, v274, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L19
	} else {
		goto L80
	}
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v142 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v146 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v259 = v249
	v261 = v255
	goto L42
L45:
	;
	v252 = v135 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v252 < v253 {
		v133 = v249
		v135 = v252
		goto L43
	} else {
		goto L79
	}
L46:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v124))) = uint8(v167)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v172 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v150 = base.I32_div_s(v142, int32(8))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v150))))
	if int32(base.Ui32(v152)>>(uint(v142&int32(7))%32))&int32(1) != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v124))) = uint8(v159)
	*(*int64)(unsafe.Add(mBase, uint32(v125+v135<<(uint(int32(3))%32)))) = int64(0)
	v249 = v133
	goto L45
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v125+v135<<(uint(int32(3))%32)))) = v201
	v203 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v203 {
		goto L63
	} else {
		goto L64
	}
L50:
	;
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.I32_popcnt(v175) != int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v201 = base.I64_extend_i32_u(v133)
	goto L49
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L19
	} else {
		goto L59
	}
L54:
	;
	switch base.I32_ctz(v175) {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L56
	case 3:
		goto L55
	default:
		goto L53
	}
L55:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v133)))
	v201 = v183
	goto L49
L56:
	;
	v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v133))))
	v201 = v182
	goto L49
L57:
	;
	v181 = int64(*(*int16)(unsafe.Add(mBase, uint32(v133))))
	v201 = v181
	goto L49
L58:
	;
	v180 = int64(*(*int8)(unsafe.Add(mBase, uint32(v133))))
	v201 = v180
	goto L49
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v175
	F_errmsg_internal(m, int32(_a_F_array_iterate_0), v15+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_array_iterate_1), int32(123), int32(_a_F_array_iterate_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v249 = (v241 + v242 - int32(1)) & (int32(0) - v242)
	goto L45
L63:
	;
	v241 = v133 + v203
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v203 == int32(-1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v209 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v236 = F_strlen(m, v133)
	mBase = m.M
	v241 = v236 + v133 + int32(1)
	goto L62
L69:
	;
	v213 = int32(18)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v215 == v213 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v227 = int32(1)
	if v209&v227 != 0 {
		v241 = v133 + int32(base.Ui32(v209)>>(uint(v227)%32))
		goto L62
	} else {
		goto L78
	}
L72:
	;
	v218 = v213
	goto L74
L73:
	;
	v218 = int32(2)
	goto L74
L74:
	;
	if base.Ui32((v215-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v225 = int32(6)
	goto L77
L76:
	;
	v225 = v218
	goto L77
L77:
	;
	v241 = v133 + v225
	goto L62
L78:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v241 = v133 + int32(base.Ui32(v232)>>(uint(int32(2))%32))
	goto L62
L79:
	;
	goto L44
L80:
	;
	v278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v278)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_u(v276)
	goto L1
}
func F_array_length(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int64
	_ = v49
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v13 == int32(-1) {
			v16 = int32(28)
		} else {
			v16 = int32(4)
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v7+v16)))
		if base.Ui32(v18-int32(7)) <= base.Ui32(int32(-7)) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			return int64(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int64(0)
			} else {
				if v13 == int32(-1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
					v43 = v40
				} else {
					v43 = v7 + int32(16)
				}
				v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43+v27<<(uint(int32(2))%32)-int32(4)))))
				return v49
			}
		}
	}
}
func F_array_positions(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v221 int64
	_ = v221
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v16 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v221
L2:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
	v221 = int64(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int64(0)
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 < int32(2) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_array_positions[0]))
	v34 = F_initArrayResult(m, int32(23), v32, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L65
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v36 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_array_positions[0]))
	v41 = F_makeArrayResult(m, v34, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v43 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v221 = v41
	goto L1
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v23+v55<<(uint(int32(2))%32))+16))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	if v63 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v46 = F_array_contains_nulls(m, v23)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v55 = v36
	v56 = v54
	goto L15
L19:
	;
	if v46 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v55 = v48
	v56 = int64(0)
	goto L15
L21:
	;
	goto L22
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_array_positions[0]))
	v52 = F_makeArrayResult(m, v34, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v221 = v52
	goto L1
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L60
	}
L25:
	;
	v105 = F_array_create_iterator(m, v23, int32(0), v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L36
	}
L26:
	;
	F_get_typlenbyvalalign(m, v61, v79+int32(4), v79+int32(6), v79+int32(7))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L32
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v68 = F_MemoryContextAlloc(m, v66, int32(48))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v77 == v61 {
		v102 = v63
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v68
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v61 ^ int32(-1)
	v79 = v73
	goto L26
L31:
	;
	v79 = v63
	goto L26
L32:
	;
	v89 = F_lookup_type_cache(m, v61, int32(32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+80))
	if v91 == int32(0) {
		goto L24
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v61
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+80))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	F_fmgr_info_cxt(m, v95, v79+int32(20), v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v102 = v79
	goto L25
L36:
	;
	v111 = F_array_iterate(m, v105, v14+int32(8), v14+int32(7))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	if v111 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = v60 - int32(1)
	v123 = v34
	goto L41
L39:
	;
	v162 = v34
	goto L40
L40:
	;
	F_array_free_iterator(m, v105)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L54
	}
L41:
	;
	v128 = int32(1)
	v129 = v118 + v128
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+7)))
	if (v130|v43)&v128 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v162 = v149
	goto L40
L43:
	;
	v154 = F_array_iterate(m, v105, v14+int32(8), v14+int32(7))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L52
	}
L44:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_array_positions[0]))
	v147 = F_accumArrayResult(m, v123, base.I64_extend_i32_s(v129), int32(0), int32(23), v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L51
	}
L45:
	;
	if v130&v43 == int32(0) {
		v149 = v123
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v138 = F_FunctionCall2Coll(m, v102+int32(20), v21, v56, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L49
	}
L48:
	;
	goto L44
L49:
	;
	if v138 == int64(0) {
		v149 = v123
		goto L43
	} else {
		goto L50
	}
L50:
	;
	goto L44
L51:
	;
	v149 = v147
	goto L43
L52:
	;
	if v154 != 0 {
		v118 = v129
		v123 = v149
		goto L41
	} else {
		goto L53
	}
L53:
	;
	goto L42
L54:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v169 != v23 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_pfree(m, v23)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_array_positions[0]))
	v175 = F_makeArrayResult(m, v162, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v221 = v175
	goto L1
L60:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v184 = F_format_type_be(m, v61)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v184
	F_errmsg(m, int32(_a_F_array_positions_0), v14)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_array_positions_1), int32(1563), int32(_a_F_array_positions_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_array_positions_3), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_array_positions_1), int32(1511), int32(_a_F_array_positions_2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_replace(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
		return int64(0)
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v24 = F_array_replace_internal(m, v18, v13, v14, v15, v16, int32(0), v23, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v24)
			}
		}
	}
}
func F_array_replace_internal(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int64
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	v5 = l4
	v9 = int32(0)
	v34 = m.G0
	v36 = v34 - int32(112)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = l0 + int32(16)
	v42 = F_ArrayGetNItemsSafe(m, v39, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L159
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L154
	}
L3:
	;
	m.G0 = v36 + int32(112)
	return v479
L4:
	;
	return int32(0)
L5:
	;
	if v42 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v479 = l0
	goto L3
L7:
	;
	goto L8
L8:
	;
	if int32(2) <= v39 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v51 = l5
	goto L11
L10:
	;
	v51 = int32(0)
	goto L11
L11:
	;
	if v51 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+10)))
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+8)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+11)))
	v68 = base.I32_extend8_s(v67)
	switch v67 - int32(99) {
	case 0:
		v90 = int32(1)
		goto L20
	case 1:
		goto L23
	default:
		goto L22
	case 6:
		goto L24
	case 16:
		goto L21
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54 == v38 {
		v64 = v53
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v57 = F_lookup_type_cache(m, v38, int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v59 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v57
	v64 = v57
	goto L13
L20:
	;
	v92 = base.B2i32(v66 != int32(-1))
	if v66 != int32(-1) {
		v104 = l1
		v105 = l3
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v90 = int32(2)
	goto L20
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	v90 = int32(8)
	goto L20
L24:
	;
	v90 = int32(4)
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v68
	F_errmsg_internal(m, int32(_a_F_array_replace_internal_0), v36+int32(16))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_1), int32(322), int32(_a_F_array_replace_internal_2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	v106 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+74)) = uint16(v106)
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+72)) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = l6
	*(*int64)(unsafe.Add(mBase, uint32(v36)+60)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+56)) = v64 + int32(76)
	v118 = F_palloc(m, v42<<(uint(int32(3))%32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L36
	}
L29:
	;
	if l2 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = F_pg_detoast_datum(m, base.I32_wrap_i64(l1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	v99 = l1
	goto L32
L32:
	;
	if v5 != 0 {
		v104 = v99
		v105 = l3
		goto L28
	} else {
		goto L34
	}
L33:
	;
	v99 = base.I64_extend_i32_u(v96)
	goto L32
L34:
	;
	v101 = F_pg_detoast_datum(m, base.I32_wrap_i64(l3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v104 = v99
	v105 = base.I64_extend_i32_u(v101)
	goto L28
L36:
	;
	v120 = F_palloc(m, v42)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = v122 << (uint(int32(3)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v127 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v128 = v41 + v124
	goto L40
L39:
	;
	v128 = int32(0)
	goto L40
L40:
	;
	v130 = int32(0) - v90
	v132 = v90 - int32(1)
	if v127 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = v127
	goto L43
L42:
	;
	v137 = (v124 + int32(23)) & int32(-8)
	goto L43
L43:
	;
	v141 = int32(1)
	v151 = int32(0)
	v152 = v141
	v154 = v128
	v156 = l0 + v137
	v163 = v9
	v171 = v9
	v172 = v9
	v175 = v9
	goto L44
L44:
	;
	if v154 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	if v378 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L46:
	;
	v383 = int32(1)
	v385 = v152 << (uint(v383) % 32)
	v387 = base.B2i32(v385 == int32(256))
	if v385 == int32(256) {
		goto L116
	} else {
		goto L117
	}
L47:
	;
	v374 = v151 + int32(1)
	v376 = v365
	v378 = v367
	v379 = v368
	v380 = v369
	goto L46
L48:
	;
	if int32(0) < v66 {
		v338 = v66
		goto L94
	} else {
		goto L95
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118+v151<<(uint(int32(3))%32)))) = v288
	v295 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v120))) = uint8(v295)
	v298 = v285
	v300 = v287
	goto L48
L50:
	;
	v280 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v120))) = uint8(v280)
	v365 = v156
	v367 = v163
	v368 = v280
	v369 = v172
	goto L47
L51:
	;
	if v65&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v152&v180 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if l2 == int32(0) {
		goto L50
	} else {
		goto L54
	}
L54:
	;
	if l5 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v374 = v151
	v376 = v156
	v378 = int32(1)
	v379 = v171
	v380 = v172
	goto L46
L56:
	;
	goto L57
L57:
	;
	if v5 != 0 {
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v285 = v156
	v287 = int32(1)
	v288 = v105
	goto L49
L59:
	;
	if int32(0) < v66 {
		v248 = v156 + v66
		goto L72
	} else {
		goto L73
	}
L60:
	;
	if base.I32_popcnt(v66) != v141 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v208 = base.I64_extend_i32_u(v156)
	goto L59
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L69
	}
L64:
	;
	switch base.I32_ctz(v66) {
	case 0:
		goto L68
	case 1:
		goto L67
	case 2:
		goto L66
	case 3:
		goto L65
	default:
		goto L63
	}
L65:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
	v208 = v191
	goto L59
L66:
	;
	v190 = int64(*(*int32)(unsafe.Add(mBase, uint32(v156))))
	v208 = v190
	goto L59
L67:
	;
	v189 = int64(*(*int16)(unsafe.Add(mBase, uint32(v156))))
	v208 = v189
	goto L59
L68:
	;
	v188 = int64(*(*int8)(unsafe.Add(mBase, uint32(v156))))
	v208 = v188
	goto L59
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v66
	F_errmsg_internal(m, int32(_a_F_array_replace_internal_3), v36+int32(32))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_1), int32(123), int32(_a_F_array_replace_internal_4))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v250 = (v248 + v132) & v130
	if l2 != 0 {
		v285 = v250
		v287 = v163
		v288 = v208
		goto L49
	} else {
		goto L87
	}
L73:
	;
	v212 = base.I32_wrap_i64(v208)
	if v92 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v215 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v242 = F_strlen(m, v212)
	mBase = m.M
	v248 = v242 + v156 + int32(1)
	goto L72
L77:
	;
	v219 = int32(18)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
	if v221 == v219 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v233 = int32(1)
	if v215&v233 != 0 {
		v248 = v156 + int32(base.Ui32(v215)>>(uint(v233)%32))
		goto L72
	} else {
		goto L86
	}
L80:
	;
	v224 = v219
	goto L82
L81:
	;
	v224 = int32(2)
	goto L82
L82:
	;
	if base.Ui32((v221-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v231 = int32(6)
	goto L85
L84:
	;
	v231 = v224
	goto L85
L85:
	;
	v248 = v156 + v231
	goto L72
L86:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v248 = v156 + int32(base.Ui32(v238)>>(uint(int32(2))%32))
	goto L72
L87:
	;
	v251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+104)) = uint8(v251)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+96)) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+88)) = uint8(v251)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+80)) = v208
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+72)) = uint8(v251)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v263 = m.T0[v262].(func(*base.Module, int32) int64)(m, v36+int32(56))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+72)))
	if v265|base.B2i32(v263 == int64(0)) != 0 {
		v285 = v250
		v287 = v163
		v288 = v208
		goto L49
	} else {
		goto L89
	}
L89:
	;
	if l5 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v374 = v151
	v376 = v250
	v378 = int32(1)
	v379 = v171
	v380 = v172
	goto L46
L91:
	;
	goto L92
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118+v151<<(uint(int32(3))%32)))) = v105
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v120))) = uint8(v5)
	v276 = int32(1)
	if v5 == int32(0) {
		v298 = v250
		v300 = v276
		goto L48
	} else {
		goto L93
	}
L93:
	;
	v365 = v250
	v367 = v276
	v368 = int32(1)
	v369 = v172
	goto L47
L94:
	;
	v342 = (v132 + v172 + v338) & v130
	if base.Ui32(v342) < base.Ui32(int32(1073741824)) {
		v365 = v298
		v367 = v300
		v368 = v171
		v369 = v342
		goto L47
	} else {
		goto L111
	}
L95:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v118+v151<<(uint(int32(3))%32))))
	if v92 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	if v311 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v335 = F_strlen(m, v308)
	mBase = m.M
	v338 = v335 + int32(1)
	goto L94
L99:
	;
	v315 = int32(18)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if v317 == v315 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	if v311&int32(1) != 0 {
		goto L108
	} else {
		goto L109
	}
L102:
	;
	v320 = v315
	goto L104
L103:
	;
	v320 = int32(2)
	goto L104
L104:
	;
	if base.Ui32((v317-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v327 = int32(6)
	goto L107
L106:
	;
	v327 = v320
	goto L107
L107:
	;
	v338 = v327
	goto L94
L108:
	;
	v338 = int32(base.Ui32(v311) >> (uint(int32(1)) % 32))
	goto L94
L109:
	;
	goto L110
L110:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v338 = int32(base.Ui32(v332) >> (uint(int32(2)) % 32))
	goto L94
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_replace_internal_5), v36+int32(48))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_6), int32(_a_F_array_replace_internal_7), int32(_a_F_array_replace_internal_8))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
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
	v388 = v383
	goto L118
L117:
	;
	v388 = v385
	goto L118
L118:
	;
	if v154 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v389 = v388
	goto L121
L120:
	;
	v389 = v152
	goto L121
L121:
	;
	if v154 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v392 = v387 + v154
	goto L124
L123:
	;
	v392 = int32(0)
	goto L124
L124:
	;
	v394 = v175 + int32(1)
	if v394 != v42 {
		v151 = v374
		v152 = v389
		v154 = v392
		v156 = v376
		v163 = v378
		v171 = v379
		v172 = v380
		v175 = v394
		goto L44
	} else {
		goto L125
	}
L125:
	;
	goto L45
L126:
	;
	F_pfree(m, v118)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if v374 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	F_pfree(m, v120)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v479 = l0
	goto L3
L131:
	;
	F_pfree(m, v118)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v417 = v39 << (uint(int32(3)) % 32)
	if v379 != 0 {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	F_pfree(m, v120)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v409 = F_palloc0(m, int32(16))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v409)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v409))) = int64(64)
	v479 = v409
	goto L3
L137:
	;
	v434 = v433 + v380
	v435 = F_palloc0(m, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L141
	}
L138:
	;
	v421 = base.I32_div_s(v374+int32(7), int32(8))
	v426 = (v417 + v421 + int32(23)) & int32(-8)
	v432 = v426
	v433 = v426
	goto L137
L139:
	;
	goto L140
L140:
	;
	v432 = int32(0)
	v433 = (v417 + int32(23)) & int32(-8)
	goto L137
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v435)+8)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v435)+4)) = v39
	v440 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v434 << (uint(v440) % 32)
	v444 = v435 + int32(16)
	v446 = v39 << (uint(v440) % 32)
	v447 = int32(0)
	v448 = base.B2i32(v446 == v447)
	if v448 == v447 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	base.MemoryCopy(m, v444, v41, v446)
	goto L144
L143:
	;
	goto L144
L144:
	;
	if v448 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	base.MemoryCopy(m, v444+v446, v41+v455<<(uint(int32(2))%32), v446)
	goto L147
L146:
	;
	goto L147
L147:
	;
	if l5 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = v374
	goto L150
L149:
	;
	goto L150
L150:
	;
	F_CopyArrayEls(m, v435, v118, v120, v374, v66, v65&int32(1), v68, int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	F_pfree(m, v118)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_pfree(m, v120)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v479 = v435
	goto L3
L154:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v514 = F_format_type_be(m, v38)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v514
	F_errmsg(m, int32(_a_F_array_replace_internal_9), v36)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_6), int32(_a_F_array_replace_internal_10), int32(_a_F_array_replace_internal_8))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(_a_F_array_replace_internal_11), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_6), int32(_a_F_array_replace_internal_12), int32(_a_F_array_replace_internal_8))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_seek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	switch l5 - int32(99) {
	case 0:
		v36 = int32(1)
		v37 = int32(0)
		if l2|base.B2i32(l4 <= v37) == v37 {
			v191 = l0 + (l4+v36-int32(1))&(int32(0)-v36)*l3
		} else {
			if l2 == int32(0) {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v54 = int32(0)
					v61 = l0
					v63 = v54
					for {
						if base.B2i32(l4 <= v54) == int32(0) {
							v108 = v61 + l4
						} else {
							if l4 == int32(-1) {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								if v76 == int32(1) {
									v80 = int32(18)
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
									if v82 == v80 {
										v85 = v80
									} else {
										v85 = int32(2)
									}
									if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v92 = int32(6)
									} else {
										v92 = v85
									}
									v108 = v61 + v92
								} else {
									v94 = int32(1)
									if v76&v94 != 0 {
										v108 = v61 + int32(base.Ui32(v76)>>(uint(v94)%32))
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v108 = v61 + int32(base.Ui32(v99)>>(uint(int32(2))%32))
									}
								}
							} else {
								v103 = F_strlen(m, v61)
								mBase = m.M
								v108 = v103 + v61 + int32(1)
							}
						}
						v110 = (v108 + (v36 - int32(1))) & (v54 - v36)
						v112 = v63 + int32(1)
						if v112 != l3 {
							v61 = v110
							v63 = v112
							continue
						} else {
							break
						}
						break
					}
					v191 = v110
				}
			} else {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v118 = int32(1)
					v121 = base.I32_div_s(l1, int32(8))
					v127 = l0
					v129 = l2 + v121
					v132 = v118 << (uint(l1&int32(7)) % 32)
					v136 = int32(0)
					for {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
						if v132&v137 != 0 {
							if int32(0) < l4 {
								v176 = v127 + l4
							} else {
								if l4 == int32(-1) {
									v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
									if v144 == int32(1) {
										v148 = int32(18)
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
										if v150 == v148 {
											v153 = v148
										} else {
											v153 = int32(2)
										}
										if base.Ui32((v150-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v160 = int32(6)
										} else {
											v160 = v153
										}
										v176 = v127 + v160
									} else {
										v162 = int32(1)
										if v144&v162 != 0 {
											v176 = v127 + int32(base.Ui32(v144)>>(uint(v162)%32))
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
											v176 = v127 + int32(base.Ui32(v167)>>(uint(int32(2))%32))
										}
									}
								} else {
									v171 = F_strlen(m, v127)
									mBase = m.M
									v176 = v171 + v127 + int32(1)
								}
							}
							v179 = (v176 + (v36 - v118)) & (int32(0) - v36)
						} else {
							v179 = v127
						}
						v181 = int32(1)
						v183 = v132 << (uint(v181) % 32)
						v185 = base.B2i32(v183 == int32(256))
						if v183 == int32(256) {
							v186 = v181
						} else {
							v186 = v183
						}
						v189 = v136 + int32(1)
						if v189 != l3 {
							v127 = v179
							v129 = v185 + v129
							v132 = v186
							v136 = v189
							continue
						} else {
							break
						}
						break
					}
					v191 = v179
				}
			}
		}
		m.G0 = v13 + int32(16)
		return v191
	case 1:
		v36 = int32(8)
		v37 = int32(0)
		if l2|base.B2i32(l4 <= v37) == v37 {
			v191 = l0 + (l4+v36-int32(1))&(int32(0)-v36)*l3
		} else {
			if l2 == int32(0) {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v54 = int32(0)
					v61 = l0
					v63 = v54
					for {
						if base.B2i32(l4 <= v54) == int32(0) {
							v108 = v61 + l4
						} else {
							if l4 == int32(-1) {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								if v76 == int32(1) {
									v80 = int32(18)
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
									if v82 == v80 {
										v85 = v80
									} else {
										v85 = int32(2)
									}
									if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v92 = int32(6)
									} else {
										v92 = v85
									}
									v108 = v61 + v92
								} else {
									v94 = int32(1)
									if v76&v94 != 0 {
										v108 = v61 + int32(base.Ui32(v76)>>(uint(v94)%32))
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v108 = v61 + int32(base.Ui32(v99)>>(uint(int32(2))%32))
									}
								}
							} else {
								v103 = F_strlen(m, v61)
								mBase = m.M
								v108 = v103 + v61 + int32(1)
							}
						}
						v110 = (v108 + (v36 - int32(1))) & (v54 - v36)
						v112 = v63 + int32(1)
						if v112 != l3 {
							v61 = v110
							v63 = v112
							continue
						} else {
							break
						}
						break
					}
					v191 = v110
				}
			} else {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v118 = int32(1)
					v121 = base.I32_div_s(l1, int32(8))
					v127 = l0
					v129 = l2 + v121
					v132 = v118 << (uint(l1&int32(7)) % 32)
					v136 = int32(0)
					for {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
						if v132&v137 != 0 {
							if int32(0) < l4 {
								v176 = v127 + l4
							} else {
								if l4 == int32(-1) {
									v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
									if v144 == int32(1) {
										v148 = int32(18)
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
										if v150 == v148 {
											v153 = v148
										} else {
											v153 = int32(2)
										}
										if base.Ui32((v150-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v160 = int32(6)
										} else {
											v160 = v153
										}
										v176 = v127 + v160
									} else {
										v162 = int32(1)
										if v144&v162 != 0 {
											v176 = v127 + int32(base.Ui32(v144)>>(uint(v162)%32))
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
											v176 = v127 + int32(base.Ui32(v167)>>(uint(int32(2))%32))
										}
									}
								} else {
									v171 = F_strlen(m, v127)
									mBase = m.M
									v176 = v171 + v127 + int32(1)
								}
							}
							v179 = (v176 + (v36 - v118)) & (int32(0) - v36)
						} else {
							v179 = v127
						}
						v181 = int32(1)
						v183 = v132 << (uint(v181) % 32)
						v185 = base.B2i32(v183 == int32(256))
						if v183 == int32(256) {
							v186 = v181
						} else {
							v186 = v183
						}
						v189 = v136 + int32(1)
						if v189 != l3 {
							v127 = v179
							v129 = v185 + v129
							v132 = v186
							v136 = v189
							continue
						} else {
							break
						}
						break
					}
					v191 = v179
				}
			}
		}
		m.G0 = v13 + int32(16)
		return v191
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = l5
			F_errmsg_internal(m, int32(_a_F_array_seek_0), v13)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_array_seek_1), int32(322), int32(_a_F_array_seek_2))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 6:
		v36 = int32(4)
		v37 = int32(0)
		if l2|base.B2i32(l4 <= v37) == v37 {
			v191 = l0 + (l4+v36-int32(1))&(int32(0)-v36)*l3
		} else {
			if l2 == int32(0) {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v54 = int32(0)
					v61 = l0
					v63 = v54
					for {
						if base.B2i32(l4 <= v54) == int32(0) {
							v108 = v61 + l4
						} else {
							if l4 == int32(-1) {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								if v76 == int32(1) {
									v80 = int32(18)
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
									if v82 == v80 {
										v85 = v80
									} else {
										v85 = int32(2)
									}
									if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v92 = int32(6)
									} else {
										v92 = v85
									}
									v108 = v61 + v92
								} else {
									v94 = int32(1)
									if v76&v94 != 0 {
										v108 = v61 + int32(base.Ui32(v76)>>(uint(v94)%32))
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v108 = v61 + int32(base.Ui32(v99)>>(uint(int32(2))%32))
									}
								}
							} else {
								v103 = F_strlen(m, v61)
								mBase = m.M
								v108 = v103 + v61 + int32(1)
							}
						}
						v110 = (v108 + (v36 - int32(1))) & (v54 - v36)
						v112 = v63 + int32(1)
						if v112 != l3 {
							v61 = v110
							v63 = v112
							continue
						} else {
							break
						}
						break
					}
					v191 = v110
				}
			} else {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v118 = int32(1)
					v121 = base.I32_div_s(l1, int32(8))
					v127 = l0
					v129 = l2 + v121
					v132 = v118 << (uint(l1&int32(7)) % 32)
					v136 = int32(0)
					for {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
						if v132&v137 != 0 {
							if int32(0) < l4 {
								v176 = v127 + l4
							} else {
								if l4 == int32(-1) {
									v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
									if v144 == int32(1) {
										v148 = int32(18)
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
										if v150 == v148 {
											v153 = v148
										} else {
											v153 = int32(2)
										}
										if base.Ui32((v150-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v160 = int32(6)
										} else {
											v160 = v153
										}
										v176 = v127 + v160
									} else {
										v162 = int32(1)
										if v144&v162 != 0 {
											v176 = v127 + int32(base.Ui32(v144)>>(uint(v162)%32))
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
											v176 = v127 + int32(base.Ui32(v167)>>(uint(int32(2))%32))
										}
									}
								} else {
									v171 = F_strlen(m, v127)
									mBase = m.M
									v176 = v171 + v127 + int32(1)
								}
							}
							v179 = (v176 + (v36 - v118)) & (int32(0) - v36)
						} else {
							v179 = v127
						}
						v181 = int32(1)
						v183 = v132 << (uint(v181) % 32)
						v185 = base.B2i32(v183 == int32(256))
						if v183 == int32(256) {
							v186 = v181
						} else {
							v186 = v183
						}
						v189 = v136 + int32(1)
						if v189 != l3 {
							v127 = v179
							v129 = v185 + v129
							v132 = v186
							v136 = v189
							continue
						} else {
							break
						}
						break
					}
					v191 = v179
				}
			}
		}
		m.G0 = v13 + int32(16)
		return v191
	case 16:
		v36 = int32(2)
		v37 = int32(0)
		if l2|base.B2i32(l4 <= v37) == v37 {
			v191 = l0 + (l4+v36-int32(1))&(int32(0)-v36)*l3
		} else {
			if l2 == int32(0) {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v54 = int32(0)
					v61 = l0
					v63 = v54
					for {
						if base.B2i32(l4 <= v54) == int32(0) {
							v108 = v61 + l4
						} else {
							if l4 == int32(-1) {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								if v76 == int32(1) {
									v80 = int32(18)
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
									if v82 == v80 {
										v85 = v80
									} else {
										v85 = int32(2)
									}
									if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v92 = int32(6)
									} else {
										v92 = v85
									}
									v108 = v61 + v92
								} else {
									v94 = int32(1)
									if v76&v94 != 0 {
										v108 = v61 + int32(base.Ui32(v76)>>(uint(v94)%32))
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v108 = v61 + int32(base.Ui32(v99)>>(uint(int32(2))%32))
									}
								}
							} else {
								v103 = F_strlen(m, v61)
								mBase = m.M
								v108 = v103 + v61 + int32(1)
							}
						}
						v110 = (v108 + (v36 - int32(1))) & (v54 - v36)
						v112 = v63 + int32(1)
						if v112 != l3 {
							v61 = v110
							v63 = v112
							continue
						} else {
							break
						}
						break
					}
					v191 = v110
				}
			} else {
				if l3 <= int32(0) {
					v191 = l0
				} else {
					v118 = int32(1)
					v121 = base.I32_div_s(l1, int32(8))
					v127 = l0
					v129 = l2 + v121
					v132 = v118 << (uint(l1&int32(7)) % 32)
					v136 = int32(0)
					for {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
						if v132&v137 != 0 {
							if int32(0) < l4 {
								v176 = v127 + l4
							} else {
								if l4 == int32(-1) {
									v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
									if v144 == int32(1) {
										v148 = int32(18)
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
										if v150 == v148 {
											v153 = v148
										} else {
											v153 = int32(2)
										}
										if base.Ui32((v150-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v160 = int32(6)
										} else {
											v160 = v153
										}
										v176 = v127 + v160
									} else {
										v162 = int32(1)
										if v144&v162 != 0 {
											v176 = v127 + int32(base.Ui32(v144)>>(uint(v162)%32))
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
											v176 = v127 + int32(base.Ui32(v167)>>(uint(int32(2))%32))
										}
									}
								} else {
									v171 = F_strlen(m, v127)
									mBase = m.M
									v176 = v171 + v127 + int32(1)
								}
							}
							v179 = (v176 + (v36 - v118)) & (int32(0) - v36)
						} else {
							v179 = v127
						}
						v181 = int32(1)
						v183 = v132 << (uint(v181) % 32)
						v185 = base.B2i32(v183 == int32(256))
						if v183 == int32(256) {
							v186 = v181
						} else {
							v186 = v183
						}
						v189 = v136 + int32(1)
						if v189 != l3 {
							v127 = v179
							v129 = v185 + v129
							v132 = v186
							v136 = v189
							continue
						} else {
							break
						}
						break
					}
					v191 = v179
				}
			}
		}
		m.G0 = v13 + int32(16)
		return v191
	}
}
func F_array_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v291 int64
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_DatumGetAnyArrayP(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v28 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = int32(40)
	goto L5
L4:
	;
	v31 = int32(12)
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22+v31)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v35 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L65
	}
L7:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v81 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	F_get_type_io_data(m, v33, int32(3), v51+int32(4), v51+int32(6), v51+int32(7), v51+int32(8), v51+int32(12), v51+int32(16))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v40 = F_MemoryContextAlloc(m, v38, int32(48))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v49 == v33 {
		v77 = v35
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v40
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v33 ^ int32(-1)
	v51 = v45
	goto L8
L13:
	;
	v51 = v35
	goto L8
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v67 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	F_fmgr_info_cxt(m, v67, v51+int32(20), v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v33
	v77 = v51
	goto L7
L17:
	;
	v84 = int32(28)
	goto L19
L18:
	;
	v84 = int32(4)
	goto L19
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v22+v84)))
	if v81 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(v77)+7)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+6)))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77)+4)))
	v102 = F_ArrayGetNItemsSafe(m, v86, v97)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v97 = v89
	v98 = v90
	goto L20
L22:
	;
	goto L23
L23:
	;
	v92 = v22 + int32(16)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v97 = v92
	v98 = v92 + v93<<(uint(int32(2))%32)
	goto L20
L24:
	;
	v105 = v17 + int32(-16)
	F_pq_begintypsend(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_enlargeStringInfo(m, v105, int32(4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v116 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v111+v112))) = base.I32_rotr(v86, int32(24))&v116 | base.I32_rotr(v86&v116, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v111 + int32(4)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v127 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v143 = v17 + int32(-16)
	F_enlargeStringInfo(m, v143, int32(4))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L34
	}
L28:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v130 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v141 = base.B2i32(v138 != int32(0))
	goto L27
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v141 = base.B2i32(v131 != int32(0))
	goto L27
L32:
	;
	goto L33
L33:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v141 = base.B2i32(v135 != int32(0))
	goto L27
L34:
	;
	v147 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	if v141 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v153 = int32(16777216)
	goto L37
L36:
	;
	v153 = v147
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148+v149))) = v153
	v155 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v148 + v155
	F_enlargeStringInfo(m, v143, v155)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v166 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v161+v162))) = base.I32_rotr(v33, int32(24))&v166 | base.I32_rotr(v33&v166, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v161 + int32(4)
	if int32(0) < v86 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v179 = v147
	goto L42
L40:
	;
	goto L41
L41:
	;
	F_array_iter_setup(m, v17+int32(-44), v22, v101, v100&int32(1), v99)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L47
	}
L42:
	;
	v196 = v179 << (uint(int32(2)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v97+v196)))
	v200 = v17 + int32(-16)
	F_enlargeStringInfo(m, v200, int32(4))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v209 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v204+v205))) = base.I32_rotr(v198, int32(24))&v209 | base.I32_rotr(v198&v209, int32(8))
	v217 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v204 + v217
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v196+v98)))
	F_enlargeStringInfo(m, v200, v217)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v230 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v225+v226))) = base.I32_rotr(v221, int32(24))&v230 | base.I32_rotr(v221&v230, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v225 + int32(4)
	v242 = v179 + int32(1)
	if v242 != v86 {
		v179 = v242
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	if int32(0) < v102 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v271 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v372 = v17 + int32(-16)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v374))) = v375 << (uint(int32(2)) % 32)
	goto L64
L51:
	;
	v291 = F_array_iter_next(m, v17+int32(-44), v17+int32(-45), v271)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+19)))
	if v293 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v353 = v271 + int32(1)
	if v353 != v102 {
		v271 = v353
		goto L51
	} else {
		goto L63
	}
L55:
	;
	F_enlargeStringInfo(m, v17+int32(-16), int32(4))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v309 = F_SendFunctionCall(m, v77+int32(20), v291)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v301+v302))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v301 + int32(4)
	goto L54
L59:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v313 = v17 + int32(-16)
	F_enlargeStringInfo(m, v313, int32(4))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v320 = int32(2)
	v322 = int32(4)
	v323 = int32(base.Ui32(v311)>>(uint(v320)%32)) - v322
	v324 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v317+v318))) = base.I32_rotr(v323&v324, int32(8)) | base.I32_rotr(v323, int32(24))&v324
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v317 + v322
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	F_appendBinaryStringInfo(m, v313, v309+v322, int32(base.Ui32(v339)>>(uint(v320)%32))-v322)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_pfree(m, v309)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L54
L63:
	;
	goto L52
L64:
	;
	m.G0 = v19 - int32(-64)
	return base.I64_extend_i32_u(v374)
L65:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v391 = F_format_type_be(m, v33)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v391
	F_errmsg(m, int32(_a_F_array_send_0), v19)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_array_send_1), int32(1594), int32(_a_F_array_send_2))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_set(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v11 = F_array_set_element(m, base.I64_extend_i32_u(l0), int32(1), l1, l2, int32(0), int32(-1), l3, l4, int32(105))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = F_pg_detoast_datum(m, base.I32_wrap_i64(v11))
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return v16
		}
	}
}
func F_array_shuffle_n(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	v6 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(80)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v26 <= v6)|base.B2i32(l1 <= v6) == v6 {
		v35 = l0 + int32(16)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
		if int32(0) < v36 {
			v45 = v26 << (uint(int32(2)) % 32)
			v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(l4)+8)))
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+10)))
			v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4)+11)))
			F_deconstruct_array(m, l0, v47, v48&int32(1), v51, v24+int32(12), v24+int32(8), v24+int32(76))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v62 = base.I32_div_s(v60, v61)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = v62
				v66 = base.I64_extend_i32_s(v61 - int32(1))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				v76 = v68
				v77 = v69
				v88 = int64(0)
				for {
					v91 = int32(_a_F_array_shuffle_n_0)
					if base.Ui64(v66) <= base.Ui64(v88) {
						v138 = v88
					} else {
						v98 = v66 - v88
						v100 = *(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[0]))
						v101 = *(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[1]))
						v104 = v101
						v106 = v100
						for {
							v110 = v104 ^ v106
							v112 = base.I64_rotl(v110, int64(37))
							v120 = v110 ^ (v110<<(uint(int64(16))%64) ^ base.I64_rotl(v104, int64(24)))
							v125 = int64(base.Ui64(base.I64_rotl(v104*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v98)) % 64))
							if base.Ui64(v98) < base.Ui64(v125) {
								v104 = v120
								v106 = v112
								continue
							} else {
								break
							}
							break
						}
						*(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[0])) = v112
						*(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[1])) = v120
						v138 = v88 + v125
					}
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
					if int32(0) < v139 {
						v143 = v139 * base.I32_wrap_i64(v138)
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
						v151 = v143 + v144
						v155 = v146 + v143<<(uint(int32(3))%32)
						v157 = v76
						v158 = v77
						v162 = int32(0)
						for {
							v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
							v173 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
							v174 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
							*(*int64)(unsafe.Add(mBase, uint32(v158))) = v174
							v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
							*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v176)
							*(*int64)(unsafe.Add(mBase, uint32(v155))) = v173
							*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v172)
							v180 = int32(1)
							v182 = int32(8)
							v185 = v157 + v180
							v187 = v158 + v182
							v189 = v162 + v180
							v190 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
							if v189 < v190 {
								v151 = v151 + v180
								v155 = v155 + v182
								v157 = v185
								v158 = v187
								v162 = v189
								continue
							} else {
								break
							}
							break
						}
						v198 = v185
						v199 = v187
					} else {
						v198 = v76
						v199 = v77
					}
					v214 = v88 + int64(1)
					if v214 != base.I64_extend_i32_u(l1) {
						v76 = v198
						v77 = v199
						v88 = v214
						continue
					} else {
						break
					}
					break
				}
				v216 = int32(0)
				v217 = base.B2i32(v45 == v216)
				if v217 == v216 {
					base.MemoryCopy(m, v24+int32(48), v35, v45)
				} else {
				}
				if v217 == int32(0) {
					base.MemoryCopy(m, v24+int32(16), v35+v45, v45)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = l1
				if l2 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(1)
				} else {
				}
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				v234 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				v241 = F_construct_md_array(m, v233, v234, v26, v24+int32(48), v24+int32(16), l3, v47, v48&int32(1), v51)
				mBase = m.M
				v242 = m.ExcPending
				if v242 != 0 {
					return int32(0)
				} else {
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
					F_pfree(m, v243)
					mBase = m.M
					v245 = m.ExcPending
					if v245 != 0 {
						return int32(0)
					} else {
						v246 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						F_pfree(m, v246)
						mBase = m.M
						v248 = m.ExcPending
						if v248 != 0 {
							return int32(0)
						} else {
							v249 = v241
							m.G0 = v24 + int32(80)
							return v249
						}
					}
				}
			}
		} else {
			v40 = F_construct_empty_array(m, l3)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v249 = v40
				m.G0 = v24 + int32(80)
				return v249
			}
		}
	} else {
		v40 = F_construct_empty_array(m, l3)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			v249 = v40
			m.G0 = v24 + int32(80)
			return v249
		}
	}
}
func F_array_smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	v4 = F_array_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 < int32(0) {
			v10 = int32(24)
		} else {
			v10 = int32(40)
		}
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_array_subscript_fetch(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+6)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)))
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+9)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v17 = F_array_get_element(m, v6, v8, v9+int32(12), v12, v13, v14, v15, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int64)(unsafe.Add(mBase, uint32(v19))) = v17
		return
	}
}
func F_array_subscript_handler_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	v4 = int64(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(472) {
		v27 = v4
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 == int32(0) {
			v27 = v4
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v14 != int32(8) {
				v27 = v4
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				if v17 != 0 {
					v27 = v4
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v18 != v19 {
						v27 = v4
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						if v21 == int32(0) {
							v27 = v4
						} else {
							v27 = base.I64_extend_i32_u(v11)
						}
					}
				}
			}
		}
	}
	return v27
}
func F_initArrayResultAny(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = int32(0)
	v5 = F_get_array_type(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v13 = F_initArrayResultArr(m, l0, int32(0), l1, int32(1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v19 = v13
				v20 = v13
				v21 = v3
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v24 = F_MemoryContextAlloc(m, v22, int32(8))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = v21
					return v24
				}
			}
		} else {
			v17 = F_initArrayResultWithSize(m, l0, l1, int32(1), int32(64))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = v17
				v20 = v3
				v21 = v17
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v24 = F_MemoryContextAlloc(m, v22, int32(8))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = v21
					return v24
				}
			}
		}
	}
}
