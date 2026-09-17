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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
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
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_palloc0(m, int32(36))
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_GinBufferInit_0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+10)))
	v27 = F_palloc0(m, v24*int32(36))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v27
	if int32(0) < v24 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v33 = int32(0)
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v46 = v43 + v33*int32(36)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_GinBufferInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v33<<(uint(int32(2))%32))))
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+20)) = uint8(v55)
	v58 = v33 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+10)) = uint16(v58)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+9)) = uint8(v55)
	if v54 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v63 = v54
	goto L12
L11:
	;
	v63 = int32(100)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v63
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+6)))
	v77 = int32(4)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67+v69*(base.I32_extend16_s(v58)-int32(1))<<(uint(int32(2))%32)+v77-v77)))
	goto L13
L13:
	;
	if v81 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v89 = v20 + v42<<(uint(int32(4))%32) + v33*int32(100)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	v92 = F_lookup_type_cache(m, v90, int32(64))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v98 = v81
	goto L16
L16:
	;
	F_PrepareSortSupportComparisonShim(m, v98, v46)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+108))
	if v94 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v98 = v94
	goto L16
L19:
	;
	if v58 != v24 {
		v33 = v58
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
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	v124 = F_format_type_be(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v124
	F_errmsg(m, int32(_a_F_GinBufferInit_1), v13)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_GinBufferInit_2), int32(1311), int32(_a_F_GinBufferInit_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
func F_GinFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v4 = l3
	v7 = l6
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v16 == int32(1) {
		v26 = l2
		v27 = base.B2i32(v4 != int32(0))
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l2
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)) = uint8(base.B2i32(v4 != v22))
		v26 = l1
		v27 = v22
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)) = uint8(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v26
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+8))
	v38 = F_index_form_tuple(m, v33, v14+int32(24), v14+int32(22))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+6)))
		v44 = v42 & int32(_a_F_GinFormTuple_0)
		if v42 < int32(0) {
			v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
			if v49 != 0 {
				v50 = int32(17)
			} else {
				v50 = int32(19)
			}
			if base.Ui32(v50) < base.Ui32(v44) {
				v52 = v44
			} else {
				v52 = v50
			}
			v53 = v52
		} else {
			v53 = v44
		}
		*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)) = uint16(v7)
		v55 = int32(_a_F_GinFormTuple_1)
		*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v55)
		v60 = (v53 + int32(1)) & int32(_a_F_GinFormTuple_2)
		*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)) = uint16(v60)
		v66 = (v60 + l5 + int32(7)) & int32(-8)
		if base.Ui32(int32(2713)) <= base.Ui32(v66) {
			if l7 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(2712)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v66
						*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v136 + int32(4)
						F_errmsg(m, int32(_a_F_GinFormTuple_3), v14)
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GinFormTuple_4), int32(111), int32(_a_F_GinFormTuple_5))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
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
				F_pfree(m, v38)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					v121 = int32(0)
					m.G0 = v14 + int32(32)
					return v121
				}
			}
		} else {
			if v66 != v44 {
				v73 = F_repalloc(m, v38, v66)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)))
					v77 = v75 & int32(_a_F_GinFormTuple_0)
					v78 = v66 - v77
					if v78 != 0 {
						base.MemoryFill(m, v73+v77, int32(0), v78)
					} else {
					}
					v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)))
					v85 = v82&int32(_a_F_GinFormTuple_6) | v66
					*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)) = uint16(v85)
					v87 = v73
					v90 = int32(0)
					if base.B2i32(l4 == v90)|base.B2i32(l5 == v90) == v90 {
						v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+2)))
						v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
						base.MemoryCopy(m, v97+(v87+v98<<(uint(int32(16))%32)&int32(2147418112)), l4, l5)
					} else {
					}
					if v4 == int32(0) {
						v121 = v87
					} else {
						v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+6)))
						if int32(0) <= v110 {
							v113 = int32(8)
						} else {
							v113 = int32(16)
						}
						v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
						if v117 != 0 {
							v118 = int32(0)
						} else {
							v118 = int32(2)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v87+v113+v118))) = uint8(v4)
						v121 = v87
					}
					m.G0 = v14 + int32(32)
					return v121
				}
			} else {
				v87 = v38
				v90 = int32(0)
				if base.B2i32(l4 == v90)|base.B2i32(l5 == v90) == v90 {
					v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+2)))
					v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
					base.MemoryCopy(m, v97+(v87+v98<<(uint(int32(16))%32)&int32(2147418112)), l4, l5)
				} else {
				}
				if v4 == int32(0) {
					v121 = v87
				} else {
					v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+6)))
					if int32(0) <= v110 {
						v113 = int32(8)
					} else {
						v113 = int32(16)
					}
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					if v117 != 0 {
						v118 = int32(0)
					} else {
						v118 = int32(2)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v87+v113+v118))) = uint8(v4)
					v121 = v87
				}
				m.G0 = v14 + int32(32)
				return v121
			}
		}
	}
}
func F_ginFillScanEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
	var v13 int32
	_ = v13
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
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
	return v126
