package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_RangeVarCallbackMaintainsTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 == int32(0) {
		m.G0 = v8 + int32(16)
		return
	} else {
		v12 = F_get_rel_relkind(m, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = v12 & int32(255)
			v17 = v15 - int32(109)
			v24 = int32(0)
			if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v17))|base.B2i32(int32(1)<<(uint(v17)%32)&int32(169) == v24) == v24 {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackMaintainsTable[0]))
				v32 = F_pg_class_aclcheck(m, l1, v30, int64(16384))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						m.G0 = v8 + int32(16)
						return
					} else {
						v36 = F_get_rel_relkind(m, l1)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							switch v36 - int32(73) {
							case 0, 32:
								v47 = int32(20)
								v49 = v47
							default:
								v47 = int32(41)
								v49 = v47
							case 10:
								v49 = int32(37)
							case 29:
								v49 = int32(18)
							case 36:
								v49 = int32(23)
							case 45:
								v49 = int32(51)
							}
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_aclcheck_error(m, v32, v49, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				if v15 == int32(0) {
					m.G0 = v8 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v62
							F_errmsg(m, int32(_a_F_RangeVarCallbackMaintainsTable_0), v8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RangeVarCallbackMaintainsTable_1), int32(_a_F_RangeVarCallbackMaintainsTable_2), int32(_a_F_RangeVarCallbackMaintainsTable_3))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
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
func F_addRangeTableEntryForSubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v4 = l3
	v5 = l4
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v22 = F_palloc0(m, int32(136))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(101)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l2
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v32 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v36 = F_makeAlias(m, int32(_a_F_addRangeTableEntryForSubquery_0), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v38 = v32
	goto L3
L8:
	;
	v38 = v36
	goto L3
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v41 = v40
	goto L11
L10:
	;
	v41 = v6
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v42 == int32(0) {
		v121 = v6
		v122 = v6
		v123 = v6
		v124 = v6
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v41 <= v121 {
		goto L33
	} else {
		goto L34
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 <= int32(0) {
		v121 = v6
		v122 = v6
		v123 = v6
		v124 = v6
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v50 = int32(0)
	v60 = v6
	v61 = v6
	v62 = v6
	v63 = v6
	goto L15
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v50<<(uint(int32(2))%32))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+26)))
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v121 = v101
	v122 = v102
	v123 = v103
	v124 = v104
	goto L12
L17:
	;
	v74 = v60 + int32(1)
	if v41 < v74 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v101 = v60
	v102 = v61
	v103 = v62
	v104 = v63
	goto L19
L19:
	;
	v107 = v50 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v107 < v108 {
		v50 = v107
		v60 = v101
		v61 = v102
		v62 = v103
		v63 = v104
		goto L15
	} else {
		goto L32
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v77 = F_pstrdup(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v87 = F_exprType(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v80 = F_makeString(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v82 = F_lappend(m, v79, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v82
	goto L22
L26:
	;
	v89 = F_lappend_oid(m, v62, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v92 = F_exprTypmod(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v94 = F_lappend_int(m, v63, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v97 = F_exprCollation(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v99 = F_lappend_oid(m, v61, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v101 = v74
	v102 = v99
	v103 = v89
	v104 = v94
	goto L19
L32:
	;
	goto L16
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+125)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+124)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v131 = F_lappend(m, v130, v22)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L41
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131
	if v131 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v136 = v134
	goto L39
L38:
	;
	v136 = int32(0)
	goto L39
L39:
	;
	v137 = F_buildNSItemFromLists(m, v22, v136, v123, v124, v122)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+20)) = uint8(base.B2i32(l2 != int32(0)))
	m.G0 = v19 + int32(16)
	return v137
L41:
	;
	F_errcode(m, int32(_a_F_addRangeTableEntryForSubquery_1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v153
	F_errmsg(m, int32(_a_F_addRangeTableEntryForSubquery_2), v19)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_addRangeTableEntryForSubquery_3), int32(1709), int32(_a_F_addRangeTableEntryForSubquery_4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_range_multirange(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13897(m, l0, int32(55))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_makeRangeVarFromAnyName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_palloc0(m, int32(28))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(3)
		if l0 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = F_NameListToString(m, l0)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
						F_errmsg(m, int32(_a_F_makeRangeVarFromAnyName_0), v8)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_scanner_errposition(m, l1, l2)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_makeRangeVarFromAnyName_1), int32(_a_F_makeRangeVarFromAnyName_2), int32(_a_F_makeRangeVarFromAnyName_3))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
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
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			switch v19 - int32(1) {
			case 0:
				*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(0)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v65 = v64
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				v69 = int32(112)
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v69)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
				m.G0 = v8 + int32(16)
				return v11
			case 1:
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v65 = v28 + int32(4)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				v69 = int32(112)
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v69)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
				m.G0 = v8 + int32(16)
				return v11
			case 2:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v33
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v65 = v39 + int32(8)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				v69 = int32(112)
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v69)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
				m.G0 = v8 + int32(16)
				return v11
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_NameListToString(m, l0)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
							F_errmsg(m, int32(_a_F_makeRangeVarFromAnyName_0), v8)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_scanner_errposition(m, l1, l2)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_makeRangeVarFromAnyName_1), int32(_a_F_makeRangeVarFromAnyName_2), int32(_a_F_makeRangeVarFromAnyName_3))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
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
}
func F_range_adjacent_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v12)>>(uint(int32(2))%32))-int32(1)))))
	if v18&int32(1) != 0 {
		v70 = v4
		m.G0 = v10 + int32(80)
		return v70
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v21 == int32(0) {
			v70 = v4
			m.G0 = v10 + int32(80)
			return v70
		} else {
			F_range_deserialize(m, l0, l1, v10+int32(72), v10-int32(-64), v10+int32(47))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v37 = v10 + int32(56)
				v39 = v10 + int32(48)
				F_multirange_get_bounds(m, l0, l2, int32(0), v37, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v10)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v42
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v10)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v44
					v51 = F_bounds_adjacent(m, l0, v10+int32(32), v10+int32(24))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v51 != 0 {
							v70 = int32(1)
							m.G0 = v10 + int32(80)
							return v70
						} else {
							if int32(2) <= v34 {
								F_multirange_get_bounds(m, l0, l2, v34-int32(1), v37, v39)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									v59 = *(*int64)(unsafe.Add(mBase, uint32(v10)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v59
									v61 = *(*int64)(unsafe.Add(mBase, uint32(v10)+72))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v61
									v67 = F_bounds_adjacent(m, l0, v10+int32(16), v10+int32(8))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v70 = v67
										m.G0 = v10 + int32(80)
										return v70
									}
								}
							} else {
								v59 = *(*int64)(unsafe.Add(mBase, uint32(v10)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v59
								v61 = *(*int64)(unsafe.Add(mBase, uint32(v10)+72))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v61
								v67 = F_bounds_adjacent(m, l0, v10+int32(16), v10+int32(8))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v70 = v67
									m.G0 = v10 + int32(80)
									return v70
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_range_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_range_deserialize(m, l2, v11, v8+int32(40), v8+int32(32), v8+int32(15))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		F_range_deserialize(m, l2, v10, v8+int32(24), v8+int32(16), v8+int32(14))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
			v33 = int32(1)
			if v30 != 0 {
				v36 = v30&v31 - v33
			} else {
				v36 = v33
			}
			if v30|v31&int32(1) != 0 {
				v163 = v36
				m.G0 = v8 + int32(48)
				return v163
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+44)))
				if v41 == int32(1) {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
					if v40&int32(1) != 0 {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
						if v47 == v44 {
							v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
							if v102 == int32(1) {
								v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
								if v101&int32(1) != 0 {
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
									if v108 == v105 {
										v163 = int32(0)
									} else {
										v112 = int32(1)
										if v105&v112 != 0 {
											v115 = int32(-1)
										} else {
											v115 = v112
										}
										v163 = v115
									}
								} else {
									v117 = int32(1)
									if v105&v117 != 0 {
										v120 = int32(-1)
									} else {
										v120 = v117
									}
									v163 = v120
								}
								m.G0 = v8 + int32(48)
								return v163
							} else {
								if v101&int32(1) != 0 {
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
									if v125 != 0 {
										v126 = int32(1)
									} else {
										v126 = int32(-1)
									}
									v163 = v126
									m.G0 = v8 + int32(48)
									return v163
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
									v132 = F_FunctionCall2Coll(m, l2+int32(212), v129, v130, v131)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										if v132 != 0 {
											v163 = v132
										} else {
											v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
											v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
											if v135 == int32(0) {
												v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
												if v134&int32(1) == int32(0) {
													v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
													if v143 == v138 {
														v163 = int32(0)
													} else {
														v146 = int32(1)
														if v138&v146 != 0 {
															v150 = v146
														} else {
															v150 = int32(-1)
														}
														v163 = v150
													}
												} else {
													v151 = int32(1)
													if v138&v151 != 0 {
														v155 = v151
													} else {
														v155 = int32(-1)
													}
													v163 = v155
												}
											} else {
												if v134&int32(1) != 0 {
													v163 = int32(0)
												} else {
													v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
													if v161 != 0 {
														v162 = int32(-1)
													} else {
														v162 = int32(1)
													}
													v163 = v162
												}
											}
										}
										m.G0 = v8 + int32(48)
										return v163
									}
								}
							}
						} else {
							v50 = int32(1)
							if v44&v50 != 0 {
								v53 = int32(-1)
							} else {
								v53 = v50
							}
							v163 = v53
							m.G0 = v8 + int32(48)
							return v163
						}
					} else {
						v55 = int32(1)
						if v44&v55 != 0 {
							v58 = int32(-1)
						} else {
							v58 = v55
						}
						v163 = v58
						m.G0 = v8 + int32(48)
						return v163
					}
				} else {
					if v40&int32(1) != 0 {
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
						if v63 != 0 {
							v64 = int32(1)
						} else {
							v64 = int32(-1)
						}
						v163 = v64
						m.G0 = v8 + int32(48)
						return v163
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
						v70 = F_FunctionCall2Coll(m, l2+int32(212), v67, v68, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							if v70 != 0 {
								v163 = v70
								m.G0 = v8 + int32(48)
								return v163
							} else {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+45)))
								if v73 == int32(0) {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
									if v72&int32(1) == int32(0) {
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
										if v81 == v76 {
											v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
											if v102 == int32(1) {
												v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
												if v101&int32(1) != 0 {
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
													if v108 == v105 {
														v163 = int32(0)
													} else {
														v112 = int32(1)
														if v105&v112 != 0 {
															v115 = int32(-1)
														} else {
															v115 = v112
														}
														v163 = v115
													}
												} else {
													v117 = int32(1)
													if v105&v117 != 0 {
														v120 = int32(-1)
													} else {
														v120 = v117
													}
													v163 = v120
												}
												m.G0 = v8 + int32(48)
												return v163
											} else {
												if v101&int32(1) != 0 {
													v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
													if v125 != 0 {
														v126 = int32(1)
													} else {
														v126 = int32(-1)
													}
													v163 = v126
													m.G0 = v8 + int32(48)
													return v163
												} else {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
													v132 = F_FunctionCall2Coll(m, l2+int32(212), v129, v130, v131)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														if v132 != 0 {
															v163 = v132
														} else {
															v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
															v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
															if v135 == int32(0) {
																v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																if v134&int32(1) == int32(0) {
																	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																	if v143 == v138 {
																		v163 = int32(0)
																	} else {
																		v146 = int32(1)
																		if v138&v146 != 0 {
																			v150 = v146
																		} else {
																			v150 = int32(-1)
																		}
																		v163 = v150
																	}
																} else {
																	v151 = int32(1)
																	if v138&v151 != 0 {
																		v155 = v151
																	} else {
																		v155 = int32(-1)
																	}
																	v163 = v155
																}
															} else {
																if v134&int32(1) != 0 {
																	v163 = int32(0)
																} else {
																	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																	if v161 != 0 {
																		v162 = int32(-1)
																	} else {
																		v162 = int32(1)
																	}
																	v163 = v162
																}
															}
														}
														m.G0 = v8 + int32(48)
														return v163
													}
												}
											}
										} else {
											v83 = int32(1)
											if v76&v83 != 0 {
												v87 = v83
											} else {
												v87 = int32(-1)
											}
											v163 = v87
											m.G0 = v8 + int32(48)
											return v163
										}
									} else {
										v88 = int32(1)
										if v76&v88 != 0 {
											v92 = v88
										} else {
											v92 = int32(-1)
										}
										v163 = v92
										m.G0 = v8 + int32(48)
										return v163
									}
								} else {
									if v72&int32(1) != 0 {
										v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
										v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
										if v102 == int32(1) {
											v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
											if v101&int32(1) != 0 {
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
												if v108 == v105 {
													v163 = int32(0)
												} else {
													v112 = int32(1)
													if v105&v112 != 0 {
														v115 = int32(-1)
													} else {
														v115 = v112
													}
													v163 = v115
												}
											} else {
												v117 = int32(1)
												if v105&v117 != 0 {
													v120 = int32(-1)
												} else {
													v120 = v117
												}
												v163 = v120
											}
											m.G0 = v8 + int32(48)
											return v163
										} else {
											if v101&int32(1) != 0 {
												v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
												if v125 != 0 {
													v126 = int32(1)
												} else {
													v126 = int32(-1)
												}
												v163 = v126
												m.G0 = v8 + int32(48)
												return v163
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
												v132 = F_FunctionCall2Coll(m, l2+int32(212), v129, v130, v131)
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int32(0)
												} else {
													if v132 != 0 {
														v163 = v132
													} else {
														v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
														v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
														if v135 == int32(0) {
															v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
															if v134&int32(1) == int32(0) {
																v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																if v143 == v138 {
																	v163 = int32(0)
																} else {
																	v146 = int32(1)
																	if v138&v146 != 0 {
																		v150 = v146
																	} else {
																		v150 = int32(-1)
																	}
																	v163 = v150
																}
															} else {
																v151 = int32(1)
																if v138&v151 != 0 {
																	v155 = v151
																} else {
																	v155 = int32(-1)
																}
																v163 = v155
															}
														} else {
															if v134&int32(1) != 0 {
																v163 = int32(0)
															} else {
																v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																if v161 != 0 {
																	v162 = int32(-1)
																} else {
																	v162 = int32(1)
																}
																v163 = v162
															}
														}
													}
													m.G0 = v8 + int32(48)
													return v163
												}
											}
										}
									} else {
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
										if v97 != 0 {
											v98 = int32(-1)
										} else {
											v98 = int32(1)
										}
										v163 = v98
										m.G0 = v8 + int32(48)
										return v163
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
func F_range_constructor3(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = F_get_fn_expr_rettype(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
		if v21 != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v22 == v16 {
				v32 = v21
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v33 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_range_constructor3_0), int32(0))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_constructor3_1), int32(424), int32(_a_F_range_constructor3_2))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
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
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v37 = F_pg_detoast_datum_packed(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = F_text_to_cstring(m, v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							if v41 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16801924))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
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
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
								if v44 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
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
									v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
									if v47 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
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
										if v41 != int32(40) {
											if v41 != int32(91) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16801924))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int32(0)
														} else {
															F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2334), int32(_a_F_range_constructor3_5))
																mBase = m.M
																v165 = m.ExcPending
																if v165 != 0 {
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
												v54 = int32(2)
												if v44 != int32(41) {
													if v44 != int32(93) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2348), int32(_a_F_range_constructor3_5))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
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
														v61 = v54 | int32(4)
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
														}
													}
												} else {
													v61 = v54
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
													if v62 != 0 {
														v65 = int32(0)
													} else {
														v65 = v14
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
													if v74 != 0 {
														v82 = v75
													} else {
														v82 = v13
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v90
													}
												}
											}
										} else {
											v54 = int32(0)
											if v44 != int32(41) {
												if v44 != int32(93) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16801924))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2348), int32(_a_F_range_constructor3_5))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
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
													v61 = v54 | int32(4)
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
													if v62 != 0 {
														v65 = int32(0)
													} else {
														v65 = v14
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
													if v74 != 0 {
														v82 = v75
													} else {
														v82 = v13
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v90
													}
												}
											} else {
												v61 = v54
												v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
												if v62 != 0 {
													v65 = int32(0)
												} else {
													v65 = v14
												}
												*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
												v67 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
												v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
												v75 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
												if v74 != 0 {
													v82 = v75
												} else {
													v82 = v13
												}
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
												v88 = int32(0)
												v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													m.G0 = v11 + int32(32)
													return v90
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
				v25 = F_lookup_type_cache(m, v16, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
							F_errmsg_internal(m, int32(_a_F_range_constructor3_6), v11)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_constructor3_1), int32(1776), int32(_a_F_range_constructor3_7))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
						if v33 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(130))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_range_constructor3_0), int32(0))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_range_constructor3_1), int32(424), int32(_a_F_range_constructor3_2))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
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
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v37 = F_pg_detoast_datum_packed(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = F_text_to_cstring(m, v37)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
									if v41 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
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
										v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
										if v44 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16801924))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
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
											v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
											if v47 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16801924))
													mBase = m.M
													v132 = m.ExcPending
													if v132 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
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
												if v41 != int32(40) {
													if v41 != int32(91) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																mBase = m.M
																v156 = m.ExcPending
																if v156 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2334), int32(_a_F_range_constructor3_5))
																		mBase = m.M
																		v165 = m.ExcPending
																		if v165 != 0 {
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
														v54 = int32(2)
														if v44 != int32(41) {
															if v44 != int32(93) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(16801924))
																	mBase = m.M
																	v172 = m.ExcPending
																	if v172 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																			mBase = m.M
																			v180 = m.ExcPending
																			if v180 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2348), int32(_a_F_range_constructor3_5))
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
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
																v61 = v54 | int32(4)
																v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
																if v62 != 0 {
																	v65 = int32(0)
																} else {
																	v65 = v14
																}
																*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
																v67 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
																v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
																v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
																v75 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
																if v74 != 0 {
																	v82 = v75
																} else {
																	v82 = v13
																}
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
																v88 = int32(0)
																v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(32)
																	return v90
																}
															}
														} else {
															v61 = v54
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
															if v62 != 0 {
																v65 = int32(0)
															} else {
																v65 = v14
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
															if v74 != 0 {
																v82 = v75
															} else {
																v82 = v13
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(32)
																return v90
															}
														}
													}
												} else {
													v54 = int32(0)
													if v44 != int32(41) {
														if v44 != int32(93) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16801924))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2348), int32(_a_F_range_constructor3_5))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
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
															v61 = v54 | int32(4)
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
															if v62 != 0 {
																v65 = int32(0)
															} else {
																v65 = v14
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
															if v74 != 0 {
																v82 = v75
															} else {
																v82 = v13
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(32)
																return v90
															}
														}
													} else {
														v61 = v54
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
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
				}
			}
		} else {
			v25 = F_lookup_type_cache(m, v16, int32(2048))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
				if v27 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
						F_errmsg_internal(m, int32(_a_F_range_constructor3_6), v11)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_range_constructor3_1), int32(1776), int32(_a_F_range_constructor3_7))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
					v32 = v25
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
					if v33 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(130))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_range_constructor3_0), int32(0))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_constructor3_1), int32(424), int32(_a_F_range_constructor3_2))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
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
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v37 = F_pg_detoast_datum_packed(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = F_text_to_cstring(m, v37)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
								if v41 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
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
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
									if v44 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
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
										v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
										if v47 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16801924))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2321), int32(_a_F_range_constructor3_5))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
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
											if v41 != int32(40) {
												if v41 != int32(91) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16801924))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2334), int32(_a_F_range_constructor3_5))
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
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
													v54 = int32(2)
													if v44 != int32(41) {
														if v44 != int32(93) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16801924))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2348), int32(_a_F_range_constructor3_5))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
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
															v61 = v54 | int32(4)
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
															if v62 != 0 {
																v65 = int32(0)
															} else {
																v65 = v14
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
															if v74 != 0 {
																v82 = v75
															} else {
																v82 = v13
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(32)
																return v90
															}
														}
													} else {
														v61 = v54
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
														}
													}
												}
											} else {
												v54 = int32(0)
												if v44 != int32(41) {
													if v44 != int32(93) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2348), int32(_a_F_range_constructor3_5))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
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
														v61 = v54 | int32(4)
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
														}
													}
												} else {
													v61 = v54
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
													if v62 != 0 {
														v65 = int32(0)
													} else {
														v65 = v14
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
													if v74 != 0 {
														v82 = v75
													} else {
														v82 = v13
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v90
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
			}
		}
	}
}
func F_range_contains_elem(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v16 {
				v30 = v19
				v31 = F_range_contains_elem_internal(m, v30, v12, v17)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v31
				}
			} else {
				v23 = F_lookup_type_cache(m, v16, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
							F_errmsg_internal(m, int32(_a_F_range_contains_elem_0), v9)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contains_elem_1), int32(1776), int32(_a_F_range_contains_elem_2))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						v31 = F_range_contains_elem_internal(m, v30, v12, v17)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v31
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v16, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
						F_errmsg_internal(m, int32(_a_F_range_contains_elem_0), v9)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_range_contains_elem_1), int32(1776), int32(_a_F_range_contains_elem_2))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					v31 = F_range_contains_elem_internal(m, v30, v12, v17)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v31
					}
				}
			}
		}
	}
}
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = l1 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+10)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)))
	v23 = base.I32_extend16_s(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v24)>>(uint(int32(2))%32))-int32(1)))))
	if v30&int32(41) != 0 {
		v89 = int32(0)
		v90 = v18
		if v30&int32(81) != 0 {
			v153 = v89
			v156 = v6
			v159 = v153
			v161 = v156
			v162 = int32(1)
			v163 = v30 & v162
			*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
			v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
			v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
			v178 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
			v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
			v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
			m.G0 = v15 + int32(32)
			return
		} else {
			if v23 == int32(-1) {
				v103 = v89
				v104 = v90
				v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
				if v106 != 0 {
					v126 = v103
					v127 = v104
					v128 = int32(-1)
				} else {
					v108 = v103
					v109 = v104
					v110 = int32(-1)
					switch v20 - int32(99) {
					case 0:
						v126 = v108
						v127 = v109
						v128 = v110
					case 1:
						v126 = v108
						v127 = (v109 + int32(7)) & int32(-8)
						v128 = v110
					default:
						v126 = v108
						v127 = (v109 + int32(1)) & int32(-2)
						v128 = v110
					case 6:
						v126 = v108
						v127 = (v109 + int32(3)) & int32(-4)
						v128 = v110
					}
				}
			} else {
				v108 = v89
				v109 = v90
				v110 = v23
				switch v20 - int32(99) {
				case 0:
					v126 = v108
					v127 = v109
					v128 = v110
				case 1:
					v126 = v108
					v127 = (v109 + int32(7)) & int32(-8)
					v128 = v110
				default:
					v126 = v108
					v127 = (v109 + int32(1)) & int32(-2)
					v128 = v110
				case 6:
					v126 = v108
					v127 = (v109 + int32(3)) & int32(-4)
					v128 = v110
				}
			}
			if v21&int32(1) == int32(0) {
				v159 = v126
				v161 = v127
				v162 = int32(1)
				v163 = v30 & v162
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
				v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
				v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
				v178 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
				v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
				v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
				m.G0 = v15 + int32(32)
				return
			} else {
				switch v128 - int32(1) {
				case 0:
					v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
					v153 = v126
					v156 = v152
					v159 = v153
					v161 = v156
					v162 = int32(1)
					v163 = v30 & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
					v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
					v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
					v178 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
					v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
					v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
					m.G0 = v15 + int32(32)
					return
				case 1:
					v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
					v153 = v126
					v156 = v135
					v159 = v153
					v161 = v156
					v162 = int32(1)
					v163 = v30 & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
					v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
					v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
					v178 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
					v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
					v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
					m.G0 = v15 + int32(32)
					return
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v128
						F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(16))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
					v153 = v126
					v156 = v136
					v159 = v153
					v161 = v156
					v162 = int32(1)
					v163 = v30 & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
					v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
					v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
					v178 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
					v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
					v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
					m.G0 = v15 + int32(32)
					return
				}
			}
		}
	} else {
		if v21&int32(1) != 0 {
			switch v22 - int32(1) {
			case 0:
				v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18))))
				v89 = v38
				v90 = v18 + v23
				if v30&int32(81) != 0 {
					v153 = v89
					v156 = v6
					v159 = v153
					v161 = v156
					v162 = int32(1)
					v163 = v30 & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
					v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
					v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
					v178 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
					v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
					v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v103 = v89
						v104 = v90
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
						if v106 != 0 {
							v126 = v103
							v127 = v104
							v128 = int32(-1)
						} else {
							v108 = v103
							v109 = v104
							v110 = int32(-1)
							switch v20 - int32(99) {
							case 0:
								v126 = v108
								v127 = v109
								v128 = v110
							case 1:
								v126 = v108
								v127 = (v109 + int32(7)) & int32(-8)
								v128 = v110
							default:
								v126 = v108
								v127 = (v109 + int32(1)) & int32(-2)
								v128 = v110
							case 6:
								v126 = v108
								v127 = (v109 + int32(3)) & int32(-4)
								v128 = v110
							}
						}
					} else {
						v108 = v89
						v109 = v90
						v110 = v23
						switch v20 - int32(99) {
						case 0:
							v126 = v108
							v127 = v109
							v128 = v110
						case 1:
							v126 = v108
							v127 = (v109 + int32(7)) & int32(-8)
							v128 = v110
						default:
							v126 = v108
							v127 = (v109 + int32(1)) & int32(-2)
							v128 = v110
						case 6:
							v126 = v108
							v127 = (v109 + int32(3)) & int32(-4)
							v128 = v110
						}
					}
					if v21&int32(1) == int32(0) {
						v159 = v126
						v161 = v127
						v162 = int32(1)
						v163 = v30 & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
						v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
						v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
						v178 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
						v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
						v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v128 - int32(1) {
						case 0:
							v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v152
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						case 1:
							v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v135
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v128
								F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(16))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
							v153 = v126
							v156 = v136
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			case 1:
				v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18))))
				v89 = v40
				v90 = v18 + v23
				if v30&int32(81) != 0 {
					v153 = v89
					v156 = v6
					v159 = v153
					v161 = v156
					v162 = int32(1)
					v163 = v30 & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
					v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
					v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
					v178 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
					v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
					v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v103 = v89
						v104 = v90
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
						if v106 != 0 {
							v126 = v103
							v127 = v104
							v128 = int32(-1)
						} else {
							v108 = v103
							v109 = v104
							v110 = int32(-1)
							switch v20 - int32(99) {
							case 0:
								v126 = v108
								v127 = v109
								v128 = v110
							case 1:
								v126 = v108
								v127 = (v109 + int32(7)) & int32(-8)
								v128 = v110
							default:
								v126 = v108
								v127 = (v109 + int32(1)) & int32(-2)
								v128 = v110
							case 6:
								v126 = v108
								v127 = (v109 + int32(3)) & int32(-4)
								v128 = v110
							}
						}
					} else {
						v108 = v89
						v109 = v90
						v110 = v23
						switch v20 - int32(99) {
						case 0:
							v126 = v108
							v127 = v109
							v128 = v110
						case 1:
							v126 = v108
							v127 = (v109 + int32(7)) & int32(-8)
							v128 = v110
						default:
							v126 = v108
							v127 = (v109 + int32(1)) & int32(-2)
							v128 = v110
						case 6:
							v126 = v108
							v127 = (v109 + int32(3)) & int32(-4)
							v128 = v110
						}
					}
					if v21&int32(1) == int32(0) {
						v159 = v126
						v161 = v127
						v162 = int32(1)
						v163 = v30 & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
						v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
						v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
						v178 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
						v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
						v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v128 - int32(1) {
						case 0:
							v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v152
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						case 1:
							v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v135
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v128
								F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(16))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
							v153 = v126
							v156 = v136
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v23
					F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 3:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v89 = v42
				v90 = v18 + v23
				if v30&int32(81) != 0 {
					v153 = v89
					v156 = v6
					v159 = v153
					v161 = v156
					v162 = int32(1)
					v163 = v30 & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
					v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
					v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
					v178 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
					v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
					v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v103 = v89
						v104 = v90
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
						if v106 != 0 {
							v126 = v103
							v127 = v104
							v128 = int32(-1)
						} else {
							v108 = v103
							v109 = v104
							v110 = int32(-1)
							switch v20 - int32(99) {
							case 0:
								v126 = v108
								v127 = v109
								v128 = v110
							case 1:
								v126 = v108
								v127 = (v109 + int32(7)) & int32(-8)
								v128 = v110
							default:
								v126 = v108
								v127 = (v109 + int32(1)) & int32(-2)
								v128 = v110
							case 6:
								v126 = v108
								v127 = (v109 + int32(3)) & int32(-4)
								v128 = v110
							}
						}
					} else {
						v108 = v89
						v109 = v90
						v110 = v23
						switch v20 - int32(99) {
						case 0:
							v126 = v108
							v127 = v109
							v128 = v110
						case 1:
							v126 = v108
							v127 = (v109 + int32(7)) & int32(-8)
							v128 = v110
						default:
							v126 = v108
							v127 = (v109 + int32(1)) & int32(-2)
							v128 = v110
						case 6:
							v126 = v108
							v127 = (v109 + int32(3)) & int32(-4)
							v128 = v110
						}
					}
					if v21&int32(1) == int32(0) {
						v159 = v126
						v161 = v127
						v162 = int32(1)
						v163 = v30 & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
						v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
						v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
						v178 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
						v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
						v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v128 - int32(1) {
						case 0:
							v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v152
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						case 1:
							v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v135
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v128
								F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(16))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
							v153 = v126
							v156 = v136
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			}
		} else {
			if int32(0) < v23 {
				v89 = v18
				v90 = v18 + v23
				if v30&int32(81) != 0 {
					v153 = v89
					v156 = v6
					v159 = v153
					v161 = v156
					v162 = int32(1)
					v163 = v30 & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
					v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
					v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
					v178 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
					v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
					v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v103 = v89
						v104 = v90
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
						if v106 != 0 {
							v126 = v103
							v127 = v104
							v128 = int32(-1)
						} else {
							v108 = v103
							v109 = v104
							v110 = int32(-1)
							switch v20 - int32(99) {
							case 0:
								v126 = v108
								v127 = v109
								v128 = v110
							case 1:
								v126 = v108
								v127 = (v109 + int32(7)) & int32(-8)
								v128 = v110
							default:
								v126 = v108
								v127 = (v109 + int32(1)) & int32(-2)
								v128 = v110
							case 6:
								v126 = v108
								v127 = (v109 + int32(3)) & int32(-4)
								v128 = v110
							}
						}
					} else {
						v108 = v89
						v109 = v90
						v110 = v23
						switch v20 - int32(99) {
						case 0:
							v126 = v108
							v127 = v109
							v128 = v110
						case 1:
							v126 = v108
							v127 = (v109 + int32(7)) & int32(-8)
							v128 = v110
						default:
							v126 = v108
							v127 = (v109 + int32(1)) & int32(-2)
							v128 = v110
						case 6:
							v126 = v108
							v127 = (v109 + int32(3)) & int32(-4)
							v128 = v110
						}
					}
					if v21&int32(1) == int32(0) {
						v159 = v126
						v161 = v127
						v162 = int32(1)
						v163 = v30 & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
						v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
						v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
						v178 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
						v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
						v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v128 - int32(1) {
						case 0:
							v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v152
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						case 1:
							v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
							v153 = v126
							v156 = v135
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v128
								F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(16))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
							v153 = v126
							v156 = v136
							v159 = v153
							v161 = v156
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			} else {
				if v23 == int32(-1) {
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					if v62 == int32(1) {
						v66 = int32(18)
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
						if v68 == v66 {
							v71 = v66
						} else {
							v71 = int32(2)
						}
						if base.Ui32((v68-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v78 = int32(6)
						} else {
							v78 = v71
						}
						v99 = v78
					} else {
						if v62&int32(1) == int32(0) {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
							v99 = int32(base.Ui32(v95) >> (uint(int32(2)) % 32))
						} else {
							v99 = int32(base.Ui32(v62) >> (uint(int32(1)) % 32))
						}
					}
					if v30&int32(80) != 0 {
						v159 = v18
						v161 = v6
						v162 = int32(1)
						v163 = v30 & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
						v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
						v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
						v178 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
						v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
						v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
						m.G0 = v15 + int32(32)
						return
					} else {
						v103 = v18
						v104 = v99 + v18
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
						if v106 != 0 {
							v126 = v103
							v127 = v104
							v128 = int32(-1)
						} else {
							v108 = v103
							v109 = v104
							v110 = int32(-1)
							switch v20 - int32(99) {
							case 0:
								v126 = v108
								v127 = v109
								v128 = v110
							case 1:
								v126 = v108
								v127 = (v109 + int32(7)) & int32(-8)
								v128 = v110
							default:
								v126 = v108
								v127 = (v109 + int32(1)) & int32(-2)
								v128 = v110
							case 6:
								v126 = v108
								v127 = (v109 + int32(3)) & int32(-4)
								v128 = v110
							}
						}
						if v21&int32(1) == int32(0) {
							v159 = v126
							v161 = v127
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						} else {
							switch v128 - int32(1) {
							case 0:
								v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
								v153 = v126
								v156 = v152
								v159 = v153
								v161 = v156
								v162 = int32(1)
								v163 = v30 & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
								v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
								v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
								v178 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
								v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
								v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
								m.G0 = v15 + int32(32)
								return
							case 1:
								v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
								v153 = v126
								v156 = v135
								v159 = v153
								v161 = v156
								v162 = int32(1)
								v163 = v30 & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
								v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
								v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
								v178 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
								v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
								v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
								m.G0 = v15 + int32(32)
								return
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v128
									F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(16))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 3:
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
								v153 = v126
								v156 = v136
								v159 = v153
								v161 = v156
								v162 = int32(1)
								v163 = v30 & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
								v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
								v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
								v178 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
								v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
								v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
								m.G0 = v15 + int32(32)
								return
							}
						}
					}
				} else {
					v85 = F_strlen(m, v18)
					mBase = m.M
					v89 = v18
					v90 = v85 + v18 + int32(1)
					if v30&int32(81) != 0 {
						v153 = v89
						v156 = v6
						v159 = v153
						v161 = v156
						v162 = int32(1)
						v163 = v30 & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
						v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
						v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
						v178 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
						v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
						v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
						m.G0 = v15 + int32(32)
						return
					} else {
						if v23 == int32(-1) {
							v103 = v89
							v104 = v90
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
							if v106 != 0 {
								v126 = v103
								v127 = v104
								v128 = int32(-1)
							} else {
								v108 = v103
								v109 = v104
								v110 = int32(-1)
								switch v20 - int32(99) {
								case 0:
									v126 = v108
									v127 = v109
									v128 = v110
								case 1:
									v126 = v108
									v127 = (v109 + int32(7)) & int32(-8)
									v128 = v110
								default:
									v126 = v108
									v127 = (v109 + int32(1)) & int32(-2)
									v128 = v110
								case 6:
									v126 = v108
									v127 = (v109 + int32(3)) & int32(-4)
									v128 = v110
								}
							}
						} else {
							v108 = v89
							v109 = v90
							v110 = v23
							switch v20 - int32(99) {
							case 0:
								v126 = v108
								v127 = v109
								v128 = v110
							case 1:
								v126 = v108
								v127 = (v109 + int32(7)) & int32(-8)
								v128 = v110
							default:
								v126 = v108
								v127 = (v109 + int32(1)) & int32(-2)
								v128 = v110
							case 6:
								v126 = v108
								v127 = (v109 + int32(3)) & int32(-4)
								v128 = v110
							}
						}
						if v21&int32(1) == int32(0) {
							v159 = v126
							v161 = v127
							v162 = int32(1)
							v163 = v30 & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
							v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
							v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
							v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
							v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
							m.G0 = v15 + int32(32)
							return
						} else {
							switch v128 - int32(1) {
							case 0:
								v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
								v153 = v126
								v156 = v152
								v159 = v153
								v161 = v156
								v162 = int32(1)
								v163 = v30 & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
								v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
								v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
								v178 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
								v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
								v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
								m.G0 = v15 + int32(32)
								return
							case 1:
								v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127))))
								v153 = v126
								v156 = v135
								v159 = v153
								v161 = v156
								v162 = int32(1)
								v163 = v30 & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
								v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
								v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
								v178 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
								v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
								v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
								m.G0 = v15 + int32(32)
								return
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v128
									F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(16))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_range_deserialize_1), int32(70), int32(_a_F_range_deserialize_2))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 3:
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
								v153 = v126
								v156 = v136
								v159 = v153
								v161 = v156
								v162 = int32(1)
								v163 = v30 & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v163)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v162)
								v170 = int32(base.Ui32(v30)>>(uint(v162)%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v170)
								v175 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v175)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
								v178 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v178)
								v183 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v183)
								v188 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v162
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v188)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v161
								m.G0 = v15 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_range_empty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v11 = int32(1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(base.Ui32(v7)>>(uint(int32(2))%32))-v11))))
		return v13 & v11
	}
}
func F_range_fast_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = F_pg_detoast_datum(m, l0)
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
	v18 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v25 = F_lookup_type_cache(m, v23, int32(2048))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v28 = v20
	goto L6
