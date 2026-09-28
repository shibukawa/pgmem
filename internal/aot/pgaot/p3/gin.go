package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinBufferStoreTuple(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
	if v13 != 0 {
		v27 = int64(0)
	} else {
		v15 = l1 + int32(24)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
		if v16 == int32(1) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v19 != 0 {
				base.MemoryCopy(m, v10+int32(8), v15, v19)
			} else {
			}
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
			v27 = v23
		} else {
			v27 = base.I64_extend_i32_u(v15)
		}
	}
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v32 = (v28 + int32(25)) & int32(-2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v38 = F_ginPostingListDecodeAllSegments(m, l1+v32, v34-v32, v10+int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v40 == int32(0) {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v43)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v45
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v47)
			v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+12)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v49)
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)) = uint8(v51)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
			if v53 != 0 {
				v59 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v59
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v67 = (v63 + v64) * int32(6)
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v68 == int32(0) {
					v71 = F_palloc(m, v67)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v75 = v71
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v75
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v86 = F_ginMergeItemPointers(m, v75+v77*int32(6), v81-v77, v38, v83, v10+int32(8))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							v90 = v88 * int32(6)
							if v90 != 0 {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								base.MemoryCopy(m, v91+v92*int32(6), v86, v90)
							} else {
							}
							F_pfree(m, v86)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v100
								F_pfree(m, v38)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v73 = F_repalloc(m, v68, v67)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						v75 = v73
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v75
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v86 = F_ginMergeItemPointers(m, v75+v77*int32(6), v81-v77, v38, v83, v10+int32(8))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							v90 = v88 * int32(6)
							if v90 != 0 {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								base.MemoryCopy(m, v91+v92*int32(6), v86, v90)
							} else {
							}
							F_pfree(m, v86)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v100
								F_pfree(m, v38)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v57 = F_datumCopy(m, v27, v51&int32(1), v49)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v59 = v57
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v59
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v67 = (v63 + v64) * int32(6)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v68 == int32(0) {
						v71 = F_palloc(m, v67)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v75 = v71
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v75
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v86 = F_ginMergeItemPointers(m, v75+v77*int32(6), v81-v77, v38, v83, v10+int32(8))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v90 = v88 * int32(6)
								if v90 != 0 {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									base.MemoryCopy(m, v91+v92*int32(6), v86, v90)
								} else {
								}
								F_pfree(m, v86)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v100
									F_pfree(m, v38)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v73 = F_repalloc(m, v68, v67)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v75 = v73
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v75
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v86 = F_ginMergeItemPointers(m, v75+v77*int32(6), v81-v77, v38, v83, v10+int32(8))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v90 = v88 * int32(6)
								if v90 != 0 {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									base.MemoryCopy(m, v91+v92*int32(6), v86, v90)
								} else {
								}
								F_pfree(m, v86)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v100
									F_pfree(m, v38)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v67 = (v63 + v64) * int32(6)
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v68 == int32(0) {
				v71 = F_palloc(m, v67)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					v75 = v71
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v75
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v86 = F_ginMergeItemPointers(m, v75+v77*int32(6), v81-v77, v38, v83, v10+int32(8))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						v90 = v88 * int32(6)
						if v90 != 0 {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							base.MemoryCopy(m, v91+v92*int32(6), v86, v90)
						} else {
						}
						F_pfree(m, v86)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v100
							F_pfree(m, v38)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v73 = F_repalloc(m, v68, v67)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					v75 = v73
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v75
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v86 = F_ginMergeItemPointers(m, v75+v77*int32(6), v81-v77, v38, v83, v10+int32(8))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						v90 = v88 * int32(6)
						if v90 != 0 {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							base.MemoryCopy(m, v91+v92*int32(6), v86, v90)
						} else {
						}
						F_pfree(m, v86)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v100
							F_pfree(m, v38)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F__gin_build_tuple(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	v1 = l0
	v2 = l1
	v4 = l3
	v5 = l4
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if v2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v77 = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v77
	if l6 != 0 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v21 = v9
	goto L4
L3:
	;
	v21 = int32(8)
	goto L4
L4:
	;
	if v2|v5 != 0 {
		v75 = v21
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(0) < v4 {
		v75 = v4
		goto L1
	} else {
		goto L6
	}
L6:
	;
	switch v4 + int32(2) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L8
	}
L7:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v75 = int32(base.Ui32(v70) >> (uint(int32(2)) % 32))
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v52 = F_strlen(m, base.I32_wrap_i64(l2))
	mBase = m.M
	v75 = v52 + int32(1)
	goto L1
L10:
	;
	v27 = base.I32_wrap_i64(l2)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(18)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if v34 == v32 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v28&int32(1) == int32(0) {
		goto L7
	} else {
		goto L20
	}
L14:
	;
	v37 = v32
	goto L16
L15:
	;
	v37 = int32(2)
	goto L16
L16:
	;
	if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = int32(6)
	goto L19
L18:
	;
	v44 = v37
	goto L19
L19:
	;
	v75 = v44
	goto L1
L20:
	;
	v75 = int32(base.Ui32(v28) >> (uint(int32(1)) % 32))
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v4
	F_errmsg_internal(m, int32(_a_F__gin_build_tuple_0), v17)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F__gin_build_tuple_1), int32(2284), int32(_a_F__gin_build_tuple_2))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v89 = int32(0)
	v94 = v9
	goto L28
L26:
	;
	v145 = v9
	goto L27
L27:
	;
	v149 = (v75 + int32(25)) & int32(-2)
	v150 = v145 + v149
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v150
	v152 = F_palloc0(m, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L21
	} else {
		goto L37
	}
L28:
	;
	v96 = F_palloc(m, int32(12))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L21
	} else {
		goto L30
	}
L29:
	;
	v145 = v117
	goto L27
L30:
	;
	v105 = F_ginCompressPostingList(m, l5+v89*int32(6), l6-v89, int32(_a_F__gin_build_tuple_3), v17+int32(4))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v105
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v109 = v108 + v89
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
	v117 = v94 + (v110+int32(1))&int32(_a_F__gin_build_tuple_4) + int32(8)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v118 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v96
	if base.Ui32(v109) < base.Ui32(l6) {
		v89 = v109
		v94 = v117
		goto L28
	} else {
		goto L36
	}
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v124 = v119
	goto L32
L34:
	;
	goto L35
L35:
	;
	v121 = v17 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v121
	v124 = v121
	goto L32
L36:
	;
	goto L29
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+15)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+4)) = uint16(v1)
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v152)+16)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v152)+8)) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+14)) = uint8(v5)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+12)) = uint16(v4)
	if v2 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v185 = int32(0)
	if base.B2i32(v184 == v185)|base.B2i32(v184 == v17+int32(8)) == v185 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	if v5 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v152)+24)) = l2
	goto L38
L41:
	;
	goto L42
L42:
	;
	if int32(0) < v4 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v4 == int32(0) {
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	switch v4 + int32(2) {
	case 0:
		goto L47
	case 1:
		goto L48
	default:
		goto L38
	}
L46:
	;
	base.MemoryCopy(m, v152+int32(24), base.I32_wrap_i64(l2), v4)
	goto L38
L47:
	;
	if v75 == int32(0) {
		goto L38
	} else {
		goto L50
	}
L48:
	;
	if v75 == int32(0) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	base.MemoryCopy(m, v152+int32(24), base.I32_wrap_i64(l2), v75)
	goto L38
L50:
	;
	base.MemoryCopy(m, v152+int32(24), base.I32_wrap_i64(l2), v75)
	goto L38
L51:
	;
	v195 = v152 + v149
	v202 = v184
	goto L54
L52:
	;
	goto L53
L53:
	;
	m.G0 = v17 + int32(16)
	return v152
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+6)))
	v216 = (v210+int32(1))&int32(_a_F__gin_build_tuple_4) + int32(8)
	if v216 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	base.MemoryCopy(m, v195, v209, v216)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218)+6)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	F_pfree(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L21
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v202)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	v235 = int32(8)
	if v208 != v17+v235 {
		v195 = v195 + (v219+int32(1))&int32(_a_F__gin_build_tuple_4) + v235
		v202 = v208
		goto L54
	} else {
		goto L61
	}
