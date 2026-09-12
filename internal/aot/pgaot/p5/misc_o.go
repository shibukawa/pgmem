package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+13)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+28)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = l1
	v16 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+22)) = uint16(v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = m.T0[v20].(func(*base.Module, int32) int32)(m, v6+int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
		if v25 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v32
				F_errmsg_internal(m, int32(558657), v6)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519609), int32(1143), int32(318845))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v6 + int32(32)
			return v21
		}
	}
}
func F_offsethash_grow(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v67 float64
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v3 = int32(0)
	v13 = int64(2)
	if base.Ui64(l1) <= base.Ui64(v13) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L11
	} else {
		goto L49
	}
L2:
	;
	v16 = v13
	goto L4
L3:
	;
	v16 = l1
	goto L4
L4:
	;
	v17 = int64(1)
	if v16&(v16-v17) == int64(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = v16
	goto L7
L6:
	;
	v27 = v17 << (uint(int64(64)-base.I64_clz(v16)) % 64)
	goto L7
L7:
	;
	if base.Ui64(v27<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v39 = F_MemoryContextAllocExtended(m, v34, base.I32_wrap_i64(v27)<<(uint(int32(3))%32), int32(5))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L46
	}
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v39
	v42 = int64(1)
	if v27&(v27-v42) == int64(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v27
	goto L15
L14:
	;
	v52 = v42 << (uint(int64(64)-base.I64_clz(v27)) % 64)
	goto L15
L15:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v52<<(uint(int64(3))%64)) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v52
	v60 = base.I32_wrap_i64(v52) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
	v67 = base.F64_mul(base.F64_convert_i64_u(v52), float64(0.9))
	if base.F64_lt(v67, float64(4.294967296e+09))&base.F64_ge(v67, float64(0)) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v52 == int64(4294967296) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v73 = base.I32_trunc_f64_u(v67)
	v75 = v73
	goto L17
L19:
	;
	goto L20
L20:
	;
	v75 = int32(0)
	goto L17
L21:
	;
	v76 = int32(-85899346)
	goto L23
L22:
	;
	v76 = v75
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v76
	if v33 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v83 = v3
	goto L28
L25:
	;
	goto L26
L26:
	;
	F_pfree(m, v32)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L11
	} else {
		goto L45
	}
L27:
	;
	v124 = v120
	v127 = v3
	goto L33
L28:
	;
	v94 = v32 + v83<<(uint(int32(3))%32)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v95 != int32(1) {
		v120 = v83
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v120 = int32(0)
	goto L27
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v99 = int32(16)
	v103 = (int32(base.Ui32(v98)>>(uint(v99)%32)) ^ v98) * int32(-2048144789)
	v108 = (int32(base.Ui32(v103)>>(uint(int32(13))%32)) ^ v103) * int32(-1028477387)
	if (int32(base.Ui32(v108)>>(uint(v99)%32))^v108)&v60 == v83 {
		v120 = v83
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v115 = v83 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v115)) < base.Ui64(v33) {
		v83 = v115
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v135 = v32 + v124<<(uint(int32(3))%32)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
	if v136 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L26
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v140 = int32(16)
	v144 = (int32(base.Ui32(v139)>>(uint(v140)%32)) ^ v139) * int32(-2048144789)
	v149 = (int32(base.Ui32(v144)>>(uint(int32(13))%32)) ^ v144) * int32(-1028477387)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = int32(base.Ui32(v149)>>(uint(v140)%32)) ^ v149
	goto L38
L36:
	;
	goto L37
L37:
	;
	v188 = v124 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v188)) < base.Ui64(v33) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v166 = v156 & v153
	v171 = v39 + v166<<(uint(int32(3))%32)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+4)))
	if v172 != 0 {
		v156 = v166 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v135)))
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = v173
	goto L37
L40:
	;
	goto L39
L41:
	;
	v192 = v188
	goto L43
L42:
	;
	v192 = int32(0)
	goto L43