L2:
	;
	v72 = F_palloc(m, int32(728))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L17
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	if base.Ui32(v13-int32(100)) < base.Ui32(int32(-99)) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = int32(0)
	v30 = v13
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1])))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v29<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v37 != 0 {
		v55 = v30
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v57 = v29 + int32(1)
	if base.Ui32(v57) < base.Ui32(v55) {
		v29 = v57
		v30 = v55
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)))
	if v38 != v7 {
		v55 = v30
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
	if v40 != v3 {
		v55 = v30
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v42 != l3 {
		v55 = v30
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+20)))
	if v44 != v2 {
		v55 = v30
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36)+4)))
	v48 = F_ginCompareEntries(m, l0+int32(4), v2, v46, v47, l4, v6)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v48 == int32(0) {
		v126 = v36
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	v55 = v54
	goto L7
L16:
	;
	goto L6
L17:
	;
	v74 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v72)+36)) = v74
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+20)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = l7
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+5)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v72)+24)) = v74
	v85 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+32)) = uint16(v85)
	*(*int64)(unsafe.Add(mBase, uint32(v72)+644)) = v74
	v89 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+640)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v72)+44)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v72)+652)) = v85
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[2])))
	if base.Ui32(v95) < base.Ui32(v96) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0]))) = v109 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v110+v109<<(uint(int32(2))%32)))) = v72
	v126 = v72
	goto L1
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1])))
	v109 = v95
	v110 = v98
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[2]))) = v96 << (uint(int32(1)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1])))
	v105 = F_repalloc(m, v102, v96<<(uint(int32(3))%32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[1]))) = v105
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFillScanEntry[0])))
	v109 = v108
	v110 = v105
	goto L18
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v5 == int32(3) {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(55)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(56)
		return
	} else {
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+76)))
		v15 = l0 + v12*int32(28)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v15 + int32(3696)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v15 + int32(2800)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0+v12<<(uint(int32(2))%32))+uint32(_c_F_ginInitConsistentFunction[0])))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v27
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(2804))))
		if v33 != 0 {
			v34 = int32(57)
		} else {
			v34 = int32(58)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v34
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(3700))))
		if v40 != 0 {
			v41 = int32(59)
		} else {
			v41 = int32(60)
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
func F_ginPrepareEntryScan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	base.MemoryFill(m, l0, v6, int32(68))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(43)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(45)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l2
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)) = uint16(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(51)
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= v3 {
		v140 = int32(0)
		m.G0 = v14 + int32(16)
		return v140
	} else {
		v21 = l0 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v21
		v23 = F_palloc(m, v16)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v23
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v28 <= int32(0) {
			} else {
				v31 = int32(0)
				if v28 != int32(1) {
					v39 = v31
					v41 = v31
					v49 = v3
					for {
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v39<<(uint(int32(3))%32)))))
						if v53 == int32(2) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v41))))
							*(*uint8)(unsafe.Add(mBase, uint32(v39+v23))) = uint8(v58)
							v62 = v41 + int32(1)
						} else {
							v62 = v41
						}
						v64 = v39 | int32(1)
						v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v64<<(uint(int32(3))%32)))))
						if v68 == int32(2) {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v62))))
							*(*uint8)(unsafe.Add(mBase, uint32(v23+v64))) = uint8(v73)
							v77 = v62 + int32(1)
						} else {
							v77 = v62
						}
						v78 = int32(2)
						v79 = v39 + v78
						v81 = v49 + v78
						if v81 != v28&int32(2147483646) {
							v39 = v79
							v41 = v77
							v49 = v81
							continue
						} else {
							break
						}
						break
					}
					if v28&int32(1) == int32(0) {
					} else {
						v85 = v79
						v87 = v77
						v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v85<<(uint(int32(3))%32)))))
						if v99 != int32(2) {
						} else {
							v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v87))))
							*(*uint8)(unsafe.Add(mBase, uint32(v85+v23))) = uint8(v104)
						}
					}
				} else {
					v85 = v31
					v87 = v31
					v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v85<<(uint(int32(3))%32)))))
					if v99 != int32(2) {
					} else {
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v87))))
						*(*uint8)(unsafe.Add(mBase, uint32(v85+v23))) = uint8(v104)
					}
				}
			}
			v120 = int32(8)
			v127 = F_execute(m, v21+v28<<(uint(int32(3))%32)-v120, v14+v120, int32(0), int32(1), int32(_a_F_gin_bool_consistent_0))
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return int32(0)
			} else {
				v140 = v127
				m.G0 = v14 + int32(16)
				return v140
			}
		}
	}
}
func F_gin_check_parent_keys_consistency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int64
	_ = v288
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
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
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v695 int32
	_ = v695
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int64
	_ = v739
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int64
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1274 int32
	_ = v1274
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1615 int32
	_ = v1615
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1661 int32
	_ = v1661
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1774 int32
	_ = v1774
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int64
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1863 int32
	_ = v1863
	var v1868 int32
	_ = v1868
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1887 int32
	_ = v1887
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1918 int32
	_ = v1918
	var v1923 int32
	_ = v1923
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1985 int32
	_ = v1985
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2026 int32
	_ = v2026
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2069 int32
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	v36 = m.G0
	v38 = v36 - int32(_a_F_gin_check_parent_keys_consistency_0)
	m.G0 = v38
	v41 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(_a_F_gin_check_parent_keys_consistency_1), int32(0), int32(_a_F_gin_check_parent_keys_consistency_2), int32(_a_F_gin_check_parent_keys_consistency_3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v51 = int32(_a_F_gin_check_parent_keys_consistency_4)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v49
	F_initGinState(m, v38+int32(564), l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v60 = F_palloc0(m, int32(20))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = int64(8589934591)
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = int64(0)
	v69 = v60
	v89 = int32(-1)
	goto L17
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L1
	} else {
		goto L387
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L1
	} else {
		goto L378
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L1
	} else {
		goto L374
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L1
	} else {
		goto L370
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L1
	} else {
		goto L366
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L362
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L1
	} else {
		goto L358
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L1
	} else {
		goto L354
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L350
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L1
	} else {
		goto L346
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L1
	} else {
		goto L342
	}
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[1]))
	if v103 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v52
	F_MemoryContextDelete(m, v49)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L341
	}
L19:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v109 = F_ReadBufferExtended(m, l0, v106, v107, v106, v41)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	F_LockBuffer(m, v109, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v114 = int32(0)
	v115 = base.B2i32(v114 <= v109)
	if v115 == v114 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+12)))
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v135)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v115 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+(v109^int32(-1))<<(uint(int32(2))%32))))
	v133 = v125
	goto L25
L27:
	;
	goto L28
L28:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v133 = v127 + v109<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+14)))
	if v157 == int32(0) {
		goto L6
	} else {
		goto L33
	}
L30:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+(v109^int32(-1))<<(uint(int32(2))%32))))
	v156 = v148
	goto L29
L31:
	;
	goto L32
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v156 = v150 + v109<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L33:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+19)))
	v161 = int32(8)
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	if (v160<<(uint(v161)%32)-v163)&int32(_a_F_gin_check_parent_keys_consistency_5) != v161 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156+v163)+6)))
	if v170&int32(4) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v134) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	if v170&int32(2) == int32(0) {
		goto L8
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+12)))
	if base.Ui32(v209) < base.Ui32(int32(25)) {
		goto L35
	} else {
		goto L45
	}
L39:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+12)))
	if base.B2i32(base.Ui32(v177) < base.Ui32(int32(25)))|base.B2i32((v177+int32(_a_F_gin_check_parent_keys_consistency_6))&int32(_a_F_gin_check_parent_keys_consistency_7) == int32(0)) != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+532)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v38)+528)) = v194 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_8), v38+int32(528))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(704), int32(_a_F_gin_check_parent_keys_consistency_10))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	if base.Ui32(int32(409)) <= base.Ui32(int32(base.Ui32(v209+int32(_a_F_gin_check_parent_keys_consistency_6))>>(uint(int32(2))%32))&int32(_a_F_gin_check_parent_keys_consistency_5)) {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	goto L35
L47:
	;
	v228 = int32(base.Ui32(v134+int32(_a_F_gin_check_parent_keys_consistency_6)) >> (uint(int32(2)) % 32))
	goto L49
L48:
	;
	v228 = int32(0)
	goto L49
L49:
	;
	v231 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v231 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+500)) = v228 & int32(_a_F_gin_check_parent_keys_consistency_5)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+496)) = v233
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_11), v38+int32(496))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v249 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(444), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v322)+6)))
	if v324&int32(2) == int32(0) {
		v334 = v89
		goto L74
	} else {
		goto L75
	}
L57:
	;
	v253 = v38 + int32(564)
	v256 = F_gintuple_get_key(m, v253, v249, v38+int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v259 = F_gintuple_get_attrnum(m, v253, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v264 = F_PageGetItemIdCareful_1(m, l0, v261, v133, v228&int32(_a_F_gin_check_parent_keys_consistency_5))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v269 = v133 + v266&int32(_a_F_gin_check_parent_keys_consistency_14)
	v270 = F_gintuple_get_attrnum(m, v253, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v274 = F_gintuple_get_key(m, v253, v269, v38+int32(_a_F_gin_check_parent_keys_consistency_15))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v137 == int32(-1) {
		goto L56
	} else {
		goto L63
	}
L63:
	;
	v278 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[4]))))
	v279 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	v280 = F_ginCompareAttEntries(m, v253, v270, v274, v278, v259, v256, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if int32(0) <= v280 {
		goto L56
	} else {
		goto L65
	}
L65:
	;
	v286 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if v286 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+480)) = base.I64_rotl(v288, int64(32))
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_16), v38+int32(480))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v303 = F_palloc(m, int32(20))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(474), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = v305
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v308 = F_CopyIndexTuple(m, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v308
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+12)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v303)+8)) = v311
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+16)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v303
	goto L56
L74:
	;
	v336 = v228 & int32(_a_F_gin_check_parent_keys_consistency_5)
	if v336 != 0 {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v89 == int32(-1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v334 = v329
	goto L74
L77:
	;
	goto L78
L78:
	;
	if v329 != v89 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	v334 = v89
	goto L74
L80:
	;
	v337 = int32(0)
	v343 = v337
	v350 = v337
	v357 = int32(1)
	goto L83
L81:
	;
	goto L82
L82:
	;
	F_LockBuffer(m, v109, int32(0))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L1
	} else {
		goto L333
	}
L83:
	;
	v376 = v38 + int32(564)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v379 = v357 & int32(_a_F_gin_check_parent_keys_consistency_5)
	v380 = F_PageGetItemIdCareful_1(m, l0, v377, v133, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	goto L82
L85:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v385 = v133 + v382&int32(_a_F_gin_check_parent_keys_consistency_14)
	v386 = F_gintuple_get_attrnum(m, v376, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v391 = int32(7)
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+6)))
	if (int32(base.Ui32(v388)>>(uint(int32(17))%32))+v391)&int32(_a_F_gin_check_parent_keys_consistency_17) == (v395&int32(_a_F_gin_check_parent_keys_consistency_18)+v391)&int32(_a_F_gin_check_parent_keys_consistency_19) {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v788 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)))
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v788)+6)))
	if v790&int32(2) == int32(0) {
		goto L152
	} else {
		goto L153
	}