L61:
	;
	goto L55
}
func F_ginGetStats(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v5 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_LockBufferInternal(m, v5, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			if v5 < int32(0) {
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_ginGetStats[0]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(v5^int32(-1))<<(uint(int32(2))%32))))
				v27 = v19
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ginGetStats[1]))
				v27 = v21 + v5<<(uint(int32(13))%32) + int32(-8192)
			}
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v34
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v27)+64))
			*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v38
			F_UnlockReleaseBuffer(m, v5)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ginInitBA(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	v8 = F_palloc(m, int32(28))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(36)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(35)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(34)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(48)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_ginInitBA_0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
		return
	}
}
func F_ginMergeItemPointers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
	v12 = l1 + l3
	v15 = F_palloc(m, v12*int32(6))
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
	if base.B2i32(l1 == v19)|base.B2i32(l3 == v19) == v19 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v257
	return v15
L4:
	;
	v66 = int32(6)
	v68 = l2 + l3*v66
	v71 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v68-int32(4)))))
	v72 = int64(32)
	v76 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v68-v66))))
	v77 = int64(48)
	v82 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v68-int32(2)))))
	v84 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v88 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui64(v84|(v85<<(uint(v72)%64)|v88<<(uint(v77)%64))) <= base.Ui64(v71<<(uint(v72)%64)|v76<<(uint(v77)%64)|v82) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v26 = int32(6)
	v28 = l0 + l1*v26
	v31 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v28-int32(4)))))
	v32 = int64(32)
	v36 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v28-v26))))
	v37 = int64(48)
	v42 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v28-int32(2)))))
	v44 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v45 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v48 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui64(v44|(v45<<(uint(v32)%64)|v48<<(uint(v37)%64))) <= base.Ui64(v31<<(uint(v32)%64)|v36<<(uint(v37)%64)|v42) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v56 = l1 * int32(6)
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	base.MemoryCopy(m, v15, l0, v56)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v59 = l3 * int32(6)
	if v59 == int32(0) {
		v257 = v12
		goto L3
	} else {
		goto L12
	}
L12:
	;
	base.MemoryCopy(m, v56+v15, l2, v59)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v12
	return v15
L13:
	;
	v99 = v15
	v100 = l0
	v101 = l2
	goto L16
L14:
	;
	goto L15
L15:
	;
	v242 = l3 * int32(6)
	if v242 != 0 {
		goto L41
	} else {
		goto L42
	}
L16:
	;
	v107 = base.I32_div_s(v101-l2, int32(6))
	if base.Ui32(v107) < base.Ui32(l3) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v164 = base.I32_div_s(v158-l0, int32(6))
	if base.Ui32(v164) < base.Ui32(l1) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v109 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
	v110 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
	v111 = int64(32)
	v113 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v100))))
	v114 = int64(48)
	v117 = v109 | (v110<<(uint(v111)%64) | v113<<(uint(v114)%64))
	v118 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	v119 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v101)+2)))
	v122 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v101))))
	v126 = v118 | (v119<<(uint(v111)%64) | v122<<(uint(v114)%64))
	if base.Ui64(v126) < base.Ui64(v117) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v157 = v99
	v158 = v100
	v159 = v101
	goto L20
L20:
	;
	goto L17
L21:
	;
	v151 = int32(6)
	v152 = v99 + v151
	v155 = base.I32_div_s(v149-l0, v151)
	if base.Ui32(v155) < base.Ui32(l1) {
		v99 = v152
		v100 = v149
		v101 = v150
		goto L16
	} else {
		goto L28
	}
L22:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)) = uint16(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v130
	v149 = v100
	v150 = v101 + int32(6)
	goto L21
L23:
	;
	goto L24
L24:
	;
	if v117 == v126 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)) = uint16(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v137
	v139 = int32(6)
	v149 = v100 + v139
	v150 = v101 + v139
	goto L21
L26:
	;
	goto L27
L27:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)) = uint16(v143)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v145
	v149 = v100 + int32(6)
	v150 = v101
	goto L21
L28:
	;
	v157 = v152
	v158 = v149
	v159 = v150
	goto L20
L29:
	;
	v171 = v157
	v172 = v158
	goto L32
L30:
	;
	v194 = v157
	goto L31
L31:
	;
	v202 = base.I32_div_s(v159-l2, int32(6))
	if base.Ui32(v202) < base.Ui32(l3) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)) = uint16(v177)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v179
	v181 = int32(6)
	v182 = v171 + v181
	v184 = v172 + v181
	v187 = base.I32_div_s(v184-l0, v181)
	if base.Ui32(v187) < base.Ui32(l1) {
		v171 = v182
		v172 = v184
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v194 = v182
	goto L31
L34:
	;
	goto L33
L35:
	;
	v209 = v194
	v211 = v159
	goto L38
L36:
	;
	v232 = v194
	goto L37
L37:
	;
	v240 = base.I32_div_s(v232-v15, int32(6))
	v257 = v240
	goto L3
L38:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v209)+4)) = uint16(v215)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v217
	v219 = int32(6)
	v220 = v209 + v219
	v222 = v211 + v219
	v225 = base.I32_div_s(v222-l2, v219)
	if base.Ui32(v225) < base.Ui32(l3) {
		v209 = v220
		v211 = v222
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v232 = v220
	goto L37
L40:
	;
	goto L39
L41:
	;
	base.MemoryCopy(m, v15, l2, v242)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v245 = l1 * int32(6)
	if v245 == int32(0) {
		v257 = v12
		goto L3
	} else {
		goto L44
	}
L44:
	;
	base.MemoryCopy(m, v242+v15, l0, v245)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v12
	return v15
}
func F_ginRedoRecompress(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
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
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v320 int32
	_ = v320
	v3 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v26 = l0 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)))
	if v27&int32(128) == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v72 = l0 + int32(32)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v73 == int32(0) {
		v302 = v72
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v34 = l0 + int32(32)
	v38 = F_ginCompressPostingList(m, v34, v32, int32(_a_F_ginRedoRecompress_0), v23+int32(12))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v52 = v25
	v56 = int32(32)
	goto L6
L6:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v56)
	v58 = l0 + v52
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
	v61 = v59 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)) = uint16(v61)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v65 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v63)+4)) = uint16(v65)
	goto L3
L7:
	;
	return
L8:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
	v44 = (v40 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	v46 = v44 + int32(8)
	if v46 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	base.MemoryCopy(m, v34, v38, v46)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v52 = v48
	v56 = v44 + int32(40)
	goto L6
L12:
	;
	v320 = v302 - v72 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v320)
	m.G0 = v23 + int32(16)
	return
L13:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v84 = v72
	v86 = v72
	v91 = v3
	v92 = v3
	v93 = l1 + int32(2)
	v95 = v72 + v78 - int32(32)
	v101 = v3
	goto L14
L14:
	;
	v102 = int32(2)
	v103 = v93 + v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v104&int32(254) == v102 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	if base.B2i32(v257 == int32(0))|base.B2i32(v284 == v258) != 0 {
		v302 = v285
		goto L12
	} else {
		goto L67
	}
L16:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if base.B2i32(v150 <= v91) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v135 = v93 + int32(4)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
	v140 = int32(0)
	v144 = v140
	v145 = v140
	v146 = v135 + v136*int32(6)
	v147 = int32(1)
	v148 = v135
	v149 = v136
	goto L16
L18:
	;
	v132 = int32(0)
	v144 = v128
	v145 = v129
	v146 = v131
	v147 = v130
	v148 = v132
	v149 = v132
	goto L16
