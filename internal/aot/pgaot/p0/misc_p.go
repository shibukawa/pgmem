package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseAbortRecord(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int64
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	v4 = int32(0)
	base.MemoryFill(m, l2, v4, int32(264))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v10
	if v4 <= base.I32_extend8_s(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v15
	if v15&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v21
	v27 = l1 + int32(20)
	goto L5
L4:
	;
	v27 = l1 + int32(12)
	goto L5
L5:
	;
	if v15&int32(2) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v32 = v27 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v30
	v38 = v32 + v30<<(uint(int32(2))%32)
	goto L8
L7:
	;
	v38 = v27
	goto L8
L8:
	;
	if v15&int32(4) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v44 = v38 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v51 = v44 + v47*int32(12)
	goto L11
L10:
	;
	v51 = v38
	goto L11
L11:
	;
	if v15&int32(256) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v57 = int32(4)
	v58 = v51 + v57
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v56
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v65 = v58 + v61<<(uint(v57)%32)
	goto L14
L13:
	;
	v65 = v51
	goto L14
L14:
	;
	if v15&int32(16) == int32(0) {
		v206 = v15
		v207 = v65
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v206&int32(32) == int32(0) {
		goto L1
	} else {
		goto L49
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+44)) = v72
	v75 = v65 + int32(4)
	if v15&int32(128) == int32(0) {
		v206 = v15
		v207 = v75
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v81 = l2 + int32(48)
	goto L21
L18:
	;
	v201 = F_strlen(m, v75)
	mBase = m.M
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v206 = v205
	v207 = v201 + v75 + int32(1)
	goto L15
L19:
	;
	v198 = F_strlen(m, v187)
	mBase = m.M
	goto L18
L21:
	;
	goto L22
L22:
	;
	v88 = int32(199)
	if (v81^v75)&int32(3) != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v191)
	goto L19
L24:
	;
	v172 = v167
	v173 = v168
	v174 = v169
	goto L45
L25:
	;
	if v162 == int32(0) {
		v187 = v160
		v188 = v161
		goto L23
	} else {
		goto L44
	}
L26:
	;
	v160 = v75
	v161 = v81
	v162 = v88
	goto L25
L27:
	;
	goto L28
L28:
	;
	v92 = int32(0)
	if base.B2i32(v75&int32(3) == v92)|int32(0) == v92 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v128 == int32(0) {
		v187 = v125
		v188 = v126
		goto L23
	} else {
		goto L38
	}
L30:
	;
	v104 = v75
	v105 = v81
	v106 = v88
	goto L33
L31:
	;
	goto L32
L32:
	;
	v125 = v75
	v126 = v81
	v127 = v88
	v128 = int32(1)
	goto L29
L33:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v108)
	if v108 == int32(0) {
		v167 = v104
		v168 = v105
		v169 = v106
		goto L24
	} else {
		goto L35
	}
L34:
	;
	v125 = v119
	v126 = v113
	v127 = v115
	v128 = v117
	goto L29
L35:
	;
	v112 = int32(1)
	v113 = v105 + v112
	v115 = v106 - v112
	v116 = int32(0)
	v117 = base.B2i32(v115 != v116)
	v119 = v104 + v112
	if v119&int32(3) == v116 {
		v125 = v119
		v126 = v113
		v127 = v115
		v128 = v117
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v115 != 0 {
		v104 = v119
		v105 = v113
		v106 = v115
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if base.B2i32(v131 == int32(0))|base.B2i32(base.Ui32(v127) < base.Ui32(int32(4))) != 0 {
		v160 = v125
		v161 = v126
		v162 = v127
		goto L25
	} else {
		goto L39
	}
L39:
	;
	v138 = v125
	v139 = v126
	v140 = v127
	goto L40
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v146 = int32(-2139062144)
	if (int32(16843008)-v143|v143)&v146 != v146 {
		v167 = v138
		v168 = v139
		v169 = v140
		goto L24
	} else {
		goto L42
	}
L41:
	;
	v160 = v154
	v161 = v152
	v162 = v156
	goto L25
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v143
	v151 = int32(4)
	v152 = v139 + v151
	v154 = v138 + v151
	v156 = v140 - v151
	if base.Ui32(int32(3)) < base.Ui32(v156) {
		v138 = v154
		v139 = v152
		v140 = v156
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v167 = v160
	v168 = v161
	v169 = v162
	goto L24
L45:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v176)
	if v176 == int32(0) {
		v187 = v172
		v188 = v173
		goto L23
	} else {
		goto L47
	}
L46:
	;
	v187 = v183
	v188 = v181
	goto L23
L47:
	;
	v180 = int32(1)
	v181 = v173 + v180
	v183 = v172 + v180
	v185 = v174 - v180
	if v185 != 0 {
		v172 = v183
		v173 = v181
		v174 = v185
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v207)))
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v207)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+256)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(l2)+248)) = v212
	goto L1
}
func F_PrepareInvalidationState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v49 int64
	_ = v49
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0]))
	if v5 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[1]))
		v19 = F_MemoryContextAllocZero(m, v17, int32(44))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v24
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[2]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v28
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0]))
			if v31 != 0 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				if v32-v33 != v35-v36 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_PrepareInvalidationState_0), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_PrepareInvalidationState_1), int32(717), int32(_a_F_PrepareInvalidationState_2))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v32
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v41
					*(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0])) = v19
					return v19
				}
			} else {
				v49 = int64(0)
				*(*int64)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[3])) = v49
				*(*int64)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[4])) = v49
				*(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0])) = v19
				return v19
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[2]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		if v8 != v11 {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[1]))
			v19 = F_MemoryContextAllocZero(m, v17, int32(44))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v24
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[2]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v28
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0]))
				if v31 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
					if v32-v33 != v35-v36 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_PrepareInvalidationState_0), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_PrepareInvalidationState_1), int32(717), int32(_a_F_PrepareInvalidationState_2))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v32
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v41
						*(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0])) = v19
						return v19
					}
				} else {
					v49 = int64(0)
					*(*int64)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[3])) = v49
					*(*int64)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[4])) = v49
					*(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0])) = v19
					return v19
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInvalidationState[0]))
			return v14
		}
	}
}
func F_PrepareSortSupportFromGistIndexRel(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	if v11 == int32(783) {
		v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+10)))
		v18 = v14<<(uint(int32(2))%32) - int32(4)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v18+v19)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v22+v18)))
		v25 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v25)
		v28 = F_get_opfamily_proc(m, v24, v21, v21, int32(11))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			if v28 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
					F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_0), v8)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_1), int32(205), int32(_a_F_PrepareSortSupportFromGistIndexRel_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v33 = F_OidFunctionCall1Coll(m, v28, int32(0), l1)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					m.G0 = v8 + int32(32)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+84))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v43
			F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_3), v8+int32(16))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_1), int32(194), int32(_a_F_PrepareSortSupportFromGistIndexRel_2))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_p_isdigit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(int32(2))%32))))
			return base.B2i32(base.Ui32(v12-int32(48)) < base.Ui32(int32(10)))
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
			return base.B2i32(base.Ui32(v24-int32(48)) < base.Ui32(int32(10)))
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v32))))
		return base.B2i32(base.Ui32((v34-int32(48))&int32(255)) < base.Ui32(int32(10)))
	}
}
func F_pairingheap_allocate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc(m, int32(12))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		return v5
	}
}
func F_pairingheap_remove_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	return v9
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v81
	return v9
L5:
	;
	v73 = v10
	goto L4
L6:
	;
	goto L7
L7:
	;
	v15 = int32(0)
	v16 = v10
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v15 == int32(0) {
		v73 = v44
		goto L4
	} else {
		goto L26
	}
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v15
	v44 = v16
	goto L10
L12:
	;
	goto L13
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, v16, v22, v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v34 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v35 = v16
	goto L18
L17:
	;
	v35 = v22
	goto L18
L18:
	;
	if v29 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v36 = v22
	goto L21
L20:
	;
	v36 = v16
	goto L21
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v35
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35
	if v26 != 0 {
		v15 = v36
		v16 = v26
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v44 = v36
	goto L10
L26:
	;
	v50 = v44
	v52 = v15
	goto L27
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v50, v52, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L29
	}
L28:
	;
	v73 = v65
	goto L4
L29:
	;
	v63 = base.B2i32(v60 < int32(0))
	if v60 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v64 = v50
	goto L32
L31:
	;
	v64 = v52
	goto L32
L32:
	;
	if v60 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v65 = v52
	goto L35
L34:
	;
	v65 = v50
	goto L35
L35:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v66 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v64
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v65
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v64
	if v57 != 0 {
		v50 = v65
		v52 = v57
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
}
func F_parse_dispatch_option(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
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
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	v2 = int32(_a_F_parse_dispatch_option_0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_dispatch_option[0])))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v5 == int32(0))|base.B2i32(v5 != v8) != 0 {
		v26 = v5
		v27 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26-v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	goto L1
L3:
	;
	v11 = v2
	v12 = l0
	goto L4
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v16 == int32(0) {
		v26 = v16
		v27 = v15
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v26 = v16
	v27 = v15
	goto L2
L6:
	;
	v19 = int32(1)
	if v16 == v15 {
		v11 = v11 + v19
		v12 = v12 + v19
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	v33 = int32(_a_F_parse_dispatch_option_1)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_dispatch_option[1])))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v36 == int32(0))|base.B2i32(v36 != v39) != 0 {
		v57 = v36
		v58 = v39
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v57-v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v42 = v33
	v43 = l0
	goto L14
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v47
		v58 = v46
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v57 = v47
	v58 = v46
	goto L12
L16:
	;
	v50 = int32(1)
	if v47 == v46 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	v64 = int32(_a_F_parse_dispatch_option_2)
	goto L23
L21:
	;
	if v102-v103 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	goto L24
L24:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_dispatch_option[2])))
	if v71 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v72 = v64
	v73 = l0
	v74 = int32(9)
	v75 = v71
	goto L29
L26:
	;
	v98 = l0
	v102 = int32(0)
	goto L27
L27:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	goto L21
L28:
	;
	v98 = v93
	v102 = v95
	goto L27
L29:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if base.B2i32(v75 != v77)|base.B2i32(v77 == int32(0)) != 0 {
		v93 = v73
		v95 = v75
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v93 = v87
	v95 = int32(0)
	goto L28
L31:
	;
	v83 = v74 - int32(1)
	if v83 == int32(0) {
		v93 = v73
		v95 = v75
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v86 = int32(1)
	v87 = v73 + v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v88 != 0 {
		v72 = v72 + v86
		v73 = v87
		v74 = v83
		v75 = v88
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	return int32(2)
L35:
	;
	goto L36
L36:
	;
	v115 = int32(_a_F_parse_dispatch_option_3)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_dispatch_option[3])))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v118 == int32(0))|base.B2i32(v118 != v121) != 0 {
		v139 = v118
		v140 = v121
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v139-v140 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	v124 = v115
	v125 = l0
	goto L40
L40:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	if v129 == int32(0) {
		v139 = v129
		v140 = v128
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v139 = v129
	v140 = v128
	goto L38
L42:
	;
	v132 = int32(1)
	if v129 == v128 {
		v124 = v124 + v132
		v125 = v125 + v132
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	return int32(3)
L45:
	;
	goto L46
L46:
	;
	v148 = int32(_a_F_parse_dispatch_option_4)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_dispatch_option[4])))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v151 == int32(0))|base.B2i32(v151 != v154) != 0 {
		v172 = v151
		v173 = v154
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v172-v173 != 0 {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	goto L47
L49:
	;
	v157 = v148
	v158 = l0
	goto L50
L50:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v162 == int32(0) {
		v172 = v162
		v173 = v161
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v172 = v162
	v173 = v161
	goto L48
L52:
	;
	v165 = int32(1)
	if v162 == v161 {
		v157 = v157 + v165
		v158 = v158 + v165
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v175 = int32(5)
	goto L56
L55:
	;
	v175 = int32(4)
	goto L56
L56:
	;
	return v175
}
func F_parse_scalar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(v8 == int32(11))|base.B2i32(base.Ui32(int32(-3)) < base.Ui32(v8&int32(-9)-int32(3))) == v3 {
		v20 = int32(11)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v23 != 0 {
			v24 = int32(10)
		} else {
			v24 = v20
		}
		if v8 == int32(12) {
			v27 = v20
		} else {
			v27 = v24
		}
		return v27
	} else {
		if v7 == int32(0) {
			v31 = F_json_lex(m, l0)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				return v31
			}
		} else {
			if v8 == int32(1) {
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
				if v38 != int32(1) {
					v69 = F_json_lex(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						if v69 != 0 {
							v87 = v69
							return v87
						} else {
							v72 = v3
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v74 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v73, v72, v8)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								if v72 == int32(0) {
									v87 = v74
									return v87
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v78&int32(4) == int32(0) {
										v87 = v74
										return v87
									} else {
										v83 = v74
										v84 = v72
										F_pfree(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = v83
											return v87
										}
									}
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					v43 = F_pstrdup(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						if v43 != 0 {
							v64 = v43
							v65 = F_json_lex(m, l0)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								if v65 == int32(0) {
									v72 = v64
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v74 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v73, v72, v8)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										if v72 == int32(0) {
											v87 = v74
											return v87
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v78&int32(4) == int32(0) {
												v87 = v74
												return v87
											} else {
												v83 = v74
												v84 = v72
												F_pfree(m, v84)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v87 = v83
													return v87
												}
											}
										}
									}
								} else {
									v83 = v65
									v84 = v64
									F_pfree(m, v84)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = v83
										return v87
									}
								}
							}
						} else {
							return int32(16)
						}
					}
				}
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v49 = v47 - v48
				v52 = F_palloc(m, v49+int32(1))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 == int32(0) {
						return int32(16)
					} else {
						if v49 != 0 {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							base.MemoryCopy(m, v52, v58, v49)
						} else {
						}
						v61 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v49+v52))) = uint8(v61)
						v64 = v52
						v65 = F_json_lex(m, l0)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v65 == int32(0) {
								v72 = v64
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v74 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v73, v72, v8)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									if v72 == int32(0) {
										v87 = v74
										return v87
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v78&int32(4) == int32(0) {
											v87 = v74
											return v87
										} else {
											v83 = v74
											v84 = v72
											F_pfree(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v87 = v83
												return v87
											}
										}
									}
								}
							} else {
								v83 = v65
								v84 = v64
								F_pfree(m, v84)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									v87 = v83
									return v87
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_partitioned_table_reloptions(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	if l0 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_partitioned_table_reloptions_0), int32(0))
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_F_partitioned_table_reloptions_1), int32(0))
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_partitioned_table_reloptions_2), int32(2026), int32(_a_F_partitioned_table_reloptions_3))
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F_pattern_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	switch l1 - int32(1) {
	case 0:
		v15 = F_like_fixed_prefix(m, l0, int32(1), l2, l3, l4)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	case 1:
		v19 = F_regex_fixed_prefix(m, l0, int32(0), l2, l3, l4)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			return v19
		}
	case 2:
		v23 = F_regex_fixed_prefix(m, l0, int32(1), l2, l3, l4)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	case 3:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
		v32 = F_datumCopy(m, v30, v31, v29)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
			v36 = F_makeConst(m, v26, v27, v28, v29, v32, v34, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v36
				if l4 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
				} else {
				}
				return int32(1)
			}
		}
	default:
		v9 = F_like_fixed_prefix(m, l0, int32(0), l2, l3, l4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_pgarch_call_module_shutdown_cb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_call_module_shutdown_cb[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_call_module_shutdown_cb[1]))
		m.T0[v5].(func(*base.Module, int32))(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_pgarch_die(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_die[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(-1)
	return
}
func F_pgl_exit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[0]))
	if v4 != 0 {
		v5 = F_fclose(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[0])) = int32(0)
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
			if v11 != 0 {
				v12 = F_fflush(m, v11)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
					v16 = F_fclose(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1])) = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
						*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
						m.Env.Exit(m, l0)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
				m.Env.Exit(m, l0)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
		if v11 != 0 {
			v12 = F_fflush(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
				v16 = F_fclose(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1])) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
					*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
					m.Env.Exit(m, l0)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
			m.Env.Exit(m, l0)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_pgl_popen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_popen[0]))
	if v5 != 0 {
		v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_popen[1])) = int32(52)
		return int32(0)
	}
}
func F_pgstatindex_impl(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int64
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v214 float64
	_ = v214
	var v216 float64
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 float64
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	v15 = int64(0)
	v24 = m.G0
	v26 = v24 - int32(240)
	m.G0 = v26
	v29 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
	if v34 != int32(105) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L91
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L87
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L83
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L79
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
	if v37 != int32(403) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+118)))
	if v40 == int32(116) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v43 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+18)))
	if v47 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v50 = int32(0)
	v53 = F_ReadBufferExtended(m, l0, v50, v50, v50, v29)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+36))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	F_ReleaseBuffer(m, v53)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L15:
	;
	if v53 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+(v53^int32(-1))<<(uint(int32(2))%32))))
	v72 = v64
	goto L14
