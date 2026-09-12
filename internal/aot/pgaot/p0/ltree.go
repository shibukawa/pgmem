package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_risparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_array_iterator(m, v6, int32(5648), v12, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v21 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								return v15
							}
						} else {
							return v15
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v21 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return v15
						}
					} else {
						return v15
					}
				}
			}
		}
	}
}
func F_ltree_addltree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_ltree_concat(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return v13
							}
						} else {
							return v13
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return v13
						}
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_ltree_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v11 == int32(0) {
		v138 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return (v138 + int32(1)) * (v11 - v10) * int32(10)
L2:
	;
	if v10 == int32(0) {
		v138 = v11
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(8)
	v20 = l0 + v16
	v21 = l1 + v16
	v24 = v11
	v28 = v10
	goto L4
L4:
	;
	v29 = int32(2)
	v30 = v20 + v29
	v32 = v21 + v29
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	if base.Ui32(v33) < base.Ui32(v34) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v138 = v119
	goto L1
L6:
	;
	v119 = v24 - int32(1)
	if v24 < int32(2) {
		v138 = v119
		goto L1
	} else {
		goto L35
	}
L7:
	;
	v36 = v33
	goto L9
L8:
	;
	v36 = v34
	goto L9
L9:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v36) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v98 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L11:
	;
	v98 = int32(0)
	goto L10
L12:
	;
	v72 = v67
	v73 = v68
	v74 = v69
	goto L22
L13:
	;
	if (v30|v32)&int32(3) != 0 {
		v67 = v30
		v68 = v32
		v69 = v36
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v60 = v30
	v61 = v32
	v62 = v36
	goto L15
L15:
	;
	if v62 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L16:
	;
	v44 = v30
	v45 = v32
	v46 = v36
	goto L17
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v49 != v50 {
		v67 = v44
		v68 = v45
		v69 = v46
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v60 = v55
	v61 = v53
	v62 = v57
	goto L15
L19:
	;
	v52 = int32(4)
	v53 = v45 + v52
	v55 = v44 + v52
	v57 = v46 - v52
	if base.Ui32(int32(3)) < base.Ui32(v57) {
		v44 = v55
		v45 = v53
		v46 = v57
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v67 = v60
	v68 = v61
	v69 = v62
	goto L12
L22:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 == v78 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v98 = v77 - v78
	goto L10
L24:
	;
	v80 = int32(1)
	v85 = v74 - v80
	if v85 != 0 {
		v72 = v72 + v80
		v73 = v73 + v80
		v74 = v85
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L11
L28:
	;
	if v33 == v34 {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v98 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v102 = int32(10)
	return (v24*v102 + v102) * (v33 - v34)
L32:
	;
	v115 = int32(-10)
	goto L34
L33:
	;
	v115 = int32(10)
	goto L34
L34:
	;
	return (v24 + int32(1)) * v115
L35:
	;
	v122 = int32(9)
	v124 = int32(131064)
	v132 = int32(1)
	if v132 < v28 {
		v20 = v20 + (v33+v122)&v124
		v21 = v21 + (v34+v122)&v124
		v24 = v119
		v28 = v28 - v132
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L5
}
func F_ltree_gist_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(76078)
			F_errmsg(m, int32(193003), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494159), int32(26), int32(279883))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
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
func F_ltree_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	if v22 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v179 != v14 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v178 = (v149 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	if v21 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v27 = int32(8)
	v33 = v14 + v27
	v35 = v22
	v38 = v19 + v27
	v42 = v21
	goto L8
L8:
	;
	v43 = int32(2)
	v44 = v33 + v43
	v46 = v38 + v43
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33))))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v47) < base.Ui32(v48) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v149 = v131
	goto L5
L10:
	;
	v131 = v35 - int32(1)
	if v35 < int32(2) {
		v149 = v131
		goto L5
	} else {
		goto L39
	}
L11:
	;
	v50 = v47
	goto L13
L12:
	;
	v50 = v48
	goto L13
L13:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v112 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v112 = int32(0)
	goto L14
L16:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L26
L17:
	;
	if (v44|v46)&int32(3) != 0 {
		v81 = v44
		v82 = v46
		v83 = v50
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v74 = v44
	v75 = v46
	v76 = v50
	goto L19
L19:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v58 = v44
	v59 = v46
	v60 = v50
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v63 != v64 {
		v81 = v58
		v82 = v59
		v83 = v60
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v74 = v69
	v75 = v67
	v76 = v71
	goto L19
L23:
	;
	v66 = int32(4)
	v67 = v59 + v66
	v69 = v58 + v66
	v71 = v60 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v58 = v69
		v59 = v67
		v60 = v71
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L16
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 == v92 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v112 = v91 - v92
	goto L14
L28:
	;
	v94 = int32(1)
	v99 = v88 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v87 + v94
		v88 = v99
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	if v47 == v48 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v112 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v116 = int32(10)
	v178 = (v35*v116 + v116) * (v47 - v48)
	goto L4
L36:
	;
	v128 = int32(-10)
	goto L38
L37:
	;
	v128 = int32(10)
	goto L38
L38:
	;
	v178 = (v35 + int32(1)) * v128
	goto L4
L39:
	;
	v134 = int32(9)
	v136 = int32(131064)
	v144 = int32(1)
	if v144 < v42 {
		v33 = v33 + (v47+v134)&v136
		v35 = v131
		v38 = v38 + (v48+v134)&v136
		v42 = v42 - v144
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L9
L41:
	;
	F_pfree(m, v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 != v19 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v19)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	return base.B2i32(v178 <= int32(0))
L48:
	;
	goto L47
}
