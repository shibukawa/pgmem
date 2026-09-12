package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinBufferInit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = F_palloc0(m, int32(36))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(10922)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+10)))
	v28 = F_palloc0(m, v25*int32(36))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v28
	if int32(0) < v25 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v36 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v14 + int32(16)
	return v17
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v50 = v47 + v36*int32(36)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v36<<(uint(int32(2))%32))))
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+20)) = uint8(v59)
	v62 = v36 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v50)+10)) = uint16(v62)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+9)) = uint8(v59)
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v67 = v58
	goto L12
L11:
	;
	v67 = int32(100)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v67
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)))
	v81 = int32(4)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v71+v73*(base.I32_extend16_s(v62)-int32(1))<<(uint(int32(2))%32)+v81-v81)))
	goto L13
L13:
	;
	if v85 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v93 = v21 + int32(88) + v46<<(uint(int32(4))%32) + v36*int32(100)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v96 = F_lookup_type_cache(m, v94, int32(64))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v102 = v85
	goto L16
L16:
	;
	F_PrepareSortSupportComparisonShim(m, v102, v50)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+108))
	if v98 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v102 = v98
	goto L16
L19:
	;
	if v62 != v25 {
		v36 = v62
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
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v129 = F_format_type_be(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v129
	F_errmsg(m, int32(188076), v14)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(488017), int32(1311), int32(99602))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+8))
	v38 = F_index_form_tuple(m, v33, v14+int32(24), v14+int32(22))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+6)))
		v44 = v42 & int32(8191)
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
		v55 = int32(32768)
		*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v55)
		v60 = (v53 + int32(1)) & int32(16382)
		*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)) = uint16(v60)
		v66 = (v60 + l5 + int32(7)) & int32(-8)
		if base.Ui32(int32(2713)) <= base.Ui32(v66) {
			if l7 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(2712)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v66
						*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v129 + int32(4)
						F_errmsg(m, int32(672351), v14)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494283), int32(111), int32(381226))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
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
					v115 = int32(0)
					m.G0 = v14 + int32(32)
					return v115
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
					v77 = v75 & int32(8191)
					v82 = F__emscripten_memset_bulkmem(m, v73+v77, base.I32_extend8_s(int32(0)), v66-v77)
					mBase = m.M
					v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)))
					v86 = v83&int32(57344) | v66
					*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)) = uint16(v86)
					v88 = v73
					if l4 != 0 {
						v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+2)))
						v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
						if l5 != 0 {
							v98 = F__emscripten_memcpy_bulkmem(m, v88+(v90|v91<<(uint(int32(16))%32)&int32(2147418112)), l4, l5)
							mBase = m.M
						} else {
						}
					} else {
					}
					if v4 == int32(0) {
						v115 = v88
					} else {
						v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+6)))
						if int32(0) <= v104 {
							v107 = int32(8)
						} else {
							v107 = int32(16)
						}
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
						if v111 != 0 {
							v112 = int32(0)
						} else {
							v112 = int32(2)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v88+v107+v112))) = uint8(v4)
						v115 = v88
					}
					m.G0 = v14 + int32(32)
					return v115
				}
			} else {
				v88 = v38
				if l4 != 0 {
					v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+2)))
					v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
					if l5 != 0 {
						v98 = F__emscripten_memcpy_bulkmem(m, v88+(v90|v91<<(uint(int32(16))%32)&int32(2147418112)), l4, l5)
						mBase = m.M
					} else {
					}
				} else {
				}
				if v4 == int32(0) {
					v115 = v88
				} else {
					v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+6)))
					if int32(0) <= v104 {
						v107 = int32(8)
					} else {
						v107 = int32(16)
					}
					v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					if v111 != 0 {
						v112 = int32(0)
					} else {
						v112 = int32(2)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v88+v107+v112))) = uint8(v4)
					v115 = v88
				}
				m.G0 = v14 + int32(32)
				return v115
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
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[17])))
	if base.Ui32(v13-int32(100)) < base.Ui32(int32(-99)) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = v13
	v30 = int32(0)
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[18])))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v37 != 0 {
		v55 = v29
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v57 = v30 + int32(1)
	if base.Ui32(v57) < base.Ui32(v55) {
		v29 = v55
		v30 = v57
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)))
	if v38 != v7 {
		v55 = v29
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
	if v40 != v3 {
		v55 = v29
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v42 != l3 {
		v55 = v29
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+20)))
	if v44 != v2 {
		v55 = v29
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
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[17])))
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
	*(*int64)(unsafe.Add(mBase, uint32(v72)+644)) = v74
	v87 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+640)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v72)+44)) = v87
	v91 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+32)) = uint16(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+652)) = v91
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[17])))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[19])))
	if base.Ui32(v95) < base.Ui32(v96) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[17]))) = v110 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(int32(2))%32)))) = v72
	v126 = v72
	goto L1
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[18])))
	v109 = v98
	v110 = v95
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[19]))) = v96 << (uint(int32(1)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[18])))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[18]))) = v105
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[17])))
	v109 = v105
	v110 = v108
	goto L18
}
func F_ginInitConsistentFunction(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v7 == int32(3) {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(55)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(56)
		return
	} else {
		v15 = l0 + int32(3724)
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+76)))
		v17 = int32(28)
		v18 = v16 * v17
		v20 = v18 - v17
		*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v15 + v20
		v24 = l0 + int32(2828)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v20 + v24
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32)+l0)+uint32(_consts[14])))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v32
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v24+v18-int32(24))))
		if v39 != 0 {
			v40 = int32(57)
		} else {
			v40 = int32(58)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v40
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v18+v15-int32(24))))
		if v47 != 0 {
			v48 = int32(59)
		} else {
			v48 = int32(60)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v48
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
	v10 = F_ginPostingListDecodeAllSegments(m, l0, (v3+int32(1))&int32(131070)+int32(8), l1)
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v35 int32
	_ = v35
	v2 = l1
	v4 = l3
	v10 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), int32(68))
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(43)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(45)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+60)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l2
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+54)) = uint16(v2)
	v35 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+52)) = uint16(v35)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+36)) = uint8(v35)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(51)
	return
}
func F_gin_bool_consistent(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= v3 {
		v135 = int32(0)
		m.G0 = v13 + int32(16)
		return v135
	} else {
		v20 = l0 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v20
		v22 = F_palloc(m, v15)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v27 <= int32(0) {
			} else {
				v30 = int32(0)
				if v27 != int32(1) {
					v36 = v30
					v38 = v30
					v43 = v3
					for {
						v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v36<<(uint(int32(3))%32)))))
						if v49 == int32(2) {
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v38))))
							*(*uint8)(unsafe.Add(mBase, uint32(v36+v22))) = uint8(v54)
							v58 = v38 + int32(1)
						} else {
							v58 = v38
						}
						v60 = v36 | int32(1)
						v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v60<<(uint(int32(3))%32)))))
						if v64 == int32(2) {
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v58))))
							*(*uint8)(unsafe.Add(mBase, uint32(v22+v60))) = uint8(v69)
							v73 = v58 + int32(1)
						} else {
							v73 = v58
						}
						v74 = int32(2)
						v75 = v36 + v74
						v77 = v43 + v74
						if v77 != v27&int32(2147483646) {
							v36 = v75
							v38 = v73
							v43 = v77
							continue
						} else {
							break
						}
						break
					}
					v79 = v75
					v81 = v73
				} else {
					v79 = v30
					v81 = v30
				}
				if v27&int32(1) == int32(0) {
				} else {
					v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v79<<(uint(int32(3))%32)))))
					if v96 != int32(2) {
					} else {
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v81))))
						*(*uint8)(unsafe.Add(mBase, uint32(v79+v22))) = uint8(v101)
					}
				}
			}
			v116 = int32(8)
			v123 = F_execute(m, v20+v27<<(uint(int32(3))%32)-v116, v13+v116, int32(0), int32(1), int32(6812))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				v135 = v123
				m.G0 = v13 + int32(16)
				return v135
			}
		}
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
	if base.Ui32(v37) < base.Ui32((v15-int32(15))&int32(65535)) {
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
	F_errmsg_internal(m, int32(477483), v10)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(491614), int32(1266), int32(318702))
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
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
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
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
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
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
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
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
	var v745 int32
	_ = v745
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v802 int32
	_ = v802
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
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
	v815 = m.ExcPending
	if v815 != 0 {
		goto L6
	} else {
		goto L159
	}
