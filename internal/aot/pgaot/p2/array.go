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
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	v9 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = (v18<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v33 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = v15
	v33 = l0 + v26<<(uint(int32(3))%32) + int32(16)
	goto L1
L5:
	;
	return
L6:
	;
	v36 = int32(1)
	v42 = int32(0)
	v49 = l0 + v32
	v50 = v36
	v51 = v33
	v52 = v9
	goto L7
L7:
	;
	if l2 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if base.B2i32(v102 == int32(0))|base.B2i32(v101 == int32(1)) != 0 {
		goto L5
	} else {
		goto L27
	}
L9:
	;
	v107 = v42 + int32(1)
	if v107 != l3 {
		v42 = v107
		v49 = v100
		v50 = v101
		v51 = v102
		v52 = v103
		goto L7
	} else {
		goto L26
	}
L10:
	;
	v92 = v50 << (uint(int32(1)) % 32)
	if v92 != int32(256) {
		v100 = v87
		v101 = v92
		v102 = v51
		v103 = v88
		goto L9
	} else {
		goto L25
	}
L11:
	;
	v77 = l1 + v42<<(uint(int32(2))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = F_ArrayCastAndSet(m, v78, l4, l5, l6, v49)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L15
	} else {
		goto L19
	}
L12:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+l2))))
	if v59 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if v51 != 0 {
		v87 = v49
		v88 = v52
		goto L10
	} else {
		goto L14
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	F_errmsg_internal(m, int32(_a_F_CopyArrayEls_0), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_CopyArrayEls_1), int32(984), int32(_a_F_CopyArrayEls_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	if l7&(l5^v36) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	F_pfree(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L15
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v84 = v50 | v52
	v85 = v49 + v79
	if v51 != 0 {
		v87 = v85
		v88 = v84
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v100 = v85
	v101 = v50
	v102 = int32(0)
	v103 = v84
	goto L9
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v88)
	v96 = int32(1)
	v100 = v87
	v101 = v96
	v102 = v51 + v96
	v103 = int32(0)
	goto L9
L26:
	;
	goto L8
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v103)
	goto L5
}
func F_array_agg_array_transfn(m *base.Module, l0 int32) int32 {
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_get_fn_expr_argtype(m, v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
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
				switch v20 - int32(429) {
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
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
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
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_array_agg_array_transfn_0), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_array_transfn_1), int32(953), int32(_a_F_array_agg_array_transfn_2))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v48 == int32(1) {
					v51 = int32(0)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					v54 = F_initArrayResultArr(m, v10, v51, v52, v51)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v57 = v54
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						v61 = F_accumArrayResultArr(m, v57, v58, v59, v10, v60)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(16)
							return v61
						}
					}
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v57 = v56
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					v61 = F_accumArrayResultArr(m, v57, v58, v59, v10, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return v61
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_array_agg_array_transfn_3), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_array_transfn_1), int32(942), int32(_a_F_array_agg_array_transfn_2))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
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
func F_array_append_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v5 != int32(464) {
		v23 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v10 == int32(0) {
			v23 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(8) {
				v23 = v2
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v16 != 0 {
					v23 = v2
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
					if v18 == v19 {
						v21 = v10
					} else {
						v21 = int32(0)
					}
					v23 = v21
				}
			}
		}
	}
	return v23
}
func F_array_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_DatumGetAnyArrayP(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = F_DatumGetAnyArrayP(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v35 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v38 = int32(28)
	goto L6
L5:
	;
	v38 = int32(4)
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v42 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = int32(28)
	goto L9
L8:
	;
	v45 = int32(4)
	goto L9
L9:
	;
	if v42 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31+v38)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26+v45)))
	if v35 == int32(-1) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v52 = v49
	goto L10
L12:
	;
	goto L13
L13:
	;
	v52 = v26 + int32(16)
	goto L10
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = F_ArrayGetNItemsSafe(m, v54, v52)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v60 = v57
	goto L14
L16:
	;
	goto L17
L17:
	;
	v60 = v31 + int32(16)
	goto L14
L18:
	;
	v64 = F_ArrayGetNItemsSafe(m, v53, v60)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v68 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L142
	}
L21:
	;
	v71 = int32(40)
	goto L23
L22:
	;
	v71 = int32(12)
	goto L23
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v26+v71)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v76 == int32(-1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = int32(40)
	goto L26
L25:
	;
	v79 = int32(12)
	goto L26
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v31+v79)))
	if v73 == v81 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if v84 != 0 {
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
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L138
	}
L30:
	;
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96)+11)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+10)))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96)+8)))
	v101 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+78)) = uint16(v101)
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+76)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v96 + int32(104)
	v111 = base.B2i32(v62 < v64)
	if v97 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v85 == v73 {
		v96 = v84
		v97 = v68
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v88 = F_lookup_type_cache(m, v73, int32(64))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+108))
	if v90 == int32(0) {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v88
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v96 = v88
	v97 = v95
	goto L30
L37:
	;
	if v62 < v64 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	if v114 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = int64(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v147 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v116
	v170 = v117
	goto L37
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = int64(0)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	if v124 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v123 + (v127<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v170 = int32(0)
	goto L37
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v123 + v124
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v170 = v123 + v139<<(uint(int32(3))%32) + int32(16)
	goto L37
L47:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v26 + (v150<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v170 = int32(0)
	goto L37
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v147 + v26
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v170 = v26 + v162<<(uint(int32(3))%32) + int32(16)
	goto L37
L50:
	;
	v171 = v62
	goto L52
L51:
	;
	v171 = v64
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v170
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v175 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = int32(1)
	if int32(0) < v171 {
		goto L67
	} else {
		goto L68
	}
L54:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	if v178 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+20)) = int64(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v211 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
	v181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v180
	v234 = v181
	goto L53
L58:
	;
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+20)) = int64(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v188 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v187 + (v191<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v234 = int32(0)
	goto L53
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v187 + v188
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v234 = v187 + v203<<(uint(int32(3))%32) + int32(16)
	goto L53
L63:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v31 + (v214<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v234 = int32(0)
	goto L53
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v211 + v31
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v234 = v31 + v226<<(uint(int32(3))%32) + int32(16)
	goto L53
L66:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v487 == int32(-1) {
		goto L130
	} else {
		goto L131
	}
L67:
	;
	v241 = v99 & int32(1)
	v244 = int32(0)
	goto L72
L68:
	;
	goto L69
L69:
	;
	if v62 != v64 {
		goto L88
	} else {
		goto L89
	}
L70:
	;
	if v320 != 0 {
		v472 = v320
		goto L66
	} else {
		goto L87
	}
L71:
	;
	v312 = int32(-1)
	if v275 != 0 {
		goto L84
	} else {
		goto L85
	}
L72:
	;
	v267 = F_array_iter_next(m, v23+int32(40), v23+int32(19), v244, v100, v241, v98)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	v320 = int32(0)
	goto L70
L74:
	;
	v273 = F_array_iter_next(m, v23+int32(20), v23+int32(18), v244, v100, v241, v98)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+19)))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+18)))
	if v275|v276&int32(1) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v309 = v244 + int32(1)
	if v309 != v171 {
		v244 = v309
		goto L72
	} else {
		goto L83
	}