L19:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+8)))
	v113 = (v109 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	v128 = v113 + int32(8)
	v129 = v103
	v130 = int32(0)
	v131 = v103 + (v113+int32(9))&int32(_a_F_ginRedoRecompress_2)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v104 == int32(4) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v124 = int32(0)
	v128 = v124
	v129 = v124
	v130 = v124
	v131 = v103
	goto L18
L23:
	;
	v156 = v84
	v158 = v86
	v163 = v91
	goto L26
L24:
	;
	v197 = v84
	v199 = v86
	v204 = v91
	goto L25
L25:
	;
	if v147 != 0 {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
	v178 = (v174 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	v180 = v178 + int32(8)
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v197 = v191
	v199 = v188
	v204 = v150
	goto L25
L28:
	;
	if v180 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v187 = v178
	goto L30
L30:
	;
	v188 = v158 + v180
	v191 = v156 + v187 + int32(8)
	v193 = v163 + int32(1)
	if v193 != v150 {
		v156 = v191
		v158 = v188
		v163 = v193
		goto L26
	} else {
		goto L34
	}
L31:
	;
	base.MemoryCopy(m, v158, v156, v180)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
	v187 = (v182 + int32(1)) & int32(_a_F_ginRedoRecompress_1)
	goto L30
L34:
	;
	goto L27
L35:
	;
	v217 = F_ginPostingListDecode(m, v197, v23+int32(12))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	v238 = v144
	v239 = v145
	v240 = v104
	goto L37
L37:
	;
	if v197 == v95 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v222 = F_ginMergeItemPointers(m, v148, v149, v217, v219, v23+int32(8))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v228 = F_ginCompressPostingList(m, v222, v224, int32(_a_F_ginRedoRecompress_0), v23+int32(4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+6)))
	v238 = (v230+int32(1))&int32(_a_F_ginRedoRecompress_1) + int32(8)
	v239 = v228
	v240 = int32(3)
	goto L37
L41:
	;
	switch v240 - int32(1) {
	case 0:
		goto L53
	case 1:
		goto L56
	case 2:
		goto L55
	default:
		goto L54
	}
L42:
	;
	v255 = v197
	v256 = int32(0)
	v257 = v92
	v258 = v95
	goto L41
L43:
	;
	goto L44
L44:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+6)))
	v249 = (v243+int32(1))&int32(_a_F_ginRedoRecompress_1) + int32(8)
	if v92 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v255 = v197
	v256 = v249
	v257 = v92
	v258 = v95
	goto L41
L46:
	;
	goto L47
L47:
	;
	v250 = v95 - v197
	v251 = F_palloc(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	if v250 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	base.MemoryCopy(m, v251, v197, v250)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v255 = v251
	v256 = v249
	v257 = v251
	v258 = v251 + v250
	goto L41
L52:
	;
	v288 = v101 + int32(1)
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if base.Ui32(v288) < base.Ui32(v289) {
		v84 = v284
		v86 = v285
		v91 = v286
		v92 = v257
		v93 = v146
		v95 = v258
		v101 = v288
		goto L14
	} else {
		goto L66
	}
L53:
	;
	v284 = v255 + v256
	v285 = v199
	v286 = v204 + int32(1)
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L63
	}
L55:
	;
	if v238 != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	if v238 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryCopy(m, v199, v239, v238)
	goto L59
L58:
	;
	goto L59
L59:
	;
	v284 = v255
	v285 = v199 + v238
	v286 = v204
	goto L52
L60:
	;
	base.MemoryCopy(m, v199, v239, v238)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v284 = v255 + v256
	v285 = v199 + v238
	v286 = v204 + int32(1)
	goto L52
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v240
	F_errmsg_internal(m, int32(_a_F_ginRedoRecompress_3), v23)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_ginRedoRecompress_4), int32(298), int32(_a_F_ginRedoRecompress_5))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	goto L15
L67:
	;
	v295 = v258 - v284
	if v295 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	base.MemoryCopy(m, v285, v284, v295)
	goto L70
L69:
	;
	goto L70
L70:
	;
	v302 = v295 + v285
	goto L12
}
func F_gin_enum_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v12 int64
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = base.I32_wrap_i64(v5)
	v7 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	if v7 == int64(0) {
		if v6 != 0 {
			v12 = int64(-1)
		} else {
			v12 = int64(0)
		}
		return v12
	} else {
		if v6 == int32(0) {
			return int64(1)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = F_CallerFInfoFunctionCall2(m, int32(4010), v19, v20, v7, v5&int64(4294967295))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				return base.I64_extend32_s(v23)
			}
		}
	}
}
func F_gin_extract_hstore_query(m *base.Module, l0 int32) int64 {
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
	var v40 int32
	_ = v40
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = base.I32_wrap_i64(v19)
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v22 = base.I32_wrap_i64(v21)
	switch v22&int32(_a_F_gin_extract_hstore_query_0) - int32(7) {
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
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return base.I64_extend_i32_u(v218)
L3:
	;
	if v22&int32(_a_F_gin_extract_hstore_query_1) != int32(10) {
		goto L1
	} else {
		goto L29
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v41 = F_pg_detoast_datum_packed(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = F_DirectFunctionCall2Coll(m, int32(_a_F_gin_extract_hstore_query_2), int32(0), v29, v19&int64(4294967295))
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
	if v36 != 0 {
		v218 = v36
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	v218 = int32(0)
	goto L2
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	v46 = F_palloc(m, int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v49 = int32(1)
	v50 = v48 & v49
	if v48 == v49 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v79 = v77 + int32(5)
	v80 = F_palloc(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L22
	}
L12:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v56 == int32(18) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v67 = int32(1)
	if v50 != 0 {
		v77 = int32(base.Ui32(v48)>>(uint(v67)%32)) - v67
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v59 = int32(16)
	goto L17
L16:
	;
	v59 = int32(0)
	goto L17
L17:
	;
	if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v66 = int32(4)
	goto L20
L19:
	;
	v66 = v59
	goto L20
L20:
	;
	v77 = v66
	goto L11
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v77 = int32(base.Ui32(v71)>>(uint(int32(2))%32)) - int32(4)
	goto L11
L22:
	;
	v82 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)) = uint8(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v79 << (uint(int32(2)) % 32)
	v87 = int32(0)
	if base.B2i32(v77 == v87)|base.B2i32(v77 <= v87) == v87 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v50 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = base.I64_extend_i32_u(v80)
	v218 = v46
	goto L2
L26:
	;
	v98 = int32(1)
	goto L28
L27:
	;
	v98 = int32(4)
	goto L28
L28:
	;
	base.MemoryCopy(m, v80+int32(5), v41+v98, v77)
	goto L25
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v108 = F_pg_detoast_datum(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_deconstruct_array_builtin(m, v108, int32(25), v16+int32(12), v16+int32(8), v16+int32(4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v122 = F_palloc(m, v119<<(uint(int32(3))%32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v124 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v128 = int32(0)
	v129 = v2
	v131 = v124
	goto L36
L34:
	;
	v193 = v2
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v193
	if base.B2i32(v22&int32(_a_F_gin_extract_hstore_query_0) != int32(11))|v193 != 0 {
		v218 = v122
		goto L2
	} else {
		goto L46
	}
L36:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v128))))
	if v143 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v193 = v183
	goto L35
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v128<<(uint(int32(3))%32))))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v153 = int32(base.Ui32(v151) >> (uint(int32(2)) % 32))
	v155 = v153 + int32(1)
	v156 = F_palloc(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	v183 = v129
	v184 = v131
	goto L40
L40:
	;
	v190 = v128 + int32(1)
	if v190 < v184 {
		v128 = v190
		v129 = v183
		v131 = v184
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v158 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+4)) = uint8(v158)
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v155 << (uint(int32(2)) % 32)
	if base.Ui32(v151) < base.Ui32(int32(20)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v122+v129<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v156)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v183 = v129 + int32(1)
	v184 = v180
	goto L40
L43:
	;
	v166 = v153 - int32(4)
	if v166 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	base.MemoryCopy(m, v156+int32(5), v150+int32(4), v166)
	goto L42
L45:
	;
	goto L37
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	v218 = v122
	goto L2
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v22 & int32(_a_F_gin_extract_hstore_query_0)
	F_errmsg_internal(m, int32(_a_F_gin_extract_hstore_query_3), v16)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_gin_extract_hstore_query_4), int32(141), int32(_a_F_gin_extract_hstore_query_5))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_query_anyenum(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_anyenum_0), int32(_a_F_gin_extract_query_anyenum_1), int32(0), int32(_a_F_gin_extract_query_anyenum_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_bytea(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_bytea_0), int32(_a_F_gin_extract_query_bytea_1), int32(0), int32(_a_F_gin_extract_query_bytea_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_int4(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_int4_0), int32(_a_F_gin_extract_query_int4_1), int32(_a_F_gin_extract_query_int4_2), int32(_a_F_gin_extract_query_int4_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_time(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_time_0), int32(_a_F_gin_extract_query_time_1), int32(0), int32(_a_F_gin_extract_query_time_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_trgm(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	switch v2 - int32(3) {
	case 0:
		v23 = F_gin_extract_value_trgm(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			return v23
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_gin_extract_trgm_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_extract_trgm_1), int32(30), int32(_a_F_gin_extract_trgm_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v5 = F_gin_extract_query_trgm(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			return v5
		}
	}
}
func F_gin_extract_value_trgm(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v98 int64
	_ = v98
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
		v13 = int32(1)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		v17 = v15 & v13
		if v17 != 0 {
			v18 = v13
		} else {
			v18 = int32(4)
		}
		if v15 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = F_generate_trgm(m, v6+v18, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
			v53 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(5)
			if base.Ui32(int32(3)) <= base.Ui32(v53) {
				v57 = base.I32_div_u_s(v53, int32(3))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v57
				v59 = int32(1)
				if base.Ui32(v57) <= base.Ui32(v59) {
					v62 = v59
				} else {
					v62 = v57
				}
				v67 = F_palloc_mul(m, int32(8), v57)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int64(0)
				} else {
					v69 = int32(0)
					v71 = v47 + int32(5)
					for {
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
						*(*int64)(unsafe.Add(mBase, uint32(v67+v69<<(uint(int32(3))%32)))) = base.I64_extend_i32_s(v76 | (v77<<(uint(int32(8))%32) | v80<<(uint(int32(16))%32)))
						v90 = v69 + int32(1)
						if v90 != v62 {
							v69 = v90
							v71 = v71 + int32(3)
							continue
						} else {
							break
						}
						break
					}
					v98 = base.I64_extend_i32_u(v67)
					return v98
				}
			} else {
				v98 = int64(0)
				return v98
			}
		}
	}
}
func F_gin_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = l0 - int32(16)
	if base.Ui32(v3) <= base.Ui32(int32(143)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v3)>>(uint(int32(2))%32))&int32(1073741820))+uint32(_c_F_gin_identify[0])))
		v12 = v10
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_gin_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v12 = v10 & int32(_a_F_gin_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v12)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v7)+6)))
	if v15&int32(4) != 0 {
		v20 = int32(0)
		base.MemoryFill(m, l0+int32(24), v20, int32(_a_F_gin_mask_1))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
		return
	} else {
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if base.Ui32(int32(25)) <= base.Ui32(v25) {
			F_mask_unused_space(m, l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_gin_redo(m *base.Module, l0 int32) {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int64
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int64
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int64
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int64
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int64
	_ = v790
	var v804 int32
	_ = v804
	var v806 int64
	_ = v806
	var v808 int64
	_ = v808
	var v810 int64
	_ = v810
	var v812 int64
	_ = v812
	var v814 int64
	_ = v814
	var v816 int64
	_ = v816
	var v818 int64
	_ = v818
	var v821 int64
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int64
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int64
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1206 int64
	_ = v1206
	var v1220 int32
	_ = v1220
	var v1222 int64
	_ = v1222
	var v1224 int64
	_ = v1224
	var v1226 int64
	_ = v1226
	var v1228 int64
	_ = v1228
	var v1230 int64
	_ = v1230
	var v1232 int64
	_ = v1232
	var v1234 int64
	_ = v1234
	var v1237 int64
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1324 int32
	_ = v1324
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = int32(_a_F_gin_redo_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+48)))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_gin_redo[0])) = v21
	v24 = v18 & int32(240)
	switch int32(base.Ui32(v24-int32(16)) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L12
	case 5:
		goto L11
	case 6:
		goto L10
	case 7:
		goto L9
	case 8:
		goto L13
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L19
	} else {
		goto L346
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L19
	} else {
		goto L343
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L19
	} else {
		goto L340
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L19
	} else {
		goto L337
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L19
	} else {
		goto L334
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L19
	} else {
		goto L331
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L19
	} else {
		goto L328
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gin_redo[0])) = v16
	v1338 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[1]))
	F_MemoryContextReset(m, v1338)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L19
	} else {
		goto L327
	}
