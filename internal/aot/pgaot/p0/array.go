package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_array_agg_array_serialize(m *base.Module, l0 int32) int64 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
		F_enlargeStringInfo(m, v7, int32(4))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
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
				return int64(0)
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
					return int64(0)
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
						return int64(0)
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						F_enlargeStringInfo(m, v7, int32(4))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int64(0)
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
								return int64(0)
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
										return int64(0)
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										F_enlargeStringInfo(m, v7, int32(4))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int64(0)
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
												return int64(0)
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
													return int64(0)
												} else {
													F_appendBinaryStringInfo(m, v7, v9+int32(56), int32(24))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int64(0)
													} else {
														v177 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v177))) = v178 << (uint(int32(2)) % 32)
														m.G0 = v7 + int32(16)
														return base.I64_extend_i32_u(v177)
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
										return int64(0)
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
											return int64(0)
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
												return int64(0)
											} else {
												F_appendBinaryStringInfo(m, v7, v9+int32(56), int32(24))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int64(0)
												} else {
													v177 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
													v178 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v177))) = v178 << (uint(int32(2)) % 32)
													m.G0 = v7 + int32(16)
													return base.I64_extend_i32_u(v177)
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
func F_array_append(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v21 int64
	_ = v21
	var v22 int64
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
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_fetch_array_arg_replace_nulls(m, l0, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v18 == int32(0) {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			v22 = v21
		} else {
			v22 = int64(0)
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		switch v23 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(1)
			v71 = int32(12)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
			v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+4)))
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
			v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+7)))
			v83 = F_array_set_element(m, base.I64_extend_i32_u(v14+v71), int32(1), v11+v71, v22, v18, int32(-1), v80, v81, v82)
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int64(0)
			} else {
				m.G0 = v11 + int32(16)
				return v83
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
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
				v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+4)))
				v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
				v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+7)))
				v83 = F_array_set_element(m, base.I64_extend_i32_u(v14+v71), int32(1), v11+v71, v22, v18, int32(-1), v80, v81, v82)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					m.G0 = v11 + int32(16)
					return v83
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_array_append_0), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_array_append_1), int32(168), int32(_a_F_array_append_2))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_array_append_3), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_append_1), int32(175), int32(_a_F_array_append_2))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
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
	}
}
func F_array_cardinality(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_DatumGetAnyArrayP(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v10 == int32(-1) {
			v13 = int32(28)
		} else {
			v13 = int32(4)
		}
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4+v13)))
		if v10 == int32(-1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
			v21 = v18
		} else {
			v21 = v4 + int32(16)
		}
		v22 = F_ArrayGetNItemsSafe(m, v15, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_s(v22)
		}
	}
}
func F_array_contain_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v252 int32
	_ = v252
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(112)
	return v252
L2:
	;
	v252 = int32(0)
	goto L1
L3:
	;
	v22 = int32(40)
	goto L5
L4:
	;
	v22 = int32(12)
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0+v22)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v27 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = int32(40)
	goto L8
L7:
	;
	v30 = int32(12)
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1+v30)))
	if v24 == v32 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v34 != 0 {
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
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L68
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L18
	} else {
		goto L63
	}
L13:
	;
	v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+11)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+10)))
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+8)))
	if v48 == int32(-1) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 == v24 {
		v47 = v34
		v48 = v27
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v38 = F_lookup_type_cache(m, v24, int32(32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v47 = v38
	v48 = v46
	goto L13
L21:
	;
	v72 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+74)) = uint16(v72)
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+72)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v15)+60)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v47 + int32(76)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v84 == int32(-1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_deconstruct_expanded_array(m, l1)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_deconstruct_array(m, l1, v51, v50&int32(1), v49, v15+int32(52), v15+int32(48), v15+int32(44))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v60
	goto L21
L26:
	;
	goto L21
L27:
	;
	v87 = int32(28)
	goto L29
L28:
	;
	v87 = int32(4)
	goto L29
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0+v87)))
	if v84 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v96 = F_ArrayGetNItemsSafe(m, v89, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L18
	} else {
		goto L34
	}
L31:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v95 = v92
	goto L30
L32:
	;
	goto L33
L33:
	;
	v95 = l0 + int32(16)
	goto L30
L34:
	;
	F_array_iter_setup(m, v15+int32(16), l0, v51, v50&int32(1), v49)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	if v96 <= int32(0) {
		v252 = l3
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v118 = int32(0)
	goto L37
L37:
	;
	v126 = F_array_iter_next(m, v15+int32(16), v15+int32(15), v118)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L39
	}
L38:
	;
	v252 = l3
	goto L1
L39:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v128 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v200 = v118 + int32(1)
	if v200 != v96 {
		v118 = v200
		goto L37
	} else {
		goto L62
	}
L41:
	;
	if l3 != 0 {
		goto L40
	} else {
		goto L61
	}
L42:
	;
	v131 = int32(0)
	if v131 < v108 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if l3 != 0 {
		goto L2
	} else {
		goto L60
	}
L45:
	;
	v134 = v131
	goto L48
L46:
	;
	goto L47
L47:
	;
	if l3 == int32(0) {
		goto L40
	} else {
		goto L59
	}
L48:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v107+v134<<(uint(int32(3))%32))))
	if v106 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L47
L50:
	;
	v170 = v134 + int32(1)
	if v170 != v108 {
		v134 = v170
		goto L48
	} else {
		goto L58
	}
L51:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v106))))
	if v151 != 0 {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+104)) = uint8(v152)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v149
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+88)) = uint8(v152)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+72)) = uint8(v152)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = m.T0[v163].(func(*base.Module, int32) int64)(m, v15+int32(56))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L18
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+72)))
	if v166 != 0 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	if v164 != int64(0) {
		goto L41
	} else {
		goto L57
	}
