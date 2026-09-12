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
	var v19 int32
	_ = v19
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
	v19 = v4
	goto L3
L3:
	;
	v24 = v19 << (uint(int32(2)) % 32)
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
	v35 = v19 + int32(1)
	if l0 != v35 {
		v19 = v35
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
	F_errmsg(m, int32(469962), v11)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, int32(0), int32(477124), int32(141), int32(395459))
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	v5 = int32(0)
	v14 = l0 - int32(1)
	if v14 < v5 {
		v92 = v5
	} else {
		v17 = int32(1)
		if v14 == int32(0) {
			v69 = v5
			v70 = v14
			v71 = v17
		} else {
			v28 = v5
			v29 = v14
			v30 = v17
			v31 = v5
			for {
				v36 = int32(2)
				v37 = v29 << (uint(v36) % 32)
				v39 = v37 - int32(4)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l3+v39)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l2+v39)))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v37+l1)))
				v47 = v46 * v30
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v37+l3)))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v37+l2)))
				v56 = (v41-v43)*v47 + ((v50-v52)*v30 + v28)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l1+v39)))
				v59 = v58 * v47
				v61 = v29 - v36
				v63 = v31 + v36
				if v63 != l0&int32(-2) {
					v28 = v56
					v29 = v61
					v30 = v59
					v31 = v63
					continue
				} else {
					break
				}
				break
			}
			v69 = v56
			v70 = v61
			v71 = v59
		}
		if l0&v17 == int32(0) {
			v92 = v69
		} else {
			v80 = v70 << (uint(int32(2)) % 32)
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l3+v80)))
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+l2)))
			v92 = (v82-v84)*v71 + v69
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
	if v102 == int32(0) {
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
	F_errmsg_internal(m, int32(427877), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(477782), int32(984), int32(146509))
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
	if v101 == int32(1) {
		goto L5
	} else {
		goto L28
	}
L28:
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
					F_errmsg_internal(m, int32(59442), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(477889), int32(953), int32(270036))
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
					F_errmsg(m, int32(357798), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(477889), int32(942), int32(270036))
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
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
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v306 int32
	_ = v306
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_DatumGetAnyArrayP(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = F_DatumGetAnyArrayP(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v34 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v37 = int32(28)
	goto L6
L5:
	;
	v37 = int32(4)
	goto L6
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v41 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = int32(28)
	goto L9
L8:
	;
	v44 = int32(4)
	goto L9
L9:
	;
	if v41 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v30+v37)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25+v44)))
	if v34 == int32(-1) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v51 = v48
	goto L10
L12:
	;
	goto L13
L13:
	;
	v51 = v25 + int32(16)
	goto L10
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v61 = F_ArrayGetNItems(m, v53, v51)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	v59 = v56
	goto L14
L16:
	;
	goto L17
L17:
	;
	v59 = v30 + int32(16)
	goto L14
L18:
	;
	v63 = F_ArrayGetNItems(m, v52, v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v67 == int32(-1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v70 = int32(40)
	goto L22
L21:
	;
	v70 = int32(12)
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v25+v70)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v75 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = int32(40)
	goto L25
L24:
	;
	v78 = int32(12)
	goto L25
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v30+v78)))
	if v72 == v80 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	if v83 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	goto L28
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L140
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L135
	}
L30:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95)+11)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+10)))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95)+8)))
	v100 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+78)) = uint16(v100)
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+76)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v95 + int32(104)
	v110 = base.B2i32(v61 < v63)
	if v96 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v84 == v72 {
		v95 = v83
		v96 = v67
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v87 = F_lookup_type_cache(m, v72, int32(64))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+108))
	if v89 == int32(0) {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v87
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v95 = v87
	v96 = v94
	goto L30
L37:
	;
	if v61 < v63 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	if v113 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = int64(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v146 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v115
	v169 = v116
	goto L37
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = int64(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v123 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v122 + (v126<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v169 = int32(0)
	goto L37
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v122 + v123
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v169 = v122 + v138<<(uint(int32(3))%32) + int32(16)
	goto L37
L47:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v25 + (v149<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v169 = int32(0)
	goto L37
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v146 + v25
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v169 = v25 + v161<<(uint(int32(3))%32) + int32(16)
	goto L37
L50:
	;
	v170 = v61
	goto L52
L51:
	;
	v170 = v63
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v169
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v174 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(1)
	if int32(0) < v170 {
		goto L67
	} else {
		goto L68
	}
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	if v177 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+20)) = int64(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v210 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
	v180 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v179
	v233 = v180
	goto L53
L58:
	;
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+20)) = int64(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if v187 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v186 + (v190<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v233 = int32(0)
	goto L53
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v186 + v187
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	v233 = v186 + v202<<(uint(int32(3))%32) + int32(16)
	goto L53
L63:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v30 + (v213<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v233 = int32(0)
	goto L53
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v210 + v30
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v233 = v30 + v225<<(uint(int32(3))%32) + int32(16)
	goto L53
L66:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v469 == int32(-1) {
		goto L127
	} else {
		goto L128
	}
L67:
	;
	v240 = v98 & int32(1)
	v243 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	if v61 != v63 {
		goto L85
	} else {
		goto L86
	}
L70:
	;
	v265 = F_array_iter_next(m, v22+int32(40), v22+int32(19), v243, v99, v240, v97)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	v271 = F_array_iter_next(m, v22+int32(20), v22+int32(18), v243, v99, v240, v97)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+19)))
	if v274 == int32(1) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v306 = v243 + int32(1)
	if v306 != v170 {
		v243 = v306
		goto L70
	} else {
		goto L84
	}
L75:
	;
	v277 = int32(1)
	if v273&v277 == int32(0) {
		v457 = v277
		goto L66
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if v273&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L74
L79:
	;
	v457 = int32(-1)
	goto L66
L80:
	;
	goto L81
L81:
	;
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+92)) = uint8(v285)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v271
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+84)) = uint8(v285)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v265
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v295 = m.T0[v294].(func(*base.Module, int32) int32)(m, v22+int32(60))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v295 == int32(0) {
		goto L74
	} else {
		goto L83
	}