L9:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v1156 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1158 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L19
	} else {
		goto L298
	}
L10:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v1007 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1009 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L19
	} else {
		goto L264
	}
L11:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v740 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v742 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L19
	} else {
		goto L208
	}
L12:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v583 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v587 = F_XLogReadBufferForRedo(m, l0, int32(2), v11+int32(-8))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L19
	} else {
		goto L163
	}
L13:
	;
	v510 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v514 = F_XLogReadBufferForRedo(m, l0, int32(0), v11+int32(-20))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L19
	} else {
		goto L140
	}
L14:
	;
	v503 = F_XLogReadBufferForRedo(m, l0, int32(0), v11+int32(-20))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L19
	} else {
		goto L137
	}
L15:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+24)))
	if v416&int32(2) != 0 {
		goto L112
	} else {
		goto L113
	}
L16:
	;
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100))))
	v103 = v101 & int32(2)
	if v103 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v32 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v52 = int32(131)
	if v32 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	return
L20:
	;
	if v32 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+(v32^int32(-1))<<(uint(int32(2))%32))))
	v51 = v43
	goto L18
L22:
	;
	goto L23
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v51 = v45 + v32<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v79 != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	F_PageInit(m, v70, int32(_a_F_gin_redo_1), int32(8))
	mBase = m.M
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
	v75 = v70 + v74
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+6)) = uint16(v52)
	goto L24
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v32^int32(-1))<<(uint(int32(2))%32))))
	v70 = v62
	goto L25
L27:
	;
	goto L28
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v70 = v64 + v32<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+64))
	base.MemoryCopy(m, v51+int32(32), v83+int32(4), v79)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29))))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = base.I64_rotl(v30, int64(32))
	v92 = v87 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+12)) = uint16(v92)
	F_MarkBufferDirty(m, v32)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v32)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L8
L34:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
	v111 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(-20))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L19
	} else {
		goto L37
	}
L35:
	;
	v155 = int32(-1)
	goto L36
L36:
	;
	v161 = F_XLogReadBufferForRedo(m, l0, int32(0), v11+int32(-24))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L19
	} else {
		goto L50
	}
L37:
	;
	if v111 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v115 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v148 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)))
	v135 = v134 + v133
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+6)))
	v138 = v136 & int32(_a_F_gin_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v135)+6)) = uint16(v138)
	*(*int64)(unsafe.Add(mBase, uint32(v133))) = base.I64_rotl(v98, int64(32))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_MarkBufferDirty(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L19
	} else {
		goto L45
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+(v115^int32(-1))<<(uint(int32(2))%32))))
	v133 = v125
	goto L41
L43:
	;
	goto L44
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v133 = v127 + v115<<(uint(int32(13))%32) + int32(-8192)
	goto L41
L45:
	;
	goto L40
L46:
	;
	F_UnlockReleaseBuffer(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L19
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v155 = v106<<(uint(int32(16))%32) | v107
	goto L36
L49:
	;
	goto L48
L50:
	;
	if v161 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v165 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v410 == int32(0) {
		goto L8
	} else {
		goto L110
	}
L54:
	;
	v184 = int32(0)
	v186 = v11 + int32(-28)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+72))
	if v189 < v184 {
		v211 = v184
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169+(v165^int32(-1))<<(uint(int32(2))%32))))
	v183 = v175
	goto L54
L56:
	;
	goto L57
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v183 = v177 + v165<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L58:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v216&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L59:
	;
	v214 = v211
	goto L58
L60:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+int32(0))+76)))
	if v194 != int32(1) {
		v211 = v184
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v198 = v188 + int32(76)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+43)))
	if v199 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v186 == int32(0) {
		v211 = v184
		goto L59
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v186 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v204 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v204
	v214 = v204
	goto L58
L66:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v207
	goto L68
L67:
	;
	goto L68
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+44))
	v211 = v209
	goto L59
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v183))) = base.I64_rotl(v98, int64(32))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	F_MarkBufferDirty(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L19
	} else {
		goto L109
	}
L70:
	;
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214))))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v339*int32(10))+22)) = base.I32_rotr(v155, int32(16))
	v347 = v214 + int32(2)
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236)+16)))
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236+v348)+4)))
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214))))
	if v351 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L71:
	;
	if v215 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	goto L73
L73:
	;
	if v215 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	if v103 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L75:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v222+(v215^int32(-1))<<(uint(int32(2))%32))))
	v236 = v228
	goto L74
L76:
	;
	goto L77