L17:
	;
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[1]))
	v72 = v66 + v53<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L19:
	;
	v79 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(v79) < base.Ui32(int32(2)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v207 = v15
	v208 = v15
	v209 = v15
	v210 = v15
	v211 = v15
	v214 = float64(0)
	v216 = float64(0)
	goto L23
L22:
	;
	v89 = int32(1)
	v99 = v15
	v100 = v15
	v101 = v15
	v102 = v15
	v103 = v15
	v104 = v15
	v105 = v15
	goto L24
L23:
	;
	F_relation_close(m, l0, int32(1))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L53
	}
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[2]))
	if v109 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v207 = v176
	v208 = v177
	v209 = v178
	v210 = v179
	v211 = v180
	v214 = base.F64_convert_i64_u(v182)
	v216 = base.F64_convert_i64_u(v181)
	goto L23
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v112 = int32(0)
	v114 = F_ReadBufferExtended(m, l0, v112, v89, v112, v29)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	F_LockBuffer(m, v114, int32(1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v114 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_LockBuffer(m, v114, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L50
	}
L33:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+16)))
	v138 = v137 + v136
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	if v139&int32(4) != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[0]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+(v114^int32(-1))<<(uint(int32(2))%32))))
	v136 = v128
	goto L33
L35:
	;
	goto L36
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[1]))
	v136 = v130 + v114<<(uint(int32(13))%32) + int32(-8192)
	goto L33
L37:
	;
	v176 = v99
	v177 = v100 + int64(1)
	v178 = v101
	v179 = v102
	v180 = v103
	v181 = v104
	v182 = v105
	goto L32
L38:
	;
	goto L39
L39:
	;
	if v139&int32(16) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v176 = v99
	v177 = v100
	v178 = v101 + int64(1)
	v179 = v102
	v180 = v103
	v181 = v104
	v182 = v105
	goto L32
L41:
	;
	goto L42
L42:
	;
	if v139&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+14)))
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	v152 = v150 - v151
	v153 = int32(0)
	if v153 < v152 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	v176 = v99
	v177 = v100
	v178 = v101
	v179 = v102
	v180 = v103 + int64(1)
	v181 = v104
	v182 = v105
	goto L32
L46:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v176 = v99 + int64(1)
	v177 = v100
	v178 = v101
	v179 = v102 + base.I64_extend_i32_s(v137-int32(24))
	v180 = v103
	v181 = v104 + base.I64_extend_i32_u(base.B2i32(v157 != int32(0))&base.B2i32(base.Ui32(v157) < base.Ui32(v89)))
	v182 = v105 + base.I64_extend_i32_u(v156)
	goto L32
L47:
	;
	v156 = v152
	goto L49
L48:
	;
	v156 = v153
	goto L49
L49:
	;
	goto L46
L50:
	;
	F_ReleaseBuffer(m, v114)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v189 = v89 + int32(1)
	if v189 != v79 {
		v89 = v189
		v99 = v176
		v100 = v177
		v101 = v178
		v102 = v179
		v103 = v180
		v104 = v181
		v105 = v182
		goto L24
	} else {
		goto L52
	}
L52:
	;
	goto L25
L53:
	;
	v223 = F_get_call_result_type(m, l1, int32(0), v26+int32(236))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v223 != int32(1) {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v75
	v231 = F_psprintf(m, int32(_a_F_pgstatindex_impl_0), v26+int32(144))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+128)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v26)+192)) = v231
	v238 = F_psprintf(m, int32(_a_F_pgstatindex_impl_0), v26+int32(128))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+196)) = v238
	*(*int64)(unsafe.Add(mBase, uint32(v26)+112)) = (v207+v211+v209+v208)<<(uint(int64(13))%64) - int64(-8192)
	v252 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(112))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v26)+200)) = v252
	v259 = F_psprintf(m, int32(_a_F_pgstatindex_impl_2), v26+int32(96))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+80)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v26)+204)) = v259
	v266 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(80))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v26)+208)) = v266
	v273 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26-int32(-64))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v26)+212)) = v273
	v280 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(48))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v26)+216)) = v280
	v287 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(32))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+220)) = v287
	if v210 != int64(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v307
	if v207 != int64(0) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v292 = float64(100)
	*(*float64)(unsafe.Add(mBase, uint32(v26)+16)) = base.F64_sub(v292, base.F64_mul(base.F64_div(v214, base.F64_convert_i64_u(v210)), v292))
	v302 = F_psprintf(m, int32(_a_F_pgstatindex_impl_3), v26+int32(16))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v305 = F_pstrdup(m, int32(_a_F_pgstatindex_impl_4))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	v307 = v302
	goto L64
L69:
	;
	v307 = v305
	goto L64
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v26)+236))
	v325 = F_TupleDescGetAttInMetadata(m, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L76
	}
L71:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v26))) = base.F64_mul(base.F64_div(v216, base.F64_convert_i64_u(v207)), float64(100))
	v317 = F_psprintf(m, int32(_a_F_pgstatindex_impl_3), v26)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v320 = F_pstrdup(m, int32(_a_F_pgstatindex_impl_4))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	v322 = v317
	goto L70
L75:
	;
	v322 = v320
	goto L70
L76:
	;
	v329 = F_BuildTupleFromCStrings(m, v325, v26+int32(192))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v329)+16))
	v332 = F_HeapTupleHeaderGetDatum(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	m.G0 = v26 + int32(240)
	return v332
L79:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+176)) = v345 + int32(4)
	F_errmsg(m, int32(_a_F_pgstatindex_impl_5), v26+int32(176))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(225), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_pgstatindex_impl_8), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(235), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+160)) = v382 + int32(4)
	F_errmsg(m, int32(_a_F_pgstatindex_impl_9), v26+int32(160))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(247), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_pgstatindex_impl_10), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(344), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pktreader_pull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = l0 + int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v64 = F_pullf_read(m, l1, l2, l3)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L22
	}
L4:
	;
	return v62
L5:
	;
	v21 = v9
	goto L8
L6:
	;
	v44 = v14
	goto L7
L7:
	;
	if l2 < v44 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	if v21 == int32(1) {
		v62 = int32(0)
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v44 = v35
	goto L7
L10:
	;
	v27 = F_parse_new_len(m, l1, v13)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v27 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return v27
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 == int32(0) {
		v21 = v27
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v47 = l2
	goto L19
L18:
	;
	v47 = v44
	goto L19
L19:
	;
	v48 = F_pullf_read(m, l1, v47, l3)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v48 <= int32(0) {
		v62 = v48
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v52 - v48
	v62 = v48
	goto L4
L22:
	;
	return v64
}
func F_plainnode(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 == v9 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 << (uint(int32(1)) % 32)
			v16 = F_repalloc(m, v7, v8*int32(24))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v16
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v20 = v16
				v21 = v19
				v24 = v21*int32(12) + v20
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v26
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				*(*int64)(unsafe.Add(mBase, uint32(v24))) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				if v31 == int32(1) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34 + int32(1)
					F_pfree(m, l1)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
					if v40 == int32(1) {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v48 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v43+v44*int32(12))+4)) = v48
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v48
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_plainnode(m, l0, v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_pfree(m, l1)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v59 + int32(1)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_plainnode(m, l0, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v66+v59*int32(12))+4)) = v70 - v59
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_plainnode(m, l0, v73)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								F_pfree(m, l1)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			v20 = v7
			v21 = v8
			v24 = v21*int32(12) + v20
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v26
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
			*(*int64)(unsafe.Add(mBase, uint32(v24))) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
			if v31 == int32(1) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34 + int32(1)
				F_pfree(m, l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
				if v40 == int32(1) {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v48 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v43+v44*int32(12))+4)) = v48
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v48
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					F_plainnode(m, l0, v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_pfree(m, l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v59 + int32(1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					F_plainnode(m, l0, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v66+v59*int32(12))+4)) = v70 - v59
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_plainnode(m, l0, v73)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							F_pfree(m, l1)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_pnstrdup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	v6 = F_memchr(m, l0, int32(0), l1)
	mBase = m.M
	if v6 != 0 {
		v8 = v6 - l0
	} else {
		v8 = l1
	}
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pnstrdup[0]))
	v11 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)) = uint8(v11)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = m.T0[v17].(func(*base.Module, int32, int32, int32) int32)(m, v10, v8+int32(1), v11)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			base.MemoryCopy(m, v18, l0, v8)
		} else {
		}
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8+v18))) = uint8(v24)
		return v18
	}
}
func F_policy_role_list_to_array(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
	return v82
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v12 = F_palloc(m, int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
	v20 = F_palloc(m, v16<<(uint(int32(2))%32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v82 = v12
	goto L1
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v22 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = v3
	goto L11
L9:
	;
	goto L10
L10:
	;
	return v20
L11:
	;
	v32 = v28 << (uint(int32(2)) % 32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v36 == int32(4) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v39 == int32(1) {
		v82 = v20
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v66 = F_get_rolespec_oid(m, v35, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L25
	}
L16:
	;
	v44 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v44 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v82 = v20
	goto L1
L21:
	;
	F_errmsg(m, int32(_a_F_policy_role_list_to_array_0), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(_a_F_policy_role_list_to_array_1), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_policy_role_list_to_array_2), int32(170), int32(_a_F_policy_role_list_to_array_3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v32))) = v66
	v70 = v28 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v70 < v71 {
		v28 = v70
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L12
}
func F_portuguese_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v437 int32
	_ = v437
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v481 int32
	_ = v481
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v558 int32
	_ = v558
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v676 int32
	_ = v676
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v710 int32
	_ = v710
	var v721 int32
	_ = v721
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v799 int32
	_ = v799
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v917 int32
	_ = v917
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1036 int32
	_ = v1036
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1069 int32
	_ = v1069
	var v1080 int32
	_ = v1080
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1157 int32
	_ = v1157
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1195 int32
	_ = v1195
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1292 int32
	_ = v1292
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1343 int32
	_ = v1343
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1376 int32
	_ = v1376
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1414 int32
	_ = v1414
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1465 int32
	_ = v1465
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1499 int32
	_ = v1499
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1539 int32
	_ = v1539
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1590 int32
	_ = v1590
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1616 int32
	_ = v1616
	var v1623 int32
	_ = v1623
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1661 int32
	_ = v1661
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1712 int32
	_ = v1712
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1738 int32
	_ = v1738
	var v1746 int32
	_ = v1746
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1766 int32
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2117 int32
	_ = v2117
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2149 int32
	_ = v2149
	var v2155 int32
	_ = v2155
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2259 int32
	_ = v2259
	var v2262 int32
	_ = v2262
	var v2266 int32
	_ = v2266
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2288 int32
	_ = v2288
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2300 int32
	_ = v2300
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v2300
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v16 = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 <= v16 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v121
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L46
L4:
	;
	goto L3
L5:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v118
	goto L2
L6:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L21
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v56 = v9
	v58 = v17
	goto L6
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v16))))
	v23 = v21 - int32(163)
	v24 = int32(0)
	if base.B2i32(v23 == v24)|base.B2i32(v23 == int32(18)) == v24 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = F_find_among(m, l0, int32(_a_F_portuguese_UTF_8_stem_0), int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
	switch v33 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L5
	}
L12:
	;
	v50 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v44 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = v37
	v58 = v41
	goto L6
L15:
	;
	if int32(0) <= v44 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v2300 = v44
	goto L1
L17:
	;
	if int32(0) <= v50 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v2300 = v50
	goto L1
L19:
	;
	if v111 < int32(0) {
		goto L4
	} else {
		goto L39
	}
L21:
	;
	goto L22
L22:
	;
	goto L23
L23:
	;
	v66 = v56
	v68 = int32(1)
	goto L26
L25:
	;
	v111 = v96
	goto L19
L26:
	;
	if v58 <= v66 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v111 = int32(-1)
	goto L19
L29:
	;
	goto L30
L30:
	;
	v73 = v66 + int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v66))))
	if base.Ui32(v75) < base.Ui32(int32(192)) {
		v96 = v73
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v97 = int32(1)
	if v97 < v68 {
		v66 = v96
		v68 = v68 - v97
		goto L26
	} else {
		goto L38
	}
L32:
	;
	if v58 <= v73 {
		v96 = v73
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v82 = v73
	goto L34
L34:
	;
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59+v82))))
	if int32(-65) < v85 {
		v96 = v82
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v96 = v58
	goto L31
L36:
	;
	v89 = v82 + int32(1)
	if v89 != v58 {
		v82 = v89
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L27
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	goto L5
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1292 = v126
	goto L302
L41:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+8)) = v1265
	goto L40
L42:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1265 = v1262 + v1261
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L175
L44:
	;
	if v244 != 0 {
		goto L43
	} else {
		goto L68
	}
L45:
	;
	v244 = v237
	goto L44
L46:
	;
	if v139 <= v126 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v237 = int32(0)
	goto L45
L48:
	;
	v244 = int32(-1)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v155 = int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v140))))
	if base.Ui32(v157) < base.Ui32(int32(192)) {
		v214 = v157
		v215 = v155
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if int32(250) < v214 {
		v237 = v215
		goto L45
	} else {
		goto L64
	}
L52:
	;
	v161 = v126 + int32(1)
	if v161 == v139 {
		v214 = v157
		v215 = v155
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v140))))
	v166 = v164 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v157) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v140))))
	v182 = v180 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v157) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v170 = v126 + int32(2)
	if v170 != v139 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v214 = v157<<(uint(int32(6))%32)&int32(1984) | v166
	v215 = int32(2)
	goto L51
L58:
	;
	goto L57
L59:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+v186))))
	v214 = v199&int32(63) | (v157<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v166<<(uint(int32(12))%32) | v182<<(uint(int32(6))%32))
	v215 = int32(4)
	goto L51
L60:
	;
	v186 = v126 + int32(3)
	if v186 != v139 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v214 = v157<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v166<<(uint(int32(6))%32) | v182
	v215 = int32(3)
	goto L51
L63:
	;
	goto L62
L64:
	;
	v219 = v214 - int32(97)
	if v219 < int32(0) {
		v237 = v215
		goto L45
	} else {
		goto L65
	}