L83:
	;
	v457 = v295>>(uint(int32(31))%32) | int32(1)
	goto L66
L84:
	;
	goto L71
L85:
	;
	if v61 < v63 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v53 == v52 {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v330 = int32(-1)
	goto L90
L89:
	;
	v330 = int32(1)
	goto L90
L90:
	;
	v457 = v330
	goto L66
L91:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v391 == int32(-1) {
		goto L110
	} else {
		goto L111
	}
L92:
	;
	v340 = v332
	goto L100
L93:
	;
	v332 = int32(0)
	if v53 <= v332 {
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v53 < v52 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L92
L97:
	;
	v338 = int32(-1)
	goto L99
L98:
	;
	v338 = int32(1)
	goto L99
L99:
	;
	v457 = v338
	goto L66
L100:
	;
	v359 = v340 << (uint(int32(2)) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v51+v359)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+v59)))
	if v361 == v363 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v361 < v363 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	v366 = v340 + int32(1)
	if v53 != v366 {
		v340 = v366
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	goto L101
L105:
	;
	goto L91
L106:
	;
	v371 = int32(-1)
	goto L108
L107:
	;
	v371 = int32(1)
	goto L108
L108:
	;
	v457 = v371
	goto L66
L109:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v402 == int32(-1) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	v401 = v394
	goto L109
L111:
	;
	goto L112
L112:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v401 = v25 + v395<<(uint(int32(2))%32) + int32(16)
	goto L109
L113:
	;
	v413 = int32(0)
	if v53 <= v413 {
		v457 = v413
		goto L66
	} else {
		goto L117
	}
L114:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v30)+36))
	v412 = v405
	goto L113
L115:
	;
	goto L116
L116:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v412 = v30 + v406<<(uint(int32(2))%32) + int32(16)
	goto L113
L117:
	;
	v418 = int32(0)
	goto L118
L118:
	;
	v437 = v418 << (uint(int32(2)) % 32)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v401+v437)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v412+v437)))
	if v439 == v441 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v439 < v441 {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	v444 = v418 + int32(1)
	if v53 != v444 {
		v418 = v444
		goto L118
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	v457 = v413
	goto L66
L124:
	;
	v449 = int32(-1)
	goto L126
L125:
	;
	v449 = int32(1)
	goto L126
L126:
	;
	v457 = v449
	goto L66
L127:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v476 == int32(-1) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v25 == v472 {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	F_pfree(m, v25)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	goto L127
L131:
	;
	m.G0 = v22 + int32(96)
	return v457
L132:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v30 == v479 {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v30)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v494 = F_format_type_be(m, v72)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v494
	F_errmsg(m, int32(182124), v22)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(477782), int32(4035), int32(227138))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(155090), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(477782), int32(4017), int32(227138))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v190
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+l2))))
	v21 = v19
	goto L7
L6:
	;
	v21 = int32(0)
	goto L7
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13+l2<<(uint(int32(2))%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v21)
	v190 = v22
	goto L1
L8:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v178 = v176 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v178
	if v178 != int32(256) {
		v190 = v174
		goto L1
	} else {
		goto L65
	}
L9:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l4 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27&v28 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
	v174 = int32(0)
	goto L8
L12:
	;
	switch l5 - int32(99) {
	case 0:
		v171 = v155
		goto L61
	case 1:
		goto L63
	default:
		goto L62
	case 6:
		goto L64
	}
L13:
	;
	switch l3 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	default:
		goto L16
	case 3:
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if int32(0) < l3 {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v155 = l3 + v35
	v156 = v42
	goto L12
L18:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35))))
	v155 = l3 + v35
	v156 = v40
	goto L12
L19:
	;
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35))))
	v155 = l3 + v35
	v156 = v38
	goto L12
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
	F_errmsg_internal(m, int32(466883), v11)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(314610), int32(70), int32(65727))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v155 = l3 + v35
	v156 = v35
	goto L12
L25:
	;
	goto L26
L26:
	;
	if l3 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v155 = v154
	v156 = v35
	goto L12
L28:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v64 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	if v35&int32(3) == int32(0) {
		v116 = v35
		goto L46
	} else {
		goto L47
	}
L31:
	;
	v67 = int32(6)
	v69 = int32(18)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v71 == v69 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v84 = int32(1)
	if v64&v84 != 0 {
		v154 = v35 + int32(base.Ui32(v64)>>(uint(v84)%32))
		goto L27
	} else {
		goto L43
	}
L34:
	;
	v74 = v69
	goto L36
L35:
	;
	v74 = int32(2)
	goto L36
L36:
	;
	if v71&int32(254) == int32(2) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v79 = v67
	goto L39
L38:
	;
	v79 = v74
	goto L39
L39:
	;
	if v71 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v82 = v67
	goto L42
L41:
	;
	v82 = v79
	goto L42
L42:
	;
	v154 = v35 + v82
	goto L27
L43:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v154 = v35 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
	goto L27
L44:
	;
	v154 = v149 + v35 + int32(1)
	goto L27
L45:
	;
	v149 = v141 - v35
	goto L44
L46:
	;
	v120 = v116
	goto L55
L47:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v100 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v149 = int32(0)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v105 = v35
	goto L51