L88:
	;
	v735 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L147
	}
L89:
	;
	F_UnlockReleaseBuffer(m, v464)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L146
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L142
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L138
	}
L92:
	;
	v405 = F_gintuple_get_key(m, v376, v385, v38+int32(563))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L134
	}
L95:
	;
	if v379 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.B2i32(v437 == int32(0))|base.B2i32(v379 != v336) != 0 {
		goto L87
	} else {
		goto L105
	}
L97:
	;
	if base.B2i32(v379 != v336)|base.B2i32(v137 != int32(-1)) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v415)+6)))
	if v417&int32(2) == int32(0) {
		goto L96
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v423 = v38 + int32(564)
	v428 = F_gintuple_get_key(m, v423, v343, v38+int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v430 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	v431 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+563)))
	v432 = F_ginCompareAttEntries(m, v423, v350&int32(_a_F_gin_check_parent_keys_consistency_5), v428, v430, v386, v405, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if int32(0) <= v432 {
		goto L91
	} else {
		goto L104
	}
L104:
	;
	goto L96
L105:
	;
	v443 = v38 + int32(564)
	v444 = F_gintuple_get_attrnum(m, v443, v437)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v449 = F_gintuple_get_key(m, v443, v446, v38+int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v451 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+563)))
	v452 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	v453 = F_ginCompareAttEntries(m, v443, v386, v405, v451, v444, v449, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v453 <= int32(0) {
		goto L87
	} else {
		goto L109
	}
L109:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	F_pfree(m, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v461 = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v464 = F_ReadBufferExtended(m, l0, v461, v462, v461, v41)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_LockBuffer(m, v464, int32(1))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v464 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v486)+16)))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487+v486)+6)))
	if v489&int32(2) != 0 {
		goto L89
	} else {
		goto L117
	}
L114:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v472+(v464^int32(-1))<<(uint(int32(2))%32))))
	v486 = v478
	goto L113
L115:
	;
	goto L116
L116:
	;
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v486 = v480 + v464<<(uint(int32(13))%32) + int32(-8192)
	goto L113
L117:
	;
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v486)+12)))
	if base.Ui32(v492) < base.Ui32(int32(25)) {
		goto L89
	} else {
		goto L118
	}
L118:
	;
	v501 = int32(base.Ui32(v492+int32(_a_F_gin_check_parent_keys_consistency_6))>>(uint(int32(2))%32)) & int32(_a_F_gin_check_parent_keys_consistency_5)
	if v501 == int32(0) {
		goto L89
	} else {
		goto L119
	}
L119:
	;
	v505 = int32(1)
	goto L120
L120:
	;
	v541 = F_PageGetItemIdCareful_1(m, l0, v462, v486, v505&int32(_a_F_gin_check_parent_keys_consistency_5))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	v558 = F_CopyIndexTuple(m, v546)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L127
	}
L122:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	v546 = v486 + v543&int32(_a_F_gin_check_parent_keys_consistency_14)
	v547 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v546))))
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v546)+2)))
	if v460 != v547<<(uint(int32(16))%32)|v550 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v554 = v505 + int32(1)
	if base.Ui32(v554&int32(_a_F_gin_check_parent_keys_consistency_5)) <= base.Ui32(v501) {
		v505 = v554
		goto L120
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L121
L126:
	;
	goto L89
L127:
	;
	F_UnlockReleaseBuffer(m, v464)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v558
	if v558 == int32(0) {
		goto L88
	} else {
		goto L129
	}
L129:
	;
	v566 = v38 + int32(564)
	v567 = F_gintuple_get_attrnum(m, v566, v558)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v572 = F_gintuple_get_key(m, v566, v569, v38+int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v574 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+563)))
	v575 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	v576 = F_ginCompareAttEntries(m, v566, v386, v405, v574, v567, v572, v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if int32(0) < v576 {
		goto L90
	} else {
		goto L133
	}
L133:
	;
	goto L87
L134:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+456)) = v357 & int32(_a_F_gin_check_parent_keys_consistency_5)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+452)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v38)+448)) = v587 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_20), v38+int32(448))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(516), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+444)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v38)+440)) = v357 & int32(_a_F_gin_check_parent_keys_consistency_5)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+436)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v38)+432)) = v613 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_21), v38+int32(432))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(542), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+424)) = v228 & int32(_a_F_gin_check_parent_keys_consistency_5)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+420)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v38)+416)) = v640 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_22), v38+int32(416))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(594), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = int32(0)
	goto L88
L147:
	;
	if v735 == int32(0) {
		goto L87
	} else {
		goto L148
	}
L148:
	;
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+400)) = base.I64_rotl(v739, int64(32))
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_23), v38+int32(400))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(576), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L87
L151:
	;
	v1697 = F_CopyIndexTuple(m, v385)
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L1
	} else {
		goto L331
	}
L152:
	;
	v796 = F_palloc(m, int32(20))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385))))
	v823 = v821 << (uint(int32(16)) % 32)
	v824 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+2)))
	v825 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+4)))
	if v825 == int32(_a_F_gin_check_parent_keys_consistency_5) {
		goto L162
	} else {
		goto L163
	}
L155:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v796))) = v798 + int32(1)
	if v379 == v336 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v796)+4)) = v808
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v796)+8)) = v810
	v812 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+2)))
	v813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385))))
	*(*int32)(unsafe.Add(mBase, uint32(v796)+12)) = v812 | v813<<(uint(int32(16))%32)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v796)+16)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v796
	goto L151
L157:
	;
	if v137 == int32(-1) {
		v808 = int32(0)
		goto L156
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v806 = F_CopyIndexTuple(m, v385)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	v808 = v806
	goto L156
L162:
	;
	v829 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v1538 = v385 + v823&int32(2147418112) + v824
	if v823 < int32(0) {
		goto L309
	} else {
		goto L310
	}
L165:
	;
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	v837 = F_AllocSetContextCreateInternal(m, v832, int32(_a_F_gin_check_parent_keys_consistency_24), int32(0), int32(_a_F_gin_check_parent_keys_consistency_2), int32(_a_F_gin_check_parent_keys_consistency_3))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v839 = int32(_a_F_gin_check_parent_keys_consistency_4)
	v840 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v837
	v844 = F_palloc0(m, int32(24))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v846 = v823 | v824
	*(*int32)(unsafe.Add(mBase, uint32(v844)+16)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v844)+12)) = int32(-1)
	v850 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v844)+8)) = uint16(v850)
	*(*int64)(unsafe.Add(mBase, uint32(v844))) = int64(-4294967296)
	v856 = F_errstart(m, int32(12), v850)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	if v856 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+352)) = v846
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_25), v38+int32(352))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v871 = int32(-1)
	v873 = v844
	goto L174
L172:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(161), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[1]))
	if v906 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[0])) = v840
	F_MemoryContextDelete(m, v837)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L306
	}
L176:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v909 = int32(0)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	v912 = F_ReadBufferExtended(m, l0, v909, v910, v909, v829)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L180
	}
L179:
	;
	goto L178
L180:
	;
	F_LockBuffer(m, v912, int32(1))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	if v912 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	F_LockBuffer(m, v912, int32(0))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L302
	}
L183:
	;
	v935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+16)))
	v936 = v935 + v934
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v936)+6)))
	if v937&int32(2) != 0 {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	v920 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[2]))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v920+(v912^int32(-1))<<(uint(int32(2))%32))))
	v934 = v926
	goto L183
L185:
	;
	goto L186
L186:
	;
	v928 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[3]))
	v934 = v928 + v912<<(uint(int32(13))%32) + int32(-8192)
	goto L183
L187:
	;
	v940 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[6]))) = uint16(v940)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[4]))) = v940
	v946 = F_errstart(m, int32(14), v940)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v936)))
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v936)+4)))
	v1150 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L232
	}