L65:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v219)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v225)>>(uint(v219&int32(7))%32))&int32(1) == int32(0) {
		v237 = v215
		goto L45
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v215 + v126
	goto L67
L67:
	;
	goto L47
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L71
L69:
	;
	if v362 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L70:
	;
	v362 = v355
	goto L69
L71:
	;
	if v258 <= v245 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v355 = int32(0)
	goto L70
L73:
	;
	v362 = int32(-1)
	goto L69
L74:
	;
	goto L75
L75:
	;
	v274 = int32(1)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v259))))
	if base.Ui32(v276) < base.Ui32(int32(192)) {
		v333 = v276
		v334 = v274
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if int32(250) < v333 {
		goto L89
	} else {
		goto L90
	}
L77:
	;
	v280 = v245 + int32(1)
	if v280 == v258 {
		v333 = v276
		v334 = v274
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v259))))
	v285 = v283 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v276) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v259))))
	v301 = v299 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v276) {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v289 = v245 + int32(2)
	if v289 != v258 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v333 = v276<<(uint(int32(6))%32)&int32(1984) | v285
	v334 = int32(2)
	goto L76
L83:
	;
	goto L82
L84:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259+v305))))
	v333 = v318&int32(63) | (v276<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v285<<(uint(int32(12))%32) | v301<<(uint(int32(6))%32))
	v334 = int32(4)
	goto L76
L85:
	;
	v305 = v245 + int32(3)
	if v305 != v258 {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v333 = v276<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v285<<(uint(int32(6))%32) | v301
	v334 = int32(3)
	goto L76
L88:
	;
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v334 + v245
	goto L93
L90:
	;
	v338 = v333 - int32(97)
	if v338 < int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v338)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v344)>>(uint(v338&int32(7))%32))&int32(1) != 0 {
		v355 = v334
		goto L70
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	goto L72
L94:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v386 = v376
	goto L99
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v245
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L125
L97:
	;
	if int32(0) <= v481 {
		v1261 = v481
		goto L42
	} else {
		goto L122
	}
L98:
	;
	v481 = v453
	goto L97
L99:
	;
	if v377 <= v386 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v481 = int32(-1)
	goto L97
L102:
	;
	goto L103
L103:
	;
	v393 = int32(1)
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v378))))
	if base.Ui32(v395) < base.Ui32(int32(192)) {
		v452 = v395
		v453 = v393
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if int32(250) < v452 {
		goto L117
	} else {
		goto L118
	}
L105:
	;
	v399 = v386 + int32(1)
	if v399 == v377 {
		v452 = v395
		v453 = v393
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399+v378))))
	v404 = v402 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v395) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v378))))
	v420 = v418 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v395) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v408 = v386 + int32(2)
	if v408 != v377 {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v452 = v395<<(uint(int32(6))%32)&int32(1984) | v404
	v453 = int32(2)
	goto L104
L111:
	;
	goto L110
L112:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378+v424))))
	v452 = v437&int32(63) | (v395<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v404<<(uint(int32(12))%32) | v420<<(uint(int32(6))%32))
	v453 = int32(4)
	goto L104
L113:
	;
	v424 = v386 + int32(3)
	if v424 != v377 {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v452 = v395<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v404<<(uint(int32(6))%32) | v420
	v453 = int32(3)
	goto L104
L116:
	;
	goto L115
L117:
	;
	v470 = v453 + v386
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v470
	v386 = v470
	goto L99
L118:
	;
	v457 = v452 - int32(97)
	if v457 < int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v463)>>(uint(v457&int32(7))%32))&int32(1) != 0 {
		goto L98
	} else {
		goto L120
	}
L120:
	;
	goto L117
L122:
	;
	goto L96
L123:
	;
	if v603 != 0 {
		goto L43
	} else {
		goto L147
	}
L124:
	;
	v603 = v596
	goto L123
L125:
	;
	if v498 <= v245 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v596 = int32(0)
	goto L124
L127:
	;
	v603 = int32(-1)
	goto L123
L128:
	;
	goto L129
L129:
	;
	v514 = int32(1)
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v499))))
	if base.Ui32(v516) < base.Ui32(int32(192)) {
		v573 = v516
		v574 = v514
		goto L130
	} else {
		goto L131
	}
L130:
	;
	if int32(250) < v573 {
		v596 = v574
		goto L124
	} else {
		goto L143
	}
L131:
	;
	v520 = v245 + int32(1)
	if v520 == v498 {
		v573 = v516
		v574 = v514
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520+v499))))
	v525 = v523 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v516) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v499))))
	v541 = v539 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v516) {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v529 = v245 + int32(2)
	if v529 != v498 {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v573 = v516<<(uint(int32(6))%32)&int32(1984) | v525
	v574 = int32(2)
	goto L130
L137:
	;
	goto L136
L138:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499+v545))))
	v573 = v558&int32(63) | (v516<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v525<<(uint(int32(12))%32) | v541<<(uint(int32(6))%32))
	v574 = int32(4)
	goto L130
L139:
	;
	v545 = v245 + int32(3)
	if v545 != v498 {
		goto L138
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v573 = v516<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v525<<(uint(int32(6))%32) | v541
	v574 = int32(3)
	goto L130
L142:
	;
	goto L141
L143:
	;
	v578 = v573 - int32(97)
	if v578 < int32(0) {
		v596 = v574
		goto L124
	} else {
		goto L144
	}
L144:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v578)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v584)>>(uint(v578&int32(7))%32))&int32(1) == int32(0) {
		v596 = v574
		goto L124
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v574 + v245
	goto L146
L146:
	;
	goto L126
L147:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v625 = v615
	goto L150
L148:
	;
	if int32(0) <= v721 {
		v1261 = v721
		goto L42
	} else {
		goto L172
	}
L149:
	;
	v721 = v692
	goto L148
L150:
	;
	if v616 <= v625 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v721 = int32(-1)
	goto L148
L153:
	;
	goto L154
L154:
	;
	v632 = int32(1)
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625+v617))))
	if base.Ui32(v634) < base.Ui32(int32(192)) {
		v691 = v634
		v692 = v632
		goto L155
	} else {
		goto L156
	}
L155:
	;
	if int32(250) < v691 {
		goto L149
	} else {
		goto L168
	}
L156:
	;
	v638 = v625 + int32(1)
	if v638 == v616 {
		v691 = v634
		v692 = v632
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638+v617))))
	v643 = v641 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v634) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647+v617))))
	v659 = v657 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v634) {
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v647 = v625 + int32(2)
	if v647 != v616 {
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v691 = v634<<(uint(int32(6))%32)&int32(1984) | v643
	v692 = int32(2)
	goto L155
L162:
	;
	goto L161
L163:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617+v663))))
	v691 = v676&int32(63) | (v634<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v643<<(uint(int32(12))%32) | v659<<(uint(int32(6))%32))
	v692 = int32(4)
	goto L155
L164:
	;
	v663 = v625 + int32(3)
	if v663 != v616 {
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v691 = v634<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v643<<(uint(int32(6))%32) | v659
	v692 = int32(3)
	goto L155
L167:
	;
	goto L166
L168:
	;
	v696 = v691 - int32(97)
	if v696 < int32(0) {
		goto L149
	} else {
		goto L169
	}
L169:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v696)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v702)>>(uint(v696&int32(7))%32))&int32(1) == int32(0) {
		goto L149
	} else {
		goto L170
	}
L170:
	;
	v710 = v692 + v625
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v710
	v625 = v710
	goto L150
L172:
	;
	goto L43
L173:
	;
	if v843 != 0 {
		goto L40
	} else {
		goto L198
	}
L174:
	;
	v843 = v836
	goto L173
L175:
	;
	if v739 <= v126 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v836 = int32(0)
	goto L174
L177:
	;
	v843 = int32(-1)
	goto L173
L178:
	;
	goto L179
L179:
	;
	v755 = int32(1)
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v740))))
	if base.Ui32(v757) < base.Ui32(int32(192)) {
		v814 = v757
		v815 = v755
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if int32(250) < v814 {
		goto L193
	} else {
		goto L194
	}
L181:
	;
	v761 = v126 + int32(1)
	if v761 == v739 {
		v814 = v757
		v815 = v755
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761+v740))))
	v766 = v764 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v757) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770+v740))))
	v782 = v780 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v757) {
		goto L189
	} else {
		goto L190
	}
L184:
	;
	v770 = v126 + int32(2)
	if v770 != v739 {
		goto L183
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v814 = v757<<(uint(int32(6))%32)&int32(1984) | v766
	v815 = int32(2)
	goto L180
L187:
	;
	goto L186
L188:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740+v786))))
	v814 = v799&int32(63) | (v757<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v766<<(uint(int32(12))%32) | v782<<(uint(int32(6))%32))
	v815 = int32(4)
	goto L180
L189:
	;
	v786 = v126 + int32(3)
	if v786 != v739 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v814 = v757<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v766<<(uint(int32(6))%32) | v782
	v815 = int32(3)
	goto L180
L192:
	;
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v815 + v126
	goto L197
L194:
	;
	v819 = v814 - int32(97)
	if v819 < int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v819)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v825)>>(uint(v819&int32(7))%32))&int32(1) != 0 {
		v836 = v815
		goto L174
	} else {
		goto L196
	}
L196:
	;
	goto L193
L197:
	;
	goto L176
L198:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L201
L199:
	;
	if v961 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L200:
	;
	v961 = v954
	goto L199
L201:
	;
	if v857 <= v844 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v954 = int32(0)
	goto L200
L203:
	;
	v961 = int32(-1)
	goto L199
L204:
	;
	goto L205
L205:
	;
	v873 = int32(1)
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844+v858))))
	if base.Ui32(v875) < base.Ui32(int32(192)) {
		v932 = v875
		v933 = v873
		goto L206
	} else {
		goto L207
	}
L206:
	;
	if int32(250) < v932 {
		goto L219
	} else {
		goto L220
	}
L207:
	;
	v879 = v844 + int32(1)
	if v879 == v857 {
		v932 = v875
		v933 = v873
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879+v858))))
	v884 = v882 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v875) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888+v858))))
	v900 = v898 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v875) {
		goto L215
	} else {
		goto L216
	}
L210:
	;
	v888 = v844 + int32(2)
	if v888 != v857 {
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v932 = v875<<(uint(int32(6))%32)&int32(1984) | v884
	v933 = int32(2)
	goto L206
L213:
	;
	goto L212
L214:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v858+v904))))
	v932 = v917&int32(63) | (v875<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v884<<(uint(int32(12))%32) | v900<<(uint(int32(6))%32))
	v933 = int32(4)
	goto L206
L215:
	;
	v904 = v844 + int32(3)
	if v904 != v857 {
		goto L214
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v932 = v875<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v884<<(uint(int32(6))%32) | v900
	v933 = int32(3)
	goto L206
L218:
	;
	goto L217
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v933 + v844
	goto L223
L220:
	;
	v937 = v932 - int32(97)
	if v937 < int32(0) {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v937)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v943)>>(uint(v937&int32(7))%32))&int32(1) != 0 {
		v954 = v933
		goto L200
	} else {
		goto L222
	}
L222:
	;
	goto L219
L223:
	;
	goto L202
L224:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v985 = v975
	goto L229
L225:
	;
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v844
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L255
L227:
	;
	if int32(0) <= v1080 {
		v1261 = v1080
		goto L42
	} else {
		goto L252
	}
L228:
	;
	v1080 = v1052
	goto L227
L229:
	;
	if v976 <= v985 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1080 = int32(-1)
	goto L227
L232:
	;
	goto L233
L233:
	;
	v992 = int32(1)
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985+v977))))
	if base.Ui32(v994) < base.Ui32(int32(192)) {
		v1051 = v994
		v1052 = v992
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if int32(250) < v1051 {
		goto L247
	} else {
		goto L248
	}
L235:
	;
	v998 = v985 + int32(1)
	if v998 == v976 {
		v1051 = v994
		v1052 = v992
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998+v977))))
	v1003 = v1001 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v994) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007+v977))))
	v1019 = v1017 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v994) {
		goto L243
	} else {
		goto L244
	}
L238:
	;
	v1007 = v985 + int32(2)
	if v1007 != v976 {
		goto L237
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1051 = v994<<(uint(int32(6))%32)&int32(1984) | v1003
	v1052 = int32(2)
	goto L234
L241:
	;
	goto L240
L242:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977+v1023))))
	v1051 = v1036&int32(63) | (v994<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1003<<(uint(int32(12))%32) | v1019<<(uint(int32(6))%32))
	v1052 = int32(4)
	goto L234
L243:
	;
	v1023 = v985 + int32(3)
	if v1023 != v976 {
		goto L242
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1051 = v994<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1003<<(uint(int32(6))%32) | v1019
	v1052 = int32(3)
	goto L234
L246:
	;
	goto L245
L247:
	;
	v1069 = v1052 + v985
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1069
	v985 = v1069
	goto L229
L248:
	;
	v1056 = v1051 - int32(97)
	if v1056 < int32(0) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1056)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1062)>>(uint(v1056&int32(7))%32))&int32(1) != 0 {
		goto L228
	} else {
		goto L250
	}
L250:
	;
	goto L247
L252:
	;
	goto L226
L253:
	;
	if v1202 != 0 {
		goto L40
	} else {
		goto L277
	}
L254:
	;
	v1202 = v1195
	goto L253
L255:
	;
	if v1097 <= v844 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1195 = int32(0)
	goto L254
L257:
	;
	v1202 = int32(-1)
	goto L253
L258:
	;
	goto L259
L259:
	;
	v1113 = int32(1)
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844+v1098))))
	if base.Ui32(v1115) < base.Ui32(int32(192)) {
		v1172 = v1115
		v1173 = v1113
		goto L260
	} else {
		goto L261
	}
L260:
	;
	if int32(250) < v1172 {
		v1195 = v1173
		goto L254
	} else {
		goto L273
	}
L261:
	;
	v1119 = v844 + int32(1)
	if v1119 == v1097 {
		v1172 = v1115
		v1173 = v1113
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119+v1098))))
	v1124 = v1122 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1115) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1128+v1098))))
	v1140 = v1138 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1115) {
		goto L269
	} else {
		goto L270
	}
L264:
	;
	v1128 = v844 + int32(2)
	if v1128 != v1097 {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1172 = v1115<<(uint(int32(6))%32)&int32(1984) | v1124
	v1173 = int32(2)
	goto L260
L267:
	;
	goto L266
L268:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1098+v1144))))
	v1172 = v1157&int32(63) | (v1115<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1124<<(uint(int32(12))%32) | v1140<<(uint(int32(6))%32))
	v1173 = int32(4)
	goto L260