L77:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v236 = v230 + v215<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	F_ginRedoRecompress(m, v236, v214)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	goto L69
L80:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214))))
	if v155 != int32(-1) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244+(v215^int32(-1))<<(uint(int32(2))%32))))
	v258 = v250
	goto L80
L82:
	;
	goto L83
L83:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v258 = v252 + v215<<(uint(int32(13))%32) + int32(-8192)
	goto L80
L84:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v258+v259<<(uint(int32(2))%32))+20))
	v268 = v258 + v265&int32(_a_F_gin_redo_3)
	v269 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+4)) = uint16(v269)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+2)) = uint16(v155)
	v273 = int32(base.Ui32(v155) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v268))) = uint16(v273)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+2)))
	if v276 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_PageIndexTupleDelete(m, v258, v259)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L19
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+10)))
	v287 = F_PageAddItemExtended(m, v258, v214+int32(4), v283&int32(_a_F_gin_redo_4), v259, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L19
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	if v287 != 0 {
		goto L69
	} else {
		goto L92
	}
L92:
	;
	v290 = v11 + int32(-20)
	if v215 < int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L19
	} else {
		goto L98
	}
L94:
	;
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+8)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v290))) = v313
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(-4)))) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(-8)))) = v319
	goto L93
L95:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[4]))
	v312 = v299 + (v215^int32(-1))*int32(56)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[5]))
	v307 = int32(56)
	v312 = v306 + v215*v307 - v307
	goto L94
L98:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+20)) = v327
	F_errmsg_internal(m, int32(_a_F_gin_redo_5), v11+int32(-48))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L19
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(104), int32(_a_F_gin_redo_7))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L19
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v377)+8)) = uint16(v379)
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
	*(*int64)(unsafe.Add(mBase, uint32(v377))) = v381
	v385 = v350 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v236+v378)+4)) = uint16(v385)
	v390 = v350*int32(10) + int32(42)
	*(*uint16)(unsafe.Add(mBase, uint32(v236)+12)) = uint16(v390)
	goto L69
L102:
	;
	v377 = v236 + v350*int32(10) + int32(32)
	v378 = v348
	goto L101
L103:
	;
	goto L104
L104:
	;
	v361 = v236 + v351*int32(10)
	v363 = v361 + int32(22)
	if v350+int32(1) == v351 {
		v377 = v363
		v378 = v348
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v368 = int32(10)
	v371 = (v350-v351)*v368 + v368
	if v371 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	base.MemoryCopy(m, v361+int32(32), v363, v371)
	goto L108
L107:
	;
	goto L108
L108:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236)+16)))
	v377 = v363
	v378 = v375
	goto L101
L109:
	;
	goto L53
L110:
	;
	F_UnlockReleaseBuffer(m, v410)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	goto L8
L112:
	;
	v471 = F_XLogReadBufferForRedo(m, l0, int32(0), v11+int32(-20))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L19
	} else {
		goto L125
	}
L113:
	;
	v419 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v423 = F_XLogReadBufferForRedo(m, l0, int32(3), v11+int32(-20))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L19
	} else {
		goto L114
	}
L114:
	;
	if v423 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v427 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L117
L117:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v460 == int32(0) {
		goto L112
	} else {
		goto L123
	}
L118:
	;
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v445)+16)))
	v447 = v446 + v445
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v447)+6)))
	v450 = v448 & int32(_a_F_gin_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v447)+6)) = uint16(v450)
	*(*int64)(unsafe.Add(mBase, uint32(v445))) = base.I64_rotl(v419, int64(32))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_MarkBufferDirty(m, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L19
	} else {
		goto L122
	}
L119:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v431+(v427^int32(-1))<<(uint(int32(2))%32))))
	v445 = v437
	goto L118
L120:
	;
	goto L121
L121:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v445 = v439 + v427<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L122:
	;
	goto L117
L123:
	;
	F_UnlockReleaseBuffer(m, v460)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L19
	} else {
		goto L124
	}
L124:
	;
	goto L112
L125:
	;
	if v471 != int32(2) {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v478 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(-4))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L19
	} else {
		goto L127
	}
L127:
	;
	if v478 != int32(2) {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	if v416&int32(4) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v487 = F_XLogReadBufferForRedo(m, l0, int32(2), v11+int32(-8))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L19
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	F_UnlockReleaseBuffer(m, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L19
	} else {
		goto L135
	}
L132:
	;
	if v487 != int32(2) {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	F_UnlockReleaseBuffer(m, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L19
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_UnlockReleaseBuffer(m, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	goto L8
L137:
	;
	if v503 != int32(2) {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_UnlockReleaseBuffer(m, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L19
	} else {
		goto L139
	}
L139:
	;
	goto L8
L140:
	;
	if v514 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v518 < int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	goto L143
L143:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v577 == int32(0) {
		goto L8
	} else {
		goto L161
	}
L144:
	;
	v537 = int32(0)
	v539 = v11 + int32(-4)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+72))
	if v542 < v537 {
		v564 = v537
		goto L149
	} else {
		goto L150
	}
L145:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v522+(v518^int32(-1))<<(uint(int32(2))%32))))
	v536 = v528
	goto L144
L146:
	;
	goto L147
L147:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v536 = v530 + v518<<(uint(int32(13))%32) + int32(-8192)
	goto L144
L148:
	;
	F_ginRedoRecompress(m, v536, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L19
	} else {
		goto L159
	}
L149:
	;
	v567 = v564
	goto L148
L150:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541+int32(0))+76)))
	if v547 != int32(1) {
		v564 = v537
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v551 = v541 + int32(76)
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+43)))
	if v552 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v539 == int32(0) {
		v564 = v537
		goto L149
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v539 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v557 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v539))) = v557
	v567 = v557
	goto L148
L156:
	;
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v539))) = v560
	goto L158
L157:
	;
	goto L158
L158:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v551)+44))
	v564 = v562
	goto L149
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v536))) = base.I64_rotl(v510, int64(32))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_MarkBufferDirty(m, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	goto L143
L161:
	;
	F_UnlockReleaseBuffer(m, v577)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	goto L8
L163:
	;
	if v587 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if v591 < int32(0) {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	goto L166
L166:
	;
	v624 = F_XLogReadBufferForRedo(m, l0, int32(0), v11+int32(-20))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L19
	} else {
		goto L172
	}
L167:
	;
	v610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v609)+16)))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v610+v609))) = v612
	*(*int64)(unsafe.Add(mBase, uint32(v609))) = base.I64_rotl(v583, int64(32))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	F_MarkBufferDirty(m, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L19
	} else {
		goto L171
	}
L168:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v595+(v591^int32(-1))<<(uint(int32(2))%32))))
	v609 = v601
	goto L167
L169:
	;
	goto L170
L170:
	;
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v609 = v603 + v591<<(uint(int32(13))%32) + int32(-8192)
	goto L167
L171:
	;
	goto L166
L172:
	;
	if v624 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v628 < int32(0) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	goto L175
L175:
	;
	v666 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(-4))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L19
	} else {
		goto L181
	}
L176:
	;
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646)+16)))
	v648 = v647 + v646
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v648)+6)))
	v651 = v649 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v648)+6)) = uint16(v651)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v582)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v646))) = base.I64_rotl(v583, int64(32))
	*(*int32)(unsafe.Add(mBase, uint32(v646)+20)) = v653
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_MarkBufferDirty(m, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L19
	} else {
		goto L180
	}
L177:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v632+(v628^int32(-1))<<(uint(int32(2))%32))))
	v646 = v638
	goto L176
L178:
	;
	goto L179
L179:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v646 = v640 + v628<<(uint(int32(13))%32) + int32(-8192)
	goto L176
L180:
	;
	goto L175
L181:
	;
	if v666 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v670 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	goto L184
L184:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if v728 != 0 {
		goto L197
	} else {
		goto L198
	}
L185:
	;
	v689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v582))))
	v692 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v688)+16)))
	v694 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v688+v692)+4)))
	if v694 != v689 {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v674+(v670^int32(-1))<<(uint(int32(2))%32))))
	v688 = v680
	goto L185
L187:
	;
	goto L188
L188:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v688 = v682 + v670<<(uint(int32(13))%32) + int32(-8192)
	goto L185
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v688))) = base.I64_rotl(v583, int64(32))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	F_MarkBufferDirty(m, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L19
	} else {
		goto L196
	}
