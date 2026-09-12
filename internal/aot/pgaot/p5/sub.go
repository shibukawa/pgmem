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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
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
	v22 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L17
	} else {
		goto L42
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
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	F_ExecInitExprRec(m, v101, l1, l2, l3)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L30
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v51 = F_ExecInitSubPlan(m, l0, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v48 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v40 <= v20 {
		v48 = v22
		goto L10
	} else {
		goto L14
	}
L14:
	;
	if v36 == int32(0) {
		v48 = v22
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v47 = v44 + v20<<(uint(int32(2))%32)
	if v47 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v48 = v22
	goto L10
L17:
	;
	return
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
	v55 = F_lappend(m, v54, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+48)) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v59 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v81 + int32(1)
	v87 = v80 + v81*int32(40)
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+24)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v87)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(103)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+32)) = v88
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v78
	v80 = v78
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(16)
	v65 = F_palloc(m, int32(640))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v67 != v59 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v78 = v65
	goto L21
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v80 = v69
	goto L20
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v59 << (uint(int32(1)) % 32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v76 = F_repalloc(m, v73, v59*int32(80))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v78 = v76
	goto L21
L30:
	;
	v104 = F_exprType(m, v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v106 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v129 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v128 + v129
	v134 = v127 + v128*int32(40)
	v135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v134)+24)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v134)+20)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v134)+16)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(55)
	*(*int64)(unsafe.Add(mBase, uint32(v134)+32)) = v135
	v20 = v20 + v129
	v22 = v104
	goto L4
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v125
	v127 = v125
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(16)
	v112 = F_palloc(m, int32(640))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v114 != v106 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v125 = v112
	goto L33
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v127 = v116
	goto L32
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v106 << (uint(int32(1)) % 32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v123 = F_repalloc(m, v120, v106*int32(80))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	v125 = v123
	goto L33
L42:
	;
	F_errmsg_internal(m, int32(283936), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(496219), int32(2831), int32(207641))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
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
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	v11 = int32(0)
	if l1 == v11 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v122 = F_make_orclause(m, v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L14
	} else {
		goto L32
	}
L2:
	;
	v106 = F_make_plain_restrictinfo(m, l0, l1, int32(0), l2, l3, l4, l5, l6, l7, l8, l9)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L14
	} else {
		goto L31
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v16 != int32(21) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v19 {
	case 0:
		goto L5
	case 1:
		goto L6
	default:
		goto L2
	}
L5:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v53 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v34 = v11
	v35 = v11
	goto L12
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(0) < v21 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v119 = v11
	goto L1
L11:
	;
	goto L10
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v35<<(uint(int32(2))%32))))
	v43 = F_make_sub_restrictinfos(m, l0, v41, l2, l3, l4, l5, l6, int32(0), l8, l9)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v119 = v47
	goto L1
L14:
	;
	return int32(0)
L15:
	;
	v47 = F_lappend(m, v34, v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v50 = v35 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v50 < v51 {
		v34 = v47
		v35 = v50
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v57 = F_make_andclause(m, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if int32(0) < v60 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	return v57
L22:
	;
	v73 = v11
	v74 = v11
	goto L25
L23:
	;
	v99 = v11
	goto L24
L24:
	;
	v102 = F_make_andclause(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L14
	} else {
		goto L30
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v74<<(uint(int32(2))%32))))
	v81 = F_make_sub_restrictinfos(m, l0, v80, l2, l3, l4, l5, l6, l7, l8, l9)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L14
	} else {
		goto L27
	}
L26:
	;
	v99 = v83
	goto L24
L27:
	;
	v83 = F_lappend(m, v73, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v86 = v74 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v86 < v87 {
		v73 = v83
		v74 = v86
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	return v102
L31:
	;
	return v106
L32:
	;
	v124 = F_make_plain_restrictinfo(m, l0, l1, v122, l2, l3, l4, l5, l6, l7, l8, l9)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	return v124
}
