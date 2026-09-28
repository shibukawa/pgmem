package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinBufferInit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_palloc0(m, int32(48))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(_a_F_GinBufferInit_0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+10)))
	v26 = F_palloc0_mul(m, int32(36), v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v26
	if int32(0) < v25 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v32 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v13 + int32(16)
	return v16
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v45 = v42 + v32*int32(36)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_GinBufferInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v32<<(uint(int32(2))%32))))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+20)) = uint8(v54)
	v57 = v32 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+10)) = uint16(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+9)) = uint8(v54)
	if v53 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v62 = v53
	goto L12
L11:
	;
	v62 = int32(100)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v62
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)))
	v76 = int32(4)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v66+v68*(base.I32_extend16_s(v57)-int32(1))<<(uint(int32(2))%32)+v76-v76)))
	goto L13
L13:
	;
	if v80 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v88 = v20 + v41<<(uint(int32(3))%32) + v32*int32(100)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+96))
	v91 = F_lookup_type_cache(m, v89, int32(64))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v97 = v80
	goto L16
L16:
	;
	F_PrepareSortSupportComparisonShim(m, v97, v45)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+108))
	if v93 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v97 = v93
	goto L16
L19:
	;
	if v57 != v25 {
		v32 = v57
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v88)+96))
	v123 = F_format_type_be(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v123
	F_errmsg(m, int32(_a_F_GinBufferInit_1), v13)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_GinBufferInit_2), int32(1325), int32(_a_F_GinBufferInit_3))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GinFormTuple(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v4 = l3
	v7 = l6
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v17 == int32(1) {
		v28 = l2
		v29 = base.B2i32(v4 != int32(0))
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = l2
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)) = uint8(base.B2i32(v4 != v23))
		v28 = base.I64_extend_i32_u(l1)
		v29 = v23
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v28
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+8))
	v40 = F_index_form_tuple(m, v35, v15+int32(16), v15+int32(14))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		return int32(0)
	} else {
		v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+6)))
		v46 = v44 & int32(_a_F_GinFormTuple_0)
		if v44 < int32(0) {
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
			if v51 != 0 {
				v52 = int32(17)
			} else {
				v52 = int32(19)
			}
			if base.Ui32(v52) < base.Ui32(v46) {
				v54 = v46
			} else {
				v54 = v52
			}
			v55 = v54
		} else {
			v55 = v46
		}
		*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v7)
		v57 = int32(_a_F_GinFormTuple_1)
		*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v57)
		v62 = (v55 + int32(1)) & int32(_a_F_GinFormTuple_2)
		*(*uint16)(unsafe.Add(mBase, uint32(v40)+2)) = uint16(v62)
		v68 = (l5 + v62 + int32(7)) & int32(-8)
		if base.Ui32(int32(2713)) <= base.Ui32(v68) {
			if l7 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(2712)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v138 + int32(4)
						F_errmsg(m, int32(_a_F_GinFormTuple_3), v15)
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GinFormTuple_4), int32(111), int32(_a_F_GinFormTuple_5))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
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
				F_pfree(m, v40)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v123 = int32(0)
					m.G0 = v15 + int32(32)
					return v123
				}
			}
		} else {
			if v68 != v46 {
				v75 = F_repalloc(m, v40, v68)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+6)))
					v79 = v77 & int32(_a_F_GinFormTuple_0)
					v80 = v68 - v79
					if v80 != 0 {
						base.MemoryFill(m, v75+v79, int32(0), v80)
					} else {
					}
					v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+6)))
					v87 = v84&int32(_a_F_GinFormTuple_6) | v68
					*(*uint16)(unsafe.Add(mBase, uint32(v75)+6)) = uint16(v87)
					v89 = v75
					v92 = int32(0)
					if base.B2i32(l4 == v92)|base.B2i32(l5 == v92) == v92 {
						v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+2)))
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89))))
						base.MemoryCopy(m, v99+(v89+v100<<(uint(int32(16))%32)&int32(2147418112)), l4, l5)
					} else {
					}
					if v4 == int32(0) {
						v123 = v89
					} else {
						v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+6)))
						if int32(0) <= v112 {
							v115 = int32(8)
						} else {
							v115 = int32(16)
						}
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
						if v119 != 0 {
							v120 = int32(0)
						} else {
							v120 = int32(2)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v89+v115+v120))) = uint8(v4)
						v123 = v89
					}
					m.G0 = v15 + int32(32)
					return v123
				}
			} else {
				v89 = v40
				v92 = int32(0)
				if base.B2i32(l4 == v92)|base.B2i32(l5 == v92) == v92 {
					v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+2)))
					v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89))))
					base.MemoryCopy(m, v99+(v89+v100<<(uint(int32(16))%32)&int32(2147418112)), l4, l5)
				} else {
				}
				if v4 == int32(0) {
					v123 = v89
				} else {
					v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+6)))
					if int32(0) <= v112 {
						v115 = int32(8)
					} else {
						v115 = int32(16)
					}
					v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					if v119 != 0 {
						v120 = int32(0)
					} else {
						v120 = int32(2)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v89+v115+v120))) = uint8(v4)
					v123 = v89
				}
				m.G0 = v15 + int32(32)
				return v123
			}
		}
	}
}
func F_ginFillScanEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32, l7 int32) int32 {
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
	var v15 int32
	_ = v15
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v145 int32
	_ = v145
	v2 = l1
	v3 = l2
	v6 = l5
	v7 = l6
	if l7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v145
L2:
	;
	v91 = F_palloc(m, int32(736))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L15
	} else {
		goto L19
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	if base.Ui32(v15-int32(100)) < base.Ui32(int32(-99)) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v41 = int32(0)
	v42 = v15
	goto L5
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1])))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v41<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v51 != 0 {
		v72 = v42
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v74 = v41 + int32(1)
	if base.Ui32(v74) < base.Ui32(v72) {
		v41 = v74
		v42 = v72
		goto L5
	} else {
		goto L18
	}
L8:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+9)))
	if v52 != v7 {
		v72 = v42
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+16)))
	if v54 != v3 {
		v72 = v42
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v56 != l3 {
		v72 = v42
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+24)))
	if v58 != v2 {
		v72 = v42
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+8)))
	if v60 != v6&int32(255) {
		v72 = v42
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v6 != 0 {
		v145 = v50
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32)+l0)+uint32(_c_F_ginFillScanEntry[2])))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
	v64 = F_FunctionCall2Coll(m, v2*int32(28)+l0+int32(116), v62, v63, l4)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if base.I32_wrap_i64(v64) == int32(0) {
		v145 = v50
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	v72 = v71
	goto L7
L18:
	;
	goto L6
L19:
	;
	v93 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+40)) = v93
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+24)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+16)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = l7
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+9)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v91)+28)) = v93
	v104 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+36)) = uint16(v104)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+648)) = v93
	v108 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+644)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v91)+656)) = v104
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[3])))
	if base.Ui32(v114) < base.Ui32(v115) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0]))) = v128 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v129+v128<<(uint(int32(2))%32)))) = v91
	v145 = v91
	goto L1
L21:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1])))
	v128 = v114
	v129 = v117
	goto L20
L22:
	;
	goto L23
