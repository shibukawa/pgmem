package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_call_enum_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v14 == int32(0) {
		v123 = v13
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L36
	}
L2:
	;
	m.G0 = v11 + int32(80)
	return v123
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[336])) = int32(50856066)
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[334])) = v21
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v21
	*(*int32)(unsafe.Add(mBase, _consts[1294])) = v21
	v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v29 != 0 {
		v123 = v13
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = F_errstart(m, l4, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	F_errcode(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_FlushErrorState(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L35
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if v86 != 0 {
		goto L26
	} else {
		goto L27
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v41
	F_errmsg_internal(m, int32(195849), v11-int32(-64))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v49 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = v49
	goto L18
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v61 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v52
	F_errmsg(m, int32(685962), v11+int32(48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L25
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v48 != v64 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = v59 + int32(12)
	if v67 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L19
L24:
	;
	v59 = v67
	goto L18
L25:
	;
	goto L12
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v86
	F_errdetail_internal(m, int32(195849), v11+int32(32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[1294]))
	if v94 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v94
	F_errhint(m, int32(195849), v11+int32(16))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_errfinish(m, int32(476640), int32(6989), int32(299835))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L10
L35:
	;
	v123 = int32(0)
	goto L2
L36:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v48
	F_errmsg_internal(m, int32(171530), v11)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(476640), int32(3036), int32(328051))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_enum_in(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v13&int32(3) == int32(0) {
		v37 = v13
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v130
L2:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v70) {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v70 = v62 - v13
	goto L2
L4:
	;
	v41 = v37
	goto L13
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v21 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v70 = int32(0)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v26 = v13
	goto L9
L9:
	;
	v30 = v26 + int32(1)
	if v30&int32(3) == int32(0) {
		v37 = v30
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v62 = v30
	goto L3
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 != 0 {
		v26 = v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v41 = v41 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v56 = v41
	goto L16
L15:
	;
	goto L14
L16:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		v56 = v56 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v62 = v56
	goto L3
L18:
	;
	goto L17
L19:
	;
	v73 = int32(0)
	v74 = F_errsave_start(m, v12)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v96 = F_SearchSysCache2(m, int32(24), v11, v13)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L22
	} else {
		goto L29
	}
L22:
	;
	return int32(0)
L23:
	;
	if v74 == int32(0) {
		v130 = v73
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v83 = F_format_type_be(m, v11)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v83
	F_errmsg(m, int32(683486), v9)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, v12, int32(474039), int32(123), int32(265794))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v130 = v73
	goto L1
L29:
	;
	if v96 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v100 = int32(0)
	v101 = F_errsave_start(m, v12)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L22
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_check_safe_enum_use(m, v96)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L22
	} else {
		goto L39
	}
L33:
	;
	if v101 == int32(0) {
		v130 = v100
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	v108 = F_format_type_be(m, v11)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v108
	F_errmsg(m, int32(683486), v9+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	F_errsave_finish(m, v12, int32(474039), int32(133), int32(265794))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v130 = v100
	goto L1
L39:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124+v125)))
	F_ReleaseCatCache(m, v96)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v130 = v127
	goto L1
}
func F_enum_larger(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_enum_cmp_internal(m, v4, v5, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if int32(0) < v6 {
			v12 = v4
		} else {
			v12 = v5
		}
		return v12
	}
}
func F_enum_range_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = F_get_fn_expr_argtype(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(351393), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(474039), int32(540), int32(290229))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
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
			v26 = int32(0)
			v28 = F_enum_range_internal(m, v4, v26, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v28
			}
		}
	}
}
func F_enum_range_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	F_ScanKeyInit(m, v16, int32(2), int32(3), int32(184), l0)
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
	v30 = F_table_open(m, int32(3501), int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = F_index_open(m, int32(3534), int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = F_systable_beginscan_ordered(m, v30, v34, int32(0), int32(1), v16)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = F_palloc(m, int32(256))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v46 = int32(0)
	v47 = int32(64)
	v48 = v41
	v52 = base.B2i32(l1 == int32(0))
	goto L7
L7:
	;
	v57 = F_systable_getnext_ordered(m, v38, int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_systable_endscan_ordered(m, v38)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L9:
	;
	if v57 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59+v60)))
	v64 = v52 | base.B2i32(l1 == v62)
	if v64&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v90 = v46
	v92 = v48
	goto L12
L12:
	;
	goto L8
L13:
	;
	F_check_safe_enum_use(m, v57)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v84 = v46
	v85 = v47
	v86 = v48
	goto L15
L15:
	;
	if l2 == int32(0) {
		v46 = v84
		v47 = v85
		v48 = v86
		v52 = v64
		goto L7
	} else {
		goto L21
	}
L16:
	;
	if v47 <= v46 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v72 = F_repalloc(m, v48, v47<<(uint(int32(3))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v76 = v47
	v77 = v48
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77+v46<<(uint(int32(2))%32)))) = v62
	v84 = v46 + int32(1)
	v85 = v76
	v86 = v77
	goto L15
L20:
	;
	v76 = v47 << (uint(int32(1)) % 32)
	v77 = v72
	goto L19
L21:
	;
	if l2 != v62 {
		v46 = v84
		v47 = v85
		v48 = v86
		v52 = v64
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v90 = v84
	v92 = v86
	goto L12
L23:
	;
	F_relation_close(m, v34, int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_sequence_close(m, v30, int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v106 = F_construct_array(m, v92, v90, l0, int32(4), int32(1), int32(105))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_pfree(m, v92)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v16 + int32(48)
	return v106
}