L269:
	;
	v1144 = v844 + int32(3)
	if v1144 != v1097 {
		goto L268
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1172 = v1115<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1124<<(uint(int32(6))%32) | v1140
	v1173 = int32(3)
	goto L260
L272:
	;
	goto L271
L273:
	;
	v1177 = v1172 - int32(97)
	if v1177 < int32(0) {
		v1195 = v1173
		goto L254
	} else {
		goto L274
	}
L274:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1177)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1183)>>(uint(v1177&int32(7))%32))&int32(1) == int32(0) {
		v1195 = v1173
		goto L254
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1173 + v844
	goto L276
L276:
	;
	goto L256
L277:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L280
L278:
	;
	if int32(0) <= v1257 {
		v1265 = v1257
		goto L41
	} else {
		goto L298
	}
L280:
	;
	goto L281
L281:
	;
	goto L282
L282:
	;
	v1212 = v1204
	v1214 = int32(1)
	goto L285
L284:
	;
	v1257 = v1242
	goto L278
L285:
	;
	if v1205 <= v1212 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	goto L284
L287:
	;
	v1257 = int32(-1)
	goto L278
L288:
	;
	goto L289
L289:
	;
	v1219 = v1212 + int32(1)
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203+v1212))))
	if base.Ui32(v1221) < base.Ui32(int32(192)) {
		v1242 = v1219
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1243 = int32(1)
	if v1243 < v1214 {
		v1212 = v1242
		v1214 = v1214 - v1243
		goto L285
	} else {
		goto L297
	}
L291:
	;
	if v1205 <= v1219 {
		v1242 = v1219
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1228 = v1219
	goto L293
L293:
	;
	v1231 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1203+v1228))))
	if int32(-65) < v1231 {
		v1242 = v1228
		goto L290
	} else {
		goto L295
	}
L294:
	;
	v1242 = v1205
	goto L290
L295:
	;
	v1235 = v1228 + int32(1)
	if v1235 != v1205 {
		v1228 = v1235
		goto L293
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	goto L286
L298:
	;
	goto L40
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1766
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1766
	if v1766-int32(2) <= v126 {
		goto L404
	} else {
		goto L405
	}
L300:
	;
	if v1387 < int32(0) {
		goto L299
	} else {
		goto L325
	}
L301:
	;
	v1387 = v1359
	goto L300
L302:
	;
	if v1283 <= v1292 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1387 = int32(-1)
	goto L300
L305:
	;
	goto L306
L306:
	;
	v1299 = int32(1)
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292+v1284))))
	if base.Ui32(v1301) < base.Ui32(int32(192)) {
		v1358 = v1301
		v1359 = v1299
		goto L307
	} else {
		goto L308
	}
L307:
	;
	if int32(250) < v1358 {
		goto L320
	} else {
		goto L321
	}
L308:
	;
	v1305 = v1292 + int32(1)
	if v1305 == v1283 {
		v1358 = v1301
		v1359 = v1299
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v1284))))
	v1310 = v1308 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1301) {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314+v1284))))
	v1326 = v1324 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1301) {
		goto L316
	} else {
		goto L317
	}
L311:
	;
	v1314 = v1292 + int32(2)
	if v1314 != v1283 {
		goto L310
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1358 = v1301<<(uint(int32(6))%32)&int32(1984) | v1310
	v1359 = int32(2)
	goto L307
L314:
	;
	goto L313
L315:
	;
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1284+v1330))))
	v1358 = v1343&int32(63) | (v1301<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1310<<(uint(int32(12))%32) | v1326<<(uint(int32(6))%32))
	v1359 = int32(4)
	goto L307
L316:
	;
	v1330 = v1292 + int32(3)
	if v1330 != v1283 {
		goto L315
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v1358 = v1301<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1310<<(uint(int32(6))%32) | v1326
	v1359 = int32(3)
	goto L307
L319:
	;
	goto L318
L320:
	;
	v1376 = v1359 + v1292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1376
	v1292 = v1376
	goto L302
L321:
	;
	v1363 = v1358 - int32(97)
	if v1363 < int32(0) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1363)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1369)>>(uint(v1363&int32(7))%32))&int32(1) != 0 {
		goto L301
	} else {
		goto L323
	}
L323:
	;
	goto L320
L325:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1391 = v1390 + v1387
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1391
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1414 = v1391
	goto L328
L326:
	;
	if v1510 < int32(0) {
		goto L299
	} else {
		goto L350
	}
L327:
	;
	v1510 = v1481
	goto L326
L328:
	;
	if v1405 <= v1414 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1510 = int32(-1)
	goto L326
L331:
	;
	goto L332
L332:
	;
	v1421 = int32(1)
	v1423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1414+v1406))))
	if base.Ui32(v1423) < base.Ui32(int32(192)) {
		v1480 = v1423
		v1481 = v1421
		goto L333
	} else {
		goto L334
	}
L333:
	;
	if int32(250) < v1480 {
		goto L327
	} else {
		goto L346
	}
L334:
	;
	v1427 = v1414 + int32(1)
	if v1427 == v1405 {
		v1480 = v1423
		v1481 = v1421
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427+v1406))))
	v1432 = v1430 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1423) {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436+v1406))))
	v1448 = v1446 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1423) {
		goto L342
	} else {
		goto L343
	}
L337:
	;
	v1436 = v1414 + int32(2)
	if v1436 != v1405 {
		goto L336
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1480 = v1423<<(uint(int32(6))%32)&int32(1984) | v1432
	v1481 = int32(2)
	goto L333
L340:
	;
	goto L339
L341:
	;
	v1465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406+v1452))))
	v1480 = v1465&int32(63) | (v1423<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1432<<(uint(int32(12))%32) | v1448<<(uint(int32(6))%32))
	v1481 = int32(4)
	goto L333
L342:
	;
	v1452 = v1414 + int32(3)
	if v1452 != v1405 {
		goto L341
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1480 = v1423<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1432<<(uint(int32(6))%32) | v1448
	v1481 = int32(3)
	goto L333
L345:
	;
	goto L344
L346:
	;
	v1485 = v1480 - int32(97)
	if v1485 < int32(0) {
		goto L327
	} else {
		goto L347
	}
L347:
	;
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1485)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1491)>>(uint(v1485&int32(7))%32))&int32(1) == int32(0) {
		goto L327
	} else {
		goto L348
	}
L348:
	;
	v1499 = v1481 + v1414
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1499
	v1414 = v1499
	goto L328
L350:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1514 = v1513 + v1510
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1514
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+4)) = v1514
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1539 = v1529
	goto L353
L351:
	;
	if v1634 < int32(0) {
		goto L299
	} else {
		goto L376
	}
L352:
	;
	v1634 = v1606
	goto L351
L353:
	;
	if v1530 <= v1539 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1634 = int32(-1)
	goto L351
L356:
	;
	goto L357
L357:
	;
	v1546 = int32(1)
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1539+v1531))))
	if base.Ui32(v1548) < base.Ui32(int32(192)) {
		v1605 = v1548
		v1606 = v1546
		goto L358
	} else {
		goto L359
	}
L358:
	;
	if int32(250) < v1605 {
		goto L371
	} else {
		goto L372
	}
L359:
	;
	v1552 = v1539 + int32(1)
	if v1552 == v1530 {
		v1605 = v1548
		v1606 = v1546
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552+v1531))))
	v1557 = v1555 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1548) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1561+v1531))))
	v1573 = v1571 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1548) {
		goto L367
	} else {
		goto L368
	}
L362:
	;
	v1561 = v1539 + int32(2)
	if v1561 != v1530 {
		goto L361
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1605 = v1548<<(uint(int32(6))%32)&int32(1984) | v1557
	v1606 = int32(2)
	goto L358
L365:
	;
	goto L364
L366:
	;
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1531+v1577))))
	v1605 = v1590&int32(63) | (v1548<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1557<<(uint(int32(12))%32) | v1573<<(uint(int32(6))%32))
	v1606 = int32(4)
	goto L358
L367:
	;
	v1577 = v1539 + int32(3)
	if v1577 != v1530 {
		goto L366
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1605 = v1548<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1557<<(uint(int32(6))%32) | v1573
	v1606 = int32(3)
	goto L358
L370:
	;
	goto L369
L371:
	;
	v1623 = v1606 + v1539
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1623
	v1539 = v1623
	goto L353
L372:
	;
	v1610 = v1605 - int32(97)
	if v1610 < int32(0) {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1610)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1616)>>(uint(v1610&int32(7))%32))&int32(1) != 0 {
		goto L352
	} else {
		goto L374
	}
L374:
	;
	goto L371
L376:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1638 = v1637 + v1634
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1638
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1661 = v1638
	goto L379
L377:
	;
	if v1757 < int32(0) {
		goto L299
	} else {
		goto L401
	}
L378:
	;
	v1757 = v1728
	goto L377
L379:
	;
	if v1652 <= v1661 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1757 = int32(-1)
	goto L377
L382:
	;
	goto L383
L383:
	;
	v1668 = int32(1)
	v1670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661+v1653))))
	if base.Ui32(v1670) < base.Ui32(int32(192)) {
		v1727 = v1670
		v1728 = v1668
		goto L384
	} else {
		goto L385
	}
L384:
	;
	if int32(250) < v1727 {
		goto L378
	} else {
		goto L397
	}
L385:
	;
	v1674 = v1661 + int32(1)
	if v1674 == v1652 {
		v1727 = v1670
		v1728 = v1668
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1674+v1653))))
	v1679 = v1677 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1670) {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	v1693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1683+v1653))))
	v1695 = v1693 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1670) {
		goto L393
	} else {
		goto L394
	}
L388:
	;
	v1683 = v1661 + int32(2)
	if v1683 != v1652 {
		goto L387
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	v1727 = v1670<<(uint(int32(6))%32)&int32(1984) | v1679
	v1728 = int32(2)
	goto L384
L391:
	;
	goto L390
L392:
	;
	v1712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653+v1699))))
	v1727 = v1712&int32(63) | (v1670<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1679<<(uint(int32(12))%32) | v1695<<(uint(int32(6))%32))
	v1728 = int32(4)
	goto L384
L393:
	;
	v1699 = v1661 + int32(3)
	if v1699 != v1652 {
		goto L392
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1727 = v1670<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1679<<(uint(int32(6))%32) | v1695
	v1728 = int32(3)
	goto L384
L396:
	;
	goto L395
L397:
	;
	v1732 = v1727 - int32(97)
	if v1732 < int32(0) {
		goto L378
	} else {
		goto L398
	}
L398:
	;
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1732)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1738)>>(uint(v1732&int32(7))%32))&int32(1) == int32(0) {
		goto L378
	} else {
		goto L399
	}
L399:
	;
	v1746 = v1728 + v1661
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1746
	v1661 = v1746
	goto L379
L401:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1760))) = v1761 + v1757
	goto L299
L402:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2117
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2117
	v2122 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_5), int32(4))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L10
	} else {
		goto L506
	}
L403:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2084
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2084
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2084 <= v2087 {
		goto L402
	} else {
		goto L498
	}
L404:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2041
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v2043)+8))
	if v2044 <= v2041 {
		goto L486
	} else {
		goto L487
	}
L405:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1774 = int32(1)
	v1776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1772+v1766-v1774))))
	if base.B2i32(v1776&int32(224) != int32(96))|base.B2i32(v1774<<(uint(v1776)%32)&int32(_a_F_portuguese_UTF_8_stem_6) == int32(0)) != 0 {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1790 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_7), int32(45))
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L10
	} else {
		goto L407
	}
L407:
	;
	if v1790 == int32(0) {
		goto L404
	} else {
		goto L408
	}
L408:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1794
	switch v1790 - int32(1) {
	case 0:
		goto L417
	case 1:
		goto L416
	case 2:
		goto L415
	case 3:
		goto L414
	case 4:
		goto L413
	case 5:
		goto L412
	case 6:
		goto L411
	case 7:
		goto L410
	case 8:
		goto L409
	default:
		goto L403
	}
L409:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+8))
	if v1794 < v2019 {
		goto L404
	} else {
		goto L480
	}
L410:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1982)))
	if v1794 < v1983 {
		goto L404
	} else {
		goto L469
	}
L411:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1940)))
	if v1794 < v1941 {
		goto L404
	} else {
		goto L459
	}
L412:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1905)))
	if v1794 < v1906 {
		goto L404
	} else {
		goto L449
	}
L413:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1832)+4))
	if v1794 < v1833 {
		goto L404
	} else {
		goto L430
	}
L414:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1823)))
	if v1794 < v1824 {
		goto L404
	} else {
		goto L427
	}
L415:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1814)))
	if v1794 < v1815 {
		goto L404
	} else {
		goto L424
	}
L416:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1805)))
	if v1794 < v1806 {
		goto L404
	} else {
		goto L421
	}
L417:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1798)))
	if v1794 < v1799 {
		goto L404
	} else {
		goto L418
	}
L418:
	;
	v1801 = F_slice_del(m, l0)
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L10
	} else {
		goto L419
	}
L419:
	;
	if int32(0) <= v1801 {
		goto L403
	} else {
		goto L420
	}
L420:
	;
	v2300 = v1801
	goto L1
L421:
	;
	v1810 = F_slice_from_s(m, l0, int32(3), int32(_a_F_portuguese_UTF_8_stem_8))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L10
	} else {
		goto L422
	}
L422:
	;
	if int32(0) <= v1810 {
		goto L403
	} else {
		goto L423
	}
L423:
	;
	v2300 = v1810
	goto L1
L424:
	;
	v1819 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_UTF_8_stem_9))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L10
	} else {
		goto L425
	}
L425:
	;
	if int32(0) <= v1819 {
		goto L403
	} else {
		goto L426
	}
L426:
	;
	v2300 = v1819
	goto L1
L427:
	;
	v1828 = F_slice_from_s(m, l0, int32(4), int32(_a_F_portuguese_UTF_8_stem_10))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L10
	} else {
		goto L428
	}
L428:
	;
	if int32(0) <= v1828 {
		goto L403
	} else {
		goto L429
	}
L429:
	;
	v2300 = v1828
	goto L1
L430:
	;
	v1835 = F_slice_del(m, l0)
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L10
	} else {
		goto L431
	}
L431:
	;
	if v1835 < int32(0) {
		v2300 = v1835
		goto L1
	} else {
		goto L432
	}
L432:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1839
	v1842 = v1839 - int32(1)
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1842 <= v1843 {
		goto L403
	} else {
		goto L433
	}
L433:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1845+v1842))))
	if base.B2i32(v1847&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1847)%32)&int32(_a_F_portuguese_UTF_8_stem_11) == int32(0)) != 0 {
		goto L403
	} else {
		goto L434
	}
L434:
	;
	v1861 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_12), int32(4))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L10
	} else {
		goto L435
	}
L435:
	;
	if v1861 == int32(0) {
		goto L403
	} else {
		goto L436
	}
L436:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1865
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v1867)))
	if v1865 < v1868 {
		goto L403
	} else {
		goto L437
	}
L437:
	;
	v1870 = F_slice_del(m, l0)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L10
	} else {
		goto L438
	}
L438:
	;
	if v1870 < int32(0) {
		v2300 = v1870
		goto L1
	} else {
		goto L439
	}
L439:
	;
	if v1861 != int32(1) {
		goto L403
	} else {
		goto L440
	}
L440:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1876
	v1878 = int32(2)
	v1880 = int32(0)
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1876-v1883 < v1878 {
		v1893 = v1880
		goto L442
	} else {
		goto L443
	}
L441:
	;
	if v1893 == int32(0) {
		goto L403
	} else {
		goto L445
	}
L442:
	;
	goto L441
L443:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1889 = F_memcmp(m, v1886+v1876-v1878, int32(_a_F_portuguese_UTF_8_stem_13), v1878)
	mBase = m.M
	if v1889 != 0 {
		v1893 = v1880
		goto L442
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1876 - v1878
	v1893 = int32(1)
	goto L442
L445:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1896
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1898)))
	if v1896 < v1899 {
		goto L403
	} else {
		goto L446
	}
L446:
	;
	v1901 = F_slice_del(m, l0)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L10
	} else {
		goto L447
	}
L447:
	;
	if int32(0) <= v1901 {
		goto L403
	} else {
		goto L448
	}
L448:
	;
	v2300 = v1901
	goto L1
L449:
	;
	v1908 = F_slice_del(m, l0)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L10
	} else {
		goto L450
	}
