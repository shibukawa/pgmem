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
	F_errmsg(m, int32(199854), v14)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(515214), int32(1311), int32(107064))
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
						F_errmsg(m, int32(722989), v14)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(522819), int32(111), int32(402983))
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
			v123 = F_execute(m, v20+v27<<(uint(int32(3))%32)-v116, v13+v116, int32(0), int32(1), int32(7619))
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
func F_gin_check_parent_keys_consistency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
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
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v740 int32
	_ = v740
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int64
	_ = v785
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v948 int32
	_ = v948
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int64
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1211 int32
	_ = v1211
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1234 int32
	_ = v1234
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1348 int32
	_ = v1348
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1414 int32
	_ = v1414
	var v1421 int32
	_ = v1421
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1477 int32
	_ = v1477
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1499 int32
	_ = v1499
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1568 int32
	_ = v1568
	var v1574 int32
	_ = v1574
	var v1598 int32
	_ = v1598
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1649 int32
	_ = v1649
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1710 int32
	_ = v1710
	var v1718 int32
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1757 int32
	_ = v1757
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1874 int32
	_ = v1874
	var v1881 int32
	_ = v1881
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1901 int32
	_ = v1901
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1930 int32
	_ = v1930
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1945 int64
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1975 int32
	_ = v1975
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v2003 int32
	_ = v2003
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2038 int32
	_ = v2038
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2065 int32
	_ = v2065
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2091 int32
	_ = v2091
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2117 int32
	_ = v2117
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2210 int32
	_ = v2210
	var v2215 int32
	_ = v2215
	var v2222 int32
	_ = v2222
	v37 = m.G0
	v39 = v37 - int32(7280)
	m.G0 = v39
	v43 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v45 = int32(4554128)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v52 = F_AllocSetContextCreateInternal(m, v46, int32(66652), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	F_initGinState(m, v39+int32(564), l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v61 = F_palloc0(m, int32(20))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = int64(8589934591)
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = int64(0)
	v74 = v61
	v89 = int32(-1)
	goto L17
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L1
	} else {
		goto L390
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L1
	} else {
		goto L381
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L1
	} else {
		goto L377
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L1
	} else {
		goto L373
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L369
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L1
	} else {
		goto L365
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L361
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L1
	} else {
		goto L357
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L353
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L1
	} else {
		goto L349
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L1
	} else {
		goto L345
	}
L17:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v105 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v54
	F_MemoryContextDelete(m, v52)
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L1
	} else {
		goto L344
	}
L19:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v108 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v111 = F_ReadBufferExtended(m, l0, v108, v109, v108, v43)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	F_LockBuffer(m, v111, int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v116 = int32(0)
	v117 = base.B2i32(v116 <= v111)
	if v117 == v116 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+12)))
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+16)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v137)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v117 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121+(v111^int32(-1))<<(uint(int32(2))%32))))
	v135 = v127
	goto L25
L27:
	;
	goto L28
L28:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v135 = v129 + v111<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+14)))
	if v159 == int32(0) {
		goto L6
	} else {
		goto L33
	}
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144+(v111^int32(-1))<<(uint(int32(2))%32))))
	v158 = v150
	goto L29
L31:
	;
	goto L32
L32:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v158 = v152 + v111<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L33:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+19)))
	v163 = int32(8)
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+16)))
	if (v162<<(uint(v163)%32)-v165)&int32(65535) != v163 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158+v165)+6)))
	if v172&int32(4) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v136) {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	if v172&int32(2) == int32(0) {
		goto L8
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+12)))
	if base.Ui32(v214) < base.Ui32(int32(25)) {
		goto L35
	} else {
		goto L46
	}
L39:
	;
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+12)))
	if base.Ui32(v179) < base.Ui32(int32(25)) {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	if (v179+int32(262120))&int32(262140) == int32(0) {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+532)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+528)) = v195 + int32(4)
	F_errmsg(m, int32(174216), v39+int32(528))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(519620), int32(704), int32(425470))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
	if base.Ui32(int32(409)) <= base.Ui32(int32(base.Ui32(v214+int32(262120))>>(uint(int32(2))%32))&int32(65535)) {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	goto L35
L48:
	;
	v233 = int32(base.Ui32(v136+int32(262120)) >> (uint(int32(2)) % 32))
	goto L50
L49:
	;
	v233 = int32(0)
	goto L50
L50:
	;
	v236 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v236 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+500)) = v233 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+496)) = v238
	F_errmsg_internal(m, int32(63810), v39+int32(496))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v258 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	F_errfinish(m, int32(519620), int32(444), int32(23383))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+16)))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v343)+6)))
	if v345&int32(2) == int32(0) {
		v355 = v89
		goto L75
	} else {
		goto L76
	}
