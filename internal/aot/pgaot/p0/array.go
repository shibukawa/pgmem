package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_array_agg_array_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
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
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
		F_enlargeStringInfo(m, v7, int32(4))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v23 = int32(16711935)
			*(*int32)(unsafe.Add(mBase, uint32(v18+v19))) = base.I32_rotr(v14, int32(24))&v23 | base.I32_rotr(v14&v23, int32(8))
			v31 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + v31
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
			F_enlargeStringInfo(m, v7, v31)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v43 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v38+v39))) = base.I32_rotr(v34, int32(24))&v43 | base.I32_rotr(v34&v43, int32(8))
				v51 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v38 + v51
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				F_enlargeStringInfo(m, v7, v51)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v63 = int32(16711935)
					*(*int32)(unsafe.Add(mBase, uint32(v58+v59))) = base.I32_rotr(v54, int32(24))&v63 | base.I32_rotr(v54&v63, int32(8))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v58 + int32(4)
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					F_appendBinaryStringInfo(m, v7, v74, v75)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						F_enlargeStringInfo(m, v7, int32(4))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							v87 = int32(16711935)
							*(*int32)(unsafe.Add(mBase, uint32(v82+v83))) = base.I32_rotr(v78, int32(24))&v87 | base.I32_rotr(v78&v87, int32(8))
							v95 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v82 + v95
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
							F_enlargeStringInfo(m, v7, v95)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								v107 = int32(16711935)
								*(*int32)(unsafe.Add(mBase, uint32(v102+v103))) = base.I32_rotr(v98, int32(24))&v107 | base.I32_rotr(v98&v107, int32(8))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v102 + int32(4)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								if v118 != 0 {
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									v123 = base.I32_div_s(v119+int32(7), int32(8))
									F_appendBinaryStringInfo(m, v7, v118, v123)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										F_enlargeStringInfo(m, v7, int32(4))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
											v135 = int32(16711935)
											*(*int32)(unsafe.Add(mBase, uint32(v130+v131))) = base.I32_rotr(v126, int32(24))&v135 | base.I32_rotr(v126&v135, int32(8))
											v143 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v130 + v143
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
											F_enlargeStringInfo(m, v7, v143)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
												v153 = int32(24)
												v155 = int32(16711935)
												*(*int32)(unsafe.Add(mBase, uint32(v150+v151))) = base.I32_rotr(v146, v153)&v155 | base.I32_rotr(v146&v155, int32(8))
												*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v150 + int32(4)
												F_appendBinaryStringInfo(m, v7, v9+int32(32), v153)
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return int32(0)
												} else {
													F_appendBinaryStringInfo(m, v7, v9+int32(56), int32(24))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														v177 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v177))) = v178 << (uint(int32(2)) % 32)
														m.G0 = v7 + int32(16)
														return v177
													}
												}
											}
										}
									}
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
									F_enlargeStringInfo(m, v7, int32(4))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
										v135 = int32(16711935)
										*(*int32)(unsafe.Add(mBase, uint32(v130+v131))) = base.I32_rotr(v126, int32(24))&v135 | base.I32_rotr(v126&v135, int32(8))
										v143 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v130 + v143
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										F_enlargeStringInfo(m, v7, v143)
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
											v151 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
											v153 = int32(24)
											v155 = int32(16711935)
											*(*int32)(unsafe.Add(mBase, uint32(v150+v151))) = base.I32_rotr(v146, v153)&v155 | base.I32_rotr(v146&v155, int32(8))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v150 + int32(4)
											F_appendBinaryStringInfo(m, v7, v9+int32(32), v153)
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return int32(0)
											} else {
												F_appendBinaryStringInfo(m, v7, v9+int32(56), int32(24))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													v177 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
													v178 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v177))) = v178 << (uint(int32(2)) % 32)
													m.G0 = v7 + int32(16)
													return v177
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
func F_array_append(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
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
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_fetch_array_arg_replace_nulls(m, l0, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v22 = v21
		} else {
			v22 = v2
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		switch v23 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(1)
			v71 = int32(12)
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
			v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78)+4)))
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+6)))
			v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78)+7)))
			v82 = F_array_set_element(m, v14+v71, int32(1), v11+v71, v22, v18, int32(-1), v79, v80, v81)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v82
			}
		case 1:
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v28 = v25 + v27
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v28
			if base.B2i32(v27 < int32(0)) == base.B2i32(v28 < v25) {
				v71 = int32(12)
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
				v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78)+4)))
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+6)))
				v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78)+7)))
				v82 = F_array_set_element(m, v14+v71, int32(1), v11+v71, v22, v18, int32(-1), v79, v80, v81)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(16)
					return v82
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_array_append_0), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_array_append_1), int32(167), int32(_a_F_array_append_2))
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
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_array_append_3), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_append_1), int32(174), int32(_a_F_array_append_2))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
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
func F_array_cardinality(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_DatumGetAnyArrayP(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		if v11 == int32(-1) {
			v14 = int32(28)
		} else {
			v14 = int32(4)
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v5+v14)))
		if v11 == int32(-1) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
			v20 = F_ArrayGetNItemsSafe(m, v16, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v20
			}
		} else {
			v25 = F_ArrayGetNItemsSafe(m, v16, v5+int32(16))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v25
			}
		}
	}
}
func F_array_contain_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
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
	var v230 int32
	_ = v230
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v315 int32
	_ = v315
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(80)
	return v315
L2:
	;
	v315 = int32(0)
	goto L1
L3:
	;
	v23 = int32(40)
	goto L5
L4:
	;
	v23 = int32(12)
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+v23)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v28 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = int32(40)
	goto L8
L7:
	;
	v31 = int32(12)
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+v31)))
	if v25 == v33 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L18
	} else {
		goto L80
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L18
	} else {
		goto L75
	}
