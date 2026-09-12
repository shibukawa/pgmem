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
	F_errmsg(m, int32(487505), v11)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, int32(0), int32(495035), int32(141), int32(410703))
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
	F_errmsg_internal(m, int32(444233), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(495703), int32(984), int32(154010))
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
					F_errmsg_internal(m, int32(61433), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495810), int32(953), int32(281052))
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
					F_errmsg(m, int32(371613), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495810), int32(942), int32(281052))
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
	F_errmsg(m, int32(190021), v22)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(495703), int32(4035), int32(236315))
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
	F_errmsg(m, int32(162748), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(495703), int32(4017), int32(236315))
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
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
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
		v134 = v22
		m.G0 = v11 + int32(16)
		return v134
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
					v99 = l3 + v35
					v100 = v38
					switch l5 - int32(99) {
					case 0:
						v115 = v99
					case 1:
						v115 = (v99 + int32(7)) & int32(-8)
					default:
						v115 = (v99 + int32(1)) & int32(-2)
					case 6:
						v115 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
					v118 = v100
					v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v122 = v120 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
					if v122 != int32(256) {
						v134 = v118
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v126 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v134 = v118
					}
					m.G0 = v11 + int32(16)
					return v134
				case 1:
					v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35))))
					v99 = l3 + v35
					v100 = v40
					switch l5 - int32(99) {
					case 0:
						v115 = v99
					case 1:
						v115 = (v99 + int32(7)) & int32(-8)
					default:
						v115 = (v99 + int32(1)) & int32(-2)
					case 6:
						v115 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
					v118 = v100
					v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v122 = v120 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
					if v122 != int32(256) {
						v134 = v118
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v126 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v134 = v118
					}
					m.G0 = v11 + int32(16)
					return v134
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
						F_errmsg_internal(m, int32(484426), v11)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(327351), int32(70), int32(68101))
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
					v99 = l3 + v35
					v100 = v42
					switch l5 - int32(99) {
					case 0:
						v115 = v99
					case 1:
						v115 = (v99 + int32(7)) & int32(-8)
					default:
						v115 = (v99 + int32(1)) & int32(-2)
					case 6:
						v115 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
					v118 = v100
					v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v122 = v120 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
					if v122 != int32(256) {
						v134 = v118
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v126 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v134 = v118
					}
					m.G0 = v11 + int32(16)
					return v134
				}
			} else {
				if int32(0) < l3 {
					v99 = l3 + v35
					v100 = v35
				} else {
					if l3 == int32(-1) {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
						if v64 == int32(1) {
							v67 = int32(6)
							v69 = int32(18)
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
							if v71 == v69 {
								v74 = v69
							} else {
								v74 = int32(2)
							}
							if v71&int32(254) == int32(2) {
								v79 = v67
							} else {
								v79 = v74
							}
							if v71 == int32(1) {
								v82 = v67
							} else {
								v82 = v79
							}
							v98 = v35 + v82
						} else {
							v84 = int32(1)
							if v64&v84 != 0 {
								v98 = v35 + int32(base.Ui32(v64)>>(uint(v84)%32))
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
								v98 = v35 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
							}
						}
					} else {
						v93 = F_strlen(m, v35)
						mBase = m.M
						v98 = v93 + v35 + int32(1)
					}
					v99 = v98
					v100 = v35
				}
				switch l5 - int32(99) {
				case 0:
					v115 = v99
				case 1:
					v115 = (v99 + int32(7)) & int32(-8)
				default:
					v115 = (v99 + int32(1)) & int32(-2)
				case 6:
					v115 = (v99 + int32(3)) & int32(-4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
				v118 = v100
				v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v122 = v120 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
				if v122 != int32(256) {
					v134 = v118
				} else {
					v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v126 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
					v134 = v118
				}
				m.G0 = v11 + int32(16)
				return v134
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
						v99 = l3 + v35
						v100 = v38
						switch l5 - int32(99) {
						case 0:
							v115 = v99
						case 1:
							v115 = (v99 + int32(7)) & int32(-8)
						default:
							v115 = (v99 + int32(1)) & int32(-2)
						case 6:
							v115 = (v99 + int32(3)) & int32(-4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
						v118 = v100
						v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v122 = v120 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
						if v122 != int32(256) {
							v134 = v118
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v126 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v134 = v118
						}
						m.G0 = v11 + int32(16)
						return v134
					case 1:
						v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35))))
						v99 = l3 + v35
						v100 = v40
						switch l5 - int32(99) {
						case 0:
							v115 = v99
						case 1:
							v115 = (v99 + int32(7)) & int32(-8)
						default:
							v115 = (v99 + int32(1)) & int32(-2)
						case 6:
							v115 = (v99 + int32(3)) & int32(-4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
						v118 = v100
						v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v122 = v120 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
						if v122 != int32(256) {
							v134 = v118
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v126 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v134 = v118
						}
						m.G0 = v11 + int32(16)
						return v134
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
							F_errmsg_internal(m, int32(484426), v11)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(327351), int32(70), int32(68101))
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
						v99 = l3 + v35
						v100 = v42
						switch l5 - int32(99) {
						case 0:
							v115 = v99
						case 1:
							v115 = (v99 + int32(7)) & int32(-8)
						default:
							v115 = (v99 + int32(1)) & int32(-2)
						case 6:
							v115 = (v99 + int32(3)) & int32(-4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
						v118 = v100
						v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v122 = v120 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
						if v122 != int32(256) {
							v134 = v118
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v126 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
							v134 = v118
						}
						m.G0 = v11 + int32(16)
						return v134
					}
				} else {
					if int32(0) < l3 {
						v99 = l3 + v35
						v100 = v35
					} else {
						if l3 == int32(-1) {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
							if v64 == int32(1) {
								v67 = int32(6)
								v69 = int32(18)
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
								if v71 == v69 {
									v74 = v69
								} else {
									v74 = int32(2)
								}
								if v71&int32(254) == int32(2) {
									v79 = v67
								} else {
									v79 = v74
								}
								if v71 == int32(1) {
									v82 = v67
								} else {
									v82 = v79
								}
								v98 = v35 + v82
							} else {
								v84 = int32(1)
								if v64&v84 != 0 {
									v98 = v35 + int32(base.Ui32(v64)>>(uint(v84)%32))
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
									v98 = v35 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
								}
							}
						} else {
							v93 = F_strlen(m, v35)
							mBase = m.M
							v98 = v93 + v35 + int32(1)
						}
						v99 = v98
						v100 = v35
					}
					switch l5 - int32(99) {
					case 0:
						v115 = v99
					case 1:
						v115 = (v99 + int32(7)) & int32(-8)
					default:
						v115 = (v99 + int32(1)) & int32(-2)
					case 6:
						v115 = (v99 + int32(3)) & int32(-4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115
					v118 = v100
					v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v122 = v120 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
					if v122 != int32(256) {
						v134 = v118
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v126 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
						v134 = v118
					}
					m.G0 = v11 + int32(16)
					return v134
				}
			} else {
				v30 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
				v118 = int32(0)
				v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v122 = v120 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v122
				if v122 != int32(256) {
					v134 = v118
				} else {
					v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v126 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126 + int32(1)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
					v134 = v118
				}
				m.G0 = v11 + int32(16)
				return v134
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
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
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
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v137 {
		goto L45
	} else {
		goto L46
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
	F_errmsg_internal(m, int32(484426), v15)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(327351), int32(70), int32(68101))
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
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	switch v115 - int32(99) {
	case 0:
		v132 = v114
		goto L41
	case 1:
		goto L43
	default:
		goto L42
	case 6:
		goto L44
	}
L22:
	;
	v114 = v42 + v74
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
	v109 = F_strlen(m, v42)
	mBase = m.M
	v114 = v109 + v42 + int32(1)
	goto L21
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
		v114 = v42 + int32(base.Ui32(v80)>>(uint(v100)%32))
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
	v114 = v42 + v98
	goto L21
L40:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v114 = v42 + int32(base.Ui32(v105)>>(uint(int32(2))%32))
	goto L21
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v132
	goto L1
L42:
	;
	v132 = (v114 + int32(1)) & int32(-2)
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = (v114 + int32(7)) & int32(-8)
	goto L1
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = (v114 + int32(3)) & int32(-4)
	goto L1
L45:
	;
	v144 = int32(0)
	v145 = v134
	goto L48
L46:
	;
	v277 = v20
	v278 = v134
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v278
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v291 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v293 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+15)))
	v294 = F_construct_md_array(m, v136, v135, v277, v287, v288, v290, v291, v292, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L17
	} else {
		goto L89
	}
L48:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v153 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v157 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v277 = v273
	v278 = v267
	goto L47
L50:
	;
	v270 = v144 + int32(1)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v270 < v271 {
		v144 = v270
		v145 = v267
		goto L48
	} else {
		goto L88
	}
L51:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144+v135))) = uint8(v178)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v183 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v161 = base.I32_div_s(v153, int32(8))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v161))))
	if int32(base.Ui32(v163)>>(uint(v153&int32(7))%32))&int32(1) != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v144+v135))) = uint8(v170)
	*(*int32)(unsafe.Add(mBase, uint32(v136+v144<<(uint(int32(2))%32)))) = int32(0)
	v267 = v145
	goto L50
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136+v144<<(uint(int32(2))%32)))) = v209
	v211 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v211 {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	switch v186 - int32(1) {
	case 0:
		goto L61
	case 1:
		goto L60
	default:
		goto L58
	case 3:
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v209 = v145
	goto L54
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L17
	} else {
		goto L62
	}
L59:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v209 = v191
	goto L54
L60:
	;
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145))))
	v209 = v190
	goto L54
