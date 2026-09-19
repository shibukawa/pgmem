package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func Fn13846(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v148
	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v148)+11)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v155
	return v154
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v67 = l4
	v69 = l3
	goto L19
L5:
	;
	if v49-v50 == int32(0) {
		v148 = v11
		goto L1
	} else {
		goto L18
	}
L7:
	;
	goto L8
L8:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = l0
	v20 = v11
	v21 = int32(10)
	v22 = v18
	goto L13
L10:
	;
	v45 = v11
	v49 = int32(0)
	goto L11
L11:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	goto L5
L12:
	;
	v45 = v40
	v49 = v42
	goto L11
L13:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v22 != v24)|base.B2i32(v24 == int32(0)) != 0 {
		v40 = v20
		v42 = v22
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v40 = v34
	v42 = int32(0)
	goto L12
L15:
	;
	v30 = v21 - int32(1)
	if v30 == int32(0) {
		v40 = v20
		v42 = v22
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v33 = int32(1)
	v34 = v20 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v35 != 0 {
		v19 = v19 + v33
		v20 = v34
		v21 = v30
		v22 = v35
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L4
L19:
	;
	v76 = v67 + (v69-v67)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76))))
	v78 = v60 - v77
	if v78 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(31)
L21:
	;
	goto L26
L22:
	;
	v129 = v78
	goto L23
L23:
	;
	v133 = base.B2i32(v129 < int32(0))
	if v129 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	if v120 == int32(0) {
		v148 = v76
		goto L1
	} else {
		goto L37
	}
L26:
	;
	goto L27
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v87 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = l0
	v89 = v76
	v90 = int32(10)
	v91 = v87
	goto L32
L29:
	;
	v114 = v76
	v118 = int32(0)
	goto L30
L30:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v120 = v118 - v119
	goto L24
L31:
	;
	v114 = v109
	v118 = v111
	goto L30