L77:
	;
	v282 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+92)) = uint8(v282)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v273
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+84)) = uint8(v282)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v267
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v292 = m.T0[v291].(func(*base.Module, int32) int32)(m, v23+int32(60))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v300 = int32(0)
	if base.B2i32(v275 == v300)|base.B2i32(v276&int32(1) == v300) != 0 {
		goto L71
	} else {
		goto L82
	}
L80:
	;
	if v292 == int32(0) {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v472 = v292>>(uint(int32(31))%32) | int32(1)
	goto L66
L82:
	;
	goto L76
L83:
	;
	goto L73
L84:
	;
	v317 = (v276 ^ v312) & int32(1)
	goto L86
L85:
	;
	v317 = v312
	goto L86
L86:
	;
	v320 = v317
	goto L70
L87:
	;
	goto L69
L88:
	;
	if v62 < v64 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v54 == v53 {
		goto L96
	} else {
		goto L97
	}
L91:
	;
	v344 = int32(-1)
	goto L93
L92:
	;
	v344 = int32(1)
	goto L93
L93:
	;
	v472 = v344
	goto L66
L94:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v407 == int32(-1) {
		goto L113
	} else {
		goto L114
	}
L95:
	;
	v354 = v346
	goto L103
L96:
	;
	v346 = int32(0)
	if v54 <= v346 {
		goto L94
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v54 < v53 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L95
L100:
	;
	v352 = int32(-1)
	goto L102
L101:
	;
	v352 = int32(1)
	goto L102
L102:
	;
	v472 = v352
	goto L66
L103:
	;
	v374 = v354 << (uint(int32(2)) % 32)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v52+v374)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374+v60)))
	if v376 == v378 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v376 < v378 {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	v381 = v354 + int32(1)
	if v54 != v381 {
		v354 = v381
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	goto L94
L109:
	;
	v386 = int32(-1)
	goto L111
L110:
	;
	v386 = int32(1)
	goto L111
L111:
	;
	v472 = v386
	goto L66
L112:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v418 == int32(-1) {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v417 = v410
	goto L112
L114:
	;
	goto L115
L115:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v417 = v26 + v411<<(uint(int32(2))%32) + int32(16)
	goto L112
L116:
	;
	v429 = int32(0)
	if v54 <= v429 {
		v472 = v429
		goto L66
	} else {
		goto L120
	}
L117:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	v428 = v421
	goto L116
L118:
	;
	goto L119
L119:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v428 = v31 + v422<<(uint(int32(2))%32) + int32(16)
	goto L116
L120:
	;
	v434 = int32(0)
	goto L121
L121:
	;
	v454 = v434 << (uint(int32(2)) % 32)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v417+v454)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v428+v454)))
	if v456 == v458 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v456 < v458 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	v461 = v434 + int32(1)
	if v54 != v461 {
		v434 = v461
		goto L121
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	v472 = v429
	goto L66
L127:
	;
	v466 = int32(-1)
	goto L129
L128:
	;
	v466 = int32(1)
	goto L129
L129:
	;
	v472 = v466
	goto L66
L130:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v494 == int32(-1) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v26 == v490 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	F_pfree(m, v26)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	m.G0 = v23 + int32(96)
	return v472
L135:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v31 == v497 {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	F_pfree(m, v31)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_array_cmp_0), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_array_cmp_1), int32(4017), int32(_a_F_array_cmp_2))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v528 = F_format_type_be(m, v73)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v528
	F_errmsg(m, int32(_a_F_array_cmp_3), v23)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_array_cmp_1), int32(4035), int32(_a_F_array_cmp_2))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_iter_next(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v17 != 0 {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+l2))))
			v21 = v19
		} else {
			v21 = int32(0)
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v13+l2<<(uint(int32(2))%32))))
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v21)
		v132 = v22
		m.G0 = v11 + int32(16)
		return v132
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v24 == int32(0) {
			v33 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v33)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if l4 != 0 {
				switch l3 - int32(1) {
				case 0:
					v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35))))
					v98 = v38
					v99 = v35 + l3
					switch l5 - int32(99) {
					case 0:
						v114 = v99
					case 1:
						v114 = (v99 + int32(7)) & int32(-8)
					default:
						v114 = (v99 + int32(1)) & int32(-2)
					case 6:
						v114 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
					v117 = v98
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v120 = v118 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
					if v120 != int32(256) {
						v132 = v117
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v124 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v132 = v117
					}
					m.G0 = v11 + int32(16)
					return v132
				case 1:
					v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35))))
					v98 = v40
					v99 = v35 + l3
					switch l5 - int32(99) {
					case 0:
						v114 = v99
					case 1:
						v114 = (v99 + int32(7)) & int32(-8)
					default:
						v114 = (v99 + int32(1)) & int32(-2)
					case 6:
						v114 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
					v117 = v98
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v120 = v118 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
					if v120 != int32(256) {
						v132 = v117
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v124 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v132 = v117
					}
					m.G0 = v11 + int32(16)
					return v132
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
						F_errmsg_internal(m, int32(_a_F_array_iter_next_0), v11)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_array_iter_next_1), int32(70), int32(_a_F_array_iter_next_2))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					v98 = v42
					v99 = v35 + l3
					switch l5 - int32(99) {
					case 0:
						v114 = v99
					case 1:
						v114 = (v99 + int32(7)) & int32(-8)
					default:
						v114 = (v99 + int32(1)) & int32(-2)
					case 6:
						v114 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
					v117 = v98
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v120 = v118 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
					if v120 != int32(256) {
						v132 = v117
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v124 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v132 = v117
					}
					m.G0 = v11 + int32(16)
					return v132
				}
			} else {
				if int32(0) < l3 {
					v98 = v35
					v99 = v35 + l3
				} else {
					if l3 == int32(-1) {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
						if v64 == int32(1) {
							v68 = int32(18)
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
							if v70 == v68 {
								v73 = v68
							} else {
								v73 = int32(2)
							}
							if base.Ui32((v70-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v80 = int32(6)
							} else {
								v80 = v73
							}
							v96 = v35 + v80
						} else {
							v82 = int32(1)
							if v64&v82 != 0 {
								v96 = v35 + int32(base.Ui32(v64)>>(uint(v82)%32))
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
								v96 = v35 + int32(base.Ui32(v87)>>(uint(int32(2))%32))
							}
						}
					} else {
						v91 = F_strlen(m, v35)
						mBase = m.M
						v96 = v91 + v35 + int32(1)
					}
					v98 = v35
					v99 = v96
				}
				switch l5 - int32(99) {
				case 0:
					v114 = v99
				case 1:
					v114 = (v99 + int32(7)) & int32(-8)
				default:
					v114 = (v99 + int32(1)) & int32(-2)
				case 6:
					v114 = (v99 + int32(3)) & int32(-4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
				v117 = v98
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v120 = v118 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
				if v120 != int32(256) {
					v132 = v117
				} else {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v124 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
					v132 = v117
				}
				m.G0 = v11 + int32(16)
				return v132
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			if v27&v28 != 0 {
				v33 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v33)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if l4 != 0 {
					switch l3 - int32(1) {
					case 0:
						v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35))))
						v98 = v38
						v99 = v35 + l3
						switch l5 - int32(99) {
						case 0:
							v114 = v99
						case 1:
							v114 = (v99 + int32(7)) & int32(-8)
						default:
							v114 = (v99 + int32(1)) & int32(-2)
						case 6:
							v114 = (v99 + int32(3)) & int32(-4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
						v117 = v98
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v120 = v118 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
						if v120 != int32(256) {
							v132 = v117
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v124 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v132 = v117
						}
						m.G0 = v11 + int32(16)
						return v132
					case 1:
						v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35))))
						v98 = v40
						v99 = v35 + l3
						switch l5 - int32(99) {
						case 0:
							v114 = v99
						case 1:
							v114 = (v99 + int32(7)) & int32(-8)
						default:
							v114 = (v99 + int32(1)) & int32(-2)
						case 6:
							v114 = (v99 + int32(3)) & int32(-4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
						v117 = v98
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v120 = v118 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
						if v120 != int32(256) {
							v132 = v117
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v124 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v132 = v117
						}
						m.G0 = v11 + int32(16)
						return v132
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
							F_errmsg_internal(m, int32(_a_F_array_iter_next_0), v11)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_array_iter_next_1), int32(70), int32(_a_F_array_iter_next_2))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 3:
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						v98 = v42
						v99 = v35 + l3
						switch l5 - int32(99) {
						case 0:
							v114 = v99
						case 1:
							v114 = (v99 + int32(7)) & int32(-8)
						default:
							v114 = (v99 + int32(1)) & int32(-2)
						case 6:
							v114 = (v99 + int32(3)) & int32(-4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
						v117 = v98
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v120 = v118 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
						if v120 != int32(256) {
							v132 = v117
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v124 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v132 = v117
						}
						m.G0 = v11 + int32(16)
						return v132
					}
				} else {
					if int32(0) < l3 {
						v98 = v35
						v99 = v35 + l3
					} else {
						if l3 == int32(-1) {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
							if v64 == int32(1) {
								v68 = int32(18)
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
								if v70 == v68 {
									v73 = v68
								} else {
									v73 = int32(2)
								}
								if base.Ui32((v70-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v80 = int32(6)
								} else {
									v80 = v73
								}
								v96 = v35 + v80
							} else {
								v82 = int32(1)
								if v64&v82 != 0 {
									v96 = v35 + int32(base.Ui32(v64)>>(uint(v82)%32))
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
									v96 = v35 + int32(base.Ui32(v87)>>(uint(int32(2))%32))
								}
							}
						} else {
							v91 = F_strlen(m, v35)
							mBase = m.M
							v96 = v91 + v35 + int32(1)
						}
						v98 = v35
						v99 = v96
					}
					switch l5 - int32(99) {
					case 0:
						v114 = v99
					case 1:
						v114 = (v99 + int32(7)) & int32(-8)
					default:
						v114 = (v99 + int32(1)) & int32(-2)
					case 6:
						v114 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v114
					v117 = v98
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v120 = v118 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
					if v120 != int32(256) {
						v132 = v117
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v124 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v132 = v117
					}
					m.G0 = v11 + int32(16)
					return v132
				}
			} else {
				v30 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
				v117 = int32(0)
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v120 = v118 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120
				if v120 != int32(256) {
					v132 = v117
				} else {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v124 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v124 + int32(1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
					v132 = v117
				}
				m.G0 = v11 + int32(16)
				return v132
			}
		}
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v17 + int32(1)
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
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v135 {
		goto L42
	} else {
		goto L43
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v72
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v74 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	switch v48&int32(_a_F_array_iterate_0) - int32(1) {
	case 0:
		goto L16
	case 1:
		goto L15
	default:
		goto L13
	case 3:
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v72 = v42
	goto L9
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v72 = v55
	goto L9
L15:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42))))
	v72 = v54
	goto L9