L51:
	;
	v109 = v105 + int32(1)
	if v109&int32(3) == int32(0) {
		v116 = v109
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v141 = v109
	goto L45
L53:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v114 != 0 {
		v105 = v109
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v129 = int32(-2139062144)
	if (int32(16843008)-v126|v126)&v129 == v129 {
		v120 = v120 + int32(4)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v135 = v120
	goto L58
L57:
	;
	goto L56
L58:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 != 0 {
		v135 = v135 + int32(1)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v141 = v135
	goto L45
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v171
	v174 = v156
	goto L8
L62:
	;
	v171 = (v155 + int32(1)) & int32(-2)
	goto L61
L63:
	;
	v171 = (v155 + int32(7)) & int32(-8)
	goto L61
L64:
	;
	v171 = (v155 + int32(3)) & int32(-4)
	goto L61
L65:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v182 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v182 + int32(1)
	goto L68
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	v190 = v174
	goto L1
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
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
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v193 {
		goto L62
	} else {
		goto L63
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
	switch v48&int32(65535) - int32(1) {
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
	F_errmsg_internal(m, int32(466883), v15)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(314610), int32(70), int32(65727))
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
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	switch v171 - int32(99) {
	case 0:
		v188 = v170
		goto L58
	case 1:
		goto L60
	default:
		goto L59
	case 6:
		goto L61
	}
L22:
	;
	v170 = v42 + v74
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
	if v42&int32(3) == int32(0) {
		v132 = v42
		goto L43
	} else {
		goto L44
	}
L28:
	;
	v83 = int32(6)
	v85 = int32(18)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v87 == v85 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v100 = int32(1)
	if v80&v100 != 0 {
		v170 = v42 + int32(base.Ui32(v80)>>(uint(v100)%32))
		goto L21
	} else {
		goto L40
	}
L31:
	;
	v90 = v85
	goto L33
L32:
	;
	v90 = int32(2)
	goto L33
L33:
	;
	if v87&int32(254) == int32(2) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v95 = v83
	goto L36
L35:
	;
	v95 = v90
	goto L36
L36:
	;
	if v87 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v98 = v83
	goto L39
L38:
	;
	v98 = v95
	goto L39
L39:
	;
	v170 = v42 + v98
	goto L21
L40:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v170 = v42 + int32(base.Ui32(v105)>>(uint(int32(2))%32))
	goto L21
L41:
	;
	v170 = v165 + v42 + int32(1)
	goto L21
L42:
	;
	v165 = v157 - v42
	goto L41
L43:
	;
	v136 = v132
	goto L52
L44:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v116 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v165 = int32(0)
	goto L41
L46:
	;
	goto L47
L47:
	;
	v121 = v42
	goto L48
L48:
	;
	v125 = v121 + int32(1)
	if v125&int32(3) == int32(0) {
		v132 = v125
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v157 = v125
	goto L42
L50:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v130 != 0 {
		v121 = v125
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v145 = int32(-2139062144)
	if (int32(16843008)-v142|v142)&v145 == v145 {
		v136 = v136 + int32(4)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v151 = v136
	goto L55
L54:
	;
	goto L53
L55:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 != 0 {
		v151 = v151 + int32(1)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v157 = v151
	goto L42
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v188
	goto L1
L59:
	;
	v188 = (v170 + int32(1)) & int32(-2)
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = (v170 + int32(7)) & int32(-8)
	goto L1
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = (v170 + int32(3)) & int32(-4)
	goto L1
L62:
	;
	v200 = int32(0)
	v201 = v190
	goto L65
L63:
	;
	v389 = v20
	v390 = v190
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v390
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v403 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v405 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+15)))
	v406 = F_construct_md_array(m, v192, v191, v389, v399, v400, v402, v403, v404, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L17
	} else {
		goto L123
	}
L65:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v209 + int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v213 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v389 = v385
	v390 = v379
	goto L64
L67:
	;
	v382 = v200 + int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v382 < v383 {
		v200 = v382
		v201 = v379
		goto L65
	} else {
		goto L122
	}
L68:
	;
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200+v191))) = uint8(v234)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v239 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v217 = base.I32_div_s(v209, int32(8))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213+v217))))
	if int32(base.Ui32(v219)>>(uint(v209&int32(7))%32))&int32(1) != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v200+v191))) = uint8(v226)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v200<<(uint(int32(2))%32)))) = int32(0)
	v379 = v201
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192+v200<<(uint(int32(2))%32)))) = v265
	v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v267 {
		goto L83
	} else {
		goto L84
	}
L72:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	switch v242 - int32(1) {
	case 0:
		goto L78
	case 1:
		goto L77
	default:
		goto L75
	case 3:
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v265 = v201
	goto L71
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L17
	} else {
		goto L79
	}
L76:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v265 = v247
	goto L71
L77:
	;
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v201))))
	v265 = v246
	goto L71
L78:
	;
	v245 = int32(*(*int8)(unsafe.Add(mBase, uint32(v201))))
	v265 = v245
	goto L71
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = base.I32_extend16_s(v242)
	F_errmsg_internal(m, int32(466883), v15+int32(16))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(314610), int32(70), int32(65727))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L17
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	switch v364 - int32(99) {
	case 0:
		v379 = v363
		goto L67
	case 1:
		goto L120
	default:
		goto L119
	case 6:
		goto L121
	}
L83:
	;
	v363 = v201 + v267
	goto L82
L84:
	;
	goto L85