L32:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if base.B2i32(v91 != v93)|base.B2i32(v93 == int32(0)) != 0 {
		v109 = v89
		v111 = v91
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v109 = v103
	v111 = int32(0)
	goto L31
L34:
	;
	v99 = v90 - int32(1)
	if v99 == int32(0) {
		v109 = v89
		v111 = v91
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v102 = int32(1)
	v103 = v89 + v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v104 != 0 {
		v88 = v88 + v102
		v89 = v103
		v90 = v99
		v91 = v104
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v129 = v120
	goto L23
L38:
	;
	v134 = v76 - int32(16)
	goto L40
L39:
	;
	v134 = v69
	goto L40
L40:
	;
	if v129 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = v67
	goto L43
L42:
	;
	v137 = v76 + int32(16)
	goto L43
L43:
	;
	if base.Ui32(v137) <= base.Ui32(v134) {
		v67 = v137
		v69 = v134
		goto L19
	} else {
		goto L44
	}
L44:
	;
	goto L20
}
func Fn13853(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
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
	v1 = l0
	v6 = F_mul_size(m, l1, v1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_add_size(m, int32(8), v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_palloc0(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v1)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v10 << (uint(int32(2)) % 32)
				return v12
			}
		}
	}
}
func Fn13859(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v6 = F_SearchCatCache2(m, v5, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11)+91)))
			if v13 == int32(0) {
				v16 = F_heap_copytuple(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = v16
					F_ReleaseCatCache(m, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = v18
						return v22
					}
				}
			} else {
				v18 = v4
				F_ReleaseCatCache(m, v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v22 = v18
					return v22
				}
			}
		} else {
			v22 = v4
			return v22
		}
	}
}
func Fn13860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = F_SearchSysCache1(m, l6, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v168
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v24)
	v168 = v8
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, l5, v16)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_Fn13860_0), l4, l3)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v40 = v36 + v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	if v41 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L47
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_Fn13860[0]))
	v46 = int32(0)
	if v45 == v46 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_Fn13860[0]))
	if v88 == int32(0) {
		v153 = v8
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v84 == int32(0) {
		v153 = v8
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v84 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v52 <= int32(0) {
		v78 = v46
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v78
	goto L18
L23:
	;
	v55 = int32(0)
	if v55 < v52 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v58 = v52
	goto L26
L25:
	;
	v58 = v55
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v61 = int32(0)
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(int32(2))%32))))
	v70 = base.B2i32(v69 == v41)
	if v69 == v41 {
		v78 = v70
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v78 = v70
	goto L22
L29:
	;
	v72 = v61 + int32(1)
	if v72 != v58 {
		v61 = v72
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v91 < v92 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v97 = v91
	goto L36
L34:
	;
	goto L35
L35:
	;
	v153 = v8
	goto L14
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v97<<(uint(int32(2))%32))))
	v116 = *(*int32)(unsafe.Add(mBase, _c_Fn13860[1]))
	if v114 != v116 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v114 == v41 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v125 = v97 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v125 < v126 {
		v97 = v125
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v153 = int32(1)
	goto L14
L42:
	;
	goto L43
L43:
	;
	v120 = int32(0)
	v122 = F_SearchSysCacheExists(m, l2, v40+int32(4), v114, v120, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v122 != 0 {
		v153 = v8
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	goto L37
L47:
	;
	v168 = v153
	goto L1
}
func Fn13871(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13871[0])))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_Fn13871_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13871_1), l3, l2)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
		return int32(0)
	}
}
func Fn13877(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v22 = v20 & int32(1)
			if v22 != 0 {
				v23 = v14
			} else {
				v23 = v9 + int32(4)
			}
			if v20 == int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v29 == int32(18) {
					v32 = int32(16)
				} else {
					v32 = int32(0)
				}
				if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v39 = int32(4)
				} else {
					v39 = v32
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v22 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v52 = v16 + v51
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v57 = v55 & v51
			if v57 != 0 {
				v58 = v52
			} else {
				v58 = v16 + int32(4)
			}
			if v55 == int32(1) {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
				if v64 == int32(18) {
					v67 = int32(16)
				} else {
					v67 = int32(0)
				}
				if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v74 = int32(4)
				} else {
					v74 = v67
				}
				v85 = v74
			} else {
				v75 = int32(1)
				if v57 != 0 {
					v85 = int32(base.Ui32(v55)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v86 = F_dotrim(m, v23, v50, v58, v85, l2, l1)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				return v86
			}
		}
	}
}
func Fn13882(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v44 int64
	_ = v44
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = int64(63)
	v16 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v23 = int64(32)
	v24 = int64(base.Ui64(v16) >> (uint(v23) % 64))
	v26 = int64(base.Ui64(v13) >> (uint(v23) % 64))
	v29 = int64(4294967295)
	v30 = v16 & v29
	v32 = v13 & v29
	v33 = v30 * v32
	v37 = int64(base.Ui64(v33)>>(uint(v23)%64)) + v30*v26
	v44 = v32*v24 + v37&v29
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v13*(v16>>(uint(v14)%64)) + v13>>(uint(v14)%64)*v16 + v24*v26 + int64(base.Ui64(v37)>>(uint(v23)%64)) + int64(base.Ui64(v44)>>(uint(v23)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v33&v29 | v44<<(uint(v23)%64)
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v55 != v56>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
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
		v74 = F_Int64GetDatum(m, v56)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v74
		}
	}
}
func Fn13884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v45 int64
	_ = v45
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = int64(63)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v24 = int64(32)
	v25 = int64(base.Ui64(v17) >> (uint(v24) % 64))
	v27 = int64(base.Ui64(v13) >> (uint(v24) % 64))
	v30 = int64(4294967295)
	v31 = v17 & v30
	v33 = v13 & v30
	v34 = v31 * v33
	v38 = int64(base.Ui64(v34)>>(uint(v24)%64)) + v31*v27
	v45 = v33*v25 + v38&v30
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v13*(v17>>(uint(v14)%64)) + v13>>(uint(v14)%64)*v17 + v25*v27 + int64(base.Ui64(v38)>>(uint(v24)%64)) + int64(base.Ui64(v45)>>(uint(v24)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v34&v30 | v45<<(uint(v24)%64)
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v56 != v57>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
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
		v75 = F_Int64GetDatum(m, v57)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v75
		}
	}
}
func Fn13888(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 == l3 {
			return int32(1)
		} else {
			v13 = F_expression_tree_walker_impl(m, l0, l2, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func Fn13897(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v6)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v16&int32(1) == v6 {
		v21 = int32(4)
		v25 = l2 + l1<<(uint(v21)%32) + v21
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		if v26 < int32(0) {
			v68 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v74 = v68
				m.G0 = v11 + int32(16)
				return v74
			}
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v31 = v15 + v29 + v26
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)))
			if v32 != int32(1) {
				v74 = v31
				m.G0 = v11 + int32(16)
				return v74
			} else {
				v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
				switch v35&int32(_a_Fn13897_0) - int32(1) {
				case 0:
					v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31))))
					v74 = v40
					m.G0 = v11 + int32(16)
					return v74
				case 1:
					v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31))))
					v74 = v41
					m.G0 = v11 + int32(16)
					return v74
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v35
						F_errmsg_internal(m, int32(_a_Fn13897_1), v11)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, int32(70), int32(_a_Fn13897_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v74 = v42
					m.G0 = v11 + int32(16)
					return v74
				}
			}
		}
	} else {
		v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+23)))
		v58 = int32(1)
		if int32(base.Ui32(v57)>>(uint(l1-v58)%32))&v58 != 0 {
			v68 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v74 = v68
				m.G0 = v11 + int32(16)
				return v74
			}
		} else {
			v63 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v63)
			v74 = int32(0)
			m.G0 = v11 + int32(16)
			return v74
		}
	}
}
func Fn13899(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l9
	v18 = int32(1)
	v22 = F_LookupFuncName(m, l0, v18, v15+int32(44), v18)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 != 0 {
			v26 = F_get_func_rettype(m, v22)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != l8 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = F_NameListToString(m, l0)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v83
								F_errmsg(m, l3, v15+int32(32))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13899_0), l2, l1)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
				} else {
					v29 = F_func_volatile(m, v22)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 != int32(118) {
							m.G0 = v15 + int32(48)
							return v22
						} else {
							v35 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									m.G0 = v15 + int32(48)
									return v22
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v42 = F_NameListToString(m, l0)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v42
											F_errmsg(m, l7, v15+int32(16))
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_Fn13899_0), l6, l1)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return int32(0)
												} else {
													m.G0 = v15 + int32(48)
													return v22
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v67 = F_func_signature_string(m, l0, int32(1), int32(0), v15+int32(44))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v67
						F_errmsg(m, int32(_a_Fn13899_1), v15)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13899_0), l5, l1)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