L23:
	;
	v119 = v115 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[3]))) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1])))
	v123 = F_repalloc_mul(m, v121, int32(4), v119)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1]))) = v123
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	v128 = v126
	v129 = v123
	goto L20
}
func F_ginInitConsistentFunction(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v5 == int32(3) {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(58)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(59)
		return
	} else {
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+84)))
		v15 = l0 + v12*int32(28)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v15 + int32(3696)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v15 + int32(2800)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0+v12<<(uint(int32(2))%32))+uint32(_c_F_ginInitConsistentFunction[0])))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v27
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(2804))))
		if v33 != 0 {
			v34 = int32(60)
		} else {
			v34 = int32(61)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v34
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(3700))))
		if v40 != 0 {
			v41 = int32(62)
		} else {
			v41 = int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v41
		return
	}
}
func F_ginPostingListDecode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v10 = F_ginPostingListDecodeAllSegments(m, l0, (v3+int32(1))&int32(_a_F_ginPostingListDecode_0)+int32(8), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_ginPrepareEntryScan(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v2 = l1
	v4 = l3
	v6 = int32(0)
	base.MemoryFill(m, l0, v6, int32(72))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(50)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(51)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(53)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)) = uint16(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(54)
	return
}
func F_gin_bool_consistent(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= v3 {
		v141 = int32(0)
		m.G0 = v14 + int32(16)
		return v141
	} else {
		v21 = l0 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v21
		v24 = F_palloc_mul(m, int32(1), v16)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v24
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v29 <= int32(0) {
			} else {
				v32 = int32(0)
				if v29 != int32(1) {
					v40 = v32
					v42 = v32
					v50 = v3
					for {
						v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v40<<(uint(int32(3))%32)))))
						if v54 == int32(2) {
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v42))))
							*(*uint8)(unsafe.Add(mBase, uint32(v40+v24))) = uint8(v59)
							v63 = v42 + int32(1)
						} else {
							v63 = v42
						}
						v65 = v40 | int32(1)
						v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v65<<(uint(int32(3))%32)))))
						if v69 == int32(2) {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v63))))
							*(*uint8)(unsafe.Add(mBase, uint32(v24+v65))) = uint8(v74)
							v78 = v63 + int32(1)
						} else {
							v78 = v63
						}
						v79 = int32(2)
						v80 = v40 + v79
						v82 = v50 + v79
						if v82 != v29&int32(2147483646) {
							v40 = v80
							v42 = v78
							v50 = v82
							continue
						} else {
							break
						}
						break
					}
					if v29&int32(1) == int32(0) {
					} else {
						v86 = v80
						v88 = v78
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v86<<(uint(int32(3))%32)))))
						if v100 != int32(2) {
						} else {
							v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v88))))
							*(*uint8)(unsafe.Add(mBase, uint32(v86+v24))) = uint8(v105)
						}
					}
				} else {
					v86 = v32
					v88 = v32
					v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v86<<(uint(int32(3))%32)))))
					if v100 != int32(2) {
					} else {
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v88))))
						*(*uint8)(unsafe.Add(mBase, uint32(v86+v24))) = uint8(v105)
					}
				}
			}
			v121 = int32(8)
			v128 = F_execute(m, v21+v29<<(uint(int32(3))%32)-v121, v14+v121, int32(0), int32(1), int32(_a_F_gin_bool_consistent_0))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				v141 = v128
				m.G0 = v14 + int32(16)
				return v141
			}
		}
	}
}
func F_gin_check_parent_keys_consistency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int64
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v440 int64
	_ = v440
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v461 int64
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int64
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int64
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v761 int32
	_ = v761
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int64
	_ = v809
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int64
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1282 int32
	_ = v1282
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1351 int32
	_ = v1351
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1495 int32
	_ = v1495
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1703 int32
	_ = v1703
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1753 int32
	_ = v1753
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1921 int32
	_ = v1921
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1934 int64
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1962 int32
	_ = v1962
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2017 int32
	_ = v2017
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2040 int32
	_ = v2040
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2062 int32
	_ = v2062
	var v2067 int32
	_ = v2067
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2084 int32
	_ = v2084
	var v2089 int32
	_ = v2089
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2101 int32
	_ = v2101
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2125 int32
	_ = v2125
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2146 int32
	_ = v2146
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2168 int32
	_ = v2168
	var v2172 int32
	_ = v2172
	var v2177 int32
	_ = v2177
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2202 int32
	_ = v2202
	var v2207 int32
	_ = v2207
	v40 = m.G0
	v42 = v40 - int32(_a_F_gin_check_parent_keys_consistency_0)
	m.G0 = v42
	v45 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	v53 = F_AllocSetContextCreateInternal(m, v48, int32(_a_F_gin_check_parent_keys_consistency_1), int32(0), int32(_a_F_gin_check_parent_keys_consistency_2), int32(_a_F_gin_check_parent_keys_consistency_3))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v55 = int32(_a_F_gin_check_parent_keys_consistency_4)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v53
	F_initGinState(m, v42+int32(564), l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v64 = F_palloc0(m, int32(20))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = int64(8589934591)
	*(*int64)(unsafe.Add(mBase, uint32(v64))) = int64(0)
	v71 = v42 + int32(_a_F_gin_check_parent_keys_consistency_5)
	v73 = v42 + int32(704)
	v80 = v64
	v100 = int32(-1)
	goto L18
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L1
	} else {
		goto L432
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L1
	} else {
		goto L423
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L1
	} else {
		goto L414
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L1
	} else {
		goto L410
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L406
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L1
	} else {
		goto L402
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L1
	} else {
		goto L398
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L1
	} else {
		goto L394
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L390
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L1
	} else {
		goto L386
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L1
	} else {
		goto L382
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L1
	} else {
		goto L378
	}
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[1]))
	if v115 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v56
	F_MemoryContextDelete(m, v53)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L1
	} else {
		goto L377
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v121 = F_ReadBufferExtended(m, l0, v118, v119, v118, v45)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	F_LockBufferInternal(m, v121, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v126 = int32(0)
	v127 = base.B2i32(v126 <= v121)
	if v127 == v126 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+12)))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+16)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145+v147)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v127 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v121^int32(-1))<<(uint(int32(2))%32))))
	v145 = v137
	goto L26
L28:
	;
	goto L29
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v145 = v139 + v121<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+14)))
	if v169 == int32(0) {
		goto L7
	} else {
		goto L34
	}
L31:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+(v121^int32(-1))<<(uint(int32(2))%32))))
	v168 = v160
	goto L30
L32:
	;
	goto L33
L33:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v168 = v162 + v121<<(uint(int32(13))%32) + int32(-8192)
	goto L30
L34:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+19)))
	v173 = int32(8)
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+16)))
	if (v172<<(uint(v173)%32)-v175)&int32(_a_F_gin_check_parent_keys_consistency_6) != v173 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168+v175)+6)))
	if v182&int32(4) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v146) {
		goto L48
	} else {
		goto L49
	}
L37:
	;
	if v182&int32(2) == int32(0) {
		goto L9
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+12)))
	if base.Ui32(v221) < base.Ui32(int32(25)) {
		goto L36
	} else {
		goto L46
	}
L40:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+12)))
	if base.B2i32(base.Ui32(v189) < base.Ui32(int32(25)))|base.B2i32((v189+int32(_a_F_gin_check_parent_keys_consistency_7))&int32(_a_F_gin_check_parent_keys_consistency_8) == int32(0)) != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+532)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v42)+528)) = v206 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_9), v42+int32(528))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(705), int32(_a_F_gin_check_parent_keys_consistency_11))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	if base.Ui32(int32(409)) <= base.Ui32(int32(base.Ui32(v221+int32(_a_F_gin_check_parent_keys_consistency_7))>>(uint(int32(2))%32))&int32(_a_F_gin_check_parent_keys_consistency_6)) {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	goto L36
L48:
	;
	v240 = int32(base.Ui32(v146+int32(_a_F_gin_check_parent_keys_consistency_7)) >> (uint(int32(2)) % 32))
	goto L50
L49:
	;
	v240 = int32(0)
	goto L50
L50:
	;
	v243 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v243 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+500)) = v240 & int32(_a_F_gin_check_parent_keys_consistency_6)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+496)) = v245
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_12), v42+int32(496))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v261 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(443), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+16)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v353)+6)))
	if v355&int32(2) == int32(0) {
		v365 = v100
		goto L85
	} else {
		goto L86
	}
L58:
	;
	v265 = v42 + int32(564)
	v268 = F_gintuple_get_key(m, v265, v261, v42+int32(_a_F_gin_check_parent_keys_consistency_14))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v271 = F_gintuple_get_attrnum(m, v265, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v276 = F_PageGetItemIdCareful_1(m, l0, v273, v145, v240&int32(_a_F_gin_check_parent_keys_consistency_6))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v281 = v145 + v278&int32(_a_F_gin_check_parent_keys_consistency_15)
	v282 = F_gintuple_get_attrnum(m, v265, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v286 = F_gintuple_get_key(m, v265, v281, v42+int32(_a_F_gin_check_parent_keys_consistency_16))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v149 == int32(-1) {
		goto L57
	} else {
		goto L64
	}
L64:
	;
	if v271 != v282 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v316 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L77
	}
L66:
	;
	if base.Ui32(v282) < base.Ui32(v271) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v292 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[4]))))
	v293 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	if v292 != v293 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L57
L70:
	;
	if v292 < v293 {
		goto L65
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v292 != 0 {
		goto L57
	} else {
		goto L74
	}
L73:
	;
	goto L57
L74:
	;
	v297 = v271 - int32(1)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v71+v297<<(uint(int32(2))%32))))
	v305 = F_FunctionCall2Coll(m, v73+v297*int32(28), v304, v286, v268)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v305&int64(2147483648) == int64(0) {
		goto L57
	} else {
		goto L76
	}
L76:
	;
	goto L65
L77:
	;
	if v316 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+480)) = base.I64_rotl(v318, int64(32))
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_17), v42+int32(480))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v333 = F_palloc(m, int32(20))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(473), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v338 = F_CopyIndexTuple(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+4)) = v338
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v333)+12)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v333)+8)) = v341
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v333)+16)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v333
	goto L57
