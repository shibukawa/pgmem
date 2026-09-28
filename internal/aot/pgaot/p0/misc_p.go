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
func F_PrepareRedoRemoveFull(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v1 = l0
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareRedoRemoveFull[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v15 <= v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return
L3:
	;
	v24 = v3
	goto L4
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8)+v24<<(uint(int32(2))%32))))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+32))
	if v1 != v32 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v39 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v35 = v24 + int32(1)
	if v15 != v35 {
		v24 = v35
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L2
L10:
	;
	return
L11:
	;
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v1)
	v43 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v43)
	F_errmsg_internal(m, int32(_a_F_PrepareRedoRemoveFull_0), v11+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+49)))
	if v55 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_errfinish(m, int32(_a_F_PrepareRedoRemoveFull_1), int32(2658), int32(_a_F_PrepareRedoRemoveFull_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_RemoveTwoPhaseFile(m, v1, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareRedoRemoveFull[0]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v62 <= int32(0) {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v66 = v61 + int32(8)
	v72 = int32(0)
	goto L22
L22:
	;
	v78 = v66 + v72<<(uint(int32(2))%32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 != v31 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v85 = v62 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v85
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v66+v85<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v90
	v92 = int32(_a_F_PrepareRedoRemoveFull_3)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareRedoRemoveFull[0]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v94
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareRedoRemoveFull[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v31
	goto L2
L24:
	;
	v82 = v72 + int32(1)
	if v62 != v82 {
		v72 = v82
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v31
	F_errmsg_internal(m, int32(_a_F_PrepareRedoRemoveFull_4), v11)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_PrepareRedoRemoveFull_1), int32(658), int32(_a_F_PrepareRedoRemoveFull_5))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
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
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
					F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_0), v8)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_1), int32(205), int32(_a_F_PrepareSortSupportFromGistIndexRel_2))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v34 = F_OidFunctionCall1Coll(m, v28, int32(0), base.I64_extend_i32_u(l1))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
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
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v44
			F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_3), v8+int32(16))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_PrepareSortSupportFromGistIndexRel_1), int32(194), int32(_a_F_PrepareSortSupportFromGistIndexRel_2))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v4 = F_pg_database_locale(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+v10<<(uint(v11)%32))))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 < v11 {
			v24 = F_pg_database_locale(m)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				if v26 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					v28 = m.T0[v27].(func(*base.Module, int32, int32) int32)(m, v14, v24)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v34 = v28
						v37 = v34
						return v37
					}
				} else {
					v34 = base.B2i32(base.Ui32(v14-int32(48)) < base.Ui32(int32(10)))
					v37 = v34
					return v37
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			if v18 != int32(1) {
				v24 = F_pg_database_locale(m)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					if v26 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
						v28 = m.T0[v27].(func(*base.Module, int32, int32) int32)(m, v14, v24)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v34 = v28
							v37 = v34
							return v37
						}
					} else {
						v34 = base.B2i32(base.Ui32(v14-int32(48)) < base.Ui32(int32(10)))
						v37 = v34
						return v37
					}
				}
			} else {
				if base.Ui32(int32(127)) < base.Ui32(v14) {
					v37 = int32(0)
					return v37
				} else {
					v24 = F_pg_database_locale(m)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						if v26 != 0 {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
							v28 = m.T0[v27].(func(*base.Module, int32, int32) int32)(m, v14, v24)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								v34 = v28
								v37 = v34
								return v37
							}
						} else {
							v34 = base.B2i32(base.Ui32(v14-int32(48)) < base.Ui32(int32(10)))
							v37 = v34
							return v37
						}
					}
				}
			}
		}
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
func F_palloc0_mul(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	v7 = base.I64_extend_i32_u(l0) * base.I64_extend_i32_u(l1)
	if int64(base.Ui64(v7)>>(uint(int64(32))%64)) == int64(0) {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_palloc0_mul[0]))
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v14)
		v16 = base.I32_wrap_i64(v7)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = m.T0[v19].(func(*base.Module, int32, int32, int32) int32)(m, v13, v16, v14)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v16&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v16)) == int32(0) {
				if v16 == int32(0) {
				} else {
					v35 = v20 + v16
					v37 = v20 + int32(4)
					if base.Ui32(v37) < base.Ui32(v35) {
						v39 = v35
					} else {
						v39 = v37
					}
					v45 = (v20^int32(-1)+v39)&int32(-4) + int32(4)
					if v45 == int32(0) {
					} else {
						base.MemoryFill(m, v20, int32(0), v45)
					}
				}
			} else {
				v45 = v16
				if v45 == int32(0) {
				} else {
					base.MemoryFill(m, v20, int32(0), v45)
				}
			}
			return v20
		}
	} else {
		F_mul_size_error(m, l0, l1)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
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
func F_partitioned_table_reloptions(m *base.Module, l0 int64) {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	if l0 != int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_partitioned_table_reloptions_0), int32(0))
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_F_partitioned_table_reloptions_1), int32(0))
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_partitioned_table_reloptions_2), int32(2135), int32(_a_F_partitioned_table_reloptions_3))
						v23 = m.ExcPending
						if v23 != 0 {
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
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v189 float64
	_ = v189
	var v192 int32
	_ = v192
	var v201 float64
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v219 int32
	_ = v219
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v235 float64
	_ = v235
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v480 float64
	_ = v480
	var v482 int32
	_ = v482
	var v492 float64
	_ = v492
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 float64
	_ = v507
	var v508 float64
	_ = v508
	var v510 int32
	_ = v510
	var v522 float64
	_ = v522
	var v523 float64
	_ = v523
	var v526 float64
	_ = v526
	var v529 int32
	_ = v529
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v562 int32
	_ = v562
	var v563 int64
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	switch l1 - int32(1) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	default:
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L12
	} else {
		goto L169
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L12
	} else {
		goto L165
	}
L3:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v561 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v563 = F_datumCopy(m, v561, v562, v560)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L12
	} else {
		goto L160
	}
L4:
	;
	v554 = F_regex_fixed_prefix(m, l0, int32(1), l2, l3, l4)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L159
	}
L5:
	;
	v550 = F_regex_fixed_prefix(m, l0, int32(0), l2, l3, l4)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L12
	} else {
		goto L158
	}
L6:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v259 = F_pg_detoast_datum_packed(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L12
	} else {
		goto L81
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 != int32(17) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v73 = F_palloc(m, v69+int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L31
	}
L9:
	;
	v18 = F_text_to_cstring(m, v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v23 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	v22 = F_strlen(m, v18)
	mBase = m.M
	v69 = v22
	v70 = v18
	goto L8
L14:
	;
	v55 = F_palloc(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L26
	}
L15:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v25 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v31 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v42 = int32(1)
	if v25&v42 != 0 {
		v54 = int32(base.Ui32(v25)>>(uint(v42)%32)) - v42
		goto L14
	} else {
		goto L25
	}
L19:
	;
	v34 = int32(16)
	goto L21
L20:
	;
	v34 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v41 = int32(4)
	goto L24
L23:
	;
	v41 = v34
	goto L24
L24:
	;
	v54 = v41
	goto L14
L25:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L14
L26:
	;
	if v54 == int32(0) {
		v69 = v54
		v70 = v55
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v59 = int32(1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v61&v59 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v64 = v59
	goto L30
L29:
	;
	v64 = int32(4)
	goto L30
L30:
	;
	base.MemoryCopy(m, v55, v23+v64, v54)
	v69 = v54
	v70 = v55
	goto L8
L31:
	;
	v75 = int32(0)
	if v69 <= v75 {
		v110 = v75
		v112 = v75
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112+v73))) = uint8(v122)
	if v15 != int32(17) {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	v79 = v75
	v81 = v75
	goto L34
L34:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v70))))
	switch v91 - int32(92) {
	case 0:
		goto L37
	case 1, 2:
		v101 = v79
		v102 = v91
		goto L36
	case 3:
		v110 = v79
		v112 = v81
		goto L32
	default:
		goto L38
	}
L35:
	;
	v110 = v108
	v112 = v106
	goto L32
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v81+v73))) = uint8(v102)
	v105 = int32(1)
	v106 = v81 + v105
	v108 = v101 + v105
	if v108 < v69 {
		v79 = v108
		v81 = v106
		goto L34
	} else {
		goto L41
	}
L37:
	;
	v97 = v79 + int32(1)
	if v69 <= v97 {
		v110 = v97
		v112 = v81
		goto L32
	} else {
		goto L40
	}
L38:
	;
	if v91 != int32(37) {
		v101 = v79
		v102 = v91
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v110 = v79
	v112 = v81
	goto L32
L40:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v70))))
	v101 = v97
	v102 = v100
	goto L36
L41:
	;
	goto L35
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v149
	if l4 != 0 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v126 = F_string_to_const(m, v73, v15)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v129 = v112 + int32(4)
	v130 = F_palloc(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L47
	}
L46:
	;
	v149 = v126
	goto L42
L47:
	;
	if v112 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	base.MemoryCopy(m, v130+int32(4), v73, v112)
	goto L50
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v129 << (uint(int32(2)) % 32)
	v139 = int32(-1)
	v140 = int32(0)
	v145 = F_makeConst(m, int32(17), v139, v140, v139, base.I64_extend_i32_u(v130), v140, v140)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v149 = v145
	goto L42
L52:
	;
	v151 = v110 + v70
	v152 = int32(0)
	v153 = v69 - v110
	if v153 <= v152 {
		v179 = v152
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	F_pfree(m, v70)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L76
	}
L55:
	;
	v232 = float64(1)
	if base.F64_gt(v231, v232) != 0 {
		goto L73
	} else {
		goto L74
	}
L56:
	;
	v189 = float64(1)
	if v153 <= v179 {
		v231 = v189
		goto L55
	} else {
		goto L62
	}
L57:
	;
	v157 = v152
	goto L58
L58:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v151))))
	if base.B2i32(v168 != int32(95))&base.B2i32(v168 != int32(37)) != 0 {
		v179 = v157
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v231 = float64(1)
	goto L55
L60:
	;
	v175 = v157 + int32(1)
	if v175 != v153 {
		v157 = v175
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v192 = v179
	v201 = v189
	goto L63
L63:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192+v151))))
	switch v203 - int32(92) {
	case 0:
		goto L67
	case 1, 2:
		v213 = v192
		goto L66
	case 3:
		goto L68
	default:
		goto L69
	}
L64:
	;
	v231 = v217
	goto L55
L65:
	;
	v217 = base.F64_mul(v201, v216)
	v219 = v215 + int32(1)
	if v219 < v153 {
		v192 = v219
		v201 = v217
		goto L63
	} else {
		goto L72
	}
L66:
	;
	v215 = v213
	v216 = float64(0.2)
	goto L65
L67:
	;
	v211 = v192 + int32(1)
	if v153 <= v211 {
		v231 = v201
		goto L55
	} else {
		goto L71
	}
L68:
	;
	v215 = v192
	v216 = float64(0.9)
	goto L65
L69:
	;
	if v203 != int32(37) {
		v213 = v192
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v215 = v192
	v216 = float64(5)
	goto L65
L71:
	;
	v213 = v211
	goto L66
L72:
	;
	goto L64
L73:
	;
	v235 = v232
	goto L75
L74:
	;
	v235 = v231
	goto L75
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v235
	goto L54
L76:
	;
	F_pfree(m, v73)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	if v110 == v69 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v256 = int32(2)
	goto L80
L79:
	;
	v256 = base.B2i32(int32(0) < v112)
	goto L80
L80:
	;
	return v256
L81:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v262 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v261 == int32(17) {
		goto L2
	} else {
		goto L93
	}
L83:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	if v268 == int32(18) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v279 = int32(1)
	if v262&v279 != 0 {
		v291 = int32(base.Ui32(v262)>>(uint(v279)%32)) - v279
		goto L82
	} else {
		goto L92
	}
L86:
	;
	v271 = int32(16)
	goto L88
L87:
	;
	v271 = int32(0)
	goto L88
L88:
	;
	if base.Ui32((v268-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v278 = int32(4)
	goto L91
L90:
	;
	v278 = v271
	goto L91
L91:
	;
	v291 = v278
	goto L82
L92:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v291 = int32(base.Ui32(v285)>>(uint(int32(2))%32)) - int32(4)
	goto L82
L93:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v296 = F_pg_newlocale_from_collation(m, l2)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	v301 = v291<<(uint(int32(2))%32) + int32(4)
	v302 = F_palloc(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	v304 = int32(1)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v306&v304 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v309 = v304
	goto L99
L98:
	;
	v309 = int32(4)
	goto L99
L99:
	;
	v311 = F_pg_mb2wchar_with_len(m, v259+v309, v302, v291)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	v313 = F_palloc(m, v301)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	v315 = int32(0)
	if v311 <= v315 {
		v376 = v315
		v377 = v315
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313+v377<<(uint(int32(2))%32)))) = int32(0)
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_pattern_fixed_prefix[0]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v393*int32(28))+uint32(_c_F_pattern_fixed_prefix[1])))
	goto L119
L103:
	;
	v320 = v315
	v321 = v315
	goto L104
L104:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v302+v320<<(uint(int32(2))%32))))
	switch v333 - int32(92) {
	case 0:
		goto L107
	case 1, 2:
		v345 = v333
		v346 = v320
		goto L106
	case 3:
		v376 = v320
		v377 = v321
		goto L102
	default:
		goto L108
	}