func Fn13907(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v48)+12)))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v54 = F_gbt_var_consistent(m, v22, v15, v19, v46, v50&int32(1), l1, v53)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			m.G0 = v11 + int32(16)
			return v54
		}
	}
}
func Fn13912(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16 + l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v16
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v28)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v10+int32(4), v12, v10+int32(14), v30&int32(1), l1, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(16)
		return v34
	}
}
func Fn13927(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_palloc(m, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v26 = F_palloc(m, int32(16))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v28
			v31 = F_palloc(m, v28)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v31
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v17
				*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v16)
				v42 = F_palloc(m, int32(4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v42))) = v26
					v47 = v16 & int32(_a_Fn13927_0)
					switch v47 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = l1
						v67 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v67)
						m.G0 = v13 + int32(16)
						return v21
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					case 3, 4:
						v50 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v50)
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = v47
							F_errmsg_internal(m, int32(_a_Fn13927_1), v13)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13927_2), int32(97), int32(_a_Fn13927_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
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
}
func Fn13932(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v16 == int32(-1) {
			m.G0 = v9 + int32(16)
			return v12
		} else {
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
			if v16 == v19 {
				m.G0 = v9 + int32(16)
				return v12
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
						F_errmsg(m, int32(_a_Fn13932_0), v9)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l2, l1, int32(_a_Fn13932_1))
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
func Fn13934(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
		v23 = F_ArrayGetIntegerTypmods(m, v17, v14+int32(12))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			if v25 == int32(1) {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				if v28 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, l7, int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, l3, l6, l1)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
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
					if base.Ui32(l9) <= base.Ui32(v28) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = l5
								F_errmsg(m, l4, v14)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l3, l2, l1)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
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
						m.G0 = v14 + int32(16)
						return v28
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_Fn13934_0), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l3, l8, l1)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
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
func Fn13943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v14 = int32(0)
	goto L3