L6:
	;
	F_range_deserialize(m, v28, v14, v12+int32(40), v12+int32(24), v12+int32(15))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v25
	v28 = v25
	goto L6
L8:
	;
	F_range_deserialize(m, v28, v18, v12+int32(32), v12+int32(16), v12+int32(14))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
	v48 = int32(1)
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v51 = v45&v46 - v48
	goto L12
L11:
	;
	v51 = v48
	goto L12
L12:
	;
	if v45|v46&int32(1) != 0 {
		v178 = v51
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l0 != v14 {
		goto L98
	} else {
		goto L99
	}
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+36)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
	if v56 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
	if v117 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L16:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+46)))
	if v55&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v55&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v62 == v59 {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v70 = int32(1)
	if v59&v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v65 = int32(1)
	if v59&v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v68 = int32(-1)
	goto L25
L24:
	;
	v68 = v65
	goto L25
L25:
	;
	v178 = v68
	goto L13
L26:
	;
	v73 = int32(-1)
	goto L28
L27:
	;
	v73 = v70
	goto L28
L28:
	;
	v178 = v73
	goto L13
L29:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v78 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v85 = F_FunctionCall2Coll(m, v28+int32(212), v82, v83, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	v79 = int32(1)
	goto L34
L33:
	;
	v79 = int32(-1)
	goto L34
L34:
	;
	v178 = v79
	goto L13
L35:
	;
	if v85 != 0 {
		v178 = v85
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+37)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+45)))
	if v88 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+46)))
	if v87&int32(1) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v87&int32(1) != 0 {
		goto L15
	} else {
		goto L50
	}