L58:
	;
	v265 = F_gintuple_get_key(m, v39+int32(564), v258, v39+int32(6240))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v270 = F_gintuple_get_attrnum(m, v39+int32(564), v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v277 = F_PageGetItemIdCareful_1(m, l0, v274, v135, v233&int32(65535))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v282 = v135 + v279&int32(32767)
	v283 = F_gintuple_get_attrnum(m, v39+int32(564), v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v289 = F_gintuple_get_key(m, v39+int32(564), v282, v39+int32(7274))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v139 == int32(-1) {
		goto L57
	} else {
		goto L64
	}
L64:
	;
	v295 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1131]))))
	v296 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132]))))
	v297 = F_ginCompareAttEntries(m, v39+int32(564), v283, v289, v295, v270, v265, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if int32(0) <= v297 {
		goto L57
	} else {
		goto L66
	}
L66:
	;
	v303 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v303 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+480)) = base.I64_rotl(v305, int64(32))
	F_errmsg_internal(m, int32(63597), v39+int32(480))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v325 = F_palloc(m, int32(20))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	F_errfinish(m, int32(519620), int32(474), int32(23383))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v330 = F_CopyIndexTuple(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+4)) = v330
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+12)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v325)+8)) = v333
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+16)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v325
	goto L57
L75:
	;
	v357 = v233 & int32(65535)
	if v357 != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v89 == int32(-1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v355 = v350
	goto L75
L78:
	;
	goto L79
L79:
	;
	if v350 != v89 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	v355 = v89
	goto L75
L81:
	;
	v359 = v135 + int32(6)
	v360 = int32(0)
	v366 = v360
	v378 = v360
	v379 = int32(1)
	goto L84
L82:
	;
	goto L83
L83:
	;
	F_LockBuffer(m, v111, int32(0))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L1
	} else {
		goto L336
	}
L84:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v403 = v379 & int32(65535)
	v404 = F_PageGetItemIdCareful_1(m, l0, v401, v135, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	goto L83
L86:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v409 = v135 + v406&int32(32767)
	v410 = F_gintuple_get_attrnum(m, v39+int32(564), v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v415 = int32(7)
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409)+6)))
	if (int32(base.Ui32(v412)>>(uint(int32(17))%32))+v415)&int32(65528) == (v419&int32(8191)+v415)&int32(16376) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+16)))
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359+v839))))
	if v841&int32(2) == int32(0) {
		goto L154
	} else {
		goto L155
	}
L89:
	;
	v781 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L149
	}
L90:
	;
	F_UnlockReleaseBuffer(m, v491)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L148
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L144
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L140
	}
L93:
	;
	v431 = F_gintuple_get_key(m, v39+int32(564), v409, v39+int32(563))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L136
	}
L96:
	;
	if v403 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v461 == int32(0) {
		goto L88
	} else {
		goto L106
	}
L98:
	;
	if v403 != v357 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v446 = v39 + int32(564)
	v453 = F_gintuple_get_key(m, v446, v366, v39+int32(6240))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L103
	}
L100:
	;
	if v139 != int32(-1) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+16)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359+v438))))
	if v440&int32(2) == int32(0) {
		goto L97
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v455 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132]))))
	v456 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+563)))
	v457 = F_ginCompareAttEntries(m, v446, v378&int32(65535), v453, v455, v410, v431, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v457 {
		goto L92
	} else {
		goto L105
	}
L105:
	;
	goto L97
L106:
	;
	if v403 != v357 {
		goto L88
	} else {
		goto L107
	}
L107:
	;
	v467 = F_gintuple_get_attrnum(m, v39+int32(564), v461)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v474 = F_gintuple_get_key(m, v39+int32(564), v471, v39+int32(6240))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v478 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+563)))
	v479 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132]))))
	v480 = F_ginCompareAttEntries(m, v39+int32(564), v410, v431, v478, v467, v474, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v480 <= int32(0) {
		goto L88
	} else {
		goto L111
	}
L111:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	F_pfree(m, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v488 = int32(0)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v491 = F_ReadBufferExtended(m, l0, v488, v489, v488, v43)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_LockBuffer(m, v491, int32(1))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if v491 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513)+16)))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514+v513)+6)))
	if v516&int32(2) != 0 {
		goto L90
	} else {
		goto L119
	}