L57:
	;
	goto L50
L58:
	;
	goto L49
L59:
	;
	goto L2
L60:
	;
	goto L40
L61:
	;
	v252 = int32(1)
	goto L1
L62:
	;
	goto L38
L63:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	v209 = F_format_type_be(m, v24)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v209
	F_errmsg(m, int32(_a_F_array_contain_compare_0), v15)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_array_contain_compare_1), int32(_a_F_array_contain_compare_2), int32(_a_F_array_contain_compare_3))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_array_contain_compare_4), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_array_contain_compare_1), int32(_a_F_array_contain_compare_5), int32(_a_F_array_contain_compare_3))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L18
	} else {
		goto L71
	}
L71:
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
func F_array_fill_with_lower_bounds(m *base.Module, l0 int32) int64 {
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
	var v22 int64
	_ = v22
	var v23 int64
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
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v7 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_array_fill_with_lower_bounds_0), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_array_fill_with_lower_bounds_1), int32(_a_F_array_fill_with_lower_bounds_2), int32(_a_F_array_fill_with_lower_bounds_3))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
		if v8 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(67108994))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_array_fill_with_lower_bounds_0), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_fill_with_lower_bounds_1), int32(_a_F_array_fill_with_lower_bounds_2), int32(_a_F_array_fill_with_lower_bounds_3))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v12 = F_pg_detoast_datum(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v17 = F_pg_detoast_datum(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v19 == int32(0) {
						v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						v23 = v22
					} else {
						v23 = int64(0)
					}
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = F_get_fn_expr_argtype(m, v24, int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						if v26 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_array_fill_with_lower_bounds_4), int32(0))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_array_fill_with_lower_bounds_1), int32(_a_F_array_fill_with_lower_bounds_5), int32(_a_F_array_fill_with_lower_bounds_3))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
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
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v30)
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_ArrayGetNItemsSafe(m, v9, l0+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 < int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
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
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L27
	}
L7:
	;
	if v20 != 0 {
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
	if int32(0) < v12 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return v84
L13:
	;
	if v19 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v84 = int32(0)
	goto L12
L16:
	;
	v32 = v19
	goto L18
L17:
	;
	v32 = (v16<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L18:
	;
	v35 = l0 + v32
	v39 = v12
	goto L19
L19:
	;
	v45 = F_DirectFunctionCall2Coll(m, l1, int32(0), base.I64_extend_i32_u(v35), base.I64_extend_i32_u(l2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	if v45 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if l3 == int32(0) {
		v84 = int32(1)
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v63 = int32(1)
	if v63 < v39 {
		v35 = v35 + (int32(base.Ui32(v55)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v39 = v39 - v63
		goto L19
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v35
	return int32(1)
L26:
	;
	goto L20
L27:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_array_iterator_0), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_array_iterator_1), int32(46), int32(_a_F_array_iterator_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_array_iterator_3), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_array_iterator_1), int32(50), int32(_a_F_array_iterator_2))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
func F_array_prepend_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	v3 = int64(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v5 != int32(472) {
		v22 = v3
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v10 == int32(0) {
			v22 = v3
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(8) {
				v22 = v3
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v16 != 0 {
					v22 = v3
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
					if v17 != v18 {
						v22 = v3
					} else {
						v22 = base.I64_extend_i32_u(v10)
					}
				}
			}
		}
	}
	return v22
}
func F_array_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v265 int64
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	v20 = m.G0
	v22 = v20 - int32(192)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = F_pq_getmsgint(m, v26, int32(4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L8
	} else {
		goto L138
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L8
	} else {
		goto L134
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L8
	} else {
		goto L130
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L8
	} else {
		goto L125
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L8
	} else {
		goto L119
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L8
	} else {
		goto L115
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L8
	} else {
		goto L111
	}
L8:
	;
	return int64(0)
L9:
	;
	if int32(0) <= v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v28) {
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
	v488 = m.ExcPending
	if v488 != 0 {
		goto L8
	} else {
		goto L107
	}
L13:
	;
	v37 = F_pq_getmsgint(m, v26, int32(4))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v37) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v42 = F_pq_getmsgint(m, v26, int32(4))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v45 = int32(_a_F_array_recv_0)
	if base.B2i32(base.B2i32(v42 == v25)|base.B2i32(base.Ui32(v45) < base.Ui32(v42)) == int32(0))&base.B2i32(base.Ui32(v25) <= base.Ui32(v45)) != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v28 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v55 = int32(0)
	goto L21
L19:
	;
	goto L20
L20:
	;
	v112 = v22 + int32(144)
	v113 = F_ArrayGetNItemsSafe(m, v28, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L26
	}
L21:
	;
	v74 = v55 << (uint(int32(2)) % 32)
	v79 = F_pq_getmsgint(m, v26, int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74+(v22+int32(144))))) = v79
	v86 = F_pq_getmsgint(m, v26, int32(4))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(112)+v74))) = v86
	v90 = v55 + int32(1)
	if v90 != v28 {
		v55 = v90
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	F_ArrayCheckBounds(m, v28, v112, v22+int32(112))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	if v120 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v113 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	F_get_type_io_data(m, v25, int32(2), v136+int32(4), v136+int32(6), v136+int32(7), v136+int32(8), v136+int32(12), v136+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L35
	}
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
	v125 = F_MemoryContextAlloc(m, v123, int32(48))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v134 == v25 {
		v162 = v120
		goto L28
	} else {
		goto L34
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v125
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v25 ^ int32(-1)
	v136 = v130
	goto L29
L34:
	;
	v136 = v120
	goto L29
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	if v152 == int32(0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	F_fmgr_info_cxt(m, v152, v136+int32(20), v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v25
	v162 = v136
	goto L28
L38:
	;
	m.G0 = v22 + int32(192)
	return base.I64_extend_i32_u(v464)
L39:
	;
	v167 = F_palloc0(m, int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v162)+4)))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
	v178 = base.I32_extend8_s(v177)
	v181 = F_palloc(m, v113<<(uint(int32(3))%32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = int64(64)
	v464 = v167
	goto L38
L43:
	;
	v183 = F_palloc(m, v113)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	switch v177 - int32(99) {
	case 0:
		v206 = int32(1)
		goto L45
	case 1:
		goto L48
	default:
		goto L47
	case 6:
		goto L49
	case 16:
		goto L46
	}
L45:
	;
	v207 = int32(0)
	if v113 <= v207 {
		v385 = v207
		goto L54
	} else {
		goto L55
	}
L46:
	;
	v206 = int32(2)
	goto L45
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	v206 = int32(8)
	goto L45
L49:
	;
	v206 = int32(4)
	goto L45
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v178
	F_errmsg_internal(m, int32(_a_F_array_recv_1), v22+int32(48))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_array_recv_2), int32(322), int32(_a_F_array_recv_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v427 = F_palloc0(m, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L97
	}
L54:
	;
	v425 = int32(0)
	v426 = v385 + (v28<<(uint(int32(3))%32)+int32(23))&int32(120)
	goto L53
L55:
	;
	v211 = v162 + int32(20)
	v216 = v207
	goto L56
L56:
	;
	v232 = F_pq_getmsgint(m, v26, int32(4))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L58
	}
L57:
	;
	v277 = int32(0)
	v287 = v277
	v288 = v277
	v292 = v277
	goto L70
L58:
	;
	if v232 < int32(-1) {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v236-v237 < v232 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v232 == int32(-1) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v275 = v216 + int32(1)
	if v275 != v113 {
		v216 = v275
		goto L56
	} else {
		goto L68
	}
L62:
	;
	v246 = F_ReceiveFunctionCall(m, v211, int32(0), v174, v24)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v237 + v252
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v232 + v237
	v265 = F_ReceiveFunctionCall(m, v211, v22+int32(176), v174, v24)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L8
	} else {
		goto L66
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v181+v216<<(uint(int32(3))%32)))) = v246
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v216+v183))) = uint8(v250)
	goto L61
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v181+v216<<(uint(int32(3))%32)))) = v265
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216+v183))) = uint8(v269)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v22)+188))
	if v271 != v232 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	goto L57
