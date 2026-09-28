package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbValueToJsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v9) < base.Ui32(int32(4)) {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
		v16 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v16
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_JsonbValueToJsonb[0]))
		v23 = F_MemoryContextAlloc(m, v21, int32(48))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v23)+40)) = uint16(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v23
			v32 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v32
			*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)) = uint8(v32)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(16)
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_JsonbValueToJsonb[0]))
			v43 = F_MemoryContextAlloc(m, v41, int32(32))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v43
				v47 = v7 + int32(8)
				F_pushJsonbValue(m, v47, int32(3), l0)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_pushJsonbValueScalar(m, v47, int32(5), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
						v57 = v7 + int32(32)
						F_initStringInfo(m, v57)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_enlargeStringInfo(m, v57, int32(4))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
								v65 = v63 + int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v65
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
								v69 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v67+v65))) = uint8(v69)
								F_convertJsonbValue(m, v57, v7+int32(28), v55, v69)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v76))) = v77 << (uint(int32(2)) % 32)
									v125 = v76
									m.G0 = v7 + int32(48)
									return v125
								}
							}
						}
					}
				}
			}
		}
	} else {
		switch v9 - int32(16) {
		case 0, 1:
			v82 = v7 + int32(8)
			F_initStringInfo(m, v82)
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				F_enlargeStringInfo(m, v82, int32(4))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v90 = v88 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v90
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v94 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v92+v90))) = uint8(v94)
					F_convertJsonbValue(m, v82, v7+int32(32), l0, v94)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v101))) = v102 << (uint(int32(2)) % 32)
						v125 = v101
						m.G0 = v7 + int32(48)
						return v125
					}
				}
			}
		default:
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v109 = F_palloc(m, v106+int32(4))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v109))) = v111<<(uint(int32(2))%32) + int32(16)
				v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v117 == int32(0) {
					v125 = v109
				} else {
					v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v109+int32(4), v122, v117)
					v125 = v109
				}
				m.G0 = v7 + int32(48)
				return v125
			}
		case 16:
			*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
			v16 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v16
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_JsonbValueToJsonb[0]))
			v23 = F_MemoryContextAlloc(m, v21, int32(48))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v23)+40)) = uint16(v27)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v23
				v32 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v32
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(16)
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_JsonbValueToJsonb[0]))
				v43 = F_MemoryContextAlloc(m, v41, int32(32))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v43
					v47 = v7 + int32(8)
					F_pushJsonbValue(m, v47, int32(3), l0)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_pushJsonbValueScalar(m, v47, int32(5), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
							v57 = v7 + int32(32)
							F_initStringInfo(m, v57)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_enlargeStringInfo(m, v57, int32(4))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
									v65 = v63 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v65
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v67+v65))) = uint8(v69)
									F_convertJsonbValue(m, v57, v7+int32(28), v55, v69)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
										*(*int32)(unsafe.Add(mBase, uint32(v76))) = v77 << (uint(int32(2)) % 32)
										v125 = v76
										m.G0 = v7 + int32(48)
										return v125
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
func F_jsonb_array_elements(m *base.Module, l0 int32) int64 {
	var v6 int32
	_ = v6
	F_elements_worker_jsonb(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_jsonb_build_array_noargs(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14321(m, l0, int32(5), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_jsonb_build_array_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int64
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v6
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v16
	F_pushJsonbValue(m, v12+int32(8), int32(4), v6)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if int32(0) < l0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v38 = v6
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_pushJsonbValue(m, v12+int32(8), int32(5), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v38))))
	if v43&int32(1) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v46 = l4
	goto L10
L9:
	;
	v46 = int32(0)
	goto L10
L10:
	;
	if v46 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l1+v38<<(uint(int32(3))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3+v38<<(uint(int32(2))%32))))
	F_add_jsonb(m, v52, (l4^int32(1))&v43, v12+int32(8), v59, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v64 = v38 + int32(1)
	if v64 != l0 {
		v38 = v64
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L7
L16:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v82 = F_JsonbValueToJsonb(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v12 + int32(32)
	return base.I64_extend_i32_u(v82)
}
func F_jsonb_build_object_noargs(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14321(m, l0, int32(7), int32(6))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_jsonb_contained(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if (v17^v18)&int32(536870912) == int32(0) {
				v26 = F_JsonbIteratorInit(m, v15+int32(4))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v26
					v31 = F_JsonbIteratorInit(m, v10+int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v31
						v38 = F_JsonbDeepContains(m, v7+int32(12), v7+int32(8))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v41 = base.I64_extend_i32_u(v38)
							m.G0 = v7 + int32(16)
							return v41
						}
					}
				}
			} else {
				v41 = int64(0)
				m.G0 = v7 + int32(16)
				return v41
			}
		}
	}
}
func F_jsonb_delete_array(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
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
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	v13 = m.G0
	v15 = v13 - int32(80)
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(0)
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v27
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v31 < int32(2) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L78
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v34&int32(268435456) != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L74
	}