L190:
	;
	if v946 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+336)) = v948
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_27), v38+int32(336))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	if v871 == int32(-1) {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(191), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+316)) = uint16(v965)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+312)) = v967
	v973 = F_GinDataLeafPageGetItems(m, v934, v38+int32(_a_F_gin_check_parent_keys_consistency_28), v38+int32(312))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L201
	}
L197:
	;
	v964 = v960
	goto L196
L198:
	;
	goto L199
L199:
	;
	if v871 != v960 {
		goto L15
	} else {
		goto L200
	}
L200:
	;
	v964 = v871
	goto L196
L201:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[7])))
	if int32(0) < v975 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	v1028 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L208
	}
L203:
	;
	v978 = int32(6)
	v980 = v973 + v975*v978
	v983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v980-int32(4)))))
	v986 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v980-v978))))
	v987 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v973)+2)))
	v988 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v973))))
	v989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v973)+4)))
	v992 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v980-int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+304)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v38)+296)) = v989
	v995 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+292)) = v987 | v988<<(uint(v995)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+288)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+300)) = v983 | v986<<(uint(v995)%32)
	v1010 = F_pg_snprintf(m, v38+int32(_a_F_gin_check_parent_keys_consistency_13), int32(1024), int32(_a_F_gin_check_parent_keys_consistency_29), v38+int32(288))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1017 = F_pg_snprintf(m, v38+int32(_a_F_gin_check_parent_keys_consistency_13), int32(1024), int32(_a_F_gin_check_parent_keys_consistency_30), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L207
	}
L206:
	;
	goto L202
L207:
	;
	goto L202
L208:
	;
	if v1025 != int32(-1) {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	if v1081 == int32(-1) {
		v1489 = v964
		goto L182
	} else {
		goto L219
	}
L210:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), v1073, int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L218
	}
L211:
	;
	if v1028 == int32(0) {
		goto L209
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	if v1028 == int32(0) {
		goto L209
	} else {
		goto L216
	}
L214:
	;
	v1035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+6)))
	v1036 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+4)))
	v1037 = *(*int64)(unsafe.Add(mBase, uint32(v873)+12))
	v1038 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+272)) = v38 + int32(_a_F_gin_check_parent_keys_consistency_13)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+268)) = v1038
	*(*int64)(unsafe.Add(mBase, uint32(v38)+256)) = base.I64_rotl(v1037, int64(32))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+264)) = v1035 | v1036<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_31), v38+int32(256))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v1073 = int32(219)
	goto L210
L216:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+240)) = v1058
	*(*int32)(unsafe.Add(mBase, uint32(v38)+244)) = v38 + int32(_a_F_gin_check_parent_keys_consistency_13)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_32), v38+int32(240))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v1073 = int32(223)
	goto L210
L218:
	;
	goto L209
L219:
	;
	v1084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+8)))
	if v1084 == int32(0) {
		v1489 = v964
		goto L182
	} else {
		goto L220
	}
L220:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[7])))
	if v1087 <= int32(0) {
		v1489 = v964
		goto L182
	} else {
		goto L221
	}
L221:
	;
	v1091 = v873 + int32(4)
	v1092 = int32(6)
	v1096 = v973 + v1087*v1092 - v1092
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+2)))
	v1101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091))))
	v1102 = int32(16)
	v1104 = v1100 | v1101<<(uint(v1102)%32)
	v1105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1096)+2)))
	v1106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1096))))
	v1109 = v1105 | v1106<<(uint(v1102)%32)
	if base.Ui32(v1104) < base.Ui32(v1109) {
		v1120 = int32(-1)
		goto L223
	} else {
		goto L224
	}
L222:
	;
	if int32(0) <= v1120 {
		v1489 = v964
		goto L182
	} else {
		goto L227
	}
L223:
	;
	goto L222
L224:
	;
	if base.Ui32(v1109) < base.Ui32(v1104) {
		v1120 = int32(1)
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+4)))
	v1115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1096)+4)))
	if base.Ui32(v1114) < base.Ui32(v1115) {
		v1120 = int32(-1)
		goto L223
	} else {
		goto L226
	}
L226:
	;
	v1120 = base.B2i32(base.Ui32(v1115) < base.Ui32(v1114))
	goto L223
L227:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+228)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v38)+224)) = v1130 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_33), v38+int32(224))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(231), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	if v1150 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+212)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v38)+208)) = v1152
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_34), v38+int32(208))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	v1169 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(246), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	if v1166 != int32(-1) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v1218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+12)))
	v1222 = base.I32_div_u_s(v1218-int32(32), int32(10))
	if v1222 != v1147 {
		goto L14
	} else {
		goto L249
	}
L240:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), v1210, int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L248
	}
L241:
	;
	if v1169 == int32(0) {
		goto L239
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	if v1169 == int32(0) {
		goto L239
	} else {
		goto L246
	}
L244:
	;
	v1176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+6)))
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+4)))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	v1180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+192)) = v1180
	*(*int32)(unsafe.Add(mBase, uint32(v38)+184)) = v1179
	*(*int32)(unsafe.Add(mBase, uint32(v38)+180)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = v1178
	*(*int32)(unsafe.Add(mBase, uint32(v38)+188)) = v1176 | v1177<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_35), v38+int32(176))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v1210 = int32(252)
	goto L240
L246:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+164)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v38)+160)) = v1197
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_36), v38+int32(160))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1210 = int32(255)
	goto L240
L248:
	;
	goto L239
L249:
	;
	v1224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[8]))) = uint16(v1224)
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v934)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5]))) = v1226
	v1229 = v873 + int32(4)
	v1230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+8)))
	v1231 = int32(0)
	if base.B2i32(v1230 == v1231)|base.B2i32(v1146 == int32(-1)) == v1231 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1229)+2)))
	v1241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1229))))
	v1242 = int32(16)
	v1245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[9]))))
	v1246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	if v1240|v1241<<(uint(v1242)%32) == v1245|v1246<<(uint(v1242)%32) {
		goto L255
	} else {
		goto L256
	}
L251:
	;
	goto L252
L252:
	;
	if v1147 == int32(0) {
		v1489 = v871
		goto L182
	} else {
		goto L260
	}
L253:
	;
	if v1256 == int32(0) {
		goto L13
	} else {
		goto L259
	}
L254:
	;
	goto L253
L255:
	;
	v1252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1229)+4)))
	v1253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[8]))))
	if v1252 == v1253 {
		v1256 = int32(1)
		goto L254
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1256 = int32(0)
	goto L254
L258:
	;
	goto L257
L259:
	;
	goto L252
L260:
	;
	v1274 = int32(1)
	goto L261
L261:
	;
	v1302 = v1274 * int32(10)
	v1303 = v934 + int32(22) + v1302
	v1306 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L263
	}
L262:
	;
	v1489 = v871
	goto L182
L263:
	;
	v1308 = base.B2i32(v1147 != v1274)
	v1309 = int32(0)
	if base.B2i32(v1308 == v1309)&base.B2i32(v1146 == int32(-1)) == v1309 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	if v1147 != v1274 {
		goto L291
	} else {
		goto L292
	}
L265:
	;
	if v1306 != 0 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L267
L267:
	;
	if v1306 != 0 {
		goto L284
	} else {
		goto L285
	}
L268:
	;
	v1316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+6)))
	v1317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+4)))
	v1318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+2)))
	v1319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303))))
	v1320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v1320
	v1322 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v1318 | v1319<<(uint(v1322)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v1316 | v1317<<(uint(v1322)%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_37), v38+int32(96))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	if v1274 == int32(1) {
		goto L264
	} else {
		goto L273
	}
L271:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(312), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	v1347 = v1303 + int32(4)
	v1350 = v934 + int32(24) + v1302 - int32(8)
	v1354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1347)+2)))
	v1355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1347))))
	v1356 = int32(16)
	v1358 = v1354 | v1355<<(uint(v1356)%32)
	v1359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1350)+2)))
	v1360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1350))))
	v1363 = v1359 | v1360<<(uint(v1356)%32)
	if base.Ui32(v1358) < base.Ui32(v1363) {
		v1374 = int32(-1)
		goto L275
	} else {
		goto L276
	}