L105:
	;
	v376 = v373
	v377 = v371
	goto L102
L106:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	if v347 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L107:
	;
	v339 = v320 + int32(1)
	if v311 <= v339 {
		v376 = v339
		v377 = v321
		goto L102
	} else {
		goto L110
	}
L108:
	;
	if v333 != int32(37) {
		v345 = v333
		v346 = v320
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v376 = v320
	v377 = v321
	goto L102
L110:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v302+v339<<(uint(int32(2))%32))))
	v345 = v344
	v346 = v339
	goto L106
L111:
	;
	if v361 != 0 {
		v376 = v346
		v377 = v321
		goto L102
	} else {
		goto L117
	}
L112:
	;
	if base.Ui32(int32(127)) < base.Ui32(v345) {
		v361 = int32(0)
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v347)+56))
	v359 = m.T0[v358].(func(*base.Module, int32, int32) int32)(m, v345, v296)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L116
	}
L115:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+uint32(_c_F_pattern_fixed_prefix[2]))))
	v361 = int32(base.Ui32(v353&int32(2)) >> (uint(int32(1)) % 32))
	goto L111
L116:
	;
	v361 = v359
	goto L111
L117:
	;
	v362 = int32(2)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v302+v346<<(uint(v362)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v313+v321<<(uint(v362)%32)))) = v368
	v370 = int32(1)
	v371 = v321 + v370
	v373 = v346 + v370
	if v373 < v311 {
		v320 = v373
		v321 = v371
		goto L104
	} else {
		goto L118
	}
L118:
	;
	goto L105
L119:
	;
	v402 = F_palloc(m, v398*v377+int32(1))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	v404 = F_pg_wchar2mb_with_len(m, v313, v402, v377)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L12
	} else {
		goto L121
	}
L121:
	;
	v407 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v404+v402))) = uint8(v407)
	F_pfree(m, v313)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L12
	} else {
		goto L122
	}
L122:
	;
	v412 = int32(-1)
	v415 = F_cstring_to_text(m, v402)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L12
	} else {
		goto L123
	}
L123:
	;
	v418 = int32(0)
	v420 = F_makeConst(m, int32(25), v412, int32(100), v412, base.I64_extend_i32_u(v415), v418, v418)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v420
	F_pfree(m, v402)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	if l4 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v425 = int32(0)
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_pattern_fixed_prefix[0]))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v431*int32(28))+uint32(_c_F_pattern_fixed_prefix[1])))
	goto L131
L127:
	;
	goto L128
L128:
	;
	F_pfree(m, v302)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L12
	} else {
		goto L154
	}
L129:
	;
	v523 = float64(1)
	if base.F64_gt(v522, v523) != 0 {
		goto L150
	} else {
		goto L151
	}
L130:
	;
	v480 = float64(1)
	if v443 <= v469 {
		v522 = v480
		goto L129
	} else {
		goto L139
	}
L131:
	;
	v437 = v311 - v376
	v441 = F_palloc(m, v436*v437+int32(1))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	v443 = F_pg_wchar2mb_with_len(m, v302+v376<<(uint(int32(2))%32), v441, v437)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L12
	} else {
		goto L133
	}
L133:
	;
	if v443 <= int32(0) {
		v469 = v425
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v447 = v425
	goto L135
L135:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v441))))
	if base.B2i32(v459 != int32(95))&base.B2i32(v459 != int32(37)) != 0 {
		v469 = v447
		goto L130
	} else {
		goto L137
	}
L136:
	;
	v522 = float64(1)
	goto L129
L137:
	;
	v466 = v447 + int32(1)
	if v466 != v443 {
		v447 = v466
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v482 = v469
	v492 = v480
	goto L140
L140:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482+v441))))
	switch v494 - int32(92) {
	case 0:
		goto L144
	case 1, 2:
		v504 = v482
		goto L143
	case 3:
		goto L145
	default:
		goto L146
	}
L141:
	;
	v522 = v508
	goto L129
L142:
	;
	v508 = base.F64_mul(v492, v507)
	v510 = v506 + int32(1)
	if v510 < v443 {
		v482 = v510
		v492 = v508
		goto L140
	} else {
		goto L149
	}
L143:
	;
	v506 = v504
	v507 = float64(0.2)
	goto L142
L144:
	;
	v502 = v482 + int32(1)
	if v443 <= v502 {
		v522 = v492
		goto L129
	} else {
		goto L148
	}
L145:
	;
	v506 = v482
	v507 = float64(0.9)
	goto L142
L146:
	;
	if v494 != int32(37) {
		v504 = v482
		goto L143
	} else {
		goto L147
	}
L147:
	;
	v506 = v482
	v507 = float64(5)
	goto L142
L148:
	;
	v504 = v502
	goto L143
L149:
	;
	goto L141
L150:
	;
	v526 = v523
	goto L152
L151:
	;
	v526 = v522
	goto L152
L152:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v526
	F_pfree(m, v441)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L12
	} else {
		goto L153
	}
L153:
	;
	goto L128
L154:
	;
	if v376 == v311 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v547 = int32(2)
	goto L157
L156:
	;
	v547 = base.B2i32(int32(0) < v377)
	goto L157
L157:
	;
	return v547
L158:
	;
	return v550
L159:
	;
	return v554
L160:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v567 = F_makeConst(m, v557, v558, v559, v560, v563, v565, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v567
	if l4 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
	goto L164
L163:
	;
	goto L164
L164:
	;
	return int32(1)
L165:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	F_errmsg(m, int32(_a_F_pattern_fixed_prefix_0), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_pattern_fixed_prefix_1), int32(1096), int32(_a_F_pattern_fixed_prefix_2))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L12
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L12
	} else {
		goto L170
	}
L170:
	;
	F_errmsg(m, int32(_a_F_pattern_fixed_prefix_3), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L12
	} else {
		goto L171
	}
L171:
	;
	F_errhint(m, int32(_a_F_pattern_fixed_prefix_4), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L12
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_pattern_fixed_prefix_1), int32(1107), int32(_a_F_pattern_fixed_prefix_2))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L12
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgarch_call_module_shutdown_cb(m *base.Module, l0 int32, l1 int64) {
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
func F_pgarch_die(m *base.Module, l0 int32, l1 int64) {
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
func F_pgstatindex_impl(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v221 float64
	_ = v221
	var v223 float64
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 float64
	_ = v301
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
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
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	v15 = int64(0)
	v24 = m.G0
	v26 = v24 - int32(256)
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
	return int64(0)
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
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L92
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L88
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L84
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L80
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
	*(*int32)(unsafe.Add(mBase, uint32(v26)+252)) = v79
	v82 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+248)) = v82
	v86 = int32(0)
	v91 = F_read_stream_begin_relation(m, int32(12), v29, l0, v86, int32(3), v26+int32(248), v86)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(v79) < base.Ui32(int32(2)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v214 = v15
	v215 = v15
	v216 = v15
	v217 = v15
	v218 = v15
	v221 = float64(0)
	v223 = float64(0)
	goto L24
L23:
	;
	v100 = v82
	v110 = v15
	v111 = v15
	v112 = v15
	v113 = v15
	v114 = v15
	v115 = v15
	v116 = v15
	goto L25
L24:
	;
	F_read_stream_end(m, v91)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L53
	}
L25:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[2]))
	if v120 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v214 = v186
	v215 = v187
	v216 = v188
	v217 = v189
	v218 = v190
	v221 = base.F64_convert_i64_u(v192)
	v223 = base.F64_convert_i64_u(v191)
	goto L24
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v124 = F_read_stream_next_buffer(m, v91, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	F_LockBufferInternal(m, v124, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v124 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_UnlockReleaseBuffer(m, v124)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L51
	}
L34:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+16)))
	v148 = v147 + v146
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+12)))
	if v149&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[0]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132+(v124^int32(-1))<<(uint(int32(2))%32))))
	v146 = v138
	goto L34
L36:
	;
	goto L37
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatindex_impl[1]))
	v146 = v140 + v124<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v186 = v110
	v187 = v111 + int64(1)
	v188 = v112
	v189 = v113
	v190 = v114
	v191 = v115
	v192 = v116
	goto L33
L39:
	;
	goto L40
L40:
	;
	if v149&int32(16) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v186 = v110
	v187 = v111
	v188 = v112 + int64(1)
	v189 = v113
	v190 = v114
	v191 = v115
	v192 = v116
	goto L33
L42:
	;
	goto L43
L43:
	;
	if v149&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+14)))
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+12)))
	v162 = v160 - v161
	v163 = int32(0)
	if v163 < v162 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	v186 = v110
	v187 = v111
	v188 = v112
	v189 = v113
	v190 = v114 + int64(1)
	v191 = v115
	v192 = v116
	goto L33
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v186 = v110 + int64(1)
	v187 = v111
	v188 = v112
	v189 = v113 + base.I64_extend_i32_s(v147-int32(24))
	v190 = v114
	v191 = v115 + base.I64_extend_i32_u(base.B2i32(v167 != int32(0))&base.B2i32(base.Ui32(v167) < base.Ui32(v100)))
	v192 = v116 + base.I64_extend_i32_u(v166)
	goto L33
L48:
	;
	v166 = v162
	goto L50
L49:
	;
	v166 = v163
	goto L50
L50:
	;
	goto L47
L51:
	;
	v196 = v100 + int32(1)
	if v196 != v79 {
		v100 = v196
		v110 = v186
		v111 = v187
		v112 = v188
		v113 = v189
		v114 = v190
		v115 = v191
		v116 = v192
		goto L25
	} else {
		goto L52
	}
L52:
	;
	goto L26
L53:
	;
	F_relation_close(m, l0, int32(1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v232 = F_get_call_result_type(m, l1, int32(0), v26+int32(244))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v232 != int32(1) {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v75
	v240 = F_psprintf(m, int32(_a_F_pgstatindex_impl_0), v26+int32(144))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+128)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v26)+192)) = v240
	v247 = F_psprintf(m, int32(_a_F_pgstatindex_impl_0), v26+int32(128))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+196)) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v26)+112)) = (v214+v218+v216+v215)<<(uint(int64(13))%64) - int64(-8192)
	v261 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(112))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v26)+200)) = v261
	v268 = F_psprintf(m, int32(_a_F_pgstatindex_impl_2), v26+int32(96))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+80)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v26)+204)) = v268
	v275 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(80))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v26)+208)) = v275
	v282 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26-int32(-64))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v26)+212)) = v282
	v289 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(48))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v26)+216)) = v289
	v296 = F_psprintf(m, int32(_a_F_pgstatindex_impl_1), v26+int32(32))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+220)) = v296
	if v217 != int64(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v316
	if v214 != int64(0) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	v301 = float64(100)
	*(*float64)(unsafe.Add(mBase, uint32(v26)+16)) = base.F64_sub(v301, base.F64_mul(base.F64_div(v221, base.F64_convert_i64_u(v217)), v301))
	v311 = F_psprintf(m, int32(_a_F_pgstatindex_impl_3), v26+int32(16))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v314 = F_pstrdup(m, int32(_a_F_pgstatindex_impl_4))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v316 = v311
	goto L65
L70:
	;
	v316 = v314
	goto L65
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v331
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v26)+244))
	v334 = F_TupleDescGetAttInMetadata(m, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L77
	}
L72:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v26))) = base.F64_mul(base.F64_div(v223, base.F64_convert_i64_u(v214)), float64(100))
	v326 = F_psprintf(m, int32(_a_F_pgstatindex_impl_3), v26)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v329 = F_pstrdup(m, int32(_a_F_pgstatindex_impl_4))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	v331 = v326
	goto L71
L76:
	;
	v331 = v329
	goto L71
L77:
	;
	v338 = F_BuildTupleFromCStrings(m, v334, v26+int32(192))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v338)+16))
	v341 = F_HeapTupleHeaderGetDatum(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	m.G0 = v26 + int32(256)
	return v341
L80:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+176)) = v354 + int32(4)
	F_errmsg(m, int32(_a_F_pgstatindex_impl_5), v26+int32(176))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(229), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_pgstatindex_impl_8), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(239), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+160)) = v391 + int32(4)
	F_errmsg(m, int32(_a_F_pgstatindex_impl_9), v26+int32(160))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(251), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errmsg_internal(m, int32(_a_F_pgstatindex_impl_10), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_pgstatindex_impl_6), int32(365), int32(_a_F_pgstatindex_impl_7))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
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
	var v8 int32
	_ = v8
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0)
	return v83
L2:
	;
	v8 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v8
	v12 = F_palloc_mul(m, int32(8), v8)
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
	v20 = F_palloc(m, v16<<(uint(int32(3))%32))
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
	v83 = v12
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
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v28<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v35 == int32(4) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v38 == int32(1) {
		v83 = v20
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v67 = F_get_rolespec_oid(m, v34, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L25
	}
L16:
	;
	v43 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v43 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
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
	v83 = v20
	goto L1
L21:
	;
	F_errmsg(m, int32(_a_F_policy_role_list_to_array_0), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(_a_F_policy_role_list_to_array_1), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_policy_role_list_to_array_2), int32(170), int32(_a_F_policy_role_list_to_array_3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20+v28<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v67)
	v72 = v28 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v72 < v73 {
		v28 = v72
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L12
}
func F_polish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = v8
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L58
L2:
	;
	if v125 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v125 = v97
	goto L2
L4:
	;
	if v6 <= v30 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v125 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v37 = int32(1)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v22))))
	if base.Ui32(v39) < base.Ui32(int32(192)) {
		v96 = v39
		v97 = v37
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(281) < v96 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v43 = v30 + int32(1)
	if v43 == v6 {
		v96 = v39
		v97 = v37
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v22))))
	v48 = v46 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v39) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v22))))
	v64 = v62 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v39) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v52 = v30 + int32(2)
	if v52 != v6 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v96 = v39<<(uint(int32(6))%32)&int32(1984) | v48
	v97 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v68))))
	v96 = v81&int32(63) | (v39<<(uint(int32(18))%32)&int32(_a_F_polish_UTF_8_stem_0) | v48<<(uint(int32(12))%32) | v64<<(uint(int32(6))%32))
	v97 = int32(4)
	goto L9