L2:
	;
	m.G0 = v15 - int32(-64)
	return v802
L3:
	;
	if v19&int32(65534) == int32(10) {
		goto L74
	} else {
		goto L75
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
	v25 = F_DirectFunctionCall2Coll(m, int32(1336), int32(0), v24, v18)
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
		v802 = v25
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v802 = v25
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
	v359 = v355 + int32(5)
	v360 = F_palloc(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L69
	}
L15:
	;
	v51 = int32(4)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v53&int32(254) == int32(2) {
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
	v62 = v51
	goto L20
L19:
	;
	v62 = base.B2i32(v53 == int32(18)) << (uint(v51) % 32)
	goto L20
L20:
	;
	if v53 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v65 = v51
	goto L23
L22:
	;
	v65 = v62
	goto L23
L23:
	;
	v355 = v65
	v356 = v32
	v357 = v48
	goto L14
L24:
	;
	v66 = int32(1)
	v75 = int32(base.Ui32(v45)>>(uint(v66)%32)) - v66
	goto L26
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v75 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L26:
	;
	if v75 < int32(126) {
		v355 = v75
		v356 = v32
		v357 = v48
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v83 = v75 - int32(1636608432)
	if v48&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v337 ^ v329 - base.I32_rotl(v337, int32(24))
	v349 = F_pg_snprintf(m, v13+int32(-10), int32(10), int32(29439), v13+int32(-32))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L68
	}
L29:
	;
	v315 = int32(14)
	v317 = v311 ^ v312 - base.I32_rotl(v311, v315)
	v321 = v317 ^ v310 - base.I32_rotl(v317, int32(11))
	v325 = v321 ^ v311 - base.I32_rotl(v321, int32(25))
	v329 = v325 ^ v317 - base.I32_rotl(v325, int32(16))
	v333 = v329 ^ v321 - base.I32_rotl(v329, int32(4))
	v337 = v333 ^ v325 - base.I32_rotl(v333, v315)
	goto L28
L30:
	;
	switch v241 - int32(1) {
	case 0:
		v303 = v242
		v304 = v243
		v305 = v244
		goto L57
	case 1:
		v296 = v242
		v297 = v243
		v298 = v244
		goto L58
	case 2:
		v289 = v242
		v290 = v243
		v291 = v244
		goto L59
	case 3:
		v283 = v243
		v284 = v244
		goto L60
	case 4:
		v279 = v243
		v280 = v244
		goto L61
	case 5:
		v273 = v243
		v274 = v244
		goto L62
	case 6:
		v267 = v243
		v268 = v244
		goto L63
	case 7:
		v262 = v244
		goto L64
	case 8:
		v257 = v244
		goto L65
	case 9:
		v252 = v244
		goto L66
	case 10:
		goto L67
	default:
		v310 = v242
		v311 = v243
		v312 = v244
		goto L29
	}
L31:
	;
	v192 = v48
	v193 = v75
	v194 = v83
	v195 = v83
	v196 = v83
	goto L54
L32:
	;
	if base.Ui32(int32(11)) < base.Ui32(v75) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(v75) < base.Ui32(int32(12)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v240 = v48
	v241 = v75
	v242 = v83
	v243 = v83
	v244 = v83
	goto L30
L36:
	;
	switch v139 - int32(1) {
	case 0:
		v189 = v140
		goto L43
	case 1:
		v184 = v140
		goto L44
	case 2:
		goto L45
	case 3:
		v177 = v141
		goto L46
	case 4:
		v174 = v141
		goto L47
	case 5:
		v169 = v141
		goto L48
	case 6:
		goto L49
	case 7:
		v160 = v142
		goto L50
	case 8:
		v155 = v142
		goto L51
	case 9:
		v150 = v142
		goto L52
	case 10:
		goto L53
	default:
		v310 = v140
		v311 = v141
		v312 = v142
		goto L29
	}
L37:
	;
	v138 = v48
	v139 = v75
	v140 = v83
	v141 = v83
	v142 = v83
	goto L36
L38:
	;
	goto L39
L39:
	;
	v90 = v48
	v91 = v75
	v92 = v83
	v93 = v83
	v94 = v83
	goto L40
L40:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v97 = v96 + v93
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v101 = v100 + v94
	v103 = int32(4)
	v105 = v98 + v92 - v101 ^ base.I32_rotl(v101, v103)
	v109 = v97 - v105 ^ base.I32_rotl(v105, int32(6))
	v110 = v101 + v97
	v111 = v105 + v110
	v112 = v109 + v111
	v116 = v110 - v109 ^ base.I32_rotl(v109, int32(8))
	v120 = v111 - v116 ^ base.I32_rotl(v116, int32(16))
	v124 = v112 - v120 ^ base.I32_rotl(v120, int32(19))
	v125 = v116 + v112
	v126 = v120 + v125
	v127 = v124 + v126
	v131 = v125 - v124 ^ base.I32_rotl(v124, v103)
	v132 = int32(12)
	v133 = v90 + v132
	v135 = v91 - v132
	if base.Ui32(int32(11)) < base.Ui32(v135) {
		v90 = v133
		v91 = v135
		v92 = v126
		v93 = v127
		v94 = v131
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v138 = v133
	v139 = v135
	v140 = v126
	v141 = v127
	v142 = v131
	goto L36
L42:
	;
	goto L41
L43:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v310 = v189 + v190
	v311 = v141
	v312 = v142
	goto L29
L44:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	v189 = v185<<(uint(int32(8))%32) + v184
	goto L43
L45:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+2)))
	v184 = v180<<(uint(int32(16))%32) + v140
	goto L44
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v310 = v178 + v140
	v311 = v177
	v312 = v142
	goto L29
L47:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+4)))
	v177 = v174 + v175
	goto L46