L274:
	;
	if int32(0) <= v1374 {
		goto L264
	} else {
		goto L279
	}
L275:
	;
	goto L274
L276:
	;
	if base.Ui32(v1363) < base.Ui32(v1358) {
		v1374 = int32(1)
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1347)+4)))
	v1369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1350)+4)))
	if base.Ui32(v1368) < base.Ui32(v1369) {
		v1374 = int32(-1)
		goto L275
	} else {
		goto L278
	}
L278:
	;
	v1374 = base.B2i32(base.Ui32(v1369) < base.Ui32(v1368))
	goto L275
L279:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v1274
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v1385
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v1384 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_38), v38+int32(80))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(341), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+2)))
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v1401 | v1402<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_37), v38-int32(-64))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L1
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+6)))
	v1422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+4)))
	if v1421|v1422 != 0 {
		goto L11
	} else {
		goto L289
	}
L287:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(315), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	goto L286
L289:
	;
	v1424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+8)))
	if v1424 != 0 {
		goto L11
	} else {
		goto L290
	}
L290:
	;
	goto L264
L291:
	;
	v1462 = F_palloc(m, int32(24))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L300
	}
L292:
	;
	v1430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+8)))
	if v1430 == int32(0) {
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v1434 = v1303 + int32(4)
	v1438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1229)+2)))
	v1439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1229))))
	v1440 = int32(16)
	v1442 = v1438 | v1439<<(uint(v1440)%32)
	v1443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1434)+2)))
	v1444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1434))))
	v1447 = v1443 | v1444<<(uint(v1440)%32)
	if base.Ui32(v1442) < base.Ui32(v1447) {
		v1458 = int32(-1)
		goto L295
	} else {
		goto L296
	}
L294:
	;
	if v1458 < int32(0) {
		goto L12
	} else {
		goto L299
	}
L295:
	;
	goto L294
L296:
	;
	if base.Ui32(v1447) < base.Ui32(v1442) {
		v1458 = int32(1)
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1229)+4)))
	v1453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1434)+4)))
	if base.Ui32(v1452) < base.Ui32(v1453) {
		v1458 = int32(-1)
		goto L295
	} else {
		goto L298
	}
L298:
	;
	v1458 = base.B2i32(base.Ui32(v1453) < base.Ui32(v1452))
	goto L295
L299:
	;
	goto L291
L300:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	v1465 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1462))) = v1464 + v1465
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1462)+4)) = v1468
	v1470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1462)+8)) = uint16(v1470)
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1462)+12)) = v1472
	v1474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+2)))
	v1475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303))))
	*(*int32)(unsafe.Add(mBase, uint32(v1462)+16)) = v1474 | v1475<<(uint(int32(16))%32)
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v873)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1462)+20)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v873)+20)) = v1462
	v1486 = (v1274 + v1465) & int32(_a_F_gin_check_parent_keys_consistency_5)
	if base.Ui32(v1486) <= base.Ui32(v1147) {
		v1274 = v1486
		goto L261
	} else {
		goto L301
	}
L301:
	;
	goto L262
L302:
	;
	F_ReleaseBuffer(m, v912)
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v873)+20))
	F_pfree(m, v873)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	if v1528 != 0 {
		v871 = v1489
		v873 = v1528
		goto L174
	} else {
		goto L305
	}
L305:
	;
	goto L175
L306:
	;
	goto L151
L307:
	;
	F_pfree(m, v1628)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L1
	} else {
		goto L330
	}
L308:
	;
	v1578 = int32(0)
	goto L326
L309:
	;
	if v825 != 0 {
		goto L312
	} else {
		goto L313
	}
L310:
	;
	goto L311
L311:
	;
	v1568 = v825 * int32(6)
	v1569 = F_palloc(m, v1568)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L1
	} else {
		goto L321
	}
L312:
	;
	v1543 = F_ginPostingListDecode(m, v1538, v38+int32(_a_F_gin_check_parent_keys_consistency_13))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L1
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v1565 = F_palloc(m, int32(0))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L1
	} else {
		goto L320
	}
L315:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5])))
	if v1545 == v825 {
		v1575 = v1543
		goto L308
	} else {
		goto L316
	}
L316:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+384)) = v825
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+388)) = v1552
	F_errmsg_internal(m, int32(_a_F_gin_check_parent_keys_consistency_39), v38+int32(384))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(113), int32(_a_F_gin_check_parent_keys_consistency_40))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	v1628 = v1565
	goto L307
L321:
	;
	if v1568 != 0 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	base.MemoryCopy(m, v1569, v1538, v1568)
	goto L324
L323:
	;
	goto L324
L324:
	;
	if v825 == int32(0) {
		v1628 = v1569
		goto L307
	} else {
		goto L325
	}
L325:
	;
	v1575 = v1569
	goto L308
L326:
	;
	v1615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1575+v1578*int32(6))+4)))
	if base.Ui32(int32(2048)) <= base.Ui32((v1615-int32(1))&int32(_a_F_gin_check_parent_keys_consistency_5)) {
		goto L16
	} else {
		goto L328
	}
L327:
	;
	v1628 = v1575
	goto L307
L328:
	;
	v1623 = v1578 + int32(1)
	if v1623 != v825 {
		v1578 = v1623
		goto L326
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	goto L151
L331:
	;
	v1700 = v357 + int32(1)
	if base.Ui32(v1700&int32(_a_F_gin_check_parent_keys_consistency_5)) <= base.Ui32(v336) {
		v343 = v1697
		v350 = v386
		v357 = v1700
		goto L83
	} else {
		goto L332
	}
L332:
	;
	goto L84
L333:
	;
	F_ReleaseBuffer(m, v109)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v1745 != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	F_pfree(m, v1745)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L1
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	F_pfree(m, v69)
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L1
	} else {
		goto L339
	}
L338:
	;
	goto L337
L339:
	;
	if v1744 != 0 {
		v69 = v1744
		v89 = v334
		goto L17
	} else {
		goto L340
	}
L340:
	;
	goto L18
L341:
	;
	m.G0 = v38 + int32(_a_F_gin_check_parent_keys_consistency_0)
	return
L342:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+372)) = v1765
	*(*int32)(unsafe.Add(mBase, uint32(v38)+368)) = v1764 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_41), v38+int32(368))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(636), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+324)) = v1788
	*(*int32)(unsafe.Add(mBase, uint32(v38)+320)) = v1787 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_42), v38+int32(320))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(199), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L350:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+156)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v38)+152)) = v1811
	*(*int32)(unsafe.Add(mBase, uint32(v38)+148)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v38)+144)) = v1810 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_43), v38+int32(144))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(274), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	v1835 = *(*int64)(unsafe.Add(mBase, uint32(v873)+12))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+6)))
	v1838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+4)))
	v1839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+136)) = v1839
	v1841 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+132)) = v1837 | v1838<<(uint(v1841)%32)
	v1845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[8]))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+120)) = v1845
	*(*int32)(unsafe.Add(mBase, uint32(v38)+112)) = v1836 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+124)) = base.I64_rotl(v1835, int64(32))
	v1853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[9]))))
	v1854 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_gin_check_parent_keys_consistency[5]))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+116)) = v1853 | v1854<<(uint(v1841)%32)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_44), v38+int32(112))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(299), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L358:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+40)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v1877
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v1876 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_45), v38+int32(32))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(354), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L362:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v1900 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+6)))
	v1901 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+4)))
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	v1904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1303)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v1904
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v1903
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v1902 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v1900 | v1901<<(uint(int32(16))%32)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_46), v38+int32(48))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(331), int32(_a_F_gin_check_parent_keys_consistency_26))
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L366:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+468)) = v1932
	*(*int32)(unsafe.Add(mBase, uint32(v38)+464)) = v1931 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_42), v38+int32(464))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(495), int32(_a_F_gin_check_parent_keys_consistency_12))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v1954 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_47), v38+int32(16))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(710), int32(_a_F_gin_check_parent_keys_consistency_10))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+516)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v38)+512)) = v1976 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_48), v38+int32(512))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(699), int32(_a_F_gin_check_parent_keys_consistency_10))
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L378:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v109 < int32(0) {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+548)) = v2017
	*(*int32)(unsafe.Add(mBase, uint32(v38)+544)) = v1998 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_49), v38+int32(544))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L1
	} else {
		goto L384
	}