L16:
	;
	v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42))))
	v72 = v53
	goto L9
L17:
	;
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v48
	F_errmsg_internal(m, int32(_a_F_array_iterate_1), v15)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_array_iterate_2), int32(70), int32(_a_F_array_iterate_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	switch v113 - int32(99) {
	case 0:
		v130 = v112
		goto L38
	case 1:
		goto L40
	default:
		goto L39
	case 6:
		goto L41
	}
L22:
	;
	v112 = v74 + v42
	goto L21
L23:
	;
	goto L24
L24:
	;
	if v74 == int32(-1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v80 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v107 = F_strlen(m, v42)
	mBase = m.M
	v112 = v107 + v42 + int32(1)
	goto L21
L28:
	;
	v84 = int32(18)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v86 == v84 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v98 = int32(1)
	if v80&v98 != 0 {
		v112 = v42 + int32(base.Ui32(v80)>>(uint(v98)%32))
		goto L21
	} else {
		goto L37
	}
L31:
	;
	v89 = v84
	goto L33
L32:
	;
	v89 = int32(2)
	goto L33
L33:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v96 = int32(6)
	goto L36
L35:
	;
	v96 = v89
	goto L36
L36:
	;
	v112 = v42 + v96
	goto L21
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v112 = v42 + int32(base.Ui32(v103)>>(uint(int32(2))%32))
	goto L21
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v130
	goto L1
L39:
	;
	v130 = (v112 + int32(1)) & int32(-2)
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = (v112 + int32(7)) & int32(-8)
	goto L1
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = (v112 + int32(3)) & int32(-4)
	goto L1
L42:
	;
	v142 = v132
	v143 = int32(0)
	goto L45
L43:
	;
	v273 = v132
	v274 = v20
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v273
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
	v287 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v289 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+15)))
	v290 = F_construct_md_array(m, v134, v133, v274, v283, v284, v286, v287, v288, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L17
	} else {
		goto L83
	}
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v151 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v155 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v273 = v263
	v274 = v269
	goto L44