L48:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+5)))
	v174 = v170<<(uint(int32(8))%32) + v169
	goto L47
L49:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+6)))
	v169 = v165<<(uint(int32(16))%32) + v141
	goto L48
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v310 = v161 + v140
	v311 = v163 + v141
	v312 = v160
	goto L29
L51:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+8)))
	v160 = v156<<(uint(int32(8))%32) + v155
	goto L50
L52:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+9)))
	v155 = v151<<(uint(int32(16))%32) + v150
	goto L51
L53:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+10)))
	v150 = v146<<(uint(int32(24))%32) + v142
	goto L52
L54:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v199 = v198 + v195
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v203 = v202 + v196
	v205 = int32(4)
	v207 = v200 + v194 - v203 ^ base.I32_rotl(v203, v205)
	v211 = v199 - v207 ^ base.I32_rotl(v207, int32(6))
	v212 = v203 + v199
	v213 = v207 + v212
	v214 = v211 + v213
	v218 = v212 - v211 ^ base.I32_rotl(v211, int32(8))
	v222 = v213 - v218 ^ base.I32_rotl(v218, int32(16))
	v226 = v214 - v222 ^ base.I32_rotl(v222, int32(19))
	v227 = v218 + v214
	v228 = v222 + v227
	v229 = v226 + v228
	v233 = v227 - v226 ^ base.I32_rotl(v226, v205)
	v234 = int32(12)
	v235 = v192 + v234
	v237 = v193 - v234
	if base.Ui32(int32(11)) < base.Ui32(v237) {
		v192 = v235
		v193 = v237
		v194 = v228
		v195 = v229
		v196 = v233
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v240 = v235
	v241 = v237
	v242 = v228
	v243 = v229
	v244 = v233
	goto L30
L56:
	;
	goto L55
L57:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v310 = v303 + v306
	v311 = v304
	v312 = v305
	goto L29
L58:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	v303 = v299<<(uint(int32(8))%32) + v296
	v304 = v297
	v305 = v298
	goto L57
L59:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+2)))
	v296 = v292<<(uint(int32(16))%32) + v289
	v297 = v290
	v298 = v291
	goto L58