L13:
	;
	v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+11)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+10)))
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+8)))
	if v49 == int32(-1) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 == v25 {
		v48 = v35
		v49 = v28
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v39 = F_lookup_type_cache(m, v25, int32(32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	if v43 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v39
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v48 = v39
	v49 = v47
	goto L13
L21:
	;
	v73 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+62)) = uint16(v73)
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+60)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v48 + int32(76)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v85 == int32(-1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_deconstruct_expanded_array(m, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_deconstruct_array(m, l1, v52, v51&int32(1), v50, v16+int32(40), v16+int32(36), v16+int32(32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v61
	goto L21
L26:
	;
	goto L21
L27:
	;
	v88 = int32(28)
	goto L29
L28:
	;
	v88 = int32(4)
	goto L29
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0+v88)))
	if v85 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v97 = F_ArrayGetNItemsSafe(m, v90, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L18
	} else {
		goto L34
	}
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v96 = v93
	goto L30
L32:
	;
	goto L33
L33:
	;
	v96 = l0 + int32(16)
	goto L30
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v99 == int32(-1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(1)
	if v97 <= int32(0) {
		v315 = l3
		goto L1
	} else {
		goto L48
	}
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v102 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v135 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v104
	v158 = v105
	goto L35
L40:
	;
	goto L41
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	if v112 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v111 + (v115<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v158 = int32(0)
	goto L35
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v111 + v112
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v158 = v111 + v127<<(uint(int32(3))%32) + int32(16)
	goto L35
L45:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0 + (v138<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v158 = int32(0)
	goto L35
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0 + v135
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v158 = l0 + v150<<(uint(int32(3))%32) + int32(16)
	goto L35
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v172 = int32(0)
	goto L49
L49:
	;
	v187 = F_array_iter_next(m, v16+int32(12), v16+int32(11), v172, v52, v51&int32(1), v50)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L18
	} else {
		goto L51
	}
L50:
	;
	v315 = l3
	goto L1
L51:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)))
	if v189 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v262 = v172 + int32(1)
	if v262 != v97 {
		v172 = v262
		goto L49
	} else {
		goto L74
	}
L53:
	;
	if l3 != 0 {
		goto L52
	} else {
		goto L73
	}
L54:
	;
	v192 = int32(0)
	if v192 < v166 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if l3 != 0 {
		goto L2
	} else {
		goto L72
	}
L57:
	;
	v195 = v192
	goto L60
L58:
	;
	goto L59
L59:
	;
	if l3 == int32(0) {
		goto L52
	} else {
		goto L71
	}
L60:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v165+v195<<(uint(int32(2))%32))))
	if v164 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L59
L62:
	;
	v230 = v195 + int32(1)
	if v230 != v166 {
		v195 = v230
		goto L60
	} else {
		goto L70
	}
L63:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v164))))
	if v213 != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v214 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v214)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v211
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+68)) = uint8(v214)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v187
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+60)) = uint8(v214)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = m.T0[v225].(func(*base.Module, int32) int32)(m, v16+int32(44))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L18
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+60)))
	if v228 != 0 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v226 != 0 {
		goto L53
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	goto L61
L71:
	;
	goto L2
L72:
	;
	goto L52
L73:
	;
	v315 = int32(1)
	goto L1
L74:
	;
	goto L50
L75:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	v271 = F_format_type_be(m, v25)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v271
	F_errmsg(m, int32(_a_F_array_contain_compare_0), v16)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_array_contain_compare_1), int32(_a_F_array_contain_compare_2), int32(_a_F_array_contain_compare_3))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(_a_F_array_contain_compare_4), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L18
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_array_contain_compare_1), int32(_a_F_array_contain_compare_5), int32(_a_F_array_contain_compare_3))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L18
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_contains_nulls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = l0 + int32(16)
	v14 = F_ArrayGetNItemsSafe(m, v11, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = v13 + v19<<(uint(int32(3))%32)
	goto L8
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	if v14 <= int32(7) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return v73
L10:
	;
	if v42 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	v42 = v14
	v44 = v24
	goto L10
L12:
	;
	goto L13
L13:
	;
	v29 = v14
	v30 = v24
	goto L14
L14:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v33 != int32(255) {
		v73 = int32(1)
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v42 = v41
	v44 = v37
	goto L10
L16:
	;
	v37 = v30 + int32(1)
	v41 = v29 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v29) {
		v29 = v41
		v30 = v37
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return int32(0)
L19:
	;
	goto L20
L20:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v53 = v42
	v54 = int32(1)
	goto L21
L21:
	;
	v60 = base.B2i32(v54&v51 == int32(0))
	if v54&v51 == int32(0) {
		v73 = v60
		goto L9
	} else {
		goto L23
	}
L22:
	;
	v73 = v60
	goto L9
L23:
	;
	v63 = int32(1)
	if v63 < v53 {
		v53 = v53 - v63
		v54 = v54 << (uint(v63) % 32)
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
}
func F_array_fill_with_lower_bounds(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_array_fill_with_lower_bounds_0), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_array_fill_with_lower_bounds_1), int32(_a_F_array_fill_with_lower_bounds_2), int32(_a_F_array_fill_with_lower_bounds_3))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
		if v8 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67108994))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_array_fill_with_lower_bounds_0), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_fill_with_lower_bounds_1), int32(_a_F_array_fill_with_lower_bounds_2), int32(_a_F_array_fill_with_lower_bounds_3))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v12 = F_pg_detoast_datum(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v17 = F_pg_detoast_datum(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v19 == int32(0) {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v23 = v22
					} else {
						v23 = int32(0)
					}
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = F_get_fn_expr_argtype(m, v24, int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_array_fill_with_lower_bounds_4), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_array_fill_with_lower_bounds_1), int32(_a_F_array_fill_with_lower_bounds_5), int32(_a_F_array_fill_with_lower_bounds_3))
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
						} else {
							v30 = F_array_fill_internal(m, v12, v17, v23, v19, v26, l0)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					}
				}
			}
		}
	}
}
func F_array_iterator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = F_ArrayGetNItemsSafe(m, v8, l0+int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 < int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L27
	}
L7:
	;
	if v19 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L11
L10:
	;
	goto L11
L11:
	;
	if int32(0) < v11 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return v76
L13:
	;
	if v18 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v76 = int32(0)
	goto L12
L16:
	;
	v31 = v18
	goto L18