L450:
	;
	if v1908 < int32(0) {
		v2300 = v1908
		goto L1
	} else {
		goto L451
	}
L451:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1912
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1912-int32(3) <= v1914 {
		goto L403
	} else {
		goto L452
	}
L452:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1918+v1912-int32(1)))))
	switch v1922 - int32(101) {
	case 0, 7:
		goto L453
	default:
		goto L403
	}
L453:
	;
	v1927 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_14), int32(3))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L10
	} else {
		goto L454
	}
L454:
	;
	if v1927 == int32(0) {
		goto L403
	} else {
		goto L455
	}
L455:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1931
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1933)))
	if v1931 < v1934 {
		goto L403
	} else {
		goto L456
	}
L456:
	;
	v1936 = F_slice_del(m, l0)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L10
	} else {
		goto L457
	}
L457:
	;
	if int32(0) <= v1936 {
		goto L403
	} else {
		goto L458
	}
L458:
	;
	v2300 = v1936
	goto L1
L459:
	;
	v1943 = F_slice_del(m, l0)
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L10
	} else {
		goto L460
	}
L460:
	;
	if v1943 < int32(0) {
		v2300 = v1943
		goto L1
	} else {
		goto L461
	}
L461:
	;
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1947
	v1950 = v1947 - int32(1)
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1950 <= v1951 {
		goto L403
	} else {
		goto L462
	}
L462:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1953+v1950))))
	if base.B2i32(v1955&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1955)%32)&int32(_a_F_portuguese_UTF_8_stem_15) == int32(0)) != 0 {
		goto L403
	} else {
		goto L463
	}
L463:
	;
	v1969 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_16), int32(3))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L10
	} else {
		goto L464
	}
L464:
	;
	if v1969 == int32(0) {
		goto L403
	} else {
		goto L465
	}
L465:
	;
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1973
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1975)))
	if v1973 < v1976 {
		goto L403
	} else {
		goto L466
	}
L466:
	;
	v1978 = F_slice_del(m, l0)
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L10
	} else {
		goto L467
	}
L467:
	;
	if int32(0) <= v1978 {
		goto L403
	} else {
		goto L468
	}
L468:
	;
	v2300 = v1978
	goto L1
L469:
	;
	v1985 = F_slice_del(m, l0)
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L10
	} else {
		goto L470
	}
L470:
	;
	if v1985 < int32(0) {
		v2300 = v1985
		goto L1
	} else {
		goto L471
	}
L471:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1989
	v1991 = int32(2)
	v1993 = int32(0)
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1989-v1996 < v1991 {
		v2006 = v1993
		goto L473
	} else {
		goto L474
	}
L472:
	;
	if v2006 == int32(0) {
		goto L403
	} else {
		goto L476
	}
L473:
	;
	goto L472
L474:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2002 = F_memcmp(m, v1999+v1989-v1991, int32(_a_F_portuguese_UTF_8_stem_17), v1991)
	mBase = m.M
	if v2002 != 0 {
		v2006 = v1993
		goto L473
	} else {
		goto L475
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1989 - v1991
	v2006 = int32(1)
	goto L473
L476:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2009
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v2011)))
	if v2009 < v2012 {
		goto L403
	} else {
		goto L477
	}
L477:
	;
	v2014 = F_slice_del(m, l0)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L10
	} else {
		goto L478
	}
L478:
	;
	if int32(0) <= v2014 {
		goto L403
	} else {
		goto L479
	}
L479:
	;
	v2300 = v2014
	goto L1
L480:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1794 <= v2021 {
		goto L404
	} else {
		goto L481
	}
L481:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2023+v1794-int32(1)))))
	if v2027 != int32(101) {
		goto L404
	} else {
		goto L482
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1794 - int32(1)
	v2035 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_18))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L10
	} else {
		goto L483
	}
L483:
	;
	if int32(0) <= v2035 {
		goto L403
	} else {
		goto L484
	}
L484:
	;
	v2300 = v2035
	goto L1
L485:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2074
	v2076 = F_slice_del(m, l0)
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L10
	} else {
		goto L496
	}
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2041
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2044
	v2051 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_19), int32(120))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L10
	} else {
		goto L489
	}
L487:
	;
	v2055 = v2041
	goto L488
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2055
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2055
	v2061 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_20), int32(7))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L10
	} else {
		goto L491
	}
L489:
	;
	if v2051 != 0 {
		goto L485
	} else {
		goto L490
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2047
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2055 = v2054
	goto L488
L491:
	;
	if v2061 == int32(0) {
		goto L402
	} else {
		goto L492
	}
L492:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2065
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+8))
	if v2065 < v2068 {
		goto L402
	} else {
		goto L493
	}
L493:
	;
	v2070 = F_slice_del(m, l0)
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L10
	} else {
		goto L494
	}
L494:
	;
	if int32(0) <= v2070 {
		goto L402
	} else {
		goto L495
	}
L495:
	;
	v2300 = v2070
	goto L1
L496:
	;
	if v2076 < int32(0) {
		v2300 = v2076
		goto L1
	} else {
		goto L497
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2047
	goto L403
L498:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2090 = v2089 + v2084
	v2093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2090-int32(1)))))
	if v2093 != int32(105) {
		goto L402
	} else {
		goto L499
	}
L499:
	;
	v2097 = v2084 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2097
	if v2097 <= v2087 {
		goto L402
	} else {
		goto L500
	}
L500:
	;
	v2103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2090-int32(2)))))
	if v2103 != int32(99) {
		goto L402
	} else {
		goto L501
	}
L501:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2106)+8))
	if v2084 <= v2107 {
		goto L402
	} else {
		goto L502
	}
L502:
	;
	v2109 = F_slice_del(m, l0)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L10
	} else {
		goto L503
	}
L503:
	;
	if v2109 < int32(0) {
		v2300 = v2109
		goto L1
	} else {
		goto L504
	}
L504:
	;
	goto L402
L505:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2193
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2197 = v2193
	v2199 = v2195
	goto L527
L506:
	;
	if v2122 == int32(0) {
		goto L505
	} else {
		goto L507
	}
L507:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2126
	switch v2122 - int32(1) {
	case 0:
		goto L509
	case 1:
		goto L508
	default:
		goto L505
	}
L508:
	;
	v2184 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_UTF_8_stem_21))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L10
	} else {
		goto L525
	}
L509:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2130)+8))
	if v2126 < v2131 {
		goto L505
	} else {
		goto L510
	}
L510:
	;
	v2133 = F_slice_del(m, l0)
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L10
	} else {
		goto L511
	}
L511:
	;
	if v2133 < int32(0) {
		v2300 = v2133
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2137
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2137 <= v2139 {
		goto L505
	} else {
		goto L513
	}
L513:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2142 = v2141 + v2137
	v2144 = v2142 - int32(1)
	v2145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2144))))
	if v2145 != int32(117) {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2173
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v2175)+8))
	if v2173 < v2176 {
		goto L505
	} else {
		goto L522
	}
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2137
	v2160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2144))))
	if v2160 != int32(105) {
		goto L505
	} else {
		goto L519
	}
L516:
	;
	v2149 = v2137 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2149
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2149
	if v2149 <= v2139 {
		goto L515
	} else {
		goto L517
	}
L517:
	;
	v2155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2142-int32(2)))))
	if v2155 == int32(103) {
		v2173 = v2149
		goto L514
	} else {
		goto L518
	}
L518:
	;
	goto L515
L519:
	;
	v2164 = v2137 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2164
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2164
	if v2164 <= v2139 {
		goto L505
	} else {
		goto L520
	}
L520:
	;
	v2170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2142-int32(2)))))
	if v2170 != int32(99) {
		goto L505
	} else {
		goto L521
	}
L521:
	;
	v2173 = v2164
	goto L514
L522:
	;
	v2178 = F_slice_del(m, l0)
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L10
	} else {
		goto L523
	}
L523:
	;
	if int32(0) <= v2178 {
		goto L505
	} else {
		goto L524
	}
L524:
	;
	v2300 = v2178
	goto L1
L525:
	;
	if v2184 < int32(0) {
		v2300 = v2184
		goto L1
	} else {
		goto L526
	}
L526:
	;
	goto L505
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2197
	v2204 = v2197 + int32(1)
	if v2204 < v2199 {
		goto L533
	} else {
		goto L534
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2193
	v2300 = int32(1)
	goto L1
L529:
	;
	goto L528
L530:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2197 = v2296
	v2199 = v2295
	goto L527
L531:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L547
L532:
	;
	v2214 = F_find_among(m, l0, int32(_a_F_portuguese_UTF_8_stem_22), int32(3))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L10
	} else {
		goto L537
	}
L533:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2206+v2204))))
	if v2208 == int32(126) {
		goto L532
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2197
	v2233 = v2197
	v2234 = v2199
	goto L531
L536:
	;
	goto L535
L537:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2216
	switch v2214 - int32(1) {
	case 0:
		goto L540
	case 1:
		goto L539
	case 2:
		goto L538
	default:
		goto L530
	}
L538:
	;
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2233 = v2216
	v2234 = v2232
	goto L531
L539:
	;
	v2228 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_23))
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L10
	} else {
		goto L543
	}
L540:
	;
	v2222 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_24))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L10
	} else {
		goto L541
	}
L541:
	;
	if int32(0) <= v2222 {
		goto L530
	} else {
		goto L542
	}
L542:
	;
	v2300 = v2222
	goto L1
L543:
	;
	if int32(0) <= v2228 {
		goto L530
	} else {
		goto L544
	}
L544:
	;
	v2300 = v2228
	goto L1
L545:
	;
	if v2288 < int32(0) {
		goto L529
	} else {
		goto L565
	}
L547:
	;
	goto L548
L548:
	;
	goto L549
L549:
	;
	v2243 = v2233
	v2245 = int32(1)
	goto L552
L551:
	;
	v2288 = v2273
	goto L545
L552:
	;
	if v2234 <= v2243 {
		goto L554
	} else {
		goto L555
	}
L553:
	;
	goto L551
L554:
	;
	v2288 = int32(-1)
	goto L545
L555:
	;
	goto L556
L556:
	;
	v2250 = v2243 + int32(1)
	v2252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236+v2243))))
	if base.Ui32(v2252) < base.Ui32(int32(192)) {
		v2273 = v2250
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v2274 = int32(1)
	if v2274 < v2245 {
		v2243 = v2273
		v2245 = v2245 - v2274
		goto L552
	} else {
		goto L564
	}
L558:
	;
	if v2234 <= v2250 {
		v2273 = v2250
		goto L557
	} else {
		goto L559
	}
L559:
	;
	v2259 = v2250
	goto L560
L560:
	;
	v2262 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2236+v2259))))
	if int32(-65) < v2262 {
		v2273 = v2259
		goto L557
	} else {
		goto L562
	}
L561:
	;
	v2273 = v2234
	goto L557
L562:
	;
	v2266 = v2259 + int32(1)
	if v2266 != v2234 {
		v2259 = v2266
		goto L560
	} else {
		goto L563
	}
L563:
	;
	goto L561
L564:
	;
	goto L553
L565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2288
	goto L530
}
func F_posix_fadvise(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	var v6 int32
	_ = v6
	v6 = m.Env.X__syscall_fadvise64(m, l0, l1, l2, l3)
	return int32(0) - v6
}
func F_preadv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = m.Wasi_snapshot_preview1.Fd_pread(m, l0, l1, l2, l3, v8+int32(12))
	mBase = m.M
	if v12 == int32(0) {
		v19 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_preadv[0])) = v12
		v19 = int32(-1)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	m.G0 = v8 + int32(16)
	if v19 != 0 {
		v25 = int32(-1)
	} else {
		v25 = v20
	}
	return v25
}
func F_prefixsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13990(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_preprocess_pubobj_list(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L52
	}
L2:
	;
	return
L3:
	;
	v9 = int32(3)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == v9 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v21 = v9
	v23 = v3
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == int32(3) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v21
	v33 = v21
	goto L10
L9:
	;
	v33 = v29
	goto L10
L10:
	;
	switch v33 {
	case 0:
		goto L17
	case 1, 2:
		goto L16
	default:
		v141 = v33
		goto L11
	}
L11:
	;
	v143 = v23 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v143 < v144 {
		v21 = v141
		v23 = v143
		goto L6
	} else {
		goto L51
	}
L12:
	;
	v138 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v138
	v141 = v138
	goto L11
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L46
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L41
	}
L15:
	;
	v86 = F_palloc0(m, int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L20
	} else {
		goto L39
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v56 != 0 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v34 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v36 != 0 {
		v141 = int32(0)
		goto L11
	} else {
		goto L19
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_preprocess_pubobj_list_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v48, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_preprocess_pubobj_list_1), int32(_a_F_preprocess_pubobj_list_2), int32(_a_F_preprocess_pubobj_list_3))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L20
	} else {
		goto L34
	}
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v57 != 0 {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v62 != 0 {
		goto L12
	} else {
		goto L33
	}
L30:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v58 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v59 == int32(0) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	goto L12
L33:
	;
	v63 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v63
	v141 = v63
	goto L11
L34:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_preprocess_pubobj_list_4), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v77, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_preprocess_pubobj_list_1), int32(_a_F_preprocess_pubobj_list_5), int32(_a_F_preprocess_pubobj_list_3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(259)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v93 = F_makeRangeVar(m, int32(0), v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v86
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v141 = v99
	goto L11
L41:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_preprocess_pubobj_list_6), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v111, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_preprocess_pubobj_list_1), int32(_a_F_preprocess_pubobj_list_7), int32(_a_F_preprocess_pubobj_list_3))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L20
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_preprocess_pubobj_list_8), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v130, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_preprocess_pubobj_list_1), int32(_a_F_preprocess_pubobj_list_9), int32(_a_F_preprocess_pubobj_list_3))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
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
	goto L7
L52:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_preprocess_pubobj_list_10), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	F_errdetail(m, int32(_a_F_preprocess_pubobj_list_11), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_scanner_errposition(m, v167, l1)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_preprocess_pubobj_list_1), int32(_a_F_preprocess_pubobj_list_12), int32(_a_F_preprocess_pubobj_list_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_process_postgres_switches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	if l2 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[0])) = int32(0)
	goto L21
L2:
	;
	v53 = l0
	v54 = l1
	v55 = int32(9)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = int32(4)
	if l0 < int32(2) {
		v53 = l0
		v54 = l1
		v55 = v16
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v22 = int32(_a_F_process_postgres_switches_0)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[1])))
	if base.B2i32(v25 == int32(0))|base.B2i32(v25 != v28) != 0 {
		v46 = v25
		v47 = v28
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v48 = v46 - v47
	goto L6
L8:
	;
	v31 = v21
	v32 = v22
	goto L9
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v36 == int32(0) {
		v46 = v36
		v47 = v35
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v46 = v36
	v47 = v35
	goto L7
L11:
	;
	v39 = int32(1)
	if v36 == v35 {
		v31 = v31 + v39
		v32 = v32 + v39
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v49 = l1
	goto L15
L14:
	;
	v49 = l1 + int32(4)
	goto L15
L15:
	;
	v53 = l0 - base.B2i32(v48 == int32(0))
	v54 = v49
	v55 = v16
	goto L1
L16:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2]))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v54+v633<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v637
	F_errmsg(m, int32(_a_F_process_postgres_switches_1), v11+int32(16))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L49
	} else {
		goto L195
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v133
	F_errmsg(m, int32(_a_F_process_postgres_switches_2), v11-int32(-64))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L49
	} else {
		goto L193
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L49
	} else {
		goto L189
	}