L40:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v96 == v91 {
		goto L15
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v103 = int32(1)
	if v91&v103 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v98 = int32(1)
	if v91&v98 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v102 = v98
	goto L46
L45:
	;
	v102 = int32(-1)
	goto L46
L46:
	;
	v178 = v102
	goto L13
L47:
	;
	v107 = v103
	goto L49
L48:
	;
	v107 = int32(-1)
	goto L49
L49:
	;
	v178 = v107
	goto L13
L50:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v112 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v113 = int32(-1)
	goto L53
L52:
	;
	v113 = int32(1)
	goto L53
L53:
	;
	v178 = v113
	goto L13
L54:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
	if v116&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if v116&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v123 == v120 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v132 = int32(1)
	if v120&v132 != 0 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v178 = int32(0)
	goto L13
L61:
	;
	goto L62
L62:
	;
	v127 = int32(1)
	if v120&v127 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v130 = int32(-1)
	goto L65
L64:
	;
	v130 = v127
	goto L65
L65:
	;
	v178 = v130
	goto L13
L66:
	;
	v135 = int32(-1)
	goto L68
L67:
	;
	v135 = v132
	goto L68
L68:
	;
	v178 = v135
	goto L13
L69:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v140 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v147 = F_FunctionCall2Coll(m, v28+int32(212), v144, v145, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L75
	}
L72:
	;
	v141 = int32(1)
	goto L74
L73:
	;
	v141 = int32(-1)
	goto L74
L74:
	;
	v178 = v141
	goto L13
L75:
	;
	if v147 != 0 {
		v178 = v147
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)))
	if v150 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
	if v149&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if v149&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v158 == v153 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v166 = int32(1)
	if v153&v166 != 0 {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v178 = int32(0)
	goto L13
L84:
	;
	goto L85
L85:
	;
	v161 = int32(1)
	if v153&v161 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v165 = v161
	goto L88
L87:
	;
	v165 = int32(-1)
	goto L88
L88:
	;
	v178 = v165
	goto L13
L89:
	;
	v170 = v166
	goto L91
L90:
	;
	v170 = int32(-1)
	goto L91
L91:
	;
	v178 = v170
	goto L13
L92:
	;
	v178 = int32(0)
	goto L13
L93:
	;
	goto L94
L94:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v176 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v177 = int32(-1)
	goto L97
L96:
	;
	v177 = int32(1)
	goto L97
L97:
	;
	v178 = v177
	goto L13
L98:
	;
	F_pfree(m, v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if l1 != v18 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	F_pfree(m, v18)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	m.G0 = v12 + int32(48)
	return v178
L105:
	;
	goto L104
}
func F_range_gist_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v78 float32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 float32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 float32
	_ = v97
	var v98 int32
	_ = v98
	var v99 float32
	_ = v99
	var v100 float32
	_ = v100
	var v101 float32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v114 float32
	_ = v114
	var v117 int32
	_ = v117
	var v118 float32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 float32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v157 float32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 float32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v192 float64
	_ = v192
	var v194 float32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 float64
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v230 float64
	_ = v230
	var v233 float64
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v250 float64
	_ = v250
	var v253 float64
	_ = v253
	var v257 float64
	_ = v257
	var v264 float32
	_ = v264
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			if v25 == v26 {
				v28 = F_range_get_typcache(m, l0, v25)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+272))
					F_range_deserialize(m, v28, v18, v12+int32(40), v12+int32(24), v12+int32(15))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_range_deserialize(m, v28, v23, v12+int32(32), v12+int32(16), v12+int32(14))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
							if v47 == int32(1) {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
								if v50 != 0 {
									v264 = float32(0)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
									v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v51)>>(uint(int32(2))%32))-int32(1)))))
									if v57&int32(-127) != 0 {
										v264 = float32(1)
									} else {
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
										v62 = int32(1)
										v64 = int32(0)
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
										if base.B2i32(v61&v62 == v64)|base.B2i32(v66 != v62) == v64 {
											v264 = float32(2)
										} else {
											if (v61|v66)&int32(1) != 0 {
												v78 = float32(3)
											} else {
												v78 = float32(4)
											}
											v264 = v78
										}
									}
								}
								*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
								m.G0 = v12 + int32(48)
								return v14
							} else {
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
								v80 = int32(1)
								v82 = int32(0)
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+36)))
								if base.B2i32(v79&v80 == v82)|base.B2i32(v84 != v80) == v82 {
									v91 = float32(2)
									v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
									v96 = v94 & int32(1)
									if v96 != 0 {
										v97 = v91
									} else {
										v97 = float32(4)
									}
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
									if v98 != 0 {
										v99 = v91
									} else {
										v99 = v97
									}
									if v96 != 0 {
										v100 = float32(0)
									} else {
										v100 = v99
									}
									if v98 != 0 {
										v101 = v100
									} else {
										v101 = v99
									}
									*(*float32)(unsafe.Add(mBase, uint32(v14))) = v101
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
									v109 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v103)>>(uint(int32(2))%32))-int32(1)))))
									if v109&int32(-127) == int32(0) {
									} else {
										v114 = *(*float32)(unsafe.Add(mBase, uint32(v14)))
										v264 = base.F32_add(v114, float32(1))
										*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
									}
									m.G0 = v12 + int32(48)
									return v14
								} else {
									v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
									if v84 != 0 {
										v118 = math.Float32frombits(uint32(0x7f800000))
										if v117&int32(1) != 0 {
											v264 = v118
											*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
											m.G0 = v12 + int32(48)
											return v14
										} else {
											v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
											if v121&int32(1) == int32(0) {
												v264 = v118
												*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
												m.G0 = v12 + int32(48)
												return v14
											} else {
												v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
												if v126 != 0 {
													v264 = float32(0)
													*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
													m.G0 = v12 + int32(48)
													return v14
												} else {
													v134 = F_range_cmp_bounds(m, v28, v12+int32(16), v12+int32(24))
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int32(0)
													} else {
														v137 = base.B2i32(v134 <= int32(0))
														if v134 <= int32(0) {
															v138 = float32(0)
														} else {
															v138 = float32(1)
														}
														if v137|base.B2i32(v30 == int32(0)) != 0 {
															v264 = v138
															*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
															m.G0 = v12 + int32(48)
															return v14
														} else {
															v144 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
															v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
															v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
															v147 = F_FunctionCall2Coll(m, v28+int32(268), v144, v145, v146)
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return int32(0)
															} else {
																v149 = *(*float64)(unsafe.Add(mBase, uint32(v147)))
																v150 = float64(0)
																if base.F64_ge(v149, v150) != 0 {
																	v153 = v149
																} else {
																	v153 = v150
																}
																v264 = base.F32_demote_f64(v153)
																*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																m.G0 = v12 + int32(48)
																return v14
															}
														}
													}
												}
											}
										}
									} else {
										if v79&int32(1) != 0 {
											v157 = math.Float32frombits(uint32(0x7f800000))
											if v117&int32(1) != 0 {
												v264 = v157
												*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
												m.G0 = v12 + int32(48)
												return v14
											} else {
												v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
												if v160&int32(1) == int32(0) {
													v264 = v157
													*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
													m.G0 = v12 + int32(48)
													return v14
												} else {
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
													if v165 != 0 {
														v264 = float32(0)
														*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
														m.G0 = v12 + int32(48)
														return v14
													} else {
														v173 = F_range_cmp_bounds(m, v28, v12+int32(32), v12+int32(40))
														mBase = m.M
														v174 = m.ExcPending
														if v174 != 0 {
															return int32(0)
														} else {
															v176 = base.B2i32(int32(0) <= v173)
															if int32(0) <= v173 {
																v177 = float32(0)
															} else {
																v177 = float32(1)
															}
															if v176|base.B2i32(v30 == int32(0)) != 0 {
																v264 = v177
																*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																m.G0 = v12 + int32(48)
																return v14
															} else {
																v183 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
																v184 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
																v185 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
																v186 = F_FunctionCall2Coll(m, v28+int32(268), v183, v184, v185)
																mBase = m.M
																v187 = m.ExcPending
																if v187 != 0 {
																	return int32(0)
																} else {
																	v188 = *(*float64)(unsafe.Add(mBase, uint32(v186)))
																	v189 = float64(0)
																	if base.F64_ge(v188, v189) != 0 {
																		v192 = v188
																	} else {
																		v192 = v189
																	}
																	v264 = base.F32_demote_f64(v192)
																	*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																	m.G0 = v12 + int32(48)
																	return v14
																}
															}
														}
													}
												}
											}
										} else {
											v194 = math.Float32frombits(uint32(0x7f800000))
											if v117&int32(1) != 0 {
												v264 = v194
												*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
												m.G0 = v12 + int32(48)
												return v14
											} else {
												v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
												if v197&int32(1) != 0 {
													v264 = v194
													*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
													m.G0 = v12 + int32(48)
													return v14
												} else {
													v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
													if v200&int32(1) != 0 {
														v264 = v194
														*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
														m.G0 = v12 + int32(48)
														return v14
													} else {
														v209 = F_range_cmp_bounds(m, v28, v12+int32(32), v12+int32(40))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int32(0)
														} else {
															v212 = base.B2i32(int32(0) <= v209)
															if int32(0) <= v209 {
																v213 = float64(0)
															} else {
																v213 = float64(1)
															}
															v214 = int32(0)
															if v212|base.B2i32(v30 == v214) == v214 {
																v221 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
																v222 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
																v223 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
																v224 = F_FunctionCall2Coll(m, v28+int32(268), v221, v222, v223)
																mBase = m.M
																v225 = m.ExcPending
																if v225 != 0 {
																	return int32(0)
																} else {
																	v226 = *(*float64)(unsafe.Add(mBase, uint32(v224)))
																	v227 = float64(0)
																	if base.F64_ge(v226, v227) != 0 {
																		v230 = v226
																	} else {
																		v230 = v227
																	}
																	v233 = base.F64_add(v230, float64(0))
																	v238 = F_range_cmp_bounds(m, v28, v12+int32(16), v12+int32(24))
																	mBase = m.M
																	v239 = m.ExcPending
																	if v239 != 0 {
																		return int32(0)
																	} else {
																		if v238 <= int32(0) {
																			v257 = v233
																			v264 = base.F32_demote_f64(v257)
																			*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																			m.G0 = v12 + int32(48)
																			return v14
																		} else {
																			if v30 != 0 {
																				v244 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
																				v245 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																				v246 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
																				v247 = F_FunctionCall2Coll(m, v28+int32(268), v244, v245, v246)
																				mBase = m.M
																				v248 = m.ExcPending
																				if v248 != 0 {
																					return int32(0)
																				} else {
																					v249 = *(*float64)(unsafe.Add(mBase, uint32(v247)))
																					v250 = float64(0)
																					if base.F64_ge(v249, v250) != 0 {
																						v253 = v249
																					} else {
																						v253 = v250
																					}
																					v257 = base.F64_add(v233, v253)
																					v264 = base.F32_demote_f64(v257)
																					*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																					m.G0 = v12 + int32(48)
																					return v14
																				}
																			} else {
																				v257 = base.F64_add(v233, float64(1))
																				v264 = base.F32_demote_f64(v257)
																				*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																				m.G0 = v12 + int32(48)
																				return v14
																			}
																		}
																	}
																}
															} else {
																v233 = v213
																v238 = F_range_cmp_bounds(m, v28, v12+int32(16), v12+int32(24))
																mBase = m.M
																v239 = m.ExcPending
																if v239 != 0 {
																	return int32(0)
																} else {
																	if v238 <= int32(0) {
																		v257 = v233
																		v264 = base.F32_demote_f64(v257)
																		*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																		m.G0 = v12 + int32(48)
																		return v14
																	} else {
																		if v30 != 0 {
																			v244 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
																			v245 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																			v246 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
																			v247 = F_FunctionCall2Coll(m, v28+int32(268), v244, v245, v246)
																			mBase = m.M
																			v248 = m.ExcPending
																			if v248 != 0 {
																				return int32(0)
																			} else {
																				v249 = *(*float64)(unsafe.Add(mBase, uint32(v247)))
																				v250 = float64(0)
																				if base.F64_ge(v249, v250) != 0 {
																					v253 = v249
																				} else {
																					v253 = v250
																				}
																				v257 = base.F64_add(v233, v253)
																				v264 = base.F32_demote_f64(v257)
																				*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																				m.G0 = v12 + int32(48)
																				return v14
																			}
																		} else {
																			v257 = base.F64_add(v233, float64(1))
																			v264 = base.F32_demote_f64(v257)
																			*(*float32)(unsafe.Add(mBase, uint32(v14))) = v264
																			m.G0 = v12 + int32(48)
																			return v14
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
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v281 = m.ExcPending
				if v281 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_range_gist_penalty_0), int32(0))
					mBase = m.M
					v285 = m.ExcPending
					if v285 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_range_gist_penalty_1), int32(379), int32(_a_F_range_gist_penalty_2))
						mBase = m.M
						v290 = m.ExcPending
						if v290 != 0 {
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
func F_range_gist_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v27 float32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 float32
	_ = v358
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v392 float32
	_ = v392
	var v394 float32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 float32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 float64
	_ = v538
	var v539 float64
	_ = v539
	var v542 float64
	_ = v542
	var v547 float32
	_ = v547
	var v552 int32
	_ = v552
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 float32
	_ = v566
	var v568 float32
	_ = v568
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v597 float32
	_ = v597
	var v599 float32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v636 float32
	_ = v636
	var v638 float32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v737 int32
	_ = v737
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 float32
	_ = v777
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 float64
	_ = v788
	var v789 float64
	_ = v789
	var v792 float64
	_ = v792
	var v797 float32
	_ = v797
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 float32
	_ = v816
	var v818 float32
	_ = v818
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 float64
	_ = v931
	var v932 float64
	_ = v932
	var v935 float64
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 float64
	_ = v941
	var v942 float64
	_ = v942
	var v945 float64
	_ = v945
	var v949 float64
	_ = v949
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1097 int32
	_ = v1097
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int64
	_ = v1131
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1172 int32
	_ = v1172
	var v1179 int64
	_ = v1179
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1201 int32
	_ = v1201
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1385 int32
	_ = v1385
	v2 = int32(0)
	v27 = float32(0)
	v32 = m.G0
	v34 = v32 - int32(96)
	m.G0 = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v39 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v44 = F_range_get_typcache(m, l0, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v46 = int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v51 = (v47 - v46) & int32(_a_F_range_gist_picksplit_0)
	v55 = v51<<(uint(v46)%32) + int32(2)
	v56 = F_palloc(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v56
	v59 = F_palloc(m, v55)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = int32(0)
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+72)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v34)+64)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v34)+56)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v34)+48)) = v64
	v73 = v37 + int32(4)
	if v47&int32(_a_F_range_gist_picksplit_0) != int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v78 = v46
	goto L9