L61:
	;
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v145))))
	v209 = v189
	goto L54
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = base.I32_extend16_s(v186)
	F_errmsg_internal(m, int32(484426), v15+int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(327351), int32(70), int32(68101))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L17
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
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	switch v252 - int32(99) {
	case 0:
		v267 = v251
		goto L50
	case 1:
		goto L86
	default:
		goto L85
	case 6:
		goto L87
	}
L66:
	;
	v251 = v145 + v211
	goto L65
L67:
	;
	goto L68
L68:
	;
	if v211 == int32(-1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v217 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v246 = F_strlen(m, v145)
	mBase = m.M
	v251 = v246 + v145 + int32(1)
	goto L65
L72:
	;
	v220 = int32(6)
	v222 = int32(18)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	if v224 == v222 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v237 = int32(1)
	if v217&v237 != 0 {
		v251 = v145 + int32(base.Ui32(v217)>>(uint(v237)%32))
		goto L65
	} else {
		goto L84
	}
L75:
	;
	v227 = v222
	goto L77
L76:
	;
	v227 = int32(2)
	goto L77
L77:
	;
	if v224&int32(254) == int32(2) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v232 = v220
	goto L80
L79:
	;
	v232 = v227
	goto L80
L80:
	;
	if v224 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v235 = v220
	goto L83
L82:
	;
	v235 = v232
	goto L83
L83:
	;
	v251 = v145 + v235
	goto L65
L84:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v251 = v145 + int32(base.Ui32(v242)>>(uint(int32(2))%32))
	goto L65
L85:
	;
	v267 = (v251 + int32(1)) & int32(-2)
	goto L50
L86:
	;
	v267 = (v251 + int32(7)) & int32(-8)
	goto L50
L87:
	;
	v267 = (v251 + int32(3)) & int32(-4)
	goto L50
L88:
	;
	goto L49
L89:
	;
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v296)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v294
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
	F_errmsg(m, int32(189810), v14)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(495810), int32(1554), int32(139335))
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
	F_errmsg(m, int32(442938), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(495810), int32(1502), int32(139335))
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
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v461 int32
	_ = v461
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
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
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L160
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L155
	}