L85:
	;
	if v267 == int32(-1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v273 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if v201&int32(3) == int32(0) {
		v325 = v201
		goto L104
	} else {
		goto L105
	}
L89:
	;
	v276 = int32(6)
	v278 = int32(18)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	if v280 == v278 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v293 = int32(1)
	if v273&v293 != 0 {
		v363 = v201 + int32(base.Ui32(v273)>>(uint(v293)%32))
		goto L82
	} else {
		goto L101
	}
L92:
	;
	v283 = v278
	goto L94
L93:
	;
	v283 = int32(2)
	goto L94
L94:
	;
	if v280&int32(254) == int32(2) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v288 = v276
	goto L97
L96:
	;
	v288 = v283
	goto L97
L97:
	;
	if v280 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v291 = v276
	goto L100
L99:
	;
	v291 = v288
	goto L100
L100:
	;
	v363 = v201 + v291
	goto L82
L101:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v363 = v201 + int32(base.Ui32(v298)>>(uint(int32(2))%32))
	goto L82
L102:
	;
	v363 = v358 + v201 + int32(1)
	goto L82
L103:
	;
	v358 = v350 - v201
	goto L102
L104:
	;
	v329 = v325
	goto L113
L105:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v309 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v358 = int32(0)
	goto L102
L107:
	;
	goto L108
L108:
	;
	v314 = v201
	goto L109
L109:
	;
	v318 = v314 + int32(1)
	if v318&int32(3) == int32(0) {
		v325 = v318
		goto L104
	} else {
		goto L111
	}
L110:
	;
	v350 = v318
	goto L103
L111:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v323 != 0 {
		v314 = v318
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v338 = int32(-2139062144)
	if (int32(16843008)-v335|v335)&v338 == v338 {
		v329 = v329 + int32(4)
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v344 = v329
	goto L116
L115:
	;
	goto L114
L116:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v348 != 0 {
		v344 = v344 + int32(1)
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v350 = v344
	goto L103
L118:
	;
	goto L117
L119:
	;
	v379 = (v363 + int32(1)) & int32(-2)
	goto L67
L120:
	;
	v379 = (v363 + int32(7)) & int32(-8)
	goto L67
L121:
	;
	v379 = (v363 + int32(3)) & int32(-4)
	goto L67
L122:
	;
	goto L66
L123:
	;
	v408 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v408)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v406
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
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
	return v214
L2:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
	v214 = int32(0)
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v200 = m.ExcPending
	if v200 != 0 {
		goto L5
	} else {
		goto L66
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
	v40 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v214 = v41
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
	v51 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v214 = v52
	goto L1
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L61
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
	v164 = v34
	goto L40
L40:
	;
	F_array_free_iterator(m, v105)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L55
	}
L41:
	;
	v128 = int32(1)
	v129 = v118 + v128
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v43|v130&v128 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v164 = v151
	goto L40
L43:
	;
	v156 = F_array_iterate(m, v105, v14+int32(12), v14+int32(11))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L53
	}
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v149 = F_accumArrayResult(m, v123, v129, int32(0), int32(23), v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L52
	}
L45:
	;
	if v43 == int32(0) {
		v151 = v123
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v141 = F_FunctionCall2Coll(m, v102+int32(20), v21, v56, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	if v130&int32(1) == int32(0) {
		v151 = v123
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	if v141 == int32(0) {
		v151 = v123
		goto L43
	} else {
		goto L51
	}
L51:
	;
	goto L44
L52:
	;
	v151 = v149
	goto L43
L53:
	;
	if v156 != 0 {
		v118 = v129
		v123 = v151
		goto L41
	} else {
		goto L54
	}
L54:
	;
	goto L42
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v171 != v23 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_pfree(m, v23)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v177 = F_makeArrayResult(m, v164, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v214 = v177
	goto L1
L61:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v186 = F_format_type_be(m, v61)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v186
	F_errmsg(m, int32(181913), v14)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(477889), int32(1554), int32(132003))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(426615), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(477889), int32(1502), int32(132003))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
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
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
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
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	v5 = l4
	v9 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(80)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = l0 + int32(16)
	v38 = F_ArrayGetNItems(m, v35, v37)
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
	v619 = m.ExcPending
	if v619 != 0 {
		goto L4
	} else {
		goto L194
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L4
	} else {
		goto L189
	}
L3:
	;
	m.G0 = v32 + int32(80)
	return v573
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
	v573 = l0
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
	if v63 != int32(65535) {
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
	v111 = int32(0)
	v119 = v111
	v120 = v106
	v121 = v99
	v123 = l0 + v104
	v126 = v111
	v131 = v9
	v135 = v9
	v138 = v9
	goto L36
L36:
	;
	if v121 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if v479 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L38:
	;
	v482 = int32(1)
	v484 = v120 << (uint(v482) % 32)
	v486 = base.B2i32(v484 == int32(256))
	if v484 == int32(256) {
		goto L149
	} else {
		goto L150
	}
L39:
	;
	v474 = v465
	v476 = v467
	v478 = v126 + int32(1)
	v479 = v469
	v480 = v470
	goto L38
L40:
	;
	if int32(0) < v64 {
		v428 = v64
		goto L108
	} else {
		goto L109
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v126<<(uint(int32(2))%32)))) = v321
	v329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v91))) = uint8(v329)
	v332 = v320
	v334 = v322
	goto L40
L42:
	;
	v320 = v285
	v321 = v272
	v322 = v131
	goto L41
L43:
	;
	v314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v91))) = uint8(v314)
	v465 = v119
	v467 = v123
	v469 = v131
	v470 = v314
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
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v120&v144 != 0 {
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
	v474 = v119
	v476 = v123
	v478 = v126
	v479 = int32(1)
	v480 = v135
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
	v320 = v123
	v321 = v75
	v322 = int32(1)
	goto L41
L52:
	;
	switch v109 {
	case 0:
		v285 = v271
		goto L96
	case 1:
		goto L98
	default:
		goto L97
	case 6:
		goto L99
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
		goto L63
	} else {
		goto L64
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L60
	}
L57:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v271 = v123 + v64
	v272 = v156
	goto L52
L58:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123))))
	v271 = v123 + v64
	v272 = v154
	goto L52
L59:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
	v271 = v123 + v64
	v272 = v152
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v64
	F_errmsg_internal(m, int32(466883), v32+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(314610), int32(70), int32(65727))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
	v271 = v123 + v64
	v272 = v123
	goto L52
L64:
	;
	goto L65
L65:
	;
	if v64 == int32(-1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v271 = v270
	v272 = v123
	goto L52
L67:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v178 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	if v123&int32(3) == int32(0) {
		v231 = v123
		goto L81
	} else {
		goto L82
	}
L70:
	;
	v270 = v206 + v123
	goto L66
L71:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	if base.Ui32((v182-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v206 = int32(6)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v197 = int32(1)
	if v178&v197 != 0 {
		v270 = v123 + int32(base.Ui32(v178)>>(uint(v197)%32))
		goto L66
	} else {
		goto L78
	}
L74:
	;
	v189 = int32(18)
	if v182&int32(255) == v189 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v195 = v189
	goto L77
L76:
	;
	v195 = int32(2)
	goto L77
L77:
	;
	v270 = v123 + v195
	goto L66
L78:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v206 = int32(base.Ui32(v202) >> (uint(int32(2)) % 32))
	goto L70
L79:
	;
	v270 = v264 + v123 + int32(1)
	goto L66
L80:
	;
	v264 = v256 - v123
	goto L79
L81:
	;
	v235 = v231
	goto L90
L82:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v215 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v264 = int32(0)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v220 = v123
	goto L86
L86:
	;
	v224 = v220 + int32(1)
	if v224&int32(3) == int32(0) {
		v231 = v224
		goto L81
	} else {
		goto L88
	}
L87:
	;
	v256 = v224
	goto L80
L88:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v229 != 0 {
		v220 = v224
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v244 = int32(-2139062144)
	if (int32(16843008)-v241|v241)&v244 == v244 {
		v235 = v235 + int32(4)
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v250 = v235
	goto L93
L92:
	;
	goto L91
L93:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v254 != 0 {
		v250 = v250 + int32(1)
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v256 = v250
	goto L80
L95:
	;
	goto L94
L96:
	;
	if l2 != 0 {
		goto L42
	} else {
		goto L100
	}
L97:
	;
	v285 = (v271 + int32(1)) & int32(-2)
	goto L96
L98:
	;
	v285 = (v271 + int32(7)) & int32(-8)
	goto L96
L99:
	;
	v285 = (v271 + int32(3)) & int32(-4)
	goto L96
L100:
	;
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v286)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v74
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v286)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v272
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+60)) = uint8(v286)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v298 = m.T0[v297].(func(*base.Module, int32) int32)(m, v32+int32(44))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+60)))
	if v300 != 0 {
		goto L42
	} else {
		goto L102
	}