L7:
	;
	v161 = v2
	v162 = v2
	v165 = v2
	v167 = v2
	v168 = v2
	v170 = v2
	v171 = v2
	v174 = v2
	v178 = v2
	goto L8
L8:
	;
	v199 = int32(0)
	v201 = base.B2i32(v161 <= v199)
	if v161 <= v199 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v73+v78<<(uint(int32(4))%32))))
	v115 = F_pg_detoast_datum(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v161 = v156
	v162 = v158
	v165 = v157
	v167 = v153
	v168 = v151
	v170 = v155
	v171 = v152
	v174 = v154
	v178 = v150
	goto L8
L11:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v123 = int32(*(*int8)(unsafe.Add(mBase, uint32(v115+int32(base.Ui32(v117)>>(uint(int32(2))%32))-int32(1)))))
	goto L12
L12:
	;
	if v123&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v137 = int32(8)
	goto L15
L14:
	;
	v127 = int32(3)
	v130 = int32(base.Ui32(v123)>>(uint(v127)%32)) & v127
	if v123 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v140 = v34 + int32(48) + v137<<(uint(int32(2))%32)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v142 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v141 + v142
	v148 = (v78 + v142) & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v148) <= base.Ui32(v51) {
		v78 = v148
		goto L9
	} else {
		goto L19
	}