L8:
	;
	if v34&int32(268435455) == int32(0) {
		v283 = v18
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v15 + int32(80)
	return base.I64_extend_i32_u(v283)
L10:
	;
	F_deconstruct_array_builtin(m, v23, int32(25), v15+int32(76), v15+int32(72), v15+int32(68))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if v50 == int32(0) {
		v283 = v18
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v55 = F_JsonbIteratorInit(m, v18+int32(4))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v55
	v63 = F_JsonbIteratorNext(m, v15+int32(44), v15+int32(8), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = v63
	goto L18
L16:
	;
	v278 = int32(0)
	goto L17
L17:
	;
	v279 = F_JsonbValueToJsonb(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L73
	}
L18:
	;
	v78 = base.B2i32(v65 != int32(1))
	if v78&base.B2i32(v65 != int32(3)) != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v278 = v264
	goto L17
L20:
	;
	v262 = F_JsonbIteratorNext(m, v15+int32(44), v15+int32(8), int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L71
	}
L21:
	;
	if v65 != int32(1) {
		goto L20
	} else {
		goto L69
	}
L22:
	;
	if base.Ui32(v65) < base.Ui32(int32(4)) {
		goto L65
	} else {
		goto L66
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v82 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if v85 <= int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v95 = int32(0)
	goto L26
L26:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v92))))
	if v106 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L22
L28:
	;
	v214 = v95 + int32(1)
	if v214 != v85 {
		v95 = v214
		goto L26
	} else {
		goto L64
	}
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91+v95<<(uint(int32(3))%32))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v112 = int32(1)
	v113 = v111 & v112
	if v111 == v112 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v140 != v90 {
		goto L28
	} else {
		goto L41
	}
L31:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v119 == int32(18) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v130 = int32(1)
	if v113 != 0 {
		v140 = int32(base.Ui32(v111)>>(uint(v130)%32)) - v130
		goto L30
	} else {
		goto L40
	}
L34:
	;
	v122 = int32(16)
	goto L36
L35:
	;
	v122 = int32(0)
	goto L36