L47:
	;
	v266 = v143 + int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v266 < v267 {
		v142 = v263
		v143 = v266
		goto L45
	} else {
		goto L82
	}
L48:
	;
	v176 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v143+v133))) = uint8(v176)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v181 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v159 = base.I32_div_s(v151, int32(8))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v159))))
	if int32(base.Ui32(v161)>>(uint(v151&int32(7))%32))&int32(1) != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v143+v133))) = uint8(v168)
	*(*int32)(unsafe.Add(mBase, uint32(v134+v143<<(uint(int32(2))%32)))) = int32(0)
	v263 = v142
	goto L47
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134+v143<<(uint(int32(2))%32)))) = v207
	v209 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v209 {
		goto L63
	} else {
		goto L64
	}
L52:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	switch v184 - int32(1) {
	case 0:
		goto L58
	case 1:
		goto L57
	default:
		goto L55
	case 3:
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v207 = v142
	goto L51
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L17
	} else {
		goto L59
	}
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v207 = v189
	goto L51
L57:
	;
	v188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142))))
	v207 = v188
	goto L51
L58:
	;
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v142))))
	v207 = v187
	goto L51
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = base.I32_extend16_s(v184)
	F_errmsg_internal(m, int32(_a_F_array_iterate_1), v15+int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_array_iterate_2), int32(70), int32(_a_F_array_iterate_3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L17
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
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	switch v248 - int32(99) {
	case 0:
		v263 = v247
		goto L47
	case 1:
		goto L80
	default:
		goto L79
	case 6:
		goto L81
	}
L63:
	;
	v247 = v142 + v209
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v209 == int32(-1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v215 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v242 = F_strlen(m, v142)
	mBase = m.M
	v247 = v242 + v142 + int32(1)
	goto L62
L69:
	;
	v219 = int32(18)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	if v221 == v219 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v233 = int32(1)
	if v215&v233 != 0 {
		v247 = v142 + int32(base.Ui32(v215)>>(uint(v233)%32))
		goto L62
	} else {
		goto L78
	}
L72:
	;
	v224 = v219
	goto L74
L73:
	;
	v224 = int32(2)
	goto L74
L74:
	;
	if base.Ui32((v221-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v231 = int32(6)
	goto L77
L76:
	;
	v231 = v224
	goto L77
L77:
	;
	v247 = v142 + v231
	goto L62
L78:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v247 = v142 + int32(base.Ui32(v238)>>(uint(int32(2))%32))
	goto L62
L79:
	;
	v263 = (v247 + int32(1)) & int32(-2)
	goto L47
L80:
	;
	v263 = (v247 + int32(7)) & int32(-8)
	goto L47
L81:
	;
	v263 = (v247 + int32(3)) & int32(-4)
	goto L47
L82:
	;
	goto L46
L83:
	;
	v292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v292)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v290
	goto L1
}
func F_array_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
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
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
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
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int32(0)
			} else {
				if v13 == int32(-1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
					v43 = v40
				} else {
					v43 = v7 + int32(16)
				}
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+v27<<(uint(int32(2))%32)-int32(4))))
				return v49
			}
		}
	}
}
func F_array_positions(m *base.Module, l0 int32) int32 {
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
	var v41 int32
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v16 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v211
L2:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
	v211 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v197 = m.ExcPending
	if v197 != 0 {
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
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v43 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v211 = v41
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
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v56 = int32(0)
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
	v211 = v52
	goto L1
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
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
	v111 = F_array_iterate(m, v105, v14+int32(12), v14+int32(11))
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
	v161 = v34
	goto L40
L40:
	;
	F_array_free_iterator(m, v105)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L54
	}
L41:
	;
	v128 = int32(1)
	v129 = v118 + v128
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if (v130|v43)&v128 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v161 = v148
	goto L40
L43:
	;
	v153 = F_array_iterate(m, v105, v14+int32(12), v14+int32(11))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L52
	}
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_array_positions[0]))
	v146 = F_accumArrayResult(m, v123, v129, int32(0), int32(23), v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L51
	}
L45:
	;
	if v130&v43 == int32(0) {
		v148 = v123
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
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
	if v138 == int32(0) {
		v148 = v123
		goto L43
	} else {
		goto L50
	}
L50:
	;
	goto L44
L51:
	;
	v148 = v146
	goto L43
L52:
	;
	if v153 != 0 {
		v118 = v129
		v123 = v148
		goto L41
	} else {
		goto L53
	}
L53:
	;
	goto L42
L54:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v168 != v23 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_pfree(m, v23)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_array_positions[0]))
	v174 = F_makeArrayResult(m, v161, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v211 = v174
	goto L1
L60:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v183 = F_format_type_be(m, v61)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v183
	F_errmsg(m, int32(_a_F_array_positions_0), v14)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_array_positions_1), int32(1554), int32(_a_F_array_positions_2))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
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
	v200 = m.ExcPending
	if v200 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_array_positions_3), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_array_positions_1), int32(1502), int32(_a_F_array_positions_2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
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
func F_array_replace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v24 = F_array_replace_internal(m, v18, v13, v14, v15, v16, int32(0), v23, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v24
			}
		}
	}
}
func F_array_replace_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v455 int32
	_ = v455
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	v5 = l4
	v9 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(80)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = l0 + int32(16)
	v38 = F_ArrayGetNItemsSafe(m, v35, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L156
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L151
	}