L69:
	;
	v372 = base.I32_div_s(v113+int32(7), int32(8))
	v379 = (v372 + v28<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v425 = v379
	v426 = v366 + v379
	goto L53
L70:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v183))))
	if v304 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v292 == int32(0) {
		v385 = v356
		goto L54
	} else {
		goto L96
	}
L72:
	;
	v305 = int32(1)
	v307 = v287 + v305
	if v307 == v113 {
		v366 = v288
		goto L69
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v176 != int32(-1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v287 = v307
	v292 = v305
	goto L70
L76:
	;
	v356 = (v288 + (v206 - int32(1)) + v352) & (v277 - v206)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v356) {
		goto L2
	} else {
		goto L94
	}
L77:
	;
	if int32(0) < v176 {
		v352 = v176
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v322 = v181 + v287<<(uint(int32(3))%32)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v324 = F_pg_detoast_datum(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L8
	} else {
		goto L81
	}
L80:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v181+v287<<(uint(int32(3))%32))))
	v317 = F_strlen(m, v316)
	mBase = m.M
	v352 = v317 + int32(1)
	goto L76
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v322))) = base.I64_extend_i32_u(v324)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v328 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v332 = int32(18)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+1)))
	if v334 == v332 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	if v328&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	v337 = v332
	goto L87
L86:
	;
	v337 = int32(2)
	goto L87
L87:
	;
	if base.Ui32((v334-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v344 = int32(6)
	goto L90
L89:
	;
	v344 = v337
	goto L90
L90:
	;
	v352 = v344
	goto L76
L91:
	;
	v352 = int32(base.Ui32(v328) >> (uint(int32(1)) % 32))
	goto L76
L92:
	;
	goto L93
L93:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v352 = int32(base.Ui32(v349) >> (uint(int32(2)) % 32))
	goto L76
L94:
	;
	v360 = v287 + int32(1)
	if v360 != v113 {
		v287 = v360
		v288 = v356
		goto L70
	} else {
		goto L95
	}
L95:
	;
	goto L71
L96:
	;
	v366 = v356
	goto L69
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v427)+8)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(v427)+4)) = v28
	v432 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v426 << (uint(v432) % 32)
	v436 = v427 + int32(16)
	v438 = v28 << (uint(v432) % 32)
	v439 = int32(0)
	v440 = base.B2i32(v438 == v439)
	if v440 == v439 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	base.MemoryCopy(m, v436, v22+int32(144), v438)
	goto L100
L99:
	;
	goto L100
L100:
	;
	if v440 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryCopy(m, v438+v436, v22+int32(112), v438)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v452 = int32(1)
	F_CopyArrayEls(m, v427, v181, v183, v113, v176, v175&v452, v178, v452)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_pfree(m, v181)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_pfree(m, v183)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	v464 = v427
	goto L38
