package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_risparent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13865(m, l0, int32(_a_F__ltree_risparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	v3 = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.B2i32(v10 == v3)|base.B2i32(v13 == v3) != 0 {
		v140 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return (v140 + int32(1)) * (v10 - v13) * int32(10)
L2:
	;
	v17 = int32(8)
	v21 = l0 + v17
	v22 = l1 + v17
	v25 = v10
	v28 = v13
	goto L3
L3:
	;
	v30 = int32(2)
	v31 = v21 + v30
	v33 = v22 + v30
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	if base.Ui32(v34) < base.Ui32(v35) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v140 = v120
	goto L1
L5:
	;
	v120 = v25 - int32(1)
	if v25 < int32(2) {
		v140 = v120
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v37 = v34
	goto L8
L7:
	;
	v37 = v35
	goto L8
L8:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v37) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if v99 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v99 = int32(0)
	goto L9
L11:
	;
	v73 = v68
	v74 = v69
	v75 = v70
	goto L21
L12:
	;
	if (v31|v33)&int32(3) != 0 {
		v68 = v31
		v69 = v33
		v70 = v37
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v61 = v31
	v62 = v33
	v63 = v37
	goto L14
L14:
	;
	if v63 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v45 = v31
	v46 = v33
	v47 = v37
	goto L16
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != v51 {
		v68 = v45
		v69 = v46
		v70 = v47
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v61 = v56
	v62 = v54
	v63 = v58
	goto L14
L18:
	;
	v53 = int32(4)
	v54 = v46 + v53
	v56 = v45 + v53
	v58 = v47 - v53
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		v45 = v56
		v46 = v54
		v47 = v58
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v68 = v61
	v69 = v62
	v70 = v63
	goto L11
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 == v79 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v99 = v78 - v79
	goto L9
L23:
	;
	v81 = int32(1)
	v86 = v75 - v81
	if v86 != 0 {
		v73 = v73 + v81
		v74 = v74 + v81
		v75 = v86
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
	goto L10
L27:
	;
	if v34 == v35 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v99 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v103 = int32(10)
	return (v25*v103 + v103) * (v34 - v35)
L31:
	;
	v116 = int32(-10)
	goto L33
L32:
	;
	v116 = int32(10)
	goto L33
L33:
	;
	return (v25 + int32(1)) * v116
L34:
	;
	v123 = int32(9)
	v125 = int32(_a_F_ltree_compare_0)
	v133 = int32(1)
	if v133 < v28 {
		v21 = v21 + (v34+v123)&v125
		v22 = v22 + (v35+v123)&v125
		v25 = v120
		v28 = v28 - v133
		goto L3
	} else {
		goto L35
	}
L35:
	;
	goto L4
}
func F_ltree_gist_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_ltree_gist_in_0), int32(26), int32(_a_F_ltree_gist_in_1), int32(_a_F_ltree_gist_in_2), int32(_a_F_ltree_gist_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
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
	v23 = int32(0)
	if base.B2i32(v22 == v23)|base.B2i32(v21 == v23) != 0 {
		v152 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v180 != v14 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v179 = (v152 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	v28 = int32(8)
	v33 = v22
	v37 = v14 + v28
	v38 = v19 + v28
	v43 = v21
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v37 + v44
	v47 = v38 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v152 = v132
	goto L5
L9:
	;
	v132 = v33 - int32(1)
	if v33 < int32(2) {
		v152 = v132
		goto L5
	} else {
		goto L38
	}
L10:
	;
	v51 = v48
	goto L12
L11:
	;
	v51 = v49
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v113 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	v113 = int32(0)
	goto L13
L15:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L25
L16:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L18
L18:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L20
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L18
L22:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L15
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v113 = v92 - v93
	goto L13
L27:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	if v48 == v49 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v113 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v117 = int32(10)
	v179 = (v33*v117 + v117) * (v48 - v49)
	goto L4
L35:
	;
	v129 = int32(-10)
	goto L37
L36:
	;
	v129 = int32(10)
	goto L37
L37:
	;
	v179 = (v33 + int32(1)) * v129
	goto L4
L38:
	;
	v135 = int32(9)
	v137 = int32(_a_F_ltree_le_0)
	v145 = int32(1)
	if v145 < v43 {
		v33 = v132
		v37 = v37 + (v48+v135)&v137
		v38 = v38 + (v49+v135)&v137
		v43 = v43 - v145
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	F_pfree(m, v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v184 != v19 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_pfree(m, v19)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	return base.B2i32(v179 <= int32(0))
L47:
	;
	goto L46
}
