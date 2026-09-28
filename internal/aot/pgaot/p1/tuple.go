package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	v3 = int32(0)
	v15 = F_palloc(m, l0*int32(108)+int32(28))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	v22 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v22
	if l0 <= v19 {
		v126 = l0
		v129 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
	return v15
L4:
	;
	v33 = v3
	goto L5
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v44 = int32(100)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1+v33<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v15+v40<<(uint(int32(3))%32)+v33*v44+int32(28), v52, v44)
	F_populate_compact_attribute(m, v15, v33)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v60 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v61 <= v60 {
		v126 = v61
		v129 = v60
		goto L3
	} else {
		goto L9
	}
L7:
	;
	v58 = v33 + int32(1)
	if v58 != l0 {
		v33 = v58
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v65 = v15 + int32(28)
	v70 = v61
	v72 = v60
	v78 = v3
	goto L11
L10:
	;
	v126 = v102
	v129 = v123
	goto L3
L11:
	;
	v81 = v65 + v61<<(uint(int32(3))%32) + v72*int32(100)
	v84 = v65 + v72<<(uint(int32(3))%32)
	if v61 != v70 {
		v102 = v70
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v123 = v61
	goto L10
L13:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+2)))
	if v103 <= int32(0) {
		v123 = v72
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+7)))
	if v86 != int32(118) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v102 = v72
	goto L13
L16:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
	if v89 != int32(1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+6)))
	if v92&int32(6) != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+2)))
	if v95 <= int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+90)))
	if v98 != int32(118) {
		v102 = v61
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+90)))
	if v106 == int32(118) {
		v123 = v72
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+5)))
	v115 = (v78 + v109 - int32(1)) & (int32(0) - v109)
	if int32(_a_F_CreateTupleDesc_0) < v115 {
		v123 = v72
		goto L10
	} else {
		goto L23
	}
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v115)
	v121 = v72 + int32(1)
	if v121 != v61 {
		v70 = v102
		v72 = v121
		v78 = v115 + v103
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L12
}
func F_FreeTupleDesc(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_pfree(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L37
	}
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v11 = v7
	goto L7
L5:
	;
	v28 = v6
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v15 = v11 - int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8+v15<<(uint(int32(3))%32))+4))
	F_pfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_pfree(m, v8)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	if base.Ui32(int32(1)) < base.Ui32(v11) {
		v11 = v15
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = v26
	goto L6
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = v33 - int32(1)
	if int32(0) <= v35 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v73 = v28
	goto L15
L15:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+14)))
	if v77 != 0 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v39 = v35
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v32)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L26
	}
L19:
	;
	v45 = v32 + v39<<(uint(int32(4))%32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	if int32(0) < v39 {
		v39 = v39 - int32(1)
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v49<<(uint(int32(3))%32)+v39*int32(100))+110)))
	if v56 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	goto L20
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v73 = v71
	goto L15
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v81 = v77
	goto L30
L28:
	;
	v105 = v73
	goto L29
L29:
	;
	F_pfree(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L36
	}