L3:
	;
	v15 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v11 = F_get_opclass_input_type(m, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v14 = v11
	goto L3
L7:
	;
	return
L8:
	;
	if v15 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v24 = l1
	v25 = int32(0)
	v29 = v14
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v25<<(uint(int32(2))%32))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L7
L13:
	;
	v63 = v25 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v63 < v64 {
		v24 = v59
		v25 = v63
		v29 = v61
		goto L11
	} else {
		goto L27
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l0
	v57 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v57)
	v59 = v24
	v61 = v29
	goto L13
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v37 != int32(1) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	if v40 != v41 {
		goto L14
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	if v40 != v29 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v44 = F_opclass_for_family_datatype(m, l4, l0, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v46 = v24
	v47 = v29
	goto L22
L22:
	;
	if v46 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v46 = v44
	v47 = v40
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v46
	v49 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v49)
	v59 = v46
	v61 = v47
	goto L13
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l0
	v52 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v52)
	v59 = int32(0)
	v61 = v47
	goto L13
L27:
	;
	goto L12
}
func Fn13956(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_JsonbExtractScalar(m, v13+int32(4), v10+int32(12))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			if v21 != 0 {
				switch v23 {
				case 0:
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v24 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v41 = int32(0)
							m.G0 = v10 + int32(32)
							return v41
						}
					} else {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v41 = int32(0)
						m.G0 = v10 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v23, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					v35 = F_DirectFunctionCall1Coll(m, l2, int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v13 == v37 {
							v41 = v35
							m.G0 = v10 + int32(32)
							return v41
						} else {
							F_pfree(m, v13)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v10 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v23, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
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
func Fn13961(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = int32(16711935)
	v10 = int32(8)
	v12 = int32(24)
	v16 = base.I32_rotr(v7&v8, v10) | base.I32_rotr(v7, v12)&v8
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v27 = base.I32_rotr(v18&v8, v10) | base.I32_rotr(v18, v12)&v8
	if base.Ui32(v16) < base.Ui32(v27) {
		v55 = l1
	} else {
		if base.Ui32(v27) < base.Ui32(v16) {
			v55 = int32(1)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v32 = int32(16711935)
			v34 = int32(8)
			v36 = int32(24)
			v40 = base.I32_rotr(v31&v32, v34) | base.I32_rotr(v31, v36)&v32
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v50 = base.I32_rotr(v41&v32, v34) | base.I32_rotr(v41, v36)&v32
			if base.Ui32(v40) < base.Ui32(v50) {
				v55 = l1
			} else {
				v55 = base.B2i32(base.Ui32(v50) < base.Ui32(v40))
			}
		}
	}
	return v55
}
func Fn13967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(7), l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = F_mic2latin_with_table(m, v9, v8, v13, l3, l2, l1, base.B2i32(v10 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func Fn13970(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_strlen(m, v6)
		mBase = m.M
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_RE_compile_and_cache(m, v8, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = F_palloc(m, v12<<(uint(int32(2))%32)+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_pg_mb2wchar_with_len(m, v6, v20, v12)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(0)
					v27 = F_RE_wchar_execute(m, v20, v22, v24, v24, v24)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return v27 ^ int32(1)
						}
					}
				}
			}
		}
	}
}
func Fn13976(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = F_strlen(m, l0)
	mBase = m.M
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L35
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L31
	}
L3:
	;
	if v61 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v61 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = v15
	v24 = l0
	v25 = v16
	v26 = v22
	goto L11
L8:
	;
	v49 = l0
	v53 = int32(0)
	goto L9
L9:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v61 = v53 - v54
	goto L3
L10:
	;
	v49 = v44
	v53 = v46
	goto L9
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(v26 != v28)|base.B2i32(v28 == int32(0)) != 0 {
		v44 = v24
		v46 = v26
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v44 = v38
	v46 = int32(0)
	goto L10
L13:
	;
	v34 = v25 - int32(1)
	if v34 == int32(0) {
		v44 = v24
		v46 = v26
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(1)
	v38 = v24 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v39 != 0 {
		v23 = v23 + v37
		v24 = v38
		v25 = v34
		v26 = v39
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v11 + int32(-4)
	v67 = v16 + v15
	v70 = F_sscanf(m, v67, l7, v11+int32(-32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L27
	}
L19:
	;
	return int32(0)
L20:
	;
	if v70 != int32(1) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v76 = int32(10)
	v77 = F___strchrnul(m, v67, v76)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == v76 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v83 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v83 = v77
	goto L25
L24:
	;
	v83 = int32(0)
	goto L25
L25:
	;
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	m.G0 = v13 - int32(-64)
	return v89
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l2
	F_errmsg(m, int32(_a_Fn13976_0), v11+int32(-16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_Fn13976_1), l6, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l2
	F_errmsg(m, int32(_a_Fn13976_0), v11+int32(-48))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_Fn13976_1), l5, l3)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg(m, int32(_a_Fn13976_0), v13)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_Fn13976_1), l4, l3)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13987(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v14 - int32(98) {
	case 0:
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
		v35 = int32(0)
		v37 = F_convert_case(m, l0, l1, l2, l3, l5, v34, v35, v35)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = v37
			m.G0 = v12 + int32(16)
			return v39
		}
	case 1:
		v17 = F_strlower_libc(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v39 = v17
			m.G0 = v12 + int32(16)
			return v39
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l6
			F_errmsg_internal(m, int32(_a_Fn13987_0), v12)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_Fn13987_1), l7, l6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
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
func Fn13989(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v11)
	v16 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v16)
	v19 = *(*int32)(unsafe.Add(mBase, _c_Fn13989[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
	v22 = F_LockAcquire(m, v8, l2, l1, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.B2i32(v22 != int32(0))
	}
}
func Fn13992(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v16 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v20 = F_pg_detoast_datum_packed(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = v20
					v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 != v9 {
							F_pfree(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v29 != v14 {
									F_pfree(m, v14)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return v23
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v22 == v36 {
												return v23
											} else {
												F_pfree(m, v22)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int32(0)
												} else {
													return v23
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = int32(0)
				v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return v23
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v22 == v36 {
									return v23
								} else {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v23
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func Fn13998(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 <= l3 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v15 = v14
		default:
			v15 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
	} else {
	}
	return int32(0)
}
func Fn14003(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 <= v12 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v370
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v364
	v370 = int32(1)
	goto L1
L3:
	;
	v364 = v20 - v9 + v151
	goto L2
L4:
	;
	v159 = v11 - v9
	v160 = v156 + v159
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160
	if v157 < v160 {
		goto L33
	} else {
		goto L34
	}
L5:
	;
	v156 = v9
	v157 = v12
	v158 = v10
	goto L4
L6:
	;
	goto L7
L7:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11-int32(1)))))
	if v17 != l1 {
		v156 = v9
		v157 = v12
		v158 = v10
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v20 = v11 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L11
L9:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v150 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L10:
	;
	v150 = v143
	goto L9
L11:
	;
	if v20 <= v35 {
		v143 = int32(-1)
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v143 = int32(0)
	goto L10
L13:
	;
	v52 = int32(1)
	v53 = v20 - v52
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36+v53))))
	v57 = v55 & int32(255)
	if base.B2i32(v53 == v35)|base.B2i32(int32(0) <= v55) != 0 {
		v115 = v57
		v119 = v52
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if int32(305) < v115 {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v64 = v57 & int32(63)
	v66 = v20 - int32(2)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v66))))
	v70 = v68 << (uint(int32(6)) % 32)
	if base.B2i32(v66 != v35)&base.B2i32(base.Ui32(v68) < base.Ui32(int32(192))) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v115 = v70&int32(1984) | v64
	v119 = int32(2)
	goto L14
L17:
	;
	goto L18
L18:
	;
	v83 = v70&int32(4032) | v64
	v85 = v20 - int32(3)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v85))))
	if base.B2i32(v85 != v35)&base.B2i32(base.Ui32(v87) < base.Ui32(int32(224))) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v115 = v87<<(uint(int32(12))%32)&int32(_a_Fn14003_0) | v83
	v119 = int32(3)
	goto L14