L381:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[10]))
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2002+(v109^int32(-1))<<(uint(int32(6))%32))+16))
	v2017 = v2008
	goto L380
L382:
	;
	goto L383
L383:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[11]))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2010+v109<<(uint(int32(6))%32)+int32(-64))+16))
	v2017 = v2016
	goto L380
L384:
	;
	F_errhint(m, int32(_a_F_gin_check_parent_keys_consistency_50), int32(0))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(691), int32(_a_F_gin_check_parent_keys_consistency_10))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L387:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v109 < int32(0) {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v2062
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v2043 + int32(4)
	F_errmsg(m, int32(_a_F_gin_check_parent_keys_consistency_51), v38)
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L1
	} else {
		goto L393
	}
L390:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[10]))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2047+(v109^int32(-1))<<(uint(int32(6))%32))+16))
	v2062 = v2053
	goto L389
L391:
	;
	goto L392
L392:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, _c_F_gin_check_parent_keys_consistency[11]))
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v2055+v109<<(uint(int32(6))%32)+int32(-64))+16))
	v2062 = v2061
	goto L389
L393:
	;
	F_errhint(m, int32(_a_F_gin_check_parent_keys_consistency_50), int32(0))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	F_errfinish(m, int32(_a_F_gin_check_parent_keys_consistency_9), int32(680), int32(_a_F_gin_check_parent_keys_consistency_10))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_consistent_jsonb_path(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v15 == int32(7) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L17
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return v57
L3:
	;
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v18)
	v21 = int32(0)
	if v13 <= v21 {
		v57 = v18
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v37 = int32(1)
	if base.Ui32(v37) < base.Ui32((v15-int32(15))&int32(_a_F_gin_consistent_jsonb_path_0)) {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v24 = v21
	goto L7
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v14))))
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v57 = int32(0)
	goto L2
L9:
	;
	v34 = v24 + int32(1)
	if v13 != v34 {
		v24 = v34
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v57 = v18
	goto L2
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v45)
	if v13 <= int32(0) {
		v57 = v37
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v50 = F_execute_jsp_gin_node(m, v49, v14)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v57 = base.B2i32(v50 != int32(0))
	goto L2
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
	F_errmsg_internal(m, int32(_a_F_gin_consistent_jsonb_path_1), v10)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_gin_consistent_jsonb_path_2), int32(1266), int32(_a_F_gin_consistent_jsonb_path_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_jsonb_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
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
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
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
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
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
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v795 int32
	_ = v795
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	switch v19 - int32(7) {
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
	v808 = m.ExcPending
	if v808 != 0 {
		goto L6
	} else {
		goto L155
	}
L2:
	;
	m.G0 = v15 - int32(-64)
	return v795
L3:
	;
	if v19&int32(_a_F_gin_extract_jsonb_query_0) == int32(10) {
		goto L73
	} else {
		goto L74
	}
L4:
	;
	v32 = int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = F_pg_detoast_datum_packed(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_DirectFunctionCall2Coll(m, int32(1320), int32(0), v24, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v29 != 0 {
		v795 = v25
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v795 = v25
	goto L2
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(1)
	v39 = F_palloc(m, int32(4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v41 = int32(1)
	v42 = v34 + v41
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v47 = v45 & v41
	if v47 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v48 = v42
	goto L13
L12:
	;
	v48 = v34 + int32(4)
	goto L13
L13:
	;
	if v45 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v356 = v352 + int32(5)
	v357 = F_palloc(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L6
	} else {
		goto L69
	}
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v54 == int32(18) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v47 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v57 = int32(16)
	goto L20
L19:
	;
	v57 = int32(0)
	goto L20
L20:
	;
	if base.Ui32((v54-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v64 = int32(4)
	goto L23
L22:
	;
	v64 = v57
	goto L23
L23:
	;
	v352 = v64
	v353 = v32
	v354 = v48
	goto L14
L24:
	;
	v65 = int32(1)
	v74 = int32(base.Ui32(v45)>>(uint(v65)%32)) - v65
	goto L26
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v74 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L26:
	;
	if v74 < int32(126) {
		v352 = v74
		v353 = v32
		v354 = v48
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v82 = v74 - int32(1636608432)
	if v48&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v336 ^ v328 - base.I32_rotl(v336, int32(24))
	v343 = v13 + int32(-10)
	v348 = F_pg_snprintf(m, v343, int32(10), int32(_a_F_gin_extract_jsonb_query_1), v13+int32(-32))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L68
	}
L29:
	;
	v314 = int32(14)
	v316 = v310 ^ v311 - base.I32_rotl(v310, v314)
	v320 = v316 ^ v309 - base.I32_rotl(v316, int32(11))
	v324 = v320 ^ v310 - base.I32_rotl(v320, int32(25))
	v328 = v324 ^ v316 - base.I32_rotl(v324, int32(16))
	v332 = v328 ^ v320 - base.I32_rotl(v328, int32(4))
	v336 = v332 ^ v324 - base.I32_rotl(v332, v314)
	goto L28
L30:
	;
	switch v240 - int32(1) {
	case 0:
		v302 = v241
		v303 = v242
		v304 = v243
		goto L57
	case 1:
		v295 = v241
		v296 = v242
		v297 = v243
		goto L58
	case 2:
		v288 = v241
		v289 = v242
		v290 = v243
		goto L59
	case 3:
		v282 = v242
		v283 = v243
		goto L60
	case 4:
		v278 = v242
		v279 = v243
		goto L61
	case 5:
		v272 = v242
		v273 = v243
		goto L62
	case 6:
		v266 = v242
		v267 = v243
		goto L63
	case 7:
		v261 = v243
		goto L64
	case 8:
		v256 = v243
		goto L65
	case 9:
		v251 = v243
		goto L66
	case 10:
		goto L67
	default:
		v309 = v241
		v310 = v242
		v311 = v243
		goto L29
	}
L31:
	;
	v191 = v48
	v192 = v74
	v193 = v82
	v194 = v82
	v195 = v82
	goto L54
L32:
	;
	if base.Ui32(int32(11)) < base.Ui32(v74) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(v74) < base.Ui32(int32(12)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v239 = v48
	v240 = v74
	v241 = v82
	v242 = v82
	v243 = v82
	goto L30
L36:
	;
	switch v138 - int32(1) {
	case 0:
		v188 = v139
		goto L43
	case 1:
		v183 = v139
		goto L44
	case 2:
		goto L45
	case 3:
		v176 = v140
		goto L46
	case 4:
		v173 = v140
		goto L47
	case 5:
		v168 = v140
		goto L48
	case 6:
		goto L49
	case 7:
		v159 = v141
		goto L50
	case 8:
		v154 = v141
		goto L51
	case 9:
		v149 = v141
		goto L52
	case 10:
		goto L53
	default:
		v309 = v139
		v310 = v140
		v311 = v141
		goto L29
	}
L37:
	;
	v137 = v48
	v138 = v74
	v139 = v82
	v140 = v82
	v141 = v82
	goto L36
L38:
	;
	goto L39
L39:
	;
	v89 = v48
	v90 = v74
	v91 = v82
	v92 = v82
	v93 = v82
	goto L40
L40:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v96 = v95 + v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v100 = v99 + v93
	v102 = int32(4)
	v104 = v97 + v91 - v100 ^ base.I32_rotl(v100, v102)
	v108 = v96 - v104 ^ base.I32_rotl(v104, int32(6))
	v109 = v100 + v96
	v110 = v104 + v109
	v111 = v108 + v110
	v115 = v109 - v108 ^ base.I32_rotl(v108, int32(8))
	v119 = v110 - v115 ^ base.I32_rotl(v115, int32(16))
	v123 = v111 - v119 ^ base.I32_rotl(v119, int32(19))
	v124 = v115 + v111
	v125 = v119 + v124
	v126 = v123 + v125
	v130 = v124 - v123 ^ base.I32_rotl(v123, v102)
	v131 = int32(12)
	v132 = v89 + v131
	v134 = v90 - v131
	if base.Ui32(int32(11)) < base.Ui32(v134) {
		v89 = v132
		v90 = v134
		v91 = v125
		v92 = v126
		v93 = v130
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v137 = v132
	v138 = v134
	v139 = v125
	v140 = v126
	v141 = v130
	goto L36
L42:
	;
	goto L41
L43:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v309 = v188 + v189
	v310 = v140
	v311 = v141
	goto L29
L44:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v188 = v184<<(uint(int32(8))%32) + v183
	goto L43
L45:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+2)))
	v183 = v179<<(uint(int32(16))%32) + v139
	goto L44
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v309 = v177 + v139
	v310 = v176
	v311 = v141
	goto L29
L47:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+4)))
	v176 = v173 + v174
	goto L46
L48:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+5)))
	v173 = v169<<(uint(int32(8))%32) + v168
	goto L47
L49:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+6)))
	v168 = v164<<(uint(int32(16))%32) + v140
	goto L48