L30:
	;
	v85 = v81 - int32(1)
	v88 = v78 + v85*int32(12)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	F_pfree(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v78)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L35
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(int32(1)) < base.Ui32(v81) {
		v81 = v85
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v105 = v99
	goto L29
L36:
	;
	goto L3
L37:
	;
	return
}
func F_TupleDescInitEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v55 int64
	_ = v55
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	v2 = l1
	v6 = l5
	v7 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = v2 * int32(100)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = v18 << (uint(int32(3)) % 32)
	v22 = v17 + (l0 + v20)
	v24 = v22 - int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v7
	v28 = v22 - int32(68)
	if l2 == v7 {
		if v28&int32(3) == int32(0) {
			v38 = l0 + v17 + v20 - int32(68)
			v39 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v38)+56)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+40)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39
		} else {
			v55 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v55
			*(*int64)(unsafe.Add(mBase, uint32(v28))) = v55
		}
	} else {
		if l2 == v28 {
		} else {
			v73 = F_strncpy(m, v28, l2, int32(64))
			mBase = m.M
			v74 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v73)+63)) = uint8(v74)
		}
	}
	v77 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+86)) = v77
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+80)) = uint16(v6)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+74)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = l4
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+90)) = uint16(v77)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+94)) = uint16(v77)
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+92)) = uint8(v86)
	v90 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		return
	} else {
		if v90 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l3
				F_errmsg_internal(m, int32(_a_F_TupleDescInitEntry_0), v14)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_TupleDescInitEntry_1), int32(963), int32(_a_F_TupleDescInitEntry_2))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v107 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = l3
			v110 = v107 + v108
			v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v111)
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v113)
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v115)
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+129)))
			v118 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v118)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v117)
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v110)+144))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v121
			F_populate_compact_attribute(m, l0, v2-int32(1))
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return
			} else {
				F_ReleaseCatCache(m, v90)
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return
				} else {
					m.G0 = v14 + int32(16)
					return
				}
			}
		}
	}
}
func F_apply_handle_tuple_routing(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
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
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int64
	_ = v468
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v720 int32
	_ = v720
	v5 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(96)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_palloc0(m, int32(264))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(402)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(0)
	v46 = F_ExecSetupPartitionTupleRouting(m, v34, v33)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v49 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v52 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v54 = v49
	goto L6
L6:
	;
	v55 = int32(_a_F_apply_handle_tuple_routing_0)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v58
	v60 = F_ExecFindPartition(m, v36, v32, v46, l1, v34)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v54 = v52
	goto L6
L8:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63)+119)))
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31)+25)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
	v67 = F_get_namespace_name(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	F_CheckSubscriptionRelkind(m, v64, v65, v67, v69+int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)+204))
	if v74 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v79 = F_table_slot_create(m, v62, v34+int32(104))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v81 = v74
	goto L13
L13:
	;
	v82 = F_ExecGetRootToChildMap(m, v60, v34)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v81 = v79
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v56
	if l3 != int32(3) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	if v82 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v85 = F_execute_attr_map_slot(m, v84, l1, v81)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	m.T0[v88].(func(*base.Module, int32, int32))(m, v81, l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v99 = v85
	v101 = v84
	goto L15
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81)+6)))
	if v92 <= v93 {
		v99 = v81
		v101 = v5
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	m.T0[v96].(func(*base.Module, int32, int32))(m, v81, v92)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v99 = v81
	v101 = v5
	goto L15
L24:
	;
	m.G0 = v29 + int32(96)
	return
L25:
	;
	v534 = F_GetTupleTransactionInfo(m, v481, v29+int32(24), v29+int32(28), v29+int32(32))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L108
	}
L26:
	;
	v106 = m.G0
	v108 = v106 + int32(-64)
	m.G0 = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v62)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+12)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[1]))
	if v114 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_InitConflictIndexes(m, v60)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L104
	}
L29:
	;
	v148 = v114
	goto L31
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2]))
	if v116 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v154 = F_hash_search(m, v148, v106+int32(-52), int32(1), v106+int32(-48))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L38
	}
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[3]))
	v126 = F_AllocSetContextCreateInternal(m, v121, int32(_a_F_apply_handle_tuple_routing_1), int32(0), int32(_a_F_apply_handle_tuple_routing_2), int32(_a_F_apply_handle_tuple_routing_3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v129 = v116
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+52)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = int64(343597383684)
	v139 = F_hash_create(m, int32(_a_F_apply_handle_tuple_routing_4), int64(64), v106+int32(-48), int32(1064))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2])) = v126
	v129 = v126
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[1])) = v139
	F_CacheRegisterRelcacheCallback(m, int32(1086))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[1]))
	v148 = v146
	goto L31