L85:
	;
	v367 = v240 & int32(_a_F_gin_check_parent_keys_consistency_6)
	if v367 != 0 {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v100 == int32(-1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v365 = v360
	goto L85
L88:
	;
	goto L89
L89:
	;
	if v360 != v100 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	v365 = v100
	goto L85
L91:
	;
	v368 = int32(0)
	v373 = v368
	v387 = int32(1)
	v393 = v368
	goto L94
L92:
	;
	goto L93
L93:
	;
	F_UnlockReleaseBuffer(m, v121)
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L1
	} else {
		goto L370
	}
L94:
	;
	v411 = v42 + int32(564)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v414 = v387 & int32(_a_F_gin_check_parent_keys_consistency_6)
	v415 = F_PageGetItemIdCareful_1(m, l0, v412, v145, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	goto L93
L96:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v420 = v145 + v417&int32(_a_F_gin_check_parent_keys_consistency_15)
	v421 = F_gintuple_get_attrnum(m, v411, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v426 = int32(7)
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+6)))
	if (int32(base.Ui32(v423)>>(uint(int32(17))%32))+v426)&int32(_a_F_gin_check_parent_keys_consistency_18) == (v430&int32(_a_F_gin_check_parent_keys_consistency_19)+v426)&int32(_a_F_gin_check_parent_keys_consistency_20) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+16)))
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v862)+6)))
	if v864&int32(2) == int32(0) {
		goto L186
	} else {
		goto L187
	}
L99:
	;
	v805 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L181
	}
L100:
	;
	F_UnlockReleaseBuffer(m, v533)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L180
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L176
	}
L102:
	;
	v440 = F_gintuple_get_key(m, v411, v420, v42+int32(563))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L172
	}
L105:
	;
	if v414 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if base.B2i32(v490 == int32(0))|base.B2i32(v414 != v367) != 0 {
		goto L98
	} else {
		goto L124
	}
L107:
	;
	if base.B2i32(v414 != v367)|base.B2i32(v149 != int32(-1)) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+16)))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v450)+6)))
	if v452&int32(2) == int32(0) {
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v461 = F_gintuple_get_key(m, v42+int32(564), v393, v42+int32(_a_F_gin_check_parent_keys_consistency_14))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v464 = v373 & int32(_a_F_gin_check_parent_keys_consistency_6)
	if v421 != v464 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	if base.Ui32(v464) < base.Ui32(v421) {
		goto L106
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v467 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	v468 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+563)))
	if v467 != v468 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L6
L117:
	;
	if v468 <= v467 {
		goto L6
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v467 != 0 {
		goto L6
	} else {
		goto L121
	}
L120:
	;
	goto L106
L121:
	;
	v472 = v464 - int32(1)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v71+v472<<(uint(int32(2))%32))))
	v480 = F_FunctionCall2Coll(m, v73+v472*int32(28), v479, v461, v440)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	if v480&int64(2147483648) == int64(0) {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	goto L106
L124:
	;
	v496 = v42 + int32(564)
	v497 = F_gintuple_get_attrnum(m, v496, v490)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v502 = F_gintuple_get_key(m, v496, v499, v42+int32(_a_F_gin_check_parent_keys_consistency_14))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v497 != v421 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	F_pfree(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L139
	}
L128:
	;
	if base.Ui32(v497) <= base.Ui32(v421) {
		goto L127
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v506 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+563)))
	v507 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	if v506 != v507 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L98
L132:
	;
	if v507 <= v506 {
		goto L127
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if v506 != 0 {
		goto L98
	} else {
		goto L136
	}
L135:
	;
	goto L98
L136:
	;
	v511 = v421 - int32(1)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v71+v511<<(uint(int32(2))%32))))
	v519 = F_FunctionCall2Coll(m, v73+v511*int32(28), v518, v440, v502)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	if base.I32_wrap_i64(v519) <= int32(0) {
		goto L98
	} else {
		goto L138
	}
L138:
	;
	goto L127
L139:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v530 = int32(0)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v533 = F_ReadBufferExtended(m, l0, v530, v531, v530, v45)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_LockBufferInternal(m, v533, int32(1))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if v533 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v555)+16)))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556+v555)+6)))
	if v558&int32(2) != 0 {
		goto L100
	} else {
		goto L146
	}
L143:
	;
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v541+(v533^int32(-1))<<(uint(int32(2))%32))))
	v555 = v547
	goto L142
L144:
	;
	goto L145
L145:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v555 = v549 + v533<<(uint(int32(13))%32) + int32(-8192)
	goto L142
L146:
	;
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v555)+12)))
	if base.Ui32(v561) < base.Ui32(int32(25)) {
		goto L100
	} else {
		goto L147
	}
L147:
	;
	v570 = int32(base.Ui32(v561+int32(_a_F_gin_check_parent_keys_consistency_7))>>(uint(int32(2))%32)) & int32(_a_F_gin_check_parent_keys_consistency_6)
	if v570 == int32(0) {
		goto L100
	} else {
		goto L148
	}
L148:
	;
	v576 = int32(1)
	goto L149
L149:
	;
	v614 = F_PageGetItemIdCareful_1(m, l0, v531, v555, v576&int32(_a_F_gin_check_parent_keys_consistency_6))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L151
	}
L150:
	;
	v631 = F_CopyIndexTuple(m, v619)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L156
	}
L151:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v619 = v555 + v616&int32(_a_F_gin_check_parent_keys_consistency_15)
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619))))
	v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619)+2)))
	if v529 != v620<<(uint(int32(16))%32)|v623 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v627 = v576 + int32(1)
	if base.Ui32(v627&int32(_a_F_gin_check_parent_keys_consistency_6)) <= base.Ui32(v570) {
		v576 = v627
		goto L149
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	goto L150
L155:
	;
	goto L100
L156:
	;
	F_UnlockReleaseBuffer(m, v533)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v631
	if v631 == int32(0) {
		goto L99
	} else {
		goto L158
	}
L158:
	;
	v639 = v42 + int32(564)
	v640 = F_gintuple_get_attrnum(m, v639, v631)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v645 = F_gintuple_get_key(m, v639, v642, v42+int32(_a_F_gin_check_parent_keys_consistency_14))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	if v640 != v421 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if base.Ui32(v421) < base.Ui32(v640) {
		goto L98
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v649 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+563)))
	v650 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	if v649 != v650 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	goto L101
L165:
	;
	if v650 <= v649 {
		goto L101
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if v649 != 0 {
		goto L98
	} else {
		goto L169
	}
L168:
	;
	goto L98
L169:
	;
	v654 = v421 - int32(1)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v71+v654<<(uint(int32(2))%32))))
	v662 = F_FunctionCall2Coll(m, v73+v654*int32(28), v661, v440, v645)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	if int32(0) < base.I32_wrap_i64(v662) {
		goto L101
	} else {
		goto L171
	}
L171:
	;
	goto L98
L172:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+456)) = v387 & int32(_a_F_gin_check_parent_keys_consistency_6)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+452)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v42)+448)) = v674 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_21), v42+int32(448))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(515), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+424)) = v240 & int32(_a_F_gin_check_parent_keys_consistency_6)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+420)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v42)+416)) = v702 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_22), v42+int32(416))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(593), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = int32(0)
	goto L99
L181:
	;
	if v805 == int32(0) {
		goto L98
	} else {
		goto L182
	}
L182:
	;
	v809 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+400)) = base.I64_rotl(v809, int64(32))
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_23), v42+int32(400))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(575), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	goto L98
L185:
	;
	if v393 != 0 {
		goto L364
	} else {
		goto L365
	}
L186:
	;
	v870 = F_palloc(m, int32(20))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v895 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420))))
	v897 = v895 << (uint(int32(16)) % 32)
	v898 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+2)))
	v899 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+4)))
	if v899 == int32(_a_F_gin_check_parent_keys_consistency_6) {
		goto L196
	} else {
		goto L197
	}
L189:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v870))) = v872 + int32(1)
	if v414 == v367 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v870)+4)) = v882
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v870)+8)) = v884
	v886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420)+2)))
	v887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420))))
	*(*int32)(unsafe.Add(mBase, uint32(v870)+12)) = v886 | v887<<(uint(int32(16))%32)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v870)+16)) = v892
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v870
	goto L185
L191:
	;
	if v149 == int32(-1) {
		v882 = int32(0)
		goto L190
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v880 = F_CopyIndexTuple(m, v420)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L195
	}
L194:
	;
	goto L193
L195:
	;
	v882 = v880
	goto L190
L196:
	;
	v903 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v1621 = v420 + v897&int32(2147418112) + v898
	if v897 < int32(0) {
		goto L342
	} else {
		goto L343
	}