L107:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v28
	F_errmsg(m, int32(_a_F_array_recv_4), v22)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1305), int32(_a_F_array_recv_6))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v28
	F_errmsg(m, int32(_a_F_array_recv_7), v22+int32(16))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1310), int32(_a_F_array_recv_6))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(_a_F_array_recv_8), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L8
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1316), int32(_a_F_array_recv_6))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	v546 = F_format_type_extended(m, v42, int32(-1), int32(2))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L8
	} else {
		goto L121
	}
L121:
	;
	v550 = F_format_type_extended(m, v25, int32(-1), int32(2))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v42
	F_errmsg(m, int32(_a_F_array_recv_9), v22+int32(96))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L8
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1344), int32(_a_F_array_recv_6))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	v573 = F_format_type_be(m, v25)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v573
	F_errmsg(m, int32(_a_F_array_recv_10), v22+int32(32))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L8
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1383), int32(_a_F_array_recv_6))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L8
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v216 + int32(1)
	F_errmsg(m, int32(_a_F_array_recv_11), v22+int32(80))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1515), int32(_a_F_array_recv_12))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L8
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_recv_13), v22-int32(-64))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L8
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1539), int32(_a_F_array_recv_12))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L8
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
	F_errcode(m, int32(50462850))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L8
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_array_recv_14), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_array_recv_5), int32(1486), int32(_a_F_array_recv_12))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L8
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_sort(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = int32(0)
		v9 = F_array_sort_internal(m, v3, v7, v7, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v9)
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
	var v104 int64
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
	var v150 int64
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
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= int32(0) {
		v191 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L52
	}
L3:
	;
	m.G0 = v14 + int32(32)
	return v191
L4:
	;
	v20 = l0 + int32(16)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 < int32(2) {
		v191 = l0
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
	v91 = F_array_iterate(m, v85, v14+int32(24), v14+int32(23))
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
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)))
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
	v112 = F_array_iterate(m, v85, v14+int32(24), v14+int32(23))
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
	v137 = F_tuplesort_getdatum(m, v81, int32(1), v131, v14+int32(24), v14+int32(23), v131)
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
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)))
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
	v163 = F_tuplesort_getdatum(m, v81, int32(1), v157, v14+int32(24), v14+int32(23), v157)
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
	v183 = F_pg_detoast_datum(m, base.I32_wrap_i64(v180))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32)+v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v183+v185<<(uint(int32(2))%32))+16)) = v189
	v191 = v183
	goto L3
L52:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v213 = F_format_type_be(m, v38)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v213
	F_errmsg(m, int32(_a_F_array_sort_internal_0), v14+int32(16))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_array_sort_internal_1), int32(1960), int32(_a_F_array_sort_internal_2))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
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
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v233 = F_format_type_be(m, v38)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v233
	F_errmsg(m, int32(_a_F_array_sort_internal_3), v14)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_array_sort_internal_1), int32(1973), int32(_a_F_array_sort_internal_2))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
func F_array_to_tsvector(m *base.Module, l0 int32) int64 {
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
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
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int64
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v250 int32
	_ = v250
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	F_deconstruct_array_builtin(m, v18, int32(25), v15+int32(44), v15+int32(40), v15+int32(36))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if v31 <= int32(0) {
		v369 = v2
		v372 = v31
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L119
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L115
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L111
	}
L7:
	;
	v384 = v369 + v372<<(uint(int32(2))%32) + int32(8)
	v385 = F_palloc0(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L97
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v39 = v2
	goto L10
L9:
	;
	v86 = int32(1)
	if v31 != v86 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v35))))
	if v49 == int32(1) {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v34+v39<<(uint(int32(3))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v60 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) - int32(4)
	if v60 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v56) < base.Ui32(int32(_a_F_array_to_tsvector_0)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = v39 + int32(1)
	if v66 == v31 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L11
L17:
	;
	v39 = v66
	goto L10
L18:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v60
	F_errmsg(m, int32(_a_F_array_to_tsvector_1), v15)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_array_to_tsvector_2), int32(784), int32(_a_F_array_to_tsvector_3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_pg_qsort(m, v89, v31, int32(8), int32(1735))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v250 = v31
	goto L24
L24:
	;
	if v250 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if base.Ui32(int32(2)) <= base.Ui32(v94) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v100 = int32(0)
	v102 = v86
	goto L29
L27:
	;
	v237 = v94
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v237
	v250 = v237
	goto L24
L29:
	;
	v111 = int32(3)
	v113 = v98 + v102<<(uint(v111)%32)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v124 = int32(1)
	v125 = v123 & v124
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v98+v100<<(uint(v111)%32))))
	if v123 == v124 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v237 = v226 + int32(1)
	goto L28
L31:
	;
	v229 = v102 + int32(1)
	if v229 != v94 {
		v100 = v226
		v102 = v229
		goto L29
	} else {
		goto L81
	}
L32:
	;
	if v215 == int32(0) {
		v226 = v100
		goto L31
	} else {
		goto L79
	}