L38:
	;
	v157 = v154 + int32(8)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+16)))
	if v158 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	m.G0 = v108 - int32(-64)
	F_check_relation_updatable(m, v157)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L80
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+48)) = v62
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+44)) = v324
	if v101 != 0 {
		goto L65
	} else {
		goto L66
	}
L41:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v198 = F_pstrdup(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L52
	}
L42:
	;
	v178 = int32(_a_F_apply_handle_tuple_routing_0)
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v182
	v185 = v154 + int32(52)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v154)+52))
	if v186 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+40)))
	if v161 != int32(1) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v165 = int32(_a_F_apply_handle_tuple_routing_0)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v169
	base.MemoryFill(m, v154, int32(0), int32(80))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v174
	v193 = v154 + int32(52)
	v194 = v166
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+48)) = v62
	goto L39
L47:
	;
	F_free_attrmap(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v191 != 0 {
		v314 = v185
		v322 = v179
		goto L40
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(0)
	goto L49
L51:
	;
	v193 = v185
	v194 = v179
	goto L41
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+12)) = v198
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v202 = F_pstrdup(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+16)) = v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = v205
	v208 = F_palloc_mul(m, int32(4), v205)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = v208
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v213 = F_palloc_mul(m, int32(4), v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v213
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if int32(0) < v216 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v237 = v5
	goto L59
L57:
	;
	goto L58
L58:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+32)) = uint8(v291)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	v294 = F_bms_copy(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L63
	}
L59:
	;
	v246 = v237 << (uint(int32(2)) % 32)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v246+v247)))
	v250 = F_pstrdup(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v154)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v252+v246))) = v250
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257+v246)))
	*(*int32)(unsafe.Add(mBase, uint32(v255+v246))) = v259
	v262 = v237 + int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v262 < v263 {
		v237 = v262
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = v294
	v314 = v193
	v322 = v194
	goto L40
L64:
	;
	F_logicalrep_rel_mark_updatable(m, v157)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L78
	}
L65:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v327 = F_make_attrmap(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v385 = F_make_attrmap(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L76
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v327
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v330 <= int32(0) {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v333 = int32(0)
	v344 = v327
	v351 = v333
	v353 = v333
	goto L70
L70:
	;
	v362 = v351 << (uint(int32(1)) % 32)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v362+v363))))
	if v365 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L64
L72:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v366+v365<<(uint(int32(1))%32)-int32(2)))))
	v374 = v372
	goto L74
L73:
	;
	v374 = int32(_a_F_apply_handle_tuple_routing_5)
	goto L74
L74:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	*(*uint16)(unsafe.Add(mBase, uint32(v375+v362))) = uint16(v374)
	v379 = v353 + int32(1)
	v380 = base.I32_extend16_s(v379)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v380 < v382 {
		v344 = v381
		v351 = v380
		v353 = v379
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v385
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v390 = v388 << (uint(int32(1)) % 32)
	if v390 == int32(0) {
		goto L64
	} else {
		goto L77
	}
L77:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	base.MemoryCopy(m, v393, v394, v390)
	goto L64
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v322
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v154)+52))
	v427 = F_FindLogicalRepLocalIndex(m, v62, v31, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v429 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+40)) = uint8(v429)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+60)) = v427
	goto L39
L80:
	;
	if l3 != int32(2) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v157)+52))
	F_apply_handle_delete_internal(m, l0, v60, v99, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v468 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v468
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v468
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v468
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v157)+52))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_TargetPrivilegesCheck(m, v62, int64(2))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	goto L24
L85:
	;
	v481 = F_table_slot_create(m, v62, v475+int32(104))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v474 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v157)+52))
	v496 = F_FindDeletedTupleInLocalRel(m, v62, v489, v99, v29+int32(24), v29+int32(28), v29+int32(32))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L96
	}
L88:
	;
	v483 = F_RelationFindReplTupleByIndex(m, v62, v474, v99, v481)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v487 = F_RelationFindReplTupleSeq(m, v62, v99, v481)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	if v483 == int32(0) {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L25
L93:
	;
	if v487 != 0 {
		goto L25
	} else {
		goto L94
	}
L94:
	;
	goto L87
L95:
	;
	F_slot_store_data(m, v481, v157, l2)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L101
	}