L102:
	;
	if v298 == int32(0) {
		v320 = v285
		v321 = v272
		v322 = v131
		goto L41
	} else {
		goto L103
	}
L103:
	;
	if l5 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v474 = v119
	v476 = v285
	v478 = v126
	v479 = int32(1)
	v480 = v135
	goto L38
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v126<<(uint(int32(2))%32)))) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v91))) = uint8(v5)
	v310 = int32(1)
	if v5 == int32(0) {
		v332 = v285
		v334 = v310
		goto L40
	} else {
		goto L107
	}
L107:
	;
	v465 = v119
	v467 = v285
	v469 = v310
	v470 = int32(1)
	goto L39
L108:
	;
	v430 = v119 + v428
	switch v109 {
	case 0:
		v443 = v430
		goto L140
	case 1:
		goto L142
	default:
		goto L141
	case 6:
		goto L143
	}
L109:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v89+v126<<(uint(int32(2))%32))))
	if v64 == int32(-1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v344 == int32(1) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v341&int32(3) == int32(0) {
		v392 = v341
		goto L125
	} else {
		goto L126
	}
L113:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	if base.Ui32((v348-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v428 = int32(6)
		goto L108
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v344&int32(1) != 0 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	v355 = int32(18)
	if v348&int32(255) == v355 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v361 = v355
	goto L119
L118:
	;
	v361 = int32(2)
	goto L119
L119:
	;
	v428 = v361
	goto L108
L120:
	;
	v428 = int32(base.Ui32(v344) >> (uint(int32(1)) % 32))
	goto L108
L121:
	;
	goto L122
L122:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v428 = int32(base.Ui32(v366) >> (uint(int32(2)) % 32))
	goto L108
L123:
	;
	v428 = v425 + int32(1)
	goto L108
L124:
	;
	v425 = v417 - v341
	goto L123
L125:
	;
	v396 = v392
	goto L134
L126:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v376 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v425 = int32(0)
	goto L123
L128:
	;
	goto L129
L129:
	;
	v381 = v341
	goto L130
L130:
	;
	v385 = v381 + int32(1)
	if v385&int32(3) == int32(0) {
		v392 = v385
		goto L125
	} else {
		goto L132
	}
L131:
	;
	v417 = v385
	goto L124
L132:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v390 != 0 {
		v381 = v385
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v405 = int32(-2139062144)
	if (int32(16843008)-v402|v402)&v405 == v405 {
		v396 = v396 + int32(4)
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v411 = v396
	goto L137
L136:
	;
	goto L135
L137:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v415 != 0 {
		v411 = v411 + int32(1)
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v417 = v411
	goto L124
L139:
	;
	goto L138
L140:
	;
	if base.Ui32(v443) < base.Ui32(int32(1073741824)) {
		v465 = v443
		v467 = v332
		v469 = v334
		v470 = v135
		goto L39
	} else {
		goto L144
	}
L141:
	;
	v443 = (v430 + int32(1)) & int32(-2)
	goto L140
L142:
	;
	v443 = (v430 + int32(7)) & int32(-8)
	goto L140
L143:
	;
	v443 = (v430 + int32(3)) & int32(-4)
	goto L140
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(1073741823)
	F_errmsg(m, int32(645032), v32+int32(32))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(477782), int32(6567), int32(300997))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	v487 = v482
	goto L151
L150:
	;
	v487 = v484
	goto L151
L151:
	;
	if v121 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v488 = v487
	goto L154
L153:
	;
	v488 = v120
	goto L154
L154:
	;
	if v121 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v491 = v121 + v486
	goto L157
L156:
	;
	v491 = int32(0)
	goto L157
L157:
	;
	v493 = v138 + int32(1)
	if v493 != v38 {
		v119 = v474
		v120 = v488
		v121 = v491
		v123 = v476
		v126 = v478
		v131 = v479
		v135 = v480
		v138 = v493
		goto L36
	} else {
		goto L158
	}
L158:
	;
	goto L37
L159:
	;
	F_pfree(m, v89)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	if v478 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	F_pfree(m, v91)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v573 = l0
	goto L3
L164:
	;
	F_pfree(m, v89)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v516 = v35 << (uint(int32(3)) % 32)
	if v480 != 0 {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	F_pfree(m, v91)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v508 = F_palloc0(m, int32(16))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v508)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v508)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v508))) = int64(64)
	v573 = v508
	goto L3
L170:
	;
	v533 = v474 + v531
	v534 = F_palloc0(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L174
	}
L171:
	;
	v520 = base.I32_div_s(v478+int32(7), int32(8))
	v525 = (v516 + v520 + int32(23)) & int32(-8)
	v531 = v525
	v532 = v525
	goto L170
L172:
	;
	goto L173
L173:
	;
	v531 = (v516 + int32(23)) & int32(-8)
	v532 = int32(0)
	goto L170
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v534)+4)) = v35
	v539 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = v533 << (uint(v539) % 32)
	v543 = v534 + int32(16)
	v545 = v35 << (uint(v539) % 32)
	if v545 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v545 != 0 {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	v546 = F__emscripten_memcpy_bulkmem(m, v543, v37, v545)
	mBase = m.M
	v547 = v546
	goto L178
L177:
	;
	v547 = v543
	goto L178
L178:
	;
	goto L175
L179:
	;
	if l5 != 0 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v553 = F__emscripten_memcpy_bulkmem(m, v547+v545, v37+v549<<(uint(int32(2))%32), v545)
	mBase = m.M
	goto L182
L181:
	;
	goto L182
L182:
	;
	goto L179
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v478
	goto L185
L184:
	;
	goto L185
L185:
	;
	F_CopyArrayEls(m, v534, v89, v91, v478, v64, v61&int32(1), base.I32_extend8_s(v62), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	F_pfree(m, v89)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_pfree(m, v91)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	v573 = v534
	goto L3
L189:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v605 = F_format_type_be(m, v34)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v605
	F_errmsg(m, int32(181913), v32)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(477782), int32(6447), int32(300997))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	F_errmsg(m, int32(426682), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(477782), int32(6431), int32(300997))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_seek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	if l4 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	switch l5 - int32(99) {
	case 0:
		v26 = l4
		goto L4
	case 1:
		goto L6
	default:
		goto L5
	case 6:
		goto L7
	}
L4:
	;
	return l0 + l3*v26
L5:
	;
	v26 = (l4 + int32(1)) & int32(-2)
	goto L4
L6:
	;
	v26 = (l4 + int32(7)) & int32(-8)
	goto L4
L7:
	;
	v26 = (l4 + int32(3)) & int32(-4)
	goto L4
L8:
	;
	return v308
L9:
	;
	if l3 <= int32(0) {
		v308 = l0
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if l3 <= int32(0) {
		v308 = l0
		goto L8
	} else {
		goto L55
	}
L12:
	;
	v34 = int32(0)
	v41 = l0
	v43 = v34
	goto L13
L13:
	;
	if base.B2i32(l4 <= v34) == int32(0) {
		v145 = v41 + l4
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v308 = v158
	goto L8
L15:
	;
	switch l5 - int32(99) {
	case 0:
		v158 = v145
		goto L50
	case 1:
		goto L52
	default:
		goto L51
	case 6:
		goto L53
	}
L16:
	;
	if base.B2i32(l4 != int32(-1)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v55 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if v41&int32(3) == int32(0) {
		v107 = v41
		goto L35
	} else {
		goto L36
	}
L20:
	;
	v58 = int32(6)
	v60 = int32(18)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v62 == v60 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v75 = int32(1)
	if v55&v75 != 0 {
		v145 = v41 + int32(base.Ui32(v55)>>(uint(v75)%32))
		goto L15
	} else {
		goto L32
	}
L23:
	;
	v65 = v60
	goto L25
L24:
	;
	v65 = int32(2)
	goto L25
L25:
	;
	if v62&int32(254) == int32(2) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v70 = v58
	goto L28
L27:
	;
	v70 = v65
	goto L28
L28:
	;
	if v62 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v73 = v58
	goto L31
L30:
	;
	v73 = v70
	goto L31
L31:
	;
	v145 = v41 + v73
	goto L15
L32:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v145 = v41 + int32(base.Ui32(v80)>>(uint(int32(2))%32))
	goto L15
L33:
	;
	v145 = v140 + v41 + int32(1)
	goto L15
L34:
	;
	v140 = v132 - v41
	goto L33
L35:
	;
	v111 = v107
	goto L44
L36:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v91 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v140 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v96 = v41
	goto L40
L40:
	;
	v100 = v96 + int32(1)
	if v100&int32(3) == int32(0) {
		v107 = v100
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v132 = v100
	goto L34
L42:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v105 != 0 {
		v96 = v100
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v120 = int32(-2139062144)
	if (int32(16843008)-v117|v117)&v120 == v120 {
		v111 = v111 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v126 = v111
	goto L47
L46:
	;
	goto L45
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 != 0 {
		v126 = v126 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v132 = v126
	goto L34
L49:
	;
	goto L48
L50:
	;
	v160 = v43 + int32(1)
	if v160 != l3 {
		v41 = v158
		v43 = v160
		goto L13
	} else {
		goto L54
	}
L51:
	;
	v158 = (v145 + int32(1)) & int32(-2)
	goto L50
L52:
	;
	v158 = (v145 + int32(7)) & int32(-8)
	goto L50
L53:
	;
	v158 = (v145 + int32(3)) & int32(-4)
	goto L50
L54:
	;
	goto L14
L55:
	;
	v165 = base.I32_div_s(l1, int32(8))
	v175 = l0
	v176 = int32(1) << (uint(l1&int32(7)) % 32)
	v177 = l2 + v165
	v182 = int32(0)
	goto L56
L56:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v176&v184 == int32(0) {
		v296 = v175
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v308 = v296
	goto L8
L58:
	;
	v298 = int32(1)
	v300 = v176 << (uint(v298) % 32)
	v302 = base.B2i32(v300 == int32(256))
	if v300 == int32(256) {
		goto L98
	} else {
		goto L99
	}
L59:
	;
	if int32(0) < l4 {
		v283 = v175 + l4
		goto L60
	} else {
		goto L61
	}
L60:
	;
	switch l5 - int32(99) {
	case 0:
		v296 = v283
		goto L58
	case 1:
		goto L96
	default:
		goto L95
	case 6:
		goto L97
	}
L61:
	;
	if base.B2i32(l4 != int32(-1)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v193 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	if v175&int32(3) == int32(0) {
		v245 = v175
		goto L80
	} else {
		goto L81
	}
L65:
	;
	v196 = int32(6)
	v198 = int32(18)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	if v200 == v198 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v213 = int32(1)
	if v193&v213 != 0 {
		v283 = v175 + int32(base.Ui32(v193)>>(uint(v213)%32))
		goto L60
	} else {
		goto L77
	}
L68:
	;
	v203 = v198
	goto L70
L69:
	;
	v203 = int32(2)
	goto L70
L70:
	;
	if v200&int32(254) == int32(2) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v208 = v196
	goto L73
L72:
	;
	v208 = v203
	goto L73
L73:
	;
	if v200 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v211 = v196
	goto L76
L75:
	;
	v211 = v208
	goto L76
L76:
	;
	v283 = v175 + v211
	goto L60
L77:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v283 = v175 + int32(base.Ui32(v218)>>(uint(int32(2))%32))
	goto L60
L78:
	;
	v283 = v278 + v175 + int32(1)
	goto L60
L79:
	;
	v278 = v270 - v175
	goto L78
L80:
	;
	v249 = v245
	goto L89
L81:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v229 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v278 = int32(0)
	goto L78
L83:
	;
	goto L84
L84:
	;
	v234 = v175
	goto L85
L85:
	;
	v238 = v234 + int32(1)
	if v238&int32(3) == int32(0) {
		v245 = v238
		goto L80
	} else {
		goto L87
	}
L86:
	;
	v270 = v238
	goto L79
L87:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v243 != 0 {
		v234 = v238
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v258 = int32(-2139062144)
	if (int32(16843008)-v255|v255)&v258 == v258 {
		v249 = v249 + int32(4)
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v264 = v249
	goto L92
L91:
	;
	goto L90
L92:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v268 != 0 {
		v264 = v264 + int32(1)
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v270 = v264
	goto L79
L94:
	;
	goto L93
L95:
	;
	v296 = (v283 + int32(1)) & int32(-2)
	goto L58
L96:
	;
	v296 = (v283 + int32(7)) & int32(-8)
	goto L58
L97:
	;
	v296 = (v283 + int32(3)) & int32(-4)
	goto L58
L98:
	;
	v303 = v298
	goto L100
L99:
	;
	v303 = v300
	goto L100
L100:
	;
	v306 = v182 + int32(1)
	if v306 != l3 {
		v175 = v296
		v176 = v303
		v177 = v177 + v302
		v182 = v306
		goto L56
	} else {
		goto L101
	}
L101:
	;
	goto L57
}
func F_array_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v77 int32
	_ = v77
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
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v357 int32
	_ = v357
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_DatumGetAnyArrayP(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v26 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = int32(40)
	goto L5
L4:
	;
	v29 = int32(12)
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20+v29)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L77
	}
L7:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v80 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	F_get_type_io_data(m, v31, int32(3), v50+int32(4), v50+int32(6), v50+int32(7), v50+int32(8), v50+int32(12), v50+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v38 = F_MemoryContextAlloc(m, v36, int32(48))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v47 == v31 {
		v77 = v33
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v31 ^ int32(-1)
	v50 = v43
	goto L8
L13:
	;
	v50 = v33
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
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v31
	v77 = v50
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
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20+v83)))
	if v80 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v77)+7)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+6)))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77)+4)))
	v101 = F_ArrayGetNItems(m, v85, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v96 = v88
	v97 = v89
	goto L20
L22:
	;
	goto L23
L23:
	;
	v91 = v20 + int32(16)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v96 = v91
	v97 = v91 + v92<<(uint(int32(2))%32)
	goto L20
L24:
	;
	F_pq_begintypsend(m, v17+int32(32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_enlargeStringInfo(m, v17+int32(32), int32(4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v115 = int32(24)
	v117 = int32(65280)
	v119 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v112+v113))) = v85<<(uint(v115)%32) | v85&v117<<(uint(v119)%32) | (int32(base.Ui32(v85)>>(uint(v119)%32))&v117 | int32(base.Ui32(v85)>>(uint(v115)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v112 + int32(4)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v134 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_enlargeStringInfo(m, v17+int32(32), int32(4))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L34
	}
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v137 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v148 = base.B2i32(v145 != int32(0))
	goto L27
L31:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v148 = base.B2i32(v138 != int32(0))
	goto L27
L32:
	;
	goto L33
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v148 = base.B2i32(v142 != int32(0))
	goto L27
L34:
	;
	v154 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v148 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v160 = int32(16777216)
	goto L37
L36:
	;
	v160 = v154
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155+v156))) = v160
	v162 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v155 + v162
	F_enlargeStringInfo(m, v17+int32(32), v162)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v173 = int32(24)
	v175 = int32(65280)
	v177 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v170+v171))) = v31<<(uint(v173)%32) | v31&v175<<(uint(v177)%32) | (int32(base.Ui32(v31)>>(uint(v177)%32))&v175 | int32(base.Ui32(v31)>>(uint(v173)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v170 + int32(4)
	if int32(0) < v85 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v194 = v154
	goto L42
L40:
	;
	goto L41
L41:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v285 == int32(-1) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v209 = v194 << (uint(int32(2)) % 32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v96+v209)))
	F_enlargeStringInfo(m, v17+int32(32), int32(4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v220 = int32(24)
	v222 = int32(65280)
	v224 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v217+v218))) = v211<<(uint(v220)%32) | v211&v222<<(uint(v224)%32) | (int32(base.Ui32(v211)>>(uint(v224)%32))&v222 | int32(base.Ui32(v211)>>(uint(v220)%32)))
	v236 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v217 + v236
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v209+v97)))
	F_enlargeStringInfo(m, v17+int32(32), v236)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v249 = int32(24)
	v251 = int32(65280)
	v253 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v246+v247))) = v240<<(uint(v249)%32) | v240&v251<<(uint(v253)%32) | (int32(base.Ui32(v240)>>(uint(v253)%32))&v251 | int32(base.Ui32(v240)>>(uint(v249)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v246 + int32(4)
	v269 = v194 + int32(1)
	if v269 != v85 {
		v194 = v269
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(1)
	if int32(0) < v101 {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v288 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+12)) = int64(0)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v321 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v291 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v290
	v344 = v291
	goto L47
L52:
	;
	goto L53
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+12)) = int64(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	if v298 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v297 + (v301<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v344 = int32(0)
	goto L47
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v298 + v297
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v344 = v297 + v313<<(uint(int32(3))%32) + int32(16)
	goto L47
L57:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v20 + (v324<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v344 = int32(0)
	goto L47
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v321 + v20
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v344 = v20 + v336<<(uint(int32(3))%32) + int32(16)
	goto L47
L60:
	;
	v357 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	v459 = v17 + int32(32)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v461))) = v462 << (uint(int32(2)) % 32)
	goto L76
L63:
	;
	v373 = F_array_iter_next(m, v17+int32(12), v17+int32(11), v357, v100, v99&int32(1), v98)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	goto L62
L65:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+11)))
	if v375 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v442 = v357 + int32(1)
	if v442 != v101 {
		v357 = v442
		goto L63
	} else {
		goto L75
	}