L60:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+3)))
	v289 = v285<<(uint(int32(24))%32) + v242
	v290 = v283
	v291 = v284
	goto L59
L61:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+4)))
	v283 = v279 + v281
	v284 = v280
	goto L60
L62:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+5)))
	v279 = v275<<(uint(int32(8))%32) + v273
	v280 = v274
	goto L61
L63:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+6)))
	v273 = v269<<(uint(int32(16))%32) + v267
	v274 = v268
	goto L62
L64:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+7)))
	v267 = v263<<(uint(int32(24))%32) + v243
	v268 = v262
	goto L63
L65:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+8)))
	v262 = v258<<(uint(int32(8))%32) + v257
	goto L64
L66:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+9)))
	v257 = v253<<(uint(int32(16))%32) + v252
	goto L65
L67:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+10)))
	v252 = v248<<(uint(int32(24))%32) + v244
	goto L66
L68:
	;
	v355 = int32(8)
	v356 = int32(17)
	v357 = v13 + int32(-10)
	goto L14
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v360)+4)) = uint8(v356)
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v359 << (uint(int32(2)) % 32)
	if v355 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v360
	v802 = v39
	goto L2
L71:
	;
	v368 = F__emscripten_memcpy_bulkmem(m, v360+int32(5), v357, v355)
	mBase = m.M
	goto L73
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v376 = F_pg_detoast_datum(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(1)) < base.Ui32((v19-int32(15))&int32(65535)) {
		goto L1
	} else {
		goto L155
	}