L199:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	v911 = F_AllocSetContextCreateInternal(m, v906, int32(_a_F_gin_check_parent_keys_consistency_24), int32(0), int32(_a_F_gin_check_parent_keys_consistency_2), int32(_a_F_gin_check_parent_keys_consistency_3))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v913 = int32(_a_F_gin_check_parent_keys_consistency_4)
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v911
	v918 = F_palloc0(m, int32(24))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v920 = v897 | v898
	*(*int32)(unsafe.Add(mBase, uint32(v918)+16)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v918)+12)) = int32(-1)
	v924 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v918)+8)) = uint16(v924)
	*(*int64)(unsafe.Add(mBase, uint32(v918))) = int64(-4294967296)
	v930 = F_errstart(m, int32(12), v924)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	if v930 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+352)) = v920
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_25), v42+int32(352))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v945 = int32(-1)
	v946 = v918
	goto L208
L206:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(161), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v984 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[1]))
	if v984 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v914
	F_MemoryContextDelete(m, v911)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L1
	} else {
		goto L339
	}
L210:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v987 = int32(0)
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	v990 = F_ReadBufferExtended(m, l0, v987, v988, v987, v903)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L1
	} else {
		goto L214
	}
L213:
	;
	goto L212
L214:
	;
	F_LockBufferInternal(m, v990, int32(1))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	if v990 < int32(0) {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	F_UnlockReleaseBuffer(m, v990)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L1
	} else {
		goto L336
	}
L217:
	;
	v1013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+16)))
	v1014 = v1013 + v1012
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+6)))
	if v1015&int32(2) != 0 {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v998+(v990^int32(-1))<<(uint(int32(2))%32))))
	v1012 = v1004
	goto L217
L219:
	;
	goto L220
L220:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v1012 = v1006 + v990<<(uint(int32(13))%32) + int32(-8192)
	goto L217
L221:
	;
	v1018 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[6]))) = uint16(v1018)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[4]))) = v1018
	v1024 = F_errstart(m, int32(14), v1018)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1014)+4)))
	v1228 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L266
	}
L224:
	;
	if v1024 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+336)) = v1026
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_27), v42+int32(336))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	if v945 == int32(-1) {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(191), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v1043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+316)) = uint16(v1043)
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+312)) = v1045
	v1051 = F_GinDataLeafPageGetItems(m, v1012, v42+int32(_a_F_gin_check_parent_keys_consistency_28), v42+int32(312))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L235
	}
L231:
	;
	v1042 = v1038
	goto L230
L232:
	;
	goto L233
L233:
	;
	if v945 != v1038 {
		goto L16
	} else {
		goto L234
	}
L234:
	;
	v1042 = v945
	goto L230
L235:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[7])))
	if int32(0) < v1053 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v946)+12))
	v1106 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L242
	}
L237:
	;
	v1056 = int32(6)
	v1058 = v1051 + v1053*v1056
	v1061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058-int32(4)))))
	v1064 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058-v1056))))
	v1065 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1051)+2)))
	v1066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1051))))
	v1067 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1051)+4)))
	v1070 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058-int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+304)) = v1070
	*(*int32)(unsafe.Add(mBase, uint32(v42)+296)) = v1067
	v1073 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+292)) = v1065 | v1066<<(uint(v1073)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+288)) = v1053
	*(*int32)(unsafe.Add(mBase, uint32(v42)+300)) = v1061 | v1064<<(uint(v1073)%32)
	v1088 = F_pg_snprintf(m, v42+int32(_a_F_gin_check_parent_keys_consistency_14), int32(1024), int32(_a_F_gin_check_parent_keys_consistency_29), v42+int32(288))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1095 = F_pg_snprintf(m, v42+int32(_a_F_gin_check_parent_keys_consistency_14), int32(1024), int32(_a_F_gin_check_parent_keys_consistency_30), int32(0))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L241
	}
L240:
	;
	goto L236
L241:
	;
	goto L236
L242:
	;
	if v1103 != int32(-1) {
		goto L245
	} else {
		goto L246
	}
L243:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v946)+12))
	if v1159 == int32(-1) {
		v1571 = v1042
		goto L216
	} else {
		goto L253
	}
L244:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), v1151, int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L252
	}
L245:
	;
	if v1106 == int32(0) {
		goto L243
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	if v1106 == int32(0) {
		goto L243
	} else {
		goto L250
	}
L248:
	;
	v1113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+6)))
	v1114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+4)))
	v1115 = *(*int64)(unsafe.Add(mBase, uint32(v946)+12))
	v1116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+272)) = v42 + int32(_a_F_gin_check_parent_keys_consistency_14)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+268)) = v1116
	*(*int64)(unsafe.Add(mBase, uint32(v42)+256)) = base.I64_rotl(v1115, int64(32))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+264)) = v1113 | v1114<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_31), v42+int32(256))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v1151 = int32(219)
	goto L244
L250:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+240)) = v1136
	*(*int32)(unsafe.Add(mBase, uint32(v42)+244)) = v42 + int32(_a_F_gin_check_parent_keys_consistency_14)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_32), v42+int32(240))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1151 = int32(223)
	goto L244
L252:
	;
	goto L243
L253:
	;
	v1162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+8)))
	if v1162 == int32(0) {
		v1571 = v1042
		goto L216
	} else {
		goto L254
	}
L254:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[7])))
	if v1165 <= int32(0) {
		v1571 = v1042
		goto L216
	} else {
		goto L255
	}
L255:
	;
	v1169 = v946 + int32(4)
	v1170 = int32(6)
	v1174 = v1051 + v1165*v1170 - v1170
	v1178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1169)+2)))
	v1179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1169))))
	v1180 = int32(16)
	v1182 = v1178 | v1179<<(uint(v1180)%32)
	v1183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174)+2)))
	v1184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174))))
	v1187 = v1183 | v1184<<(uint(v1180)%32)
	if base.Ui32(v1182) < base.Ui32(v1187) {
		v1198 = int32(-1)
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if int32(0) <= v1198 {
		v1571 = v1042
		goto L216
	} else {
		goto L261
	}
L257:
	;
	goto L256
L258:
	;
	if base.Ui32(v1187) < base.Ui32(v1182) {
		v1198 = int32(1)
		goto L257
	} else {
		goto L259
	}
L259:
	;
	v1192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1169)+4)))
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174)+4)))
	if base.Ui32(v1192) < base.Ui32(v1193) {
		v1198 = int32(-1)
		goto L257
	} else {
		goto L260
	}
L260:
	;
	v1198 = base.B2i32(base.Ui32(v1193) < base.Ui32(v1192))
	goto L257
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+228)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v1208 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_33), v42+int32(224))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(231), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	if v1228 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+212)) = v1225
	*(*int32)(unsafe.Add(mBase, uint32(v42)+208)) = v1230
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_34), v42+int32(208))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v946)+12))
	v1247 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(246), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	if v1244 != int32(-1) {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	v1296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+12)))
	v1300 = base.I32_div_u_s(v1296-int32(32), int32(10))
	if v1300 != v1225 {
		goto L15
	} else {
		goto L283
	}
L274:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), v1288, int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L1
	} else {
		goto L282
	}
L275:
	;
	if v1247 == int32(0) {
		goto L273
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	if v1247 == int32(0) {
		goto L273
	} else {
		goto L280
	}
L278:
	;
	v1254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+6)))
	v1255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+4)))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v946)+12))
	v1258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+192)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v42)+184)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v42)+180)) = v1225
	*(*int32)(unsafe.Add(mBase, uint32(v42)+176)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v42)+188)) = v1254 | v1255<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_35), v42+int32(176))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v1288 = int32(252)
	goto L274
L280:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+164)) = v1225
	*(*int32)(unsafe.Add(mBase, uint32(v42)+160)) = v1275
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_36), v42+int32(160))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v1288 = int32(255)
	goto L274
L282:
	;
	goto L273
L283:
	;
	v1302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[8]))) = uint16(v1302)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5]))) = v1304
	v1307 = v946 + int32(4)
	v1308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+8)))
	v1309 = int32(0)
	if base.B2i32(v1308 == v1309)|base.B2i32(v1224 == int32(-1)) == v1309 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1307)+2)))
	v1319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1307))))
	v1320 = int32(16)
	v1323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[9]))))
	v1324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	if v1318|v1319<<(uint(v1320)%32) == v1323|v1324<<(uint(v1320)%32) {
		goto L289
	} else {
		goto L290
	}
L285:
	;
	goto L286
L286:
	;
	if v1225 == int32(0) {
		v1571 = v945
		goto L216
	} else {
		goto L294
	}
L287:
	;
	if v1334 == int32(0) {
		goto L14
	} else {
		goto L293
	}
L288:
	;
	goto L287