L19:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[3])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L49
	} else {
		goto L183
	}
L20:
	;
	v553 = int32(_a_F_process_postgres_switches_3)
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2])) = v555 - int32(1)
	goto L19
L21:
	;
	v69 = F_getopt(m, v53, v54, int32(_a_F_process_postgres_switches_4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L49
	} else {
		goto L50
	}
L22:
	;
	if l3 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L23:
	;
	goto L22
L24:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_5), v507, l2, v55)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L49
	} else {
		goto L173
	}
L25:
	;
	if l2 != int32(1) {
		goto L21
	} else {
		goto L156
	}
L26:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	switch v439 - int32(101) {
	case 0:
		v447 = int32(_a_F_process_postgres_switches_6)
		goto L152
	default:
		goto L20
	case 11:
		goto L153
	}
L27:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_7), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L49
	} else {
		goto L151
	}
L28:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_9), v429, l2, v55)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L49
	} else {
		goto L150
	}
L29:
	;
	if l2 != int32(1) {
		goto L21
	} else {
		goto L118
	}
L30:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_10), v300, l2, v55)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L49
	} else {
		goto L117
	}
L31:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_11), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L49
	} else {
		goto L116
	}
L32:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_12), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L49
	} else {
		goto L115
	}
L33:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_13), v287, l2, v55)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L49
	} else {
		goto L114
	}
L34:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_14), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L49
	} else {
		goto L113
	}
L35:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_15), v278, l2, v55)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L49
	} else {
		goto L112
	}
L36:
	;
	if l2 != int32(1) {
		goto L21
	} else {
		goto L111
	}
L37:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_16), int32(_a_F_process_postgres_switches_17), l2, v55)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L49
	} else {
		goto L110
	}
L38:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_16), v264, l2, v55)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L49
	} else {
		goto L109
	}
L39:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v242 = v240 - int32(98)
	v244 = v242 & int32(255)
	if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v244))|base.B2i32(int32(base.Ui32(int32(_a_F_process_postgres_switches_18))>>(uint(v244)%32))&int32(1) == int32(0)) != 0 {
		goto L20
	} else {
		goto L107
	}
L40:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_19), int32(_a_F_process_postgres_switches_20), l2, v55)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L49
	} else {
		goto L106
	}
L41:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_21), int32(_a_F_process_postgres_switches_22), l2, v55)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L49
	} else {
		goto L105
	}
L42:
	;
	if l2 != int32(1) {
		goto L21
	} else {
		goto L104
	}
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	v178 = v174
	goto L88
L44:
	;
	if l2 != int32(1) {
		goto L21
	} else {
		goto L82
	}
L45:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_ParseLongOption(m, v115, v11+int32(108), v11+int32(104))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L49
	} else {
		goto L70
	}
L46:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	v86 = F_strcmp(m, int32(_a_F_process_postgres_switches_23), v84)
	mBase = m.M
	if v86 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	if l2 != int32(1) {
		goto L21
	} else {
		goto L52
	}
L48:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_24), v75, l2, v55)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L49
	} else {
		goto L51
	}
L49:
	;
	return
L50:
	;
	switch v69 + int32(1) {
	case 0:
		goto L23
	default:
		goto L20
	case 46:
		goto L46
	case 67:
		goto L48
	case 68, 85, 111:
		goto L21
	case 69:
		goto L44
	case 70:
		goto L42
	case 71:
		goto L40
	case 79:
		goto L33
	case 80:
		goto L32
	case 81:
		goto L31
	case 84:
		goto L28
	case 88:
		goto L24
	case 99:
		goto L47
	case 100:
		goto L45
	case 101:
		goto L43
	case 102:
		goto L41
	case 103:
		goto L39
	case 105:
		goto L38
	case 106:
		goto L37
	case 107:
		goto L36
	case 108:
		goto L35
	case 109:
		goto L34
	case 113:
		goto L30
	case 115:
		goto L29
	case 116:
		goto L27
	case 117:
		goto L26
	case 119:
		goto L25
	}
L51:
	;
	goto L21
L52:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[5])) = uint8(v81)
	goto L21
L53:
	;
	if v111 != int32(5) {
		goto L18
	} else {
		goto L69
	}
L54:
	;
	v111 = int32(0)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v91 = F_strcmp(m, int32(_a_F_process_postgres_switches_25), v84)
	mBase = m.M
	if v91 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v111 = int32(1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v97 = F_strncmp(m, int32(_a_F_process_postgres_switches_26), v84, int32(9))
	mBase = m.M
	if v97 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v111 = int32(2)
	goto L53
L61:
	;
	goto L62
L62:
	;
	v102 = F_strcmp(m, int32(_a_F_process_postgres_switches_27), v84)
	mBase = m.M
	if v102 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v111 = int32(3)
	goto L53
L64:
	;
	goto L65
L65:
	;
	v109 = F_strcmp(m, int32(_a_F_process_postgres_switches_28), v84)
	mBase = m.M
	if v109 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v110 = int32(5)
	goto L68
L67:
	;
	v110 = int32(4)
	goto L68
L68:
	;
	v111 = v110
	goto L53
L69:
	;
	goto L45
L70:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	if v122 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L49
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	F_SetConfigOption(m, v147, v122, l2, v55)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L49
	} else {
		goto L79
	}
L74:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L49
	} else {
		goto L75
	}
L75:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	if v69 == int32(45) {
		goto L17
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v133
	F_errmsg(m, int32(_a_F_process_postgres_switches_29), v11+int32(80))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L49
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(3985), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L49
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	F_pfree(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L49
	} else {
		goto L80
	}
L80:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	F_pfree(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L49
	} else {
		goto L81
	}
L81:
	;
	goto L21
L82:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	v163 = F_strlen(m, v160)
	mBase = m.M
	v165 = v163 + int32(1)
	v166 = F_emscripten_builtin_malloc(m, v165)
	mBase = m.M
	if v166 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[6])) = v171
	goto L21
L84:
	;
	v171 = int32(0)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v170 = F___memcpy(m, v166, v160, v165)
	mBase = m.M
	v171 = v170
	goto L83
L87:
	;
	F_set_debug_options(m, v222, l2, v55)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L49
	} else {
		goto L103
	}
L88:
	;
	v183 = v178 + int32(1)
	v184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178))))
	v185 = F___isspace(m, v184)
	mBase = m.M
	if v185 != 0 {
		v178 = v183
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v186 = int32(1)
	switch v184&int32(255) - int32(43) {
	case 0:
		v192 = v186
		goto L92
	default:
		v194 = v184
		v195 = v178
		v196 = v186
		goto L91
	case 2:
		goto L93
	}
L90:
	;
	goto L89
L91:
	;
	v197 = int32(0)
	v199 = v194 - int32(48)
	if base.Ui32(v199) <= base.Ui32(int32(9)) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v183))))
	v194 = v193
	v195 = v183
	v196 = v192
	goto L91
L93:
	;
	v192 = int32(0)
	goto L92
L94:
	;
	v202 = v197
	v203 = v199
	v204 = v195
	goto L97
L95:
	;
	v216 = v197
	goto L96
L96:
	;
	if v196 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v206 = int32(10)
	v208 = v202*v206 - v203
	v209 = int32(*(*int8)(unsafe.Add(mBase, uint32(v204)+1)))
	v213 = v209 - int32(48)
	if base.Ui32(v213) < base.Ui32(v206) {
		v202 = v208
		v203 = v213
		v204 = v204 + int32(1)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v216 = v208
	goto L96
L99:
	;
	goto L98
L100:
	;
	v222 = int32(0) - v216
	goto L102
L101:
	;
	v222 = v216
	goto L102
L102:
	;
	goto L87
L103:
	;
	goto L21
L104:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[7])) = uint8(v228)
	goto L21
L105:
	;
	goto L21
L106:
	;
	goto L21
L107:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v242&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_process_postgres_switches[8])))
	F_SetConfigOption(m, v258, int32(_a_F_process_postgres_switches_20), l2, v55)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L49
	} else {
		goto L108
	}
L108:
	;
	goto L21
L109:
	;
	goto L21
L110:
	;
	goto L21
L111:
	;
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[9])) = uint8(v274)
	goto L21
L112:
	;
	goto L21
L113:
	;
	goto L21
L114:
	;
	goto L21
L115:
	;
	goto L21
L116:
	;
	goto L21
L117:
	;
	goto L21
L118:
	;
	v305 = int32(_a_F_process_postgres_switches_32)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	goto L122
L119:
	;
	goto L21
L120:
	;
	v424 = F_strlen(m, v413)
	mBase = m.M
	goto L119
L122:
	;
	goto L123
L123:
	;
	v314 = int32(1023)
	if (v305^v307)&int32(3) != 0 {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v414))) = uint8(v417)
	goto L120
L125:
	;
	v398 = v393
	v399 = v394
	v400 = v395
	goto L146
L126:
	;
	if v388 == int32(0) {
		v413 = v386
		v414 = v387
		goto L124
	} else {
		goto L145
	}
L127:
	;
	v386 = v307
	v387 = v305
	v388 = v314
	goto L126
L128:
	;
	goto L129
L129:
	;
	v318 = int32(0)
	if base.B2i32(v307&int32(3) == v318)|int32(0) == v318 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v354 == int32(0) {
		v413 = v351
		v414 = v352
		goto L124
	} else {
		goto L139
	}
L131:
	;
	v330 = v307
	v331 = v305
	v332 = v314
	goto L134
L132:
	;
	goto L133
L133:
	;
	v351 = v307
	v352 = v305
	v353 = v314
	v354 = int32(1)
	goto L130
L134:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v334)
	if v334 == int32(0) {
		v393 = v330
		v394 = v331
		v395 = v332
		goto L125
	} else {
		goto L136
	}
L135:
	;
	v351 = v345
	v352 = v339
	v353 = v341
	v354 = v343
	goto L130
L136:
	;
	v338 = int32(1)
	v339 = v331 + v338
	v341 = v332 - v338
	v342 = int32(0)
	v343 = base.B2i32(v341 != v342)
	v345 = v330 + v338
	if v345&int32(3) == v342 {
		v351 = v345
		v352 = v339
		v353 = v341
		v354 = v343
		goto L130
	} else {
		goto L137
	}
L137:
	;
	if v341 != 0 {
		v330 = v345
		v331 = v339
		v332 = v341
		goto L134
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	if base.B2i32(v357 == int32(0))|base.B2i32(base.Ui32(v353) < base.Ui32(int32(4))) != 0 {
		v386 = v351
		v387 = v352
		v388 = v353
		goto L126
	} else {
		goto L140
	}
L140:
	;
	v364 = v351
	v365 = v352
	v366 = v353
	goto L141
L141:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v372 = int32(-2139062144)
	if (int32(16843008)-v369|v369)&v372 != v372 {
		v393 = v364
		v394 = v365
		v395 = v366
		goto L125
	} else {
		goto L143
	}
L142:
	;
	v386 = v380
	v387 = v378
	v388 = v382
	goto L126
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v369
	v377 = int32(4)
	v378 = v365 + v377
	v380 = v364 + v377
	v382 = v366 - v377
	if base.Ui32(int32(3)) < base.Ui32(v382) {
		v364 = v380
		v365 = v378
		v366 = v382
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v393 = v386
	v394 = v387
	v395 = v388
	goto L125
L146:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	*(*uint8)(unsafe.Add(mBase, uint32(v399))) = uint8(v402)
	if v402 == int32(0) {
		v413 = v398
		v414 = v399
		goto L124
	} else {
		goto L148
	}
L147:
	;
	v413 = v409
	v414 = v407
	goto L124
L148:
	;
	v406 = int32(1)
	v407 = v399 + v406
	v409 = v398 + v406
	v411 = v400 - v406
	if v411 != 0 {
		v398 = v409
		v399 = v407
		v400 = v411
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	goto L21
L151:
	;
	goto L21
L152:
	;
	F_SetConfigOption(m, v447, int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L49
	} else {
		goto L155
	}
L153:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
	switch v443 - int32(97) {
	case 0:
		v447 = int32(_a_F_process_postgres_switches_33)
		goto L152
	default:
		goto L20
	case 11:
		goto L154
	}
L154:
	;
	v447 = int32(_a_F_process_postgres_switches_34)
	goto L152
L155:
	;
	goto L21
L156:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	v459 = v455
	goto L158
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[10])) = v503
	goto L21
L158:
	;
	v464 = v459 + int32(1)
	v465 = int32(*(*int8)(unsafe.Add(mBase, uint32(v459))))
	v466 = F___isspace(m, v465)
	mBase = m.M
	if v466 != 0 {
		v459 = v464
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v467 = int32(1)
	switch v465&int32(255) - int32(43) {
	case 0:
		v473 = v467
		goto L162
	default:
		v475 = v465
		v476 = v459
		v477 = v467
		goto L161
	case 2:
		goto L163
	}
L160:
	;
	goto L159
L161:
	;
	v478 = int32(0)
	v480 = v475 - int32(48)
	if base.Ui32(v480) <= base.Ui32(int32(9)) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v474 = int32(*(*int8)(unsafe.Add(mBase, uint32(v464))))
	v475 = v474
	v476 = v464
	v477 = v473
	goto L161
L163:
	;
	v473 = int32(0)
	goto L162
L164:
	;
	v483 = v478
	v484 = v480
	v485 = v476
	goto L167
L165:
	;
	v497 = v478
	goto L166
L166:
	;
	if v477 != 0 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v487 = int32(10)
	v489 = v483*v487 - v484
	v490 = int32(*(*int8)(unsafe.Add(mBase, uint32(v485)+1)))
	v494 = v490 - int32(48)
	if base.Ui32(v494) < base.Ui32(v487) {
		v483 = v489
		v484 = v494
		v485 = v485 + int32(1)
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v497 = v489
	goto L166
L169:
	;
	goto L168
L170:
	;
	v503 = int32(0) - v497
	goto L172
L171:
	;
	v503 = v497
	goto L172
L172:
	;
	goto L157
L173:
	;
	goto L21
L174:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2]))
	if v53 != v540 {
		goto L19
	} else {
		goto L182
	}
L175:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v512 != 0 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2]))
	if v53-v514 <= int32(0) {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v519 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2])) = v514 + v519
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v54+v514<<(uint(int32(2))%32))))
	v528 = F_strlen(m, v525)
	mBase = m.M
	v530 = v528 + v519
	v531 = F_emscripten_builtin_malloc(m, v530)
	mBase = m.M
	if v531 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v536
	goto L174
L179:
	;
	v536 = int32(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v535 = F___memcpy(m, v531, v525, v530)
	mBase = m.M
	v536 = v535
	goto L178
L182:
	;
	v543 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[11])) = v543
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2])) = v543
	m.G0 = v11 + int32(112)
	return
L183:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L49
	} else {
		goto L184
	}
L184:
	;
	if v563 == int32(1) {
		goto L16
	} else {
		goto L185
	}
L185:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[2]))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v54+v574<<(uint(int32(2))%32))))
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v578
	F_errmsg(m, int32(_a_F_process_postgres_switches_35), v11+int32(48))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L49
	} else {
		goto L186
	}
L186:
	;
	v589 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v589
	F_errhint(m, int32(_a_F_process_postgres_switches_36), v11+int32(32))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L49
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(_a_F_process_postgres_switches_37), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L49
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L49
	} else {
		goto L190
	}