L77:
	;
	F_deconstruct_array_builtin(m, v376, int32(25), v13+int32(-16), v13+int32(-20), v13+int32(-24))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v390 = F_palloc(m, v387<<(uint(int32(2))%32))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v392 = int32(0)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v393 <= v392 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v767
	if v19 != int32(11) {
		v802 = v390
		goto L2
	} else {
		goto L153
	}
L81:
	;
	v767 = v2
	goto L80
L82:
	;
	goto L83
L83:
	;
	v396 = v392
	v397 = v393
	v400 = v2
	goto L84
L84:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v396))))
	if v410 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v767 = v756
	goto L80
L86:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v413+v396<<(uint(int32(2))%32))))
	v418 = int32(1)
	v419 = v417 + v418
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417))))
	v424 = v422 & v418
	if v424 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v754 = v397
	v756 = v400
	goto L88
L88:
	;
	v761 = v396 + int32(1)
	if v761 < v754 {
		v396 = v761
		v397 = v754
		v400 = v756
		goto L84
	} else {
		goto L152
	}
L89:
	;
	v425 = v419
	goto L91
L90:
	;
	v425 = v417 + int32(4)
	goto L91
L91:
	;
	if v422 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v736 = v732 + int32(5)
	v737 = F_palloc(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L147
	}
L93:
	;
	v428 = int32(4)
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	if v430&int32(254) == int32(2) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if v424 != 0 {
		goto L102
	} else {
		goto L103
	}
L96:
	;
	v439 = v428
	goto L98
L97:
	;
	v439 = base.B2i32(v430 == int32(18)) << (uint(v428) % 32)
	goto L98
L98:
	;
	if v430 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v442 = v428
	goto L101
L100:
	;
	v442 = v439
	goto L101
L101:
	;
	v732 = v442
	v733 = v425
	v734 = int32(1)
	goto L92
L102:
	;
	v444 = int32(1)
	v453 = int32(base.Ui32(v422)>>(uint(v444)%32)) - v444
	goto L104
L103:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v453 = int32(base.Ui32(v448)>>(uint(int32(2))%32)) - int32(4)
	goto L104
L104:
	;
	if v453 < int32(126) {
		v732 = v453
		v733 = v425
		v734 = int32(1)
		goto L92
	} else {
		goto L105
	}
L105:
	;
	v462 = v453 - int32(1636608432)
	if v425&int32(3) != 0 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v716 ^ v708 - base.I32_rotl(v716, int32(24))
	v726 = F_pg_snprintf(m, v13+int32(-10), int32(10), int32(29439), v15)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L6
	} else {
		goto L146
	}
L107:
	;
	v694 = int32(14)
	v696 = v690 ^ v691 - base.I32_rotl(v690, v694)
	v700 = v696 ^ v689 - base.I32_rotl(v696, int32(11))
	v704 = v700 ^ v690 - base.I32_rotl(v700, int32(25))
	v708 = v704 ^ v696 - base.I32_rotl(v704, int32(16))
	v712 = v708 ^ v700 - base.I32_rotl(v708, int32(4))
	v716 = v712 ^ v704 - base.I32_rotl(v712, v694)
	goto L106