L20:
	;
	goto L21
L21:
	;
	v105 = int32(4)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v36-v105))))
	v115 = v87<<(uint(int32(12))%32)&int32(_a_Fn14003_1) | v107&int32(7)<<(uint(int32(18))%32) | v83
	v119 = v105
	goto L14
L22:
	;
	v150 = v119
	goto L9
L23:
	;
	goto L24
L24:
	;
	v121 = v115 - int32(97)
	if v121 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v150 = v119
	goto L9
L26:
	;
	goto L27
L27:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v121)>>(uint(int32(3))%32)))+uint32(_c_Fn14003[0]))))
	if int32(base.Ui32(v127)>>(uint(v121&int32(7))%32))&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v150 = v119
	goto L9
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 - v119
	goto L31
L31:
	;
	goto L12
L32:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = v151
	v157 = v155
	v158 = v154
	goto L4
L33:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v158-int32(1)))))
	if v167 == l1 {
		v370 = int32(0)
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v170 = int32(0)
	goto L39
L36:
	;
	goto L35
L37:
	;
	if v222 < int32(0) {
		v370 = v170
		goto L1
	} else {
		goto L56
	}
L39:
	;
	goto L40
L40:
	;
	goto L41
L41:
	;
	v177 = v160
	v179 = int32(1)
	goto L44
L43:
	;
	v222 = v204
	goto L37
L44:
	;
	if v177 <= v157 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L43
L46:
	;
	v222 = int32(-1)
	goto L37
L47:
	;
	goto L48
L48:
	;
	v184 = v177 - int32(1)
	v186 = int32(*(*int8)(unsafe.Add(mBase, uint32(v158+v184))))
	if base.B2i32(int32(0) <= v186)|base.B2i32(v184 <= v157) != 0 {
		v204 = v184
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v208 = int32(1)
	if v208 < v179 {
		v177 = v204
		v179 = v179 - v208
		goto L44
	} else {
		goto L55
	}
L50:
	;
	v192 = v184
	goto L51
L51:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158+v192))))
	if base.Ui32(int32(191)) < base.Ui32(v197) {
		v204 = v192
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v204 = v157
	goto L49
L53:
	;
	v201 = v192 - int32(1)
	if v157 < v201 {
		v192 = v201
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L45
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L59
L57:
	;
	if v354 != 0 {
		v370 = v170
		goto L1
	} else {
		goto L80
	}
L58:
	;
	v354 = v347
	goto L57
L59:
	;
	if v222 <= v239 {
		v347 = int32(-1)
		goto L58
	} else {
		goto L61
	}
L60:
	;
	v347 = int32(0)
	goto L58
L61:
	;
	v256 = int32(1)
	v257 = v222 - v256
	v259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v240+v257))))
	v261 = v259 & int32(255)
	if base.B2i32(v257 == v239)|base.B2i32(int32(0) <= v259) != 0 {
		v319 = v261
		v323 = v256
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if int32(305) < v319 {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	v268 = v261 & int32(63)
	v270 = v222 - int32(2)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+v270))))
	v274 = v272 << (uint(int32(6)) % 32)
	if base.B2i32(v270 != v239)&base.B2i32(base.Ui32(v272) < base.Ui32(int32(192))) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v319 = v274&int32(1984) | v268
	v323 = int32(2)
	goto L62