L96:
	;
	if v496 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+28)))
	v501 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[4])))
	if v499 != v501 {
		v504 = int32(3)
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v504 = int32(4)
	goto L95
L100:
	;
	goto L99
L101:
	;
	v508 = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v508
	v515 = F_list_make1_impl(m, int32(1), v29+int32(4))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_ReportApplyConflict(m, v34, v60, int32(15), v504, v99, v481, v515)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L24
L104:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v522, int64(1))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_ExecSimpleRelationInsert(m, v60, v519, v99)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	goto L24
L107:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v562 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L108:
	;
	if v534 == int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+28)))
	v540 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[4])))
	if v538 == v540 {
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v544 = F_table_slot_create(m, v62, v34+int32(104))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_slot_store_data(m, v544, v157, l2)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v481
	v550 = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v550
	v554 = int32(1)
	v556 = F_list_make1_impl(m, v554, v29)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_ReportApplyConflict(m, v34, v60, int32(15), v554, v99, v544, v556)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L107
L115:
	;
	v565 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	v567 = v562
	goto L117
L117:
	;
	v568 = int32(_a_F_apply_handle_tuple_routing_0)
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v571
	F_slot_modify_data(m, v99, v481, v157, l2)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	v567 = v565
	goto L117
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v569
	v579 = int32(0)
	F_EvalPlanQualInit(m, v29+int32(44), v34, v579, v579, int32(-1), v579)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+131)))
	if v586 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	F_EvalPlanQualEnd(m, v29+int32(44))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L169
	}
L122:
	;
	if v82 != 0 {
		goto L132
	} else {
		goto L133
	}
L123:
	;
	v590 = F_ExecPartitionCheck(m, v60, v99, v34, int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	F_InitConflictIndexes(m, v60)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	if v590 == int32(0) {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v99
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v597, int64(4))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_ExecSimpleRelationUpdate(m, v60, v34, v29+int32(44), v481, v99)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	goto L121
L131:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v626 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L132:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v62)+52))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v607 = F_convert_tuples_by_name(m, v605, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)+32))
	m.T0[v613].(func(*base.Module, int32, int32))(m, l1, v99)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v607)+8))
	v610 = F_execute_attr_map_slot(m, v609, v99, l1)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v624 = v610
	goto L131
L137:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v618 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v617 <= v618 {
		v624 = l1
		goto L131
	} else {
		goto L138
	}
L138:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)+16))
	m.T0[v621].(func(*base.Module, int32, int32))(m, l1, v617)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v624 = l1
	goto L131
L140:
	;
	v629 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	v631 = v626
	goto L142
L142:
	;
	v632 = int32(_a_F_apply_handle_tuple_routing_0)
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v635
	v637 = F_ExecFindPartition(m, v36, v32, v46, v624, v34)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	v631 = v629
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v633
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+48))
	v643 = int32(*(*int8)(unsafe.Add(mBase, uint32(v642)+119)))
	v644 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31)+25)))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v642)+68))
	v646 = F_get_namespace_name(m, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v641)+48))
	F_CheckSubscriptionRelkind(m, v643, v644, v646, v648+int32(4))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v481
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_TargetPrivilegesCheck(m, v654, int64(8))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_ExecSimpleRelationDelete(m, v60, v34, v29+int32(44), v481)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
	if v662 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v665 = F_MakePerTupleExprContext(m, v34)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	v667 = v662
	goto L151
L151:
	;
	v668 = int32(_a_F_apply_handle_tuple_routing_0)
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0]))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v667)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v671
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v637)+204))
	if v673 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v667 = v665
	goto L151
L153:
	;
	v678 = F_table_slot_create(m, v641, v34+int32(104))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	v680 = v673
	goto L155