L3:
	;
	m.G0 = v32 + int32(80)
	return v461
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
	v461 = l0
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
	if v367 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L38:
	;
	v370 = int32(1)
	v372 = v120 << (uint(v370) % 32)
	v374 = base.B2i32(v372 == int32(256))
	if v372 == int32(256) {
		goto L115
	} else {
		goto L116
	}
L39:
	;
	v362 = v353
	v364 = v355
	v366 = v126 + int32(1)
	v367 = v357
	v368 = v358
	goto L38
L40:
	;
	if int32(0) < v64 {
		v316 = v64
		goto L91
	} else {
		goto L92
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v126<<(uint(int32(2))%32)))) = v265
	v273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v91))) = uint8(v273)
	v276 = v264
	v278 = v266
	goto L40
L42:
	;
	v264 = v229
	v265 = v216
	v266 = v131
	goto L41
L43:
	;
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v91))) = uint8(v258)
	v353 = v119
	v355 = v123
	v357 = v131
	v358 = v258
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
	v362 = v119
	v364 = v123
	v366 = v126
	v367 = int32(1)
	v368 = v135
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
	v264 = v123
	v265 = v75
	v266 = int32(1)
	goto L41
L52:
	;
	switch v109 {
	case 0:
		v229 = v215
		goto L79
	case 1:
		goto L81
	default:
		goto L80
	case 6:
		goto L82
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
	v215 = v123 + v64
	v216 = v156
	goto L52
L58:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123))))
	v215 = v123 + v64
	v216 = v154
	goto L52