L18:
	;
	v68 = v30 + int32(3)
	if v68 != v6 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v96 = v39<<(uint(int32(12))%32)&int32(_a_F_polish_UTF_8_stem_1) | v48<<(uint(int32(6))%32) | v64
	v97 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v114 = v97 + v30
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
	v30 = v114
	goto L4
L23:
	;
	v101 = v96 - int32(97)
	if v101 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(3))%32)))+uint32(_c_F_polish_UTF_8_stem[0]))))
	if int32(base.Ui32(v107)>>(uint(v101&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v129 = v128 + v125
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v129
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = v129
	goto L30
L28:
	;
	if v248 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v248 = v219
	goto L28
L30:
	;
	if v143 <= v152 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v248 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v159 = int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v144))))
	if base.Ui32(v161) < base.Ui32(int32(192)) {
		v218 = v161
		v219 = v159
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(281) < v218 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v165 = v152 + int32(1)
	if v165 == v143 {
		v218 = v161
		v219 = v159
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v144))))
	v170 = v168 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v161) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v144))))
	v186 = v184 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v161) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v174 = v152 + int32(2)
	if v174 != v143 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v218 = v161<<(uint(int32(6))%32)&int32(1984) | v170
	v219 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v190))))
	v218 = v203&int32(63) | (v161<<(uint(int32(18))%32)&int32(_a_F_polish_UTF_8_stem_0) | v170<<(uint(int32(12))%32) | v186<<(uint(int32(6))%32))
	v219 = int32(4)
	goto L35
L44:
	;
	v190 = v152 + int32(3)
	if v190 != v143 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v218 = v161<<(uint(int32(12))%32)&int32(_a_F_polish_UTF_8_stem_1) | v170<<(uint(int32(6))%32) | v186
	v219 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v223 = v218 - int32(97)
	if v223 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v223)>>(uint(int32(3))%32)))+uint32(_c_F_polish_UTF_8_stem[0]))))
	if int32(base.Ui32(v229)>>(uint(v223&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v237 = v219 + v152
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v237
	v152 = v237
	goto L30
L52:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v251 + v248
	goto L1
L53:
	;
	return v465
L54:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v462
	v465 = int32(1)
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v417
	v420 = int32(0)
	v424 = F_find_among_b(m, l0, int32(_a_F_polish_UTF_8_stem_2), int32(4), v420)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L79
	} else {
		goto L113
	}
L56:
	;
	if v309 < int32(0) {
		goto L55
	} else {
		goto L76
	}
L58:
	;
	goto L59
L59:
	;
	goto L60
L60:
	;
	v264 = v8
	v266 = int32(2)
	goto L63
L62:
	;
	v309 = v294
	goto L56
L63:
	;
	if v257 <= v264 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L62
L65:
	;
	v309 = int32(-1)
	goto L56
L66:
	;
	goto L67
L67:
	;
	v271 = v264 + int32(1)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256+v264))))
	if base.Ui32(v273) < base.Ui32(int32(192)) {
		v294 = v271
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v295 = int32(1)
	if v295 < v266 {
		v264 = v294
		v266 = v266 - v295
		goto L63
	} else {
		goto L75
	}
L69:
	;
	if v257 <= v271 {
		v294 = v271
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v280 = v271
	goto L71
L71:
	;
	v283 = int32(*(*int8)(unsafe.Add(mBase, uint32(v256+v280))))
	if int32(-65) < v283 {
		v294 = v280
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v294 = v257
	goto L68
L73:
	;
	v287 = v280 + int32(1)
	if v287 != v257 {
		v280 = v287
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L64
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v309
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v313 < v315 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v336
	v342 = F_find_among_b(m, l0, int32(_a_F_polish_UTF_8_stem_3), int32(118), int32(_a_F_polish_UTF_8_stem_4))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L79
	} else {
		goto L85
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v315
	v322 = F_find_among_b(m, l0, int32(_a_F_polish_UTF_8_stem_5), int32(5), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	return int32(0)
L80:
	;
	if v322 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v309
	goto L77
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v309
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v330
	v332 = F_slice_del(m, l0)
	mBase = m.M
	if v332 < int32(0) {
		v465 = v332
		goto L53
	} else {
		goto L84
	}
L84:
	;
	goto L77
L85:
	;
	if v342 == int32(0) {
		goto L55
	} else {
		goto L86
	}
L86:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v346
	switch v342 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	default:
		goto L54
	}
L87:
	;
	v376 = F_slice_del(m, l0)
	mBase = m.M
	if v376 < int32(0) {
		v465 = v376
		goto L53
	} else {
		goto L103
	}
L88:
	;
	v372 = F_slice_from_s(m, l0, int32(2), int32(_a_F_polish_UTF_8_stem_6))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L79
	} else {
		goto L101
	}
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v359 <= v346 {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v355 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_UTF_8_stem_7))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L79
	} else {
		goto L93
	}
L91:
	;
	v350 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v350 {
		goto L54
	} else {
		goto L92
	}
L92:
	;
	v465 = v350
	goto L53
L93:
	;
	if int32(0) <= v355 {
		goto L54
	} else {
		goto L94
	}
L94:
	;
	v465 = v355
	goto L53
L95:
	;
	v361 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v361 {
		goto L54
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v366 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_UTF_8_stem_8))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L79
	} else {
		goto L99
	}
L98:
	;
	v465 = v361
	goto L53
L99:
	;
	if int32(0) <= v366 {
		goto L54
	} else {
		goto L100
	}
L100:
	;
	v465 = v366
	goto L53
L101:
	;
	if int32(0) <= v372 {
		goto L54
	} else {
		goto L102
	}
L102:
	;
	v465 = v372
	goto L53
L103:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v379
	v382 = v379 - int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v382 <= v383 {
		goto L54
	} else {
		goto L104
	}
L104:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v382))))
	if base.B2i32(v387 != int32(122))&base.B2i32(v387 != int32(99)) != 0 {
		goto L54
	} else {
		goto L105
	}
L105:
	;
	v396 = F_find_among_b(m, l0, int32(_a_F_polish_UTF_8_stem_9), int32(5), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L79
	} else {
		goto L106
	}
L106:
	;
	if v396 == int32(0) {
		goto L54
	} else {
		goto L107
	}
L107:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v400
	switch v396 - int32(1) {
	case 0:
		goto L109
	case 1:
		goto L108
	default:
		goto L54
	}
L108:
	;
	v409 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_UTF_8_stem_10))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L79
	} else {
		goto L111
	}
L109:
	;
	v404 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v404 {
		goto L54
	} else {
		goto L110
	}
L110:
	;
	v465 = v404
	goto L53
L111:
	;
	if int32(0) <= v409 {
		goto L54
	} else {
		goto L112
	}
L112:
	;
	v465 = v409
	goto L53
L113:
	;
	if v424 == int32(0) {
		v465 = v420
		goto L53
	} else {
		goto L114
	}
L114:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v428 <= v430 {
		v465 = v420
		goto L53
	} else {
		goto L115
	}
L115:
	;
	switch v424 - int32(1) {
	case 0:
		goto L119
	case 1:
		goto L118
	case 2:
		goto L117
	case 3:
		goto L116
	default:
		goto L54
	}
L116:
	;
	v454 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_UTF_8_stem_11))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L79
	} else {
		goto L126
	}
L117:
	;
	v448 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_UTF_8_stem_12))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L79
	} else {
		goto L124
	}
L118:
	;
	v442 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_UTF_8_stem_13))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L79
	} else {
		goto L122
	}
L119:
	;
	v436 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_UTF_8_stem_14))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L79
	} else {
		goto L120
	}
L120:
	;
	if int32(0) <= v436 {
		goto L54
	} else {
		goto L121
	}
L121:
	;
	v465 = v436
	goto L53
L122:
	;
	if int32(0) <= v442 {
		goto L54
	} else {
		goto L123
	}
L123:
	;
	v465 = v442
	goto L53
L124:
	;
	if int32(0) <= v448 {
		goto L54
	} else {
		goto L125
	}
L125:
	;
	v465 = v448
	goto L53
L126:
	;
	if v454 < int32(0) {
		v465 = v454
		goto L53
	} else {
		goto L127
	}
L127:
	;
	goto L54
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v316 int32
	_ = v316
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v435 int32
	_ = v435
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v479 int32
	_ = v479
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v556 int32
	_ = v556
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v674 int32
	_ = v674
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v719 int32
	_ = v719
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v797 int32
	_ = v797
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v915 int32
	_ = v915
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v941 int32
	_ = v941
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v983 int32
	_ = v983
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1034 int32
	_ = v1034
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	var v1078 int32
	_ = v1078
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1155 int32
	_ = v1155
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1289 int32
	_ = v1289
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1340 int32
	_ = v1340
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1366 int32
	_ = v1366
	var v1373 int32
	_ = v1373
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1411 int32
	_ = v1411
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1462 int32
	_ = v1462
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1488 int32
	_ = v1488
	var v1496 int32
	_ = v1496
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1535 int32
	_ = v1535
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1586 int32
	_ = v1586
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1612 int32
	_ = v1612
	var v1619 int32
	_ = v1619
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1657 int32
	_ = v1657
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1708 int32
	_ = v1708
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1742 int32
	_ = v1742
	var v1753 int32
	_ = v1753
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2021 int32
	_ = v2021
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2088 int32
	_ = v2088
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2257 int32
	_ = v2257
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v2269
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v121
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L46
L4:
	;
	goto L3
L5:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v119
	goto L2
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L21
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v57 = v9
	v59 = v17
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
	v34 = F_find_among(m, l0, int32(_a_F_portuguese_UTF_8_stem_0), int32(3), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v38
	switch v34 - int32(1) {
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
	v51 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v45 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v57 = v38
	v59 = v42
	goto L6
L15:
	;
	if int32(0) <= v45 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v2269 = v45
	goto L1
L17:
	;
	if int32(0) <= v51 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v2269 = v51
	goto L1
L19:
	;
	if v112 < int32(0) {
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
	v67 = v57
	v69 = int32(1)
	goto L26
L25:
	;
	v112 = v97
	goto L19
L26:
	;
	if v59 <= v67 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v112 = int32(-1)
	goto L19
L29:
	;
	goto L30
L30:
	;
	v74 = v67 + int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v67))))
	if base.Ui32(v76) < base.Ui32(int32(192)) {
		v97 = v74
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v98 = int32(1)
	if v98 < v69 {
		v67 = v97
		v69 = v69 - v98
		goto L26
	} else {
		goto L38
	}
L32:
	;
	if v59 <= v74 {
		v97 = v74
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v83 = v74
	goto L34
L34:
	;
	v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60+v83))))
	if int32(-65) < v86 {
		v97 = v83
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v97 = v59
	goto L31
L36:
	;
	v90 = v83 + int32(1)
	if v90 != v59 {
		v83 = v90
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v112
	goto L5
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1289 = v7
	goto L302
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1262
	goto L40
L42:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1262 = v1260 + v1258
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L175
L44:
	;
	if v242 != 0 {
		goto L43
	} else {
		goto L68
	}
L45:
	;
	v242 = v235
	goto L44
L46:
	;
	if v121 <= v7 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v235 = int32(0)
	goto L45
L48:
	;
	v242 = int32(-1)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v153 = int32(1)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v138))))
	if base.Ui32(v155) < base.Ui32(int32(192)) {
		v212 = v155
		v213 = v153
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if int32(250) < v212 {
		v235 = v213
		goto L45
	} else {
		goto L64
	}
L52:
	;
	v159 = v7 + int32(1)
	if v159 == v121 {
		v212 = v155
		v213 = v153
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v138))))
	v164 = v162 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v155) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v138))))
	v180 = v178 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v155) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v168 = v7 + int32(2)
	if v168 != v121 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v212 = v155<<(uint(int32(6))%32)&int32(1984) | v164
	v213 = int32(2)
	goto L51
L58:
	;
	goto L57
L59:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+v184))))
	v212 = v197&int32(63) | (v155<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v164<<(uint(int32(12))%32) | v180<<(uint(int32(6))%32))
	v213 = int32(4)
	goto L51
L60:
	;
	v184 = v7 + int32(3)
	if v184 != v121 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v212 = v155<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v164<<(uint(int32(6))%32) | v180
	v213 = int32(3)
	goto L51
L63:
	;
	goto L62
L64:
	;
	v217 = v212 - int32(97)
	if v217 < int32(0) {
		v235 = v213
		goto L45
	} else {
		goto L65
	}