L116:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v499+(v491^int32(-1))<<(uint(int32(2))%32))))
	v513 = v505
	goto L115
L117:
	;
	goto L118
L118:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v513 = v507 + v491<<(uint(int32(13))%32) + int32(-8192)
	goto L115
L119:
	;
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513)+12)))
	if base.Ui32(v519) < base.Ui32(int32(25)) {
		goto L90
	} else {
		goto L120
	}
L120:
	;
	v528 = int32(base.Ui32(v519+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v528 == int32(0) {
		goto L90
	} else {
		goto L121
	}
L121:
	;
	v532 = int32(1)
	goto L122
L122:
	;
	v569 = F_PageGetItemIdCareful_1(m, l0, v489, v513, v532&int32(65535))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	v586 = F_CopyIndexTuple(m, v574)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L129
	}
L124:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v574 = v513 + v571&int32(32767)
	v575 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574))))
	v578 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574)+2)))
	if v487 != v575<<(uint(int32(16))%32)|v578 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v582 = v532 + int32(1)
	if base.Ui32(v582&int32(65535)) <= base.Ui32(v528) {
		v532 = v582
		goto L122
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	goto L123
L128:
	;
	goto L90
L129:
	;
	F_UnlockReleaseBuffer(m, v491)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v586
	if v586 == int32(0) {
		goto L89
	} else {
		goto L131
	}
L131:
	;
	v595 = F_gintuple_get_attrnum(m, v39+int32(564), v586)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v602 = F_gintuple_get_key(m, v39+int32(564), v599, v39+int32(6240))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v606 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+563)))
	v607 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132]))))
	v608 = F_ginCompareAttEntries(m, v39+int32(564), v410, v431, v606, v595, v602, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	if int32(0) < v608 {
		goto L91
	} else {
		goto L135
	}
L135:
	;
	goto L88
L136:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+456)) = v379 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+452)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v39)+448)) = v619 + int32(4)
	F_errmsg(m, int32(44882), v39+int32(448))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(519620), int32(516), int32(23383))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+444)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+440)) = v379 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+436)) = v650
	*(*int32)(unsafe.Add(mBase, uint32(v39)+432)) = v649 + int32(4)
	F_errmsg(m, int32(51410), v39+int32(432))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(519620), int32(542), int32(23383))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+424)) = v233 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+420)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v39)+416)) = v680 + int32(4)
	F_errmsg(m, int32(44055), v39+int32(416))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(519620), int32(594), int32(23383))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = int32(0)
	goto L89
L149:
	;
	if v781 == int32(0) {
		goto L88
	} else {
		goto L150
	}
L150:
	;
	v785 = *(*int64)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+400)) = base.I64_rotl(v785, int64(32))
	F_errmsg_internal(m, int32(109113), v39+int32(400))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(519620), int32(576), int32(23383))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	goto L88
L153:
	;
	v1794 = F_CopyIndexTuple(m, v409)
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L1
	} else {
		goto L334
	}
L154:
	;
	v847 = F_palloc(m, int32(20))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v872 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409))))
	v874 = v872 << (uint(int32(16)) % 32)
	v875 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409)+2)))
	v876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409)+4)))
	if v876 == int32(65535) {
		goto L164
	} else {
		goto L165
	}
L157:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v847))) = v849 + int32(1)
	if v403 == v357 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+4)) = v859
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v847)+8)) = v861
	v863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409)+2)))
	v864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409))))
	*(*int32)(unsafe.Add(mBase, uint32(v847)+12)) = v863 | v864<<(uint(int32(16))%32)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v847)+16)) = v869
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v847
	goto L153