L17:
	;
	v31 = (v15<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L18:
	;
	v33 = l0 + v31
	v37 = v11
	goto L19
L19:
	;
	v41 = F_DirectFunctionCall2Coll(m, l1, int32(0), v33, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	if v41 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if l3 == int32(0) {
		v76 = int32(1)
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v57 = int32(1)
	if v57 < v37 {
		v33 = v33 + (int32(base.Ui32(v49)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v37 = v37 - v57
		goto L19
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	return int32(1)
L26:
	;
	goto L20
L27:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_array_iterator_0), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_array_iterator_1), int32(46), int32(_a_F_array_iterator_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_array_iterator_3), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_array_iterator_1), int32(50), int32(_a_F_array_iterator_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_prepend_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v5 != int32(464) {
		v23 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v10 == int32(0) {
			v23 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(8) {
				v23 = v2
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v16 != 0 {
					v23 = v2
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
					if v18 == v19 {
						v21 = v10
					} else {
						v21 = int32(0)
					}
					v23 = v21
				}
			}
		}
	}
	return v23
}
func F_array_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	v19 = m.G0
	v21 = v19 - int32(176)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L8
	} else {
		goto L135
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L8
	} else {
		goto L131
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L8
	} else {
		goto L127
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L8
	} else {
		goto L122
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L8
	} else {
		goto L116
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L8
	} else {
		goto L112
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L8
	} else {
		goto L108
	}
L8:
	;
	return int32(0)
L9:
	;
	if int32(0) <= v27 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v27) {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L8
	} else {
		goto L104
	}
L13:
	;
	v36 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v36) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v41 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v44 = int32(_a_F_array_recv_0)
	if base.B2i32(base.B2i32(v41 == v24)|base.B2i32(base.Ui32(v44) < base.Ui32(v41)) == int32(0))&base.B2i32(base.Ui32(v24) <= base.Ui32(v44)) != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v27 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = int32(0)
	goto L21
L19:
	;
	goto L20
L20:
	;
	v109 = v21 + int32(128)
	v110 = F_ArrayGetNItemsSafe(m, v27, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L26
	}
L21:
	;
	v72 = v54 << (uint(int32(2)) % 32)
	v77 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72+(v21+int32(128))))) = v77
	v84 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(96)+v72))) = v84
	v88 = v54 + int32(1)
	if v88 != v27 {
		v54 = v88
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	F_ArrayCheckBounds(m, v27, v109, v21+int32(96))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	if v117 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v110 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	F_get_type_io_data(m, v24, int32(2), v133+int32(4), v133+int32(6), v133+int32(7), v133+int32(8), v133+int32(12), v133+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L35
	}
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v122 = F_MemoryContextAlloc(m, v120, int32(48))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v131 == v24 {
		v159 = v117
		goto L28
	} else {
		goto L34
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+16)) = v122
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v24 ^ int32(-1)
	v133 = v127
	goto L29
L34:
	;
	v133 = v117
	goto L29
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	if v149 == int32(0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	F_fmgr_info_cxt(m, v149, v133+int32(20), v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v24
	v159 = v133
	goto L28
L38:
	;
	m.G0 = v21 + int32(176)
	return v497
L39:
	;
	v164 = F_palloc0(m, int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+6)))
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v159)+4)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+7)))
	v177 = F_palloc(m, v110<<(uint(int32(2))%32))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = int64(64)
	v497 = v164
	goto L38
L43:
	;
	v179 = F_palloc(m, v110)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v181 = int32(0)
	if v110 <= v181 {
		v419 = v181
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v459 = F_palloc0(m, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L8
	} else {
		goto L94
	}
L46:
	;
	v457 = int32(0)
	v458 = v419 + (v27<<(uint(int32(3))%32)+int32(23))&int32(120)
	goto L45
L47:
	;
	v185 = v159 + int32(20)
	v190 = v181
	goto L48
L48:
	;
	v205 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L50
	}
L49:
	;
	v254 = int32(0)
	v260 = v254
	v261 = v254
	v266 = v254
	goto L62
L50:
	;
	if v205 < int32(-1) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v209-v210 < v205 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	if v205 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v248 = v190 + int32(1)
	if v248 != v110 {
		v190 = v248
		goto L48
	} else {
		goto L60
	}
L54:
	;
	v219 = F_ReceiveFunctionCall(m, v185, int32(0), v171, v23)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+168)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v210 + v225
	*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v205 + v210
	v238 = F_ReceiveFunctionCall(m, v185, v21+int32(160), v171, v23)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177+v190<<(uint(int32(2))%32)))) = v219
	v223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190+v179))) = uint8(v223)
	goto L53
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177+v190<<(uint(int32(2))%32)))) = v238
	v242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v190+v179))) = uint8(v242)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v21)+172))
	if v244 != v205 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	goto L49
L61:
	;
	v406 = base.I32_div_s(v110+int32(7), int32(8))
	v413 = (v406 + v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v457 = v413
	v458 = v389 + v413
	goto L45
L62:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+v179))))
	if v276 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v312 == int32(0) {
		v419 = v377
		goto L46
	} else {
		goto L93
	}
L64:
	;
	v282 = v260
	goto L67
L65:
	;
	v306 = v260
	v312 = v266
	goto L66
L66:
	;
	if base.B2i32(v173 == int32(-1)) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v298 = v282 + int32(1)
	if v298 == v110 {
		v389 = v261
		goto L61
	} else {
		goto L69
	}
L68:
	;
	v306 = v298
	v312 = int32(1)
	goto L66
L69:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298+v179))))
	if v301 != 0 {
		v282 = v298
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v364 = v362 + v261
	switch v174 - int32(99) {
	case 0:
		v377 = v364
		goto L87
	case 1:
		goto L89
	default:
		goto L88
	case 6:
		goto L90
	}
L72:
	;
	if int32(0) < v173 {
		v362 = v173
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v334 = v177 + v306<<(uint(int32(2))%32)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	v336 = F_pg_detoast_datum(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L8
	} else {
		goto L76
	}
L75:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v177+v306<<(uint(int32(2))%32))))
	v329 = F_strlen(m, v328)
	mBase = m.M
	v362 = v329 + int32(1)
	goto L71
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334))) = v336
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	if v339 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	if base.Ui32((v343-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v362 = int32(6)
		goto L71
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v339&int32(1) != 0 {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	v350 = int32(18)
	if v343 == v350 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v354 = v350
	goto L83
L82:
	;
	v354 = int32(2)
	goto L83
L83:
	;
	v362 = v354
	goto L71
L84:
	;
	v362 = int32(base.Ui32(v339) >> (uint(int32(1)) % 32))
	goto L71
L85:
	;
	goto L86
L86:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v362 = int32(base.Ui32(v359) >> (uint(int32(2)) % 32))
	goto L71
L87:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v377) {
		goto L1
	} else {
		goto L91
	}
L88:
	;
	v377 = (v364 + int32(1)) & int32(-2)
	goto L87
L89:
	;
	v377 = (v364 + int32(7)) & int32(-8)
	goto L87
L90:
	;
	v377 = (v364 + int32(3)) & int32(-4)
	goto L87
L91:
	;
	v381 = v306 + int32(1)
	if v381 != v110 {
		v260 = v381
		v261 = v377
		v266 = v312
		goto L62
	} else {
		goto L92
	}
L92:
	;
	goto L63
L93:
	;
	v389 = v377
	goto L61
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v459)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v459)+8)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v459)+4)) = v27
	v464 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v458 << (uint(v464) % 32)
	v468 = v459 + int32(16)
	v470 = v27 << (uint(v464) % 32)
	v471 = int32(0)
	v472 = base.B2i32(v470 == v471)
	if v472 == v471 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	base.MemoryCopy(m, v468, v21+int32(128), v470)
	goto L97