L65:
	;
	goto L66
L66:
	;
	v287 = v274&int32(4032) | v268
	v289 = v222 - int32(3)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+v289))))
	if base.B2i32(v289 != v239)&base.B2i32(base.Ui32(v291) < base.Ui32(int32(224))) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v319 = v291<<(uint(int32(12))%32)&int32(_a_Fn14003_0) | v287
	v323 = int32(3)
	goto L62
L68:
	;
	goto L69
L69:
	;
	v309 = int32(4)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+v240-v309))))
	v319 = v291<<(uint(int32(12))%32)&int32(_a_Fn14003_1) | v311&int32(7)<<(uint(int32(18))%32) | v287
	v323 = v309
	goto L62
L70:
	;
	v354 = v323
	goto L57
L71:
	;
	goto L72
L72:
	;
	v325 = v319 - int32(97)
	if v325 < int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v354 = v323
	goto L57
L74:
	;
	goto L75
L75:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v325)>>(uint(int32(3))%32)))+uint32(_c_Fn14003[0]))))
	if int32(base.Ui32(v331)>>(uint(v325&int32(7))%32))&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v354 = v323
	goto L57
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222 - v323
	goto L79
L79:
	;
	goto L60
L80:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v364 = v355 + v159
	goto L2
}
func Fn14016(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v4 = l3
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 != 0 {
			v18 = F_array_contains_nulls(m, v13)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_Fn14016_0), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn14016_1), l2, l1)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
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
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v26 != 0 {
							v34 = v26
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v34+v13, v23, v10+int32(15))
						mBase = m.M
						m.G0 = v10 + int32(16)
						return v13
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v26 != 0 {
					v34 = v26
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				F_isort(m, v34+v13, v23, v10+int32(15))
				mBase = m.M
				m.G0 = v10 + int32(16)
				return v13
			}
		}
	}
}
func Fn14021(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if base.Ui64(int64(2)) <= base.Ui64(v15-int64(9223372036854775807)) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
		if v21 != 0 {
			if v21 != int32(2147483647) {
				if v21 != int32(-2147483648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
								F_errmsg(m, int32(_a_Fn14021_0), v12)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn14021_1), l2, l1)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
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
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
					if v26 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
									F_errmsg(m, int32(_a_Fn14021_0), v12)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn14021_1), l2, l1)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
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
					} else {
						v29 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
						if v29 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v49 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
										F_errmsg(m, int32(_a_Fn14021_2), v12+int32(16))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_Fn14021_1), l5, l1)
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
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
										F_errmsg(m, int32(_a_Fn14021_0), v12)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_Fn14021_1), l2, l1)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
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
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
				if v32 != int32(2147483647) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
								F_errmsg(m, int32(_a_Fn14021_0), v12)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn14021_1), l2, l1)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
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
				} else {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
					if v35 != int64(9223372036854775807) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
									F_errmsg(m, int32(_a_Fn14021_0), v12)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn14021_1), l2, l1)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v49 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
									F_errmsg(m, int32(_a_Fn14021_2), v12+int32(16))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn14021_1), l5, l1)
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
						}
					}
				}
			}
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			if v60 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						v103 = F_DirectFunctionCall1Coll(m, int32(1273), int32(0), v20)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
							F_errmsg(m, int32(_a_Fn14021_0), v12)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn14021_1), l2, l1)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
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
			} else {
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
				v62 = base.I64_div_s(v61, l4)
				v66 = base.I64_extend32_s(v62)*int64(-1000000) + v15
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v66+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_Fn14021_3), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn14021_1), l3, l1)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
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
					v73 = v66
					v74 = F_Int64GetDatum(m, v73)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(32)
						return v74
					}
				}
			}
		}
	} else {
		v73 = v15
		v74 = F_Int64GetDatum(m, v73)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v12 + int32(32)
			return v74
		}
	}
}
func Fn14027(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_text_to_cstring(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_Fn14027[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v18
			v21 = *(*int64)(unsafe.Add(mBase, _c_Fn14027[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v21
			v26 = F_DirectInputFunctionCallSafe(m, l1, v15, int32(-1), v8, v8+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					v33 = v32
				}
				m.G0 = v8 + int32(16)
				return v33
			}
		}
	}
}
func Fn14032(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v24) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errmsg(m, int32(_a_Fn14032_0), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_Fn14032_1), l2, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	return v144
L10:
	;
	v144 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_Fn14032[0]))
	if v35 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v144 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_Fn14032[1]))
	if v39 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v144 = v136
	goto L9
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_Fn14032[2]))
	if v43 == int32(0) {
		v136 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_Fn14032[3]))
	v107 = int32(0)
	v109 = v39 - int32(1)
	goto L39
