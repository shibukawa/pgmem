package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FindLogicalRepLocalIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
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
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
	if v12 == int32(112) {
		v216 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v216
L2:
	;
	v15 = F_RelationGetReplicaIndex(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v15 != 0 {
		v216 = v15
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = F_RelationGetPrimaryKeyIndex(m, l0, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v20 != 0 {
		v216 = v20
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v22 != int32(102) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	v28 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if v28 == int32(0) {
		v216 = int32(0)
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v32 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v38 = int32(0)
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v38<<(uint(int32(2))%32))))
	v54 = F_index_open(m, v52, int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	F_relation_close(m, v54, int32(1))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L48
	}
L19:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+196))
	v60 = F_heap_attisnull(m, v57, int32(21), v56)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L21
	}
L20:
	;
	v201 = v189
	goto L18
L21:
	;
	if v60 == int32(0) {
		v189 = v56
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)+196))
	v67 = F_SysCacheGetAttrNotNull(m, int32(34), v65, int32(18))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+192))
	v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+10)))
	if int32(0) < v70 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = int32(0)
	goto L27
L25:
	;
	v114 = v69
	goto L26
L26:
	;
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v115 < v117 {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(24)+v79<<(uint(int32(2))%32))))
	v90 = F_get_opclass_family(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L29
	}
L28:
	;
	v114 = v102
	goto L26
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+84))
	v96 = F_IndexAmTranslateCompareType(m, int32(3), v94, v90, int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	if v96 == int32(0) {
		v189 = v56
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v101 = v79 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v54)+192))
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+10)))
	if v101 < v103 {
		v79 = v101
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v123 = v115
	v126 = v117
	v127 = v116
	goto L37
L34:
	;
	v158 = v114
	goto L35
L35:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v158)+48)))
	if v159 == int32(0) {
		v189 = v56
		goto L20
	} else {
		goto L44
	}
L36:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v54)+192))
	v158 = v147
	goto L35
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v127+v126<<(uint(int32(4))%32)+v123*int32(100))+88))
	v138 = F_lookup_type_cache(m, v136, int32(32))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L39
	}
L38:
	;
	v201 = int32(0)
	goto L18
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+80))
	if v140 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v142 = v123 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v142 < v144 {
		v123 = v142
		v126 = v144
		v127 = v143
		goto L37
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L38
L43:
	;
	goto L36
L44:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v162 < v159 {
		v189 = v56
		goto L20
	} else {
		goto L45
	}
L45:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v164+v159<<(uint(int32(1))%32)-int32(2)))))
	if v170 < int32(0) {
		v189 = v56
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+84))
	v176 = F_GetIndexAmRoutineByAmId(m, v174, int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176)+100))
	v189 = base.B2i32(v178 != int32(0))
	goto L20