L96:
	;
	goto L97
L97:
	;
	if v472 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	base.MemoryCopy(m, v470+v468, v21+int32(96), v470)
	goto L100
L99:
	;
	goto L100
L100:
	;
	v484 = int32(1)
	F_CopyArrayEls(m, v459, v177, v179, v110, v173, v172&v484, base.I32_extend8_s(v174), v484)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	F_pfree(m, v177)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_pfree(m, v179)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	v497 = v459
	goto L38
L104:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v27
	F_errmsg(m, int32(_a_F_array_recv_1), v21)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1301), int32(_a_F_array_recv_3))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v27
	F_errmsg(m, int32(_a_F_array_recv_4), v21+int32(16))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1306), int32(_a_F_array_recv_3))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(_a_F_array_recv_5), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1312), int32(_a_F_array_recv_3))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L8
	} else {
		goto L117
	}
L117:
	;
	v577 = F_format_type_extended(m, v41, int32(-1), int32(2))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	v581 = F_format_type_extended(m, v24, int32(-1), int32(2))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v577
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v41
	F_errmsg(m, int32(_a_F_array_recv_6), v21+int32(80))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1340), int32(_a_F_array_recv_3))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L8
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L8
	} else {
		goto L123
	}
L123:
	;
	v604 = F_format_type_be(m, v24)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v604
	F_errmsg(m, int32(_a_F_array_recv_7), v21+int32(32))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1379), int32(_a_F_array_recv_3))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L8
	} else {
		goto L128
	}
L128:
	;
	F_errmsg(m, int32(_a_F_array_recv_8), int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1481), int32(_a_F_array_recv_9))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L8
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v190 + int32(1)
	F_errmsg(m, int32(_a_F_array_recv_10), v21-int32(-64))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1510), int32(_a_F_array_recv_9))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L8
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_recv_11), v21+int32(48))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L8
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(1534), int32(_a_F_array_recv_9))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L8
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_sort(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		v9 = F_array_sort_internal(m, v3, v7, v7, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_array_sort_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= int32(0) {
		v190 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L9
	} else {
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L52
	}
L3:
	;
	m.G0 = v14 + int32(32)
	return v190
L4:
	;
	v20 = l0 + int32(16)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 < int32(2) {
		v190 = l0
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v26 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v31 = F_MemoryContextAllocZero(m, v29, int32(60))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v37 = v26
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 != v39 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return int32(0)
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v31
	v37 = v31
	goto L8
L11:
	;
	v42 = F_lookup_type_cache(m, v38, int32(6))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v16 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v38
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v45)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+6)) = uint8(v47)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+7)) = uint8(v49)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+52)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+56)) = v55
	goto L13
L15:
	;
	v77 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_array_sort_internal[0]))
	v81 = F_tuplesort_begin_datum(m, v74, v73, v24, l2, v79, v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L27
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+56))
	if v60 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if l1 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v65 = int32(1073)
	goto L22
L21:
	;
	v65 = int32(1072)
	goto L22
L22:
	;
	v73 = v65
	v74 = v60
	goto L15
L23:
	;
	v68 = int32(52)
	goto L25
L24:
	;
	v68 = int32(48)
	goto L25
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37+v68)))
	if v70 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v73 = v70
	v74 = v38
	goto L15
L27:
	;
	v85 = F_array_create_iterator(m, l0, v16-int32(1), v37)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v91 = F_array_iterate(m, v85, v14+int32(28), v14+int32(27))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v91 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	goto L33
L31:
	;
	goto L32
L32:
	;
	F_array_free_iterator(m, v85)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L38
	}
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	F_tuplesort_putdatum(m, v81, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	v112 = F_array_iterate(m, v85, v14+int32(28), v14+int32(27))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v112 != 0 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	F_tuplesort_performsort(m, v81)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v131 = int32(0)
	v137 = F_tuplesort_getdatum(m, v81, int32(1), v131, v14+int32(28), v14+int32(27), v131)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	if v137 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v140 = v77
	goto L44
L42:
	;
	v166 = v77
	goto L43
L43:
	;
	F_tuplesort_end(m, v81)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L9
	} else {
		goto L49
	}
L44:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_array_sort_internal[1]))
	v154 = F_accumArrayResultAny(m, v140, v150, v151, v74, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L46
	}
L45:
	;
	v166 = v154
	goto L43
L46:
	;
	v157 = int32(0)
	v163 = F_tuplesort_getdatum(m, v81, int32(1), v157, v14+int32(28), v14+int32(27), v157)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	if v163 != 0 {
		v140 = v154
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_array_sort_internal[1]))
	v180 = F_makeArrayResultAny(m, v166, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v182 = F_pg_detoast_datum(m, v180)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32)+v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v182+v184<<(uint(int32(2))%32))+16)) = v188
	v190 = v182
	goto L3