L43:
	;
	v194 = v127 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v194)) < base.Ui64(v33) {
		v124 = v192
		v127 = v194
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	return
L46:
	;
	F_errmsg_internal(m, int32(419840), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(343318), int32(327), int32(358465))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L11
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
	F_errmsg_internal(m, int32(419840), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(343318), int32(327), int32(358465))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_oidgt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v3) < base.Ui32(v2))
}
func F_oidle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v2) <= base.Ui32(v3))
}
func F_oidvectorhashfast(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_DirectFunctionCall1Coll(m, int32(1591), int32(0), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_operationPriority(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = l0 - int32(4)
	if base.Ui32(v5) <= base.Ui32(int32(37)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_consts[1064])))
		v13 = v12
	} else {
		v13 = int32(6)
	}
	return v13
}
func F_overlaps_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	v11 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 == v11 {
		if v14&int32(1) == int32(0) {
			v37 = v17
			v39 = v17
			v40 = v11
			if v13&int32(1) != 0 {
				if v12&int32(1) != 0 {
					v107 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
					v117 = int32(0)
					return v117
				} else {
					v48 = F_DirectFunctionCall2Coll(m, int32(1516), int32(0), v39, v15)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v15
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				}
			} else {
				if v12&int32(1) != 0 {
					v55 = F_DirectFunctionCall2Coll(m, int32(1516), int32(0), v39, v16)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v16
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(1516)
					v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 != 0 {
							v65 = v15
						} else {
							v65 = v16
						}
						v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							if v66 != 0 {
								if v63 != 0 {
									v91 = v16
								} else {
									v91 = v15
								}
								v92 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v40&base.B2i32(v92 == int32(0)) != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										return base.B2i32(v92 != int32(0))
									}
								}
							} else {
								v68 = v65
								v69 = v58
								v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				}
			}
		} else {
			v107 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
			v117 = int32(0)
			return v117
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14&int32(1) != 0 {
			v37 = v17
			v39 = v25
			v40 = v11
			if v13&int32(1) != 0 {
				if v12&int32(1) != 0 {
					v107 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
					v117 = int32(0)
					return v117
				} else {
					v48 = F_DirectFunctionCall2Coll(m, int32(1516), int32(0), v39, v15)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v15
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				}
			} else {
				if v12&int32(1) != 0 {
					v55 = F_DirectFunctionCall2Coll(m, int32(1516), int32(0), v39, v16)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
							v117 = int32(0)
							return v117
						} else {
							v68 = v16
							v69 = int32(1)
							v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != 0 {
									if v40 != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v69&base.B2i32(v77 == int32(0)) != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												return base.B2i32(v77 != int32(0))
											}
										}
									}
								} else {
									v85 = int32(1)
									if v69|v40 != v85 {
										v117 = v85
									} else {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
									}
									return v117
								}
							}
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(1516)
					v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 != 0 {
							v65 = v15
						} else {
							v65 = v16
						}
						v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							if v66 != 0 {
								if v63 != 0 {
									v91 = v16
								} else {
									v91 = v15
								}
								v92 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v40&base.B2i32(v92 == int32(0)) != 0 {
										v107 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
										v117 = int32(0)
										return v117
									} else {
										return base.B2i32(v92 != int32(0))
									}
								}
							} else {
								v68 = v65
								v69 = v58
								v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				}
			}
		} else {
			v28 = int32(0)
			v31 = F_DirectFunctionCall2Coll(m, int32(1516), v28, v25, v17)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v31 != 0 {
					v35 = v17
				} else {
					v35 = v25
				}
				if v31 != 0 {
					v36 = v25
				} else {
					v36 = v17
				}
				v37 = v36
				v39 = v35
				v40 = v28
				if v13&int32(1) != 0 {
					if v12&int32(1) != 0 {
						v107 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
						v117 = int32(0)
						return v117
					} else {
						v48 = F_DirectFunctionCall2Coll(m, int32(1516), int32(0), v39, v15)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 != 0 {
								v107 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
								v117 = int32(0)
								return v117
							} else {
								v68 = v15
								v69 = int32(1)
								v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					}
				} else {
					if v12&int32(1) != 0 {
						v55 = F_DirectFunctionCall2Coll(m, int32(1516), int32(0), v39, v16)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v107 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
								v117 = int32(0)
								return v117
							} else {
								v68 = v16
								v69 = int32(1)
								v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										if v40 != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v69&base.B2i32(v77 == int32(0)) != 0 {
													v107 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
													v117 = int32(0)
													return v117
												} else {
													return base.B2i32(v77 != int32(0))
												}
											}
										}
									} else {
										v85 = int32(1)
										if v69|v40 != v85 {
											v117 = v85
										} else {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
										}
										return v117
									}
								}
							}
						}
					} else {
						v58 = int32(0)
						v59 = int32(1516)
						v63 = F_DirectFunctionCall2Coll(m, v59, v58, v16, v15)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v63 != 0 {
								v65 = v15
							} else {
								v65 = v16
							}
							v66 = F_DirectFunctionCall2Coll(m, v59, v58, v39, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								if v66 != 0 {
									if v63 != 0 {
										v91 = v16
									} else {
										v91 = v15
									}
									v92 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v91)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										if v40&base.B2i32(v92 == int32(0)) != 0 {
											v107 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
											v117 = int32(0)
											return v117
										} else {
											return base.B2i32(v92 != int32(0))
										}
									}
								} else {
									v68 = v65
									v69 = v58
									v73 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v39, v68)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										if v73 != 0 {
											if v40 != 0 {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
												return v117
											} else {
												v77 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v68, v37)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													if v69&base.B2i32(v77 == int32(0)) != 0 {
														v107 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
														v117 = int32(0)
														return v117
													} else {
														return base.B2i32(v77 != int32(0))
													}
												}
											}
										} else {
											v85 = int32(1)
											if v69|v40 != v85 {
												v117 = v85
											} else {
												v107 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v107)
												v117 = int32(0)
											}
											return v117
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