L48:
	;
	if v201 != 0 {
		v216 = v52
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v206 = v38 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v206 < v207 {
		v38 = v206
		goto L16
	} else {
		goto L50
	}
L50:
	;
	goto L17
}
func F_LogicalIncreaseRestartDecodingForSlot(m *base.Module, l0 int64, l1 int64) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v2 int64
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v1 = l0
	v2 = l1
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
	if v15 != 0 {
		F_s_lock(m, v14, int32(498761), int32(1757), int32(86429))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v14)+104))
			if base.Ui64(v2) <= base.Ui64(v23) {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
				m.G0 = v11 - int32(-64)
				return
			} else {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
				if base.Ui64(v1) <= base.Ui64(v27) {
					*(*int64)(unsafe.Add(mBase, uint32(v14)+256)) = v2
					*(*int64)(unsafe.Add(mBase, uint32(v14)+248)) = v1
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
					F_LogicalConfirmReceivedLocation(m, v27)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v11 - int32(-64)
						return
					}
				} else {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v14)+248))
					if v35 == int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v14)+256)) = v2
						*(*int64)(unsafe.Add(mBase, uint32(v14)+248)) = v1
						v40 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v40
						v44 = F_errstart(m, int32(14), v40)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							if v44 == int32(0) {
								m.G0 = v11 - int32(-64)
								return
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+12)) = uint32(v1)
								v49 = int64(32)
								v50 = int64(base.Ui64(v1) >> (uint(v49) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v50)
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v2)
								v54 = int64(base.Ui64(v2) >> (uint(v49) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v54)
								F_errmsg_internal(m, int32(514980), v11)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(498761), int32(1792), int32(86429))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										m.G0 = v11 - int32(-64)
										return
									}
								}
							}
						}
					} else {
						v64 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v64
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v14)+256))
						v69 = F_errstart(m, int32(14), v64)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							if v69 == int32(0) {
								m.G0 = v11 - int32(-64)
								return
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+52)) = uint32(v27)
								v74 = int64(32)
								v75 = int64(base.Ui64(v27) >> (uint(v74) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+48)) = uint32(v75)
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+44)) = uint32(v35)
								v79 = int64(base.Ui64(v35) >> (uint(v74) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+40)) = uint32(v79)
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+36)) = uint32(v66)
								v83 = int64(base.Ui64(v66) >> (uint(v74) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+32)) = uint32(v83)
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+28)) = uint32(v1)
								v87 = int64(base.Ui64(v1) >> (uint(v74) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+24)) = uint32(v87)
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v2)
								v91 = int64(base.Ui64(v2) >> (uint(v74) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v91)
								F_errmsg_internal(m, int32(515657), v9+int32(-48))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									F_errfinish(m, int32(498761), int32(1810), int32(86429))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return
									} else {
										m.G0 = v11 - int32(-64)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v14)+104))
		if base.Ui64(v2) <= base.Ui64(v23) {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
			m.G0 = v11 - int32(-64)
			return
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
			if base.Ui64(v1) <= base.Ui64(v27) {
				*(*int64)(unsafe.Add(mBase, uint32(v14)+256)) = v2
				*(*int64)(unsafe.Add(mBase, uint32(v14)+248)) = v1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
				F_LogicalConfirmReceivedLocation(m, v27)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					m.G0 = v11 - int32(-64)
					return
				}
			} else {
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v14)+248))
				if v35 == int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v14)+256)) = v2
					*(*int64)(unsafe.Add(mBase, uint32(v14)+248)) = v1
					v40 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v40
					v44 = F_errstart(m, int32(14), v40)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						if v44 == int32(0) {
							m.G0 = v11 - int32(-64)
							return
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+12)) = uint32(v1)
							v49 = int64(32)
							v50 = int64(base.Ui64(v1) >> (uint(v49) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v50)
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v2)
							v54 = int64(base.Ui64(v2) >> (uint(v49) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v54)
							F_errmsg_internal(m, int32(514980), v11)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(498761), int32(1792), int32(86429))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v11 - int32(-64)
									return
								}
							}
						}
					}
				} else {
					v64 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v64
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v14)+256))
					v69 = F_errstart(m, int32(14), v64)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						if v69 == int32(0) {
							m.G0 = v11 - int32(-64)
							return
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+52)) = uint32(v27)
							v74 = int64(32)
							v75 = int64(base.Ui64(v27) >> (uint(v74) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+48)) = uint32(v75)
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+44)) = uint32(v35)
							v79 = int64(base.Ui64(v35) >> (uint(v74) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+40)) = uint32(v79)
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+36)) = uint32(v66)
							v83 = int64(base.Ui64(v66) >> (uint(v74) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+32)) = uint32(v83)
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+28)) = uint32(v1)
							v87 = int64(base.Ui64(v1) >> (uint(v74) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+24)) = uint32(v87)
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v2)
							v91 = int64(base.Ui64(v2) >> (uint(v74) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v91)
							F_errmsg_internal(m, int32(515657), v9+int32(-48))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								F_errfinish(m, int32(498761), int32(1810), int32(86429))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									m.G0 = v11 - int32(-64)
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
func F_LogicalOutputPrepareWrite(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v7
	return
}
func F_LogicalRepWorkersWakeupAtCommit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = int32(4520272)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[669]))
	v11 = F_list_append_unique_oid(m, v10, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v4
		*(*int32)(unsafe.Add(mBase, _consts[669])) = v11
		return
	}
}
func F_LogicalTapeWrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = F_palloc(m, int32(8192))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v18 == int64(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v12
	goto L3
L6:
	;
	v21 = F_ltsGetBlock(m, v7, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v21
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1243]))) = int64(-1)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L33
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v33 = l1
	v34 = l2
	v35 = v31
	goto L14