L159:
	;
	if v139 == int32(-1) {
		v859 = int32(0)
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v857 = F_CopyIndexTuple(m, v409)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v859 = v857
	goto L158
L164:
	;
	v880 = int32(4554128)
	v882 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v1627 = v409 + (v874&int32(2147418112) | v875)
	if v874 < int32(0) {
		goto L311
	} else {
		goto L312
	}
L167:
	;
	v884 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v890 = F_AllocSetContextCreateInternal(m, v884, int32(66686), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v892 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v890
	v895 = F_palloc0(m, int32(24))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v897 = v874 | v875
	*(*int32)(unsafe.Add(mBase, uint32(v895)+16)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v895)+12)) = int32(-1)
	v901 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v895)+8)) = uint16(v901)
	*(*int64)(unsafe.Add(mBase, uint32(v895))) = int64(-4294967296)
	v907 = F_errstart(m, int32(12), v901)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	if v907 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+352)) = v897
	F_errmsg_internal(m, int32(51526), v39+int32(352))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v929 = v895
	v948 = int32(-1)
	goto L176
L174:
	;
	F_errfinish(m, int32(519620), int32(161), int32(23417))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v963 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v892
	F_MemoryContextDelete(m, v890)
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L1
	} else {
		goto L308
	}
L178:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v966 = int32(0)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	v969 = F_ReadBufferExtended(m, l0, v966, v967, v966, v882)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L182
	}
L181:
	;
	goto L180
L182:
	;
	F_LockBuffer(m, v969, int32(1))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	if v969 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	F_LockBuffer(m, v969, int32(0))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L1
	} else {
		goto L304
	}
L185:
	;
	v993 = v991 + int32(16)
	v994 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v993))))
	v995 = v994 + v991
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+6)))
	if v996&int32(2) != 0 {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v977 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v977+(v969^int32(-1))<<(uint(int32(2))%32))))
	v991 = v983
	goto L185
L187:
	;
	goto L188
L188:
	;
	v985 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v991 = v985 + v969<<(uint(int32(13))%32) + int32(-8192)
	goto L185
L189:
	;
	v999 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1133]))) = uint16(v999)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1131]))) = v999
	v1005 = F_errstart(m, int32(14), v999)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	v1220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995)+4)))
	v1223 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L234
	}
L192:
	;
	if v1005 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+336)) = v1007
	F_errmsg_internal(m, int32(356070), v39+int32(336))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v929)))
	if v948 == int32(-1) {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	F_errfinish(m, int32(519620), int32(191), int32(23417))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v1031 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1133]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+316)) = uint16(v1031)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1131])))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+312)) = v1033
	v1039 = F_GinDataLeafPageGetItems(m, v991, v39+int32(7268), v39+int32(312))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L203
	}
L199:
	;
	v1028 = v1024
	goto L198
L200:
	;
	goto L201
L201:
	;
	if v1024 != v948 {
		goto L15
	} else {
		goto L202
	}
L202:
	;
	v1028 = v948
	goto L198
L203:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1134])))
	if int32(0) < v1041 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v929)+12))
	v1094 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L210
	}
L205:
	;
	v1044 = int32(6)
	v1046 = v1039 + v1041*v1044
	v1049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046-int32(4)))))
	v1052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046-v1044))))
	v1053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1039)+2)))
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1039))))
	v1055 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1039)+4)))
	v1058 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046-int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+304)) = v1058
	*(*int32)(unsafe.Add(mBase, uint32(v39)+296)) = v1055
	v1061 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+292)) = v1053 | v1054<<(uint(v1061)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+288)) = v1041
	*(*int32)(unsafe.Add(mBase, uint32(v39)+300)) = v1049 | v1052<<(uint(v1061)%32)
	v1076 = F_pg_snprintf(m, v39+int32(6240), int32(1024), int32(701357), v39+int32(288))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v1083 = F_pg_snprintf(m, v39+int32(6240), int32(1024), int32(183420), int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L209
	}
L208:
	;
	goto L204
L209:
	;
	goto L204
L210:
	;
	if v1091 != int32(-1) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v929)+12))
	if v1150 == int32(-1) {
		v1598 = v1028
		goto L184
	} else {
		goto L221
	}
L212:
	;
	F_errfinish(m, int32(519620), v1138, int32(23417))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L220
	}