L16:
	;
	v135 = v130 | int32(4)
	goto L18
L17:
	;
	v135 = v130
	goto L18
L18:
	;
	v137 = v135
	goto L15
L19:
	;
	goto L10
L20:
	;
	v202 = int32(-1)
	goto L22
L21:
	;
	v202 = v199
	goto L22
L22:
	;
	v203 = int32(0)
	v205 = base.B2i32(v203 < v161)
	if v203 < v161 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v206 = v161
	goto L25
L24:
	;
	v206 = v203
	goto L25
L25:
	;
	v207 = base.B2i32(v206 < v170)
	if v206 < v170 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v208 = int32(1)
	goto L28
L27:
	;
	v208 = v202
	goto L28
L28:
	;
	v210 = base.B2i32(int32(0) < v170)
	if int32(0) < v170 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v211 = v208
	goto L31
L30:
	;
	v211 = v202
	goto L31
L31:
	;
	if v206 < v170 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v212 = v170
	goto L34
L33:
	;
	v212 = v206
	goto L34
L34:
	;
	if int32(0) < v170 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v213 = v212
	goto L37
L36:
	;
	v213 = v206
	goto L37
L37:
	;
	v214 = base.B2i32(v213 < v174)
	if v213 < v174 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v215 = int32(2)
	goto L40
L39:
	;
	v215 = v211
	goto L40
L40:
	;
	v217 = base.B2i32(int32(0) < v174)
	if int32(0) < v174 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v218 = v215
	goto L43
L42:
	;
	v218 = v211
	goto L43
L43:
	;
	if v213 < v174 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v219 = v174
	goto L46
L45:
	;
	v219 = v213
	goto L46
L46:
	;
	if int32(0) < v174 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v220 = v219
	goto L49
L48:
	;
	v220 = v213
	goto L49
L49:
	;
	v221 = base.B2i32(v220 < v167)
	if v220 < v167 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v222 = int32(3)
	goto L52
L51:
	;
	v222 = v218
	goto L52
L52:
	;
	v224 = base.B2i32(int32(0) < v167)
	if int32(0) < v167 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v225 = v222
	goto L55
L54:
	;
	v225 = v218
	goto L55
L55:
	;
	if v220 < v167 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v226 = v167
	goto L58
L57:
	;
	v226 = v220
	goto L58
L58:
	;
	if int32(0) < v167 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v227 = v226
	goto L61
L60:
	;
	v227 = v220
	goto L61
L61:
	;
	v228 = base.B2i32(v227 < v165)
	if v227 < v165 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v229 = int32(4)
	goto L64
L63:
	;
	v229 = v225
	goto L64
L64:
	;
	v231 = base.B2i32(int32(0) < v165)
	if int32(0) < v165 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v232 = v229
	goto L67
L66:
	;
	v232 = v225
	goto L67
L67:
	;
	if v227 < v165 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v233 = v165
	goto L70
L69:
	;
	v233 = v227
	goto L70
L70:
	;
	if int32(0) < v165 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v234 = v233
	goto L73
L72:
	;
	v234 = v227
	goto L73
L73:
	;
	v235 = base.B2i32(v234 < v171)
	if v234 < v171 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v236 = int32(5)
	goto L76
L75:
	;
	v236 = v232
	goto L76
L76:
	;
	v238 = base.B2i32(int32(0) < v171)
	if int32(0) < v171 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v239 = v236
	goto L79
L78:
	;
	v239 = v232
	goto L79
L79:
	;
	if v234 < v171 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v240 = v171
	goto L82
L81:
	;
	v240 = v234
	goto L82
L82:
	;
	if int32(0) < v171 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v241 = v240
	goto L85
L84:
	;
	v241 = v234
	goto L85
L85:
	;
	v242 = base.B2i32(v241 < v162)
	if v241 < v162 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v243 = int32(6)
	goto L88
L87:
	;
	v243 = v239
	goto L88