L190:
	;
	v698 = (v694 - v689) * int32(10)
	if v698 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v711 = v692
	goto L192
L192:
	;
	v714 = v694 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v711+v688)+4)) = uint16(v714)
	v719 = v694*int32(10) + int32(22)
	*(*uint16)(unsafe.Add(mBase, uint32(v688)+12)) = uint16(v719)
	goto L189
L193:
	;
	v701 = v688 + v689*int32(10)
	base.MemoryCopy(m, v701+int32(22), v701+int32(32), v698)
	goto L195
L194:
	;
	goto L195
L195:
	;
	v708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v688)+16)))
	v711 = v708
	goto L192
L196:
	;
	goto L184
L197:
	;
	F_UnlockReleaseBuffer(m, v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L19
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v731 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	goto L199
L201:
	;
	F_UnlockReleaseBuffer(m, v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L19
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v734 == int32(0) {
		goto L8
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	F_UnlockReleaseBuffer(m, v734)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L19
	} else {
		goto L206
	}
L206:
	;
	goto L8
L207:
	;
	if v742 < int32(0) {
		goto L214
	} else {
		goto L215
	}
L208:
	;
	if v742 < int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v747+(v742^int32(-1))<<(uint(int32(2))%32))))
	v761 = v753
	goto L207
L210:
	;
	goto L211
L211:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v761 = v755 + v742<<(uint(int32(13))%32) + int32(-8192)
	goto L207
L212:
	;
	v806 = *(*int64)(unsafe.Add(mBase, uint32(v739)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+72)) = v806
	v808 = *(*int64)(unsafe.Add(mBase, uint32(v739)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+64)) = v808
	v810 = *(*int64)(unsafe.Add(mBase, uint32(v739)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+56)) = v810
	v812 = *(*int64)(unsafe.Add(mBase, uint32(v739)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+48)) = v812
	v814 = *(*int64)(unsafe.Add(mBase, uint32(v739)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+40)) = v814
	v816 = *(*int64)(unsafe.Add(mBase, uint32(v739)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+32)) = v816
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v739)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+24)) = v818
	v821 = base.I64_rotl(v740, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v761))) = v821
	F_MarkBufferDirty(m, v742)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L19
	} else {
		goto L217
	}
L213:
	;
	v782 = int32(8)
	F_PageInit(m, v780, int32(_a_F_gin_redo_1), v782)
	mBase = m.M
	v784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v780)+16)))
	v785 = v780 + v784
	*(*int32)(unsafe.Add(mBase, uint32(v785))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v785)+6)) = uint16(v782)
	v790 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v780)+64)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v780)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v780)+32)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v780)+40)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v780)+48)) = v790
	*(*int32)(unsafe.Add(mBase, uint32(v780)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v780)+72)) = int32(2)
	v804 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v780)+12)) = uint16(v804)
	goto L212
L214:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v766+(v742^int32(-1))<<(uint(int32(2))%32))))
	v780 = v772
	goto L213
L215:
	;
	goto L216
L216:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v780 = v774 + v742<<(uint(int32(13))%32) + int32(-8192)
	goto L213
L217:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v739)+80))
	if int32(0) < v825 {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	F_UnlockReleaseBuffer(m, v742)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L19
	} else {
		goto L262
	}
L219:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v989 == int32(0) {
		goto L218
	} else {
		goto L260
	}
L220:
	;
	v831 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(-20))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L19
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v739)+72))
	if v944 == int32(-1) {
		goto L218
	} else {
		goto L252
	}
L223:
	;
	if v831 != 0 {
		goto L219
	} else {
		goto L224
	}
L224:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v833 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v854 = v11 + int32(-4)
	v855 = int32(0)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v856)+72))
	if v857 < int32(1) {
		v879 = v855
		goto L230
	} else {
		goto L231
	}
L226:
	;
	v837 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v837+(v833^int32(-1))<<(uint(int32(2))%32))))
	v851 = v843
	goto L225
L227:
	;
	goto L228
L228:
	;
	v845 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v851 = v845 + v833<<(uint(int32(13))%32) + int32(-8192)
	goto L225
L229:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v739)+80))
	if int32(0) < v883 {
		goto L240
	} else {
		goto L241
	}
L230:
	;
	v882 = v879
	goto L229
L231:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856+int32(52))+76)))
	if v862 != int32(1) {
		v879 = v855
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v866 = v856 + int32(128)
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+43)))
	if v867 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v854 == int32(0) {
		v879 = v855
		goto L230
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	if v854 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v872 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = v872
	v882 = v872
	goto L229
L237:
	;
	v875 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v866)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = v875
	goto L239
L238:
	;
	goto L239
L239:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v866)+44))
	v879 = v877
	goto L230
L240:
	;
	v886 = int32(1)
	v887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v851)+12)))
	if base.Ui32(v887) < base.Ui32(int32(25)) {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	goto L242
L242:
	;
	v934 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v851)+16)))
	v935 = v851 + v934
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+4)))
	v938 = v936 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v935)+4)) = uint16(v938)
	*(*int64)(unsafe.Add(mBase, uint32(v851))) = v821
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_MarkBufferDirty(m, v941)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L19
	} else {
		goto L251
	}
L243:
	;
	v896 = v886
	goto L245
L244:
	;
	v896 = int32(base.Ui32(v887+int32(_a_F_gin_redo_8))>>(uint(int32(2))%32)) + v886
	goto L245
L245:
	;
	v897 = v896
	v898 = v882
	v901 = int32(0)
	goto L246
L246:
	;
	v907 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v898)+6)))
	v909 = v907 & int32(_a_F_gin_redo_4)
	v913 = F_PageAddItemExtended(m, v851, v898, v909, v897&int32(_a_F_gin_redo_9), int32(0))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L19
	} else {
		goto L248
	}
L247:
	;
	goto L242
L248:
	;
	if v913 == int32(0) {
		goto L3
	} else {
		goto L249
	}
L249:
	;
	v917 = int32(1)
	v921 = v901 + v917
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v739)+80))
	if v921 < v922 {
		v897 = v897 + v917
		v898 = v898 + v909
		v901 = v921
		goto L246
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	goto L219
L252:
	;
	v950 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(-20))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L19
	} else {
		goto L253
	}
L253:
	;
	if v950 != 0 {
		goto L219
	} else {
		goto L254
	}
L254:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v952 < int32(0) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v971 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v970)+16)))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v739)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v971+v970))) = v973
	*(*int64)(unsafe.Add(mBase, uint32(v970))) = v821
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_MarkBufferDirty(m, v976)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L19
	} else {
		goto L259
	}
L256:
	;
	v956 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v956+(v952^int32(-1))<<(uint(int32(2))%32))))
	v970 = v962
	goto L255
L257:
	;
	goto L258
L258:
	;
	v964 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v970 = v964 + v952<<(uint(int32(13))%32) + int32(-8192)
	goto L255
L259:
	;
	goto L219
L260:
	;
	F_UnlockReleaseBuffer(m, v989)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L19
	} else {
		goto L261
	}
L261:
	;
	goto L218
L262:
	;
	goto L8
L263:
	;
	v1029 = int32(16)
	if v1009 < int32(0) {
		goto L270
	} else {
		goto L271
	}
L264:
	;
	if v1009 < int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1014+(v1009^int32(-1))<<(uint(int32(2))%32))))
	v1028 = v1020
	goto L263
L266:
	;
	goto L267
L267:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1028 = v1022 + v1009<<(uint(int32(13))%32) + int32(-8192)
	goto L263
L268:
	;
	v1056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1028)+16)))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1006)))
	*(*int32)(unsafe.Add(mBase, uint32(v1028+v1056))) = v1058
	v1060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1028)+16)))
	if v1058 != int32(-1) {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	F_PageInit(m, v1047, int32(_a_F_gin_redo_1), int32(8))
	mBase = m.M
	v1051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1047)+16)))
	v1052 = v1047 + v1051
	*(*int32)(unsafe.Add(mBase, uint32(v1052))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1052)+6)) = uint16(v1029)
	goto L268
L270:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1033+(v1009^int32(-1))<<(uint(int32(2))%32))))
	v1047 = v1039
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1047 = v1041 + v1009<<(uint(int32(13))%32) + int32(-8192)
	goto L269
L273:
	;
	v1072 = v1060
	v1073 = int32(0)
	goto L275