L289:
	;
	v1330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1307)+4)))
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[8]))))
	if v1330 == v1331 {
		v1334 = int32(1)
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1334 = int32(0)
	goto L288
L292:
	;
	goto L291
L293:
	;
	goto L286
L294:
	;
	v1351 = int32(1)
	goto L295
L295:
	;
	v1384 = v1351 * int32(10)
	v1385 = v1012 + int32(22) + v1384
	v1388 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L297
	}
L296:
	;
	v1571 = v945
	goto L216
L297:
	;
	v1390 = base.B2i32(v1225 != v1351)
	v1391 = int32(0)
	if base.B2i32(v1390 == v1391)&base.B2i32(v1224 == int32(-1)) == v1391 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	if v1225 != v1351 {
		goto L325
	} else {
		goto L326
	}
L299:
	;
	if v1388 != 0 {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	goto L301
L301:
	;
	if v1388 != 0 {
		goto L318
	} else {
		goto L319
	}
L302:
	;
	v1398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+6)))
	v1399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+4)))
	v1400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+2)))
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385))))
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+100)) = v1402
	v1404 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+104)) = v1400 | v1401<<(uint(v1404)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v1398 | v1399<<(uint(v1404)%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_37), v42+int32(96))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	if v1351 == int32(1) {
		goto L298
	} else {
		goto L307
	}
L305:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(312), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	goto L304
L307:
	;
	v1429 = v1385 + int32(4)
	v1432 = v1012 + int32(24) + v1384 - int32(8)
	v1436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1429)+2)))
	v1437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1429))))
	v1438 = int32(16)
	v1440 = v1436 | v1437<<(uint(v1438)%32)
	v1441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1432)+2)))
	v1442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1432))))
	v1445 = v1441 | v1442<<(uint(v1438)%32)
	if base.Ui32(v1440) < base.Ui32(v1445) {
		v1456 = int32(-1)
		goto L309
	} else {
		goto L310
	}
L308:
	;
	if int32(0) <= v1456 {
		goto L298
	} else {
		goto L313
	}
L309:
	;
	goto L308
L310:
	;
	if base.Ui32(v1445) < base.Ui32(v1440) {
		v1456 = int32(1)
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1429)+4)))
	v1451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1432)+4)))
	if base.Ui32(v1450) < base.Ui32(v1451) {
		v1456 = int32(-1)
		goto L309
	} else {
		goto L312
	}
L312:
	;
	v1456 = base.B2i32(base.Ui32(v1451) < base.Ui32(v1450))
	goto L309
L313:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+88)) = v1351
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = v1467
	*(*int32)(unsafe.Add(mBase, uint32(v42)+80)) = v1466 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_38), v42+int32(80))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(341), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L318:
	;
	v1483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+2)))
	v1484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385))))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+72)) = v1483 | v1484<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_37), v42-int32(-64))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v1503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+6)))
	v1504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+4)))
	if v1503|v1504 != 0 {
		goto L12
	} else {
		goto L323
	}
L321:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(315), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	goto L320
L323:
	;
	v1506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+8)))
	if v1506 != 0 {
		goto L12
	} else {
		goto L324
	}
L324:
	;
	goto L298
L325:
	;
	v1544 = F_palloc(m, int32(24))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L1
	} else {
		goto L334
	}
L326:
	;
	v1512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+8)))
	if v1512 == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1516 = v1385 + int32(4)
	v1520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1307)+2)))
	v1521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1307))))
	v1522 = int32(16)
	v1524 = v1520 | v1521<<(uint(v1522)%32)
	v1525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1516)+2)))
	v1526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1516))))
	v1529 = v1525 | v1526<<(uint(v1522)%32)
	if base.Ui32(v1524) < base.Ui32(v1529) {
		v1540 = int32(-1)
		goto L329
	} else {
		goto L330
	}
L328:
	;
	if v1540 < int32(0) {
		goto L13
	} else {
		goto L333
	}
L329:
	;
	goto L328
L330:
	;
	if base.Ui32(v1529) < base.Ui32(v1524) {
		v1540 = int32(1)
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1534 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1307)+4)))
	v1535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1516)+4)))
	if base.Ui32(v1534) < base.Ui32(v1535) {
		v1540 = int32(-1)
		goto L329
	} else {
		goto L332
	}
L332:
	;
	v1540 = base.B2i32(base.Ui32(v1535) < base.Ui32(v1534))
	goto L329
L333:
	;
	goto L325
L334:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v1547 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1544))) = v1546 + v1547
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+4)) = v1550
	v1552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1544)+8)) = uint16(v1552)
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+12)) = v1554
	v1556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+2)))
	v1557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385))))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+16)) = v1556 | v1557<<(uint(int32(16))%32)
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v946)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+20)) = v1562
	*(*int32)(unsafe.Add(mBase, uint32(v946)+20)) = v1544
	v1568 = (v1351 + v1547) & int32(_a_F_gin_check_parent_keys_consistency_6)
	if base.Ui32(v1568) <= base.Ui32(v1225) {
		v1351 = v1568
		goto L295
	} else {
		goto L335
	}
L335:
	;
	goto L296
L336:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v946)+20))
	F_pfree(m, v946)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	if v1611 != 0 {
		v945 = v1571
		v946 = v1611
		goto L208
	} else {
		goto L338
	}
L338:
	;
	goto L209
L339:
	;
	goto L185
L340:
	;
	F_pfree(m, v1715)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L363
	}
L341:
	;
	v1662 = int32(0)
	goto L359
L342:
	;
	if v899 != 0 {
		goto L345
	} else {
		goto L346
	}
L343:
	;
	goto L344
L344:
	;
	v1651 = F_palloc_mul(m, int32(6), v899)
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L1
	} else {
		goto L354
	}
L345:
	;
	v1626 = F_ginPostingListDecode(m, v1621, v42+int32(_a_F_gin_check_parent_keys_consistency_14))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v1648 = F_palloc(m, int32(0))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L1
	} else {
		goto L353
	}
L348:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5])))
	if v1628 == v899 {
		v1659 = v1626
		goto L341
	} else {
		goto L349
	}
L349:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+384)) = v899
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+388)) = v1635
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_39), v42+int32(384))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(113), int32(_a_F_gin_check_parent_keys_consistency_40))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L353:
	;
	v1715 = v1648
	goto L340
L354:
	;
	v1654 = v899 * int32(6)
	if v1654 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	base.MemoryCopy(m, v1651, v1621, v1654)
	goto L357
L356:
	;
	goto L357
L357:
	;
	if v899 == int32(0) {
		v1715 = v1651
		goto L340
	} else {
		goto L358
	}
L358:
	;
	v1659 = v1651
	goto L341
L359:
	;
	v1703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1659+v1662*int32(6))+4)))
	if base.Ui32(int32(2048)) <= base.Ui32((v1703-int32(1))&int32(_a_F_gin_check_parent_keys_consistency_6)) {
		goto L17
	} else {
		goto L361
	}
L360:
	;
	v1715 = v1659
	goto L340
L361:
	;
	v1711 = v1662 + int32(1)
	if v1711 != v899 {
		v1662 = v1711
		goto L359
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	goto L185
L364:
	;
	F_pfree(m, v393)
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L1
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v1795 = F_CopyIndexTuple(m, v420)
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L1
	} else {
		goto L368
	}
L367:
	;
	goto L366
L368:
	;
	v1798 = v387 + int32(1)
	if base.Ui32(v1798&int32(_a_F_gin_check_parent_keys_consistency_6)) <= base.Ui32(v367) {
		v373 = v421
		v387 = v1798
		v393 = v1795
		goto L94
	} else {
		goto L369
	}
L369:
	;
	goto L95
L370:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v1844 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	F_pfree(m, v1844)
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L1
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	F_pfree(m, v80)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L1
	} else {
		goto L375
	}
L374:
	;
	goto L373
L375:
	;
	if v1843 != 0 {
		v80 = v1843
		v100 = v365
		goto L18
	} else {
		goto L376
	}
L376:
	;
	goto L19
L377:
	;
	m.G0 = v42 + int32(_a_F_gin_check_parent_keys_consistency_0)
	return