L20:
	;
	v48 = v43
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v53 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v136 = int32(0)
	goto L16
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	if v100 != 0 {
		v48 = v100
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v56 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v59 = int32(1)
	if v24 == v56 {
		v136 = v59
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v63 = v61 - int32(1)
	if v63 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v68 = int32(0)
	v70 = v63
	goto L28
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v76 = int32(2)
	v77 = base.I32_div_s(v70-v68, v76)
	v78 = v77 + v68
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74+v78<<(uint(v76)%32))))
	if v82 == v24 {
		v136 = v59
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v86 = F_TransactionIdPrecedes(m, v82, v24)
	mBase = m.M
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v87 = v78 + int32(1)
	goto L33
L32:
	;
	v87 = v68
	goto L33
L33:
	;
	if v86 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v90 = v70
	goto L36
L35:
	;
	v90 = v78 - int32(1)
	goto L36
L36:
	;
	if v87 <= v90 {
		v68 = v87
		v70 = v90
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	goto L22
L39:
	;
	v114 = int32(2)
	v115 = base.I32_div_s(v109-v107, v114)
	v116 = v115 + v107
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v105+v116<<(uint(v114)%32))))
	v121 = base.B2i32(v120 == v24)
	if v120 == v24 {
		v136 = v121
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v136 = v121
	goto L16
L41:
	;
	v124 = base.B2i32(base.Ui32(v120) < base.Ui32(v24))
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v125 = v116 + int32(1)
	goto L44
L43:
	;
	v125 = v107
	goto L44
L44:
	;
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v128 = v109
	goto L47
L46:
	;
	v128 = v116 - int32(1)
	goto L47
L47:
	;
	if v125 <= v128 {
		v107 = v125
		v109 = v128
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
}
func Fn14038(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+29)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v20
	v24 = F_DirectFunctionCall1Coll(m, int32(3376), int32(0), v10)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(48)
		return v24
	}
}
func Fn14041(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 float32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v23 = v21 & int32(1)
			if v23 != 0 {
				v24 = v17
			} else {
				v24 = v10 + int32(4)
			}
			if v21 == int32(1) {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v30 == int32(18) {
					v33 = int32(16)
				} else {
					v33 = int32(0)
				}
				if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v40 = int32(4)
				} else {
					v40 = v33
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v52 = int32(1)
			v53 = v19 + v52
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v58 = v56 & v52
			if v58 != 0 {
				v59 = v53
			} else {
				v59 = v19 + int32(4)
			}
			if v56 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v86 = v75
			} else {
				v76 = int32(1)
				if v58 != 0 {
					v86 = int32(base.Ui32(v56)>>(uint(v76)%32)) - v76
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_calc_word_similarity(m, v24, v51, v59, v86, l1)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v89 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v93 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(v87)
							}
						} else {
							return base.I32_reinterpret_f32(v87)
						}
					}
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v93 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v87)
						}
					} else {
						return base.I32_reinterpret_f32(v87)
					}
				}
			}
		}
	}
}