L108:
	;
	switch v620 - int32(1) {
	case 0:
		v682 = v621
		v683 = v622
		v684 = v623
		goto L135
	case 1:
		v675 = v621
		v676 = v622
		v677 = v623
		goto L136
	case 2:
		v668 = v621
		v669 = v622
		v670 = v623
		goto L137
	case 3:
		v662 = v622
		v663 = v623
		goto L138
	case 4:
		v658 = v622
		v659 = v623
		goto L139
	case 5:
		v652 = v622
		v653 = v623
		goto L140
	case 6:
		v646 = v622
		v647 = v623
		goto L141
	case 7:
		v641 = v623
		goto L142
	case 8:
		v636 = v623
		goto L143
	case 9:
		v631 = v623
		goto L144
	case 10:
		goto L145
	default:
		v689 = v621
		v690 = v622
		v691 = v623
		goto L107
	}
L109:
	;
	v571 = v425
	v572 = v453
	v573 = v462
	v574 = v462
	v575 = v462
	goto L132
L110:
	;
	if base.Ui32(int32(11)) < base.Ui32(v453) {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if base.Ui32(v453) < base.Ui32(int32(12)) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v619 = v425
	v620 = v453
	v621 = v462
	v622 = v462
	v623 = v462
	goto L108
L114:
	;
	switch v518 - int32(1) {
	case 0:
		v568 = v519
		goto L121
	case 1:
		v563 = v519
		goto L122
	case 2:
		goto L123
	case 3:
		v556 = v520
		goto L124
	case 4:
		v553 = v520
		goto L125
	case 5:
		v548 = v520
		goto L126
	case 6:
		goto L127
	case 7:
		v539 = v521
		goto L128
	case 8:
		v534 = v521
		goto L129
	case 9:
		v529 = v521
		goto L130
	case 10:
		goto L131
	default:
		v689 = v519
		v690 = v520
		v691 = v521
		goto L107
	}
L115:
	;
	v517 = v425
	v518 = v453
	v519 = v462
	v520 = v462
	v521 = v462
	goto L114
L116:
	;
	goto L117
L117:
	;
	v469 = v425
	v470 = v453
	v471 = v462
	v472 = v462
	v473 = v462
	goto L118
L118:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	v476 = v475 + v472
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v469)+8))
	v480 = v479 + v473
	v482 = int32(4)
	v484 = v477 + v471 - v480 ^ base.I32_rotl(v480, v482)
	v488 = v476 - v484 ^ base.I32_rotl(v484, int32(6))
	v489 = v480 + v476
	v490 = v484 + v489
	v491 = v488 + v490
	v495 = v489 - v488 ^ base.I32_rotl(v488, int32(8))
	v499 = v490 - v495 ^ base.I32_rotl(v495, int32(16))
	v503 = v491 - v499 ^ base.I32_rotl(v499, int32(19))
	v504 = v495 + v491
	v505 = v499 + v504
	v506 = v503 + v505
	v510 = v504 - v503 ^ base.I32_rotl(v503, v482)
	v511 = int32(12)
	v512 = v469 + v511
	v514 = v470 - v511
	if base.Ui32(int32(11)) < base.Ui32(v514) {
		v469 = v512
		v470 = v514
		v471 = v505
		v472 = v506
		v473 = v510
		goto L118
	} else {
		goto L120
	}
L119:
	;
	v517 = v512
	v518 = v514
	v519 = v505
	v520 = v506
	v521 = v510
	goto L114
L120:
	;
	goto L119
L121:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	v689 = v568 + v569
	v690 = v520
	v691 = v521
	goto L107
L122:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)))
	v568 = v564<<(uint(int32(8))%32) + v563
	goto L121
L123:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+2)))
	v563 = v559<<(uint(int32(16))%32) + v519
	goto L122
L124:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	v689 = v557 + v519
	v690 = v556
	v691 = v521
	goto L107
L125:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+4)))
	v556 = v553 + v554
	goto L124
L126:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+5)))
	v553 = v549<<(uint(int32(8))%32) + v548
	goto L125