L274:
	;
	v1065 = v1060 + v1028
	v1066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1065)+6)))
	v1068 = v1066 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1065)+6)) = uint16(v1068)
	v1070 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1028)+16)))
	v1072 = v1070
	v1073 = int32(1)
	goto L275
L275:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1072+v1028)+4)) = uint16(v1073)
	v1076 = int32(0)
	v1078 = v11 + int32(-20)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+72))
	if v1081 < v1076 {
		v1103 = v1076
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	if int32(0) < v1108 {
		goto L287
	} else {
		goto L288
	}
L277:
	;
	v1106 = v1103
	goto L276
L278:
	;
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1080+int32(0))+76)))
	if v1086 != int32(1) {
		v1103 = v1076
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v1090 = v1080 + int32(76)
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1090)+43)))
	if v1091 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	if v1078 == int32(0) {
		v1103 = v1076
		goto L277
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	if v1078 != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1096 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1078))) = v1096
	v1106 = v1096
	goto L276
L284:
	;
	v1099 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1090)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1078))) = v1099
	goto L286
L285:
	;
	goto L286
L286:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+44))
	v1103 = v1101
	goto L277
L287:
	;
	v1111 = int32(1)
	v1112 = v1106
	v1113 = int32(0)
	goto L290
L288:
	;
	goto L289
L289:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1028))) = base.I64_rotl(v1007, int64(32))
	F_MarkBufferDirty(m, v1009)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L19
	} else {
		goto L295
	}
L290:
	;
	v1121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1112)+6)))
	v1123 = v1121 & int32(_a_F_gin_redo_4)
	v1127 = F_PageAddItemExtended(m, v1028, v1112, v1123, v1111&int32(_a_F_gin_redo_9), int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L19
	} else {
		goto L292
	}
L291:
	;
	goto L289
L292:
	;
	if v1127 == int32(0) {
		goto L2
	} else {
		goto L293
	}
L293:
	;
	v1131 = int32(1)
	v1135 = v1113 + v1131
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	if v1135 < v1136 {
		v1111 = v1111 + v1131
		v1112 = v1112 + v1123
		v1113 = v1135
		goto L290
	} else {
		goto L294
	}
L294:
	;
	goto L291
L295:
	;
	F_UnlockReleaseBuffer(m, v1009)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L19
	} else {
		goto L296
	}
L296:
	;
	goto L8
L297:
	;
	if v1158 < int32(0) {
		goto L304
	} else {
		goto L305
	}
L298:
	;
	if v1158 < int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1163+(v1158^int32(-1))<<(uint(int32(2))%32))))
	v1177 = v1169
	goto L297
L300:
	;
	goto L301
L301:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1177 = v1171 + v1158<<(uint(int32(13))%32) + int32(-8192)
	goto L297
L302:
	;
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(v1155)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+72)) = v1222
	v1224 = *(*int64)(unsafe.Add(mBase, uint32(v1155)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+64)) = v1224
	v1226 = *(*int64)(unsafe.Add(mBase, uint32(v1155)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+56)) = v1226
	v1228 = *(*int64)(unsafe.Add(mBase, uint32(v1155)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+48)) = v1228
	v1230 = *(*int64)(unsafe.Add(mBase, uint32(v1155)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+40)) = v1230
	v1232 = *(*int64)(unsafe.Add(mBase, uint32(v1155)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+32)) = v1232
	v1234 = *(*int64)(unsafe.Add(mBase, uint32(v1155)))
	*(*int64)(unsafe.Add(mBase, uint32(v1177)+24)) = v1234
	v1237 = base.I64_rotl(v1156, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v1177))) = v1237
	F_MarkBufferDirty(m, v1158)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L19
	} else {
		goto L307
	}
L303:
	;
	v1198 = int32(8)
	F_PageInit(m, v1196, int32(_a_F_gin_redo_1), v1198)
	mBase = m.M
	v1200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1196)+16)))
	v1201 = v1196 + v1200
	*(*int32)(unsafe.Add(mBase, uint32(v1201))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1201)+6)) = uint16(v1198)
	v1206 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1196)+64)) = v1206
	*(*int64)(unsafe.Add(mBase, uint32(v1196)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1196)+32)) = v1206
	*(*int64)(unsafe.Add(mBase, uint32(v1196)+40)) = v1206
	*(*int64)(unsafe.Add(mBase, uint32(v1196)+48)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+72)) = int32(2)
	v1220 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v1196)+12)) = uint16(v1220)
	goto L302
L304:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1182+(v1158^int32(-1))<<(uint(int32(2))%32))))
	v1196 = v1188
	goto L303
L305:
	;
	goto L306
L306:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1196 = v1190 + v1158<<(uint(int32(13))%32) + int32(-8192)
	goto L303
L307:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+56))
	if int32(0) < v1241 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1247 = int32(0)
	goto L311
L309:
	;
	goto L310
L310:
	;
	F_UnlockReleaseBuffer(m, v1158)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L19
	} else {
		goto L326
	}
L311:
	;
	v1256 = v1247 + int32(1)
	v1259 = F_XLogInitBufferForRedo(m, l0, v1256&int32(255))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L19
	} else {
		goto L314
	}
L312:
	;
	goto L310
L313:
	;
	v1279 = int32(4)
	if v1259 < int32(0) {
		goto L320
	} else {
		goto L321
	}
L314:
	;
	if v1259 < int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1264+(v1259^int32(-1))<<(uint(int32(2))%32))))
	v1278 = v1270
	goto L313
L316:
	;
	goto L317
L317:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1278 = v1272 + v1259<<(uint(int32(13))%32) + int32(-8192)
	goto L313
L318:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1278))) = v1237
	F_MarkBufferDirty(m, v1259)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L19
	} else {
		goto L323
	}
L319:
	;
	F_PageInit(m, v1297, int32(_a_F_gin_redo_1), int32(8))
	mBase = m.M
	v1301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1297)+16)))
	v1302 = v1297 + v1301
	*(*int32)(unsafe.Add(mBase, uint32(v1302))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1302)+6)) = uint16(v1279)
	goto L318
L320:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[2]))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1283+(v1259^int32(-1))<<(uint(int32(2))%32))))
	v1297 = v1289
	goto L319
L321:
	;
	goto L322
L322:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, _c_F_gin_redo[3]))
	v1297 = v1291 + v1259<<(uint(int32(13))%32) + int32(-8192)
	goto L319
L323:
	;
	F_UnlockReleaseBuffer(m, v1259)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L19
	} else {
		goto L324
	}
L324:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+56))
	if v1256 < v1311 {
		v1247 = v1256
		goto L311
	} else {
		goto L325
	}
L325:
	;
	goto L312
L326:
	;
	goto L8
L327:
	;
	m.G0 = v13 - int32(-64)
	return
L328:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_10), int32(0))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L19
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(418), int32(_a_F_gin_redo_11))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L19
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_12), int32(0))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L19
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(421), int32(_a_F_gin_redo_11))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L19
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_13), int32(0))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L19
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(426), int32(_a_F_gin_redo_11))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L19
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_14), int32(0))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L19
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(445), int32(_a_F_gin_redo_15))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L19
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_16), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L19
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(577), int32(_a_F_gin_redo_17))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L19
	} else {
		goto L342
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	F_errmsg_internal(m, int32(_a_F_gin_redo_16), int32(0))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L19
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(659), int32(_a_F_gin_redo_18))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L19
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
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v24
	F_errmsg_internal(m, int32(_a_F_gin_redo_19), v13)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L19
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_gin_redo_6), int32(766), int32(_a_F_gin_redo_20))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L19
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_trgm_triconsistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v42 float64
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
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
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v349 int32
	_ = v349
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
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
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int64
	_ = v469
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v497 float32
	_ = v497
	var v504 int64
	_ = v504
	var v514 int64
	_ = v514
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = base.I32_wrap_i64(v17)
	v20 = v18 & int32(_a_F_gin_trgm_triconsistent_0)
	if base.Ui32(int32(11)) < base.Ui32(v20) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v514
L2:
	;
	if base.F64_le(v42, base.F64_promote_f32(base.F32_div(v497, base.F32_convert_i32_s(v23)))) != 0 {
		goto L91
	} else {
		goto L92
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L13
	} else {
		goto L88
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = int32(1) << (uint(v20) % 32)
	if v26&int32(642) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v164 = v23 & int32(3)
	v166 = F_palloc_mul(m, int32(1), v23)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L37
	}