L3:
	;
	m.G0 = v32 + int32(80)
	return v455
L4:
	;
	return int32(0)
L5:
	;
	if v38 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v455 = l0
	goto L3
L7:
	;
	goto L8
L8:
	;
	if int32(2) <= v35 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = l5
	goto L11
L10:
	;
	v47 = int32(0)
	goto L11
L11:
	;
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+11)))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+8)))
	v64 = base.I32_extend16_s(v63)
	if v63 != int32(_a_F_array_replace_internal_0) {
		v74 = l1
		v75 = l3
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 == v34 {
		v60 = v49
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v53 = F_lookup_type_cache(m, v34, int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	if v55 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v53
	v60 = v53
	goto L13
L20:
	;
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+60)) = uint8(v77)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = l6
	*(*int64)(unsafe.Add(mBase, uint32(v32)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v60 + int32(76)
	v85 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+62)) = uint16(v85)
	v89 = F_palloc(m, v38<<(uint(v85)%32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L28
	}
L21:
	;
	if l2 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	v71 = l1
	goto L24
L24:
	;
	if v5 != 0 {
		v74 = v71
		v75 = l3
		goto L20
	} else {
		goto L26
	}
L25:
	;
	v71 = v69
	goto L24
L26:
	;
	v72 = F_pg_detoast_datum(m, l3)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v74 = v71
	v75 = v72
	goto L20
L28:
	;
	v91 = F_palloc(m, v38)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v95 = v93 << (uint(int32(3)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v98 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v99 = v37 + v95
	goto L32
L31:
	;
	v99 = int32(0)
	goto L32
L32:
	;
	if v98 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v104 = v98
	goto L35
L34:
	;
	v104 = (v95 + int32(23)) & int32(-8)
	goto L35
L35:
	;
	v106 = int32(1)
	v109 = v62 - int32(99)
	v118 = int32(0)
	v119 = v106
	v120 = v99
	v124 = l0 + v104
	v125 = v9
	v131 = v9
	v136 = v9
	v140 = v9
	goto L36
L36:
	;
	if v120 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if v358 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L38:
	;
	v360 = int32(1)
	v362 = v119 << (uint(v360) % 32)
	v364 = base.B2i32(v362 == int32(256))
	if v362 == int32(256) {
		goto L113
	} else {
		goto L114
	}
L39:
	;
	v353 = v345
	v356 = v348
	v357 = v125 + int32(1)
	v358 = v349
	v359 = v350
	goto L38
L40:
	;
	if int32(0) < v64 {
		v309 = v64
		goto L89
	} else {
		goto L90
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v125<<(uint(int32(2))%32)))) = v259
	v268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v91))) = uint8(v268)
	v272 = v261
	v273 = v262
	goto L40
L42:
	;
	v259 = v211
	v261 = v226
	v262 = v131
	goto L41
L43:
	;
	v255 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v91))) = uint8(v255)
	v345 = v118
	v348 = v124
	v349 = v131
	v350 = v255
	goto L39
L44:
	;
	if v61&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v119&v143 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	if l2 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if l5 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v353 = v118
	v356 = v124
	v357 = v125
	v358 = int32(1)
	v359 = v136
	goto L38
L49:
	;
	goto L50
L50:
	;
	if v5 != 0 {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	v259 = v75
	v261 = v124
	v262 = int32(1)
	goto L41
L52:
	;
	switch v109 {
	case 0:
		v226 = v213
		goto L77
	case 1:
		goto L79
	default:
		goto L78
	case 6:
		goto L80
	}
L53:
	;
	switch v63 - v106 {
	case 0:
		goto L59
	case 1:
		goto L58
	default:
		goto L56
	case 3:
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if int32(0) < v64 {
		v211 = v124
		v213 = v124 + v64
		goto L52
	} else {
		goto L63
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L60
	}
L57:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v211 = v155
	v213 = v124 + v64
	goto L52
L58:
	;
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v124))))
	v211 = v153
	v213 = v124 + v64
	goto L52
L59:
	;
	v151 = int32(*(*int8)(unsafe.Add(mBase, uint32(v124))))
	v211 = v151
	v213 = v124 + v64
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v64
	F_errmsg_internal(m, int32(_a_F_array_replace_internal_1), v32+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_2), int32(70), int32(_a_F_array_replace_internal_3))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	if v64 == int32(-1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v211 = v124
	v213 = v210
	goto L52
L65:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v177 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	v205 = F_strlen(m, v124)
	mBase = m.M
	v210 = v205 + v124 + int32(1)
	goto L64
L68:
	;
	v210 = v203 + v124
	goto L64
L69:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	if base.Ui32((v181-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v203 = int32(6)
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v194 = int32(1)
	if v177&v194 != 0 {
		v210 = v124 + int32(base.Ui32(v177)>>(uint(v194)%32))
		goto L64
	} else {
		goto L76
	}
L72:
	;
	v188 = int32(18)
	if v181 == v188 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v192 = v188
	goto L75
L74:
	;
	v192 = int32(2)
	goto L75
L75:
	;
	v210 = v124 + v192
	goto L64
L76:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v203 = int32(base.Ui32(v199) >> (uint(int32(2)) % 32))
	goto L68
L77:
	;
	if l2 != 0 {
		goto L42
	} else {
		goto L81
	}
L78:
	;
	v226 = (v213 + int32(1)) & int32(-2)
	goto L77
L79:
	;
	v226 = (v213 + int32(7)) & int32(-8)
	goto L77
L80:
	;
	v226 = (v213 + int32(3)) & int32(-4)
	goto L77
L81:
	;
	v227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v227)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v74
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v227)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v211
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+60)) = uint8(v227)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v239 = m.T0[v238].(func(*base.Module, int32) int32)(m, v32+int32(44))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+60)))
	if v241 != 0 {
		goto L42
	} else {
		goto L83
	}
L83:
	;
	if v239 == int32(0) {
		v259 = v211
		v261 = v226
		v262 = v131
		goto L41
	} else {
		goto L84
	}