L378:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+372)) = v1864
	*(*int32)(unsafe.Add(mBase, uint32(v42)+368)) = v1863 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_41), v42+int32(368))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(635), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L382:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+324)) = v1887
	*(*int32)(unsafe.Add(mBase, uint32(v42)+320)) = v1886 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_42), v42+int32(320))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(199), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L386:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+156)) = v1225
	*(*int32)(unsafe.Add(mBase, uint32(v42)+152)) = v1910
	*(*int32)(unsafe.Add(mBase, uint32(v42)+148)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v42)+144)) = v1909 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_43), v42+int32(144))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(274), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	v1934 = *(*int64)(unsafe.Add(mBase, uint32(v946)+12))
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+6)))
	v1937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+4)))
	v1938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+136)) = v1938
	v1940 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+132)) = v1936 | v1937<<(uint(v1940)%32)
	v1944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+120)) = v1944
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v1935 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+124)) = base.I64_rotl(v1934, int64(32))
	v1952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[9]))))
	v1953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+116)) = v1952 | v1953<<(uint(v1940)%32)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_44), v42+int32(112))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(299), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+40)) = v1225
	*(*int32)(unsafe.Add(mBase, uint32(v42)+36)) = v1976
	*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = v1975 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_45), v42+int32(32))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(354), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L398:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	v1999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+6)))
	v2000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+4)))
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	v2003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1385)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+60)) = v2003
	*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = v2002
	*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v2001 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+56)) = v1999 | v2000<<(uint(int32(16))%32)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_46), v42+int32(48))
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(331), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+468)) = v2031
	*(*int32)(unsafe.Add(mBase, uint32(v42)+464)) = v2030 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_42), v42+int32(464))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(494), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L406:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v2053 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_47), v42+int32(16))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(711), int32(_a_F_gin_check_parent_keys_consistency_11))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L410:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+516)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v42)+512)) = v2075 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_48), v42+int32(512))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(700), int32(_a_F_gin_check_parent_keys_consistency_11))
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L414:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v121 < int32(0) {
		goto L417
	} else {
		goto L418
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+548)) = v2116
	*(*int32)(unsafe.Add(mBase, uint32(v42)+544)) = v2097 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_49), v42+int32(544))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L1
	} else {
		goto L420
	}
L417:
	;
	v2101 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[10]))
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2101+(v121^int32(-1))*int32(56))+16))
	v2116 = v2107
	goto L416
L418:
	;
	goto L419
L419:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[11]))
	v2110 = int32(56)
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2109+v121*v2110-v2110)+16))
	v2116 = v2115
	goto L416
L420:
	;
	F_errhint(m, int32(_a_F_gin_check_parent_keys_consistency_50), int32(0))
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(692), int32(_a_F_gin_check_parent_keys_consistency_11))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L423:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v121 < int32(0) {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v2161
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v2142 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_51), v42)
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L1
	} else {
		goto L429
	}
L426:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[10]))
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2146+(v121^int32(-1))*int32(56))+16))
	v2161 = v2152
	goto L425
L427:
	;
	goto L428
L428:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[11]))
	v2155 = int32(56)
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2154+v121*v2155-v2155)+16))
	v2161 = v2160
	goto L425
L429:
	;
	F_errhint(m, int32(_a_F_gin_check_parent_keys_consistency_50), int32(0))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(681), int32(_a_F_gin_check_parent_keys_consistency_11))
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L432:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+444)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v42)+440)) = v387 & int32(_a_F_gin_check_parent_keys_consistency_6)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+436)) = v2189
	*(*int32)(unsafe.Add(mBase, uint32(v42)+432)) = v2188 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_52), v42+int32(432))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_10), int32(541), int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_consistent_jsonb_path(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = base.I32_wrap_i64(v15)
	if v16&int32(_a_F_gin_consistent_jsonb_path_0) == int32(7) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v76
L2:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v21)
	v23 = int32(0)
	v24 = int64(1)
	if v13 <= v23 {
		v76 = v24
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if base.Ui32((v16-int32(15))&int32(_a_F_gin_consistent_jsonb_path_0)) <= base.Ui32(int32(1)) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v27 = v23
	goto L6
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v14))))
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v76 = int64(0)
	goto L1
L8:
	;
	v37 = v27 + int32(1)
	if v13 != v37 {
		v27 = v37
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	v76 = v24
	goto L1
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v70 = F_execute_jsp_gin_node(m, v69, v14)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L21
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v47)
	if int32(0) < v13 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v76 = int64(1)
	goto L1
L17:
	;
	return int64(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v16 & int32(_a_F_gin_consistent_jsonb_path_0)
	F_errmsg_internal(m, int32(_a_F_gin_consistent_jsonb_path_1), v10)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_gin_consistent_jsonb_path_2), int32(1267), int32(_a_F_gin_consistent_jsonb_path_3))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
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
	v76 = base.I64_extend_i32_u(base.B2i32(v70 != int32(0)))
	goto L1
}
func F_gin_extract_jsonb_query(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
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
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
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
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v808 int32
	_ = v808
	var v823 int32
	_ = v823
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = base.I32_wrap_i64(v19)
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v22 = base.I32_wrap_i64(v21)
	switch v22&int32(_a_F_gin_extract_jsonb_query_0) - int32(7) {
	case 0:
		goto L5
	default:
		goto L3
	case 2:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L6
	} else {
		goto L155
	}
L2:
	;
	m.G0 = v16 - int32(-64)
	return base.I64_extend_i32_u(v808)
L3:
	;
	if v22&int32(_a_F_gin_extract_jsonb_query_1) == int32(10) {
		goto L73
	} else {
		goto L74
	}
L4:
	;
	v40 = int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v42 = F_pg_detoast_datum_packed(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = F_DirectFunctionCall2Coll(m, int32(1460), int32(0), v29, v19&int64(4294967295))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int64(0)
L7:
	;
	v36 = base.I32_wrap_i64(v32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v37 != 0 {
		v808 = v36
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	v808 = v36
	goto L2
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	v47 = F_palloc(m, int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v49 = int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v53 = v51 & v49
	if v53 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = v49
	goto L13
L12:
	;
	v54 = int32(4)
	goto L13
L13:
	;
	v55 = v42 + v54
	if v51 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v363 = v359 + int32(5)
	v364 = F_palloc(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L69
	}
L15:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v61 == int32(18) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v53 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v64 = int32(16)
	goto L20
L19:
	;
	v64 = int32(0)
	goto L20
L20:
	;
	if base.Ui32((v61-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(4)
	goto L23
L22:
	;
	v71 = v64
	goto L23
L23:
	;
	v359 = v71
	v360 = v40
	v361 = v55
	goto L14
L24:
	;
	v72 = int32(1)
	v81 = int32(base.Ui32(v51)>>(uint(v72)%32)) - v72
	goto L26
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v81 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L26:
	;
	if v81 < int32(126) {
		v359 = v81
		v360 = v40
		v361 = v55
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v89 = v81 - int32(1636608432)
	if v55&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v343 ^ v335 - base.I32_rotl(v343, int32(24))
	v350 = v14 + int32(-10)
	v355 = F_pg_snprintf(m, v350, int32(10), int32(_a_F_gin_extract_jsonb_query_2), v14+int32(-32))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L68
	}
L29:
	;
	v321 = int32(14)
	v323 = v317 ^ v318 - base.I32_rotl(v317, v321)
	v327 = v323 ^ v316 - base.I32_rotl(v323, int32(11))
	v331 = v327 ^ v317 - base.I32_rotl(v327, int32(25))
	v335 = v331 ^ v323 - base.I32_rotl(v331, int32(16))
	v339 = v335 ^ v327 - base.I32_rotl(v335, int32(4))
	v343 = v339 ^ v331 - base.I32_rotl(v339, v321)
	goto L28
L30:
	;
	switch v247 - int32(1) {
	case 0:
		v309 = v248
		v310 = v249
		v311 = v250
		goto L57
	case 1:
		v302 = v248
		v303 = v249
		v304 = v250
		goto L58
	case 2:
		v295 = v248
		v296 = v249
		v297 = v250
		goto L59
	case 3:
		v289 = v249
		v290 = v250
		goto L60
	case 4:
		v285 = v249
		v286 = v250
		goto L61
	case 5:
		v279 = v249
		v280 = v250
		goto L62
	case 6:
		v273 = v249
		v274 = v250
		goto L63
	case 7:
		v268 = v250
		goto L64
	case 8:
		v263 = v250
		goto L65
	case 9:
		v258 = v250
		goto L66
	case 10:
		goto L67
	default:
		v316 = v248
		v317 = v249
		v318 = v250
		goto L29
	}
L31:
	;
	v198 = v55
	v199 = v81
	v200 = v89
	v201 = v89
	v202 = v89
	goto L54
L32:
	;
	if base.Ui32(int32(11)) < base.Ui32(v81) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(v81) < base.Ui32(int32(12)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v246 = v55
	v247 = v81
	v248 = v89
	v249 = v89
	v250 = v89
	goto L30
L36:
	;
	switch v145 - int32(1) {
	case 0:
		v195 = v146
		goto L43
	case 1:
		v190 = v146
		goto L44
	case 2:
		goto L45
	case 3:
		v183 = v147
		goto L46
	case 4:
		v180 = v147
		goto L47
	case 5:
		v175 = v147
		goto L48
	case 6:
		goto L49
	case 7:
		v166 = v148
		goto L50
	case 8:
		v161 = v148
		goto L51
	case 9:
		v156 = v148
		goto L52
	case 10:
		goto L53
	default:
		v316 = v146
		v317 = v147
		v318 = v148
		goto L29
	}
L37:
	;
	v144 = v55
	v145 = v81
	v146 = v89
	v147 = v89
	v148 = v89
	goto L36
L38:
	;
	goto L39
L39:
	;
	v96 = v55
	v97 = v81
	v98 = v89
	v99 = v89
	v100 = v89
	goto L40
L40:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v103 = v102 + v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v107 = v106 + v100
	v109 = int32(4)
	v111 = v104 + v98 - v107 ^ base.I32_rotl(v107, v109)
	v115 = v103 - v111 ^ base.I32_rotl(v111, int32(6))
	v116 = v107 + v103
	v117 = v111 + v116
	v118 = v115 + v117
	v122 = v116 - v115 ^ base.I32_rotl(v115, int32(8))
	v126 = v117 - v122 ^ base.I32_rotl(v122, int32(16))
	v130 = v118 - v126 ^ base.I32_rotl(v126, int32(19))
	v131 = v122 + v118
	v132 = v126 + v131
	v133 = v130 + v132
	v137 = v131 - v130 ^ base.I32_rotl(v130, v109)
	v138 = int32(12)
	v139 = v96 + v138
	v141 = v97 - v138
	if base.Ui32(int32(11)) < base.Ui32(v141) {
		v96 = v139
		v97 = v141
		v98 = v132
		v99 = v133
		v100 = v137
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v144 = v139
	v145 = v141
	v146 = v132
	v147 = v133
	v148 = v137
	goto L36
L42:
	;
	goto L41
L43:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v316 = v195 + v196
	v317 = v147
	v318 = v148
	goto L29
L44:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	v195 = v191<<(uint(int32(8))%32) + v190
	goto L43
L45:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+2)))
	v190 = v186<<(uint(int32(16))%32) + v146
	goto L44
L46:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v316 = v184 + v146
	v317 = v183
	v318 = v148
	goto L29
L47:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+4)))
	v183 = v180 + v181
	goto L46
L48:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+5)))
	v180 = v176<<(uint(int32(8))%32) + v175
	goto L47