L59:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
	v215 = v123 + v64
	v216 = v152
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v64
	F_errmsg_internal(m, int32(484426), v32+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(327351), int32(70), int32(68101))
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
	v215 = v123 + v64
	v216 = v123
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
	v215 = v214
	v216 = v123
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
	v208 = F_strlen(m, v123)
	mBase = m.M
	v214 = v208 + v123 + int32(1)
	goto L66
L70:
	;
	v214 = v206 + v123
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
		v214 = v123 + int32(base.Ui32(v178)>>(uint(v197)%32))
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
	v214 = v123 + v195
	goto L66
L78:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v206 = int32(base.Ui32(v202) >> (uint(int32(2)) % 32))
	goto L70
L79:
	;
	if l2 != 0 {
		goto L42
	} else {
		goto L83
	}
L80:
	;
	v229 = (v215 + int32(1)) & int32(-2)
	goto L79
L81:
	;
	v229 = (v215 + int32(7)) & int32(-8)
	goto L79
L82:
	;
	v229 = (v215 + int32(3)) & int32(-4)
	goto L79
L83:
	;
	v230 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v74
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v216
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+60)) = uint8(v230)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v242 = m.T0[v241].(func(*base.Module, int32) int32)(m, v32+int32(44))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+60)))
	if v244 != 0 {
		goto L42
	} else {
		goto L85
	}
L85:
	;
	if v242 == int32(0) {
		v264 = v229
		v265 = v216
		v266 = v131
		goto L41
	} else {
		goto L86
	}
L86:
	;
	if l5 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v362 = v119
	v364 = v229
	v366 = v126
	v367 = int32(1)
	v368 = v135
	goto L38
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v126<<(uint(int32(2))%32)))) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v91))) = uint8(v5)
	v254 = int32(1)
	if v5 == int32(0) {
		v276 = v229
		v278 = v254
		goto L40
	} else {
		goto L90
	}
L90:
	;
	v353 = v119
	v355 = v229
	v357 = v254
	v358 = int32(1)
	goto L39
L91:
	;
	v318 = v119 + v316
	switch v109 {
	case 0:
		v331 = v318
		goto L106
	case 1:
		goto L108
	default:
		goto L107
	case 6:
		goto L109
	}