L84:
	;
	if l5 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v353 = v118
	v356 = v226
	v357 = v125
	v358 = int32(1)
	v359 = v136
	goto L38
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v125<<(uint(int32(2))%32)))) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v91))) = uint8(v5)
	v251 = int32(1)
	if v5 == int32(0) {
		v272 = v226
		v273 = v251
		goto L40
	} else {
		goto L88
	}
L88:
	;
	v345 = v118
	v348 = v226
	v349 = v251
	v350 = int32(1)
	goto L39
L89:
	;
	v310 = v118 + v309
	switch v109 {
	case 0:
		v323 = v310
		goto L104
	case 1:
		goto L106
	default:
		goto L105
	case 6:
		goto L107
	}
L90:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v89+v125<<(uint(int32(2))%32))))
	if v64 == int32(-1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v282 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v305 = F_strlen(m, v279)
	mBase = m.M
	v309 = v305 + int32(1)
	goto L89
L94:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if base.Ui32((v286-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v309 = int32(6)
		goto L89
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if v282&int32(1) != 0 {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	v293 = int32(18)
	if v286 == v293 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v297 = v293
	goto L100
L99:
	;
	v297 = int32(2)
	goto L100
L100:
	;
	v309 = v297
	goto L89
L101:
	;
	v309 = int32(base.Ui32(v282) >> (uint(int32(1)) % 32))
	goto L89
L102:
	;
	goto L103
L103:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v309 = int32(base.Ui32(v302) >> (uint(int32(2)) % 32))
	goto L89
L104:
	;
	if base.Ui32(v323) < base.Ui32(int32(1073741824)) {
		v345 = v323
		v348 = v272
		v349 = v273
		v350 = v136
		goto L39
	} else {
		goto L108
	}
L105:
	;
	v323 = (v310 + int32(1)) & int32(-2)
	goto L104
L106:
	;
	v323 = (v310 + int32(7)) & int32(-8)
	goto L104
L107:
	;
	v323 = (v310 + int32(3)) & int32(-4)
	goto L104
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_replace_internal_4), v32+int32(32))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_5), int32(_a_F_array_replace_internal_6), int32(_a_F_array_replace_internal_7))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v365 = v360
	goto L115
L114:
	;
	v365 = v362
	goto L115
L115:
	;
	if v120 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v366 = v365
	goto L118
L117:
	;
	v366 = v119
	goto L118
L118:
	;
	if v120 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v369 = v120 + v364
	goto L121
L120:
	;
	v369 = int32(0)
	goto L121
L121:
	;
	v371 = v140 + int32(1)
	if v371 != v38 {
		v118 = v353
		v119 = v366
		v120 = v369
		v124 = v356
		v125 = v357
		v131 = v358
		v136 = v359
		v140 = v371
		goto L36
	} else {
		goto L122
	}
L122:
	;
	goto L37
L123:
	;
	F_pfree(m, v89)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v357 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	F_pfree(m, v91)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v455 = l0
	goto L3
L128:
	;
	F_pfree(m, v89)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v394 = v35 << (uint(int32(3)) % 32)
	if v359 != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	F_pfree(m, v91)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v386 = F_palloc0(m, int32(16))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v386)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v386))) = int64(64)
	v455 = v386
	goto L3
L134:
	;
	v411 = v353 + v409
	v412 = F_palloc0(m, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L138
	}
L135:
	;
	v398 = base.I32_div_s(v357+int32(7), int32(8))
	v403 = (v394 + v398 + int32(23)) & int32(-8)
	v409 = v403
	v410 = v403
	goto L134
L136:
	;
	goto L137
L137:
	;
	v409 = (v394 + int32(23)) & int32(-8)
	v410 = int32(0)
	goto L134
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v412)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v412)+8)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v412)+4)) = v35
	v417 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = v411 << (uint(v417) % 32)
	v421 = v412 + int32(16)
	v423 = v35 << (uint(v417) % 32)
	v424 = int32(0)
	v425 = base.B2i32(v423 == v424)
	if v425 == v424 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	base.MemoryCopy(m, v421, v37, v423)
	goto L141
L140:
	;
	goto L141
L141:
	;
	if v425 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	base.MemoryCopy(m, v421+v423, v37+v432<<(uint(int32(2))%32), v423)
	goto L144
L143:
	;
	goto L144
L144:
	;
	if l5 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = v357
	goto L147
L146:
	;
	goto L147