L52:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v212 = F_format_type_be(m, v38)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v212
	F_errmsg(m, int32(_a_F_array_sort_internal_0), v14+int32(16))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_array_sort_internal_1), int32(1951), int32(_a_F_array_sort_internal_2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v232 = F_format_type_be(m, v38)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v232
	F_errmsg(m, int32(_a_F_array_sort_internal_3), v14)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_array_sort_internal_1), int32(1964), int32(_a_F_array_sort_internal_2))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_to_tsvector(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_deconstruct_array_builtin(m, v18, int32(25), v15+int32(12), v15+int32(8), v15+int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v32 <= int32(0) {
		v368 = v32
		v377 = int32(8)
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L105
	}
L5:
	;
	v380 = v377 + v368<<(uint(int32(2))%32)
	v381 = F_palloc0(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L91
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v38 = v2
	goto L8
L7:
	;
	v81 = int32(1)
	if v32 != v81 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v36))))
	if v50 == int32(1) {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v35+v38<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57&int32(-4) != int32(16) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v63 = v38 + int32(1)
	if v63 == v32 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L9
L14:
	;
	v38 = v63
	goto L8
L15:
	;
	F_errcode(m, int32(369098882))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_array_to_tsvector_0), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_array_to_tsvector_1), int32(776), int32(_a_F_array_to_tsvector_2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_pg_qsort(m, v84, v32, int32(4), int32(1522))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v246 = v32
	goto L21
L21:
	;
	if v246 <= int32(0) {
		v368 = v246
		v377 = int32(8)
		goto L5
	} else {
		goto L79
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v89) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v95 = v81
	v96 = int32(0)
	goto L26
L24:
	;
	v233 = v89
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v233
	v246 = v233
	goto L21
L26:
	;
	v106 = int32(2)
	v108 = v93 + v95<<(uint(v106)%32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v119 = int32(1)
	v120 = v118 + v119
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v123 = v121 & v119
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v93+v96<<(uint(v106)%32))))
	if v121 == v119 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v233 = v223 + int32(1)
	goto L25
L28:
	;
	v226 = v95 + int32(1)
	if v226 != v89 {
		v95 = v226
		v96 = v223
		goto L26
	} else {
		goto L78
	}
L29:
	;
	if v212 == int32(0) {
		v223 = v96
		goto L28
	} else {
		goto L76
	}
L30:
	;
	v152 = int32(1)
	v153 = v124 + v152
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v156 = v154 & v152
	if v154 == v152 {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v130 == int32(18) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v141 = int32(1)
	if v123 != 0 {
		v151 = int32(base.Ui32(v121)>>(uint(v141)%32)) - v141
		goto L30
	} else {
		goto L40
	}
L34:
	;
	v133 = int32(16)
	goto L36
L35:
	;
	v133 = int32(0)
	goto L36
L36:
	;
	if base.Ui32((v130-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v140 = int32(4)
	goto L39
L38:
	;
	v140 = v133
	goto L39
L39:
	;
	v151 = v140
	goto L30
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v151 = int32(base.Ui32(v145)>>(uint(int32(2))%32)) - int32(4)
	goto L30
L41:
	;
	if v151 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v162 == int32(18) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v173 = int32(1)
	if v156 != 0 {
		v183 = int32(base.Ui32(v154)>>(uint(v173)%32)) - v173
		goto L41
	} else {
		goto L51
	}
L45:
	;
	v165 = int32(16)
	goto L47
L46:
	;
	v165 = int32(0)
	goto L47
L47:
	;
	if base.Ui32((v162-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v172 = int32(4)
	goto L50
L49:
	;
	v172 = v165
	goto L50
L50:
	;
	v183 = v172
	goto L41
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v183 = int32(base.Ui32(v177)>>(uint(int32(2))%32)) - int32(4)
	goto L41
L52:
	;
	v187 = int32(0)
	if v187 < v183 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v183 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v190 = int32(-1)
	goto L57
L56:
	;
	v190 = v187
	goto L57
L57:
	;
	v212 = v190
	goto L29
L58:
	;
	v212 = base.B2i32(int32(0) < v151)
	goto L29
L59:
	;
	goto L60
L60:
	;
	if v123 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v212 = v210
	goto L29
L62:
	;
	v197 = v120
	goto L64
L63:
	;
	v197 = v118 + int32(4)
	goto L64
L64:
	;
	if v156 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v200 = v153
	goto L67
L66:
	;
	v200 = v124 + int32(4)
	goto L67
L67:
	;
	if base.Ui32(v151) < base.Ui32(v183) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v202 = v151
	goto L70
L69:
	;
	v202 = v183
	goto L70
L70:
	;
	v203 = F_memcmp(m, v197, v200, v202)
	mBase = m.M
	if v203 != 0 {
		v210 = v203
		goto L61
	} else {
		goto L71
	}
L71:
	;
	if v151 == v183 {
		v210 = int32(0)
		goto L61
	} else {
		goto L72
	}
L72:
	;
	if v151 < v183 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v209 = int32(-1)
	goto L75
L74:
	;
	v209 = int32(1)
	goto L75
L75:
	;
	v210 = v209
	goto L61
L76:
	;
	v216 = v96 + int32(1)
	if v95 == v216 {
		v223 = v95
		goto L28
	} else {
		goto L77
	}
L77:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v93+v216<<(uint(int32(2))%32)))) = v221
	v223 = v216
	goto L28
L78:
	;
	goto L27
L79:
	;
	v259 = v246 & int32(3)
	v260 = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v246) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v368 = v246
	v377 = v353 + int32(8)
	goto L5
L81:
	;
	v269 = v260
	v270 = v260
	v279 = v2
	goto L84
L82:
	;
	v313 = v260
	v314 = v260
	goto L83
L83:
	;
	v325 = v313
	v326 = v314
	v331 = v260
	goto L88
L84:
	;
	v280 = int32(2)
	v282 = v261 + v269<<(uint(v280)%32)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v304 = v270 + int32(base.Ui32(v284)>>(uint(v280)%32)) + int32(base.Ui32(v289)>>(uint(v280)%32)) + int32(base.Ui32(v294)>>(uint(v280)%32)) + int32(base.Ui32(v299)>>(uint(v280)%32)) - int32(16)
	v305 = int32(4)
	v306 = v269 + v305
	v308 = v279 + v305
	if v308 != v246&int32(2147483644) {
		v269 = v306
		v270 = v304
		v279 = v308
		goto L84
	} else {
		goto L86
	}
L85:
	;
	if v259 == int32(0) {
		v353 = v304
		goto L80
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v313 = v306
	v314 = v304
	goto L83
L88:
	;
	v336 = int32(2)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v261+v325<<(uint(v336)%32))))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v345 = v326 + int32(base.Ui32(v340)>>(uint(v336)%32)) - int32(4)
	v346 = int32(1)
	v349 = v331 + v346
	if v349 != v259 {
		v325 = v325 + v346
		v326 = v345
		v331 = v349
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v353 = v345
	goto L80
L90:
	;
	goto L89
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = v380 << (uint(int32(2)) % 32)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v381)+4)) = v386
	if int32(0) < v386 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v391 = v381 + int32(8)
	v397 = v391 + v386<<(uint(int32(2))%32)
	v399 = int32(0)
	goto L95