L213:
	;
	if v1094 == int32(0) {
		goto L211
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	if v1094 == int32(0) {
		goto L211
	} else {
		goto L218
	}
L216:
	;
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+6)))
	v1101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+4)))
	v1102 = *(*int64)(unsafe.Add(mBase, uint32(v929)+12))
	v1103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+272)) = v39 + int32(6240)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+268)) = v1103
	*(*int64)(unsafe.Add(mBase, uint32(v39)+256)) = base.I64_rotl(v1102, int64(32))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+264)) = v1100 | v1101<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(216083), v39+int32(256))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v1138 = int32(219)
	goto L212
L218:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+240)) = v1123
	*(*int32)(unsafe.Add(mBase, uint32(v39)+244)) = v39 + int32(6240)
	F_errmsg_internal(m, int32(216061), v39+int32(240))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v1138 = int32(223)
	goto L212
L220:
	;
	goto L211
L221:
	;
	v1153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+8)))
	if v1153 == int32(0) {
		v1598 = v1028
		goto L184
	} else {
		goto L222
	}
L222:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1134])))
	if v1156 <= int32(0) {
		v1598 = v1028
		goto L184
	} else {
		goto L223
	}
L223:
	;
	v1160 = v929 + int32(4)
	v1161 = int32(6)
	v1165 = v1039 + v1156*v1161 - v1161
	v1169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1160)+2)))
	v1170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1160))))
	v1171 = int32(16)
	v1173 = v1169 | v1170<<(uint(v1171)%32)
	v1174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165)+2)))
	v1175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165))))
	v1178 = v1174 | v1175<<(uint(v1171)%32)
	if base.Ui32(v1173) < base.Ui32(v1178) {
		v1189 = int32(-1)
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if int32(0) <= v1189 {
		v1598 = v1028
		goto L184
	} else {
		goto L229
	}
L225:
	;
	goto L224
L226:
	;
	if base.Ui32(v1178) < base.Ui32(v1173) {
		v1189 = int32(1)
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v1183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1160)+4)))
	v1184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165)+4)))
	if base.Ui32(v1183) < base.Ui32(v1184) {
		v1189 = int32(-1)
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v1189 = base.B2i32(base.Ui32(v1184) < base.Ui32(v1183))
	goto L225
L229:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+228)) = v1200
	*(*int32)(unsafe.Add(mBase, uint32(v39)+224)) = v1199 + int32(4)
	F_errmsg(m, int32(52123), v39+int32(224))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(519620), int32(231), int32(23417))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	if v1223 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+212)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v39)+208)) = v1225
	F_errmsg_internal(m, int32(496026), v39+int32(208))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v929)+12))
	v1246 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	F_errfinish(m, int32(519620), int32(246), int32(23417))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	if v1243 != int32(-1) {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	v1298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v991)+12)))
	v1302 = base.I32_div_u_s(v1298-int32(32), int32(10))
	if v1302 != v1220 {
		goto L14
	} else {
		goto L251
	}
L242:
	;
	F_errfinish(m, int32(519620), v1286, int32(23417))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L250
	}
L243:
	;
	if v1246 == int32(0) {
		goto L241
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	if v1246 == int32(0) {
		goto L241
	} else {
		goto L248
	}
L246:
	;
	v1252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+6)))
	v1253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+4)))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v929)+12))
	v1256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+192)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v39)+184)) = v1255
	*(*int32)(unsafe.Add(mBase, uint32(v39)+180)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v39)+176)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v39)+188)) = v1252 | v1253<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(701198), v39+int32(176))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1286 = int32(252)
	goto L242
L248:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+164)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v39)+160)) = v1273
	F_errmsg_internal(m, int32(160040), v39+int32(160))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v1286 = int32(255)
	goto L242
L250:
	;
	goto L241
L251:
	;
	v1306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v991)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1135]))) = uint16(v1306)
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v991)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132]))) = v1308
	v1311 = v929 + int32(4)
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+8)))
	if v1312 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	if v1220 == int32(0) {
		v1598 = v948
		goto L184
	} else {
		goto L262
	}