L147:
	;
	F_CopyArrayEls(m, v412, v89, v91, v357, v64, v61&int32(1), base.I32_extend8_s(v62), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	F_pfree(m, v89)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_pfree(m, v91)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v455 = v412
	goto L3
L151:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v487 = F_format_type_be(m, v34)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v487
	F_errmsg(m, int32(_a_F_array_replace_internal_8), v32)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_5), int32(_a_F_array_replace_internal_9), int32(_a_F_array_replace_internal_7))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_array_replace_internal_10), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_array_replace_internal_5), int32(_a_F_array_replace_internal_11), int32(_a_F_array_replace_internal_7))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_seek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	v7 = int32(0)
	if l2|base.B2i32(l4 <= v7) == v7 {
		switch l5 - int32(99) {
		case 0:
			v29 = l4
		case 1:
			v29 = (l4 + int32(7)) & int32(-8)
		default:
			v29 = (l4 + int32(1)) & int32(-2)
		case 6:
			v29 = (l4 + int32(3)) & int32(-4)
		}
		return l0 + l3*v29
	} else {
		if l2 == int32(0) {
			if l3 <= int32(0) {
				v195 = l0
			} else {
				v37 = int32(0)
				v44 = l0
				v46 = v37
				for {
					if base.B2i32(l4 <= v37) == int32(0) {
						v90 = v44 + l4
					} else {
						if base.B2i32(l4 != int32(-1)) == int32(0) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
							if v58 == int32(1) {
								v62 = int32(18)
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
								if v64 == v62 {
									v67 = v62
								} else {
									v67 = int32(2)
								}
								if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v74 = int32(6)
								} else {
									v74 = v67
								}
								v90 = v44 + v74
							} else {
								v76 = int32(1)
								if v58&v76 != 0 {
									v90 = v44 + int32(base.Ui32(v58)>>(uint(v76)%32))
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
									v90 = v44 + int32(base.Ui32(v81)>>(uint(int32(2))%32))
								}
							}
						} else {
							v85 = F_strlen(m, v44)
							mBase = m.M
							v90 = v85 + v44 + int32(1)
						}
					}
					switch l5 - int32(99) {
					case 0:
						v103 = v90
					case 1:
						v103 = (v90 + int32(7)) & int32(-8)
					default:
						v103 = (v90 + int32(1)) & int32(-2)
					case 6:
						v103 = (v90 + int32(3)) & int32(-4)
					}
					v105 = v46 + int32(1)
					if v105 != l3 {
						v44 = v103
						v46 = v105
						continue
					} else {
						break
					}
					break
				}
				v195 = v103
			}
		} else {
			if l3 <= int32(0) {
				v195 = l0
			} else {
				v110 = base.I32_div_s(l1, int32(8))
				v120 = l0
				v121 = int32(1) << (uint(l1&int32(7)) % 32)
				v122 = l2 + v110
				v128 = v7
				for {
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
					if v121&v129 == int32(0) {
						v183 = v120
					} else {
						if int32(0) < l4 {
							v170 = v120 + l4
						} else {
							if base.B2i32(l4 != int32(-1)) == int32(0) {
								v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
								if v138 == int32(1) {
									v142 = int32(18)
									v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
									if v144 == v142 {
										v147 = v142
									} else {
										v147 = int32(2)
									}
									if base.Ui32((v144-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v154 = int32(6)
									} else {
										v154 = v147
									}
									v170 = v120 + v154
								} else {
									v156 = int32(1)
									if v138&v156 != 0 {
										v170 = v120 + int32(base.Ui32(v138)>>(uint(v156)%32))
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
										v170 = v120 + int32(base.Ui32(v161)>>(uint(int32(2))%32))
									}
								}
							} else {
								v165 = F_strlen(m, v120)
								mBase = m.M
								v170 = v165 + v120 + int32(1)
							}
						}
						switch l5 - int32(99) {
						case 0:
							v183 = v170
						case 1:
							v183 = (v170 + int32(7)) & int32(-8)
						default:
							v183 = (v170 + int32(1)) & int32(-2)
						case 6:
							v183 = (v170 + int32(3)) & int32(-4)
						}
					}
					v185 = int32(1)
					v187 = v121 << (uint(v185) % 32)
					v189 = base.B2i32(v187 == int32(256))
					if v187 == int32(256) {
						v190 = v185
					} else {
						v190 = v187
					}
					v193 = v128 + int32(1)
					if v193 != l3 {
						v120 = v183
						v121 = v190
						v122 = v122 + v189
						v128 = v193
						continue
					} else {
						break
					}
					break
				}
				v195 = v183
			}
		}
		return v195
	}
}
func F_array_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v327 int32
	_ = v327
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_DatumGetAnyArrayP(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v27 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = int32(40)
	goto L5
L4:
	;
	v30 = int32(12)
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21+v30)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L77
	}
L7:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v80 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	F_get_type_io_data(m, v32, int32(3), v50+int32(4), v50+int32(6), v50+int32(7), v50+int32(8), v50+int32(12), v50+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v39 = F_MemoryContextAlloc(m, v37, int32(48))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v48 == v32 {
		v76 = v34
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v32 ^ int32(-1)
	v50 = v44
	goto L8
L13:
	;
	v50 = v34
	goto L8
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v66 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	F_fmgr_info_cxt(m, v66, v50+int32(20), v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v32
	v76 = v50
	goto L7
L17:
	;
	v83 = int32(28)
	goto L19
L18:
	;
	v83 = int32(4)
	goto L19
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v21+v83)))
	if v80 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76)+7)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+4)))
	v101 = F_ArrayGetNItemsSafe(m, v85, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v96 = v88
	v97 = v89
	goto L20
L22:
	;
	goto L23
L23:
	;
	v91 = v21 + int32(16)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v96 = v91
	v97 = v91 + v92<<(uint(int32(2))%32)
	goto L20
L24:
	;
	v104 = v18 + int32(32)
	F_pq_begintypsend(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_enlargeStringInfo(m, v104, int32(4))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v115 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v110+v111))) = base.I32_rotr(v85, int32(24))&v115 | base.I32_rotr(v85&v115, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v110 + int32(4)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v126 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v142 = v18 + int32(32)
	F_enlargeStringInfo(m, v142, int32(4))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
	}
L28:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	if v129 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v140 = base.B2i32(v137 != int32(0))
	goto L27
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v140 = base.B2i32(v130 != int32(0))
	goto L27
L32:
	;
	goto L33
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v140 = base.B2i32(v134 != int32(0))
	goto L27
L34:
	;
	v146 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	if v140 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v152 = int32(16777216)
	goto L37
L36:
	;
	v152 = v146
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147+v148))) = v152
	v154 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v147 + v154
	F_enlargeStringInfo(m, v142, v154)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v165 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v160+v161))) = base.I32_rotr(v32, int32(24))&v165 | base.I32_rotr(v32&v165, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v160 + int32(4)
	if int32(0) < v85 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v178 = v146
	goto L42
L40:
	;
	goto L41
L41:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v257 == int32(-1) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v194 = v178 << (uint(int32(2)) % 32)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v96+v194)))
	v198 = v18 + int32(32)
	F_enlargeStringInfo(m, v198, int32(4))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v207 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v202+v203))) = base.I32_rotr(v196, int32(24))&v207 | base.I32_rotr(v196&v207, int32(8))
	v215 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v202 + v215
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v194+v97)))
	F_enlargeStringInfo(m, v198, v215)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v228 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v223+v224))) = base.I32_rotr(v219, int32(24))&v228 | base.I32_rotr(v219&v228, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v223 + int32(4)
	v240 = v178 + int32(1)
	if v240 != v85 {
		v178 = v240
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(1)
	if int32(0) < v101 {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	if v260 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+12)) = int64(0)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v293 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v263 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v262
	v316 = v263
	goto L47
L52:
	;
	goto L53
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+12)) = int64(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+8))
	if v270 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v269 + (v273<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v316 = int32(0)
	goto L47
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v269 + v270
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v316 = v269 + v285<<(uint(int32(3))%32) + int32(16)
	goto L47
L57:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v21 + (v296<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v316 = int32(0)
	goto L47
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v293 + v21
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v316 = v21 + v308<<(uint(int32(3))%32) + int32(16)
	goto L47
L60:
	;
	v327 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	v426 = v18 + int32(32)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v429 << (uint(int32(2)) % 32)
	goto L76
L63:
	;
	v346 = F_array_iter_next(m, v18+int32(12), v18+int32(11), v327, v100, v99&int32(1), v98)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	goto L62
L65:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)))
	if v348 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v408 = v327 + int32(1)
	if v408 != v101 {
		v327 = v408
		goto L63
	} else {
		goto L75
	}