L36:
	;
	if base.Ui32((v119-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v129 = int32(4)
	goto L39
L38:
	;
	v129 = v122
	goto L39
L39:
	;
	v140 = v129
	goto L30
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v140 = int32(base.Ui32(v134)>>(uint(int32(2))%32)) - int32(4)
	goto L30
L41:
	;
	if v113 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v144 = int32(1)
	goto L44
L43:
	;
	v144 = int32(4)
	goto L44
L44:
	;
	v145 = v110 + v144
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if v207 == int32(0) {
		goto L21
	} else {
		goto L63
	}
L46:
	;
	v207 = int32(0)
	goto L45
L47:
	;
	v181 = v176
	v182 = v177
	v183 = v178
	goto L57
L48:
	;
	if (v145|v89)&int32(3) != 0 {
		v176 = v145
		v177 = v89
		v178 = v90
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v169 = v145
	v170 = v89
	v171 = v90
	goto L50
L50:
	;
	if v171 == int32(0) {
		goto L46
	} else {
		goto L56
	}
L51:
	;
	v153 = v145
	v154 = v89
	v155 = v90
	goto L52
L52:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v158 != v159 {
		v176 = v153
		v177 = v154
		v178 = v155
		goto L47
	} else {
		goto L54
	}
L53:
	;
	v169 = v164
	v170 = v162
	v171 = v166
	goto L50
L54:
	;
	v161 = int32(4)
	v162 = v154 + v161
	v164 = v153 + v161
	v166 = v155 - v161
	if base.Ui32(int32(3)) < base.Ui32(v166) {
		v153 = v164
		v154 = v162
		v155 = v166
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v176 = v169
	v177 = v170
	v178 = v171
	goto L47
L57:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v186 == v187 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v207 = v186 - v187
	goto L45
L59:
	;
	v189 = int32(1)
	v194 = v183 - v189
	if v194 != 0 {
		v181 = v181 + v189
		v182 = v182 + v189
		v183 = v194
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	goto L46
L63:
	;
	goto L28
L64:
	;
	goto L27
L65:
	;
	v235 = v15 + int32(8)
	goto L67
L66:
	;
	v235 = int32(0)
	goto L67
L67:
	;
	F_pushJsonbValue(m, v15+int32(48), v65, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	goto L20
L69:
	;
	v243 = F_JsonbIteratorNext(m, v15+int32(44), v15+int32(8), int32(1))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L20
L71:
	;
	if v262 != 0 {
		v65 = v262
		goto L18
	} else {
		goto L72
	}
L72:
	;
	goto L19
L73:
	;
	v283 = v279
	goto L9
L74:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_jsonb_delete_array_0), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_jsonb_delete_array_1), int32(_a_F_jsonb_delete_array_2), int32(_a_F_jsonb_delete_array_3))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_jsonb_delete_array_4), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_jsonb_delete_array_1), int32(_a_F_jsonb_delete_array_5), int32(_a_F_jsonb_delete_array_3))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_exec_setup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v14 = F_palloc0(m, v9*int32(12)+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v16)
	v19 = v14 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v19 + v9<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v14
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v26 == v16 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = int32(1527)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(1528)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = int32(1529)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1530)
	return
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = int32(0)
	goto L6
L6:
	;
	v40 = v37 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41)))
	v44 = F_exprType(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46+v40))) = v44
	v50 = v37 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v50 < v51 {
		v37 = v50
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_jsonb_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v6)
	v19 = l0 + int32(4)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = int32(base.Ui32(v21&int32(536870912)) >> (uint(int32(29)) % 32))
	if v25 != 0 {
		v40 = v6
		v41 = v6
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_u(v234)
L2:
	;
	if l4 == int32(0) {
		v234 = l0
		goto L1
	} else {
		goto L71
	}
L3:
	;
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v221)
	v234 = int32(0)
	goto L1
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L11
	} else {
		goto L68
	}
L5:
	;
	if l4 != 0 {
		goto L60
	} else {
		goto L61
	}
L6:
	;
	v53 = int32(0)
	v59 = v48
	v60 = v19
	v61 = v25
	goto L15
L7:
	;
	v42 = int32(0)
	v45 = base.B2i32(l2 <= v42)
	if base.B2i32(v41 == v42)&v45 != 0 {
		goto L2
	} else {
		goto L13
	}