L253:
	;
	if v1219 == int32(-1) {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v1319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+2)))
	v1320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311))))
	v1321 = int32(16)
	v1324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1136]))))
	v1325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132]))))
	if v1319|v1320<<(uint(v1321)%32) == v1324|v1325<<(uint(v1321)%32) {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	if v1335 == int32(0) {
		goto L13
	} else {
		goto L261
	}
L256:
	;
	goto L255
L257:
	;
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+4)))
	v1332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1135]))))
	if v1331 == v1332 {
		v1335 = int32(1)
		goto L256
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v1335 = int32(0)
	goto L256
L260:
	;
	goto L259
L261:
	;
	goto L252
L262:
	;
	v1348 = int32(1)
	goto L263
L263:
	;
	v1380 = v1348 * int32(10)
	v1381 = v991 + int32(22) + v1380
	v1384 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L265
	}
L264:
	;
	v1598 = v948
	goto L184
L265:
	;
	v1386 = base.B2i32(v1348 != v1220)
	v1387 = int32(0)
	if base.B2i32(v1386 == v1387)&base.B2i32(v1219 == int32(-1)) == v1387 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	if v1348 != v1220 {
		goto L293
	} else {
		goto L294
	}
L267:
	;
	if v1384 != 0 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	if v1384 != 0 {
		goto L286
	} else {
		goto L287
	}
L270:
	;
	v1394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+6)))
	v1395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+4)))
	v1396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+2)))
	v1397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381))))
	v1398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+100)) = v1398
	v1400 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = v1396 | v1397<<(uint(v1400)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v1394 | v1395<<(uint(v1400)%32)
	F_errmsg_internal(m, int32(62276), v39+int32(96))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	if v1348 == int32(1) {
		goto L266
	} else {
		goto L275
	}
L273:
	;
	F_errfinish(m, int32(519620), int32(312), int32(23417))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	goto L272
L275:
	;
	v1429 = v1381 + int32(4)
	v1430 = v1380 + v993
	v1434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1429)+2)))
	v1435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1429))))
	v1436 = int32(16)
	v1438 = v1434 | v1435<<(uint(v1436)%32)
	v1439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1430)+2)))
	v1440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1430))))
	v1443 = v1439 | v1440<<(uint(v1436)%32)
	if base.Ui32(v1438) < base.Ui32(v1443) {
		v1454 = int32(-1)
		goto L277
	} else {
		goto L278
	}
L276:
	;
	if int32(0) <= v1454 {
		goto L266
	} else {
		goto L281
	}
L277:
	;
	goto L276
L278:
	;
	if base.Ui32(v1443) < base.Ui32(v1438) {
		v1454 = int32(1)
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v1448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1429)+4)))
	v1449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1430)+4)))
	if base.Ui32(v1448) < base.Ui32(v1449) {
		v1454 = int32(-1)
		goto L277
	} else {
		goto L280
	}
L280:
	;
	v1454 = base.B2i32(base.Ui32(v1449) < base.Ui32(v1448))
	goto L277
L281:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+88)) = v1348
	*(*int32)(unsafe.Add(mBase, uint32(v39)+84)) = v1465
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v1464 + int32(4)
	F_errmsg(m, int32(44943), v39+int32(80))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(519620), int32(341), int32(23417))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	v1485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+2)))
	v1486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381))))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v1485 | v1486<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(62276), v39-int32(-64))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v1509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+6)))
	v1510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+4)))
	if v1509|v1510 != 0 {
		goto L11
	} else {
		goto L291
	}
L289:
	;
	F_errfinish(m, int32(519620), int32(315), int32(23417))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	goto L288
L291:
	;
	v1512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+8)))
	if v1512 != 0 {
		goto L11
	} else {
		goto L292
	}
L292:
	;
	goto L266
L293:
	;
	v1550 = F_palloc(m, int32(24))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L302
	}
L294:
	;
	v1518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+8)))
	if v1518 == int32(0) {
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1522 = v1381 + int32(4)
	v1526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+2)))
	v1527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311))))
	v1528 = int32(16)
	v1530 = v1526 | v1527<<(uint(v1528)%32)
	v1531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1522)+2)))
	v1532 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1522))))
	v1535 = v1531 | v1532<<(uint(v1528)%32)
	if base.Ui32(v1530) < base.Ui32(v1535) {
		v1546 = int32(-1)
		goto L297
	} else {
		goto L298
	}
L296:
	;
	if v1546 < int32(0) {
		goto L12
	} else {
		goto L301
	}