L12:
	;
	goto L13
L13:
	;
	return
L14:
	;
	if int32(8176) <= v35 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	v64 = v35
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v69 = int32(8176) - v64
	if base.Ui32(v69) < base.Ui32(v34) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = F_ltsGetBlock(m, v43, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[1244]))) = v44
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ltsWriteBlock(m, v50, v51, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v55)+uint32(_consts[1243]))) = v58
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v44
	v64 = int32(0)
	goto L18
L22:
	;
	v71 = v69
	goto L24
L23:
	;
	v71 = v34
	goto L24
L24:
	;
	if v71 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v77 = v76 + v71
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v79 < v77 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v72 = F__emscripten_memcpy_bulkmem(m, v66+v64, v33, v71)
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v77
	goto L31
L30:
	;
	goto L31
L31:
	;
	v83 = v34 - v71
	if v83 != 0 {
		v33 = v33 + v71
		v34 = v83
		v35 = v77
		goto L14
	} else {
		goto L32
	}
L32:
	;
	goto L15
L33:
	;
	F_errmsg_internal(m, int32(8367), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(500012), int32(797), int32(351041))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logical_rewrite_log_mapping(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v10 = m.G0
	v12 = v10 - int32(1088)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1084)) = l1
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v23 = F_hash_search(m, v17, v12+int32(1084), int32(1), v12+int32(1083))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1083)))
		if v25 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+117)))
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1084))
			v35 = F_GetCurrentTransactionId(m)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v34
				*(*uint32)(unsafe.Add(mBase, uint32(v12)+32)) = uint32(v33)
				v41 = int64(base.Ui64(v33) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v12)+28)) = uint32(v41)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v16
				if v32 != 0 {
					v45 = int32(0)
				} else {
					v45 = v29
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v45
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(156630)
				v55 = F_pg_snprintf(m, v12+int32(48), int32(1024), int32(30130), v12+int32(16))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
					v60 = v23 + int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v60
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v60
					*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = int64(0)
					v70 = F__emscripten_memcpy_bulkmem(m, v23+int32(28), v12+int32(48), int32(1024))
					mBase = m.M
					v75 = F_PathNameOpenFile(m, v12+int32(48), int32(193))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v75
						if v75 < int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 + int32(48)
									F_errmsg(m, int32(300018), v12)
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return
									} else {
										F_errfinish(m, int32(497182), int32(977), int32(334885))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v87 = F_MemoryContextAlloc(m, v85, int32(44))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v89
								v91 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v87)+24)) = v91
								v93 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v93
								v95 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v95
								v97 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
								*(*int64)(unsafe.Add(mBase, uint32(v87))) = v97
								v100 = v23 + int32(16)
								v102 = v87 + int32(36)
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
								if v103 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v100
									*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v100
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v100
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+36)) = v111
								*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v102
								*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v102
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
								v116 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v115 + v116
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v121 = v119 + v116
								*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v121
								if base.Ui32(int32(1000)) <= base.Ui32(v121) {
									F_logical_heap_rewrite_flush_mappings(m, l0)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return
									} else {
										m.G0 = v12 + int32(1088)
										return
									}
								} else {
									m.G0 = v12 + int32(1088)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v87 = F_MemoryContextAlloc(m, v85, int32(44))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v89
				v91 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v87)+24)) = v91
				v93 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v93
				v95 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v95
				v97 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v87))) = v97
				v100 = v23 + int32(16)
				v102 = v87 + int32(36)
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
				if v103 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v100
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v100
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v100
				v111 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v87)+36)) = v111
				*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v102
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v102
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
				v116 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v115 + v116
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				v121 = v119 + v116
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v121
				if base.Ui32(int32(1000)) <= base.Ui32(v121) {
					F_logical_heap_rewrite_flush_mappings(m, l0)
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return
					} else {
						m.G0 = v12 + int32(1088)
						return
					}
				} else {
					m.G0 = v12 + int32(1088)
					return
				}
			}
		}
	}
}