L88:
	;
	v245 = base.B2i32(int32(0) < v162)
	if int32(0) < v162 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v246 = v243
	goto L91
L90:
	;
	v246 = v239
	goto L91
L91:
	;
	if v241 < v162 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v247 = v162
	goto L94
L93:
	;
	v247 = v241
	goto L94
L94:
	;
	if int32(0) < v162 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v248 = v247
	goto L97
L96:
	;
	v248 = v241
	goto L97
L97:
	;
	v249 = base.B2i32(v248 < v168)
	if v248 < v168 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v250 = int32(7)
	goto L100
L99:
	;
	v250 = v246
	goto L100
L100:
	;
	v252 = base.B2i32(int32(0) < v168)
	if int32(0) < v168 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v253 = v250
	goto L103
L102:
	;
	v253 = v246
	goto L103
L103:
	;
	if v248 < v168 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v254 = v168
	goto L106
L105:
	;
	v254 = v248
	goto L106
L106:
	;
	if int32(0) < v168 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v255 = v254
	goto L109
L108:
	;
	v255 = v248
	goto L109
L109:
	;
	if v255 < v178 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v257 = int32(8)
	goto L112
L111:
	;
	v257 = v253
	goto L112
L112:
	;
	v259 = base.B2i32(int32(0) < v178)
	if int32(0) < v178 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v260 = v257
	goto L115
L114:
	;
	v260 = v253
	goto L115
L115:
	;
	if v205+v210+v217+v224+v231+v238+v245+v252+v259 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	m.G0 = v34 + int32(96)
	return v36
L117:
	;
	F_range_gist_fallback_split(m, v44, v37, v36)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L339
	}
L118:
	;
	if v281 != 0 {
		goto L334
	} else {
		goto L335
	}
L119:
	;
	switch v260 & int32(-5) {
	case 0:
		goto L125
	case 1:
		goto L124
	case 2:
		goto L123
	default:
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v1129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v1129
	v1131 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v1131
	*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v1131
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v1131
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v1131
	if v201 == v1129 {
		goto L296
	} else {
		goto L297
	}
L122:
	;
	F_range_gist_fallback_split(m, v44, v37, v36)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L294
	}
L123:
	;
	F_range_gist_single_sorting_split(m, v44, v37, v36, int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L293
	}
L124:
	;
	F_range_gist_single_sorting_split(m, v44, v37, v36, int32(1))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L292
	}
L125:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v44)+272))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v276 = int32(_a_F_range_gist_picksplit_0)
	v277 = v275 + v276
	v279 = v277 & v276
	v281 = v279 << (uint(int32(4)) % 32)
	v282 = F_palloc(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v284 = F_palloc(m, v281)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v275&int32(_a_F_range_gist_picksplit_0) == int32(1) {
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v291 = v279 - int32(1)
	v294 = int32(1)
	goto L129
L129:
	;
	v324 = v294 << (uint(int32(4)) % 32)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v73+v324)))
	v327 = F_pg_detoast_datum(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	if v281 != 0 {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v329 = v324 + v282
	F_range_deserialize(m, v44, v327, v329-int32(16), v329-int32(8), v34)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v339 = (v294 + int32(1)) & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v339) <= base.Ui32(v279) {
		v294 = v339
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	base.MemoryCopy(m, v284, v282, v281)
	goto L136
L135:
	;
	goto L136
L136:
	;
	F_qsort_arg(m, v282, v279, int32(16), int32(1469), v44)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_qsort_arg(m, v284, v279, int32(16), int32(1470), v44)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v351 = v44 + int32(268)
	v352 = int32(1)
	v353 = int32(base.Ui32(v279) >> (uint(v352) % 32))
	v357 = int32(base.Ui32(v279+v352) >> (uint(v352) % 32))
	v358 = base.F32_convert_i32_u(v279)
	v359 = int32(0)
	v366 = v282
	v367 = v284
	v368 = v359
	v371 = v359
	v376 = v359
	v377 = v359
	v380 = v352
	v387 = v359
	v392 = v27
	v394 = v27
	goto L139
L139:
	;
	v398 = v367
	v402 = v371
	goto L142
L140:
	;
	v602 = v291 << (uint(int32(4)) % 32)
	v604 = int32(8)
	v609 = v291
	v610 = v602 + v284 + v604
	v611 = v282 + v602 + v604
	v612 = v291
	v620 = v581
	v621 = v582
	v624 = v585
	v631 = v592
	v636 = v597
	v638 = v599
	goto L184
L141:
	;
	goto L140
L142:
	;
	v429 = v282 + v402<<(uint(int32(4))%32)
	v430 = F_range_cmp_bounds(m, v44, v366, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	if v279 <= v368 {
		v491 = v368
		goto L153
	} else {
		goto L154
	}
L144:
	;
	if v430 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v435 = v429 + int32(8)
	v436 = F_range_cmp_bounds(m, v44, v435, v398)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	goto L143
L148:
	;
	if int32(0) < v436 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v440 = v435
	goto L151
L150:
	;
	v440 = v398
	goto L151
L151:
	;
	v442 = v402 + int32(1)
	if v442 < v279 {
		v398 = v440
		v402 = v442
		goto L142
	} else {
		goto L152
	}
L152:
	;
	v581 = v376
	v582 = v377
	v585 = v380
	v592 = v387
	v597 = v392
	v599 = v394
	goto L141
L153:
	;
	if v491 < v353 {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	v448 = v368
	goto L155
L155:
	;
	v481 = F_range_cmp_bounds(m, v44, v284+v448<<(uint(int32(4))%32)+int32(8), v398)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v491 = v279
	goto L153
L157:
	;
	if int32(0) < v481 {
		v491 = v448
		goto L153
	} else {
		goto L158
	}
L158:
	;
	v486 = v448 + int32(1)
	if v486 != v279 {
		v448 = v486
		goto L155
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	if v402 < v279 {
		v366 = v429
		v367 = v398
		v368 = v491
		v371 = v402
		v376 = v561
		v377 = v562
		v380 = v563
		v387 = v564
		v392 = v566
		v394 = v568
		goto L139
	} else {
		goto L183
	}
L161:
	;
	v520 = v491
	goto L163
L162:
	;
	v520 = v353
	goto L163
L163:
	;
	if v402 < v357 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v522 = v520
	goto L166
L165:
	;
	v522 = v402
	goto L166
L166:
	;
	v523 = v279 - v522
	if v522 < v523 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v525 = v522
	goto L169
L168:
	;
	v525 = v523
	goto L169
L169:
	;
	v527 = base.F32_div(base.F32_convert_i32_s(v525), v358)
	if base.F64_gt(base.F64_promote_f32(v527), float64(0.3)) == int32(0) {
		v561 = v376
		v562 = v377
		v563 = v380
		v564 = v387
		v566 = v392
		v568 = v394
		goto L160
	} else {
		goto L170
	}
L170:
	;
	if v274 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if base.F32_lt(v547, v392)|v380 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	v536 = F_FunctionCall2Coll(m, v351, v533, v534, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v547 = base.F32_convert_i32_s(v491 - v402)
	goto L171
L175:
	;
	v538 = *(*float64)(unsafe.Add(mBase, uint32(v536)))
	v539 = float64(0)
	if base.F64_ge(v538, v539) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v542 = v538
	goto L178
L177:
	;
	v542 = v539
	goto L178
L178:
	;
	v547 = base.F32_demote_f64(v542)
	goto L171
L179:
	;
	v552 = int32(0)
	if base.B2i32(base.F32_gt(v527, v394) == v552)|base.F32_ne(v392, v547) != 0 {
		v561 = v376
		v562 = v377
		v563 = v552
		v564 = v387
		v566 = v392
		v568 = v394
		goto L160
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v561 = v429
	v562 = v398
	v563 = int32(0)
	v564 = v491 - v522
	v566 = v547
	v568 = v527
	goto L160
L182:
	;
	goto L181
L183:
	;
	v581 = v561
	v582 = v562
	v585 = v563
	v592 = v564
	v597 = v566
	v599 = v568
	goto L141
L184:
	;
	v640 = v609
	v642 = v611
	goto L187
L185:
	;
	if v836 != 0 {
		goto L117
	} else {
		goto L231
	}
L186:
	;
	goto L185
L187:
	;
	v673 = v284 + v640<<(uint(int32(4))%32)
	v675 = v673 + int32(8)
	v676 = F_range_cmp_bounds(m, v44, v610, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L189
	}
L188:
	;
	if v612 < int32(0) {
		v737 = v612
		goto L198
	} else {
		goto L199
	}
L189:
	;
	if v676 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v680 = F_range_cmp_bounds(m, v44, v673, v642)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L188
L193:
	;
	if v680 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v684 = v673
	goto L196
L195:
	;
	v684 = v642
	goto L196
L196:
	;
	if int32(0) < v640 {
		v640 = v640 - int32(1)
		v642 = v684
		goto L187
	} else {
		goto L197
	}
L197:
	;
	v832 = v620
	v833 = v621
	v836 = v624
	v843 = v631
	goto L186
L198:
	;
	v766 = v640 + int32(1)
	if v766 < v353 {
		goto L208
	} else {
		goto L209
	}
L199:
	;
	v697 = v612
	goto L200
L200:
	;
	v725 = F_range_cmp_bounds(m, v44, v282+v697<<(uint(int32(4))%32), v642)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L202
	}
L201:
	;
	v737 = int32(-1)
	goto L198
L202:
	;
	if v725 < int32(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v737 = v697
	goto L198
L204:
	;
	goto L205
L205:
	;
	if int32(0) < v697 {
		v697 = v697 - int32(1)
		goto L200
	} else {
		goto L206
	}
L206:
	;
	goto L201
L207:
	;
	if int32(0) <= v640 {
		v609 = v640
		v610 = v675
		v611 = v642
		v612 = v737
		v620 = v811
		v621 = v812
		v624 = v813
		v631 = v814
		v636 = v816
		v638 = v818
		goto L184
	} else {
		goto L230
	}
L208:
	;
	v768 = v766
	goto L210
L209:
	;
	v768 = v353
	goto L210
L210:
	;
	v770 = v737 + int32(1)
	if v770 < v357 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v772 = v768
	goto L213
L212:
	;
	v772 = v770
	goto L213
L213:
	;
	v773 = v279 - v772
	if v772 < v773 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v775 = v772
	goto L216
L215:
	;
	v775 = v773
	goto L216
L216:
	;
	v777 = base.F32_div(base.F32_convert_i32_s(v775), v358)
	if base.F64_gt(base.F64_promote_f32(v777), float64(0.3)) == int32(0) {
		v811 = v620
		v812 = v621
		v813 = v624
		v814 = v631
		v816 = v636
		v818 = v638
		goto L207
	} else {
		goto L217
	}
L217:
	;
	if v274 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	if base.F32_lt(v797, v636)|v624 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L219:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
	v786 = F_FunctionCall2Coll(m, v351, v783, v784, v785)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v797 = base.F32_convert_i32_s(v640 - v737)
	goto L218
L222:
	;
	v788 = *(*float64)(unsafe.Add(mBase, uint32(v786)))
	v789 = float64(0)
	if base.F64_ge(v788, v789) != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v792 = v788
	goto L225
L224:
	;
	v792 = v789
	goto L225
L225:
	;
	v797 = base.F32_demote_f64(v792)
	goto L218
L226:
	;
	v802 = int32(0)
	if base.B2i32(base.F32_gt(v777, v638) == v802)|base.F32_ne(v636, v797) != 0 {
		v811 = v620
		v812 = v621
		v813 = v802
		v814 = v631
		v816 = v636
		v818 = v638
		goto L207
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v811 = v642
	v812 = v675
	v813 = int32(0)
	v814 = v766 - v772
	v816 = v797
	v818 = v777
	goto L207
L229:
	;
	goto L228
L230:
	;
	v832 = v811
	v833 = v812
	v836 = v813
	v843 = v814
	goto L186
L231:
	;
	v853 = v279 << (uint(int32(1)) % 32)
	v854 = F_palloc(m, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v854
	v857 = F_palloc(m, v853)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v859 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v859
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v859
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v857
	v864 = F_palloc(m, v281)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v866 = int32(1)
	v868 = int32(0)
	v871 = v866
	v872 = v868
	v873 = v866
	v879 = v868
	v880 = v868
	goto L235
L235:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v73+v873<<(uint(int32(4))%32))))
	v906 = F_pg_detoast_datum(m, v905)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L237
	}
L236:
	;
	if int32(0) < v988 {
		goto L270
	} else {
		goto L271
	}
L237:
	;
	v909 = v34 + int32(88)
	F_range_deserialize(m, v44, v906, v34, v909, v34+int32(87))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v914 = F_range_cmp_bounds(m, v44, v909, v833)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L240
	}
L239:
	;
	v992 = v871 + int32(1)
	v993 = int32(_a_F_range_gist_picksplit_0)
	v994 = v992 & v993
	if base.Ui32(v994) <= base.Ui32(v277&v993) {
		v871 = v992
		v872 = v985
		v873 = v994
		v879 = v988
		v880 = v989
		goto L235
	} else {
		goto L269
	}
L240:
	;
	if v914 <= int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v918 = F_range_cmp_bounds(m, v44, v34, v832)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if v969 <= int32(0) {
		goto L265
	} else {
		goto L266
	}
L244:
	;
	if int32(0) <= v918 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v924 = v864 + v879<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v924))) = v873
	if v274 != 0 {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	goto L247
L247:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v953 <= int32(0) {
		goto L260
	} else {
		goto L261
	}
L248:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	v929 = F_FunctionCall2Coll(m, v351, v926, v927, v928)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	v949 = float64(0)
	goto L250
L250:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v924)+8)) = v949
	v985 = v872
	v988 = v879 + int32(1)
	v989 = v880
	goto L239