L33:
	;
	v154 = base.I32_wrap_i64(v126)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v156 = int32(1)
	v157 = v155 & v156
	if v155 == v156 {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	if v132 == int32(18) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v143 = int32(1)
	if v125 != 0 {
		v153 = int32(base.Ui32(v123)>>(uint(v143)%32)) - v143
		goto L33
	} else {
		goto L43
	}
L37:
	;
	v135 = int32(16)
	goto L39
L38:
	;
	v135 = int32(0)
	goto L39
L39:
	;
	if base.Ui32((v132-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v142 = int32(4)
	goto L42
L41:
	;
	v142 = v135
	goto L42
L42:
	;
	v153 = v142
	goto L33
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v153 = int32(base.Ui32(v147)>>(uint(int32(2))%32)) - int32(4)
	goto L33
L44:
	;
	if v153 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v163 == int32(18) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v174 = int32(1)
	if v157 != 0 {
		v184 = int32(base.Ui32(v155)>>(uint(v174)%32)) - v174
		goto L44
	} else {
		goto L54
	}
L48:
	;
	v166 = int32(16)
	goto L50
L49:
	;
	v166 = int32(0)
	goto L50
L50:
	;
	if base.Ui32((v163-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v173 = int32(4)
	goto L53
L52:
	;
	v173 = v166
	goto L53
L53:
	;
	v184 = v173
	goto L44
L54:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v184 = int32(base.Ui32(v178)>>(uint(int32(2))%32)) - int32(4)
	goto L44
L55:
	;
	v188 = int32(0)
	if v188 < v184 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if v184 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v191 = int32(-1)
	goto L60
L59:
	;
	v191 = v188
	goto L60
L60:
	;
	v215 = v191
	goto L32
L61:
	;
	v215 = base.B2i32(int32(0) < v153)
	goto L32
L62:
	;
	goto L63
L63:
	;
	if v125 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v215 = v213
	goto L32
L65:
	;
	v198 = int32(1)
	goto L67
L66:
	;
	v198 = int32(4)
	goto L67
L67:
	;
	if v157 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v202 = int32(1)
	goto L70
L69:
	;
	v202 = int32(4)
	goto L70
L70:
	;
	if base.Ui32(v153) < base.Ui32(v184) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v205 = v153
	goto L73
L72:
	;
	v205 = v184
	goto L73
L73:
	;
	v206 = F_memcmp(m, v122+v198, v154+v202, v205)
	mBase = m.M
	if v206 != 0 {
		v213 = v206
		goto L64
	} else {
		goto L74
	}
L74:
	;
	if v153 == v184 {
		v213 = int32(0)
		goto L64
	} else {
		goto L75
	}
L75:
	;
	if v153 < v184 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v212 = int32(-1)
	goto L78
L77:
	;
	v212 = int32(1)
	goto L78
L78:
	;
	v213 = v212
	goto L64
L79:
	;
	v219 = v100 + int32(1)
	if v219 == v102 {
		v226 = v102
		goto L31
	} else {
		goto L80
	}
L80:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
	*(*int64)(unsafe.Add(mBase, uint32(v98+v219<<(uint(int32(3))%32)))) = v224
	v226 = v219
	goto L31
L81:
	;
	goto L30
L82:
	;
	v369 = int32(0)
	v372 = v250
	goto L7
L83:
	;
	goto L84
L84:
	;
	v262 = v250 & int32(3)
	v263 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if base.Ui32(int32(4)) <= base.Ui32(v250) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if int32(_a_F_array_to_tsvector_4) <= v355 {
		goto L4
	} else {
		goto L96
	}
L86:
	;
	v272 = v263
	v274 = v263
	v282 = v2
	goto L89
L87:
	;
	v316 = v263
	v318 = v263
	goto L88
L88:
	;
	v328 = v316
	v330 = v318
	v334 = v263
	goto L93
L89:
	;
	v285 = v264 + v274<<(uint(int32(3))%32)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v288 = int32(2)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v285)+24))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v307 = v272 + int32(base.Ui32(v287)>>(uint(v288)%32)) + int32(base.Ui32(v292)>>(uint(v288)%32)) + int32(base.Ui32(v297)>>(uint(v288)%32)) + int32(base.Ui32(v302)>>(uint(v288)%32)) - int32(16)
	v308 = int32(4)
	v309 = v274 + v308
	v311 = v282 + v308
	if v311 != v250&int32(2147483644) {
		v272 = v307
		v274 = v309
		v282 = v311
		goto L89
	} else {
		goto L91
	}
L90:
	;
	if v262 == int32(0) {
		v355 = v307
		goto L85
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	v316 = v307
	v318 = v309
	goto L88
L93:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v264+v330<<(uint(int32(3))%32))))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v348 = v328 + int32(base.Ui32(v343)>>(uint(int32(2))%32)) - int32(4)
	v349 = int32(1)
	v352 = v334 + v349
	if v352 != v262 {
		v328 = v348
		v330 = v330 + v349
		v334 = v352
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v355 = v348
	goto L85
L95:
	;
	goto L94
L96:
	;
	v369 = v355
	v372 = v250
	goto L7
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v384 << (uint(int32(2)) % 32)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v385)+4)) = v390
	if int32(0) < v390 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v395 = v385 + int32(8)
	v401 = v395 + v390<<(uint(int32(2))%32)
	v403 = int32(0)
	goto L101
L99:
	;
	goto L100
L100:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v458 != v18 {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v412+v403<<(uint(int32(3))%32))))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v421 = int32(base.Ui32(v417)>>(uint(int32(2))%32)) - int32(4)
	if v421 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L100
L103:
	;
	base.MemoryCopy(m, v401, v416+int32(4), v421)
	goto L105
L104:
	;
	goto L105
L105:
	;
	v425 = int32(2)
	v428 = int32(1)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v395+v403<<(uint(v425)%32)))) = v421<<(uint(v428)%32)&int32(4094) | (v401-(v395+v432<<(uint(v425)%32)))<<(uint(int32(12))%32)
	v443 = v403 + v428
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if v443 < v444 {
		v401 = v401 + v421
		v403 = v443
		goto L101
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	F_pfree(m, v18)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	m.G0 = v15 + int32(48)
	return base.I64_extend_i32_u(v385)
L110:
	;
	goto L109
L111:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_array_to_tsvector_5), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_array_to_tsvector_2), int32(772), int32(_a_F_array_to_tsvector_3))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(369098882))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(_a_F_array_to_tsvector_6), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_array_to_tsvector_2), int32(778), int32(_a_F_array_to_tsvector_3))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(_a_F_array_to_tsvector_7)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v355
	F_errmsg(m, int32(_a_F_array_to_tsvector_8), v15+int32(16))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_array_to_tsvector_2), int32(802), int32(_a_F_array_to_tsvector_3))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	switch l3 - int32(99) {
	case 0:
		v39 = int32(1)
		goto L1
	case 1:
		goto L4
	default:
		goto L3
	case 6:
		goto L5
	case 16:
		goto L2
	}