L65:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v217)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v223)>>(uint(v217&int32(7))%32))&int32(1) == int32(0) {
		v235 = v213
		goto L45
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v213 + v7
	goto L67
L67:
	;
	goto L47
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L71
L69:
	;
	if v360 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L70:
	;
	v360 = v353
	goto L69
L71:
	;
	if v256 <= v243 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v353 = int32(0)
	goto L70
L73:
	;
	v360 = int32(-1)
	goto L69
L74:
	;
	goto L75
L75:
	;
	v272 = int32(1)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v257))))
	if base.Ui32(v274) < base.Ui32(int32(192)) {
		v331 = v274
		v332 = v272
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if int32(250) < v331 {
		goto L89
	} else {
		goto L90
	}
L77:
	;
	v278 = v243 + int32(1)
	if v278 == v256 {
		v331 = v274
		v332 = v272
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v257))))
	v283 = v281 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v274) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v257))))
	v299 = v297 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v274) {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v287 = v243 + int32(2)
	if v287 != v256 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v331 = v274<<(uint(int32(6))%32)&int32(1984) | v283
	v332 = int32(2)
	goto L76
L83:
	;
	goto L82
L84:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257+v303))))
	v331 = v316&int32(63) | (v274<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v283<<(uint(int32(12))%32) | v299<<(uint(int32(6))%32))
	v332 = int32(4)
	goto L76
L85:
	;
	v303 = v243 + int32(3)
	if v303 != v256 {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v331 = v274<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v283<<(uint(int32(6))%32) | v299
	v332 = int32(3)
	goto L76
L88:
	;
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v332 + v243
	goto L93
L90:
	;
	v336 = v331 - int32(97)
	if v336 < int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v336)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v342)>>(uint(v336&int32(7))%32))&int32(1) != 0 {
		v353 = v332
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
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v384 = v374
	goto L99
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v243
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L125
L97:
	;
	if int32(0) <= v479 {
		v1258 = v479
		goto L42
	} else {
		goto L122
	}
L98:
	;
	v479 = v451
	goto L97
L99:
	;
	if v375 <= v384 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v479 = int32(-1)
	goto L97
L102:
	;
	goto L103
L103:
	;
	v391 = int32(1)
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+v376))))
	if base.Ui32(v393) < base.Ui32(int32(192)) {
		v450 = v393
		v451 = v391
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if int32(250) < v450 {
		goto L117
	} else {
		goto L118
	}
L105:
	;
	v397 = v384 + int32(1)
	if v397 == v375 {
		v450 = v393
		v451 = v391
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v376))))
	v402 = v400 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v393) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v376))))
	v418 = v416 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v393) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v406 = v384 + int32(2)
	if v406 != v375 {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v450 = v393<<(uint(int32(6))%32)&int32(1984) | v402
	v451 = int32(2)
	goto L104
L111:
	;
	goto L110
L112:
	;
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v422))))
	v450 = v435&int32(63) | (v393<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v402<<(uint(int32(12))%32) | v418<<(uint(int32(6))%32))
	v451 = int32(4)
	goto L104
L113:
	;
	v422 = v384 + int32(3)
	if v422 != v375 {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v450 = v393<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v402<<(uint(int32(6))%32) | v418
	v451 = int32(3)
	goto L104
L116:
	;
	goto L115
L117:
	;
	v468 = v451 + v384
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v468
	v384 = v468
	goto L99
L118:
	;
	v455 = v450 - int32(97)
	if v455 < int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v455)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v461)>>(uint(v455&int32(7))%32))&int32(1) != 0 {
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
	if v601 != 0 {
		goto L43
	} else {
		goto L147
	}
L124:
	;
	v601 = v594
	goto L123
L125:
	;
	if v496 <= v243 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v594 = int32(0)
	goto L124
L127:
	;
	v601 = int32(-1)
	goto L123
L128:
	;
	goto L129
L129:
	;
	v512 = int32(1)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v497))))
	if base.Ui32(v514) < base.Ui32(int32(192)) {
		v571 = v514
		v572 = v512
		goto L130
	} else {
		goto L131
	}
L130:
	;
	if int32(250) < v571 {
		v594 = v572
		goto L124
	} else {
		goto L143
	}
L131:
	;
	v518 = v243 + int32(1)
	if v518 == v496 {
		v571 = v514
		v572 = v512
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518+v497))))
	v523 = v521 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v514) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527+v497))))
	v539 = v537 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v514) {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v527 = v243 + int32(2)
	if v527 != v496 {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v571 = v514<<(uint(int32(6))%32)&int32(1984) | v523
	v572 = int32(2)
	goto L130
L137:
	;
	goto L136
L138:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+v543))))
	v571 = v556&int32(63) | (v514<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v523<<(uint(int32(12))%32) | v539<<(uint(int32(6))%32))
	v572 = int32(4)
	goto L130
L139:
	;
	v543 = v243 + int32(3)
	if v543 != v496 {
		goto L138
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v571 = v514<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v523<<(uint(int32(6))%32) | v539
	v572 = int32(3)
	goto L130
L142:
	;
	goto L141
L143:
	;
	v576 = v571 - int32(97)
	if v576 < int32(0) {
		v594 = v572
		goto L124
	} else {
		goto L144
	}
L144:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v576)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v582)>>(uint(v576&int32(7))%32))&int32(1) == int32(0) {
		v594 = v572
		goto L124
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v572 + v243
	goto L146
L146:
	;
	goto L126
L147:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v623 = v613
	goto L150
L148:
	;
	if int32(0) <= v719 {
		v1258 = v719
		goto L42
	} else {
		goto L172
	}
L149:
	;
	v719 = v690
	goto L148
L150:
	;
	if v614 <= v623 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v719 = int32(-1)
	goto L148
L153:
	;
	goto L154
L154:
	;
	v630 = int32(1)
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623+v615))))
	if base.Ui32(v632) < base.Ui32(int32(192)) {
		v689 = v632
		v690 = v630
		goto L155
	} else {
		goto L156
	}
L155:
	;
	if int32(250) < v689 {
		goto L149
	} else {
		goto L168
	}
L156:
	;
	v636 = v623 + int32(1)
	if v636 == v614 {
		v689 = v632
		v690 = v630
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636+v615))))
	v641 = v639 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v632) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645+v615))))
	v657 = v655 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v632) {
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v645 = v623 + int32(2)
	if v645 != v614 {
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v689 = v632<<(uint(int32(6))%32)&int32(1984) | v641
	v690 = int32(2)
	goto L155
L162:
	;
	goto L161
L163:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615+v661))))
	v689 = v674&int32(63) | (v632<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v641<<(uint(int32(12))%32) | v657<<(uint(int32(6))%32))
	v690 = int32(4)
	goto L155
L164:
	;
	v661 = v623 + int32(3)
	if v661 != v614 {
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v689 = v632<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v641<<(uint(int32(6))%32) | v657
	v690 = int32(3)
	goto L155
L167:
	;
	goto L166
L168:
	;
	v694 = v689 - int32(97)
	if v694 < int32(0) {
		goto L149
	} else {
		goto L169
	}
L169:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v694)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v700)>>(uint(v694&int32(7))%32))&int32(1) == int32(0) {
		goto L149
	} else {
		goto L170
	}
L170:
	;
	v708 = v690 + v623
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v708
	v623 = v708
	goto L150
L172:
	;
	goto L43
L173:
	;
	if v841 != 0 {
		goto L40
	} else {
		goto L198
	}
L174:
	;
	v841 = v834
	goto L173
L175:
	;
	if v737 <= v7 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v834 = int32(0)
	goto L174
L177:
	;
	v841 = int32(-1)
	goto L173
L178:
	;
	goto L179
L179:
	;
	v753 = int32(1)
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v738))))
	if base.Ui32(v755) < base.Ui32(int32(192)) {
		v812 = v755
		v813 = v753
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if int32(250) < v812 {
		goto L193
	} else {
		goto L194
	}
L181:
	;
	v759 = v7 + int32(1)
	if v759 == v737 {
		v812 = v755
		v813 = v753
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759+v738))))
	v764 = v762 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v755) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768+v738))))
	v780 = v778 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v755) {
		goto L189
	} else {
		goto L190
	}
L184:
	;
	v768 = v7 + int32(2)
	if v768 != v737 {
		goto L183
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v812 = v755<<(uint(int32(6))%32)&int32(1984) | v764
	v813 = int32(2)
	goto L180
L187:
	;
	goto L186
L188:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738+v784))))
	v812 = v797&int32(63) | (v755<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v764<<(uint(int32(12))%32) | v780<<(uint(int32(6))%32))
	v813 = int32(4)
	goto L180
L189:
	;
	v784 = v7 + int32(3)
	if v784 != v737 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v812 = v755<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v764<<(uint(int32(6))%32) | v780
	v813 = int32(3)
	goto L180
L192:
	;
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 + v7
	goto L197
L194:
	;
	v817 = v812 - int32(97)
	if v817 < int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v817)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v823)>>(uint(v817&int32(7))%32))&int32(1) != 0 {
		v834 = v813
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
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L201
L199:
	;
	if v959 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L200:
	;
	v959 = v952
	goto L199
L201:
	;
	if v855 <= v842 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v952 = int32(0)
	goto L200
L203:
	;
	v959 = int32(-1)
	goto L199
L204:
	;
	goto L205
L205:
	;
	v871 = int32(1)
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842+v856))))
	if base.Ui32(v873) < base.Ui32(int32(192)) {
		v930 = v873
		v931 = v871
		goto L206
	} else {
		goto L207
	}
L206:
	;
	if int32(250) < v930 {
		goto L219
	} else {
		goto L220
	}
L207:
	;
	v877 = v842 + int32(1)
	if v877 == v855 {
		v930 = v873
		v931 = v871
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877+v856))))
	v882 = v880 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v873) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886+v856))))
	v898 = v896 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v873) {
		goto L215
	} else {
		goto L216
	}
L210:
	;
	v886 = v842 + int32(2)
	if v886 != v855 {
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v930 = v873<<(uint(int32(6))%32)&int32(1984) | v882
	v931 = int32(2)
	goto L206
L213:
	;
	goto L212
L214:
	;
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856+v902))))
	v930 = v915&int32(63) | (v873<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v882<<(uint(int32(12))%32) | v898<<(uint(int32(6))%32))
	v931 = int32(4)
	goto L206
L215:
	;
	v902 = v842 + int32(3)
	if v902 != v855 {
		goto L214
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v930 = v873<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v882<<(uint(int32(6))%32) | v898
	v931 = int32(3)
	goto L206
L218:
	;
	goto L217
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v931 + v842
	goto L223
L220:
	;
	v935 = v930 - int32(97)
	if v935 < int32(0) {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v935)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v941)>>(uint(v935&int32(7))%32))&int32(1) != 0 {
		v952 = v931
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
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v983 = v973
	goto L229
L225:
	;
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v842
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L255
L227:
	;
	if int32(0) <= v1078 {
		v1258 = v1078
		goto L42
	} else {
		goto L252
	}
L228:
	;
	v1078 = v1050
	goto L227
L229:
	;
	if v974 <= v983 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1078 = int32(-1)
	goto L227
L232:
	;
	goto L233
L233:
	;
	v990 = int32(1)
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983+v975))))
	if base.Ui32(v992) < base.Ui32(int32(192)) {
		v1049 = v992
		v1050 = v990
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if int32(250) < v1049 {
		goto L247
	} else {
		goto L248
	}
L235:
	;
	v996 = v983 + int32(1)
	if v996 == v974 {
		v1049 = v992
		v1050 = v990
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996+v975))))
	v1001 = v999 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v992) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005+v975))))
	v1017 = v1015 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v992) {
		goto L243
	} else {
		goto L244
	}
L238:
	;
	v1005 = v983 + int32(2)
	if v1005 != v974 {
		goto L237
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1049 = v992<<(uint(int32(6))%32)&int32(1984) | v1001
	v1050 = int32(2)
	goto L234
L241:
	;
	goto L240
L242:
	;
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975+v1021))))
	v1049 = v1034&int32(63) | (v992<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1001<<(uint(int32(12))%32) | v1017<<(uint(int32(6))%32))
	v1050 = int32(4)
	goto L234
L243:
	;
	v1021 = v983 + int32(3)
	if v1021 != v974 {
		goto L242
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1049 = v992<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1001<<(uint(int32(6))%32) | v1017
	v1050 = int32(3)
	goto L234
L246:
	;
	goto L245
L247:
	;
	v1067 = v1050 + v983
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1067
	v983 = v1067
	goto L229
L248:
	;
	v1054 = v1049 - int32(97)
	if v1054 < int32(0) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1054)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1060)>>(uint(v1054&int32(7))%32))&int32(1) != 0 {
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
	if v1200 != 0 {
		goto L40
	} else {
		goto L277
	}
L254:
	;
	v1200 = v1193
	goto L253
L255:
	;
	if v1095 <= v842 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1193 = int32(0)
	goto L254
L257:
	;
	v1200 = int32(-1)
	goto L253
L258:
	;
	goto L259
L259:
	;
	v1111 = int32(1)
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842+v1096))))
	if base.Ui32(v1113) < base.Ui32(int32(192)) {
		v1170 = v1113
		v1171 = v1111
		goto L260
	} else {
		goto L261
	}
L260:
	;
	if int32(250) < v1170 {
		v1193 = v1171
		goto L254
	} else {
		goto L273
	}