L155:
	;
	v681 = F_ExecGetRootToChildMap(m, v637, v34)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	v680 = v678
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_tuple_routing[0])) = v669
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_InitConflictIndexes(m, v637)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L166
	}
L158:
	;
	if v681 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v681)+8))
	v684 = F_execute_attr_map_slot(m, v683, v624, v680)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v680)+8))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v686)+32))
	m.T0[v687].(func(*base.Module, int32, int32))(m, v680, v624)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	v698 = v684
	goto L157
L163:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v624)+12))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	v692 = int32(*(*int16)(unsafe.Add(mBase, uint32(v624)+6)))
	if v691 <= v692 {
		v698 = v680
		goto L157
	} else {
		goto L164
	}
L164:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v624)+8))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+16))
	m.T0[v695].(func(*base.Module, int32, int32))(m, v624, v691)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v698 = v680
	goto L157
L166:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	F_TargetPrivilegesCheck(m, v705, int64(1))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_ExecSimpleRelationInsert(m, v637, v702, v698)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	goto L121
L169:
	;
	goto L24
}
func F_restore_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	F_BufFileReadExact(m, l0, v12+int32(28), int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v22 = F_palloc(m, v19+int32(24))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = v22 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_BufFileReadExact(m, l0, v25, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v31 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+8)) = uint16(v31)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v30
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v36
	F_ExecForceStoreHeapTuple(m, v22, l2, v31)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_BufFileReadExact(m, l0, v12+int32(24), int32(4))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if int32(0) < v46 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L49
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L46
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 <= int32(0) {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	m.G0 = v12 + int32(32)
	return
L12:
	;
	v57 = int32(0)
	goto L13
L13:
	;
	v66 = v57 << (uint(int32(3)) % 32)
	v67 = v49 + int32(28) + v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+6)))
	if v68&int32(4) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v158 != 0 {
		goto L7
	} else {
		goto L45
	}
L15:
	;
	v155 = v57 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v155 < v156 {
		v57 = v155
		goto L13
	} else {
		goto L44
	}
L16:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+2)))
	if v71 != int32(_a_F_restore_tuple_0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v75 = v57 + int32(1)
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v76 <= v57 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	m.T0[v79].(func(*base.Module, int32, int32))(m, l2, v75)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v57))))
	if v84 != 0 {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85+v66)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 != int32(1) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v91 != int32(1) {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v94 <= v57 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	m.T0[v97].(func(*base.Module, int32, int32))(m, l2, v75)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_BufFileReadExact(m, l0, v12+int32(16), int32(4))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if v105 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v131 = F_palloc(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L31:
	;
	v109 = int32(18)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+17)))
	if v111 == v109 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v122 = int32(1)
	if v105&v122 != 0 {
		v130 = int32(base.Ui32(v105) >> (uint(v122) % 32))
		goto L30
	} else {
		goto L40
	}
L34:
	;
	v114 = v109
	goto L36
L35:
	;
	v114 = int32(2)
	goto L36
L36:
	;
	if base.Ui32((v111-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v121 = int32(6)
	goto L39
L38:
	;
	v121 = v114
	goto L39
L39:
	;
	v130 = v121
	goto L30
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v130 = int32(base.Ui32(v126) >> (uint(int32(2)) % 32))
	goto L30
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v133
	v135 = int32(4)
	F_BufFileReadExact(m, l0, v131+v135, v130-v135)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v141+v66))) = base.I64_extend_i32_u(v131)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v147 = v145 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v147
	if v147 < int32(0) {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	goto L15
L44:
	;
	goto L14
L45:
	;
	goto L11
L46:
	;
	F_errmsg_internal(m, int32(_a_F_restore_tuple_1), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_restore_tuple_2), int32(2962), int32(_a_F_restore_tuple_3))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v197
	F_errmsg_internal(m, int32(_a_F_restore_tuple_4), v12)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_restore_tuple_2), int32(2968), int32(_a_F_restore_tuple_3))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