L190:
	;
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v609
	F_errmsg(m, int32(_a_F_process_postgres_switches_38), v11+int32(96))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L49
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(3965), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L49
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(3980), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L49
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v645
	F_errhint(m, int32(_a_F_process_postgres_switches_36), v11)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L49
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(_a_F_process_postgres_switches_39), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L49
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_procsignal_sigusr1_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[1]))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	if v411 != 0 {
		goto L123
	} else {
		goto L124
	}
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+40))
	if v6 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[2])) = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[1]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v61 = v3
	goto L5
L5:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v58 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L7:
	;
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v20 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[3]))
	if v24 == v20 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4]))
	if v31 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v54 = F_pgmem_kill(m, v20, int32(23))
	mBase = m.M
	goto L7
L14:
	;
	m.G0 = v28 + int32(16)
	goto L6
L15:
	;
	v34 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)) = uint8(v34)
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[5]))
	v42 = F_write(m, v38, v28+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v42 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	goto L14
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[6]))
	if v46 == int32(27) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v61 = v58
	goto L5
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[7])) = int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[1]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v117 = v61
	goto L23
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+48))
	if v118 != 0 {
		goto L39
	} else {
		goto L40
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v114 == int32(0) {
		goto L1
	} else {
		goto L38
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v73 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v76 == int32(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[3]))
	if v80 == v76 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = m.G0
	v84 = v82 - int32(16)
	m.G0 = v84
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4]))
	if v87 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v110 = F_pgmem_kill(m, v76, int32(23))
	mBase = m.M
	goto L25
L32:
	;
	m.G0 = v84 + int32(16)
	goto L24
L33:
	;
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+15)) = uint8(v90)
	goto L34
L34:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[5]))
	v98 = F_write(m, v94, v84+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v98 {
		goto L32
	} else {
		goto L36
	}
L35:
	;
	goto L32
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[6]))
	if v102 == int32(27) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v117 = v114
	goto L23
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+48)) = int32(0)
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v122
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[9])) = v122
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[1]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v129 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v176 = v117
	goto L41
L41:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+52))
	if v177 != 0 {
		goto L57
	} else {
		goto L58
	}
L42:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v173 == int32(0) {
		goto L1
	} else {
		goto L56
	}
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v132 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v135 == int32(0) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[3]))
	if v139 == v135 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v141 = m.G0
	v143 = v141 - int32(16)
	m.G0 = v143
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4]))
	if v146 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v169 = F_pgmem_kill(m, v135, int32(23))
	mBase = m.M
	goto L43
L50:
	;
	m.G0 = v143 + int32(16)
	goto L42
L51:
	;
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+15)) = uint8(v149)
	goto L52
L52:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[5]))
	v157 = F_write(m, v153, v143+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v157 {
		goto L50
	} else {
		goto L54
	}
L53:
	;
	goto L50
L54:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[6]))
	if v161 == int32(27) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v176 = v173
	goto L41
L57:
	;
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v176)+52)) = v178
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[10]))
	if v181 == v178 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v195 = v176
	goto L59
L59:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+56))
	if v196 != 0 {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v192 == int32(0) {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[3]))
	v187 = F_pgmem_kill(m, v185, int32(15))
	mBase = m.M
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[11])) = int32(1)
	goto L60
L64:
	;
	v195 = v192
	goto L59
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+56)) = int32(0)
	v200 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v200
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[12])) = v200
	goto L67
L66:
	;
	goto L67
L67:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)+60))
	if v205 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v206 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+60)) = v206
	v209 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v209
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[13])) = v209
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v215 == v206 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	v218 = v195
	goto L70
L70:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+64))
	if v219 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v218 = v215
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+64)) = int32(0)
	v223 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v223
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[14])) = v223
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[1]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v230 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v277 = v218
	goto L74
L74:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+68))
	if v278 != 0 {
		goto L90
	} else {
		goto L91
	}
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v274 == int32(0) {
		goto L1
	} else {
		goto L89
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v233 == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	if v236 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[3]))
	if v240 == v236 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v242 = m.G0
	v244 = v242 - int32(16)
	m.G0 = v244
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4]))
	if v247 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v270 = F_pgmem_kill(m, v236, int32(23))
	mBase = m.M
	goto L76
L83:
	;
	m.G0 = v244 + int32(16)
	goto L75
L84:
	;
	v250 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+15)) = uint8(v250)
	goto L85
L85:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[5]))
	v258 = F_write(m, v254, v244+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v258 {
		goto L83
	} else {
		goto L87
	}
L86:
	;
	goto L83
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[6]))
	if v262 == int32(27) {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v277 = v274
	goto L74
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+68)) = int32(0)
	v284 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[15])) = v284
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v284
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v284
	goto L93
L91:
	;
	v296 = v277
	goto L92
L92:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+72))
	if v297 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v293 == int32(0) {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v296 = v293
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+72)) = int32(0)
	v303 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[17])) = v303
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v303
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v303
	goto L98
L96:
	;
	v315 = v296
	goto L97
L97:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+76))
	if v316 != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v312 == int32(0) {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v315 = v312
	goto L97
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+76)) = int32(0)
	v322 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[18])) = v322
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v322
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v322
	goto L103
L101:
	;
	v334 = v315
	goto L102
L102:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+80))
	if v335 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v331 == int32(0) {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v334 = v331
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+80)) = int32(0)
	v341 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[19])) = v341
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v341
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v341
	goto L108
L106:
	;
	v353 = v334
	goto L107
L107:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+84))
	if v354 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v350 == int32(0) {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v353 = v350
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+84)) = int32(0)
	v360 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[20])) = v360
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v360
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v360
	goto L113
L111:
	;
	v372 = v353
	goto L112
L112:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+92))
	if v373 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v369 == int32(0) {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v372 = v369
	goto L112
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v372)+92)) = int32(0)
	v379 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[21])) = v379
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v379
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v379
	goto L118
L116:
	;
	v391 = v372
	goto L117
L117:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+88))
	if v392 == int32(0) {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v388 == int32(0) {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v391 = v388
	goto L117
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+88)) = int32(0)
	v400 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[22])) = v400
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v400
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = v400
	goto L121
L121:
	;
	goto L1
L122:
	;
	return
L123:
	;
	goto L122
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = int32(1)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v414 == int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v410)+12))
	if v417 == int32(0) {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[3]))
	if v421 == v417 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v423 = m.G0
	v425 = v423 - int32(16)
	m.G0 = v425
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4]))
	if v428 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v451 = F_pgmem_kill(m, v417, int32(23))
	mBase = m.M
	goto L123
L130:
	;
	m.G0 = v425 + int32(16)
	goto L122
L131:
	;
	v431 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v425)+15)) = uint8(v431)
	goto L132
L132:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[5]))
	v439 = F_write(m, v435, v425+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v439 {
		goto L130
	} else {
		goto L134
	}
L133:
	;
	goto L130
L134:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[6]))
	if v443 == int32(27) {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
}
func F_prs_process_call(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 < v12 {
		v15 = v8 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v11<<(uint(int32(3))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v21
		v24 = F_pg_sprintf(m, v15, int32(_a_F_prs_process_call_0), v8)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(3))%32))+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v33
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v38 = F_BuildTupleFromCStrings(m, v35, v8+int32(40))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
				v41 = F_HeapTupleHeaderGetDatum(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
					F_pfree(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v46 + int32(1)
						v50 = v41
						m.G0 = v8 + int32(48)
						return v50
					}
				}
			}
		}
	} else {
		v50 = int32(0)
		m.G0 = v8 + int32(48)
		return v50
	}
}
func F_prs_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = F_lookup_ts_parser_cache(m, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v18
	v22 = int32(_a_F_prs_setup_firstcall_0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_prs_setup_firstcall[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_prs_setup_firstcall[0])) = v25
	v28 = F_palloc(m, int32(12))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = int64(68719476736)
	v33 = F_palloc(m, int32(128))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v33
	v37 = v16 + int32(56)
	v38 = int32(0)
	v42 = int32(1)
	v43 = l3 + v42
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v48 = v46 & v42
	if v48 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v49 = v43
	goto L7
L6:
	;
	v49 = l3 + int32(4)
	goto L7
L7:
	;
	if v46 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v77 = F_FunctionCall2Coll(m, v16+int32(28), v38, v49, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v55 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v66 = int32(1)
	if v48 != 0 {
		v76 = int32(base.Ui32(v46)>>(uint(v66)%32)) - v66
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v58 = int32(16)
	goto L14
L13:
	;
	v58 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = int32(4)
	goto L17
L16:
	;
	v65 = v58
	goto L17
L17:
	;
	v76 = v65
	goto L8
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	v83 = F_FunctionCall3Coll(m, v37, v38, v77, v14+int32(8), v14+int32(4))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v83 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v88 = v83
	goto L24
L22:
	;
	goto L23
L23:
	;
	v169 = F_FunctionCall1Coll(m, v16+int32(84), int32(0), v77)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v96 <= v97 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v96 << (uint(int32(1)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v105 = F_repalloc(m, v102, v96<<(uint(int32(4))%32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v111 = F_palloc(m, v108+int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v105
	goto L28
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v113+v114<<(uint(int32(3))%32))+4)) = v111
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v119 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120+v121<<(uint(int32(3))%32))+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	base.MemoryCopy(m, v125, v126, v119)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v130 = int32(3)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128+v129<<(uint(v130)%32))+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v134))) = uint8(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v139<<(uint(v130)%32)))) = v88
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v144 + int32(1)
	v153 = F_FunctionCall3Coll(m, v37, v136, v77, v14+int32(8), v14+int32(4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v153 != 0 {
		v88 = v153
		goto L24
	} else {
		goto L35
	}
L35:
	;
	goto L25
L36:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v171
	v173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v28
	v179 = F_get_call_result_type(m, l1, v173, v14+int32(12))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v179 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v196
	v198 = F_TupleDescGetAttInMetadata(m, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	F_errmsg_internal(m, int32(_a_F_prs_setup_firstcall_1), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_prs_setup_firstcall_2), int32(209), int32(_a_F_prs_setup_firstcall_3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v198
	*(*int32)(unsafe.Add(mBase, _c_F_prs_setup_firstcall[0])) = v23
	m.G0 = v14 + int32(16)
	return
}
func F_pull_ands(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v2 = int32(0)
	if l0 == v2 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v9 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = v2
	v15 = v2
	goto L7
L5:
	;
	v43 = v2
	goto L6
L6:
	;
	return v43
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v15<<(uint(int32(2))%32))))
	if v20 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v43 = v36
	goto L6
L9:
	;
	v38 = v15 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 < v39 {
		v14 = v36
		v15 = v38
		goto L7
	} else {
		goto L18
	}
L10:
	;
	v34 = F_lappend(m, v14, v20)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L17
	}
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v23 != int32(21) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v26 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v28 = F_pull_ands(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v32 = F_list_concat(m, v14, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v36 = v32
	goto L9
L17:
	;
	v36 = v34
	goto L9
L18:
	;
	goto L8
}
func F_pull_varnos_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
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
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
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
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	v3 = int32(0)
	if l0 == v3 {
		v218 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v218
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v7 - int32(58) {
	case 0:
		goto L6
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L3
	case 9:
		goto L4
	default:
		goto L7
	}
L3:
	;
	v215 = F_expression_tree_walker_impl(m, l0, int32(896), l1)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L61
	}
L4:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v199 + int32(1)
	v205 = F_query_tree_walker_impl(m, l0, int32(896), l1, int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L60
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v38 != v39 {
		goto L3
	} else {
		goto L16
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v30 != 0 {
		v218 = v3
		goto L1
	} else {
		goto L14
	}
L7:
	;
	if v7 == int32(319) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v7 != int32(6) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 != v15 {
		v218 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = F_bms_add_member(m, v17, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_bms_add_members(m, v19, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
	return int32(0)
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = F_bms_add_member(m, v31, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
	return int32(0)
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if v38 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v194 = F_bms_add_members(m, v191, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L59
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v61 = int32(0)
	if base.B2i32(v58 == v61)|base.B2i32(v60 == v61) != 0 {
		v107 = base.B2i32(v58|v60 == v61)
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = F_bms_add_members(m, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L24
	}
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+148))
	if base.Ui32(v45) <= base.Ui32(v44) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+144))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44<<(uint(int32(2))%32))))
	if v51 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v191 = v56
	goto L18
L25:
	;
	if v107 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	goto L25
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v75 != v76 {
		v107 = int32(0)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v78 = int32(1)
	if v75 <= v78 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = v78
	goto L31
L30:
	;
	v81 = v75
	goto L31
L31:
	;
	v82 = int32(8)
	v87 = int32(0)
	goto L32
L32:
	;
	v95 = v87 << (uint(int32(2)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v58+v82+v95)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60+v82+v95)))
	v100 = base.B2i32(v97 == v99)
	if v97 != v99 {
		v107 = v100
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v107 = v100
	goto L26
L34:
	;
	v103 = v87 + int32(1)
	if v103 != v81 {
		v87 = v103
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v114 = F_bms_add_members(m, v112, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L11
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v119 = F_bms_difference(m, v117, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L11
	} else {
		goto L40
	}
L39:
	;
	v191 = v114
	goto L18
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v122 = F_bms_difference(m, v121, v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v125 = int32(0)
	if base.B2i32(v122 == v125)|base.B2i32(v124 == v125) != 0 {
		v171 = base.B2i32(v122|v124 == v125)
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v171 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L43:
	;
	goto L42
L44:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v139 != v140 {
		v171 = int32(0)
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v142 = int32(1)
	if v139 <= v142 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v145 = v142
	goto L48
L47:
	;
	v145 = v139
	goto L48
L48:
	;
	v146 = int32(8)
	v151 = int32(0)
	goto L49
L49:
	;
	v159 = v151 << (uint(int32(2)) % 32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v122+v146+v159)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v124+v146+v159)))
	v164 = base.B2i32(v161 == v163)
	if v161 != v163 {
		v171 = v164
		goto L43
	} else {
		goto L51
	}
L50:
	;
	v171 = v164
	goto L43
L51:
	;
	v167 = v151 + int32(1)
	if v167 != v145 {
		v151 = v167
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	v181 = F_bms_difference(m, v178, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L11
	} else {
		goto L56
	}
L54:
	;
	v185 = v122
	goto L55
L55:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v187 = F_bms_join(m, v186, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L58
	}
L56:
	;
	v183 = F_bms_join(m, v122, v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v185 = v183
	goto L55
L58:
	;
	v191 = v187
	goto L18
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v194
	return int32(0)
L60:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v207 - int32(1)
	return v205
L61:
	;
	v218 = v215
	goto L1
}
func F_pullf_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		m.T0[v4].(func(*base.Module, int32))(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v8 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v10 != 0 {
					base.MemoryFill(m, v8, int32(0), v10)
				} else {
				}
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_pfree(m, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					base.MemoryFill(m, l0, int32(0), int32(24))
					F_pfree(m, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				base.MemoryFill(m, l0, int32(0), int32(24))
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v10 != 0 {
				base.MemoryFill(m, v8, int32(0), v10)
			} else {
			}
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pfree(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				base.MemoryFill(m, l0, int32(0), int32(24))
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			base.MemoryFill(m, l0, int32(0), int32(24))
			F_pfree(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pullf_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v7 == int32(0) {
		v10 = l0
		for {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			if v17 == int32(0) {
				v10 = v15
				continue
			} else {
				break
			}
			break
		}
		v20 = v15
		v24 = v17
	} else {
		v20 = l0
		v24 = v7
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if l1 < v27 {
		v29 = l1
	} else {
		v29 = v27
	}
	if v27 != 0 {
		v30 = v29
	} else {
		v30 = l1
	}
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v32 = m.T0[v24].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v25, v26, v30, l2, v31, v27)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		return v32
	}
}
func F_pullup_replace_vars_callback(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v18 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	return v12
L6:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v574 == int32(0) {
		v752 = v567
		goto L154
	} else {
		goto L155
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v48 = int32(0)
	v50 = F_ReplaceVarFromTargetList(m, l0, v46, v24, v47, v48, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L20
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v41 = int32(0)
	v43 = F_ReplaceVarFromTargetList(m, l0, v38, v39, v40, v41, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L19
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v21 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = v25
	goto L15
L14:
	;
	v27 = int32(0)
	goto L15
L15:
	;
	if v27 < v9 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v9<<(uint(int32(2))%32))))
	if v33 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v36 = F_copyObjectImpl(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v567 = v36
	goto L6
L19:
	;
	v567 = v43
	goto L6
L20:
	;
	if v9 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v550 = F_bms_make_singleton(m, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L147
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v54 == int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v50 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+124)))
	if v337 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v59 != int32(319) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v59 != int32(6) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v141 != 0 {
		goto L24
	} else {
		goto L49
	}
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	if v64 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+124)))
	if v66 != int32(1) {
		v567 = v50
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v71 = F_bms_is_member(m, v69, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v71 != 0 {
		v567 = v50
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v76 = int32(2)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74+v75<<(uint(v76)%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74+v80<<(uint(v76)%32))))
	v85 = int32(0)
	if v79 == v85 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v138 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L35:
	;
	v138 = int32(1)
	goto L34
L36:
	;
	goto L37
L37:
	;
	if v84 == int32(0) {
		v131 = v85
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v138 = v131
	goto L34
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v95 < v94 {
		v131 = v85
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v97 = int32(1)
	if v94 <= v97 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v100 = v97
	goto L43
L42:
	;
	v100 = v94
	goto L43
L43:
	;
	v101 = int32(8)
	v106 = int32(0)
	goto L44
L44:
	;
	v113 = v106 << (uint(int32(2)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v79+v101+v113)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v84+v101+v113)))
	v120 = v115 & (v117 ^ int32(-1))
	v122 = base.B2i32(v120 == int32(0))
	if v120 != 0 {
		v131 = v122
		goto L38
	} else {
		goto L46
	}
L45:
	;
	v131 = v122
	goto L38
L46:
	;
	v124 = v106 + int32(1)
	if v124 != v100 {
		v106 = v124
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v567 = v50
	goto L6
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+124)))
	if v143 != int32(1) {
		v567 = v50
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v148 = int32(0)
	if v146 == v148 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v201 != 0 {
		v567 = v50
		goto L6
	} else {
		goto L65
	}
L52:
	;
	v201 = int32(1)
	goto L51
L53:
	;
	goto L54
L54:
	;
	if v147 == int32(0) {
		v194 = v148
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v201 = v194
	goto L51
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v158 < v157 {
		v194 = v148
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v160 = int32(1)
	if v157 <= v160 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v163 = v160
	goto L60
L59:
	;
	v163 = v157
	goto L60
L60:
	;
	v164 = int32(8)
	v169 = int32(0)
	goto L61
L61:
	;
	v176 = v169 << (uint(int32(2)) % 32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v146+v164+v176)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v147+v164+v176)))
	v183 = v178 & (v180 ^ int32(-1))
	v185 = base.B2i32(v183 == int32(0))
	if v183 != 0 {
		v194 = v185
		goto L55
	} else {
		goto L63
	}
L62:
	;
	v194 = v185
	goto L55
L63:
	;
	v187 = v169 + int32(1)
	if v187 != v163 {
		v169 = v187
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v208 = int32(-1)
	goto L66
L66:
	;
	if v202 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L21
L68:
	;
	if v268 < int32(0) {
		v567 = v50
		goto L6
	} else {
		goto L79
	}
L69:
	;
	v268 = base.I32_ctz(v254) | v255<<(uint(int32(5))%32)
	goto L68
L70:
	;
	v268 = int32(-2)
	goto L68
L71:
	;
	v219 = v208 + int32(1)
	v221 = base.I32_div_s(v219, int32(32))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v222 <= v221 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v225 = v202 + int32(8)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v221<<(uint(int32(2))%32))))
	v232 = v229 & (int32(-1) << (uint(v219) % 32))
	if v232 != 0 {
		v254 = v232
		v255 = v221
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v234 = v221 + int32(1)
	if v234 == v222 {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v237 = v234
	goto L75
L75:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v225+v237<<(uint(int32(2))%32))))
	if v244 != 0 {
		v254 = v244
		v255 = v237
		goto L69
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v246 = v237 + int32(1)
	if v246 != v222 {
		v237 = v246
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v273 = int32(2)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v271+v272<<(uint(v273)%32))))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v271+v268<<(uint(v273)%32))))
	v281 = int32(0)
	if v276 == v281 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v334 != 0 {
		v208 = v268
		goto L66
	} else {
		goto L94
	}