L261:
	;
	v1117 = v842 + int32(1)
	if v1117 == v1095 {
		v1170 = v1113
		v1171 = v1111
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1117+v1096))))
	v1122 = v1120 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1113) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126+v1096))))
	v1138 = v1136 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1113) {
		goto L269
	} else {
		goto L270
	}
L264:
	;
	v1126 = v842 + int32(2)
	if v1126 != v1095 {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1170 = v1113<<(uint(int32(6))%32)&int32(1984) | v1122
	v1171 = int32(2)
	goto L260
L267:
	;
	goto L266
L268:
	;
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096+v1142))))
	v1170 = v1155&int32(63) | (v1113<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1122<<(uint(int32(12))%32) | v1138<<(uint(int32(6))%32))
	v1171 = int32(4)
	goto L260
L269:
	;
	v1142 = v842 + int32(3)
	if v1142 != v1095 {
		goto L268
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1170 = v1113<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1122<<(uint(int32(6))%32) | v1138
	v1171 = int32(3)
	goto L260
L272:
	;
	goto L271
L273:
	;
	v1175 = v1170 - int32(97)
	if v1175 < int32(0) {
		v1193 = v1171
		goto L254
	} else {
		goto L274
	}
L274:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1175)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1181)>>(uint(v1175&int32(7))%32))&int32(1) == int32(0) {
		v1193 = v1171
		goto L254
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1171 + v842
	goto L276
L276:
	;
	goto L256
L277:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L280
L278:
	;
	if int32(0) <= v1255 {
		v1262 = v1255
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
	v1210 = v1202
	v1212 = int32(1)
	goto L285
L284:
	;
	v1255 = v1240
	goto L278
L285:
	;
	if v1203 <= v1210 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	goto L284
L287:
	;
	v1255 = int32(-1)
	goto L278
L288:
	;
	goto L289
L289:
	;
	v1217 = v1210 + int32(1)
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201+v1210))))
	if base.Ui32(v1219) < base.Ui32(int32(192)) {
		v1240 = v1217
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1241 = int32(1)
	if v1241 < v1212 {
		v1210 = v1240
		v1212 = v1212 - v1241
		goto L285
	} else {
		goto L297
	}
L291:
	;
	if v1203 <= v1217 {
		v1240 = v1217
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1226 = v1217
	goto L293
L293:
	;
	v1229 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1201+v1226))))
	if int32(-65) < v1229 {
		v1240 = v1226
		goto L290
	} else {
		goto L295
	}
L294:
	;
	v1240 = v1203
	goto L290
L295:
	;
	v1233 = v1226 + int32(1)
	if v1233 != v1203 {
		v1226 = v1233
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1761
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1761
	if v1761-int32(2) <= v7 {
		goto L404
	} else {
		goto L405
	}
L300:
	;
	if v1384 < int32(0) {
		goto L299
	} else {
		goto L325
	}
L301:
	;
	v1384 = v1356
	goto L300
L302:
	;
	if v1280 <= v1289 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1384 = int32(-1)
	goto L300
L305:
	;
	goto L306
L306:
	;
	v1296 = int32(1)
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1289+v1281))))
	if base.Ui32(v1298) < base.Ui32(int32(192)) {
		v1355 = v1298
		v1356 = v1296
		goto L307
	} else {
		goto L308
	}
L307:
	;
	if int32(250) < v1355 {
		goto L320
	} else {
		goto L321
	}
L308:
	;
	v1302 = v1289 + int32(1)
	if v1302 == v1280 {
		v1355 = v1298
		v1356 = v1296
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302+v1281))))
	v1307 = v1305 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1298) {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v1321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311+v1281))))
	v1323 = v1321 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1298) {
		goto L316
	} else {
		goto L317
	}
L311:
	;
	v1311 = v1289 + int32(2)
	if v1311 != v1280 {
		goto L310
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1355 = v1298<<(uint(int32(6))%32)&int32(1984) | v1307
	v1356 = int32(2)
	goto L307
L314:
	;
	goto L313
L315:
	;
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1281+v1327))))
	v1355 = v1340&int32(63) | (v1298<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1307<<(uint(int32(12))%32) | v1323<<(uint(int32(6))%32))
	v1356 = int32(4)
	goto L307
L316:
	;
	v1327 = v1289 + int32(3)
	if v1327 != v1280 {
		goto L315
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v1355 = v1298<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1307<<(uint(int32(6))%32) | v1323
	v1356 = int32(3)
	goto L307
L319:
	;
	goto L318
L320:
	;
	v1373 = v1356 + v1289
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1373
	v1289 = v1373
	goto L302
L321:
	;
	v1360 = v1355 - int32(97)
	if v1360 < int32(0) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1360)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1366)>>(uint(v1360&int32(7))%32))&int32(1) != 0 {
		goto L301
	} else {
		goto L323
	}
L323:
	;
	goto L320
L325:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1388 = v1387 + v1384
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1388
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1411 = v1388
	goto L328
L326:
	;
	if v1507 < int32(0) {
		goto L299
	} else {
		goto L350
	}
L327:
	;
	v1507 = v1478
	goto L326
L328:
	;
	if v1402 <= v1411 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1507 = int32(-1)
	goto L326
L331:
	;
	goto L332
L332:
	;
	v1418 = int32(1)
	v1420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1411+v1403))))
	if base.Ui32(v1420) < base.Ui32(int32(192)) {
		v1477 = v1420
		v1478 = v1418
		goto L333
	} else {
		goto L334
	}
L333:
	;
	if int32(250) < v1477 {
		goto L327
	} else {
		goto L346
	}
L334:
	;
	v1424 = v1411 + int32(1)
	if v1424 == v1402 {
		v1477 = v1420
		v1478 = v1418
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1424+v1403))))
	v1429 = v1427 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1420) {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1433+v1403))))
	v1445 = v1443 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1420) {
		goto L342
	} else {
		goto L343
	}
L337:
	;
	v1433 = v1411 + int32(2)
	if v1433 != v1402 {
		goto L336
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1477 = v1420<<(uint(int32(6))%32)&int32(1984) | v1429
	v1478 = int32(2)
	goto L333
L340:
	;
	goto L339
L341:
	;
	v1462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403+v1449))))
	v1477 = v1462&int32(63) | (v1420<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1429<<(uint(int32(12))%32) | v1445<<(uint(int32(6))%32))
	v1478 = int32(4)
	goto L333
L342:
	;
	v1449 = v1411 + int32(3)
	if v1449 != v1402 {
		goto L341
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1477 = v1420<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1429<<(uint(int32(6))%32) | v1445
	v1478 = int32(3)
	goto L333
L345:
	;
	goto L344
L346:
	;
	v1482 = v1477 - int32(97)
	if v1482 < int32(0) {
		goto L327
	} else {
		goto L347
	}
L347:
	;
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1482)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1488)>>(uint(v1482&int32(7))%32))&int32(1) == int32(0) {
		goto L327
	} else {
		goto L348
	}
L348:
	;
	v1496 = v1478 + v1411
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1496
	v1411 = v1496
	goto L328
L350:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1511 = v1510 + v1507
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1511
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1511
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1535 = v1511
	goto L353
L351:
	;
	if v1630 < int32(0) {
		goto L299
	} else {
		goto L376
	}
L352:
	;
	v1630 = v1602
	goto L351
L353:
	;
	if v1526 <= v1535 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1630 = int32(-1)
	goto L351
L356:
	;
	goto L357
L357:
	;
	v1542 = int32(1)
	v1544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535+v1527))))
	if base.Ui32(v1544) < base.Ui32(int32(192)) {
		v1601 = v1544
		v1602 = v1542
		goto L358
	} else {
		goto L359
	}
L358:
	;
	if int32(250) < v1601 {
		goto L371
	} else {
		goto L372
	}
L359:
	;
	v1548 = v1535 + int32(1)
	if v1548 == v1526 {
		v1601 = v1544
		v1602 = v1542
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548+v1527))))
	v1553 = v1551 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1544) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1557+v1527))))
	v1569 = v1567 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1544) {
		goto L367
	} else {
		goto L368
	}
L362:
	;
	v1557 = v1535 + int32(2)
	if v1557 != v1526 {
		goto L361
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1601 = v1544<<(uint(int32(6))%32)&int32(1984) | v1553
	v1602 = int32(2)
	goto L358
L365:
	;
	goto L364
L366:
	;
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527+v1573))))
	v1601 = v1586&int32(63) | (v1544<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1553<<(uint(int32(12))%32) | v1569<<(uint(int32(6))%32))
	v1602 = int32(4)
	goto L358
L367:
	;
	v1573 = v1535 + int32(3)
	if v1573 != v1526 {
		goto L366
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1601 = v1544<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1553<<(uint(int32(6))%32) | v1569
	v1602 = int32(3)
	goto L358
L370:
	;
	goto L369
L371:
	;
	v1619 = v1602 + v1535
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1619
	v1535 = v1619
	goto L353
L372:
	;
	v1606 = v1601 - int32(97)
	if v1606 < int32(0) {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1606)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1612)>>(uint(v1606&int32(7))%32))&int32(1) != 0 {
		goto L352
	} else {
		goto L374
	}
L374:
	;
	goto L371
L376:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1634 = v1633 + v1630
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1634
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1657 = v1634
	goto L379
L377:
	;
	if v1753 < int32(0) {
		goto L299
	} else {
		goto L401
	}
L378:
	;
	v1753 = v1724
	goto L377
L379:
	;
	if v1648 <= v1657 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1753 = int32(-1)
	goto L377
L382:
	;
	goto L383
L383:
	;
	v1664 = int32(1)
	v1666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1657+v1649))))
	if base.Ui32(v1666) < base.Ui32(int32(192)) {
		v1723 = v1666
		v1724 = v1664
		goto L384
	} else {
		goto L385
	}
L384:
	;
	if int32(250) < v1723 {
		goto L378
	} else {
		goto L397
	}
L385:
	;
	v1670 = v1657 + int32(1)
	if v1670 == v1648 {
		v1723 = v1666
		v1724 = v1664
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1670+v1649))))
	v1675 = v1673 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1666) {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	v1689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1679+v1649))))
	v1691 = v1689 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1666) {
		goto L393
	} else {
		goto L394
	}
L388:
	;
	v1679 = v1657 + int32(2)
	if v1679 != v1648 {
		goto L387
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	v1723 = v1666<<(uint(int32(6))%32)&int32(1984) | v1675
	v1724 = int32(2)
	goto L384
L391:
	;
	goto L390
L392:
	;
	v1708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1649+v1695))))
	v1723 = v1708&int32(63) | (v1666<<(uint(int32(18))%32)&int32(_a_F_portuguese_UTF_8_stem_3) | v1675<<(uint(int32(12))%32) | v1691<<(uint(int32(6))%32))
	v1724 = int32(4)
	goto L384
L393:
	;
	v1695 = v1657 + int32(3)
	if v1695 != v1648 {
		goto L392
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1723 = v1666<<(uint(int32(12))%32)&int32(_a_F_portuguese_UTF_8_stem_4) | v1675<<(uint(int32(6))%32) | v1691
	v1724 = int32(3)
	goto L384
L396:
	;
	goto L395
L397:
	;
	v1728 = v1723 - int32(97)
	if v1728 < int32(0) {
		goto L378
	} else {
		goto L398
	}
L398:
	;
	v1734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1728)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_UTF_8_stem[0]))))
	if int32(base.Ui32(v1734)>>(uint(v1728&int32(7))%32))&int32(1) == int32(0) {
		goto L378
	} else {
		goto L399
	}
L399:
	;
	v1742 = v1724 + v1657
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1742
	v1657 = v1742
	goto L379
L401:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1756 + v1753
	goto L299
L402:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2088
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2088
	v2094 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_5), int32(4), int32(0))
	mBase = m.M
	v2095 = m.ExcPending
	if v2095 != 0 {
		goto L10
	} else {
		goto L493
	}
L403:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2057
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2057
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2057 <= v2060 {
		goto L402
	} else {
		goto L486
	}
L404:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2016
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2018 <= v2016 {
		goto L476
	} else {
		goto L477
	}
L405:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1769 = int32(1)
	v1771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1767+v1761-v1769))))
	if base.B2i32(v1771&int32(224) != int32(96))|base.B2i32(v1769<<(uint(v1771)%32)&int32(_a_F_portuguese_UTF_8_stem_6) == int32(0)) != 0 {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1786 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_7), int32(45), int32(0))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L10
	} else {
		goto L407
	}
L407:
	;
	if v1786 == int32(0) {
		goto L404
	} else {
		goto L408
	}
L408:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1790
	switch v1786 - int32(1) {
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
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1790 < v1994 {
		goto L404
	} else {
		goto L470
	}
L410:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1790 < v1962 {
		goto L404
	} else {
		goto L461
	}
L411:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1790 < v1923 {
		goto L404
	} else {
		goto L453
	}
L412:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1790 < v1891 {
		goto L404
	} else {
		goto L445
	}
L413:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1790 < v1823 {
		goto L404
	} else {
		goto L429
	}
L414:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1790 < v1815 {
		goto L404
	} else {
		goto L426
	}
L415:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1790 < v1807 {
		goto L404
	} else {
		goto L423
	}
L416:
	;
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1790 < v1799 {
		goto L404
	} else {
		goto L420
	}
L417:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1790 < v1794 {
		goto L404
	} else {
		goto L418
	}
L418:
	;
	v1796 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1796 {
		goto L403
	} else {
		goto L419
	}
L419:
	;
	v2269 = v1796
	goto L1
L420:
	;
	v1803 = F_slice_from_s(m, l0, int32(3), int32(_a_F_portuguese_UTF_8_stem_8))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L10
	} else {
		goto L421
	}
