package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_changeDependencyFor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
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
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l2
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l4
	v52 = base.B2i32(base.B2i32(l2 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_changeDependencyFor_0)) < base.Ui32(l4)) == int32(0)) & ((base.B2i32(l2 != int32(2615)) | base.B2i32(l4 != int32(2200))) & base.B2i32(l2 != int32(1262)))
	goto L2
L2:
	;
	if base.B2i32(base.B2i32(l2 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_changeDependencyFor_0)) < base.Ui32(l3)) == v6)&((base.B2i32(l2 != int32(2615))|base.B2i32(l3 != int32(2200)))&base.B2i32(l2 != int32(1262))) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(128)
	return v166
L4:
	;
	v53 = int32(1)
	if v52 != 0 {
		v166 = v53
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v52 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	F_recordMultipleDependencies(m, v13+int32(16), v13+int32(4), int32(1), int32(110))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v166 = v53
	goto L3
L10:
	;
	F_dependencyLockAndCheckObject(m, l2, l4)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v74 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v77 = v13 + int32(16)
	F_ScanKeyInit(m, v77, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	F_ScanKeyInit(m, v13+int32(72), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v96 = F_systable_beginscan(m, v74, int32(2673), int32(1), int32(0), int32(2), v77)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v98 = F_systable_getnext(m, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	if v98 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v100 = v98
	v109 = v6
	goto L22
L20:
	;
	v151 = v6
	goto L21
L21:
	;
	F_systable_endscan(m, v96)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L37
	}
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+22)))
	v112 = v110 + v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if v113 != l2 {
		v139 = v109
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v151 = v139
	goto L21
L24:
	;
	v140 = F_systable_getnext(m, v96)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L35
	}
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v115 != l3 {
		v139 = v109
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if v52 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v139 = v109 + int32(1)
	goto L24
L28:
	;
	F_simple_heap_delete(m, v74, v100+int32(4))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v121 = F_heap_copytuple(m, v100)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L8
	} else {
		goto L32
	}
L31:
	;
	goto L27
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v123+v124)+16)) = l4
	F_CatalogTupleUpdate(m, v74, v121+int32(4), v121)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v121)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	goto L27
L35:
	;
	if v140 != 0 {
		v100 = v140
		v109 = v139
		goto L22
	} else {
		goto L36
	}
L36:
	;
	goto L23
L37:
	;
	F_relation_close(m, v74, int32(3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v166 = v151
	goto L3
}
func F_change_useless_for_repack(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[0]))
	if v9 == v2 {
		v78 = v2
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v13 = int32(0)
		v15 = v6 + int32(4)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
		if v21 < v13 {
			v45 = v13
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(0))+76)))
			if v26 != int32(1) {
				v45 = v13
			} else {
				v30 = v20 + int32(76)
				if v15 != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v31
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v30)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v33
				} else {
				}
				v45 = int32(1)
			}
		}
		if v45 == int32(0) {
			v78 = v2
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v50 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[0]))
			if v48 != v50 {
				v61 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[1]))
				if base.B2i32(v61 == int32(0))|base.B2i32(v48 != v61) != 0 {
					v78 = int32(1)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					v69 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[2]))
					if v67 != v69 {
						v78 = int32(1)
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						v73 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[3]))
						if v71 == v73 {
							v78 = int32(0)
						} else {
							v78 = int32(1)
						}
					}
				}
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[4]))
				if v52 != v54 {
					v61 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[1]))
					if base.B2i32(v61 == int32(0))|base.B2i32(v48 != v61) != 0 {
						v78 = int32(1)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						v69 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[2]))
						if v67 != v69 {
							v78 = int32(1)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
							v73 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[3]))
							if v71 == v73 {
								v78 = int32(0)
							} else {
								v78 = int32(1)
							}
						}
					}
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[5]))
					if v56 == v58 {
						v78 = v2
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[1]))
						if base.B2i32(v61 == int32(0))|base.B2i32(v48 != v61) != 0 {
							v78 = int32(1)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							v69 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[2]))
							if v67 != v69 {
								v78 = int32(1)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
								v73 = *(*int32)(unsafe.Add(mBase, _c_F_change_useless_for_repack[3]))
								if v71 == v73 {
									v78 = int32(0)
								} else {
									v78 = int32(1)
								}
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v78
}