L251:
	;
	v931 = *(*float64)(unsafe.Add(mBase, uint32(v929)))
	v932 = float64(0)
	if base.F64_ge(v931, v932) != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v935 = v931
	goto L254
L253:
	;
	v935 = v932
	goto L254
L254:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	v939 = F_FunctionCall2Coll(m, v351, v936, v937, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v941 = *(*float64)(unsafe.Add(mBase, uint32(v939)))
	v942 = float64(0)
	if base.F64_ge(v941, v942) != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v945 = v941
	goto L258
L257:
	;
	v945 = v942
	goto L258
L258:
	;
	v949 = base.F64_sub(v935, v945)
	goto L250
L259:
	;
	v961 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v959 + v961
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v964+v959<<(uint(v961)%32)))) = uint16(v871)
	v985 = v872
	v988 = v879
	v989 = v960
	goto L239
L260:
	;
	v959 = v953
	v960 = v906
	goto L259
L261:
	;
	goto L262
L262:
	;
	v956 = F_range_super_union(m, v44, v880, v906)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v959 = v958
	v960 = v956
	goto L259
L264:
	;
	v977 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v976 + v977
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v980+v976<<(uint(v977)%32)))) = uint16(v871)
	v985 = v975
	v988 = v879
	v989 = v880
	goto L239
L265:
	;
	v975 = v906
	v976 = v969
	goto L264
L266:
	;
	goto L267
L267:
	;
	v972 = F_range_super_union(m, v44, v872, v906)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v975 = v972
	v976 = v974
	goto L264
L269:
	;
	goto L236
L270:
	;
	F_pg_qsort(m, v864, v988, int32(16), int32(1471))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	v1089 = v985
	v1097 = v989
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v1089
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v1097
	goto L116
L273:
	;
	v1004 = int32(0)
	v1006 = v1004
	v1007 = v985
	v1008 = v1004
	v1015 = v989
	goto L274
L274:
	;
	v1037 = int32(4)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v864+v1006<<(uint(v1037)%32))))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v73+v1040<<(uint(v1037)%32))))
	v1045 = F_pg_detoast_datum(m, v1044)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L276
	}
L275:
	;
	v1089 = v1081
	v1097 = v1082
	goto L272
L276:
	;
	if v1006 < v843 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1084 = v1008 + int32(1)
	v1086 = v1084 & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v1086) < base.Ui32(v988) {
		v1006 = v1086
		v1007 = v1081
		v1008 = v1084
		v1015 = v1082
		goto L274
	} else {
		goto L291
	}
L278:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v1048 <= int32(0) {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	goto L280
L280:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if v1064 <= int32(0) {
		goto L287
	} else {
		goto L288
	}
L281:
	;
	v1056 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v1054 + v1056
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1059+v1054<<(uint(v1056)%32)))) = uint16(v1040)
	v1081 = v1007
	v1082 = v1055
	goto L277
L282:
	;
	v1054 = v1048
	v1055 = v1045
	goto L281
L283:
	;
	goto L284
L284:
	;
	v1051 = F_range_super_union(m, v44, v1015, v1045)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1054 = v1053
	v1055 = v1051
	goto L281
L286:
	;
	v1072 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1070 + v1072
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1075+v1070<<(uint(v1072)%32)))) = uint16(v1040)
	v1081 = v1071
	v1082 = v1015
	goto L277
L287:
	;
	v1070 = v1064
	v1071 = v1045
	goto L286
L288:
	;
	goto L289
L289:
	;
	v1067 = F_range_super_union(m, v44, v1007, v1045)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v1070 = v1069
	v1071 = v1067
	goto L286
L291:
	;
	goto L275
L292:
	;
	goto L116
L293:
	;
	goto L116
L294:
	;
	goto L116
L295:
	;
	v1192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v1193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v1193
	if v1192 != int32(1) {
		goto L306
	} else {
		goto L307
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(1)
	goto L295
L297:
	;
	goto L298
L298:
	;
	v1145 = v161 + v170 + v174 + v167
	v1146 = v51 - v1145
	v1148 = v161 + v165 + v178
	if v1148 <= int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1172 = int32(0)
	if base.B2i32(v1145 <= v1172)|base.B2i32(v1146 <= v1172) == v1172 {
		goto L303
	} else {
		goto L304
	}
L300:
	;
	v1151 = v51 - v1148
	if v1151 <= int32(0) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v1154 = v1151 - v1148
	v1155 = int32(31)
	v1156 = v1154 >> (uint(v1155) % 32)
	v1159 = v1146 - v1145
	v1161 = v1159 >> (uint(v1155) % 32)
	if v1159^v1161-v1161 < v1154^v1156-v1156 {
		goto L299
	} else {
		goto L302
	}
L302:
	;
	v1165 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v1165
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1165
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1165
	goto L295
L303:
	;
	v1179 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v1179
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v1179
	goto L295
L304:
	;
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+v260<<(uint(int32(2))%32)))) = int32(1)
	goto L295
L306:
	;
	v1201 = int32(1)
	v1207 = v1201
	v1213 = v1201
	v1215 = v1193
	v1218 = v1193
	goto L309
L307:
	;
	v1319 = v1193
	v1322 = v1193
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v1322
	goto L116
L309:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v73+v1213<<(uint(int32(4))%32))))
	v1242 = F_pg_detoast_datum(m, v1241)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	v1319 = v1304
	v1322 = v1305
	goto L308
L311:
	;
	v1307 = v1207 + int32(1)
	v1309 = v1307 & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v1309) <= base.Ui32((v1192-v1201)&int32(_a_F_range_gist_picksplit_0)) {
		v1207 = v1307
		v1213 = v1309
		v1215 = v1304
		v1218 = v1305
		goto L309
	} else {
		goto L333
	}
L312:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1242)))
	v1250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1242+int32(base.Ui32(v1244)>>(uint(int32(2))%32))-int32(1)))))
	goto L313
L313:
	;
	if v1250&int32(1) != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1264 = int32(8)
	goto L316
L315:
	;
	v1254 = int32(3)
	v1257 = int32(base.Ui32(v1250)>>(uint(v1254)%32)) & v1254
	if v1250 < int32(0) {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v34+v1264<<(uint(int32(2))%32))))
	if v1268 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	v1262 = v1257 | int32(4)
	goto L319
L318:
	;
	v1262 = v1257
	goto L319
L319:
	;
	v1264 = v1262
	goto L316
L320:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v1271 <= int32(0) {
		goto L324
	} else {
		goto L325
	}
L321:
	;
	goto L322
L322:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if v1287 <= int32(0) {
		goto L329
	} else {
		goto L330
	}
L323:
	;
	v1279 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v1277 + v1279
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1282+v1277<<(uint(v1279)%32)))) = uint16(v1207)
	v1304 = v1215
	v1305 = v1278
	goto L311
L324:
	;
	v1277 = v1271
	v1278 = v1242
	goto L323
L325:
	;
	goto L326
L326:
	;
	v1274 = F_range_super_union(m, v44, v1218, v1242)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1277 = v1276
	v1278 = v1274
	goto L323
L328:
	;
	v1295 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1293 + v1295
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1298+v1293<<(uint(v1295)%32)))) = uint16(v1207)
	v1304 = v1294
	v1305 = v1218
	goto L311
L329:
	;
	v1293 = v1287
	v1294 = v1242
	goto L328
L330:
	;
	goto L331
L331:
	;
	v1290 = F_range_super_union(m, v44, v1215, v1242)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v1293 = v1292
	v1294 = v1290
	goto L328
L333:
	;
	goto L310
L334:
	;
	base.MemoryCopy(m, v284, v282, v281)
	goto L336
L335:
	;
	goto L336
L336:
	;
	F_qsort_arg(m, v282, v279, int32(16), int32(1469), v44)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	F_qsort_arg(m, v284, v279, int32(16), int32(1470), v44)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	goto L117
L339:
	;
	goto L116
}
func F_range_gt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_range_in(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
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
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v339 int32
	_ = v339
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = F_get_range_io_data(m, l0, v18, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = v17
	goto L4
L4:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v36-int32(9)))&base.B2i32(v36 != int32(32)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v52 = v27
	v53 = int32(_a_F_range_in_0)
	v54 = int32(5)
	goto L15
L6:
	;
	v27 = v27 + int32(1)
	goto L4
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	m.G0 = v13 + int32(32)
	return v339
L10:
	;
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v296)
	v298 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)) = uint8(v298)
	v303 = int32(base.Ui32(v290)>>(uint(int32(3))%32)) & v298
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)) = uint8(v303)
	v308 = int32(base.Ui32(v290)>>(uint(int32(2))%32)) & v298
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v308)
	v313 = int32(base.Ui32(v290&int32(240)) >> (uint(int32(4)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v313)
	v318 = int32(base.Ui32(v290)>>(uint(v298)%32)) & v298
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+21)) = uint8(v318)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v327 = F_make_range(m, v320, v13+int32(16), v13+int32(8), v290&v298, v15)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L82
	}
L11:
	;
	v283 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v283)
	v339 = int32(0)
	goto L9
L12:
	;
	if v181&int32(41) != 0 {
		goto L75
	} else {
		goto L76
	}
L13:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L71
	}