L67:
	;
	F_enlargeStringInfo(m, v17+int32(32), int32(4))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v391 = F_SendFunctionCall(m, v77+int32(20), v373)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v383+v384))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v383 + int32(4)
	goto L66
L71:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F_enlargeStringInfo(m, v17+int32(32), int32(4))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v402 = int32(2)
	v404 = int32(4)
	v405 = int32(base.Ui32(v393)>>(uint(v402)%32)) - v404
	v406 = int32(24)
	v408 = int32(65280)
	v410 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v399+v400))) = v405<<(uint(v406)%32) | v405&v408<<(uint(v410)%32) | (int32(base.Ui32(v405)>>(uint(v410)%32))&v408 | int32(base.Ui32(v405)>>(uint(v406)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v399 + v404
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F_pq_sendbytes(m, v17+int32(32), v391+v404, int32(base.Ui32(v429)>>(uint(v402)%32))-v404)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v391)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
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
	m.G0 = v17 + int32(48)
	return v461
L77:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v477 = F_format_type_be(m, v31)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v477
	F_errmsg(m, int32(182281), v17)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(477782), int32(1589), int32(411175))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	v26 = m.G0
	v28 = v26 - int32(80)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 <= int32(0) {
		v41 = F_construct_empty_array(m, l3)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v258 = v41
			m.G0 = v28 + int32(80)
			return v258
		}
	} else {
		if l1 <= int32(0) {
			v41 = F_construct_empty_array(m, l3)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v258 = v41
				m.G0 = v28 + int32(80)
				return v258
			}
		} else {
			v36 = l0 + int32(16)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if int32(0) < v37 {
				v46 = v30 << (uint(int32(2)) % 32)
				v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l4)+8)))
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+10)))
				v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4)+11)))
				F_deconstruct_array(m, l0, v48, v49&int32(1), v52, v28+int32(12), v28+int32(8), v28+int32(76))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v28)+76))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v63 = base.I32_div_s(v61, v62)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v63
					v67 = base.I64_extend_i32_s(v62 - int32(1))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
					v77 = v69
					v78 = v70
					v92 = int64(0)
					for {
						v96 = int32(4538720)
						if base.Ui64(v67) <= base.Ui64(v92) {
							v143 = v92
						} else {
							v103 = v67 - v92
							v105 = *(*int64)(unsafe.Add(mBase, _consts[88]))
							v106 = *(*int64)(unsafe.Add(mBase, _consts[87]))
							v109 = v106
							v111 = v105
							for {
								v115 = v109 ^ v111
								v117 = base.I64_rotl(v115, int64(37))
								v125 = v115 ^ (v115<<(uint(int64(16))%64) ^ base.I64_rotl(v109, int64(24)))
								v130 = int64(base.Ui64(base.I64_rotl(v109*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v103)) % 64))
								if base.Ui64(v103) < base.Ui64(v130) {
									v109 = v125
									v111 = v117
									continue
								} else {
									break
								}
								break
							}
							*(*int64)(unsafe.Add(mBase, _consts[88])) = v117
							*(*int64)(unsafe.Add(mBase, _consts[87])) = v125
							v143 = v92 + v130
						}
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v28)+76))
						if int32(0) < v144 {
							v148 = v144 * base.I32_wrap_i64(v143)
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
							v156 = v148 + v149
							v160 = v151 + v148<<(uint(int32(2))%32)
							v162 = v77
							v163 = v78
							v167 = int32(0)
							for {
								v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
								v182 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
								*(*int32)(unsafe.Add(mBase, uint32(v163))) = v183
								v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
								*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v185)
								*(*int32)(unsafe.Add(mBase, uint32(v160))) = v182
								*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v181)
								v189 = int32(1)
								v191 = int32(4)
								v194 = v162 + v189
								v196 = v163 + v191
								v198 = v167 + v189
								v199 = *(*int32)(unsafe.Add(mBase, uint32(v28)+76))
								if v198 < v199 {
									v156 = v156 + v189
									v160 = v160 + v191
									v162 = v194
									v163 = v196
									v167 = v198
									continue
								} else {
									break
								}
								break
							}
							v207 = v194
							v208 = v196
						} else {
							v207 = v77
							v208 = v78
						}
						v227 = v92 + int64(1)
						if v227 != base.I64_extend_i32_u(l1) {
							v77 = v207
							v78 = v208
							v92 = v227
							continue
						} else {
							break
						}
						break
					}
					if v46 != 0 {
						v231 = F__emscripten_memcpy_bulkmem(m, v28+int32(48), v36, v46)
						mBase = m.M
					} else {
					}
					if v46 != 0 {
						v235 = F__emscripten_memcpy_bulkmem(m, v28+int32(16), v36+v46, v46)
						mBase = m.M
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = l1
					if l2 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(1)
					} else {
					}
					v242 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v250 = F_construct_md_array(m, v242, v243, v30, v28+int32(48), v28+int32(16), l3, v48, v49&int32(1), v52)
					mBase = m.M
					v251 = m.ExcPending
					if v251 != 0 {
						return int32(0)
					} else {
						v252 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
						F_pfree(m, v252)
						mBase = m.M
						v254 = m.ExcPending
						if v254 != 0 {
							return int32(0)
						} else {
							v255 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
							F_pfree(m, v255)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								v258 = v250
								m.G0 = v28 + int32(80)
								return v258
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
					v258 = v41
					m.G0 = v28 + int32(80)
					return v258
				}
			}
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
	var v25 int32
	_ = v25
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(464) {
		v25 = v2
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 == int32(0) {
			v25 = v2
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v14 != int32(8) {
				v25 = v2
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				if v17 != 0 {
					v25 = v2
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v18 != v19 {
						v25 = v2
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						if v22 != 0 {
							v23 = v11
						} else {
							v23 = int32(0)
						}
						v25 = v23
					}
				}
			}
		}
	}
	return v25
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