L92:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v89+v126<<(uint(int32(2))%32))))
	if v64 == int32(-1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v288 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v313 = F_strlen(m, v285)
	mBase = m.M
	v316 = v313 + int32(1)
	goto L91
L96:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	if base.Ui32((v292-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v316 = int32(6)
		goto L91
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v288&int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v299 = int32(18)
	if v292&int32(255) == v299 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v305 = v299
	goto L102
L101:
	;
	v305 = int32(2)
	goto L102
L102:
	;
	v316 = v305
	goto L91
L103:
	;
	v316 = int32(base.Ui32(v288) >> (uint(int32(1)) % 32))
	goto L91
L104:
	;
	goto L105
L105:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v316 = int32(base.Ui32(v310) >> (uint(int32(2)) % 32))
	goto L91
L106:
	;
	if base.Ui32(v331) < base.Ui32(int32(1073741824)) {
		v353 = v331
		v355 = v276
		v357 = v278
		v358 = v135
		goto L39
	} else {
		goto L110
	}
L107:
	;
	v331 = (v318 + int32(1)) & int32(-2)
	goto L106
L108:
	;
	v331 = (v318 + int32(7)) & int32(-8)
	goto L106
L109:
	;
	v331 = (v318 + int32(3)) & int32(-4)
	goto L106
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(1073741823)
	F_errmsg(m, int32(680940), v32+int32(32))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(495703), int32(6567), int32(313068))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	v375 = v370
	goto L117
L116:
	;
	v375 = v372
	goto L117
L117:
	;
	if v121 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v376 = v375
	goto L120
L119:
	;
	v376 = v120
	goto L120
L120:
	;
	if v121 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v379 = v121 + v374
	goto L123
L122:
	;
	v379 = int32(0)
	goto L123
L123:
	;
	v381 = v138 + int32(1)
	if v381 != v38 {
		v119 = v362
		v120 = v376
		v121 = v379
		v123 = v364
		v126 = v366
		v131 = v367
		v135 = v368
		v138 = v381
		goto L36
	} else {
		goto L124
	}
L124:
	;
	goto L37
L125:
	;
	F_pfree(m, v89)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if v366 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	F_pfree(m, v91)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v461 = l0
	goto L3
L130:
	;
	F_pfree(m, v89)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v404 = v35 << (uint(int32(3)) % 32)
	if v368 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	F_pfree(m, v91)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	v396 = F_palloc0(m, int32(16))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v396)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v396))) = int64(64)
	v461 = v396
	goto L3
L136:
	;
	v421 = v362 + v419
	v422 = F_palloc0(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L140
	}
L137:
	;
	v408 = base.I32_div_s(v366+int32(7), int32(8))
	v413 = (v404 + v408 + int32(23)) & int32(-8)
	v419 = v413
	v420 = v413
	goto L136
L138:
	;
	goto L139
L139:
	;
	v419 = (v404 + int32(23)) & int32(-8)
	v420 = int32(0)
	goto L136
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v422)+8)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v422)+4)) = v35
	v427 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = v421 << (uint(v427) % 32)
	v431 = v422 + int32(16)
	v433 = v35 << (uint(v427) % 32)
	if v433 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v433 != 0 {
		goto L146
	} else {
		goto L147
	}
L142:
	;
	v434 = F__emscripten_memcpy_bulkmem(m, v431, v37, v433)
	mBase = m.M
	v435 = v434
	goto L144
L143:
	;
	v435 = v431
	goto L144
L144:
	;
	goto L141
L145:
	;
	if l5 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v441 = F__emscripten_memcpy_bulkmem(m, v435+v433, v37+v437<<(uint(int32(2))%32), v433)
	mBase = m.M
	goto L148
L147:
	;
	goto L148
L148:
	;
	goto L145
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v366
	goto L151
L150:
	;
	goto L151