L421:
	;
	if int32(0) <= v1803 {
		goto L403
	} else {
		goto L422
	}
L422:
	;
	v2269 = v1803
	goto L1
L423:
	;
	v1811 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_UTF_8_stem_9))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L10
	} else {
		goto L424
	}
L424:
	;
	if int32(0) <= v1811 {
		goto L403
	} else {
		goto L425
	}
L425:
	;
	v2269 = v1811
	goto L1
L426:
	;
	v1819 = F_slice_from_s(m, l0, int32(4), int32(_a_F_portuguese_UTF_8_stem_10))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L10
	} else {
		goto L427
	}
L427:
	;
	if int32(0) <= v1819 {
		goto L403
	} else {
		goto L428
	}
L428:
	;
	v2269 = v1819
	goto L1
L429:
	;
	v1825 = F_slice_del(m, l0)
	mBase = m.M
	if v1825 < int32(0) {
		v2269 = v1825
		goto L1
	} else {
		goto L430
	}
L430:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1828
	v1831 = v1828 - int32(1)
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1831 <= v1832 {
		goto L403
	} else {
		goto L431
	}
L431:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1834+v1831))))
	if base.B2i32(v1836&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1836)%32)&int32(_a_F_portuguese_UTF_8_stem_11) == int32(0)) != 0 {
		goto L403
	} else {
		goto L432
	}
L432:
	;
	v1851 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_12), int32(4), int32(0))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L10
	} else {
		goto L433
	}
L433:
	;
	if v1851 == int32(0) {
		goto L403
	} else {
		goto L434
	}
L434:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1855
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1855 < v1857 {
		goto L403
	} else {
		goto L435
	}
L435:
	;
	v1859 = F_slice_del(m, l0)
	mBase = m.M
	if v1859 < int32(0) {
		v2269 = v1859
		goto L1
	} else {
		goto L436
	}
L436:
	;
	if v1851 != int32(1) {
		goto L403
	} else {
		goto L437
	}
L437:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1864
	v1866 = int32(2)
	v1868 = int32(0)
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1864-v1871 < v1866 {
		v1881 = v1868
		goto L439
	} else {
		goto L440
	}
L438:
	;
	if v1881 == int32(0) {
		goto L403
	} else {
		goto L442
	}
L439:
	;
	goto L438
L440:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1877 = F_memcmp(m, v1874+v1864-v1866, int32(_a_F_portuguese_UTF_8_stem_13), v1866)
	mBase = m.M
	if v1877 != 0 {
		v1881 = v1868
		goto L439
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1864 - v1866
	v1881 = int32(1)
	goto L439
L442:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1884
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1884 < v1886 {
		goto L403
	} else {
		goto L443
	}
L443:
	;
	v1888 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1888 {
		goto L403
	} else {
		goto L444
	}
L444:
	;
	v2269 = v1888
	goto L1
L445:
	;
	v1893 = F_slice_del(m, l0)
	mBase = m.M
	if v1893 < int32(0) {
		v2269 = v1893
		goto L1
	} else {
		goto L446
	}
L446:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1896
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1896-int32(3) <= v1898 {
		goto L403
	} else {
		goto L447
	}
L447:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1902+v1896-int32(1)))))
	switch v1906 - int32(101) {
	case 0, 7:
		goto L448
	default:
		goto L403
	}
L448:
	;
	v1912 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_14), int32(3), int32(0))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L10
	} else {
		goto L449
	}
L449:
	;
	if v1912 == int32(0) {
		goto L403
	} else {
		goto L450
	}
L450:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1916
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1916 < v1918 {
		goto L403
	} else {
		goto L451
	}
L451:
	;
	v1920 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1920 {
		goto L403
	} else {
		goto L452
	}
L452:
	;
	v2269 = v1920
	goto L1
L453:
	;
	v1925 = F_slice_del(m, l0)
	mBase = m.M
	if v1925 < int32(0) {
		v2269 = v1925
		goto L1
	} else {
		goto L454
	}
L454:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1928
	v1931 = v1928 - int32(1)
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1931 <= v1932 {
		goto L403
	} else {
		goto L455
	}
L455:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1934+v1931))))
	if base.B2i32(v1936&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1936)%32)&int32(_a_F_portuguese_UTF_8_stem_15) == int32(0)) != 0 {
		goto L403
	} else {
		goto L456
	}
L456:
	;
	v1951 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_16), int32(3), int32(0))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L10
	} else {
		goto L457
	}
L457:
	;
	if v1951 == int32(0) {
		goto L403
	} else {
		goto L458
	}
L458:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1955
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1955 < v1957 {
		goto L403
	} else {
		goto L459
	}
L459:
	;
	v1959 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1959 {
		goto L403
	} else {
		goto L460
	}
L460:
	;
	v2269 = v1959
	goto L1
L461:
	;
	v1964 = F_slice_del(m, l0)
	mBase = m.M
	if v1964 < int32(0) {
		v2269 = v1964
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1967
	v1969 = int32(2)
	v1971 = int32(0)
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1967-v1974 < v1969 {
		v1984 = v1971
		goto L464
	} else {
		goto L465
	}
L463:
	;
	if v1984 == int32(0) {
		goto L403
	} else {
		goto L467
	}
L464:
	;
	goto L463
L465:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1980 = F_memcmp(m, v1977+v1967-v1969, int32(_a_F_portuguese_UTF_8_stem_17), v1969)
	mBase = m.M
	if v1980 != 0 {
		v1984 = v1971
		goto L464
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1967 - v1969
	v1984 = int32(1)
	goto L464
L467:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1987
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1987 < v1989 {
		goto L403
	} else {
		goto L468
	}
L468:
	;
	v1991 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1991 {
		goto L403
	} else {
		goto L469
	}
L469:
	;
	v2269 = v1991
	goto L1
L470:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1790 <= v1996 {
		goto L404
	} else {
		goto L471
	}
L471:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1998+v1790-int32(1)))))
	if v2002 != int32(101) {
		goto L404
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1790 - int32(1)
	v2010 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_18))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L10
	} else {
		goto L473
	}
L473:
	;
	if int32(0) <= v2010 {
		goto L403
	} else {
		goto L474
	}
L474:
	;
	v2269 = v2010
	goto L1
L475:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2048
	v2050 = F_slice_del(m, l0)
	mBase = m.M
	if v2050 < int32(0) {
		v2269 = v2050
		goto L1
	} else {
		goto L485
	}
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2016
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2018
	v2026 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_19), int32(120), int32(0))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L10
	} else {
		goto L479
	}
L477:
	;
	v2030 = v2016
	goto L478
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2030
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2030
	v2037 = F_find_among_b(m, l0, int32(_a_F_portuguese_UTF_8_stem_20), int32(7), int32(0))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L10
	} else {
		goto L481
	}
L479:
	;
	if v2026 != 0 {
		goto L475
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2021
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2030 = v2029
	goto L478
L481:
	;
	if v2037 == int32(0) {
		goto L402
	} else {
		goto L482
	}
L482:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2041
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2041 < v2043 {
		goto L402
	} else {
		goto L483
	}
L483:
	;
	v2045 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2045 {
		goto L402
	} else {
		goto L484
	}
L484:
	;
	v2269 = v2045
	goto L1
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2021
	goto L403
L486:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2063 = v2062 + v2057
	v2066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2063-int32(1)))))
	if v2066 != int32(105) {
		goto L402
	} else {
		goto L487
	}
L487:
	;
	v2070 = v2057 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2070
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2070
	if v2070 <= v2060 {
		goto L402
	} else {
		goto L488
	}
L488:
	;
	v2076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2063-int32(2)))))
	if v2076 != int32(99) {
		goto L402
	} else {
		goto L489
	}
L489:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2057 <= v2079 {
		goto L402
	} else {
		goto L490
	}
L490:
	;
	v2081 = F_slice_del(m, l0)
	mBase = m.M
	if v2081 < int32(0) {
		v2269 = v2081
		goto L1
	} else {
		goto L491
	}
L491:
	;
	goto L402
L492:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2161
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2165 = v2161
	v2168 = v2163
	goto L512
L493:
	;
	if v2094 == int32(0) {
		goto L492
	} else {
		goto L494
	}
L494:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2098
	switch v2094 - int32(1) {
	case 0:
		goto L496
	case 1:
		goto L495
	default:
		goto L492
	}
L495:
	;
	v2152 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_UTF_8_stem_21))
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L10
	} else {
		goto L510
	}
L496:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2098 < v2102 {
		goto L492
	} else {
		goto L497
	}
L497:
	;
	v2104 = F_slice_del(m, l0)
	mBase = m.M
	if v2104 < int32(0) {
		v2269 = v2104
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2107
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2107 <= v2109 {
		goto L492
	} else {
		goto L499
	}
L499:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2112 = v2111 + v2107
	v2114 = v2112 - int32(1)
	v2115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2114))))
	if v2115 != int32(117) {
		goto L501
	} else {
		goto L502
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2143
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2143 < v2145 {
		goto L492
	} else {
		goto L508
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2107
	v2130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2114))))
	if v2130 != int32(105) {
		goto L492
	} else {
		goto L505
	}
L502:
	;
	v2119 = v2107 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2119
	if v2119 <= v2109 {
		goto L501
	} else {
		goto L503
	}
L503:
	;
	v2125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2112-int32(2)))))
	if v2125 == int32(103) {
		v2143 = v2119
		goto L500
	} else {
		goto L504
	}
L504:
	;
	goto L501
L505:
	;
	v2134 = v2107 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2134
	if v2134 <= v2109 {
		goto L492
	} else {
		goto L506
	}
L506:
	;
	v2140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2112-int32(2)))))
	if v2140 != int32(99) {
		goto L492
	} else {
		goto L507
	}
L507:
	;
	v2143 = v2134
	goto L500
L508:
	;
	v2147 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2147 {
		goto L492
	} else {
		goto L509
	}
L509:
	;
	v2269 = v2147
	goto L1
L510:
	;
	if v2152 < int32(0) {
		v2269 = v2152
		goto L1
	} else {
		goto L511
	}
L511:
	;
	goto L492
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2165
	v2172 = v2165 + int32(1)
	if v2172 < v2168 {
		goto L518
	} else {
		goto L519
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2161
	v2269 = int32(1)
	goto L1
L514:
	;
	goto L513
L515:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2165 = v2265
	v2168 = v2264
	goto L512
L516:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L532
L517:
	;
	v2183 = F_find_among(m, l0, int32(_a_F_portuguese_UTF_8_stem_22), int32(3), int32(0))
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L10
	} else {
		goto L522
	}
L518:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2174+v2172))))
	if v2176 == int32(126) {
		goto L517
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2165
	v2202 = v2165
	v2204 = v2168
	goto L516
L521:
	;
	goto L520
L522:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2185
	switch v2183 - int32(1) {
	case 0:
		goto L525
	case 1:
		goto L524
	case 2:
		goto L523
	default:
		goto L515
	}
L523:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2202 = v2185
	v2204 = v2201
	goto L516
L524:
	;
	v2197 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_23))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L10
	} else {
		goto L528
	}
L525:
	;
	v2191 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_UTF_8_stem_24))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L10
	} else {
		goto L526
	}
L526:
	;
	if int32(0) <= v2191 {
		goto L515
	} else {
		goto L527
	}
L527:
	;
	v2269 = v2191
	goto L1
L528:
	;
	if int32(0) <= v2197 {
		goto L515
	} else {
		goto L529
	}
L529:
	;
	v2269 = v2197
	goto L1
L530:
	;
	if v2257 < int32(0) {
		goto L514
	} else {
		goto L550
	}
L532:
	;
	goto L533
L533:
	;
	goto L534
L534:
	;
	v2212 = v2202
	v2214 = int32(1)
	goto L537
L536:
	;
	v2257 = v2242
	goto L530
L537:
	;
	if v2204 <= v2212 {
		goto L539
	} else {
		goto L540
	}
L538:
	;
	goto L536
L539:
	;
	v2257 = int32(-1)
	goto L530
L540:
	;
	goto L541
L541:
	;
	v2219 = v2212 + int32(1)
	v2221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2205+v2212))))
	if base.Ui32(v2221) < base.Ui32(int32(192)) {
		v2242 = v2219
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v2243 = int32(1)
	if v2243 < v2214 {
		v2212 = v2242
		v2214 = v2214 - v2243
		goto L537
	} else {
		goto L549
	}
L543:
	;
	if v2204 <= v2219 {
		v2242 = v2219
		goto L542
	} else {
		goto L544
	}
L544:
	;
	v2228 = v2219
	goto L545
L545:
	;
	v2231 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2205+v2228))))
	if int32(-65) < v2231 {
		v2242 = v2228
		goto L542
	} else {
		goto L547
	}
L546:
	;
	v2242 = v2204
	goto L542
L547:
	;
	v2235 = v2228 + int32(1)
	if v2235 != v2204 {
		v2228 = v2235
		goto L545
	} else {
		goto L548
	}
L548:
	;
	goto L546