L50:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v309 = v160 + v139
	v310 = v162 + v140
	v311 = v159
	goto L29
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+8)))
	v159 = v155<<(uint(int32(8))%32) + v154
	goto L50
L52:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+9)))
	v154 = v150<<(uint(int32(16))%32) + v149
	goto L51
L53:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+10)))
	v149 = v145<<(uint(int32(24))%32) + v141
	goto L52
L54:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v198 = v197 + v194
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	v202 = v201 + v195
	v204 = int32(4)
	v206 = v199 + v193 - v202 ^ base.I32_rotl(v202, v204)
	v210 = v198 - v206 ^ base.I32_rotl(v206, int32(6))
	v211 = v202 + v198
	v212 = v206 + v211
	v213 = v210 + v212
	v217 = v211 - v210 ^ base.I32_rotl(v210, int32(8))
	v221 = v212 - v217 ^ base.I32_rotl(v217, int32(16))
	v225 = v213 - v221 ^ base.I32_rotl(v221, int32(19))
	v226 = v217 + v213
	v227 = v221 + v226
	v228 = v225 + v227
	v232 = v226 - v225 ^ base.I32_rotl(v225, v204)
	v233 = int32(12)
	v234 = v191 + v233
	v236 = v192 - v233
	if base.Ui32(int32(11)) < base.Ui32(v236) {
		v191 = v234
		v192 = v236
		v193 = v227
		v194 = v228
		v195 = v232
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v239 = v234
	v240 = v236
	v241 = v227
	v242 = v228
	v243 = v232
	goto L30
L56:
	;
	goto L55
L57:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v309 = v302 + v305
	v310 = v303
	v311 = v304
	goto L29
L58:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	v302 = v298<<(uint(int32(8))%32) + v295
	v303 = v296
	v304 = v297
	goto L57
L59:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+2)))
	v295 = v291<<(uint(int32(16))%32) + v288
	v296 = v289
	v297 = v290
	goto L58
L60:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+3)))
	v288 = v284<<(uint(int32(24))%32) + v241
	v289 = v282
	v290 = v283
	goto L59
L61:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+4)))
	v282 = v278 + v280
	v283 = v279
	goto L60
L62:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+5)))
	v278 = v274<<(uint(int32(8))%32) + v272
	v279 = v273
	goto L61
L63:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+6)))
	v272 = v268<<(uint(int32(16))%32) + v266
	v273 = v267
	goto L62
L64:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+7)))
	v266 = v262<<(uint(int32(24))%32) + v242
	v267 = v261
	goto L63
L65:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+8)))
	v261 = v257<<(uint(int32(8))%32) + v256
	goto L64
L66:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+9)))
	v256 = v252<<(uint(int32(16))%32) + v251
	goto L65
L67:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+10)))
	v251 = v247<<(uint(int32(24))%32) + v243
	goto L66
L68:
	;
	v352 = int32(8)
	v353 = int32(17)
	v354 = v343
	goto L14
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v357)+4)) = uint8(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v356 << (uint(int32(2)) % 32)
	if v352 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	base.MemoryCopy(m, v357+int32(5), v354, v352)
	goto L72
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v357
	v795 = v39
	goto L2
L73:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v372 = F_pg_detoast_datum(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(1)) < base.Ui32((v19-int32(15))&int32(_a_F_gin_extract_jsonb_query_2)) {
		goto L1
	} else {
		goto L151
	}
L76:
	;
	F_deconstruct_array_builtin(m, v372, int32(25), v13+int32(-16), v13+int32(-20), v13+int32(-24))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v386 = F_palloc(m, v383<<(uint(int32(2))%32))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if int32(0) < v388 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v392 = int32(0)
	v393 = v388
	v397 = v2
	goto L82
L80:
	;
	v760 = v2
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v760
	if base.B2i32(v19 != int32(11))|v760 != 0 {
		v795 = v386
		goto L2
	} else {
		goto L150
	}
L82:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v392))))
	if v406 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v760 = v747
	goto L81
L84:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v409+v392<<(uint(int32(2))%32))))
	v414 = int32(1)
	v415 = v413 + v414
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	v420 = v418 & v414
	if v420 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v746 = v393
	v747 = v397
	goto L86
L86:
	;
	v753 = v392 + int32(1)
	if v753 < v746 {
		v392 = v753
		v393 = v746
		v397 = v747
		goto L82
	} else {
		goto L149
	}
L87:
	;
	v421 = v415
	goto L89
L88:
	;
	v421 = v413 + int32(4)
	goto L89
L89:
	;
	if v418 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v729 = v725 + int32(5)
	v730 = F_palloc(m, v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L6
	} else {
		goto L145
	}
L91:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	if v427 == int32(18) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	if v420 != 0 {
		goto L100
	} else {
		goto L101
	}
L94:
	;
	v430 = int32(16)
	goto L96
L95:
	;
	v430 = int32(0)
	goto L96