L67:
	;
	F_enlargeStringInfo(m, v18+int32(32), int32(4))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v364 = F_SendFunctionCall(m, v76+int32(20), v346)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v356+v357))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v356 + int32(4)
	goto L66
L71:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v368 = v18 + int32(32)
	F_enlargeStringInfo(m, v368, int32(4))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v375 = int32(2)
	v377 = int32(4)
	v378 = int32(base.Ui32(v366)>>(uint(v375)%32)) - v377
	v379 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v372+v373))) = base.I32_rotr(v378&v379, int32(8)) | base.I32_rotr(v378, int32(24))&v379
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v372 + v377
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	F_appendBinaryStringInfo(m, v368, v364+v377, int32(base.Ui32(v394)>>(uint(v375)%32))-v377)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v364)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L66
L75:
	;
	goto L64
L76:
	;
	m.G0 = v18 + int32(48)
	return v428
L77:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v444 = F_format_type_be(m, v32)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v444
	F_errmsg(m, int32(_a_F_array_send_0), v18)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_array_send_1), int32(1589), int32(_a_F_array_send_2))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v10 = F_array_set_element(m, l0, int32(1), l1, l2, int32(0), int32(-1), l3, l4, int32(105))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_pg_detoast_datum(m, v10)
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_array_shuffle_n(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v122 int64
	_ = v122
	var v127 int64
	_ = v127
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	v6 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(80)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v27 <= v6)|base.B2i32(l1 <= v6) == v6 {
		v36 = l0 + int32(16)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
		if int32(0) < v37 {
			v46 = v27 << (uint(int32(2)) % 32)
			v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l4)+8)))
			v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+10)))
			v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4)+11)))
			F_deconstruct_array(m, l0, v48, v49&int32(1), v52, v25+int32(12), v25+int32(8), v25+int32(76))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v63 = base.I32_div_s(v61, v62)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+76)) = v63
				v67 = base.I64_extend_i32_s(v62 - int32(1))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				v77 = v69
				v78 = v70
				v89 = int64(0)
				for {
					v93 = int32(_a_F_array_shuffle_n_0)
					if base.Ui64(v67) <= base.Ui64(v89) {
						v140 = v89
					} else {
						v100 = v67 - v89
						v102 = *(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[0]))
						v103 = *(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[1]))
						v106 = v103
						v108 = v102
						for {
							v112 = v106 ^ v108
							v114 = base.I64_rotl(v112, int64(37))
							v122 = v112 ^ (v112<<(uint(int64(16))%64) ^ base.I64_rotl(v106, int64(24)))
							v127 = int64(base.Ui64(base.I64_rotl(v106*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v100)) % 64))
							if base.Ui64(v100) < base.Ui64(v127) {
								v106 = v122
								v108 = v114
								continue
							} else {
								break
							}
							break
						}
						*(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[0])) = v114
						*(*int64)(unsafe.Add(mBase, _c_F_array_shuffle_n[1])) = v122
						v140 = v89 + v127
					}
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
					if int32(0) < v141 {
						v145 = v141 * base.I32_wrap_i64(v140)
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
						v148 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
						v153 = v145 + v146
						v157 = v148 + v145<<(uint(int32(2))%32)
						v159 = v77
						v160 = v78
						v164 = int32(0)
						for {
							v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
							v176 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
							v177 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
							*(*int32)(unsafe.Add(mBase, uint32(v160))) = v177
							v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
							*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v179)
							*(*int32)(unsafe.Add(mBase, uint32(v157))) = v176
							*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v175)
							v183 = int32(1)
							v185 = int32(4)
							v188 = v159 + v183
							v190 = v160 + v185
							v192 = v164 + v183
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
							if v192 < v193 {
								v153 = v153 + v183
								v157 = v157 + v185
								v159 = v188
								v160 = v190
								v164 = v192
								continue
							} else {
								break
							}
							break
						}
						v201 = v188
						v202 = v190
					} else {
						v201 = v77
						v202 = v78
					}
					v218 = v89 + int64(1)
					if v218 != base.I64_extend_i32_u(l1) {
						v77 = v201
						v78 = v202
						v89 = v218
						continue
					} else {
						break
					}
					break
				}
				v220 = int32(0)
				v221 = base.B2i32(v46 == v220)
				if v221 == v220 {
					base.MemoryCopy(m, v25+int32(48), v36, v46)
				} else {
				}
				if v221 == int32(0) {
					base.MemoryCopy(m, v25+int32(16), v36+v46, v46)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = l1
				if l2 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(1)
				} else {
				}
				v237 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				v238 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v245 = F_construct_md_array(m, v237, v238, v27, v25+int32(48), v25+int32(16), l3, v48, v49&int32(1), v52)
				mBase = m.M
				v246 = m.ExcPending
				if v246 != 0 {
					return int32(0)
				} else {
					v247 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
					F_pfree(m, v247)
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return int32(0)
					} else {
						v250 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
						F_pfree(m, v250)
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							v253 = v245
							m.G0 = v25 + int32(80)
							return v253
						}
					}
				}
			}
		} else {
			v41 = F_construct_empty_array(m, l3)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v253 = v41
				m.G0 = v25 + int32(80)
				return v253
			}
		}
	} else {
		v41 = F_construct_empty_array(m, l3)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v253 = v41
			m.G0 = v25 + int32(80)
			return v253
		}
	}
}
func F_array_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = F_array_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 < int32(0) {
			v10 = int32(20)
		} else {
			v10 = int32(28)
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_array_subscript_fetch(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+4)))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+6)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)))
	v14 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+9)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = F_array_get_element(m, v5, v7, v8+int32(12), v11, v12, v13, v14, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v16
		return
	}
}
func F_array_subscript_handler_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(464) {
		v26 = v2
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 == int32(0) {
			v26 = v2
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v14 != int32(8) {
				v26 = v2
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				if v17 != 0 {
					v26 = v2
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v18 != v19 {
						v26 = v2
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						if v22 != 0 {
							v23 = v11
						} else {
							v23 = int32(0)
						}
						v26 = v23
					}
				}
			}
		}
	}
	return v26
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
