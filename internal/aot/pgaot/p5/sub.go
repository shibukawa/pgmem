package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitSubPlanExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v5 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v20 = v5
	v23 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L16
	} else {
		goto L41
	}
L4:
	;
	v26 = int32(0)
	if v14 == v26 {
		v36 = v26
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v13 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v30 <= v20 {
		v36 = int32(0)
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v36 = v32 + v20<<(uint(int32(2))%32)
	goto L6
L9:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v45+v20<<(uint(int32(2))%32))))
	F_ExecInitExprRec(m, v102, l1, l2, l3)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L16
	} else {
		goto L29
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v49 = F_ExecInitSubPlan(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v46 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if base.B2i32(v36 == int32(0))|base.B2i32(v42 <= v20) != 0 {
		v46 = v23
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v45 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v46 = v23
	goto L10
L16:
	;
	return
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+48))
	v53 = F_lappend(m, v52, v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+48)) = v53
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v57 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v79 + int32(1)
	v85 = v78 + v79*int32(40)
	v86 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(103)
	*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v86
	return
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v76
	v78 = v76
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(16)
	v63 = F_palloc(m, int32(640))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v65 != v57 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v76 = v63
	goto L20
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v78 = v67
	goto L19
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v57 << (uint(int32(1)) % 32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v74 = F_repalloc(m, v71, v57*int32(80))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v76 = v74
	goto L20
L29:
	;
	v105 = F_exprType(m, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v107 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v130 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v129 + v130
	v135 = v128 + v129*int32(40)
	v136 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v135)+24)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v135)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v135)+16)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(55)
	*(*int64)(unsafe.Add(mBase, uint32(v135)+32)) = v136
	v20 = v20 + v130
	v23 = v105
	goto L4
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v126
	v128 = v126
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(16)
	v113 = F_palloc(m, int32(640))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L16
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v115 != v107 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v126 = v113
	goto L32
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v128 = v117
	goto L31
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v107 << (uint(int32(1)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v124 = F_repalloc(m, v121, v107*int32(80))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v126 = v124
	goto L32
L41:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInitSubPlanExpr_0), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_ExecInitSubPlanExpr_1), int32(2831), int32(_a_F_ExecInitSubPlanExpr_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__equalSubLink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != v5 {
		v28 = v3
		return v28
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v7 != v8 {
			v28 = v3
			return v28
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v12 = F_equal(m, v10, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v12 == int32(0) {
					v28 = v3
					return v28
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v20 = F_equal(m, v18, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if v20 == int32(0) {
							v28 = v3
							return v28
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v26 = F_equal(m, v24, v25)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								v28 = v26
								return v28
							}
						}
					}
				}
			}
		}
	}
}
func F_make_sub_restrictinfos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v11 = int32(0)
	if l1 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v126 = F_make_plain_restrictinfo(m, l0, l1, int32(0), l2, l3, l4, l5, l6, l7, l8, l9)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L30
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v16 != int32(21) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v19 {
	case 0:
		goto L4
	case 1:
		goto L5
	default:
		goto L1
	}
L4:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v73 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v20 == int32(0) {
		v65 = v11
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v68 = F_make_orclause(m, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L11
	} else {
		goto L15
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 <= int32(0) {
		v65 = v11
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v36 = v11
	v37 = v11
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v37<<(uint(int32(2))%32))))
	v45 = F_make_sub_restrictinfos(m, l0, v43, l2, l3, l4, l5, l6, int32(0), l8, l9)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v65 = v49
	goto L6
L11:
	;
	return int32(0)
L12:
	;
	v49 = F_lappend(m, v36, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v52 = v37 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v52 < v53 {
		v36 = v49
		v37 = v52
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v70 = F_make_plain_restrictinfo(m, l0, l1, v68, l2, l3, l4, l5, l6, l7, l8, l9)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	return v70
L17:
	;
	v77 = F_make_andclause(m, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if int32(0) < v80 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	return v77
L21:
	;
	v93 = v11
	v94 = v11
	goto L24
L22:
	;
	v119 = v11
	goto L23
L23:
	;
	v122 = F_make_andclause(m, v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L29
	}
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v94<<(uint(int32(2))%32))))
	v101 = F_make_sub_restrictinfos(m, l0, v100, l2, l3, l4, l5, l6, l7, l8, l9)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L26
	}
L25:
	;
	v119 = v103
	goto L23
L26:
	;
	v103 = F_lappend(m, v93, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v106 = v94 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v106 < v107 {
		v93 = v103
		v94 = v106
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	return v122
L30:
	;
	return v126
}