L549:
	;
	goto L538
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2257
	goto L515
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
func F_prefixsel(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14370(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
	var v165 int32
	_ = v165
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
	v9 = int32(4)
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
	if v29 == int32(4) {
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
	default:
		v141 = v33
		goto L11
	case 2, 3:
		goto L16
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
	v138 = int32(2)
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
	v86 = F_palloc0(m, int32(20))
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
	v63 = int32(3)
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
	v165 = F_errdetail(m, int32(_a_F_preprocess_pubobj_list_11), int32(0))
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
	var v58 int32
	_ = v58
	var v65 int64
	_ = v65
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
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
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
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
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	if l2 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v58 = v11 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = int32(_a_F_process_postgres_switches_0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(_a_F_process_postgres_switches_1)
	v65 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+20)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = v65
	goto L16
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
	v22 = int32(_a_F_process_postgres_switches_2)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[0])))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = int32(0)
	goto L22
L17:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v54+v615<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v619
	F_errmsg(m, int32(_a_F_process_postgres_switches_3), v11+int32(16))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L50
	} else {
		goto L196
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v141
	F_errmsg(m, int32(_a_F_process_postgres_switches_4), v11-int32(-64))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L50
	} else {
		goto L194
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L50
	} else {
		goto L190
	}
L20:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[1])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L50
	} else {
		goto L184
	}
L21:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v540 - int32(1)
	goto L20
L22:
	;
	v81 = F_pg_getopt_next(m, v11+int32(112))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L50
	} else {
		goto L51
	}
L23:
	;
	if l3 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L24:
	;
	goto L23
L25:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_5), v503, l2, v55)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L50
	} else {
		goto L174
	}
L26:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L157
	}
L27:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	switch v437 - int32(101) {
	case 0:
		v445 = int32(_a_F_process_postgres_switches_6)
		goto L153
	default:
		goto L21
	case 11:
		goto L154
	}
L28:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_7), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L50
	} else {
		goto L152
	}
L29:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_9), v428, l2, v55)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L50
	} else {
		goto L151
	}
L30:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L119
	}
L31:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_10), v301, l2, v55)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L50
	} else {
		goto L118
	}
L32:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_11), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L50
	} else {
		goto L117
	}
L33:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_12), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L50
	} else {
		goto L116
	}
L34:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_13), v289, l2, v55)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L50
	} else {
		goto L115
	}
L35:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_14), int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L50
	} else {
		goto L114
	}
L36:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_15), v281, l2, v55)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L50
	} else {
		goto L113
	}
L37:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L112
	}
L38:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_16), int32(_a_F_process_postgres_switches_17), l2, v55)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L50
	} else {
		goto L111
	}
L39:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_16), v268, l2, v55)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L50
	} else {
		goto L110
	}
L40:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v247 = v245 - int32(98)
	v249 = v247 & int32(255)
	if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v249))|base.B2i32(int32(base.Ui32(int32(_a_F_process_postgres_switches_18))>>(uint(v249)%32))&int32(1) == int32(0)) != 0 {
		goto L21
	} else {
		goto L108
	}
L41:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_19), int32(_a_F_process_postgres_switches_20), l2, v55)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L50
	} else {
		goto L107
	}
L42:
	;
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_21), int32(_a_F_process_postgres_switches_22), l2, v55)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L50
	} else {
		goto L106
	}
L43:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L105
	}
L44:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v184 = v180
	goto L89
L45:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L83
	}
L46:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_ParseLongOption(m, v124, v11+int32(108), v11+int32(104))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L50
	} else {
		goto L71
	}
L47:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v96 = F_strcmp(m, int32(_a_F_process_postgres_switches_23), v94)
	mBase = m.M
	if v96 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L53
	}
L49:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_SetConfigOption(m, int32(_a_F_process_postgres_switches_24), v86, l2, v55)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L50
	} else {
		goto L52
	}
L50:
	;
	return
L51:
	;
	switch v81 + int32(1) {
	case 0:
		goto L24
	default:
		goto L21
	case 46:
		goto L47
	case 67:
		goto L49
	case 68, 85, 111:
		goto L22
	case 69:
		goto L45
	case 70:
		goto L43
	case 71:
		goto L41
	case 79:
		goto L34
	case 80:
		goto L33
	case 81:
		goto L32
	case 84:
		goto L29
	case 88:
		goto L25
	case 99:
		goto L48
	case 100:
		goto L46
	case 101:
		goto L44
	case 102:
		goto L42
	case 103:
		goto L40
	case 105:
		goto L39
	case 106:
		goto L38
	case 107:
		goto L37
	case 108:
		goto L36
	case 109:
		goto L35
	case 113:
		goto L31
	case 115:
		goto L30
	case 116:
		goto L28
	case 117:
		goto L27
	case 119:
		goto L26
	}
L52:
	;
	goto L22
L53:
	;
	v92 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[2])) = uint8(v92)
	goto L22
L54:
	;
	if v121 != int32(5) {
		goto L19
	} else {
		goto L70
	}
L55:
	;
	v121 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v101 = F_strcmp(m, int32(_a_F_process_postgres_switches_25), v94)
	mBase = m.M
	if v101 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v121 = int32(1)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v107 = F_strncmp(m, int32(_a_F_process_postgres_switches_26), v94, int32(9))
	mBase = m.M
	if v107 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v121 = int32(2)
	goto L54
L62:
	;
	goto L63
L63:
	;
	v112 = F_strcmp(m, int32(_a_F_process_postgres_switches_27), v94)
	mBase = m.M
	if v112 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v121 = int32(3)
	goto L54
L65:
	;
	goto L66
L66:
	;
	v119 = F_strcmp(m, int32(_a_F_process_postgres_switches_28), v94)
	mBase = m.M
	if v119 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v120 = int32(5)
	goto L69
L68:
	;
	v120 = int32(4)
	goto L69
L69:
	;
	v121 = v120
	goto L54
L70:
	;
	goto L46
L71:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	if v131 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L50
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	F_SetConfigOption(m, v155, v131, l2, v55)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L50
	} else {
		goto L80
	}
L75:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L50
	} else {
		goto L76
	}
L76:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	if v81 == int32(45) {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v141
	F_errmsg(m, int32(_a_F_process_postgres_switches_29), v11+int32(80))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L50
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(4071), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L50
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
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	F_pfree(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L50
	} else {
		goto L81
	}
L81:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	F_pfree(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L50
	} else {
		goto L82
	}
L82:
	;
	goto L22
L83:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v170 = F_strlen(m, v167)
	mBase = m.M
	v172 = v170 + int32(1)
	v173 = F_emscripten_builtin_malloc(m, v172)
	mBase = m.M
	if v173 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[3])) = v178
	goto L22
L85:
	;
	v178 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v177 = F___memcpy(m, v173, v167, v172)
	mBase = m.M
	v178 = v177
	goto L84
L88:
	;
	F_set_debug_options(m, v228, l2, v55)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L50
	} else {
		goto L104
	}
L89:
	;
	v189 = v184 + int32(1)
	v190 = int32(*(*int8)(unsafe.Add(mBase, uint32(v184))))
	v191 = F___isspace(m, v190)
	mBase = m.M
	if v191 != 0 {
		v184 = v189
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v192 = int32(1)
	switch v190&int32(255) - int32(43) {
	case 0:
		v198 = v192
		goto L93
	default:
		v200 = v190
		v201 = v184
		v202 = v192
		goto L92
	case 2:
		goto L94
	}
L91:
	;
	goto L90
L92:
	;
	v203 = int32(0)
	v205 = v200 - int32(48)
	if base.Ui32(v205) <= base.Ui32(int32(9)) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v199 = int32(*(*int8)(unsafe.Add(mBase, uint32(v189))))
	v200 = v199
	v201 = v189
	v202 = v198
	goto L92
L94:
	;
	v198 = int32(0)
	goto L93
L95:
	;
	v208 = v203
	v209 = v205
	v210 = v201
	goto L98
L96:
	;
	v222 = v203
	goto L97
L97:
	;
	if v202 != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v212 = int32(10)
	v214 = v208*v212 - v209
	v215 = int32(*(*int8)(unsafe.Add(mBase, uint32(v210)+1)))
	v219 = v215 - int32(48)
	if base.Ui32(v219) < base.Ui32(v212) {
		v208 = v214
		v209 = v219
		v210 = v210 + int32(1)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v222 = v214
	goto L97
L100:
	;
	goto L99
L101:
	;
	v228 = int32(0) - v222
	goto L103
L102:
	;
	v228 = v222
	goto L103
L103:
	;
	goto L88
L104:
	;
	goto L22
L105:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[4])) = uint8(v234)
	goto L22
L106:
	;
	goto L22
L107:
	;
	goto L22
L108:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v247&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_process_postgres_switches[5])))
	F_SetConfigOption(m, v263, int32(_a_F_process_postgres_switches_20), l2, v55)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L50
	} else {
		goto L109
	}
L109:
	;
	goto L22
L110:
	;
	goto L22
L111:
	;
	goto L22
L112:
	;
	v278 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_postgres_switches[6])) = uint8(v278)
	goto L22
L113:
	;
	goto L22
L114:
	;
	goto L22
L115:
	;
	goto L22
L116:
	;
	goto L22
L117:
	;
	goto L22
L118:
	;
	goto L22
L119:
	;
	v306 = int32(_a_F_process_postgres_switches_32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	goto L123
L120:
	;
	goto L22
L121:
	;
	v424 = F_strlen(m, v413)
	mBase = m.M
	goto L120
L123:
	;
	goto L124
L124:
	;
	v314 = int32(1023)
	if (v306^v307)&int32(3) != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v414))) = uint8(v417)
	goto L121
L126:
	;
	v398 = v393
	v399 = v394
	v400 = v395
	goto L147
L127:
	;
	if v388 == int32(0) {
		v413 = v386
		v414 = v387
		goto L125
	} else {
		goto L146
	}
L128:
	;
	v386 = v307
	v387 = v306
	v388 = v314
	goto L127
L129:
	;
	goto L130
L130:
	;
	v318 = int32(0)
	if base.B2i32(v307&int32(3) == v318)|int32(0) == v318 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v354 == int32(0) {
		v413 = v351
		v414 = v352
		goto L125
	} else {
		goto L140
	}
L132:
	;
	v330 = v307
	v331 = v306
	v332 = v314
	goto L135
L133:
	;
	goto L134
L134:
	;
	v351 = v307
	v352 = v306
	v353 = v314
	v354 = int32(1)
	goto L131
L135:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v334)
	if v334 == int32(0) {
		v393 = v330
		v394 = v331
		v395 = v332
		goto L126
	} else {
		goto L137
	}
L136:
	;
	v351 = v345
	v352 = v339
	v353 = v341
	v354 = v343
	goto L131
L137:
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
		goto L131
	} else {
		goto L138
	}
L138:
	;
	if v341 != 0 {
		v330 = v345
		v331 = v339
		v332 = v341
		goto L135
	} else {
		goto L139
	}
L139:
	;
	goto L136
L140:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	if base.B2i32(v357 == int32(0))|base.B2i32(base.Ui32(v353) < base.Ui32(int32(4))) != 0 {
		v386 = v351
		v387 = v352
		v388 = v353
		goto L127
	} else {
		goto L141
	}
L141:
	;
	v364 = v351
	v365 = v352
	v366 = v353
	goto L142
L142:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v372 = int32(-2139062144)
	if (int32(16843008)-v369|v369)&v372 != v372 {
		v393 = v364
		v394 = v365
		v395 = v366
		goto L126
	} else {
		goto L144
	}
L143:
	;
	v386 = v380
	v387 = v378
	v388 = v382
	goto L127
L144:
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
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v393 = v386
	v394 = v387
	v395 = v388
	goto L126
L147:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	*(*uint8)(unsafe.Add(mBase, uint32(v399))) = uint8(v402)
	if v402 == int32(0) {
		v413 = v398
		v414 = v399
		goto L125
	} else {
		goto L149
	}
L148:
	;
	v413 = v409
	v414 = v407
	goto L125
L149:
	;
	v406 = int32(1)
	v407 = v399 + v406
	v409 = v398 + v406
	v411 = v400 - v406
	if v411 != 0 {
		v398 = v409
		v399 = v407
		v400 = v411
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	goto L22
L152:
	;
	goto L22
L153:
	;
	F_SetConfigOption(m, v445, int32(_a_F_process_postgres_switches_8), l2, v55)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L50
	} else {
		goto L156
	}
L154:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+1)))
	switch v441 - int32(97) {
	case 0:
		v445 = int32(_a_F_process_postgres_switches_33)
		goto L153
	default:
		goto L21
	case 11:
		goto L155
	}
L155:
	;
	v445 = int32(_a_F_process_postgres_switches_34)
	goto L153
L156:
	;
	goto L22
L157:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v456 = v452
	goto L159
L158:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[7])) = v500
	goto L22