L297:
	;
	goto L296
L298:
	;
	if base.Ui32(v1535) < base.Ui32(v1530) {
		v1546 = int32(1)
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1540 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+4)))
	v1541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1522)+4)))
	if base.Ui32(v1540) < base.Ui32(v1541) {
		v1546 = int32(-1)
		goto L297
	} else {
		goto L300
	}
L300:
	;
	v1546 = base.B2i32(base.Ui32(v1541) < base.Ui32(v1540))
	goto L297
L301:
	;
	goto L293
L302:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v929)))
	v1553 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1550))) = v1552 + v1553
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1550)+4)) = v1556
	v1558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1550)+8)) = uint16(v1558)
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1550)+12)) = v1560
	v1562 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+2)))
	v1563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381))))
	*(*int32)(unsafe.Add(mBase, uint32(v1550)+16)) = v1562 | v1563<<(uint(int32(16))%32)
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v929)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1550)+20)) = v1568
	*(*int32)(unsafe.Add(mBase, uint32(v929)+20)) = v1550
	v1574 = (v1348 + v1553) & int32(65535)
	if base.Ui32(v1574) <= base.Ui32(v1220) {
		v1348 = v1574
		goto L263
	} else {
		goto L303
	}
L303:
	;
	goto L264
L304:
	;
	F_ReleaseBuffer(m, v969)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v929)+20))
	F_pfree(m, v929)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	if v1617 != 0 {
		v929 = v1617
		v948 = v1598
		goto L176
	} else {
		goto L307
	}
L307:
	;
	goto L177
L308:
	;
	goto L153
L309:
	;
	F_pfree(m, v1723)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L1
	} else {
		goto L333
	}
L310:
	;
	v1672 = int32(0)
	goto L329
L311:
	;
	if v876 != 0 {
		goto L314
	} else {
		goto L315
	}
L312:
	;
	goto L313
L313:
	;
	v1661 = v876 * int32(6)
	v1662 = F_palloc(m, v1661)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L1
	} else {
		goto L323
	}
L314:
	;
	v1632 = F_ginPostingListDecode(m, v1627, v39+int32(6240))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L1
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v1658 = F_palloc(m, int32(0))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L1
	} else {
		goto L322
	}
L317:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132])))
	if v1634 == v876 {
		v1669 = v1632
		goto L310
	} else {
		goto L318
	}
L318:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+384)) = v876
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132])))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+388)) = v1641
	F_errmsg_internal(m, int32(480850), v39+int32(384))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(519620), int32(113), int32(370485))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L322:
	;
	v1723 = v1658
	goto L309
L323:
	;
	if v1661 != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	if v876 == int32(0) {
		v1723 = v1662
		goto L309
	} else {
		goto L328
	}
L325:
	;
	v1664 = F__emscripten_memcpy_bulkmem(m, v1662, v1627, v1661)
	mBase = m.M
	goto L327
L326:
	;
	goto L327
L327:
	;
	goto L324
L328:
	;
	v1669 = v1662
	goto L310
L329:
	;
	v1710 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669+v1672*int32(6))+4)))
	if base.Ui32(int32(2048)) <= base.Ui32((v1710-int32(1))&int32(65535)) {
		goto L16
	} else {
		goto L331
	}
L330:
	;
	v1723 = v1669
	goto L309
L331:
	;
	v1718 = v1672 + int32(1)
	if v1718 != v876 {
		v1672 = v1718
		goto L329
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	goto L153
L334:
	;
	v1797 = v379 + int32(1)
	if base.Ui32(v1797&int32(65535)) <= base.Ui32(v357) {
		v366 = v1794
		v378 = v410
		v379 = v1797
		goto L84
	} else {
		goto L335
	}
L335:
	;
	goto L85
L336:
	;
	F_ReleaseBuffer(m, v111)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v1843 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	F_pfree(m, v1843)
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L1
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	F_pfree(m, v74)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L342
	}
L341:
	;
	goto L340
L342:
	;
	if v1842 != 0 {
		v74 = v1842
		v89 = v355
		goto L17
	} else {
		goto L343
	}
L343:
	;
	goto L18
L344:
	;
	m.G0 = v39 + int32(7280)
	return