L49:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+6)))
	v175 = v171<<(uint(int32(16))%32) + v147
	goto L48
L50:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v316 = v167 + v146
	v317 = v169 + v147
	v318 = v166
	goto L29
L51:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+8)))
	v166 = v162<<(uint(int32(8))%32) + v161
	goto L50
L52:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+9)))
	v161 = v157<<(uint(int32(16))%32) + v156
	goto L51
L53:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+10)))
	v156 = v152<<(uint(int32(24))%32) + v148
	goto L52
L54:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v205 = v204 + v201
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	v209 = v208 + v202
	v211 = int32(4)
	v213 = v206 + v200 - v209 ^ base.I32_rotl(v209, v211)
	v217 = v205 - v213 ^ base.I32_rotl(v213, int32(6))
	v218 = v209 + v205
	v219 = v213 + v218
	v220 = v217 + v219
	v224 = v218 - v217 ^ base.I32_rotl(v217, int32(8))
	v228 = v219 - v224 ^ base.I32_rotl(v224, int32(16))
	v232 = v220 - v228 ^ base.I32_rotl(v228, int32(19))
	v233 = v224 + v220
	v234 = v228 + v233
	v235 = v232 + v234
	v239 = v233 - v232 ^ base.I32_rotl(v232, v211)
	v240 = int32(12)
	v241 = v198 + v240
	v243 = v199 - v240
	if base.Ui32(int32(11)) < base.Ui32(v243) {
		v198 = v241
		v199 = v243
		v200 = v234
		v201 = v235
		v202 = v239
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v246 = v241
	v247 = v243
	v248 = v234
	v249 = v235
	v250 = v239
	goto L30
L56:
	;
	goto L55
L57:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v316 = v309 + v312
	v317 = v310
	v318 = v311
	goto L29
L58:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	v309 = v305<<(uint(int32(8))%32) + v302
	v310 = v303
	v311 = v304
	goto L57
L59:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+2)))
	v302 = v298<<(uint(int32(16))%32) + v295
	v303 = v296
	v304 = v297
	goto L58
L60:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+3)))
	v295 = v291<<(uint(int32(24))%32) + v248
	v296 = v289
	v297 = v290
	goto L59
L61:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+4)))
	v289 = v285 + v287
	v290 = v286
	goto L60
L62:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+5)))
	v285 = v281<<(uint(int32(8))%32) + v279
	v286 = v280
	goto L61
L63:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+6)))
	v279 = v275<<(uint(int32(16))%32) + v273
	v280 = v274
	goto L62
L64:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+7)))
	v273 = v269<<(uint(int32(24))%32) + v249
	v274 = v268
	goto L63
L65:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+8)))
	v268 = v264<<(uint(int32(8))%32) + v263
	goto L64
L66:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+9)))
	v263 = v259<<(uint(int32(16))%32) + v258
	goto L65
L67:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+10)))
	v258 = v254<<(uint(int32(24))%32) + v250
	goto L66
L68:
	;
	v359 = int32(8)
	v360 = int32(17)
	v361 = v350
	goto L14
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+4)) = uint8(v360)
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v363 << (uint(int32(2)) % 32)
	if v359 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	base.MemoryCopy(m, v364+int32(5), v361, v359)
	goto L72
L71:
	;
	goto L72
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = base.I64_extend_i32_u(v364)
	v808 = v47
	goto L2
L73:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v380 = F_pg_detoast_datum(m, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(1)) < base.Ui32((v22-int32(15))&int32(_a_F_gin_extract_jsonb_query_0)) {
		goto L1
	} else {
		goto L151
	}
L76:
	;
	F_deconstruct_array_builtin(m, v380, int32(25), v14+int32(-16), v14+int32(-20), v14+int32(-24))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v393 = F_palloc_mul(m, int32(8), v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if int32(0) < v395 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v399 = int32(0)
	v400 = v395
	v404 = v2
	goto L82
L80:
	;
	v768 = v2
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v768
	if base.B2i32(v22&int32(_a_F_gin_extract_jsonb_query_0) != int32(11))|v768 != 0 {
		v808 = v393
		goto L2
	} else {
		goto L150
	}
L82:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+v399))))
	if v414 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v768 = v755
	goto L81
L84:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v417+v399<<(uint(int32(3))%32))))
	v422 = int32(1)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	v426 = v424 & v422
	if v426 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v754 = v400
	v755 = v404
	goto L86
L86:
	;
	v761 = v399 + int32(1)
	if v761 < v754 {
		v399 = v761
		v400 = v754
		v404 = v755
		goto L82
	} else {
		goto L149
	}
L87:
	;
	v427 = v422
	goto L89
L88:
	;
	v427 = int32(4)
	goto L89
L89:
	;
	v428 = v421 + v427
	if v424 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v736 = v732 + int32(5)
	v737 = F_palloc(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L145
	}
L91:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421)+1)))
	if v434 == int32(18) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	if v426 != 0 {
		goto L100
	} else {
		goto L101
	}
L94:
	;
	v437 = int32(16)
	goto L96
L95:
	;
	v437 = int32(0)
	goto L96