L93:
	;
	goto L94
L94:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v452 != v18 {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v408 = int32(2)
	v409 = v399 << (uint(v408) % 32)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v409+v410)))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v417 = int32(base.Ui32(v413)>>(uint(v408)%32)) - int32(4)
	if v417 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L94
L97:
	;
	base.MemoryCopy(m, v397, v412+int32(4), v417)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v422 = int32(1)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v391+v409))) = v417<<(uint(v422)%32)&int32(4094) | (v397-(v391+v426<<(uint(int32(2))%32)))<<(uint(int32(12))%32)
	v437 = v399 + v422
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v437 < v438 {
		v397 = v397 + v417
		v399 = v437
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L96
L101:
	;
	F_pfree(m, v18)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	m.G0 = v15 + int32(16)
	return v381
L104:
	;
	goto L103
L105:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_array_to_tsvector_3), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_array_to_tsvector_1), int32(771), int32(_a_F_array_to_tsvector_2))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	v14 = m.G0
	v15 = int32(16)
	v16 = v14 - v15
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = l0 + v15
	v21 = F_ArrayGetNItemsSafe(m, v18, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = F_palloc(m, v21<<(uint(int32(2))%32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v25
	if l5 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = F_palloc0(m, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v31 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v21
	if int32(0) < v21 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v28
	v31 = v28
	goto L6
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v35 << (uint(int32(3)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
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
	m.G0 = v16 + int32(16)
	return
L11:
	;
	v41 = v20 + v37
	goto L13
L12:
	;
	v41 = int32(0)
	goto L13
L13:
	;
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = v40
	goto L16
L15:
	;
	v46 = (v37 + int32(23)) & int32(-8)
	goto L16
L16:
	;
	v48 = int32(1)
	v54 = v48
	v58 = l0 + v46
	v59 = v41
	v60 = int32(0)
	goto L17
L17:
	;
	if v59 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L10
L19:
	;
	v173 = int32(1)
	v175 = v54 << (uint(v173) % 32)
	v177 = base.B2i32(v175 == int32(256))
	if v175 == int32(256) {
		goto L59
	} else {
		goto L60
	}
L20:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v54&v69 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v60<<(uint(int32(2))%32)))) = int32(0)
	if v31 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v31))) = uint8(v77)
	v172 = v58
	goto L19
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_deconstruct_array_0), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_deconstruct_array_1), int32(3669), int32(_a_F_deconstruct_array_2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	switch l3 - int32(99) {
	case 0:
		v172 = v158
		goto L19
	case 1:
		goto L57
	default:
		goto L56
	case 6:
		goto L58
	}
L31:
	;
	switch l1 - v48 {
	case 0:
		goto L35
	case 1:
		goto L38
	default:
		goto L36
	case 3:
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v60<<(uint(int32(2))%32)))) = v58
	if int32(0) < l1 {
		v158 = l1 + v58
		goto L30
	} else {
		goto L42
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v60<<(uint(int32(2))%32)))) = v114
	v158 = l1 + v58
	goto L30
L35:
	;
	v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v58))))
	v114 = v113
	goto L34
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v114 = v99
	goto L34
L38:
	;
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58))))
	v114 = v98
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg_internal(m, int32(_a_F_deconstruct_array_3), v16)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_deconstruct_array_4), int32(70), int32(_a_F_deconstruct_array_5))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	if l1 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v126 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v153 = F_strlen(m, v58)
	mBase = m.M
	v158 = v153 + v58 + int32(1)
	goto L30
L46:
	;
	v130 = int32(18)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v132 == v130 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v144 = int32(1)
	if v126&v144 != 0 {
		v158 = v58 + int32(base.Ui32(v126)>>(uint(v144)%32))
		goto L30
	} else {
		goto L55
	}
L49:
	;
	v135 = v130
	goto L51
L50:
	;
	v135 = int32(2)
	goto L51