L151:
	;
	F_CopyArrayEls(m, v422, v89, v91, v366, v64, v61&int32(1), base.I32_extend8_s(v62), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_pfree(m, v89)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	F_pfree(m, v91)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v461 = v422
	goto L3
L155:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v493 = F_format_type_be(m, v34)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v493
	F_errmsg(m, int32(189810), v32)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(495703), int32(6447), int32(313068))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(443005), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(495703), int32(6431), int32(313068))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
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
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	if l2 != 0 {
		if l2 == int32(0) {
			if l3 <= int32(0) {
				v196 = l0
			} else {
				v34 = int32(0)
				v41 = l0
				v43 = v34
				for {
					if base.B2i32(l4 <= v34) == int32(0) {
						v89 = v41 + l4
					} else {
						if base.B2i32(l4 != int32(-1)) == int32(0) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
							if v55 == int32(1) {
								v58 = int32(6)
								v60 = int32(18)
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
								if v62 == v60 {
									v65 = v60
								} else {
									v65 = int32(2)
								}
								if v62&int32(254) == int32(2) {
									v70 = v58
								} else {
									v70 = v65
								}
								if v62 == int32(1) {
									v73 = v58
								} else {
									v73 = v70
								}
								v89 = v41 + v73
							} else {
								v75 = int32(1)
								if v55&v75 != 0 {
									v89 = v41 + int32(base.Ui32(v55)>>(uint(v75)%32))
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
									v89 = v41 + int32(base.Ui32(v80)>>(uint(int32(2))%32))
								}
							}
						} else {
							v84 = F_strlen(m, v41)
							mBase = m.M
							v89 = v84 + v41 + int32(1)
						}
					}
					switch l5 - int32(99) {
					case 0:
						v102 = v89
					case 1:
						v102 = (v89 + int32(7)) & int32(-8)
					default:
						v102 = (v89 + int32(1)) & int32(-2)
					case 6:
						v102 = (v89 + int32(3)) & int32(-4)
					}
					v104 = v43 + int32(1)
					if v104 != l3 {
						v41 = v102
						v43 = v104
						continue
					} else {
						break
					}
					break
				}
				v196 = v102
			}
		} else {
			if l3 <= int32(0) {
				v196 = l0
			} else {
				v109 = base.I32_div_s(l1, int32(8))
				v119 = l0
				v120 = int32(1) << (uint(l1&int32(7)) % 32)
				v121 = l2 + v109
				v126 = int32(0)
				for {
					v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
					if v120&v128 == int32(0) {
						v184 = v119
					} else {
						if int32(0) < l4 {
							v171 = v119 + l4
						} else {
							if base.B2i32(l4 != int32(-1)) == int32(0) {
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
								if v137 == int32(1) {
									v140 = int32(6)
									v142 = int32(18)
									v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
									if v144 == v142 {
										v147 = v142
									} else {
										v147 = int32(2)
									}
									if v144&int32(254) == int32(2) {
										v152 = v140
									} else {
										v152 = v147
									}
									if v144 == int32(1) {
										v155 = v140
									} else {
										v155 = v152
									}
									v171 = v119 + v155
								} else {
									v157 = int32(1)
									if v137&v157 != 0 {
										v171 = v119 + int32(base.Ui32(v137)>>(uint(v157)%32))
									} else {
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
										v171 = v119 + int32(base.Ui32(v162)>>(uint(int32(2))%32))
									}
								}
							} else {
								v166 = F_strlen(m, v119)
								mBase = m.M
								v171 = v166 + v119 + int32(1)
							}
						}
						switch l5 - int32(99) {
						case 0:
							v184 = v171
						case 1:
							v184 = (v171 + int32(7)) & int32(-8)
						default:
							v184 = (v171 + int32(1)) & int32(-2)
						case 6:
							v184 = (v171 + int32(3)) & int32(-4)
						}
					}
					v186 = int32(1)
					v188 = v120 << (uint(v186) % 32)
					v190 = base.B2i32(v188 == int32(256))
					if v188 == int32(256) {
						v191 = v186
					} else {
						v191 = v188
					}
					v194 = v126 + int32(1)
					if v194 != l3 {
						v119 = v184
						v120 = v191
						v121 = v121 + v190
						v126 = v194
						continue
					} else {
						break
					}
					break
				}
				v196 = v184
			}
		}
		return v196
	} else {
		if l4 <= int32(0) {
			if l2 == int32(0) {
				if l3 <= int32(0) {
					v196 = l0
				} else {
					v34 = int32(0)
					v41 = l0
					v43 = v34
					for {
						if base.B2i32(l4 <= v34) == int32(0) {
							v89 = v41 + l4
						} else {
							if base.B2i32(l4 != int32(-1)) == int32(0) {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
								if v55 == int32(1) {
									v58 = int32(6)
									v60 = int32(18)
									v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
									if v62 == v60 {
										v65 = v60
									} else {
										v65 = int32(2)
									}
									if v62&int32(254) == int32(2) {
										v70 = v58
									} else {
										v70 = v65
									}
									if v62 == int32(1) {
										v73 = v58
									} else {
										v73 = v70
									}
									v89 = v41 + v73
								} else {
									v75 = int32(1)
									if v55&v75 != 0 {
										v89 = v41 + int32(base.Ui32(v55)>>(uint(v75)%32))
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
										v89 = v41 + int32(base.Ui32(v80)>>(uint(int32(2))%32))
									}
								}
							} else {
								v84 = F_strlen(m, v41)
								mBase = m.M
								v89 = v84 + v41 + int32(1)
							}
						}
						switch l5 - int32(99) {
						case 0:
							v102 = v89
						case 1:
							v102 = (v89 + int32(7)) & int32(-8)
						default:
							v102 = (v89 + int32(1)) & int32(-2)
						case 6:
							v102 = (v89 + int32(3)) & int32(-4)
						}
						v104 = v43 + int32(1)
						if v104 != l3 {
							v41 = v102
							v43 = v104
							continue
						} else {
							break
						}
						break
					}
					v196 = v102
				}
			} else {
				if l3 <= int32(0) {
					v196 = l0
				} else {
					v109 = base.I32_div_s(l1, int32(8))
					v119 = l0
					v120 = int32(1) << (uint(l1&int32(7)) % 32)
					v121 = l2 + v109
					v126 = int32(0)
					for {
						v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
						if v120&v128 == int32(0) {
							v184 = v119
						} else {
							if int32(0) < l4 {
								v171 = v119 + l4
							} else {
								if base.B2i32(l4 != int32(-1)) == int32(0) {
									v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
									if v137 == int32(1) {
										v140 = int32(6)
										v142 = int32(18)
										v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
										if v144 == v142 {
											v147 = v142
										} else {
											v147 = int32(2)
										}
										if v144&int32(254) == int32(2) {
											v152 = v140
										} else {
											v152 = v147
										}
										if v144 == int32(1) {
											v155 = v140
										} else {
											v155 = v152
										}
										v171 = v119 + v155
									} else {
										v157 = int32(1)
										if v137&v157 != 0 {
											v171 = v119 + int32(base.Ui32(v137)>>(uint(v157)%32))
										} else {
											v162 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
											v171 = v119 + int32(base.Ui32(v162)>>(uint(int32(2))%32))
										}
									}
								} else {
									v166 = F_strlen(m, v119)
									mBase = m.M
									v171 = v166 + v119 + int32(1)
								}
							}
							switch l5 - int32(99) {
							case 0:
								v184 = v171
							case 1:
								v184 = (v171 + int32(7)) & int32(-8)
							default:
								v184 = (v171 + int32(1)) & int32(-2)
							case 6:
								v184 = (v171 + int32(3)) & int32(-4)
							}
						}
						v186 = int32(1)
						v188 = v120 << (uint(v186) % 32)
						v190 = base.B2i32(v188 == int32(256))
						if v188 == int32(256) {
							v191 = v186
						} else {
							v191 = v188
						}
						v194 = v126 + int32(1)
						if v194 != l3 {
							v119 = v184
							v120 = v191
							v121 = v121 + v190
							v126 = v194
							continue
						} else {
							break
						}
						break
					}
					v196 = v184
				}
			}
			return v196
		} else {
			switch l5 - int32(99) {
			case 0:
				v26 = l4
			case 1:
				v26 = (l4 + int32(7)) & int32(-8)
			default:
				v26 = (l4 + int32(1)) & int32(-2)
			case 6:
				v26 = (l4 + int32(3)) & int32(-4)
			}
			return l0 + l3*v26
		}
	}
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
	F_errmsg(m, int32(190178), v17)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(495703), int32(1589), int32(427146))
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
						v96 = int32(4603776)
						if base.Ui64(v67) <= base.Ui64(v92) {
							v143 = v92
						} else {
							v103 = v67 - v92
							v105 = *(*int64)(unsafe.Add(mBase, _consts[92]))
							v106 = *(*int64)(unsafe.Add(mBase, _consts[91]))
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
							*(*int64)(unsafe.Add(mBase, _consts[92])) = v117
							*(*int64)(unsafe.Add(mBase, _consts[91])) = v125
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