L1:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = l0 + int32(16)
	v44 = F_ArrayGetNItemsSafe(m, v41, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L10
	}
L2:
	;
	v39 = int32(2)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v39 = int32(8)
	goto L1
L5:
	;
	v39 = int32(4)
	goto L1
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l3
	F_errmsg_internal(m, int32(_a_F_deconstruct_array_0), v18)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_deconstruct_array_1), int32(322), int32(_a_F_deconstruct_array_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v46 = F_palloc_mul(m, int32(8), v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v46
	if l5 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = F_palloc0_mul(m, int32(1), v44)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v53 = int32(0)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v44
	if int32(0) < v44 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v50
	v53 = v50
	goto L14
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v59 = v57 << (uint(int32(3)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	m.G0 = v18 + int32(32)
	return
L19:
	;
	v63 = v43 + v59
	goto L21
L20:
	;
	v63 = int32(0)
	goto L21
L21:
	;
	if v62 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = v62
	goto L24
L23:
	;
	v72 = (v59 + int32(23)) & int32(-8)
	goto L24
L24:
	;
	v76 = int32(1)
	v80 = v76
	v83 = l0 + v72
	v85 = v63
	v87 = int32(0)
	goto L25
L25:
	;
	if v85 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L18
L27:
	;
	v190 = int32(1)
	v192 = v80 << (uint(v190) % 32)
	v194 = base.B2i32(v192 == int32(256))
	if v192 == int32(256) {
		goto L66
	} else {
		goto L67
	}
L28:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v80&v97 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46+v87<<(uint(int32(3))%32)))) = int64(0)
	if v53 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v53))) = uint8(v105)
	v188 = v83
	goto L27
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_deconstruct_array_3), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_deconstruct_array_4), int32(3677), int32(_a_F_deconstruct_array_5))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46+v87<<(uint(int32(3))%32)))) = v146
	if int32(0) < l1 {
		v185 = l1 + v83
		goto L51
	} else {
		goto L52
	}
L39:
	;
	if base.I32_popcnt(l1) != v76 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v146 = base.I64_extend_i32_u(v83)
	goto L38
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L48
	}
L43:
	;
	switch base.I32_ctz(l1) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	default:
		goto L42
	}
L44:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v146 = v129
	goto L38
L45:
	;
	v128 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83))))
	v146 = v128
	goto L38
L46:
	;
	v127 = int64(*(*int16)(unsafe.Add(mBase, uint32(v83))))
	v146 = v127
	goto L38
L47:
	;
	v126 = int64(*(*int8)(unsafe.Add(mBase, uint32(v83))))
	v146 = v126
	goto L38
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_deconstruct_array_6), v18+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_deconstruct_array_1), int32(123), int32(_a_F_deconstruct_array_7))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v188 = (v185 + (v39 - int32(1))) & (int32(0) - v39)
	goto L27