L8:
	;
	if v21&int32(1342177280) == int32(1073741824) {
		v40 = int32(1)
		v41 = int32(0)
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = int32(0)
	if v32 < l2 {
		v48 = v32
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v36 = F_getIthJsonbValueFromContainer(m, v19, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int64(0)
L12:
	;
	v40 = v32
	v41 = v36
	goto L7
L13:
	;
	if l2 <= v42 {
		v184 = v41
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v48 = v40
	goto L6
L15:
	;
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v154 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1+v53<<(uint(int32(3))%32))))
	v68 = F_pg_detoast_datum_packed(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v59 == int32(0) {
		goto L3
	} else {
		goto L40
	}
L21:
	;
	v70 = int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v74 = v72 & v70
	if v74 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = v70
	goto L24
L23:
	;
	v75 = int32(4)
	goto L24
L24:
	;
	v76 = v68 + v75
	if v72 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	if v82 == int32(18) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v74 != 0 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v85 = int32(16)
	goto L30
L29:
	;
	v85 = int32(0)
	goto L30
L30:
	;
	if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v92 = int32(4)
	goto L33
L32:
	;
	v92 = v85
	goto L33
L33:
	;
	v94 = F_getKeyJsonValueFromContainer(m, v60, v76, v92, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v154 = v94
	goto L17
L35:
	;
	v96 = int32(1)
	v101 = F_getKeyJsonValueFromContainer(m, v60, v76, int32(base.Ui32(v72)>>(uint(v96)%32))-v96, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v109 = F_getKeyJsonValueFromContainer(m, v60, v76, int32(base.Ui32(v103)>>(uint(int32(2))%32))-int32(4), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L39
	}
L38:
	;
	v154 = v101
	goto L17
L39:
	;
	v154 = v109
	goto L17
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+v53<<(uint(int32(3))%32))))
	v117 = F_text_to_cstring(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_jsonb_get_element[0])) = int32(0)
	v125 = F_strtol(m, v117, v14+int32(12), int32(10))
	mBase = m.M
	goto L42
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v117 == v126 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v128 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_get_element[0]))
	if v130 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	if v125 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v133&int32(1073741824) == int32(0) {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	v147 = v125
	goto L48
L48:
	;
	v148 = F_getIthJsonbValueFromContainer(m, v60, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L11
	} else {
		goto L52
	}
L49:
	;
	if v125 == int32(-2147483648) {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v141 = v133 & int32(268435455)
	if base.Ui32(v141) < base.Ui32(int32(0)-v125) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v147 = v125 + v141
	goto L48
L52:
	;
	v154 = v148
	goto L17
L53:
	;
	goto L3
L54:
	;
	goto L55
L55:
	;
	if v53 == l2-int32(1) {
		v184 = v154
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v158 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v160 == int32(18) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v174 = int32(base.Ui32(v164&int32(1073741824)) >> (uint(int32(30)) % 32))
	v175 = v163
	v176 = int32(base.Ui32(v164&int32(536870912)) >> (uint(int32(29)) % 32))
	goto L59
L58:
	;
	v174 = v158
	v175 = v60
	v176 = v158
	goto L59
L59:
	;
	v53 = v53 + int32(1)
	v59 = v174
	v60 = v175
	v61 = v176
	goto L15
L60:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v190 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v195 = F_JsonbValueToJsonb(m, v184)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L11
	} else {
		goto L67
	}
L63:
	;
	goto L3
L64:
	;
	goto L65
L65:
	;
	v193 = F_JsonbValueAsText(m, v184)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	v234 = v193
	goto L1
L67:
	;
	v234 = v195
	goto L1
L68:
	;
	F_errmsg_internal(m, int32(_a_F_jsonb_get_element_0), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_jsonb_get_element_1), int32(1617), int32(_a_F_jsonb_get_element_2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v230 = F_JsonbToCString(m, int32(0), v19, int32(base.Ui32(v227)>>(uint(int32(2))%32)))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v232 = F_cstring_to_text(m, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v234 = v232
	goto L1
}
func F_jsonb_in_array_end(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pushJsonbValue(m, l0, int32(5), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_int4(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14322(m, l0, int32(1458), int32(_a_F_jsonb_int4_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_jsonb_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v17 <= int32(0)))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(v17 <= int32(0)))
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v17 <= int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v17 <= int32(0)))
					}
				}
			}
		}
	}
}
func F_jsonb_path_match(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_match_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_subscript_assign(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
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
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	if v16 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
		if v38 == int32(1) {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v42 == int32(1) {
				v45 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v45)
				v48 = int32(16)
			} else {
				v48 = int32(17)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v48
			v52 = F_JsonbValueToJsonb(m, v12)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v55 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
				v61 = v52
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				v65 = v10 + int32(-32)
				v66 = m.G0
				v68 = v66 - int32(32)
				m.G0 = v68
				*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(0)
				v72 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v72
				v77 = F_palloc0_mul(m, int32(1), v63)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					if v79 != int32(16) {
					} else {
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+16)))
						if v82 != int32(1) {
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
							v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							*(*int64)(unsafe.Add(mBase, uint32(v65))) = v86
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v88
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v90
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v92
						}
					}
					v97 = F_JsonbIteratorInit(m, v61+int32(4))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v97
						F_setPath(m, v68+int32(4), v62, v77, v63, v68+int32(8), int32(0), v65, int32(97))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							F_pfree(m, v77)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
								v111 = F_JsonbValueToJsonb(m, v110)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									m.G0 = v68 + int32(32)
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v116))) = base.I64_extend_i32_u(v111)
									m.G0 = v12 - int32(-64)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			v59 = F_pg_detoast_datum(m, v58)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				v61 = v59
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				v65 = v10 + int32(-32)
				v66 = m.G0
				v68 = v66 - int32(32)
				m.G0 = v68
				*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(0)
				v72 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v72
				v77 = F_palloc0_mul(m, int32(1), v63)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					if v79 != int32(16) {
					} else {
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+16)))
						if v82 != int32(1) {
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
							v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							*(*int64)(unsafe.Add(mBase, uint32(v65))) = v86
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v88
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v90
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v92
						}
					}
					v97 = F_JsonbIteratorInit(m, v61+int32(4))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v97
						F_setPath(m, v68+int32(4), v62, v77, v63, v68+int32(8), int32(0), v65, int32(97))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							F_pfree(m, v77)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
								v111 = F_JsonbValueToJsonb(m, v110)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									m.G0 = v68 + int32(32)
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v116))) = base.I64_extend_i32_u(v111)
									m.G0 = v12 - int32(-64)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v25 = v10 + int32(-32)
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(18)
			v28 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v22 + v28
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - v28
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
			if v38 == int32(1) {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v42 == int32(1) {
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v45)
					v48 = int32(16)
				} else {
					v48 = int32(17)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v48
				v52 = F_JsonbValueToJsonb(m, v12)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v55 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
					v61 = v52
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					v65 = v10 + int32(-32)
					v66 = m.G0
					v68 = v66 - int32(32)
					m.G0 = v68
					*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(0)
					v72 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v72
					v77 = F_palloc0_mul(m, int32(1), v63)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						if v79 != int32(16) {
						} else {
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+16)))
							if v82 != int32(1) {
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
								v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								*(*int64)(unsafe.Add(mBase, uint32(v65))) = v86
								v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v88
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v90
								v92 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v92
							}
						}
						v97 = F_JsonbIteratorInit(m, v61+int32(4))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v97
							F_setPath(m, v68+int32(4), v62, v77, v63, v68+int32(8), int32(0), v65, int32(97))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								F_pfree(m, v77)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
									v111 = F_JsonbValueToJsonb(m, v110)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										m.G0 = v68 + int32(32)
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int64)(unsafe.Add(mBase, uint32(v116))) = base.I64_extend_i32_u(v111)
										m.G0 = v12 - int32(-64)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				v59 = F_pg_detoast_datum(m, v58)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v61 = v59
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					v65 = v10 + int32(-32)
					v66 = m.G0
					v68 = v66 - int32(32)
					m.G0 = v68
					*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(0)
					v72 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v72
					v77 = F_palloc0_mul(m, int32(1), v63)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						if v79 != int32(16) {
						} else {
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+16)))
							if v82 != int32(1) {
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
								v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								*(*int64)(unsafe.Add(mBase, uint32(v65))) = v86
								v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v88
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v90
								v92 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v92
							}
						}
						v97 = F_JsonbIteratorInit(m, v61+int32(4))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v97
							F_setPath(m, v68+int32(4), v62, v77, v63, v68+int32(8), int32(0), v65, int32(97))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								F_pfree(m, v77)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
									v111 = F_JsonbValueToJsonb(m, v110)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										m.G0 = v68 + int32(32)
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int64)(unsafe.Add(mBase, uint32(v116))) = base.I64_extend_i32_u(v111)
										m.G0 = v12 - int32(-64)
										return
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
func F_jsonb_to_recordset(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = int32(0)
	F_populate_recordset_worker(m, l0, int32(_a_F_jsonb_to_recordset_0), v3, v3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