L96:
	;
	if base.Ui32((v434-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v444 = int32(4)
	goto L99
L98:
	;
	v444 = v437
	goto L99
L99:
	;
	v732 = v444
	v733 = v428
	v734 = int32(1)
	goto L90
L100:
	;
	v447 = int32(1)
	v456 = int32(base.Ui32(v424)>>(uint(v447)%32)) - v447
	goto L102
L101:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	v456 = int32(base.Ui32(v451)>>(uint(int32(2))%32)) - int32(4)
	goto L102
L102:
	;
	if v456 < int32(126) {
		v732 = v456
		v733 = v428
		v734 = int32(1)
		goto L90
	} else {
		goto L103
	}
L103:
	;
	v464 = v456 - int32(1636608432)
	if v428&int32(3) != 0 {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v718 ^ v710 - base.I32_rotl(v718, int32(24))
	v725 = v14 + int32(-10)
	v728 = F_pg_snprintf(m, v725, int32(10), int32(_a_F_gin_extract_jsonb_query_2), v16)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L6
	} else {
		goto L144
	}
L105:
	;
	v696 = int32(14)
	v698 = v692 ^ v693 - base.I32_rotl(v692, v696)
	v702 = v698 ^ v691 - base.I32_rotl(v698, int32(11))
	v706 = v702 ^ v692 - base.I32_rotl(v702, int32(25))
	v710 = v706 ^ v698 - base.I32_rotl(v706, int32(16))
	v714 = v710 ^ v702 - base.I32_rotl(v710, int32(4))
	v718 = v714 ^ v706 - base.I32_rotl(v714, v696)
	goto L104
L106:
	;
	switch v622 - int32(1) {
	case 0:
		v684 = v623
		v685 = v624
		v686 = v625
		goto L133
	case 1:
		v677 = v623
		v678 = v624
		v679 = v625
		goto L134
	case 2:
		v670 = v623
		v671 = v624
		v672 = v625
		goto L135
	case 3:
		v664 = v624
		v665 = v625
		goto L136
	case 4:
		v660 = v624
		v661 = v625
		goto L137
	case 5:
		v654 = v624
		v655 = v625
		goto L138
	case 6:
		v648 = v624
		v649 = v625
		goto L139
	case 7:
		v643 = v625
		goto L140
	case 8:
		v638 = v625
		goto L141
	case 9:
		v633 = v625
		goto L142
	case 10:
		goto L143
	default:
		v691 = v623
		v692 = v624
		v693 = v625
		goto L105
	}
L107:
	;
	v573 = v428
	v574 = v456
	v575 = v464
	v576 = v464
	v577 = v464
	goto L130
L108:
	;
	if base.Ui32(int32(11)) < base.Ui32(v456) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	if base.Ui32(v456) < base.Ui32(int32(12)) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v621 = v428
	v622 = v456
	v623 = v464
	v624 = v464
	v625 = v464
	goto L106
L112:
	;
	switch v520 - int32(1) {
	case 0:
		v570 = v521
		goto L119
	case 1:
		v565 = v521
		goto L120
	case 2:
		goto L121
	case 3:
		v558 = v522
		goto L122
	case 4:
		v555 = v522
		goto L123
	case 5:
		v550 = v522
		goto L124
	case 6:
		goto L125
	case 7:
		v541 = v523
		goto L126
	case 8:
		v536 = v523
		goto L127
	case 9:
		v531 = v523
		goto L128
	case 10:
		goto L129
	default:
		v691 = v521
		v692 = v522
		v693 = v523
		goto L105
	}
L113:
	;
	v519 = v428
	v520 = v456
	v521 = v464
	v522 = v464
	v523 = v464
	goto L112
L114:
	;
	goto L115
L115:
	;
	v471 = v428
	v472 = v456
	v473 = v464
	v474 = v464
	v475 = v464
	goto L116
L116:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	v478 = v477 + v474
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v471)+8))
	v482 = v481 + v475
	v484 = int32(4)
	v486 = v479 + v473 - v482 ^ base.I32_rotl(v482, v484)
	v490 = v478 - v486 ^ base.I32_rotl(v486, int32(6))
	v491 = v482 + v478
	v492 = v486 + v491
	v493 = v490 + v492
	v497 = v491 - v490 ^ base.I32_rotl(v490, int32(8))
	v501 = v492 - v497 ^ base.I32_rotl(v497, int32(16))
	v505 = v493 - v501 ^ base.I32_rotl(v501, int32(19))
	v506 = v497 + v493
	v507 = v501 + v506
	v508 = v505 + v507
	v512 = v506 - v505 ^ base.I32_rotl(v505, v484)
	v513 = int32(12)
	v514 = v471 + v513
	v516 = v472 - v513
	if base.Ui32(int32(11)) < base.Ui32(v516) {
		v471 = v514
		v472 = v516
		v473 = v507
		v474 = v508
		v475 = v512
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v519 = v514
	v520 = v516
	v521 = v507
	v522 = v508
	v523 = v512
	goto L112
L118:
	;
	goto L117
L119:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	v691 = v570 + v571
	v692 = v522
	v693 = v523
	goto L105
L120:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+1)))
	v570 = v566<<(uint(int32(8))%32) + v565
	goto L119
L121:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+2)))
	v565 = v561<<(uint(int32(16))%32) + v521
	goto L120
L122:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v691 = v559 + v521
	v692 = v558
	v693 = v523
	goto L105
L123:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+4)))
	v558 = v555 + v556
	goto L122
L124:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+5)))
	v555 = v551<<(uint(int32(8))%32) + v550
	goto L123
L125:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+6)))
	v550 = v546<<(uint(int32(16))%32) + v522
	goto L124
L126:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	v691 = v542 + v521
	v692 = v544 + v522
	v693 = v541
	goto L105
L127:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+8)))
	v541 = v537<<(uint(int32(8))%32) + v536
	goto L126
L128:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+9)))
	v536 = v532<<(uint(int32(16))%32) + v531
	goto L127
L129:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+10)))
	v531 = v527<<(uint(int32(24))%32) + v523
	goto L128
L130:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v580 = v579 + v576
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v573)+8))
	v584 = v583 + v577
	v586 = int32(4)
	v588 = v581 + v575 - v584 ^ base.I32_rotl(v584, v586)
	v592 = v580 - v588 ^ base.I32_rotl(v588, int32(6))
	v593 = v584 + v580
	v594 = v588 + v593
	v595 = v592 + v594
	v599 = v593 - v592 ^ base.I32_rotl(v592, int32(8))
	v603 = v594 - v599 ^ base.I32_rotl(v599, int32(16))
	v607 = v595 - v603 ^ base.I32_rotl(v603, int32(19))
	v608 = v599 + v595
	v609 = v603 + v608
	v610 = v607 + v609
	v614 = v608 - v607 ^ base.I32_rotl(v607, v586)
	v615 = int32(12)
	v616 = v573 + v615
	v618 = v574 - v615
	if base.Ui32(int32(11)) < base.Ui32(v618) {
		v573 = v616
		v574 = v618
		v575 = v609
		v576 = v610
		v577 = v614
		goto L130
	} else {
		goto L132
	}
L131:
	;
	v621 = v616
	v622 = v618
	v623 = v609
	v624 = v610
	v625 = v614
	goto L106
L132:
	;
	goto L131
L133:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	v691 = v684 + v687
	v692 = v685
	v693 = v686
	goto L105
L134:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+1)))
	v684 = v680<<(uint(int32(8))%32) + v677
	v685 = v678
	v686 = v679
	goto L133
L135:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+2)))
	v677 = v673<<(uint(int32(16))%32) + v670
	v678 = v671
	v679 = v672
	goto L134
L136:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+3)))
	v670 = v666<<(uint(int32(24))%32) + v623
	v671 = v664
	v672 = v665
	goto L135
L137:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+4)))
	v664 = v660 + v662
	v665 = v661
	goto L136
L138:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+5)))
	v660 = v656<<(uint(int32(8))%32) + v654
	v661 = v655
	goto L137
L139:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+6)))
	v654 = v650<<(uint(int32(16))%32) + v648
	v655 = v649
	goto L138
L140:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+7)))
	v648 = v644<<(uint(int32(24))%32) + v624
	v649 = v643
	goto L139
L141:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+8)))
	v643 = v639<<(uint(int32(8))%32) + v638
	goto L140
L142:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+9)))
	v638 = v634<<(uint(int32(16))%32) + v633
	goto L141
L143:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+10)))
	v633 = v629<<(uint(int32(24))%32) + v625
	goto L142
L144:
	;
	v732 = int32(8)
	v733 = v725
	v734 = int32(17)
	goto L90
L145:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v737)+4)) = uint8(v734)
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = v736 << (uint(int32(2)) % 32)
	if v732 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	base.MemoryCopy(m, v737+int32(5), v733, v732)
	goto L148
L147:
	;
	goto L148
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v393+v404<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v737)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v754 = v753
	v755 = v404 + int32(1)
	goto L86
L149:
	;
	goto L83
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	v808 = v393
	goto L2
L151:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v791 = F_pg_detoast_datum(m, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v797 = F_extract_jsp_query(m, v791, v22&int32(_a_F_gin_extract_jsonb_query_0), int32(0), v20, v796)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	if v797 != 0 {
		v808 = v797
		goto L2
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	v808 = int32(0)
	goto L2
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v22 & int32(_a_F_gin_extract_jsonb_query_0)
	F_errmsg_internal(m, int32(_a_F_gin_extract_jsonb_query_3), v14+int32(-48))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_gin_extract_jsonb_query_4), int32(922), int32(_a_F_gin_extract_jsonb_query_5))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_query_cidr(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_cidr_0), int32(_a_F_gin_extract_query_cidr_1), int32(0), int32(_a_F_gin_extract_query_cidr_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_varbit(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_varbit_0), int32(_a_F_gin_extract_query_varbit_1), int32(0), int32(_a_F_gin_extract_query_varbit_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_index_check(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_amcheck_lock_relation_and_check(m, v2, int32(2742), int32(_a_F_gin_index_check_0), int32(1), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