L52:
	;
	if l1 == int32(-1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v153 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v180 = F_strlen(m, v83)
	mBase = m.M
	v185 = v180 + v83 + int32(1)
	goto L51
L56:
	;
	v157 = int32(18)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v159 == v157 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v171 = int32(1)
	if v153&v171 != 0 {
		v185 = v83 + int32(base.Ui32(v153)>>(uint(v171)%32))
		goto L51
	} else {
		goto L65
	}
L59:
	;
	v162 = v157
	goto L61
L60:
	;
	v162 = int32(2)
	goto L61
L61:
	;
	if base.Ui32((v159-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v169 = int32(6)
	goto L64
L63:
	;
	v169 = v162
	goto L64
L64:
	;
	v185 = v83 + v169
	goto L51
L65:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v185 = v83 + int32(base.Ui32(v176)>>(uint(int32(2))%32))
	goto L51
L66:
	;
	v195 = v190
	goto L68
L67:
	;
	v195 = v192
	goto L68
L68:
	;
	if v85 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v196 = v195
	goto L71
L70:
	;
	v196 = v80
	goto L71
L71:
	;
	if v85 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v199 = v85 + v194
	goto L74
L73:
	;
	v199 = int32(0)
	goto L74
L74:
	;
	v201 = v87 + int32(1)
	if v201 != v44 {
		v80 = v196
		v83 = v188
		v85 = v199
		v87 = v201
		goto L25
	} else {
		goto L75
	}
L75:
	;
	goto L26
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(99)
	v15 = int32(1)
	switch l1 - int32(18) {
	case 0:
		v50 = v15
		v51 = v14
		v52 = v15
		F_deconstruct_array(m, l0, v50, v52, v51, l2, l3, l4)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
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
				F_errfinish(m, int32(_a_F_deconstruct_array_builtin_1), int32(3764), int32(_a_F_deconstruct_array_builtin_2))
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
		v50 = int32(2)
		v51 = int32(115)
		v52 = v15
		F_deconstruct_array(m, l0, v50, v52, v51, l2, l3, l4)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 5, 8:
		v50 = int32(4)
		v51 = int32(105)
		v52 = v15
		F_deconstruct_array(m, l0, v50, v52, v51, l2, l3, l4)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 7:
		v50 = int32(-1)
		v51 = int32(105)
		v52 = int32(0)
		F_deconstruct_array(m, l0, v50, v52, v51, l2, l3, l4)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 9:
		v50 = int32(6)
		v51 = int32(115)
		v52 = int32(0)
		F_deconstruct_array(m, l0, v50, v52, v51, l2, l3, l4)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	default:
		if l1 == int32(701) {
			v50 = int32(8)
			v51 = int32(100)
			v52 = v15
			F_deconstruct_array(m, l0, v50, v52, v51, l2, l3, l4)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
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
						F_errfinish(m, int32(_a_F_deconstruct_array_builtin_1), int32(3764), int32(_a_F_deconstruct_array_builtin_2))
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
				v50 = int32(-2)
				v51 = v14
				v52 = int32(0)
				F_deconstruct_array(m, l0, v50, v52, v51, l2, l3, l4)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	}
}
func F_expand_array(m *base.Module, l0 int64, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v18 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_expand_array_0), int32(0), int32(1024), int32(_a_F_expand_array_1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v23 = F_MemoryContextAlloc(m, v18, int32(80))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v26 = int32(513)
			*(*uint16)(unsafe.Add(mBase, uint32(v23)+18)) = uint16(v26)
			v28 = int32(769)
			*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)) = uint16(v28)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(_a_F_expand_array_2)
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v23)+14)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(689375833)
			v38 = base.I32_wrap_i64(l0)
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
			if v39 != int32(1) {
				v122 = l2
				v124 = int32(_a_F_expand_array_3)
				v125 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v18
				v128 = F_pg_detoast_datum_copy(m, v38)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v125
					v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
					v134 = v128 + int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v134
					*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v132
					v137 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v134 + v137<<(uint(int32(2))%32)
					v142 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v142
					if v122 != 0 {
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
						if v142 == v144 {
							v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)) = uint16(v146)
							v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
							*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)) = uint8(v148)
							v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
							*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)) = uint8(v150)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
							v178 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
							*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
							*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
							if v183 != 0 {
								v191 = v183
							} else {
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
								v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
							v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
							v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
							*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
							m.G0 = v12 + int32(48)
							return base.I64_extend_i32_u(v23 + int32(12))
						} else {
							F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
							mBase = m.M
							v159 = m.ExcPending
							if v159 != 0 {
								return int64(0)
							} else {
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v122))) = v160
								v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)) = uint16(v162)
								v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)))
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)) = uint8(v164)
								v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)))
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)) = uint8(v166)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
								v178 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
								*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
								*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
								if v183 != 0 {
									v191 = v183
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
									v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
								v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
								m.G0 = v12 + int32(48)
								return base.I64_extend_i32_u(v23 + int32(12))
							}
						}
					} else {
						F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
							v178 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
							*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
							*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
							if v183 != 0 {
								v191 = v183
							} else {
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
								v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
							v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
							v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
							*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
							m.G0 = v12 + int32(48)
							return base.I64_extend_i32_u(v23 + int32(12))
						}
					}
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
				if v42&int32(254) != int32(2) {
					v122 = l2
					v124 = int32(_a_F_expand_array_3)
					v125 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v18
					v128 = F_pg_detoast_datum_copy(m, v38)
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v125
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
						v134 = v128 + int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v132
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v134 + v137<<(uint(int32(2))%32)
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v142
						if v122 != 0 {
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
							if v142 == v144 {
								v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)) = uint16(v146)
								v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
								*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)) = uint8(v148)
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
								*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)) = uint8(v150)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
								v178 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
								*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
								*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
								if v183 != 0 {
									v191 = v183
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
									v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
								v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
								m.G0 = v12 + int32(48)
								return base.I64_extend_i32_u(v23 + int32(12))
							} else {
								F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int64(0)
								} else {
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v122))) = v160
									v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)))
									*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)) = uint16(v162)
									v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)))
									*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)) = uint8(v164)
									v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)))
									*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)) = uint8(v166)
									*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
									v178 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
									*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
									*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
									if v183 != 0 {
										v191 = v183
									} else {
										v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
										v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
									v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
									m.G0 = v12 + int32(48)
									return base.I64_extend_i32_u(v23 + int32(12))
								}
							}
						} else {
							F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
								v178 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
								*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
								*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
								if v183 != 0 {
									v191 = v183
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
									v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
								v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
								m.G0 = v12 + int32(48)
								return base.I64_extend_i32_u(v23 + int32(12))
							}
						}
					}
				} else {
					if l2 != 0 {
						v47 = l2
					} else {
						v47 = v12
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = v50
					v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+44)))
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)) = uint16(v52)
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+46)))
					*(*uint8)(unsafe.Add(mBase, uint32(v47)+6)) = uint8(v54)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+47)))
					*(*uint8)(unsafe.Add(mBase, uint32(v47)+7)) = uint8(v56)
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+46)))
					if v58 != int32(1) {
						v122 = v47
						v124 = int32(_a_F_expand_array_3)
						v125 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v18
						v128 = F_pg_detoast_datum_copy(m, v38)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v125
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
							v134 = v128 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v132
							v137 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v134 + v137<<(uint(int32(2))%32)
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v142
							if v122 != 0 {
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
								if v142 == v144 {
									v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)) = uint16(v146)
									v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
									*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)) = uint8(v148)
									v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
									*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)) = uint8(v150)
									*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
									v178 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
									*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
									*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
									if v183 != 0 {
										v191 = v183
									} else {
										v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
										v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
									v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
									m.G0 = v12 + int32(48)
									return base.I64_extend_i32_u(v23 + int32(12))
								} else {
									F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int64(0)
									} else {
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v122))) = v160
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)))
										*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)) = uint16(v162)
										v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)))
										*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)) = uint8(v164)
										v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)))
										*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)) = uint8(v166)
										*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
										v178 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
										*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
										*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
										v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
										if v183 != 0 {
											v191 = v183
										} else {
											v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
											v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
										v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
										v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
										m.G0 = v12 + int32(48)
										return base.I64_extend_i32_u(v23 + int32(12))
									}
								}
							} else {
								F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
									v178 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
									*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
									*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
									if v183 != 0 {
										v191 = v183
									} else {
										v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
										v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
									v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
									m.G0 = v12 + int32(48)
									return base.I64_extend_i32_u(v23 + int32(12))
								}
							}
						}
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
						if v61 == int32(0) {
							v122 = v47
							v124 = int32(_a_F_expand_array_3)
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_expand_array[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v18
							v128 = F_pg_detoast_datum_copy(m, v38)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_expand_array[0])) = v125
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
								v134 = v128 + int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v132
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v134 + v137<<(uint(int32(2))%32)
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v142
								if v122 != 0 {
									v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
									if v142 == v144 {
										v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)) = uint16(v146)
										v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
										*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)) = uint8(v148)
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
										*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)) = uint8(v150)
										*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
										v178 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
										*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
										*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
										v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
										if v183 != 0 {
											v191 = v183
										} else {
											v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
											v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
										v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
										v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
										m.G0 = v12 + int32(48)
										return base.I64_extend_i32_u(v23 + int32(12))
									} else {
										F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int64(0)
										} else {
											v160 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v122))) = v160
											v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)))
											*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)) = uint16(v162)
											v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)))
											*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)) = uint8(v164)
											v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)))
											*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)) = uint8(v166)
											*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
											v178 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
											*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
											*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
											v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
											if v183 != 0 {
												v191 = v183
											} else {
												v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
												v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
											v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
											v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
											*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
											m.G0 = v12 + int32(48)
											return base.I64_extend_i32_u(v23 + int32(12))
										}
									}
								} else {
									F_get_typlenbyvalalign(m, v142, v23+int32(44), v23+int32(46), v23+int32(47))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
										v178 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v178
										*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v178
										*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v128
										v183 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
										if v183 != 0 {
											v191 = v183
										} else {
											v184 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
											v191 = (v184<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v191 + v128
										v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
										v204 = v128 + int32(base.Ui32(v194)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
										m.G0 = v12 + int32(48)
										return base.I64_extend_i32_u(v23 + int32(12))
									}
								}
							}
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+56))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v65
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
							v70 = F_MemoryContextAlloc(m, v67, v65<<(uint(int32(3))%32))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v70
								v74 = v65 << (uint(int32(2)) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v70 + v74
								v77 = int32(0)
								v78 = base.B2i32(v74 == v77)
								if v78 == v77 {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v49)+32))
									base.MemoryCopy(m, v70, v81, v74)
								} else {
								}
								if v78 == int32(0) {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v49)+36))
									base.MemoryCopy(m, v85, v86, v74)
								} else {
								}
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v49)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v88
								v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)) = uint16(v90)
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+46)))
								*(*uint8)(unsafe.Add(mBase, uint32(v23)+46)) = uint8(v92)
								v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+47)))
								*(*uint8)(unsafe.Add(mBase, uint32(v23)+47)) = uint8(v94)
								v97 = v64 << (uint(int32(3)) % 32)
								v98 = F_MemoryContextAlloc(m, v67, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v98
									if v97 != 0 {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
										base.MemoryCopy(m, v98, v101, v97)
									} else {
									}
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v49)+52))
									if v103 != 0 {
										v104 = F_MemoryContextAlloc(m, v67, v64)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v104
											if v64 == int32(0) {
											} else {
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v49)+52))
												base.MemoryCopy(m, v104, v109, v64)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v64
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v115
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
											*(*int64)(unsafe.Add(mBase, uint32(v23)+68)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v117
											v204 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
											m.G0 = v12 + int32(48)
											return base.I64_extend_i32_u(v23 + int32(12))
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v64
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v115
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v23)+68)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v117
										v204 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v204
										m.G0 = v12 + int32(48)
										return base.I64_extend_i32_u(v23 + int32(12))
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
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int64 {
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
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResult[0])) = v17
		if v15 == int32(1) {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_MemoryContextDelete(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				m.G0 = v9 + int32(16)
				return base.I64_extend_i32_u(v32)
			}
		} else {
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(v32)
		}
	}
}
func F_makeArrayTypeName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v14 = F_makeObjectName(m, int32(_a_F_makeArrayTypeName_0), l0, v3)
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
	v19 = base.I64_extend_i32_u(l1)
	v20 = int64(0)
	v22 = F_SearchSysCacheExists(m, int32(81), base.I64_extend_i32_u(v14), v19, v20, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = v14
	v28 = v3
	goto L7
L5:
	;
	v53 = v14
	goto L6
L6:
	;
	m.G0 = v9 + int32(80)
	return v53
L7:
	;
	F_pfree(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v53 = v43
	goto L6
L9:
	;
	v33 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v33
	v36 = v9 + int32(16)
	v39 = F_pg_snprintf(m, v36, int32(64), int32(_a_F_makeArrayTypeName_1), v9)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v43 = F_makeObjectName(m, int32(_a_F_makeArrayTypeName_0), l0, v36)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v46 = int64(0)
	v48 = F_SearchSysCacheExists(m, int32(81), base.I64_extend_i32_u(v43), v19, v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v48 != 0 {
		v27 = v43
		v28 = v33
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