L6:
	;
	v140 = int32(0)
	v141 = int64(2)
	if v23 <= v140 {
		v514 = v141
		goto L1
	} else {
		goto L30
	}
L7:
	;
	if v26&int32(2072) != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v42 = F_index_strategy_get_limit(m, v18&int32(_a_F_gin_trgm_triconsistent_0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v26&int32(96) == int32(0) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if int32(0) < v23 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v514 = int64(2)
	goto L1
L13:
	;
	return int64(0)
L14:
	;
	if int32(0) < v23 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = v23 & int32(3)
	v50 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v23) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if v23 != 0 {
		v497 = float32(0)
		goto L2
	} else {
		goto L29
	}
L18:
	;
	v497 = base.F32_convert_i32_u(v128)
	goto L2
L19:
	;
	v56 = v50
	v57 = v50
	v63 = int32(0)
	goto L22
L20:
	;
	v92 = v50
	v93 = v50
	goto L21
L21:
	;
	v105 = v92
	v106 = v93
	v107 = int32(0)
	goto L26
L22:
	;
	v68 = v56 + v24
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v70 = int32(0)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+2)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+3)))
	v84 = v57 + base.B2i32(v69 != v70) + base.B2i32(v73 != v70) + base.B2i32(v77 != v70) + base.B2i32(v81 != v70)
	v85 = int32(4)
	v86 = v56 + v85
	v88 = v63 + v85
	if v88 != v23&int32(2147483644) {
		v56 = v86
		v57 = v84
		v63 = v88
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v49 == int32(0) {
		v128 = v84
		goto L18
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v92 = v86
	v93 = v84
	goto L21
L26:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v24))))
	v121 = v106 + base.B2i32(v118 != int32(0))
	v122 = int32(1)
	v125 = v107 + v122
	if v125 != v49 {
		v105 = v105 + v122
		v106 = v121
		v107 = v125
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v128 = v121
	goto L18
L28:
	;
	goto L27
L29:
	;
	v514 = int64(0)
	goto L1
L30:
	;
	v144 = v140
	goto L31
L31:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v24))))
	if v157 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v514 = int64(0)
	goto L1
L33:
	;
	v159 = v144 + int32(1)
	if v23 != v159 {
		v144 = v159
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v514 = v141
	goto L1
L37:
	;
	v168 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v23) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v272 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v282 != 0 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v174 = v168
	v176 = int32(0)
	goto L42
L40:
	;
	v223 = v168
	goto L41
L41:
	;
	v236 = v223
	v238 = int32(0)
	goto L46
L42:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v24))))
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v174+v166))) = uint8(base.B2i32(v188 != v189))
	v193 = v174 | int32(1)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v193))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166+v193))) = uint8(base.B2i32(v196 != v189))
	v201 = v174 | int32(2)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v201))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166+v201))) = uint8(base.B2i32(v204 != v189))
	v209 = v174 | int32(3)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v209))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166+v209))) = uint8(base.B2i32(v212 != v189))
	v216 = int32(4)
	v217 = v174 + v216
	v219 = v176 + v216
	if v219 != v23&int32(2147483644) {
		v174 = v217
		v176 = v219
		goto L42
	} else {
		goto L44
	}
L43:
	;
	if v164 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v223 = v217
	goto L41
L46:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v24))))
	*(*uint8)(unsafe.Add(mBase, uint32(v236+v166))) = uint8(base.B2i32(v250 != int32(0)))
	v254 = int32(1)
	v257 = v238 + v254
	if v257 != v164 {
		v236 = v236 + v254
		v238 = v257
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L38
L48:
	;
	goto L47
L49:
	;
	F_pfree(m, v166)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L13
	} else {
		goto L84
	}
L50:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v271)+16))
	base.MemoryFill(m, v283, int32(0), v282)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	if v286 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	base.MemoryFill(m, v287, int32(0), v286)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if int32(0) < v290 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	v297 = v272
	v298 = v272
	goto L59
L57:
	;
	goto L58
L58:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	v364 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v363))) = uint8(v364)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v271)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v366))) = int32(0)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v376 = v364
	v381 = v272
	goto L71
L59:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v293+v297<<(uint(int32(2))%32))))
	v310 = v309 + v298
	if v309 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L58
L61:
	;
	v349 = v297 + int32(1)
	if v349 != v290 {
		v297 = v349
		v298 = v310
		goto L59
	} else {
		goto L69
	}
L62:
	;
	v315 = v298
	goto L63
L63:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v315))))
	if v326 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v271)+16))
	v334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v332+v297))) = uint8(v334)
	goto L61
L65:
	;
	v330 = v315 + int32(1)
	if v330 < v310 {
		v315 = v330
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L61
L69:
	;
	goto L60
L70:
	;
	goto L49
L71:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v366+v381<<(uint(int32(2))%32))))
	v389 = v369 + v386<<(uint(int32(3))%32)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	if int32(0) < v390 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v464 = int32(0)
	goto L70
L73:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v271)+16))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v398 = int32(0)
	v401 = v376
	goto L76
L74:
	;
	v441 = v376
	goto L75
L75:
	;
	v449 = v381 + int32(1)
	if v449 < v441 {
		v376 = v441
		v381 = v449
		goto L71
	} else {
		goto L83
	}
L76:
	;
	v410 = v394 + v398<<(uint(int32(3))%32)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v411))))
	if v413 != int32(1) {
		v431 = v401
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v441 = v431
	goto L75
L78:
	;
	v434 = v398 + int32(1)
	if v434 != v390 {
		v398 = v434
		v401 = v431
		goto L76
	} else {
		goto L82
	}
L79:
	;
	v416 = int32(1)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	if v417 == v416 {
		v464 = v416
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v420 = v363 + v417
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	if v421 != 0 {
		v431 = v401
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v422 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v422)
	*(*int32)(unsafe.Add(mBase, uint32(v366+v401<<(uint(int32(2))%32)))) = v417
	v431 = v401 + v422
	goto L78
L82:
	;
	goto L77
L83:
	;
	goto L72
L84:
	;
	if v464 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v469 = int64(2)
	goto L87
L86:
	;
	v469 = int64(0)
	goto L87
L87:
	;
	v514 = v469
	goto L1
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v20
	F_errmsg_internal(m, int32(_a_F_gin_trgm_triconsistent_1), v15)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_gin_trgm_triconsistent_2), int32(354), int32(_a_F_gin_trgm_triconsistent_3))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v504 = int64(2)
	goto L93
L92:
	;
	v504 = int64(0)
	goto L93
L93:
	;
	v514 = v504
	goto L1
}
func F_gin_tsquery_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v2)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v16 <= v2 {
		v40 = int64(0)
		m.G0 = v8 + int32(16)
		return v40
	} else {
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v11)
		v22 = v12 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v22
		v25 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v10))))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v25
		v30 = F_TS_execute_ternary(m, v22, v8+int32(4))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int64(0)
		} else {
			switch v30 - int32(1) {
			case 0:
				v40 = int64(1)
			case 1:
				v36 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v36)
				v40 = int64(1)
			default:
				v40 = int64(0)
			}
			m.G0 = v8 + int32(16)
			return v40
		}
	}
}
func F_gin_tsquery_consistent_6args(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(7) < v10 {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v17)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		if v19 <= v17 {
			v43 = int64(0)
			m.G0 = v8 + int32(16)
			return v43
		} else {
			*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v14)
			v25 = v15 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v13))))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v28
			v33 = F_TS_execute_ternary(m, v25, v8+int32(4))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				switch v33 - int32(1) {
				case 0:
					v43 = int64(1)
				case 1:
					v39 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v39)
					v43 = int64(1)
				default:
					v43 = int64(0)
				}
				m.G0 = v8 + int32(16)
				return v43
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_gin_tsquery_consistent_6args_0), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_tsquery_consistent_6args_1), int32(337), int32(_a_F_gin_tsquery_consistent_6args_2))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_gin_xlog_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_startup[0]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(_a_F_gin_xlog_startup_0), int32(0), int32(_a_F_gin_xlog_startup_1), int32(_a_F_gin_xlog_startup_2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_gin_xlog_startup[1])) = v8
		return
	}
}