L159:
	;
	v461 = v456 + int32(1)
	v462 = int32(*(*int8)(unsafe.Add(mBase, uint32(v456))))
	v463 = F___isspace(m, v462)
	mBase = m.M
	if v463 != 0 {
		v456 = v461
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v464 = int32(1)
	switch v462&int32(255) - int32(43) {
	case 0:
		v470 = v464
		goto L163
	default:
		v472 = v462
		v473 = v456
		v474 = v464
		goto L162
	case 2:
		goto L164
	}
L161:
	;
	goto L160
L162:
	;
	v475 = int32(0)
	v477 = v472 - int32(48)
	if base.Ui32(v477) <= base.Ui32(int32(9)) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v471 = int32(*(*int8)(unsafe.Add(mBase, uint32(v461))))
	v472 = v471
	v473 = v461
	v474 = v470
	goto L162
L164:
	;
	v470 = int32(0)
	goto L163
L165:
	;
	v480 = v475
	v481 = v477
	v482 = v473
	goto L168
L166:
	;
	v494 = v475
	goto L167
L167:
	;
	if v474 != 0 {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v484 = int32(10)
	v486 = v480*v484 - v481
	v487 = int32(*(*int8)(unsafe.Add(mBase, uint32(v482)+1)))
	v491 = v487 - int32(48)
	if base.Ui32(v491) < base.Ui32(v484) {
		v480 = v486
		v481 = v491
		v482 = v482 + int32(1)
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v494 = v486
	goto L167
L170:
	;
	goto L169
L171:
	;
	v500 = int32(0) - v494
	goto L173
L172:
	;
	v500 = v494
	goto L173
L173:
	;
	goto L158
L174:
	;
	goto L22
L175:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v53 != v533 {
		goto L20
	} else {
		goto L183
	}
L176:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v508 != 0 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v53-v509 <= int32(0) {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v513 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v509 + v513
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v54+v509<<(uint(int32(2))%32))))
	v522 = F_strlen(m, v519)
	mBase = m.M
	v524 = v522 + v513
	v525 = F_emscripten_builtin_malloc(m, v524)
	mBase = m.M
	if v525 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v530
	goto L175
L180:
	;
	v530 = int32(0)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v529 = F___memcpy(m, v525, v519, v524)
	mBase = m.M
	v530 = v529
	goto L179
L183:
	;
	m.G0 = v11 + int32(144)
	return
L184:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L50
	} else {
		goto L185
	}
L185:
	;
	if v548 == int32(1) {
		goto L17
	} else {
		goto L186
	}
L186:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v54+v558<<(uint(int32(2))%32))))
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v562
	F_errmsg(m, int32(_a_F_process_postgres_switches_35), v11+int32(48))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L50
	} else {
		goto L187
	}
L187:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v573
	F_errhint(m, int32(_a_F_process_postgres_switches_36), v11+int32(32))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L50
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(_a_F_process_postgres_switches_37), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L50
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L50
	} else {
		goto L191
	}
L191:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v592
	F_errmsg(m, int32(_a_F_process_postgres_switches_38), v11+int32(96))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L50
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(4051), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L50
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(4066), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L50
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_process_postgres_switches[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v627
	F_errhint(m, int32(_a_F_process_postgres_switches_36), v11)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L50
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_process_postgres_switches_30), int32(_a_F_process_postgres_switches_39), int32(_a_F_process_postgres_switches_31))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L50
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_procsignal_sigusr1_handler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[1]))
	v194 = int32(0)
	v197 = base.AtomicRmwOr32(m, v194, int32(_a_F_procsignal_sigusr1_handler_0), v194)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v198 != 0 {
		goto L61
	} else {
		goto L62
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+40))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+40)) = v8
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[2])) = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v14 == v8 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v17 = v4
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v17 = v14
	goto L5
L7:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v19
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[3])) = int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v25 == v19 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v28 = v17
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v28 = v25
	goto L9
L11:
	;
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v30
	v33 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4])) = v33
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[5])) = v33
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v39 == v30 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v42 = v28
	goto L13
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v42 = v39
	goto L13
L15:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[6]))
	if v47 == v44 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v61 = v42
	goto L17
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
	if v62 != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v58 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[7]))
	v53 = F_pgmem_kill(m, v51, int32(15))
	mBase = m.M
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[8])) = int32(1)
	goto L18
L22:
	;
	v61 = v58
	goto L17
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+56)) = int32(0)
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4])) = v66
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[9])) = v66
	goto L25
L24:
	;
	goto L25
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+60))
	if v71 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v72 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+60)) = v72
	v75 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4])) = v75
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[10])) = v75
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v81 == v72 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v84 = v61
	goto L28
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+64))
	if v85 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v84 = v81
	goto L28
L30:
	;
	v86 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+64)) = v86
	v89 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4])) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[11])) = v89
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v95 == v86 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	v98 = v84
	goto L32
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
	if v99 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v98 = v95
	goto L32
L34:
	;
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+72)) = v100
	v103 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4])) = v103
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[12])) = v103
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[1]))
	v113 = base.AtomicRmwOr32(m, v100, int32(_a_F_procsignal_sigusr1_handler_0), v100)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v114 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v165 = v98
	goto L36
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+68))
	if v166 != 0 {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v162 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(1)
	v117 = int32(0)
	v120 = base.AtomicRmwOr32(m, v117, int32(_a_F_procsignal_sigusr1_handler_0), v117)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v121 == v117 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v124 == int32(0) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[7]))
	if v128 == v124 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = m.G0
	v132 = v130 - int32(16)
	m.G0 = v132
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[13]))
	if v135 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v158 = F_pgmem_kill(m, v124, int32(23))
	mBase = m.M
	goto L38
L45:
	;
	m.G0 = v132 + int32(16)
	goto L37
L46:
	;
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v132)+15)) = uint8(v138)
	goto L47
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[14]))
	v146 = F_write(m, v142, v132+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v146 {
		goto L45
	} else {
		goto L49
	}
L48:
	;
	goto L45
L49:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[15]))
	if v150 == int32(27) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v165 = v162
	goto L36
L52:
	;
	v167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+68)) = v167
	v170 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4])) = v170
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[16])) = v170
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[0]))
	if v176 == v167 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v179 = v165
	goto L54
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+76))
	if v180 == int32(0) {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v179 = v176
	goto L54
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+76)) = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[17]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+340))
	if v187 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[4])) = int32(1)
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L1
L60:
	;
	return
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = int32(1)
	v201 = int32(0)
	v204 = base.AtomicRmwOr32(m, v201, int32(_a_F_procsignal_sigusr1_handler_0), v201)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v205 == v201 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	if v208 == int32(0) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[7]))
	if v212 == v208 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v214 = m.G0
	v216 = v214 - int32(16)
	m.G0 = v216
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[13]))
	if v219 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v242 = F_pgmem_kill(m, v208, int32(23))
	mBase = m.M
	goto L61
L68:
	;
	m.G0 = v216 + int32(16)
	goto L60
L69:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+15)) = uint8(v222)
	goto L70
L70:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[14]))
	v230 = F_write(m, v226, v216+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v230 {
		goto L68
	} else {
		goto L72
	}
L71:
	;
	goto L68
L72:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_procsignal_sigusr1_handler[15]))
	if v234 == int32(27) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
}
func F_prs_process_call(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 < v13 {
		v16 = v9 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v12<<(uint(int32(3))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v22
		v25 = F_pg_sprintf(m, v16, int32(_a_F_prs_process_call_0), v9)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(int32(3))%32))+4))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v39 = F_BuildTupleFromCStrings(m, v36, v9+int32(40))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				v42 = F_HeapTupleHeaderGetDatum(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
					F_pfree(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v47 + int32(1)
						v52 = v42
						m.G0 = v9 + int32(48)
						return v52
					}
				}
			}
		}
	} else {
		v52 = int64(0)
		m.G0 = v9 + int32(48)
		return v52
	}
}
func F_prs_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = F_lookup_ts_parser_cache(m, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v20
	v24 = int32(_a_F_prs_setup_firstcall_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_prs_setup_firstcall[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_prs_setup_firstcall[0])) = v27
	v30 = F_palloc(m, int32(12))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = int64(68719476736)
	v36 = F_palloc_mul(m, int32(8), int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v36
	v40 = v18 + int32(56)
	v41 = int32(0)
	v45 = int32(1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v49 = v47 & v45
	if v49 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v50 = v45
	goto L7
L6:
	;
	v50 = int32(4)
	goto L7
L7:
	;
	if v47 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v81 = F_FunctionCall2Coll(m, v18+int32(28), v41, base.I64_extend_i32_u(l3+v50), base.I64_extend_i32_s(v79))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v58 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v69 = int32(1)
	if v49 != 0 {
		v79 = int32(base.Ui32(v47)>>(uint(v69)%32)) - v69
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v61 = int32(16)
	goto L14
L13:
	;
	v61 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v68 = int32(4)
	goto L17
L16:
	;
	v68 = v61
	goto L17
L17:
	;
	v79 = v68
	goto L8
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	v84 = v81 & int64(4294967295)
	v87 = base.I64_extend_i32_u(v16 + int32(8))
	v90 = base.I64_extend_i32_u(v16 + int32(4))
	v91 = F_FunctionCall3Coll(m, v40, v41, v84, v87, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v93 = base.I32_wrap_i64(v91)
	if v93 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v97 = v93
	goto L24
L22:
	;
	goto L23
L23:
	;
	v179 = F_FunctionCall1Coll(m, v18+int32(84), int32(0), v84)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v107 <= v108 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v107 << (uint(int32(1)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v116 = F_repalloc(m, v113, v107<<(uint(int32(4))%32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v122 = F_palloc(m, v119+int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v116
	goto L28
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v124+v125<<(uint(int32(3))%32))+4)) = v122
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v130 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131+v132<<(uint(int32(3))%32))+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	base.MemoryCopy(m, v136, v137, v130)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v141 = int32(3)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139+v140<<(uint(v141)%32))+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144+v145))) = uint8(v147)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v149+v150<<(uint(v141)%32)))) = v97
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v155 + int32(1)
	v160 = F_FunctionCall3Coll(m, v40, v147, v84, v87, v90)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v162 = base.I32_wrap_i64(v160)
	if v162 != 0 {
		v97 = v162
		goto L24
	} else {
		goto L35
	}
L35:
	;
	goto L25
L36:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v181
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v183
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v30
	v189 = F_get_call_result_type(m, l1, v183, v16+int32(12))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v189 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v206
	v208 = F_TupleDescGetAttInMetadata(m, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	F_errmsg_internal(m, int32(_a_F_prs_setup_firstcall_1), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_prs_setup_firstcall_2), int32(209), int32(_a_F_prs_setup_firstcall_3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v208
	*(*int32)(unsafe.Add(mBase, _c_F_prs_setup_firstcall[0])) = v25
	m.G0 = v16 + int32(16)
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
	v215 = F_expression_tree_walker_impl(m, l0, int32(947), l1)
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
	v205 = F_query_tree_walker_impl(m, l0, int32(947), l1, int32(0))
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
	if v7 == int32(321) {
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+168))
	if base.Ui32(v45) <= base.Ui32(v44) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+164))
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
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
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
	if v59 != int32(321) {
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
	v221 = int32(base.Ui32(v219) >> (uint(int32(5)) % 32))
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
	v411 = int32(base.Ui32(v409) >> (uint(int32(5)) % 32))
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
	if v577 != int32(321) {
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
	v608 = int32(0)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v611 <= v608 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v614 = v599 + int32(8)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v621 = v618 & int32(-1)
	if v621 != 0 {
		v643 = v621
		v644 = v608
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v622 = int32(1)
	if v622 == v611 {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v626 = v622
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
	v689 = int32(base.Ui32(v687) >> (uint(int32(5)) % 32))
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
						*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v66
						*(*int32)(unsafe.Add(mBase, _c_F_push_old_value[3])) = l0 + int32(76)
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
							*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v66
							*(*int32)(unsafe.Add(mBase, _c_F_push_old_value[3])) = l0 + int32(76)
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
func F_pushdown_var_grouping_eqop(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v13 != 0 {
		v204 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v204
L2:
	;
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v14 <= int32(0) {
		v204 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v17 == int32(0) {
		v204 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 < v14 {
		v204 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+v14<<(uint(int32(2))%32)-int32(4))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	if v29 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	if v73 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v32 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v38 = int32(0)
	goto L9
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v36+v38<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v35 != v54 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	return v59
L11:
	;
	v57 = v38 + int32(1)
	if v57 != v32 {
		v38 = v57
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L6
L15:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	if v146 == int32(0) {
		v204 = v3
		goto L1
	} else {
		goto L30
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v76 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		v204 = v3
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v91 = v3
	goto L19
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v82+v91<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	if v99 == int32(0) {
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	return v132
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v102 <= int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v108 = int32(0)
	goto L23
L23:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v106+v108<<(uint(int32(2))%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v105 != v124 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v130 = v91 + int32(1)
	if v130 != v79 {
		v91 = v130
		goto L19
	} else {
		goto L29
	}
L25:
	;
	v127 = v108 + int32(1)
	if v127 != v102 {
		v108 = v127
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L15
L29:
	;
	goto L20
L30:
	;
	if v146 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v204 = v192
	goto L1
L32:
	;
	v192 = v186
	goto L31
L33:
	;
	v186 = int32(0)
	goto L32
L34:
	;
	v155 = v146
	goto L35
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v159 != int32(142) {
		goto L33
	} else {
		goto L37
	}
L36:
	;
	goto L33
L37:
	;
	if v14 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v177 = F_setop_column_grouping_eqop(m, v176, v14)
	mBase = m.M
	if v177 != 0 {
		v186 = v177
		goto L32
	} else {
		goto L42
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)+32))
	if v162 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v165 < v14 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167+v14<<(uint(int32(2))%32)-int32(4))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v192 = v174
	goto L31
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	if v178 != 0 {
		v155 = v178
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L36
}