L51:
	;
	if base.Ui32((v132-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v142 = int32(6)
	goto L54
L53:
	;
	v142 = v135
	goto L54
L54:
	;
	v158 = v58 + v142
	goto L30
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v158 = v58 + int32(base.Ui32(v149)>>(uint(int32(2))%32))
	goto L30
L56:
	;
	v172 = (v158 + int32(1)) & int32(-2)
	goto L19
L57:
	;
	v172 = (v158 + int32(7)) & int32(-8)
	goto L19
L58:
	;
	v172 = (v158 + int32(3)) & int32(-4)
	goto L19
L59:
	;
	v178 = v173
	goto L61
L60:
	;
	v178 = v175
	goto L61
L61:
	;
	if v59 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v179 = v178
	goto L64
L63:
	;
	v179 = v54
	goto L64
L64:
	;
	if v59 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v182 = v177 + v59
	goto L67
L66:
	;
	v182 = int32(0)
	goto L67
L67:
	;
	v184 = v60 + int32(1)
	if v184 != v21 {
		v54 = v179
		v58 = v172
		v59 = v182
		v60 = v184
		goto L17
	} else {
		goto L68
	}
L68:
	;
	goto L18
}
func F_deconstruct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(99)
	v15 = int32(1)
	switch l1 - int32(18) {
	case 0:
		v51 = v15
		v52 = v14
		v53 = v15
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 1, 2, 4, 6:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
			F_errmsg_internal(m, int32(_a_F_deconstruct_array_builtin_0), v12)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_deconstruct_array_builtin_1), int32(3756), int32(_a_F_deconstruct_array_builtin_2))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v51 = int32(2)
		v52 = int32(115)
		v53 = v15
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 5, 8:
		v51 = int32(4)
		v52 = int32(105)
		v53 = v15
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 7:
		v51 = int32(-1)
		v52 = int32(105)
		v53 = int32(0)
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 9:
		v51 = int32(6)
		v52 = int32(115)
		v53 = int32(0)
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	default:
		if l1 == int32(701) {
			v51 = int32(8)
			v52 = int32(100)
			v53 = int32(0)
			F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			if l1 != int32(2275) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
					F_errmsg_internal(m, int32(_a_F_deconstruct_array_builtin_0), v12)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_deconstruct_array_builtin_1), int32(3756), int32(_a_F_deconstruct_array_builtin_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = int32(-2)
				v52 = v14
				v53 = int32(0)
				F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	}
}
func F_expand_array(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v17 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_expand_array_0), int32(0), int32(1024), int32(_a_F_expand_array_1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = F_MemoryContextAlloc(m, v17, int32(80))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v25 = int32(513)
			*(*uint16)(unsafe.Add(mBase, uint32(v22)+18)) = uint16(v25)
			v27 = int32(769)
			*(*uint16)(unsafe.Add(mBase, uint32(v22)+12)) = uint16(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(_a_F_expand_array_2)
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v22)+14)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(689375833)
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v37 != int32(1) {
				v119 = l2
				v121 = int32(_a_F_expand_array_3)
				v122 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v17
				v125 = F_pg_detoast_datum_copy(m, l0)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v122
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
					v131 = v125 + int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v129
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v131 + v134<<(uint(int32(2))%32)
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v139
					if v119 != 0 {
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
						if v139 == v141 {
							v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)) = uint16(v143)
							v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
							*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)) = uint8(v145)
							v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
							*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)) = uint8(v147)
							*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
							v175 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
							*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
							*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
							v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
							if v180 != 0 {
								v188 = v180
							} else {
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
								v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
							v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
							m.G0 = v11 + int32(48)
							return v22 + int32(12)
						} else {
							F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return int32(0)
							} else {
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v119))) = v157
								v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)) = uint16(v159)
								v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)))
								*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)) = uint8(v161)
								v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)))
								*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)) = uint8(v163)
								*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
								v175 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
								*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
								*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
								if v180 != 0 {
									v188 = v180
								} else {
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
									v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
								v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
								m.G0 = v11 + int32(48)
								return v22 + int32(12)
							}
						}
					} else {
						F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
						mBase = m.M
						v172 = m.ExcPending
						if v172 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
							v175 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
							*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
							*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
							v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
							if v180 != 0 {
								v188 = v180
							} else {
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
								v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
							v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
							m.G0 = v11 + int32(48)
							return v22 + int32(12)
						}
					}
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v40&int32(254) != int32(2) {
					v119 = l2
					v121 = int32(_a_F_expand_array_3)
					v122 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v17
					v125 = F_pg_detoast_datum_copy(m, l0)
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v122
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
						v131 = v125 + int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v131
						*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v129
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v131 + v134<<(uint(int32(2))%32)
						v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v139
						if v119 != 0 {
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
							if v139 == v141 {
								v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)) = uint16(v143)
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)) = uint8(v145)
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)) = uint8(v147)
								*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
								v175 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
								*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
								*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
								if v180 != 0 {
									v188 = v180
								} else {
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
									v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
								v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
								m.G0 = v11 + int32(48)
								return v22 + int32(12)
							} else {
								F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return int32(0)
								} else {
									v157 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v119))) = v157
									v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)))
									*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)) = uint16(v159)
									v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)))
									*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)) = uint8(v161)
									v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)))
									*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)) = uint8(v163)
									*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
									v175 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
									*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
									*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
									if v180 != 0 {
										v188 = v180
									} else {
										v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
										v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
									v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
									m.G0 = v11 + int32(48)
									return v22 + int32(12)
								}
							}
						} else {
							F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
								v175 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
								*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
								*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
								if v180 != 0 {
									v188 = v180
								} else {
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
									v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
								v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
								m.G0 = v11 + int32(48)
								return v22 + int32(12)
							}
						}
					}
				} else {
					if l2 != 0 {
						v45 = l2
					} else {
						v45 = v11
					}
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v45))) = v47
					v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+44)))
					*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)) = uint16(v49)
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+46)))
					*(*uint8)(unsafe.Add(mBase, uint32(v45)+6)) = uint8(v51)
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+47)))
					*(*uint8)(unsafe.Add(mBase, uint32(v45)+7)) = uint8(v53)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+46)))
					if v55 != int32(1) {
						v119 = v45
						v121 = int32(_a_F_expand_array_3)
						v122 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v17
						v125 = F_pg_detoast_datum_copy(m, l0)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v122
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
							v131 = v125 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v129
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v131 + v134<<(uint(int32(2))%32)
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v139
							if v119 != 0 {
								v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
								if v139 == v141 {
									v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)) = uint16(v143)
									v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
									*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)) = uint8(v145)
									v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
									*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)) = uint8(v147)
									*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
									v175 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
									*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
									*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
									if v180 != 0 {
										v188 = v180
									} else {
										v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
										v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
									v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
									m.G0 = v11 + int32(48)
									return v22 + int32(12)
								} else {
									F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return int32(0)
									} else {
										v157 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v119))) = v157
										v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)))
										*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)) = uint16(v159)
										v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)))
										*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)) = uint8(v161)
										v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)))
										*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)) = uint8(v163)
										*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
										v175 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
										*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
										*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
										v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
										if v180 != 0 {
											v188 = v180
										} else {
											v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
											v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
										v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
										m.G0 = v11 + int32(48)
										return v22 + int32(12)
									}
								}
							} else {
								F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
									v175 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
									*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
									*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
									if v180 != 0 {
										v188 = v180
									} else {
										v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
										v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
									v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
									m.G0 = v11 + int32(48)
									return v22 + int32(12)
								}
							}
						}
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
						if v58 == int32(0) {
							v119 = v45
							v121 = int32(_a_F_expand_array_3)
							v122 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v17
							v125 = F_pg_detoast_datum_copy(m, l0)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v122
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
								v131 = v125 + int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v131
								*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v129
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v131 + v134<<(uint(int32(2))%32)
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v139
								if v119 != 0 {
									v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
									if v139 == v141 {
										v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)) = uint16(v143)
										v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
										*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)) = uint8(v145)
										v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
										*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)) = uint8(v147)
										*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
										v175 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
										*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
										*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
										v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
										if v180 != 0 {
											v188 = v180
										} else {
											v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
											v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
										v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
										m.G0 = v11 + int32(48)
										return v22 + int32(12)
									} else {
										F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v157 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v119))) = v157
											v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)))
											*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)) = uint16(v159)
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)))
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)) = uint8(v161)
											v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)))
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)) = uint8(v163)
											*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
											v175 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
											*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
											*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
											v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
											if v180 != 0 {
												v188 = v180
											} else {
												v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
												v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
											v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
											*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
											m.G0 = v11 + int32(48)
											return v22 + int32(12)
										}
									}
								} else {
									F_get_typlenbyvalalign(m, v139, v22+int32(44), v22+int32(46), v22+int32(47))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
										v175 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v175
										*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v175
										*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v125
										v180 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
										if v180 != 0 {
											v188 = v180
										} else {
											v181 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
											v188 = (v181<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v188 + v125
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
										v201 = v125 + int32(base.Ui32(v191)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
										m.G0 = v11 + int32(48)
										return v22 + int32(12)
									}
								}
							}
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v46)+56))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v62
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
							v67 = F_MemoryContextAlloc(m, v64, v62<<(uint(int32(3))%32))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v67
								v71 = v62 << (uint(int32(2)) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v67 + v71
								v74 = int32(0)
								v75 = base.B2i32(v71 == v74)
								if v75 == v74 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
									base.MemoryCopy(m, v67, v78, v71)
								} else {
								}
								if v75 == int32(0) {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v46)+36))
									base.MemoryCopy(m, v82, v83, v71)
								} else {
								}
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v85
								v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)) = uint16(v87)
								v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+46)))
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)) = uint8(v89)
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+47)))
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)) = uint8(v91)
								v94 = v61 << (uint(int32(2)) % 32)
								v95 = F_MemoryContextAlloc(m, v64, v94)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v95
									if v94 != 0 {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
										base.MemoryCopy(m, v95, v98, v94)
									} else {
									}
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
									if v100 != 0 {
										v101 = F_MemoryContextAlloc(m, v64, v61)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v101
											if v61 == int32(0) {
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
												base.MemoryCopy(m, v101, v106, v61)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v61
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v46)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v112
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
											*(*int64)(unsafe.Add(mBase, uint32(v22)+68)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v114
											v201 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
											m.G0 = v11 + int32(48)
											return v22 + int32(12)
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v61
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v46)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v112
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v22)+68)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v114
										v201 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v201
										m.G0 = v11 + int32(48)
										return v22 + int32(12)
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
func F_get_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 <= v9 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v16 = v11 + v8<<(uint(int32(2))%32) - int32(4)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 + int32(1)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v22 = v21
	} else {
		v22 = v9
	}
	if v22 < v8 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v25 = int32(1)
		v26 = v8 - v25
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v26))))
		if v28 != v25 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v31 == int32(0) {
				return int32(0)
			} else {
				v35 = v26 << (uint(int32(2)) % 32)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v35+v31)))
				if v38 != v40 {
					return int32(0)
				} else {
					if v8 < v22 {
						v44 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v8+v24))) = uint8(v44)
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
						if v51 != int32(1) {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
							if v54 != int32(1) {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61
								return int32(0)
							} else {
								v57 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v57)
								return int32(0)
							}
						}
					}
				}
			}
		}
	}
}
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	v16 = int32(_a_F_makeArrayResult_0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_makeArrayResult[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResult[0])) = l1
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+27)))
	v32 = F_construct_md_array(m, v20, v21, base.B2i32(int32(0) < v11), v9+int32(12), v9+int32(8), v28, v29, v30, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResult[0])) = v17
		if v15 == int32(1) {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_MemoryContextDelete(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v32
			}
		} else {
			m.G0 = v9 + int32(16)
			return v32
		}
	}
}
func F_makeArrayTypeName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v13 = F_makeObjectName(m, int32(_a_F_makeArrayTypeName_0), l0, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = int32(0)
	v19 = F_SearchSysCacheExists(m, int32(81), v13, l1, v17, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v13
	v25 = v3
	goto L7
L5:
	;
	v47 = v13
	goto L6
L6:
	;
	m.G0 = v8 + int32(80)
	return v47
L7:
	;
	F_pfree(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v47 = v39
	goto L6
L9:
	;
	v29 = v25 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
	v32 = v8 + int32(16)
	v35 = F_pg_snprintf(m, v32, int32(64), int32(_a_F_makeArrayTypeName_1), v8)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v39 = F_makeObjectName(m, int32(_a_F_makeArrayTypeName_0), l0, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v41 = int32(0)
	v43 = F_SearchSysCacheExists(m, int32(81), v39, l1, v41, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v43 != 0 {
		v23 = v39
		v25 = v29
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
}
func F_parse_array_element(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v12 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, v9, base.B2i32(v6 == int32(11)))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v33 = v12
				return v33
			} else {
				switch v6 - int32(3) {
				case 0:
					v19 = F_parse_object(m, l0, l1)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v25 = v19
						if v25 != 0 {
							v33 = v25
							return v33
						} else {
							if v7 != 0 {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										v33 = v29
									} else {
										v33 = int32(0)
									}
									return v33
								}
							} else {
								v33 = int32(0)
								return v33
							}
						}
					}
				default:
					v23 = F_parse_scalar(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = v23
						if v25 != 0 {
							v33 = v25
							return v33
						} else {
							if v7 != 0 {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										v33 = v29
									} else {
										v33 = int32(0)
									}
									return v33
								}
							} else {
								v33 = int32(0)
								return v33
							}
						}
					}
				case 2:
					v21 = F_parse_array(m, l0, l1)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v25 = v21
						if v25 != 0 {
							v33 = v25
							return v33
						} else {
							if v7 != 0 {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										v33 = v29
									} else {
										v33 = int32(0)
									}
									return v33
								}
							} else {
								v33 = int32(0)
								return v33
							}
						}
					}
				}
			}
		}
	} else {
		switch v6 - int32(3) {
		case 0:
			v19 = F_parse_object(m, l0, l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v25 = v19
				if v25 != 0 {
					v33 = v25
					return v33
				} else {
					if v7 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = v29
							} else {
								v33 = int32(0)
							}
							return v33
						}
					} else {
						v33 = int32(0)
						return v33
					}
				}
			}
		default:
			v23 = F_parse_scalar(m, l0, l1)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v23
				if v25 != 0 {
					v33 = v25
					return v33
				} else {
					if v7 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = v29
							} else {
								v33 = int32(0)
							}
							return v33
						}
					} else {
						v33 = int32(0)
						return v33
					}
				}
			}
		case 2:
			v21 = F_parse_array(m, l0, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v25 = v21
				if v25 != 0 {
					v33 = v25
					return v33
				} else {
					if v7 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = v29
							} else {
								v33 = int32(0)
							}
							return v33
						}
					} else {
						v33 = int32(0)
						return v33
					}
				}
			}
		}
	}
}