L81:
	;
	v334 = int32(1)
	goto L80
L82:
	;
	goto L83
L83:
	;
	if v280 == int32(0) {
		v327 = v281
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v334 = v327
	goto L80
L85:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v291 < v290 {
		v327 = v281
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v293 = int32(1)
	if v290 <= v293 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v296 = v293
	goto L89
L88:
	;
	v296 = v290
	goto L89
L89:
	;
	v297 = int32(8)
	v302 = int32(0)
	goto L90
L90:
	;
	v309 = v302 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v276+v297+v309)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v280+v297+v309)))
	v316 = v311 & (v313 ^ int32(-1))
	v318 = base.B2i32(v316 == int32(0))
	if v316 != 0 {
		v327 = v318
		goto L84
	} else {
		goto L92
	}
L91:
	;
	v327 = v318
	goto L84
L92:
	;
	v320 = v302 + int32(1)
	if v320 != v296 {
		v302 = v320
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	goto L67
L95:
	;
	v536 = F_contain_nonstrict_functions_walker(m, v50, int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L145
	}
L96:
	;
	v341 = F_contain_vars_of_level(m, v50, int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v344 = F_pull_varnos(m, v343, v50)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	if v341 != 0 {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L21
L101:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v347 = int32(0)
	if base.B2i32(v344 == v347)|base.B2i32(v346 == v347) != 0 {
		v392 = v347
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v392 != 0 {
		goto L95
	} else {
		goto L115
	}
L103:
	;
	goto L102
L104:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v357 < v358 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v360 = v357
	goto L107
L106:
	;
	v360 = v358
	goto L107
L107:
	;
	if v360 <= int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v363 = int32(1)
	goto L110
L109:
	;
	v363 = v360
	goto L110
L110:
	;
	v364 = int32(8)
	v369 = int32(0)
	goto L111
L111:
	;
	v376 = v369 << (uint(int32(2)) % 32)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v346+v364+v376)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v344+v364+v376)))
	v381 = v378 & v380
	v383 = base.B2i32(v381 != int32(0))
	if v381 != 0 {
		v392 = v383
		goto L103
	} else {
		goto L113
	}
L112:
	;
	v392 = v383
	goto L103
L113:
	;
	v385 = v369 + int32(1)
	if v385 != v363 {
		v369 = v385
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v398 = int32(-1)
	goto L116
L116:
	;
	if v344 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	goto L95
L118:
	;
	if v458 < int32(0) {
		goto L21
	} else {
		goto L129
	}
L119:
	;
	v458 = base.I32_ctz(v444) | v445<<(uint(int32(5))%32)
	goto L118
L120:
	;
	v458 = int32(-2)
	goto L118
L121:
	;
	v409 = v398 + int32(1)
	v411 = base.I32_div_s(v409, int32(32))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v412 <= v411 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v415 = v344 + int32(8)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415+v411<<(uint(int32(2))%32))))
	v422 = v419 & (int32(-1) << (uint(v409) % 32))
	if v422 != 0 {
		v444 = v422
		v445 = v411
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v424 = v411 + int32(1)
	if v424 == v412 {
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v427 = v424
	goto L125
L125:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v415+v427<<(uint(int32(2))%32))))
	if v434 != 0 {
		v444 = v434
		v445 = v427
		goto L119
	} else {
		goto L127
	}
L126:
	;
	goto L120
L127:
	;
	v436 = v427 + int32(1)
	if v436 != v412 {
		v427 = v436
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v463 = int32(2)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v461+v462<<(uint(v463)%32))))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v461+v458<<(uint(v463)%32))))
	v471 = int32(0)
	if v466 == v471 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v524 == int32(0) {
		v398 = v458
		goto L116
	} else {
		goto L144
	}
L131:
	;
	v524 = int32(1)
	goto L130
L132:
	;
	goto L133
L133:
	;
	if v470 == int32(0) {
		v517 = v471
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v524 = v517
	goto L130
L135:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v481 < v480 {
		v517 = v471
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v483 = int32(1)
	if v480 <= v483 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v486 = v483
	goto L139
L138:
	;
	v486 = v480
	goto L139
L139:
	;
	v487 = int32(8)
	v492 = int32(0)
	goto L140
L140:
	;
	v499 = v492 << (uint(int32(2)) % 32)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v466+v487+v499)))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v470+v487+v499)))
	v506 = v501 & (v503 ^ int32(-1))
	v508 = base.B2i32(v506 == int32(0))
	if v506 != 0 {
		v517 = v508
		goto L134
	} else {
		goto L142
	}
L141:
	;
	v517 = v508
	goto L134
L142:
	;
	v510 = v492 + int32(1)
	if v510 != v486 {
		v492 = v510
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	goto L117
L145:
	;
	if v536 == int32(0) {
		v567 = v50
		goto L6
	} else {
		goto L146
	}
L146:
	;
	goto L21
L147:
	;
	v552 = F_make_placeholder_expr(m, v548, v50, v550)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v554 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+4))
	v557 = v555
	goto L151
L150:
	;
	v557 = int32(0)
	goto L151
L151:
	;
	if v557 < v9 {
		v567 = v552
		goto L6
	} else {
		goto L152
	}
L152:
	;
	v559 = F_copyObjectImpl(m, v552)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v561+v9<<(uint(int32(2))%32)))) = v559
	v567 = v552
	goto L6
L154:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v759 != 0 {
		goto L200
	} else {
		goto L201
	}
L155:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if v577 != int32(319) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+124)))
	if v591 != int32(1) {
		v740 = v567
		goto L163
	} else {
		goto L164
	}
L157:
	;
	if v577 != int32(6) {
		goto L156
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v587 = F_bms_add_members(m, v586, v574)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v567)+24))
	v583 = F_bms_add_members(m, v582, v574)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+24)) = v583
	v752 = v567
	goto L154
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+12)) = v587
	v752 = v567
	goto L154
L163:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v749 = F_add_nulling_relids(m, v740, v747, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L4
	} else {
		goto L199
	}
L164:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v596 = F_pull_varnos(m, v595, v567)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v599 = F_bms_del_members(m, v596, v598)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	if v599 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	if v657 < int32(0) {
		v740 = v567
		goto L163
	} else {
		goto L178
	}
L168:
	;
	v657 = base.I32_ctz(v643) | v644<<(uint(int32(5))%32)
	goto L167
L169:
	;
	v657 = int32(-2)
	goto L167
L170:
	;
	v610 = base.I32_div_s(int32(0), int32(32))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v611 <= v610 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v614 = v599 + int32(8)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v614+v610<<(uint(int32(2))%32))))
	v621 = v618 & int32(-1)
	if v621 != 0 {
		v643 = v621
		v644 = v610
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v623 = v610 + int32(1)
	if v623 == v611 {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v626 = v623
	goto L174
L174:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v614+v626<<(uint(int32(2))%32))))
	if v633 != 0 {
		v643 = v633
		v644 = v626
		goto L168
	} else {
		goto L176
	}
L175:
	;
	goto L169
L176:
	;
	v635 = v626 + int32(1)
	if v635 != v611 {
		v626 = v635
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v661 = v567
	v663 = v657
	goto L179
L179:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v669+v663<<(uint(int32(2))%32))))
	v674 = F_bms_intersect(m, v668, v673)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L4
	} else {
		goto L181
	}
L180:
	;
	v740 = v680
	goto L163
L181:
	;
	if v674 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v676 = F_bms_make_singleton(m, v663)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L185
	}
L183:
	;
	v680 = v661
	goto L184
L184:
	;
	if v599 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v678 = F_add_nulling_relids(m, v661, v676, v674)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	v680 = v678
	goto L184
L187:
	;
	if int32(0) <= v736 {
		v661 = v680
		v663 = v736
		goto L179
	} else {
		goto L198
	}
L188:
	;
	v736 = base.I32_ctz(v722) | v723<<(uint(int32(5))%32)
	goto L187
L189:
	;
	v736 = int32(-2)
	goto L187
L190:
	;
	v687 = v663 + int32(1)
	v689 = base.I32_div_s(v687, int32(32))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v690 <= v689 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v693 = v599 + int32(8)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v693+v689<<(uint(int32(2))%32))))
	v700 = v697 & (int32(-1) << (uint(v687) % 32))
	if v700 != 0 {
		v722 = v700
		v723 = v689
		goto L188
	} else {
		goto L192
	}
L192:
	;
	v702 = v689 + int32(1)
	if v702 == v690 {
		goto L189
	} else {
		goto L193
	}
L193:
	;
	v705 = v702
	goto L194
L194:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v693+v705<<(uint(int32(2))%32))))
	if v712 != 0 {
		v722 = v712
		v723 = v705
		goto L188
	} else {
		goto L196
	}
L195:
	;
	goto L189
L196:
	;
	v714 = v705 + int32(1)
	if v714 != v690 {
		v705 = v714
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	goto L180
L199:
	;
	v752 = v749
	goto L154
L200:
	;
	F_IncrementVarSublevelsUp(m, v752, v759, int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L4
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	return v752
L203:
	;
	goto L202
}
func F_push_old_value(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_push_old_value[0]))
	if v6 == int32(0) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if v9 == int32(0) {
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_push_old_value[1]))
			v39 = F_MemoryContextAllocZero(m, v37, int32(64))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = v41
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_push_old_value[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v44
				if base.Ui32(l1) <= base.Ui32(int32(2)) {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_c_F_push_old_value[2])))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v50
				} else {
				}
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v52
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v56
				F_set_stack_value(m, l0, v39+int32(32))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v62 == int32(0) {
						v65 = int32(_a_F_push_old_value_0)
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_push_old_value[3]))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v66
						*(*int32)(unsafe.Add(mBase, _c_F_push_old_value[3])) = l0 + int32(72)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v39
					return
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v12 < v6 {
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_push_old_value[1]))
				v39 = F_MemoryContextAllocZero(m, v37, int32(64))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = v41
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_push_old_value[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v44
					if base.Ui32(l1) <= base.Ui32(int32(2)) {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_c_F_push_old_value[2])))
						*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v50
					} else {
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v56
					F_set_stack_value(m, l0, v39+int32(32))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						if v62 == int32(0) {
							v65 = int32(_a_F_push_old_value_0)
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_push_old_value[3]))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v66
							*(*int32)(unsafe.Add(mBase, _c_F_push_old_value[3])) = l0 + int32(72)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v39
						return
					}
				}
			} else {
				switch l1 {
				case 0:
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					if v14 == int32(3) {
						F_discard_stack_value(m, l0, v9+int32(48))
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1)
						return
					}
				case 1:
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					if v23 != int32(1) {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v28
						F_set_stack_value(m, l0, v9+int32(48))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
							return
						}
					}
				default:
					return
				}
			}
		}
	}
}