L14:
	;
	if v99 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v99 = int32(0)
	goto L14
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 == v58 {
		v80 = v57
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	v82 = int32(1)
	if v80 != 0 {
		v52 = v52 + v82
		v53 = v53 + v82
		v54 = v54 - v82
		goto L15
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v57-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v68 = v57 | int32(32)
	goto L24
L23:
	;
	v68 = v57
	goto L24
L24:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v77 = v58 | int32(32)
	goto L27
L26:
	;
	v77 = v58
	goto L27
L27:
	;
	if v68 == v77 {
		v80 = v68
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v99 = v68 - v77
	goto L14
L29:
	;
	goto L19
L30:
	;
	v102 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v102
	v109 = v27 + int32(5)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v137 == int32(40) {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if base.B2i32(base.Ui32(v118-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v118 == int32(32)) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v131 = F_errsave_start(m, v15)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	v109 = v109 + int32(1)
	goto L33
L36:
	;
	if v118 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L34
L38:
	;
	v290 = int32(1)
	goto L10
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	if v131 == int32(0) {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	v229 = int32(_a_F_range_in_1)
	v238 = int32(2418)
	goto L13
L43:
	;
	v222 = F_errsave_start(m, v15)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L69
	}
L44:
	;
	v216 = F_errsave_start(m, v15)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L67
	}
L45:
	;
	v210 = F_errsave_start(m, v15)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L65
	}
L46:
	;
	v144 = int32(0)
	goto L48
L47:
	;
	if v137 != int32(91) {
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v150 = v13 + int32(16)
	v151 = F_range_parse_bound(m, v17, v27+int32(1), v13+int32(28), v150, v15)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v144 = int32(2)
	goto L48
L50:
	;
	if v151 == int32(0) {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 != int32(44) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)))
	v163 = F_range_parse_bound(m, v17, v151+int32(1), v13+int32(24), v150, v15)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v163 == int32(0) {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)))
	v173 = v158<<(uint(int32(3))%32) | v144 | v170<<(uint(int32(4))%32)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v174 != int32(41) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v174 != int32(93) {
		goto L43
	} else {
		goto L58
	}
L56:
	;
	v181 = v173
	goto L57
L57:
	;
	v183 = v163
	goto L59
L58:
	;
	v181 = v173 | int32(4)
	goto L57
L59:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if base.B2i32(base.Ui32(v194-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v194 == int32(32)) != 0 {
		v183 = v183 + int32(1)
		goto L59
	} else {
		goto L61
	}
L60:
	;
	if v194 == int32(0) {
		goto L12
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v204 = F_errsave_start(m, v15)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v204 == int32(0) {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v229 = int32(_a_F_range_in_2)
	v238 = int32(2481)
	goto L13
L65:
	;
	if v210 == int32(0) {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	v229 = int32(_a_F_range_in_3)
	v238 = int32(2435)
	goto L13
L67:
	;
	if v216 == int32(0) {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v229 = int32(_a_F_range_in_4)
	v238 = int32(2450)
	goto L13
L69:
	;
	if v222 == int32(0) {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	v229 = int32(_a_F_range_in_5)
	v238 = int32(2470)
	goto L13
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
	F_errmsg(m, int32(_a_F_range_in_6), v13)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errdetail(m, v229, int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errsave_finish(m, v15, int32(_a_F_range_in_7), v238, int32(_a_F_range_in_8))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L11
L75:
	;
	if v181&int32(81) != 0 {
		v290 = v181
		goto L10
	} else {
		goto L79
	}
L76:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v261 = F_InputFunctionCallSafe(m, v24+int32(4), v257, v258, v16, v15, v13+int32(16))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v261 != 0 {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L11
L79:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v271 = F_InputFunctionCallSafe(m, v24+int32(4), v267, v268, v16, v15, v13+int32(8))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v271 != 0 {
		v290 = v181
		goto L10
	} else {
		goto L81
	}
L81:
	;
	goto L11
L82:
	;
	v339 = v327
	goto L9
}
func F_range_lower_inf(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13985(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_range_ne(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_eq_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33 ^ int32(1)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_ne_0), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_ne_1), int32(1776), int32(_a_F_range_ne_2))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_eq_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33 ^ int32(1)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_ne_0), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_ne_1), int32(1776), int32(_a_F_range_ne_2))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_eq_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33 ^ int32(1)
						}
					}
				}
			}
		}
	}
}
func F_range_overright_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v45 = v4
		m.G0 = v8 + int32(48)
		return v45
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v45 = v4
			m.G0 = v8 + int32(48)
			return v45
		} else {
			v23 = v8 + int32(40)
			F_range_deserialize(m, l0, l1, v23, v8+int32(32), v8+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v34 = v8 + int32(24)
				F_multirange_get_bounds(m, l0, l2, int32(0), v34, v8+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = F_range_cmp_bounds(m, l0, v23, v34)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v45 = base.B2i32(int32(0) <= v39)
						m.G0 = v8 + int32(48)
						return v45
					}
				}
			}
		}
	}
}
func F_range_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(1468)
	return v4
}
func F_range_super_union(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	F_range_deserialize(m, l0, l1, v12+int32(40), v12+int32(24), v12+int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		F_range_deserialize(m, l0, l2, v12+int32(32), v12+int32(16), v12+int32(14))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v32)>>(uint(int32(2))%32))-int32(1)))))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v39)>>(uint(int32(2))%32))-int32(1)))))
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
			if v46 == int32(1) {
				if v45&int32(-127) == int32(0) {
					v58 = l2
					v61 = F_datumCopy(m, v58, int32(0), int32(-1))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v117 = v61
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
						v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
						v129 = v127 | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
						v132 = v117
						m.G0 = v12 + int32(48)
						return v132
					}
				} else {
					v132 = l2
					m.G0 = v12 + int32(48)
					return v132
				}
			} else {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
				if v53 != int32(1) {
					v67 = F_range_cmp_bounds(m, l0, v12+int32(40), v12+int32(32))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v70 = base.B2i32(v67 <= int32(0))
						v75 = F_range_cmp_bounds(m, l0, v12+int32(24), v12+int32(16))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							v77 = int32(0)
							if base.B2i32(v75 < v77)|base.B2i32(v77 < v67) == v77 {
								if v38 < int32(0) {
									v132 = l1
									m.G0 = v12 + int32(48)
									return v132
								} else {
									if v45 < int32(0) {
										if v67 <= int32(0) {
											v101 = v12 + int32(40)
										} else {
											v101 = v12 + int32(32)
										}
										if int32(0) <= v75 {
											v108 = v12 + int32(24)
										} else {
											v108 = v12 + int32(16)
										}
										v109 = int32(0)
										v111 = F_make_range(m, l0, v101, v108, v109, v109)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											if v38 < int32(0) {
												v117 = v111
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
												v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
												v129 = v127 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
												v132 = v117
											} else {
												if int32(0) <= v45 {
													v132 = v111
												} else {
													v117 = v111
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
													v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
													v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
													v129 = v127 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
													v132 = v117
												}
											}
											m.G0 = v12 + int32(48)
											return v132
										}
									} else {
										v132 = l1
										m.G0 = v12 + int32(48)
										return v132
									}
								}
							} else {
								if v70|base.B2i32(int32(0) <= v75) != 0 {
									if v67 <= int32(0) {
										v101 = v12 + int32(40)
									} else {
										v101 = v12 + int32(32)
									}
									if int32(0) <= v75 {
										v108 = v12 + int32(24)
									} else {
										v108 = v12 + int32(16)
									}
									v109 = int32(0)
									v111 = F_make_range(m, l0, v101, v108, v109, v109)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										if v38 < int32(0) {
											v117 = v111
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
											v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
											v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
											v129 = v127 | int32(128)
											*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
											v132 = v117
										} else {
											if int32(0) <= v45 {
												v132 = v111
											} else {
												v117 = v111
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
												v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
												v129 = v127 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
												v132 = v117
											}
										}
										m.G0 = v12 + int32(48)
										return v132
									}
								} else {
									v91 = int32(0)
									if base.B2i32(v45 < v91)|base.B2i32(v91 <= v38) != 0 {
										v132 = l2
										m.G0 = v12 + int32(48)
										return v132
									} else {
										if v67 <= int32(0) {
											v101 = v12 + int32(40)
										} else {
											v101 = v12 + int32(32)
										}
										if int32(0) <= v75 {
											v108 = v12 + int32(24)
										} else {
											v108 = v12 + int32(16)
										}
										v109 = int32(0)
										v111 = F_make_range(m, l0, v101, v108, v109, v109)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											if v38 < int32(0) {
												v117 = v111
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
												v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
												v129 = v127 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
												v132 = v117
											} else {
												if int32(0) <= v45 {
													v132 = v111
												} else {
													v117 = v111
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
													v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
													v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
													v129 = v127 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
													v132 = v117
												}
											}
											m.G0 = v12 + int32(48)
											return v132
										}
									}
								}
							}
						}
					}
				} else {
					if v38&int32(-127) != 0 {
						v132 = l1
						m.G0 = v12 + int32(48)
						return v132
					} else {
						v58 = l1
						v61 = F_datumCopy(m, v58, int32(0), int32(-1))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v117 = v61
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
							v126 = v117 + int32(base.Ui32(v121)>>(uint(int32(2))%32)) - int32(1)
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
							v129 = v127 | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
							v132 = v117
							m.G0 = v12 + int32(48)
							return v132
						}
					}
				}
			}
		}
	}
}
func F_range_table_entry_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	if l3&int32(16) != 0 {
		v7 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v7 != 0 {
				return int32(1)
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				switch v11 {
				case 0:
					v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v13 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v12, l2)
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return int32(0)
					} else {
						if v13 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							return int32(1)
						}
					}
				case 1:
					if l3&int32(1) != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v20 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v19, l2)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							if v20 != 0 {
								return int32(1)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != 0 {
										return int32(1)
									} else {
										if l3&int32(32) != 0 {
											v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 != 0 {
													return int32(1)
												} else {
													return int32(0)
												}
											}
										} else {
											return int32(0)
										}
									}
								}
							}
						}
					}
				case 2:
					if l3&int32(4) != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v25 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v24, l2)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							if v25 == int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != 0 {
										return int32(1)
									} else {
										if l3&int32(32) != 0 {
											v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 != 0 {
													return int32(1)
												} else {
													return int32(0)
												}
											}
										} else {
											return int32(0)
										}
									}
								}
							} else {
								return int32(1)
							}
						}
					}
				case 3:
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v30 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v29, l2)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							return int32(1)
						}
					}
				case 4:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v35 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v34, l2)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							return int32(1)
						}
					}
				case 5:
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					v40 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v39, l2)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v40 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							return int32(1)
						}
					}
				default:
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				case 9:
					if l3&int32(256) != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v47 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v46, l2)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							if v47 == int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != 0 {
										return int32(1)
									} else {
										if l3&int32(32) != 0 {
											v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 != 0 {
													return int32(1)
												} else {
													return int32(0)
												}
											}
										} else {
											return int32(0)
										}
									}
								}
							} else {
								return int32(1)
							}
						}
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		switch v11 {
		case 0:
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v13 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v12, l2)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					return int32(1)
				}
			}
		case 1:
			if l3&int32(1) != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != 0 {
						return int32(1)
					} else {
						if l3&int32(32) != 0 {
							v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									return int32(1)
								} else {
									return int32(0)
								}
							}
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v20 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v19, l2)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != 0 {
						return int32(1)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					}
				}
			}
		case 2:
			if l3&int32(4) != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != 0 {
						return int32(1)
					} else {
						if l3&int32(32) != 0 {
							v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									return int32(1)
								} else {
									return int32(0)
								}
							}
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v25 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v24, l2)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						return int32(1)
					}
				}
			}
		case 3:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v30 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v29, l2)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					return int32(1)
				}
			}
		case 4:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v35 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v34, l2)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				if v35 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					return int32(1)
				}
			}
		case 5:
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			v40 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v39, l2)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if v40 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					return int32(1)
				}
			}
		default:
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if v52 != 0 {
					return int32(1)
				} else {
					if l3&int32(32) != 0 {
						v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							if v56 != 0 {
								return int32(1)
							} else {
								return int32(0)
							}
						}
					} else {
						return int32(0)
					}
				}
			}
		case 9:
			if l3&int32(256) != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != 0 {
						return int32(1)
					} else {
						if l3&int32(32) != 0 {
							v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									return int32(1)
								} else {
									return int32(0)
								}
							}
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v47 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v46, l2)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					if v47 == int32(0) {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						return int32(1)
					}
				}
			}
		}
	}
}