L127:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+6)))
	v548 = v544<<(uint(int32(16))%32) + v520
	goto L126
L128:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
	v689 = v540 + v519
	v690 = v542 + v520
	v691 = v539
	goto L107
L129:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+8)))
	v539 = v535<<(uint(int32(8))%32) + v534
	goto L128
L130:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+9)))
	v534 = v530<<(uint(int32(16))%32) + v529
	goto L129
L131:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+10)))
	v529 = v525<<(uint(int32(24))%32) + v521
	goto L130
L132:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	v578 = v577 + v574
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v571)+8))
	v582 = v581 + v575
	v584 = int32(4)
	v586 = v579 + v573 - v582 ^ base.I32_rotl(v582, v584)
	v590 = v578 - v586 ^ base.I32_rotl(v586, int32(6))
	v591 = v582 + v578
	v592 = v586 + v591
	v593 = v590 + v592
	v597 = v591 - v590 ^ base.I32_rotl(v590, int32(8))
	v601 = v592 - v597 ^ base.I32_rotl(v597, int32(16))
	v605 = v593 - v601 ^ base.I32_rotl(v601, int32(19))
	v606 = v597 + v593
	v607 = v601 + v606
	v608 = v605 + v607
	v612 = v606 - v605 ^ base.I32_rotl(v605, v584)
	v613 = int32(12)
	v614 = v571 + v613
	v616 = v572 - v613
	if base.Ui32(int32(11)) < base.Ui32(v616) {
		v571 = v614
		v572 = v616
		v573 = v607
		v574 = v608
		v575 = v612
		goto L132
	} else {
		goto L134
	}
L133:
	;
	v619 = v614
	v620 = v616
	v621 = v607
	v622 = v608
	v623 = v612
	goto L108
L134:
	;
	goto L133
L135:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	v689 = v682 + v685
	v690 = v683
	v691 = v684
	goto L107
L136:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+1)))
	v682 = v678<<(uint(int32(8))%32) + v675
	v683 = v676
	v684 = v677
	goto L135
L137:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+2)))
	v675 = v671<<(uint(int32(16))%32) + v668
	v676 = v669
	v677 = v670
	goto L136
L138:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+3)))
	v668 = v664<<(uint(int32(24))%32) + v621
	v669 = v662
	v670 = v663
	goto L137
L139:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+4)))
	v662 = v658 + v660
	v663 = v659
	goto L138
L140:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+5)))
	v658 = v654<<(uint(int32(8))%32) + v652
	v659 = v653
	goto L139
L141:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+6)))
	v652 = v648<<(uint(int32(16))%32) + v646
	v653 = v647
	goto L140
L142:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+7)))
	v646 = v642<<(uint(int32(24))%32) + v622
	v647 = v641
	goto L141
L143:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+8)))
	v641 = v637<<(uint(int32(8))%32) + v636
	goto L142
L144:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+9)))
	v636 = v632<<(uint(int32(16))%32) + v631
	goto L143
L145:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+10)))
	v631 = v627<<(uint(int32(24))%32) + v623
	goto L144
L146:
	;
	v732 = int32(8)
	v733 = v13 + int32(-10)
	v734 = int32(17)
	goto L92
L147:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v737)+4)) = uint8(v734)
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = v736 << (uint(int32(2)) % 32)
	if v732 != 0 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390+v400<<(uint(int32(2))%32)))) = v737
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v754 = v753
	v756 = v400 + int32(1)
	goto L88
L149:
	;
	v745 = F__emscripten_memcpy_bulkmem(m, v737+int32(5), v733, v732)
	mBase = m.M
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L148
L152:
	;
	goto L85
L153:
	;
	if v767 != 0 {
		v802 = v390
		goto L2
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v802 = v390
	goto L2
L155:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v787 = F_pg_detoast_datum(m, v786)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v791 = F_extract_jsp_query(m, v787, v19, int32(0), v18, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	if v791 != 0 {
		v802 = v791
		goto L2
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v802 = int32(0)
	goto L2
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
	F_errmsg_internal(m, int32(477483), v13+int32(-48))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(491614), int32(921), int32(15940))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
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
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(6809), int32(2661))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