L96:
	;
	if base.Ui32((v427-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v437 = int32(4)
	goto L99
L98:
	;
	v437 = v430
	goto L99
L99:
	;
	v725 = v437
	v726 = v421
	v727 = int32(1)
	goto L90
L100:
	;
	v440 = int32(1)
	v449 = int32(base.Ui32(v418)>>(uint(v440)%32)) - v440
	goto L102
L101:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v449 = int32(base.Ui32(v444)>>(uint(int32(2))%32)) - int32(4)
	goto L102
L102:
	;
	if v449 < int32(126) {
		v725 = v449
		v726 = v421
		v727 = int32(1)
		goto L90
	} else {
		goto L103
	}
L103:
	;
	v457 = v449 - int32(1636608432)
	if v421&int32(3) != 0 {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v711 ^ v703 - base.I32_rotl(v711, int32(24))
	v718 = v13 + int32(-10)
	v721 = F_pg_snprintf(m, v718, int32(10), int32(_a_F_gin_extract_jsonb_query_1), v15)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L6
	} else {
		goto L144
	}
L105:
	;
	v689 = int32(14)
	v691 = v685 ^ v686 - base.I32_rotl(v685, v689)
	v695 = v691 ^ v684 - base.I32_rotl(v691, int32(11))
	v699 = v695 ^ v685 - base.I32_rotl(v695, int32(25))
	v703 = v699 ^ v691 - base.I32_rotl(v699, int32(16))
	v707 = v703 ^ v695 - base.I32_rotl(v703, int32(4))
	v711 = v707 ^ v699 - base.I32_rotl(v707, v689)
	goto L104
L106:
	;
	switch v615 - int32(1) {
	case 0:
		v677 = v616
		v678 = v617
		v679 = v618
		goto L133
	case 1:
		v670 = v616
		v671 = v617
		v672 = v618
		goto L134
	case 2:
		v663 = v616
		v664 = v617
		v665 = v618
		goto L135
	case 3:
		v657 = v617
		v658 = v618
		goto L136
	case 4:
		v653 = v617
		v654 = v618
		goto L137
	case 5:
		v647 = v617
		v648 = v618
		goto L138
	case 6:
		v641 = v617
		v642 = v618
		goto L139
	case 7:
		v636 = v618
		goto L140
	case 8:
		v631 = v618
		goto L141
	case 9:
		v626 = v618
		goto L142
	case 10:
		goto L143
	default:
		v684 = v616
		v685 = v617
		v686 = v618
		goto L105
	}
L107:
	;
	v566 = v421
	v567 = v449
	v568 = v457
	v569 = v457
	v570 = v457
	goto L130
L108:
	;
	if base.Ui32(int32(11)) < base.Ui32(v449) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	if base.Ui32(v449) < base.Ui32(int32(12)) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v614 = v421
	v615 = v449
	v616 = v457
	v617 = v457
	v618 = v457
	goto L106
L112:
	;
	switch v513 - int32(1) {
	case 0:
		v563 = v514
		goto L119
	case 1:
		v558 = v514
		goto L120
	case 2:
		goto L121
	case 3:
		v551 = v515
		goto L122
	case 4:
		v548 = v515
		goto L123
	case 5:
		v543 = v515
		goto L124
	case 6:
		goto L125
	case 7:
		v534 = v516
		goto L126
	case 8:
		v529 = v516
		goto L127
	case 9:
		v524 = v516
		goto L128
	case 10:
		goto L129
	default:
		v684 = v514
		v685 = v515
		v686 = v516
		goto L105
	}
L113:
	;
	v512 = v421
	v513 = v449
	v514 = v457
	v515 = v457
	v516 = v457
	goto L112
L114:
	;
	goto L115
L115:
	;
	v464 = v421
	v465 = v449
	v466 = v457
	v467 = v457
	v468 = v457
	goto L116
L116:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	v471 = v470 + v467
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v464)+8))
	v475 = v474 + v468
	v477 = int32(4)
	v479 = v472 + v466 - v475 ^ base.I32_rotl(v475, v477)
	v483 = v471 - v479 ^ base.I32_rotl(v479, int32(6))
	v484 = v475 + v471
	v485 = v479 + v484
	v486 = v483 + v485
	v490 = v484 - v483 ^ base.I32_rotl(v483, int32(8))
	v494 = v485 - v490 ^ base.I32_rotl(v490, int32(16))
	v498 = v486 - v494 ^ base.I32_rotl(v494, int32(19))
	v499 = v490 + v486
	v500 = v494 + v499
	v501 = v498 + v500
	v505 = v499 - v498 ^ base.I32_rotl(v498, v477)
	v506 = int32(12)
	v507 = v464 + v506
	v509 = v465 - v506
	if base.Ui32(int32(11)) < base.Ui32(v509) {
		v464 = v507
		v465 = v509
		v466 = v500
		v467 = v501
		v468 = v505
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v512 = v507
	v513 = v509
	v514 = v500
	v515 = v501
	v516 = v505
	goto L112
L118:
	;
	goto L117
L119:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v684 = v563 + v564
	v685 = v515
	v686 = v516
	goto L105
L120:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+1)))
	v563 = v559<<(uint(int32(8))%32) + v558
	goto L119
L121:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+2)))
	v558 = v554<<(uint(int32(16))%32) + v514
	goto L120
L122:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v684 = v552 + v514
	v685 = v551
	v686 = v516
	goto L105
L123:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+4)))
	v551 = v548 + v549
	goto L122
L124:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+5)))
	v548 = v544<<(uint(int32(8))%32) + v543
	goto L123
L125:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+6)))
	v543 = v539<<(uint(int32(16))%32) + v515
	goto L124
L126:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v684 = v535 + v514
	v685 = v537 + v515
	v686 = v534
	goto L105
L127:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+8)))
	v534 = v530<<(uint(int32(8))%32) + v529
	goto L126
L128:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+9)))
	v529 = v525<<(uint(int32(16))%32) + v524
	goto L127
L129:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+10)))
	v524 = v520<<(uint(int32(24))%32) + v516
	goto L128
L130:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v573 = v572 + v569
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v577 = v576 + v570
	v579 = int32(4)
	v581 = v574 + v568 - v577 ^ base.I32_rotl(v577, v579)
	v585 = v573 - v581 ^ base.I32_rotl(v581, int32(6))
	v586 = v577 + v573
	v587 = v581 + v586
	v588 = v585 + v587
	v592 = v586 - v585 ^ base.I32_rotl(v585, int32(8))
	v596 = v587 - v592 ^ base.I32_rotl(v592, int32(16))
	v600 = v588 - v596 ^ base.I32_rotl(v596, int32(19))
	v601 = v592 + v588
	v602 = v596 + v601
	v603 = v600 + v602
	v607 = v601 - v600 ^ base.I32_rotl(v600, v579)
	v608 = int32(12)
	v609 = v566 + v608
	v611 = v567 - v608
	if base.Ui32(int32(11)) < base.Ui32(v611) {
		v566 = v609
		v567 = v611
		v568 = v602
		v569 = v603
		v570 = v607
		goto L130
	} else {
		goto L132
	}
L131:
	;
	v614 = v609
	v615 = v611
	v616 = v602
	v617 = v603
	v618 = v607
	goto L106
L132:
	;
	goto L131
L133:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v684 = v677 + v680
	v685 = v678
	v686 = v679
	goto L105
L134:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+1)))
	v677 = v673<<(uint(int32(8))%32) + v670
	v678 = v671
	v679 = v672
	goto L133
L135:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+2)))
	v670 = v666<<(uint(int32(16))%32) + v663
	v671 = v664
	v672 = v665
	goto L134
L136:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+3)))
	v663 = v659<<(uint(int32(24))%32) + v616
	v664 = v657
	v665 = v658
	goto L135
L137:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+4)))
	v657 = v653 + v655
	v658 = v654
	goto L136
L138:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+5)))
	v653 = v649<<(uint(int32(8))%32) + v647
	v654 = v648
	goto L137
L139:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+6)))
	v647 = v643<<(uint(int32(16))%32) + v641
	v648 = v642
	goto L138
L140:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+7)))
	v641 = v637<<(uint(int32(24))%32) + v617
	v642 = v636
	goto L139
L141:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+8)))
	v636 = v632<<(uint(int32(8))%32) + v631
	goto L140
L142:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+9)))
	v631 = v627<<(uint(int32(16))%32) + v626
	goto L141
L143:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+10)))
	v626 = v622<<(uint(int32(24))%32) + v618
	goto L142
L144:
	;
	v725 = int32(8)
	v726 = v718
	v727 = int32(17)
	goto L90
L145:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v730)+4)) = uint8(v727)
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v729 << (uint(int32(2)) % 32)
	if v725 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	base.MemoryCopy(m, v730+int32(5), v726, v725)
	goto L148
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386+v397<<(uint(int32(2))%32)))) = v730
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v746 = v745
	v747 = v397 + int32(1)
	goto L86
L149:
	;
	goto L83
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v795 = v386
	goto L2
L151:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v780 = F_pg_detoast_datum(m, v779)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v784 = F_extract_jsp_query(m, v780, v19, int32(0), v18, v783)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	if v784 != 0 {
		v795 = v784
		goto L2
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v795 = int32(0)
	goto L2
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
	F_errmsg_internal(m, int32(_a_F_gin_extract_jsonb_query_3), v13+int32(-48))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_gin_extract_jsonb_query_4), int32(921), int32(_a_F_gin_extract_jsonb_query_5))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
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
func F_gin_extract_query_varbit(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(_a_F_gin_extract_query_varbit_0), int32(2645))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gin_index_check(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_amcheck_lock_relation_and_check(m, v2, int32(2742), int32(_a_F_gin_index_check_0), int32(1), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