L345:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+372)) = v1863
	*(*int32)(unsafe.Add(mBase, uint32(v39)+368)) = v1862 + int32(4)
	F_errmsg(m, int32(52056), v39+int32(368))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(519620), int32(636), int32(23383))
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+324)) = v1890
	*(*int32)(unsafe.Add(mBase, uint32(v39)+320)) = v1889 + int32(4)
	F_errmsg(m, int32(51972), v39+int32(320))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(519620), int32(199), int32(23417))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
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
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+156)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v39)+152)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v39)+148)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v39)+144)) = v1916 + int32(4)
	F_errmsg(m, int32(700711), v39+int32(144))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(519620), int32(274), int32(23417))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	v1945 = *(*int64)(unsafe.Add(mBase, uint32(v929)+12))
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+6)))
	v1948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+4)))
	v1949 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = v1949
	v1951 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = v1947 | v1948<<(uint(v1951)%32)
	v1955 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1135]))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = v1955
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v1946 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+124)) = base.I64_rotl(v1945, int64(32))
	v1963 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1136]))))
	v1964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1132]))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+116)) = v1963 | v1964<<(uint(v1951)%32)
	F_errmsg(m, int32(712946), v39+int32(112))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(519620), int32(299), int32(23417))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v1991
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v1990 + int32(4)
	F_errmsg(m, int32(43953), v39+int32(32))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	F_errfinish(m, int32(519620), int32(354), int32(23417))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L365:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	v2018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+6)))
	v2019 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+4)))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v929)+16))
	v2022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v2022
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v2021
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v2020 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v2018 | v2019<<(uint(int32(16))%32)
	F_errmsg(m, int32(701275), v39+int32(48))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(519620), int32(331), int32(23417))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+468)) = v2054
	*(*int32)(unsafe.Add(mBase, uint32(v39)+464)) = v2053 + int32(4)
	F_errmsg(m, int32(51972), v39+int32(464))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(519620), int32(495), int32(23383))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L373:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v2080 + int32(4)
	F_errmsg(m, int32(174408), v39+int32(16))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	F_errfinish(m, int32(519620), int32(710), int32(425470))
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L377:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+516)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+512)) = v2106 + int32(4)
	F_errmsg(m, int32(56925), v39+int32(512))
	mBase = m.M
	v2117 = m.ExcPending
	if v2117 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(519620), int32(699), int32(425470))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v111 < int32(0) {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+548)) = v2151
	*(*int32)(unsafe.Add(mBase, uint32(v39)+544)) = v2132 + int32(4)
	F_errmsg(m, int32(51817), v39+int32(544))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L1
	} else {
		goto L387
	}
L384:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v2136+(v111^int32(-1))<<(uint(int32(6))%32))+16))
	v2151 = v2142
	goto L383
L385:
	;
	goto L386
L386:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v2144+v111<<(uint(int32(6))%32)+int32(-64))+16))
	v2151 = v2150
	goto L383
L387:
	;
	F_errhint(m, int32(608395), int32(0))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(519620), int32(691), int32(425470))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
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
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v111 < int32(0) {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v2201
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v2182 + int32(4)
	F_errmsg(m, int32(51764), v39)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L1
	} else {
		goto L396
	}
L393:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2186+(v111^int32(-1))<<(uint(int32(6))%32))+16))
	v2201 = v2192
	goto L392
L394:
	;
	goto L395
L395:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v2194+v111<<(uint(int32(6))%32)+int32(-64))+16))
	v2201 = v2200
	goto L392
L396:
	;
	F_errhint(m, int32(608395), int32(0))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(519620), int32(680), int32(425470))
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
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
	F_errmsg_internal(m, int32(503985), v10)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(519680), int32(1266), int32(336807))
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
	v349 = F_pg_snprintf(m, v13+int32(-10), int32(10), int32(30746), v13+int32(-32))
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
	v726 = F_pg_snprintf(m, v13+int32(-10), int32(10), int32(30746), v15)
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
	F_errmsg_internal(m, int32(503985), v13+int32(-48))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(519680), int32(921), int32(16134))
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
	v5 = F_gin_btree_extract_query(m, l0, int32(1), int32(7616), int32(2661))
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
	F_amcheck_lock_relation_and_check(m, v2, int32(2742), int32(7639), int32(1), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
